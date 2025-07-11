---
title: Docker-Compose
date: 2022-04-30 11:11:58
categories:
  - [tools]
  - [linux, Docker]
tags:
  - Docker
# description: docker 网络采用 veth-pair 技术, 每次启动容器时会自动创建一对虚拟网络设备接口, 一端连着网络协议栈, 一端彼此相连, 停止容器时自动删除, docker0 网卡作为中间的桥梁, 常见的网络模式包含 bridge, host, none, container, overlay等. Docker Compose 是定义和运行多容器 Docker 应用程序的工具, 运行部分命令时需要在 compose.yaml/docker-compose.yaml 文件所在目录中, 以应用目录名_服务名_服务名数量编号为规则命名容器, 配置文件使用 yaml 语法, yaml 是一个可读性高，用来表达数据序列化的格式.
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

## Docker Compose

- Docker Compose 以服务为单位, 将为每一个服务部署一个容器
- 默认以 `应用名称-服务名称-数字` 方式作为容器名称
- 默认以 `应用名称_数据卷名` 方式作为数据卷名称
- 默认以 `应用名称_网络名` 方式作为网络名称

Docker Compose 是定义和运行多容器 Docker 应用程序的工具, 运行部分命令时需要在 `compose.yaml/docker-compose.yaml` 文件所在目录中, 以 `应用名称-服务名称-数字` 编号为规则命名容器, 配置文件使用 yaml 语法, yaml 是一个可读性高，用来表达数据序列化的格式.

yaml 文件中不能使用 tab 缩进, 只能使用空格

```yaml
${VAR:-default}   如果 变量名 对应的环境变量未设置或为空, 则使用 默认值, 否则使用环境变量的值 
${VAR-default}    仅当 变量名 对应的环境变量 完全未定义 时才用默认值(为空时不会替换)
${VAR:?error}     严格模式, 如果 变量名 对应的环境变量 未设置/为空, 则报错并显示错误信息 error
${VAR:+alternate} 反向逻辑, 如果 变量名 已设置且非空, 则使用 alternate 值
```

```bash
# 启动指定服务, 不加参数则默认启动所有服务
docker compose -f -p -c --env-file .env.development up [service_name]

# 以下的命令不带服务名称则默认对所有服务执行相同操作
```

### 参数

- \-\-all-resources 引入所有的资源, 即使未被服务使用
- -f, \-\-file stringArray 指定配置文件

- -p, \-\-project-name string 指定项目名称
- \-\-project-directory string 指定项目工作目录

- -c, \-\-context 指定上下文环境名称
- \-\-env-file stringArray 指定环境变量配置文件
- \-\-parallel 设置并行
- \-\-compatibility 运行 compose 兼容模式

- \-\-profile 启用指定的服务, web 服务默认启动

```bash
docker compose --profile db up -d # 只启动 db, web 服务
```

### 命令

- attach 连接运行服务的标准输入输出
- build 构建服务
- commit 从服务容器创建一个新的镜像
  - -a, \-\-author 作者
  - -m, \-\-message string 提交信息
  - -c, \-\-change list 应用 Dockerfile 指令创建镜像
  - \-\-index int 指定如果有多个副本的服务的容器
  - -p, \-\-pause 提交过程中是否中断容器运行, 默认为 true
- config 解析验证 compose.yaml 配置文件
  - \-\-environment 打印环境变量的插值
  - \-\-format string 格式化输出, 值可选 yaml(default) | json
  - \-\-images 输出镜像名称

  - \-\-hash  输出服务的配置 hash, 一个一行
  - \-\-profiles  输出指定的服务名, 一个一行
  
  - -o, \-\-output string 保存到指定文件中, 默认是标准输出流
  - -q, \-\-quiet 仅验证配置项, 不输出任何信息
  - \-\-services 输出服务名称
  - \-\-volumes  输出数据卷名称
- cp 在容器和本地文件系统之间拷贝文件
- events 接收一个来自容器的真实的事件
- export 导出容器文件系统为归档文件
- images 列出创建容器的镜像

