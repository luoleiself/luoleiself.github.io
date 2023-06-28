---
title: Docker-下
date: 2022-04-30 11:11:58
categories:
  - [tools]
  - [linux, Docker]
tags:
  - Docker
# description: docker 网络采用 veth-pair 技术, 每次启动容器时会自动创建一对虚拟网络设备接口, 一端连着网络协议栈, 一端彼此相连, 停止容器时自动删除, docker0 网卡作为中间的桥梁, 常见的网络模式包含 bridge, host, none, container, overlay等. Docker Compose 是定义和运行多容器 Docker 应用程序的工具, 运行部分命令时需要在 docker-compose yml 文件所在目录中, 以应用目录名_服务名_服务名数量编号为规则命名容器, 配置文件使用 yaml 语法, yaml 是一个可读性高，用来表达数据序列化的格式.
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

```shell
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

```shell
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

```shell
docker run ... --network host ...
```

#### none 模式

禁用网络功能, 不会为 docker 容器配置任何网络配置, 容器无法访问外部网络, 也无法被外部网络访问

```shell
docker run ... --network none ...
```

#### container 模式

新创建的容器不会创建自己的网卡和配置自己的 IP, 而是共享一个指定容器的 IP 和端口等

- \-\-network 指定容器运行的网络模式, 默认为 docker0

```shell
docker run ... centos02 --network centos01 ...
```

#### overlay 模式

将多个 docker 守护进程连接起来, 使 swarm 服务之间能够互相通信, 一般用于 swarm 集群

### 容器通信

#### 借助 docker0 路由功能

- 自定义网络 不适用 docker0
- 此方式需要查看容器的 ip 信息
- docker0 **不支持** 容器名 **别名** 连接访问

```shell
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

```shell
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

```shell
[root@localhost ~]# docker network create --subnet 192.168.0.0/16 --gateway 192.168.0.1 my-docker-net
[root@localhost ~]# docker network ls
NETWORK ID     NAME            DRIVER    SCOPE
2cad59b7f47d   bridge          bridge    local
1db6eddeb99f   host            host      local
b97686d7890c   my-docker-net   bridge    local
b2f23805b5e5   none            null      local
```

2. 查看宿主机 ip 相关信息 <!-- markdownlint-disable-line -->

