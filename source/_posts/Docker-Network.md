---
title: Docker
date: 2022-03-22 18:37:24
categories:
  - [tools]
  - [linux, Docker]
tags:
  - Docker
# description: docker 网络采用 veth-pair 技术, 每次启动容器时会自动创建一对虚拟网络设备接口, 一端连着网络协议栈, 一端彼此相连, 停止容器时自动删除, docker0 网卡作为中间的桥梁, 常见的网络模式包含 bridge, host, none, container, overlay等. Docker Compose 是定义和运行多容器 Docker 应用程序的工具, 运行部分命令时需要在 compose.yaml/docker-compose.yaml 文件所在目录中, 以应用目录名_服务名_服务名数量编号为规则命名容器, 配置文件使用 yaml 语法, yaml 是一个可读性高，用来表达数据序列化的格式
---

## Docker 网络

Docker 网络架构源自一种叫作容器网络模型(CNM)的方案, 该方案是开源的并且支持插接式连接

Libnetwork 是 Docker 对 CNM 的一种实现, 提供了 Docker 核心网络架构的全部功能. 不同的驱动可以通过插拔的方式接入 Libnetwork 来提供定制化的网络拓扑

CNM 定义了 3 个基本要素：沙盒(Sandbox)、终端(Endpoint)和网络(Network)

- 沙盒是一个独立的网络栈, 其中包括以太网接口、端口、路由表以及 DNS 配置
- 终端就是虚拟网络接口。就像普通网络接口一样，终端主要职责是负责创建连接. 在 CNM 中, 终端负责将沙盒连接到网络
- 网络是 802.1d 网桥(类似大家熟知的交换机)的软件实现. 因此, 网络就是需要交互的终端的集合, 并且终端之间相互独立

docker 网络采用 veth-pair 技术, 每次启动容器时会自动创建一对虚拟网络设备接口, 一端连着网络协议栈, 一端彼此相连, 停止容器时自动删除, docker0 网卡作为中间的桥梁, 常见的网络模式包含 bridge, host, none, container, overlay 等.

<!-- more -->

```bash
[root@localhost ~]# docker run -tid --name centos01 centos /bin/bash
[root@localhost ~]# docker run -tid --name centos02 centos /bin/bash
[root@localhost ~]# ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:5f:bb:e6 brd ff:ff:ff:ff:ff:ff
    inet 10.0.2.15/24 brd 10.0.2.255 scope global eth0
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fe5f:bbe6/64 scope link
       valid_lft forever preferred_lft forever
3: docker0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:95:38:bb:e7 brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
    inet6 fe80::42:95ff:fe38:bbe7/64 scope link
       valid_lft forever preferred_lft forever
21: veth8d4dd5d@if20: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default
    link/ether 1a:9d:58:47:ed:41 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet6 fe80::189d:58ff:fe47:ed41/64 scope link
       valid_lft forever preferred_lft forever
25: veth60f8b26@if24: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default
    link/ether 9a:48:44:23:44:18 brd ff:ff:ff:ff:ff:ff link-netnsid 1
    inet6 fe80::9848:44ff:fe23:4418/64 scope link
       valid_lft forever preferred_lft forever
```

### veth-pair

veth-pair 就是一对的虚拟设备接口，和 tap/tun 设备不同的是，它都是成对出现的。一端连着协议栈，一端彼此相连着. 由于它的这个特性，常常被用于构建虚拟网络拓扑。例如连接两个不同的网络命名空间(netns)，连接 docker 容器，连接网桥(Bridge)等

![docker-2](/images/docker-2.png)

### [网络模式](https://docs.docker.com/network/)

- ls 查看网络列表
  - \-f, \-\-filter \<FILTER_TYPE\>=\<VALUE\> 根据指定条件过滤网络
    - driver 按网络模式过滤
    - id  按网络 id 过滤
    - label 按标签过滤
    - name  按网络名称过滤
    - scope 按网络作用域过滤
    - type 按网络类型过滤

```bash
[root@localhost ~]# docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
2cad59b7f47d   bridge    bridge    local
1db6eddeb99f   host      host      local
b2f23805b5e5   none      null      local
```

- bridge 桥接模式, 默认
- host 和宿主机共享网络
- none 不配置网络
- container 容器内网络连通(缺点较多)

![docker-4](/images/docker-4.png)

#### bridge 模式

Docker 容器默认使用的网络模式, Docker 为每个容器创建一个虚拟网桥, 并将每个容器连接到该虚拟网桥上, 容器之间通过虚拟网桥进行通信