- start 启动服务
- restart 重启服务
- wait 阻塞直到所有的服务容器停止
- stop 停止服务
- kill 强制停止容器
- pause 暂停服务
- unpasue 取消暂停服务
- rm 移除已经停止的容器, 默认情况下附加到容器上的匿名数据卷不会被移除
  - -f, \-\-force 不询问确认操作直接移除
  - -s, \-\-stop 在移除之前停止容器
  - -v, \-\-volumes 移除所有附加到容器上的匿名数据卷
- down 停止并移除容器, 网络
  - \-\-rmi string 移除服务使用的镜像
  - -t, \-\-timeout int 延迟关机的时长秒
  - -v, \-\-volumes 移除在 compose.yaml 中 volumes 顶层指令中声明的具名和匿名数据卷

- logs 查看服务输出日志
  - \-f, \-\-follow  监听日志输出
  - \-\-index int  指定哪个容器执行命令, 如果服务有多个副本时
  - \-n, \-\-tail  输出容器日志的最后几行
  - \-\-since  显示指定时间开始的日志
  - \-\-until  显示截止到指定时间的日志
- stats 查看容器资源的使用情况

- ls 列出正在运行的 compose 项目
- port 查看公共端口绑定信息
  - \-\-index int 指定如果有多个副本的服务的容器
  - \-\-protocol string  指定协议, tcp(default) | udp
- ps 查看所有容器
  - -a, \-\-all 列出所有容器
  - \-f, \-\-filter \<FILTER_TYPE\>=\<VALUE\> 根据指定条件过滤服务容器
  - \-\-format string 使用自定义模板格式化输出, table(default) | table TEMPLATE | json | TEMPLATE
  - \-\-services 显示服务名称
  - \-\-status stringArray 通过状态过滤服务, [paused | restarting | removing | running | dead | created | exited]
- pull 拉取服务镜像
- push 推送服务镜像

- top 显示运行的进程信息
- exec 在运行的容器中执行命令
  - -d, \-\-detach  后台运行命令
  - -e, \-\-env stringArray 设置环境变量
  - \-\-index int  指定哪个容器执行命令, 如果服务有多个副本时
  - -T, \-\-no-TTY  不分配伪 TTY, 默认每次执行命令时都分配 TTY
  - -w, \-\-workdir string 设置本次命令的工作目录
- run 在服务上运行一次性命令
  - \-\-build 启动容器之前构建镜像
  - -d, \-\-detach
  - -i, \-\-interactive 交互式运行
  - -e, \-\-env stringArray 设置环境变量
  - -l, \-\-label stringArray 添加或覆盖 label
  - -P, \-\-service-ports 运行命令所有服务的端口都能映射到宿主机
  - \-\-name  给容器定义名字
  - \-\-pull  运行之前拉取镜像
  - -p, \-\-publish stringArray 发布容器的端口到宿主机
  - \-\-rm 当容器退出时自动移除
  - -v, \-\-volume stringArray 挂载数据卷
  - -w, \-\-workdir string 设置容器内的工作目录

- scale 调整服务
  - \-\-no-deps 不启动关联的服务

- version 查看版本信息
- watch 监听文件系统更新时服务容器重构/重启的构建上下文

- create 为服务创建容器

  - \-\-build 启动容器之前构建镜像
  - \-\-no-build 不构建镜像即使镜像不存在
  - \-\-force-recreate 即使配置项或镜像没有改变也要重新创建容器
  - \-\-no-recreate 如果容器存在则不创建新的容器
  - \-\-pull 创建之前拉取镜像
  - \-\-scale scale 调整服务实例数量, 并覆盖 `compose.yaml` 配置文件中的 scale 配置

- up 创建服务并启动容器

  - -d, \-\-detach 后台运行容器
  - \-\-attach 连接服务的输出
  - \-\-no-attach 不连接服务的输出
  - \-\-build 启动容器之前构建镜像
  - \-\-no-build 不构建镜像即使镜像不存在
  - \-\-no-start 创建服务之后不启动它
  - \-\-no-deps 不启动关联的服务
  - \-\-pull 启动之前拉取镜像
  - \-\-scale scale 调整服务实例数量, 覆盖 `compose.yaml` 配置文件中的 scale 配置
  - \-\-no-log-prefix 打印日志时不适用前缀
  - \-\-no-recreate 如果容器存在则不创建新的容器
  - -y 非交互式运行命令, 所有的提示都回答 yes

  - -t, \-\-timeout int 延迟关闭容器