```shell
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

```shell
[root@localhost ~]# docker run -tid --name my-docker-net01 --network my-docker-net centos /bin/bash
[root@localhost ~]# docker run -tid --name my-docker-net02 --network my-docker-net centos /bin/bash
```

4. 查看自定义网络信息 <!-- markdownlint-disable-line -->

```shell
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
      "Name": "my-docker-net01",
      "EndpointID": "e97619b1e847dc2a1e86fc9820ddfcd06a2e8d1d2685aa74d7f4bf850103bcfa",
      "MacAddress": "02:42:c0:a8:00:02",
      "IPv4Address": "192.168.0.2/16",
      "IPv6Address": ""
    },
    "d0f5bf60fb85ef35e3ca14c357fb2a6c0ab74c4d1c7679311a0afa8c41b3af79": {
      "Name": "my-docker-net02",
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

```shell
[root@localhost ~]# docker exec -it my-docker-net01 ping my-docker-net02
PING my-docker-net02 (192.168.0.3) 56(84) bytes of data.
64 bytes from my-docker-net02.my-docker-net (192.168.0.3): icmp_seq=1 ttl=64 time=0.110 ms
64 bytes from my-docker-net02.my-docker-net (192.168.0.3): icmp_seq=2 ttl=64 time=0.046 ms
64 bytes from my-docker-net02.my-docker-net (192.168.0.3): icmp_seq=3 ttl=64 time=0.046 ms
64 bytes from my-docker-net02.my-docker-net (192.168.0.3): icmp_seq=4 ttl=64 time=0.046 ms
^C
--- my-docker-net02 ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3043ms
rtt min/avg/max/mdev = 0.046/0.062/0.110/0.027 ms

[root@localhost ~]# docker exec -it my-docker-net02 ping my-docker-net01
PING my-docker-net01 (192.168.0.2) 56(84) bytes of data.
64 bytes from my-docker-net01.my-docker-net (192.168.0.2): icmp_seq=1 ttl=64 time=0.045 ms
64 bytes from my-docker-net01.my-docker-net (192.168.0.2): icmp_seq=2 ttl=64 time=0.095 ms
64 bytes from my-docker-net01.my-docker-net (192.168.0.2): icmp_seq=3 ttl=64 time=0.109 ms
64 bytes from my-docker-net01.my-docker-net (192.168.0.2): icmp_seq=4 ttl=64 time=0.047 ms
^C
--- my-docker-net01 ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3061ms
rtt min/avg/max/mdev = 0.045/0.074/0.109/0.028 ms
```

### 跨网络模式容器通信

> 自定义 bridge 和 docker0 结合使用

- my-docker-net01 和 my-docker-net02 运行在 my-docker-net 自定义网络模式下
- centos01 运行在 docker0 网络模式下

原理: 自定义网络模式分配 ip 信息给连接到此网络的容器

1. 查看自定义网络模式信息

```shell
[root@localhost ~]# docker network inspect my-docker-net
...
"Containers": {
  "ad9cdd7a0edf2c76710388fcc71de4df129b8fddd9d5e816795689022a62b141": {
    "Name": "my-docker-net01",
    "EndpointID": "52c0425c9b4274681f523e8a3ba7748dbc586ff80eafafa39897d577abb1edb6",
    "MacAddress": "02:42:c0:a8:00:02",
    "IPv4Address": "192.168.0.2/16",
    "IPv6Address": ""
  },
  "d0f5bf60fb85ef35e3ca14c357fb2a6c0ab74c4d1c7679311a0afa8c41b3af79": {
    "Name": "my-docker-net02",
    "EndpointID": "211d41485985263ca5131423b3784fec72a4923d21170212e4e311e569fe2d26",
    "MacAddress": "02:42:c0:a8:00:03",
    "IPv4Address": "192.168.0.3/16",
    "IPv6Address": ""
  }
}
...

# 创建基于 docker0 容器
[root@localhost ~]# docker run -tid --name centos01 centos
[root@localhost ~]# docker ps -a
CONTAINER ID   IMAGE     COMMAND       CREATED         STATUS        PORTS     NAMES
e2a9d9f19bd3   centos    "/bin/bash"   3 seconds ago   Up 1 second             centos01
d0f5bf60fb85   centos    "/bin/bash"   22 hours ago    Up 6 hours              my-docker-net02
ad9cdd7a0edf   centos    "/bin/bash"   22 hours ago    Up 6 hours              my-docker-net01
```

2. connect 命令连接容器到自定义网络 <!-- markdownlint-disable-line -->

```shell
docker network connect [OPTIONS] NETWORK CONTAINER
```

```shell
# 使用命令将不同网络模式中的容器加入到当前网络模式中
# 连接 centos01 到 自定义网络 my-docker-net
[root@localhost ~]# docker network connect my-docker-net centos01
# 查看 my-docker-net 自定义网络状态
[root@localhost ~]# docker network inspect my-docker-net
...
"Containers": {
  "ad9cdd7a0edf2c76710388fcc71de4df129b8fddd9d5e816795689022a62b141": {
    "Name": "my-docker-net01",
    "EndpointID": "52c0425c9b4274681f523e8a3ba7748dbc586ff80eafafa39897d577abb1edb6",
    "MacAddress": "02:42:c0:a8:00:02",
    "IPv4Address": "192.168.0.2/16",
    "IPv6Address": ""
  },
  "d0f5bf60fb85ef35e3ca14c357fb2a6c0ab74c4d1c7679311a0afa8c41b3af79": {
    "Name": "my-docker-net02",
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

3. 与自定义网络模式中的容器通信 <!-- markdownlint-disable-line -->

```shell
# 默认网络模式中容器 ping 自定义网络模式中容器
[root@localhost ~]# docker exec -it centos01 ping my-docker-net01
PING my-docker-net01 (192.168.0.2) 56(84) bytes of data.
64 bytes from my-docker-net01.my-docker-net (192.168.0.2): icmp_seq=1 ttl=64 time=0.060 ms
64 bytes from my-docker-net01.my-docker-net (192.168.0.2): icmp_seq=2 ttl=64 time=0.056 ms
^C
--- my-docker-net01 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 1002ms
rtt min/avg/max/mdev = 0.056/0.058/0.060/0.002 ms
[root@localhost ~]# docker exec -it centos01 ping my-docker-net02
PING my-docker-net02 (192.168.0.3) 56(84) bytes of data.
64 bytes from my-docker-net02.my-docker-net (192.168.0.3): icmp_seq=1 ttl=64 time=0.049 ms
64 bytes from my-docker-net02.my-docker-net (192.168.0.3): icmp_seq=2 ttl=64 time=0.115 ms
64 bytes from my-docker-net02.my-docker-net (192.168.0.3): icmp_seq=3 ttl=64 time=0.044 ms
^C
--- my-docker-net02 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2026ms
rtt min/avg/max/mdev = 0.044/0.069/0.115/0.033 ms

# 自定义网络模式中容器 ping 默认网络模式中容器
[root@localhost ~]# docker exec -it my-docker-net01 ping centos01
PING centos01 (192.168.0.4) 56(84) bytes of data.
64 bytes from centos01.my-docker-net (192.168.0.4): icmp_seq=1 ttl=64 time=0.066 ms
64 bytes from centos01.my-docker-net (192.168.0.4): icmp_seq=2 ttl=64 time=0.050 ms
^C
--- centos01 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 1027ms
rtt min/avg/max/mdev = 0.050/0.058/0.066/0.008 ms
[root@localhost ~]# docker exec -it my-docker-net02 ping centos01
PING centos01 (192.168.0.4) 56(84) bytes of data.
64 bytes from centos01.my-docker-net (192.168.0.4): icmp_seq=1 ttl=64 time=0.000 ms
64 bytes from centos01.my-docker-net (192.168.0.4): icmp_seq=2 ttl=64 time=0.048 ms
64 bytes from centos01.my-docker-net (192.168.0.4): icmp_seq=3 ttl=64 time=0.045 ms
^C
--- centos01 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2011ms
rtt min/avg/max/mdev = 0.000/0.031/0.048/0.021 ms
```

4. 断开容器到另一个网络的连接 <!-- markdownlint-disable-line -->

```shell
# 断开自定义网络和 centos01 的连接
[root@localhost ~]# docker network disconnect my-docker-net centos01
[root@localhost ~]# docker exec -it centos01 ping my-docker-net01
ping: my-docker-net01: Name or service not known
```

### 容器与外网互联

- 容器访问外网使用宿主机 `NAT` 转换 IP
- 外网访问容器使用 `docker proxy` 代理监听宿主机容端口映射容器端口

![docker-6](/images/docker-6.jpg)

### 跨主机容器通信

## 应用部署

### 部署 nginx

```shell
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

```shell
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

```shell
[root@localhost workspace]# docker ps -a
CONTAINER ID   IMAGE       COMMAND                  CREATED          STATUS             PORTS                                                  NAMES
6b6d19282ca8   nginx       "/docker-entrypoint.…"   27 minutes ago   Up 27 minutes      0.0.0.0:80->80/tcp, :::80->80/tcp                      c_nginx
c136f18229c3   mysql:5.7   "docker-entrypoint.s…"   15 hours ago     Up About an hour   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp   c_mysql
```

## Docker Compose

- Docker Compose 以服务为单位, 将为每一个服务部署一个容器
- 默认以 `应用目录名\_服务名\_数字` 方式作为容器名称
- 默认以 `应用目录名\_数据卷名` 方式作为数据卷名称
- 默认以 `应用目录名\_网络名` 方式作为网络名称

Docker Compose 是定义和运行多容器 Docker 应用程序的工具, 运行部分命令时需要在 `docker-compose.yml` 文件所在目录中, 以 `应用目录名\_服务名\_数字` 编号为规则命名容器, 配置文件使用 yaml 语法, yaml 是一个可读性高，用来表达数据序列化的格式.

yaml 文件中不能使用 tab 缩进, 只能使用空格

```shell
# 启动指定服务, 不加参数则默认启动所有服务
docker-compose -f -p -c --env-file up [service_name]

# 以下的命令不带服务名称则默认对所有服务执行相同操作
```

### 参数

- -f, \-\-file 指定配置文件
- -p, \-\-project-name 指定项目名称
- \-\-project-directory 指定项目工作目录
- -c, \-\-context 指定上下文环境名称
- \-\-env-file 指定环境变量配置文件

### 命令

- version 查看版本信息
- build 构建服务
- config 验证 docker-compose 配置文件
- cp 在容器和本地文件系统之间拷贝文件
- events 接收一个来自容器的真实的事件
- exec 在运行的容器中打开命令行
- down 停止并移除资源
- kill 关闭容器
- top 显示运行的进程信息
- images 查看所有镜像
- ls 列出正在运行的 compose 项目
- logs 查看容器日志
- ps 查看所有容器
- port 查看公共端口绑定信息
- pull 拉取服务镜像
- push 推送服务镜像
- start 启动服务
- stop 停止服务
- restart 重启服务
- rm 移除已经停止的容器
- run 运行命令
- pause 暂停服务
- unpasue 取消暂停服务
- create 创建容器, deprecated, Use the `up` command with `--no-start` instead

  - \-\-build 启动容器之前构建镜像
  - \-\-no-build 不构建镜像即使镜像不存在
  - \-\-force-recreate 即使配置项或镜像没有改变也要重新创建容器
  - \-\-no-recreate 如果容器存在则不创建新的容器
  - \-\-scale 调整服务实例数量, 并覆盖配置文件中的 scale 配置

- up 创建服务并启动容器

  - -f 指定配置文件
  - -d, \-\-detach 后台运行容器
  - \-\-attach 连接服务的输出
  - \-\-no-attach 不连接服务的输出
  - \-\-build 启动容器之前构建镜像
  - \-\-no-build 不构建镜像即使镜像不存在
  - \-\-no-start 创建服务之后不启动它
  - \-\-no-deps 不启动关联的服务
  - \-\-scale 调整服务实例数量, 覆盖配置文件中的 scale 配置

```shell
[root@localhost ~]# docker-compose up service_id # 启动指定服务

# 调整指定服务实例数量, 先去掉 docker-compose.yml 配置文件 service 指定的端口, 在单机中会出现端口占用问题
[root@localhost ~]# docker compose up --scale web=5 -d
```

### 配置文件

```yaml
# docker-compose.yml
version: 3.9   # 版本
services:
  web:   # 服务名称
    build: .
      context: './web'  # 指定构建 web 服务的镜像的上下文环境目录
      dockerfile: Dockerfile  # 指定构建镜像的配置文件名称
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
    container_name: my-web  # 容器名称
    env_file: .env # 环境变量配置文件
    environment:  # 设置容器内环境变量
      RACK_ENV: development
      SHOW: 'true'
      USER_INPUT:
    command: ['bundle', 'exec', 'thin', '-p', '3000']  # 覆盖镜像配置文件(Dockerfile)中的CMD指令
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
    networks:  # 自定义网络模式
      - my-web-network
    depends_on:  # 服务启动依赖
      - db
      - redis
    deploy:  # 部署
      # 外部客户端连接服务的方式
      # vip(Virtual IP) 为服务分配虚拟 IP, 客户端使用虚拟 IP 连接
      # dnsrr 平台配置 dns 条目, 使用服务名称查询 IP 地址列表连接
      endpoint_mode: vip
      mode: replicated # 服务运行模式, global | replicaated(default)
      replicas: 6 # 副本
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
  redis: # 服务名称
    image: redis
    volumes:   # 挂载数据卷
      - /home/workspace   # 定义匿名数据卷
    networks:   # 自定义网络模式
      - my-web-network
    links:    # 定义网络连接另一个服务的容器
      - db:mysql   # 可以直接使用 服务名, 或者使用 服务名:别名 方式
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
      - dbata:/var/lib/mysql   # 定义具名数据卷
    networks:
      - my-web-network
volumes:
  dbData:   # 声明卷名, compose 自动创建该卷名并会添加项目名前缀
  data:
    name: 'my-app-data'
networks:
  front-tier:
  back-tier:
  my-web-network:   # 声明自定义网络模式, compose 自动创建该网络并会添加项目名前缀
    driver: bridge
    enable_ipv6: true
    external: true
configs: # 允许服务调整其行为而无须重新构建 docker 镜像
  http_config:
    file: ./httpd.conf
secrets: # 针对敏感数据的配置
  server-certificate:
    file: ./server.cert
  token:
    environment: 'OAUTH_TOKEN'
```

#### version 支持的 Docker 引擎版本

#### volumes

- 使用路径方式挂载数据卷

- 使用卷名方式挂载数据卷, 需要在 `一级配置项` 中声明, compose 会自动创建以项目名为前缀的卷名, 如果不需要卷名前缀, 则使用 `external: true` 指定卷名, 但是需要手动创建该卷名

#### depends_on 服务启动依赖

- 当前服务启动的依赖优先于当前服务启动
- 当前服务关闭优先于当前服务的依赖关闭

## Docker Swarm

### 介绍

Swarm 是 Docker 官方提供的一款集群管理工具, 其主要作用是把若干台 Docker 主机抽象为一个整体, 并且通过一个入口统一管理这些 Docker 主机上的各种 Docker 资源

从集群角度来说, 一个 Swarm 由一个或多个 Docker 节点组成. 这些节点可以是物理服务器、虚拟机、树莓派(Raspberry Pi)或云实例. 唯一的前提就是要求所有节点通过可靠的网络相连

节点会被配置为管理节点(Manager)或工作节点(Worker). 管理节点负责集群控制面(Control Plane), 进行诸如监控集群状态、分发任务至工作节点等操作. 工作节点接收来自管理节点的任务并执行.

Swarm 的配置和状态信息保存在一套位于所有管理节点上的分布式 etcd 数据库中. 该数据库运行于内存中, 并保持数据的最新状态. 关于该数据库最棒的是, 它几乎不需要任何配置, 作为 Swarm 的一部分被安装, 无须管理

- Swarm 和 Kubernetes 比较类似, 但是更加轻, 具有的功能也较 kubernetes 更少一些
- Docker Swarm 包含两方面：一个企业级的 Docker 安全集群, 以及一个微服务应用编排引擎
- Swarm 默认内置有加密的分布式集群存储(encrypted distributed cluster store)、加密网络(Encrypted Network)、公用 TLS(Mutual TLS)、安全集群接入令牌 Secure Cluster Join Token)以及一套简化数字证书管理的 PKI(Public Key Infrastructure). 我们可以自如地添加或删除节点
- 编排方面, Swarm 提供了一套丰富的 API 使得部署和管理复杂的微服务应用变得易如反掌. 通过将应用定义在声明式配置文件中, 就可以使用原生的 Docker 命令完成部署

![docker-7](/images/docker-7.gif)

### 令牌格式

```shell
PREFIX - VERSION - SWARM ID - TOKEN

SWMTKN-1-5uqag7ddbx6jp9l273blxmda6308l5cn23487hbwsnw71w6dsh-eh4h7yhzchi0p6cy2ihg539jh
```

- PREFIX 令牌前缀,便于区分 固定为 SWMTKN
- VERSION Swarm 的版本信息
- SWARM ID Swarm 认证信息的一个哈希值
- TOKEN 标识管理节点还是工作节点的准入令牌

### 初始化 init

- \-\-advertise-addr 指定其他节点用来连接到当前管理节点的 IP 和端口, 当节点上有多个 IP 时指定
- \-\-listen-addr 指定用于承载 Swarm 流量的 IP 和端口. 其设置通常与 `--advertise-addr` 相匹配, 但是当节点上有多个 IP 的时候,可用于指定具体某个 IP
- \-\-autolock 启用锁

#### 开放端口

- 每个节点都需要安装 Docker, 并且能够与 Swarm 的其他节点通信

需要在路由器和防火墙中开放如下端口

- 2377/tcp: 用于客户端与 Swarm 进行安全通信
- 7946/tcp 与 7946/udp: 用于控制面 gossip 分发
- 4789/udp: 用于基于 VXLAN 的覆盖网络

```shell
[root@localhost ~]# docker swarm init --advertise-addr 192.168.1.2 --listen-addr 192.168.1.2
Swarm initialized: current node (5r1q8c5jaawi9w1wd8yr7w3u2) is now a manager.

To add a worker to this swarm, run the following command:

  docker swarm join --token SWMTKN-1-5uqag7ddbx6jp9l273blxmda6308l5cn23487hbwsnw71w6dsh-eh4h7yhzchi0p6cy2ihg539jh 192.168.1.2:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.
```

### 生成节点令牌

- join-token

#### 生成管理节点令牌

```shell
[root@localhost ~]# docker swarm join-token manager
To add a manager to this swarm, run the following command:

  docker swarm join --token SWMTKN-1-5uqag7ddbx6jp9l273blxmda6308l5cn23487hbwsnw71w6dsh-58yur8457jq0ghy45qnvislbi 192.168.1.2:2377
```

#### 生成工作节点令牌

```shell
[root@localhost ~]# docker swarm join-token worker
To add a worker to this swarm, run the following command:

  docker swarm join --token SWMTKN-1-5uqag7ddbx6jp9l273blxmda6308l5cn23487hbwsnw71w6dsh-eh4h7yhzchi0p6cy2ihg539jh 192.168.1.2:2377
```

### 更新节点令牌

- \-\-rotate

#### 更新管理节点令牌

```shell
[root@localhost ~]# docker swarm join-token --rotate manager
```

#### 更新工作节点令牌

```shell
[root@localhost ~]# docker swarm join-token --rotate worker
```

### 添加节点 join

- \-\-token

```shell
[root@localhost ~]# docker swarm join --token TOKEN HOST:PORT
```

#### 添加管理节点

```shell
[root@localhost ~]# docker swarm join --token SWMTKN-1-5uqag7ddbx6jp9l273blxmda6308l5cn23487hbwsnw71w6dsh-58yur8457jq0ghy45qnvislbi 192.168.1.2:2377
```

#### 添加工作节点

```shell
[root@localhost ~]# docker swarm join --token SWMTKN-1-5uqag7ddbx6jp9l273blxmda6308l5cn23487hbwsnw71w6dsh-eh4h7yhzchi0p6cy2ihg539jh 192.168.1.2:2377
```

### 移除节点

```shell
[root@localhost ~]# docker swarm leave
```