- link 使用参数方式添加连接到另一个容器，不再推荐使用
  - 自定义网络 不适用 docker0
  - docker0 不支持容器名别名连接访问
- network 自定义网络模式

桥接模式一般使用于单机模式

- 默认的桥接网络模式 `docker0` 不支持通过 Docker DNS 服务进行域名解析
- 自定义的桥接网络模式可以支持, 具体例子可见下方的 [自定义网络模式](#zidingyiwangluomoshi)

#### host 模式

容器不会获得一个独立的 Network Namespace, 而是和宿主机共享 IP 地址和端口

```bash
docker run ... --network host ...
```

#### none 模式

禁用网络功能, 不会为 docker 容器配置任何网络配置, 容器无法访问外部网络, 也无法被外部网络访问

```bash
docker run ... --network none ...
```

#### container 模式

新创建的容器不会创建自己的网卡和配置自己的 IP, 而是共享一个指定容器的 IP 和端口等

- \-\-network 指定容器运行的网络模式, 默认为 docker0

```bash
docker run ... centos02 --network centos01 ...
```

#### overlay 模式

将多个 docker 守护进程连接起来, 使 swarm 服务之间能够互相通信, 一般用于 swarm 集群

### 容器通信

#### 借助 docker0 路由功能

- 自定义网络 不适用 docker0
- 此方式需要查看容器的 ip 信息
- docker0 **不支持** 容器名 **别名** 连接访问

```bash
# 查看 centos01 的 ip 信息
[root@localhost ~]# docker exec -it centos01 ip addr
... 172.17.0.2
# 查看 centos02 的 ip 信息
[root@localhost ~]# docker exec -it centos02 ip addr
... 172.17.0.3

# centos01 ping centos02 只能通过 ip 测试
[root@localhost ~]# docker exec -it centos01 ping 172.17.0.3
PING 172.17.0.3 (172.17.0.3) 56(84) bytes of data.
64 bytes from 172.17.0.3: icmp_seq=1 ttl=64 time=0.065 ms
64 bytes from 172.17.0.3: icmp_seq=2 ttl=64 time=0.047 ms
64 bytes from 172.17.0.3: icmp_seq=3 ttl=64 time=0.041 ms
64 bytes from 172.17.0.3: icmp_seq=4 ttl=64 time=0.042 ms
--- 172.17.0.3 ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3081ms
rtt min/avg/max/mdev = 0.041/0.048/0.065/0.012 ms
```

#### \-\-link 参数方式

本质上是在容器内部 hosts 文件中添加 ip 映射, 只能单向通信

在未知对方容器 ip 信息的情况下不能通信

```bash
# 创建 centos02 连接到 centos01
[root@localhost ~]# docker run -tid --name centos02 --link centos01 centos /bin/bash
# centos02 ping centos01
[root@localhost ~]# docker exec -it centos02 ping centos01
PING centos01 (172.17.0.2) 56(84) bytes of data.
64 bytes from centos01 (172.17.0.2): icmp_seq=1 ttl=64 time=0.076 ms
64 bytes from centos01 (172.17.0.2): icmp_seq=2 ttl=64 time=0.044 ms
--- centos01 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 1002ms
rtt min/avg/max/mdev = 0.044/0.060/0.076/0.016 ms

# 查看 centos02 的 hosts
[root@localhost ~]# docker exec -it centos02 cat /etc/hosts
127.0.0.1       localhost
::1     localhost ip6-localhost ip6-loopback
fe00::0 ip6-localnet
ff00::0 ip6-mcastprefix
ff02::1 ip6-allnodes
ff02::2 ip6-allrouters
172.17.0.2      centos01 ab2b4ae51950
172.17.0.3      8f49dabd3c2c

# centos01 ping centos02 失败
[root@localhost ~]# docker exec -it centos01 ping centos02
ping: centos02: Name or service not known
```

### 自定义网络模式 <em id="zidingyiwangluomoshi"></em> <!-- markdownlint-disable-line -->

- ls 显示所有网络模式状态
- create 创建网络模式
- inspect 查看指定网络模式的信息
- rm 删除网络模式
- prune 删除所有未使用的网络模式
- connect 连接容器到另一个网络
- disconnect 断开容器到另一个网络的连接

#### overlay 模式 <!-- markdownlint-disable-line -->

- docker 运行在 swarm 模式
- 使用键值存储的 docker 主机集群

#### bridge 模式 <!-- markdownlint-disable-line -->

- \-\-driver 网络模式, 默认 bridge
- \-\-subnet CIDR 格式的子网网段
- \-\-gateway 子网的 IPV4 或者 IPV6 网关
- \-\-config-from 基于配置文件创建一个网络
- \-\-ip-range 设置自网络的 ip 范围
- \-\-label 设置网络的元数据

1. 创建自定义网络

```bash
[root@localhost ~]# docker network create --subnet 192.168.0.0/16 --gateway 192.168.0.1 my-docker-net
[root@localhost ~]# docker network ls
NETWORK ID     NAME            DRIVER    SCOPE
2cad59b7f47d   bridge          bridge    local
1db6eddeb99f   host            host      local
b97686d7890c   my-docker-net   bridge    local
b2f23805b5e5   none            null      local
```

2. 查看宿主机 ip 相关信息 <!-- markdownlint-disable-line -->

```bash
[root@localhost ~]# ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:d3:b5:cd brd ff:ff:ff:ff:ff:ff
    inet 10.0.2.15/24 brd 10.0.2.255 scope global noprefixroute dynamic eth0
       valid_lft 83646sec preferred_lft 83646sec
3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default
    link/ether 02:42:71:92:ec:8c brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
4: br-b97686d7890c: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default
    link/ether 02:42:58:8d:a4:42 brd ff:ff:ff:ff:ff:ff
    inet 192.168.0.1/16 brd 192.168.255.255 scope global br-b97686d7890c
       valid_lft forever preferred_lft forever
```

3. 基于自定义网络模式创建容器 <!-- markdownlint-disable-line -->

```bash
[root@localhost ~]# docker run -tid --name my-docker-net-centos-01 --network my-docker-net centos /bin/bash
[root@localhost ~]# docker run -tid --name my-docker-net-centos-02 --network my-docker-net centos /bin/bash
```

4. 查看自定义网络信息 <!-- markdownlint-disable-line -->

```bash
[root@localhost ~]# docker network inspect my-docker-net
[{
  "Name": "my-docker-net",
  "Id": "b97686d7890ce2da3ce1e53927a4f5c02f898f18e22d9c4766a5e761d4275568",
  "Created": "2022-04-08T10:52:46.474584772Z",
  "Scope": "local",
  "Driver": "bridge",
  "EnableIPv6": false,
  "IPAM": {
    "Driver": "default",
    "Options": {},
    "Config": [{ "Subnet": "192.168.0.0/16", "Gateway": "192.168.0.1" }]
  },
  "Internal": false,
  "Attachable": false,
  "Ingress": false,
  "ConfigFrom": {"Network": ""},
  "ConfigOnly": false,
  "Containers": {
    "ad9cdd7a0edf2c76710388fcc71de4df129b8fddd9d5e816795689022a62b141": {
      "Name": "my-docker-net-centos-01",
      "EndpointID": "e97619b1e847dc2a1e86fc9820ddfcd06a2e8d1d2685aa74d7f4bf850103bcfa",
      "MacAddress": "02:42:c0:a8:00:02",
      "IPv4Address": "192.168.0.2/16",
      "IPv6Address": ""
    },
    "d0f5bf60fb85ef35e3ca14c357fb2a6c0ab74c4d1c7679311a0afa8c41b3af79": {
      "Name": "my-docker-net-centos-02",
      "EndpointID": "1158ee5f27aaaac418cf63ab9e16a92d83d0fadb2af4c7943800628be8cf9ce3",
      "MacAddress": "02:42:c0:a8:00:03",
      "IPv4Address": "192.168.0.3/16",
      "IPv6Address": ""
    }
  },
  "Options": {}
}]
```

5. 自定义网络模式容器通信 <!-- markdownlint-disable-line -->

```bash
[root@localhost ~]# docker exec -it my-docker-net-centos-01 ping my-docker-net-centos-02
PING my-docker-net-centos-02 (192.168.0.3) 56(84) bytes of data.
64 bytes from my-docker-net-centos-02.my-docker-net (192.168.0.3): icmp_seq=1 ttl=64 time=0.110 ms
64 bytes from my-docker-net-centos-02.my-docker-net (192.168.0.3): icmp_seq=2 ttl=64 time=0.046 ms
64 bytes from my-docker-net-centos-02.my-docker-net (192.168.0.3): icmp_seq=3 ttl=64 time=0.046 ms
64 bytes from my-docker-net-centos-02.my-docker-net (192.168.0.3): icmp_seq=4 ttl=64 time=0.046 ms
^C
--- my-docker-net-centos-02 ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3043ms
rtt min/avg/max/mdev = 0.046/0.062/0.110/0.027 ms

[root@localhost ~]# docker exec -it my-docker-net-centos-02 ping my-docker-net-centos-01
PING my-docker-net-centos-01 (192.168.0.2) 56(84) bytes of data.
64 bytes from my-docker-net-centos-01.my-docker-net (192.168.0.2): icmp_seq=1 ttl=64 time=0.045 ms
64 bytes from my-docker-net-centos-01.my-docker-net (192.168.0.2): icmp_seq=2 ttl=64 time=0.095 ms
64 bytes from my-docker-net-centos-01.my-docker-net (192.168.0.2): icmp_seq=3 ttl=64 time=0.109 ms
64 bytes from my-docker-net-centos-01.my-docker-net (192.168.0.2): icmp_seq=4 ttl=64 time=0.047 ms
^C
--- my-docker-net-centos-01 ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3061ms
rtt min/avg/max/mdev = 0.045/0.074/0.109/0.028 ms
```

### 跨网络模式容器通信

> 自定义 bridge 和 docker0 结合使用

- my-docker-net-centos-01 和 my-docker-net-centos-02 运行在 my-docker-net 自定义网络模式下
- centos01 运行在 docker0 网络模式下

原理: 自定义网络模式分配 ip 信息给连接到此网络的容器

1. network connect 子命令连接容器到自定义网络 <!-- markdownlint-disable-line -->

```bash
docker network connect [OPTIONS] NETWORK CONTAINER
```

```bash
# 使用命令将容器连接到当前网络模式中
# 连接 centos01 到 自定义网络 my-docker-net
[root@localhost ~]# docker network connect my-docker-net centos01

# 查看 my-docker-net 自定义网络状态
[root@localhost ~]# docker network inspect my-docker-net
...
"Containers": {
  "ad9cdd7a0edf2c76710388fcc71de4df129b8fddd9d5e816795689022a62b141": {
    "Name": "my-docker-net-centos-01",
    "EndpointID": "52c0425c9b4274681f523e8a3ba7748dbc586ff80eafafa39897d577abb1edb6",
    "MacAddress": "02:42:c0:a8:00:02",
    "IPv4Address": "192.168.0.2/16",
    "IPv6Address": ""
  },
  "d0f5bf60fb85ef35e3ca14c357fb2a6c0ab74c4d1c7679311a0afa8c41b3af79": {
    "Name": "my-docker-net-centos-02",
    "EndpointID": "211d41485985263ca5131423b3784fec72a4923d21170212e4e311e569fe2d26",
    "MacAddress": "02:42:c0:a8:00:03",
    "IPv4Address": "192.168.0.3/16",
    "IPv6Address": ""
  },
  "e2a9d9f19bd318a5c3a9a488a170071b1ac68b3193d33118e691306189306f3e": {
    "Name": "centos01",
    "EndpointID": "249d9dad83b236b0d565ee00e94c27d4c7d0008d660478a251d251367c92e772",
    "MacAddress": "02:42:c0:a8:00:04",
    "IPv4Address": "192.168.0.4/16",
    "IPv6Address": ""
  }
}
...

# 查看 centos01 ip 信息
[root@localhost ~]# docker exec -it centos01 ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
      valid_lft forever preferred_lft forever
9: eth0@if10: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
      valid_lft forever preferred_lft forever
11: eth1@if12: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:c0:a8:00:04 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 192.168.0.4/16 brd 192.168.255.255 scope global eth1
      valid_lft forever preferred_lft forever
```

2. 与自定义网络模式中的容器通信 <!-- markdownlint-disable-line -->

```bash
# docker0 中容器 ping 自定义网络模式中容器
[root@localhost ~]# docker exec -it centos01 ping my-docker-net-centos-01
PING my-docker-net-centos-01 (192.168.0.2) 56(84) bytes of data.
64 bytes from my-docker-net-centos-01.my-docker-net (192.168.0.2): icmp_seq=1 ttl=64 time=0.060 ms
64 bytes from my-docker-net-centos-01.my-docker-net (192.168.0.2): icmp_seq=2 ttl=64 time=0.056 ms
^C
--- my-docker-net-centos-01 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 1002ms
rtt min/avg/max/mdev = 0.056/0.058/0.060/0.002 ms

[root@localhost ~]# docker exec -it centos01 ping my-docker-net-centos-02
PING my-docker-net-centos-02 (192.168.0.3) 56(84) bytes of data.
64 bytes from my-docker-net-centos-02.my-docker-net (192.168.0.3): icmp_seq=1 ttl=64 time=0.049 ms
64 bytes from my-docker-net-centos-02.my-docker-net (192.168.0.3): icmp_seq=2 ttl=64 time=0.115 ms
64 bytes from my-docker-net-centos-02.my-docker-net (192.168.0.3): icmp_seq=3 ttl=64 time=0.044 ms
^C
--- my-docker-net-centos-02 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2026ms
rtt min/avg/max/mdev = 0.044/0.069/0.115/0.033 ms

# 自定义网络模式中容器 ping docker0 中容器
[root@localhost ~]# docker exec -it my-docker-net-centos-01 ping centos01
PING centos01 (192.168.0.4) 56(84) bytes of data.
64 bytes from centos01.my-docker-net (192.168.0.4): icmp_seq=1 ttl=64 time=0.066 ms
64 bytes from centos01.my-docker-net (192.168.0.4): icmp_seq=2 ttl=64 time=0.050 ms
^C
--- centos01 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 1027ms
rtt min/avg/max/mdev = 0.050/0.058/0.066/0.008 ms

[root@localhost ~]# docker exec -it my-docker-net-centos-02 ping centos01
PING centos01 (192.168.0.4) 56(84) bytes of data.
64 bytes from centos01.my-docker-net (192.168.0.4): icmp_seq=1 ttl=64 time=0.000 ms
64 bytes from centos01.my-docker-net (192.168.0.4): icmp_seq=2 ttl=64 time=0.048 ms
64 bytes from centos01.my-docker-net (192.168.0.4): icmp_seq=3 ttl=64 time=0.045 ms
^C
--- centos01 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2011ms
rtt min/avg/max/mdev = 0.000/0.031/0.048/0.021 ms
```

3. network disconnect 子命令断开容器与其他网络模式的连接 <!-- markdownlint-disable-line -->

```bash
docker network disconnect [OPTIONS] NETWORK CONTAINER
```

```bash
# 断开自定义网络和 centos01 的连接
[root@localhost ~]# docker network disconnect my-docker-net centos01
[root@localhost ~]# docker exec -it centos01 ping my-docker-net-centos-01
ping: my-docker-net-centos-01: Name or service not known
```

### 容器与外网互联

- 容器访问外网使用宿主机 `NAT` 转换 IP
- 外网访问容器使用 `docker proxy` 代理监听宿主机容端口映射容器端口

![docker-6](/images/docker-6.jpg)

### 跨主机容器通信

## 应用部署

### 部署 nginx

```bash
[root@localhost workspace]#
# 创建宿主机挂载目录 nginx/conf.d nginx/log html

# 创建 nginx 配置文件 nginx/conf.d/default.conf
server {
  listen 80;
  server_name localhost;
  location /{
    root /usr/share/nginx/html;
    index index.html index.htm;
  }
}

# 创建 html/index.html 文件

[root@localhost workspace]# docker run -id --name c_nginx -p 80:80 \
> -v ${PWD}/nginx/conf.d:/etc/nginx/conf.d \
> -v ${PWD}/nginx/log:/var/log/nginx \
> -v ${PWD}/html:/usr/share/nginx/html nginx
```

### 部署 mysql

```bash
[root@localhost workspace]
# 创建宿主机挂载目录 mysql/db mysql/log mysql/conf

# 创建数据库配置文件 touch /mysql/conf/my.cnf
# my.cnf
[client]
default-character-set=utf8
[mysql]
default-character-set=utf8
[mysqld]
init_connect='SET collation_connection = utf8_unicode_ci'
init_connect='SET NAMES utf8'
character-set-server=utf8
collation-server=utf8_unicode_ci
skip-character-set-client-handshake
skip-name-resolve

[root@localhost workspace]# docker run -tid -p 3306:3306 --name c_mysql \
> -e MYSQL_ROOT_PASSWORD=123456 \
> -v ${PWD}/mysql/db:/var/lib/mysql \
> -v ${PWD}/mysql/log:/var/log/mysql \
> -v ${PWD}/mysql/conf:/etc/mysql \
> mysql:5.7
```

```mysql
show variables like '%char%'; # 查看字符集
```

```bash
[root@localhost workspace]# docker ps -a
CONTAINER ID   IMAGE       COMMAND                  CREATED          STATUS             PORTS                                                  NAMES
6b6d19282ca8   nginx       "/docker-entrypoint.…"   27 minutes ago   Up 27 minutes      0.0.0.0:80->80/tcp, :::80->80/tcp                      c_nginx
c136f18229c3   mysql:5.7   "docker-entrypoint.s…"   15 hours ago     Up About an hour   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp   c_mysql
```