```bash
[root@localhost ~]# docker compose up service_id # 启动指定服务

# 调整指定服务实例数量, 先去掉 compose.yaml/docker-compose.yaml 配置文件 service 指定的端口, 在单机中会出现端口占用问题
[root@localhost ~]# docker compose up --scale web=5 -d
```

### 配置文件

- 使用副本不能指定容器名称, Compose 自动使用 应用名称-服务名称-数字 形式命名容器

```yaml
# compose.yaml/docker-compose.yaml
version: 3.9   # 版本, obsolete(已过时)
name: myapp # 定义默认的项目名称, 将以环境变量 ${COMPOSE_PROJECT_NAME} 的方式公开
services:
  web:   # 服务名称
    annotations: # 容器声明, 可以是 arr 或 map
      com.example.foo: bar
      # - com.example.foo=bar
    attach: false # 设置为 false 时不会主动收集服务日志, 默认为 false, v2.20.0 以上支持
    build:
      context: './web'  # 指定构建 web 服务的镜像的上下文环境目录
      dockerfile: Dockerfile  # 指定构建镜像的配置文件名称
    command: ['bundle', 'exec', 'thin', '-p', '3000']  # 覆盖镜像配置文件(Dockerfile)中的CMD指令
    ports: # 端口映射
      - '5000:5000'
      - '0.0.0.0:80:80/tcp' # 指定 ip 地址和协议, 或修改 /etc/docker/daemon.json 配置项"ipv6":false
      - target: 80
        host_ip: 127.0.0.1
        published: 8080
        protocol: tcp
        mode: host
    privileged: true   # 配置容器目录权限
    read_only: true    # 开启容器文件系统只读模式
    restart: always    # 定义容器重启模式 "no" | always | on-failure | unless-stopped
    container_name: my-web  # 容器名称, 使用副本不能指定容器名称, Compose 自动使用 应用名称-服务名称-数字 形式命名容器
    env_file: .env # 环境变量配置文件
    environment:  # 设置容器内环境变量
      RACK_ENV: development
      SHOW: 'true'
      USER_INPUT:
    entrypoint:   # 覆盖镜像配置文件(Dockerfile)中的 ENTRYPOINT 指令
      - php
      - -d
      - vendor/bin/phpunit
    external_links:   # 将服务容器连接到 compose 应用管理以外的服务, 作用同 links
      - redis
      - database:mysql
      - database:postgresql
    extra_hosts:  # 添加主机 ip 映射关系到容器网络接口配置中(/etc/hosts)
      - 'somehost:162.242.195.82'
      - 'otherhost:50.31.209.229'
    volumes: # 挂载数据卷
      - type: volume
        source: db-data
        target: /data
        read_only: true
        volume:
          nocopy: true
          subpath: sub
        tempfs:
          size: 1024
          mode: 755
      - type: bind
        source: /home/workspace
        target: /home/workVolume
      - /home/workspace:/var/workspace  # 定义指定路径数据卷
    tmpfs:
      - /run  # 挂载容器内临时文件系统
      - /tmp
    volumes_from:     # 挂载共享数据卷
      - service_name
      - service_name:ro
      - container:container_name
      - container:container_name:rw
    network_mode: "host|none|service:[service name]" # 设置服务容器的网络模式
    networks:  # 自定义网络模式
      - my-web-network
    platform: linux/amd64 # 设置服务容器运行的目标平台
    # 当前服务启动的依赖优先于当前服务启动
    # 当前服务关闭优先于当前服务的依赖关闭
    depends_on:  # 服务启动依赖, 
      db:
        condition: service_healthy
        restart: true
      redis:
        condition: service_healthy
      # - db
      # - redis
    deploy:  # 部署
      # 外部客户端连接服务的方式
      # vip(Virtual IP) 为服务分配虚拟 IP, 客户端使用虚拟 IP 连接
      # dnsrr 平台配置 dns 条目, 使用服务名称查询 IP 地址列表连接
      endpoint_mode: vip
      labels:
        com.example.description: 'This label will appear on the web server' # 服务元数据
      mode: replicated # 服务运行模式, global | replicaated(default) | replicated-job | global-job
      replicas: 6 # 实例
      restart_policy: # 服务重启策略, 如果缺失, compose 会使用服务 restart 项
        condition: on-failure
        delay: 5s
        max_attempts: 3
        window: 120s
      rollback_config: # 服务回滚设置
      update_config: # 服务升级设置
        parallelism: 2
        delay: 10s
        order: stop-first
    dns:
      - 8.8.8.8  # 自定义网络的 DNS 服务器
    extends:
      file: common.yml   # 当前配置中扩展另一个服务
    labels:    # 添加容器元数据
      - 'com.example.description=Accounting webapp'
    develop: # 定义开发模式容器同步
      devices:
      dns:
      dns_opt:
      dns_search:
    healthcheck: # 服务健康检查
      test: ["CMD", "curl", "-f", "http://localhost"]
      interval: 1m30s
      timeout: 10s
      retries: 3
      start_period: 40s
      start_interval: 5s

  redis: # 服务名称 
    # 当同时使用 image 和 build 指令构建时, 将按照 pull_policy 的定义进行构建, 
    # 如果未定义 pull_policy 时, Compose 会尝试先拉取镜像, 如果在镜像仓库或缓存中找不到镜像时从源构建
    pull_policy: always   # Compose 总是从镜像仓库拉取镜像
                 never    # Compose 不会从镜像仓库拉取镜像而是依赖缓存中的镜像, 如果缓存中不存在则报告错误
                 missing  # 默认选项, Compose 仅当缓存中镜像不可用时从镜像仓库拉取
                 build    # Compose 构建镜像如果镜像已经存在则重新构建
    image: redis
    build:
      context: redis
      dockerfile: /redis.Dockerfile
    volumes:   # 挂载数据卷
      - /home/workspace   # 定义匿名数据卷
      - db-data:/var/lib/redis  # 挂载公共数据卷 db-data
    networks:   # 自定义网络模式
      - my-web-network
    links:    # 定义网络连接另一个服务的容器
      - db:mysql   # 可以直接使用 服务名, 或者使用 服务名:别名 方式
    scale: 6  # 设置容器数量, 如果 scale 和 deploy.replicas 同时存在则必须保持一致
    profiles: # 指定启动时的 profile
      - 'db'

  db:
    image: mysql
    ports:
      - '3306:3306'
    environment:
      MYSQL_ROOT_PASSWORD: 123456
      MYSQL_DATABASE: test
      MYSQL_USER: test
      MYSQL_PASSWORD: test123
    volumes:
      - db-data:/var/lib/mysql   # 挂载公共数据卷 db-data
    networks:
      - my-web-network
    profiles: # 指定启动时的 profile
      - 'debug'

# 跨服务共享数据卷定义到顶层指令 volumes
volumes:
  db-data:   # 声明卷名, compose 自动创建该卷名并会添加项目名前缀
  data:
    name: 'my-app-data'

networks:
  front-tier:
  back-tier:
  backend:
    driver: custom-driver
  my-web-network:   # 声明自定义网络模式, compose 自动创建该网络并会添加项目名前缀
    driver: bridge
    name: myapp-network # 定义网络名称, 在 docker network 列表中显示
    attachable: true  # 允许独立的容器连接到此网络
    enable_ipv6: true
    external: true  # 指定此网络的生命周期在应用程序的生命周期之外进行维护, Compose 不会尝试创建这些网络, 如果不存在则返回错误
configs: # 允许服务调整其行为而无须重新构建 docker 镜像
  http_config:
    file: ./httpd.conf
secrets: # 针对敏感数据的配置
  server-certificate:
    file: ./server.cert
  token:
    environment: 'OAUTH_TOKEN'
```

#### volumes

- 使用路径方式挂载数据卷

- 使用卷名方式挂载数据卷, 需要在 `一级配置项` 中声明, compose 会自动创建以项目名为前缀的卷名, 如果不需要卷名前缀, 则使用 `external: true` 指定卷名, 但是需要手动创建该卷名

### 多实例 Web 应用

```yaml
# compose.yaml
name: myapp
service:
  nginx:
    image: nginx:latest
    port:
      - '80:80'
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf  # 自定义 nginx 配置文件
    depends_on:
      - web
    network:
      - webnet
  web:
    build:
      context: web
      dockerfile: ./my-web-app.Dockerfile
    deploy:
      replicas: 3 # 启动 3 个 web 实例
    environment:
      - ENV=production
    network:
      - webnet
networks:
  webnet:
    driver: bridge
```

nginx

```conf
# nginx.conf
http {
  upstream web_backend {
    server web:80;
    server web:80;
    server web:80;
  }
  server {
    listen 80;
    location / {
      proxy_pass http://web_backend;

      proxy_set_header Host $host;
      proxy_set_header X-Reap-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
  }
}
```

web 应用

```python
from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello, World!'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=80)
```

```Dockerfile
# my-web-app.Dockerfile
# 使用官方 Python 镜像作为基础镜像
FROM python:3.9-slim
# 设置工作目录
WORKDIR /app

# 复制依赖文件并安装依赖
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# 复制应用代码
COPY . .
EXPOSE 80
CMD ["python", "app.py"]
```

### docker compose 数据库

```yaml
name: myapp # ${COMPOSE_PROJECT_NAME} 以环境变量形式访问项目名称
services:
  redis:
    image: redis:7
#   使用副本不能指定容器名称, Compose 自动使用 应用名称-服务名称-数字 形式命名容器
#   container_name: redis-container
    ports:
      - '6379:6379'
    command: ['redis-server', '--appendonly yes', '--logfile /data/redis.log']
    volumes:
      - /var/lib/redis:/data
      - /var/lib/redis/redis.conf:/usr/local/etc/redis/redis.conf
    networks:
      - my-app-network
#   deploy:
#     replicas: 3
#       labels:
#         com.myapp.redis.description: 'This label will appear on the redis server'
  mysql:
    image: mysql:latest
    container_name: mysql-container
    ports:
      - '3306:3306'
    environment:
      MYSQL_ROOT_PASSWORD: 123456
      MYSQL_DATABASE: test
      MYSQL_USER: test
      MYSQL_PASSWORD: test123
    volumes:
      - /var/lib/mysql:/var/lib/mysql
    networks:
      - my-app-network
  mongodb:
#   mongodb 6.0 以上 docker 镜像不再包含 mongo shell 工具. 只包含 mongod 数据库服务器
#   手动下载 mongosh 工具, 或者使用 mongodb 6.0 之前的版本
    image: mongo:latest
    container_name: mongodb-container
    ports:
      - '27017:27017'
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: 123456
      MONGODB_BIND_IP: 0.0.0.0
    command: ['mongod', '--logpath', '/var/log/mongodb/mongod.log']
    volumes:
      - /var/lib/mongodb:/data/db
      - /var/lib/mongodb/logs:/var/log/mongodb
    networks:
      - my-app-network
#   挂载临时文件系统添加初始化脚本安装 mongosh
    tmpfs:
      - /tmp
#   首次创建容器需要先执行 entrypoint 命令安装 mongosh, 然后再注释 entrypoint
#   entrypoint: ['/bin/sh', '-c', 'apt-get update && apt-get install -y wget && wget -qO- https://downloads.mongodb.com/compass/mongosh-1.5.4-linux-x64.tgz | tar -xz -C /usr/local/bin --strip-components 1 mongosh-1.5.4-linux-x64/bin/mongosh && exec docker-entrypoint.sh $$MONGO_INITDB_ROOT_USERNAME $$MONGO_INITDB_ROOT_PASSWORD']
# sqllite:
#   image: sqllite3:latest
#   command: ['sqllite3', '/data/database.db'] # 启动 sqllite 并指向数据库文件
#   volumes:
#     - /var/lib/sqllite:/data
#   networks:
#     - my-app-network
volumes:
  db-data:
    labels:
      - 'com.myapp.volumes.description=share data vaolume'
networks:
  my-app-network:
    driver: bridge
    name: myapp-network # 定义网络名称, 在 docker network 列表中显示
    attachable: true  # 允许独立的容器连接到此网络
    enable_ipv6: true
```
