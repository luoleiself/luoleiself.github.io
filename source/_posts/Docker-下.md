---
title: Docker-下
date: 2022-04-30 11:11:58
categories:
  - tools
tags:
  - Docker
# description: docker 网络采用 veth-pair 技术, 每次启动容器时会自动创建一对虚拟网络设备接口, 一端连着网络协议栈, 一端彼此相连, 停止容器时自动删除, docker0 网卡作为中间的桥梁, 常见的网络模式包含 bridge, host, none, container, overlay等. Docker Compose 是定义和运行多容器 Docker 应用程序的工具, 运行部分命令时需要在 docker-compose yaml 文件所在目录中, 以应用目录名_服务名_服务名数量编号为规则命名容器, 配置文件使用 yaml 语法, yaml 是一个可读性高，用来表达数据序列化的格式.
---

## Docker 网络

docker 网络采用 veth-pair 技术, 每次启动容器时会自动创建一对虚拟网络设备接口, 一端连着网络协议栈, 一端彼此相连, 停止容器时自动删除, docker0 网卡作为中间的桥梁, 常见的网络模式包含 bridge, host, none, container, overlay 等.

- link 使用参数方式添加连接到另一个容器，不再推荐使用
  - 自定义网络 不适用 docker0
  - docker0 不支持容器名别名连接访问
- network 自定义网络模式

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

<!-- more -->

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

桥接模式一般使用于单机模式

- 默认的桥接网络模式 `docker0` 不支持通过 Docker DNS 服务进行域名解析
- 自定义的桥接网络模式可以支持, 具体例子可见下方的 [自定义网络模式](#zidingyiwangluomoshi)

#### host 模式

容器不会获得一个独立的 Network Namespace, 而是和宿主机共用一个 Network Namespace, 容器将不会虚拟出自己的网卡而是使用宿主机的 IP 和端口

```shell
docker run ... --network host ...
```

#### none 模式

禁用网络功能, 不会为 docker 容器配置任何网络配置

```shell
docker run ... --network none ...
```

#### container 模式

新创建的容器不会创建自己的网卡和配置自己的 IP, 而是共享一个指定容器的 IP 和端口等

- \-\-network 连接一个容器的网络

```shell
docker run ... centos02 --network centos01 ...
```

#### overlay 模式

将多个 docker 守护进程连接起来, 使 swarm 服务之间能够互相通信, 一般用于 swarm 集群

### bridge 容器互联通信

#### 借助 docker0 路由功能

此方式需要查看容器的 ip 信息

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

本质上是在容器内部 hosts 文件中添加 ip 映射, 可以单向使用容器别名通信

在未知容器 ip 信息的情况下不能使用别名通信

- 自定义网络 不适用 docker0
- docker0 不支持容器名别名连接访问

```shell
# 创建 centos02 连接到 centos01
[root@localhost ~]# docker run -tid --name centos02 --link centos01 centos /bin/bash
[root@localhost ~]# docker exec -it centos02 ping centos01 # centos02 ping centos01
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

### network 自定义网络模式 <em id="zidingyiwangluomoshi"></em>

- ls 显示所有网络模式状态
- create 创建网络模式
- inspect 查看指定网络模式的信息
- rm 删除网络模式
- prune 删除所有未使用的网络模式
- connect 连接容器到另一个网络
- disconnect 断开容器到另一个网络的连接

#### bridge 模式

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
8d4dbf9c3f56   my-docker-net   bridge    local
b2f23805b5e5   none            null      local
```

2. 创建基于自定义网络模式容器

```shell
[root@localhost ~]# docker run -tid --name my-docker-net01 --net my-docker-net centos /bin/bash
[root@localhost ~]# docker run -tid --name my-docker-net02 --net my-docker-net centos /bin/bash
```

3. 查看宿主机 ip 相关信息

```shell
[root@localhost ~]# ip addr
...
3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default
  link/ether 02:42:c2:ab:51:b5 brd ff:ff:ff:ff:ff:ff
  inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
    valid_lft forever preferred_lft forever
  inet6 fe80::42:c2ff:feab:51b5/64 scope link
    valid_lft forever preferred_lft forever
14: br-8d4dbf9c3f56: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
  link/ether 02:42:f5:dc:26:98 brd ff:ff:ff:ff:ff:ff
  inet 192.168.0.1/16 brd 192.168.255.255 scope global br-8d4dbf9c3f56
    valid_lft forever preferred_lft forever
  inet6 fe80::42:f5ff:fedc:2698/64 scope link
    valid_lft forever preferred_lft forever
16: veth3b7468a@if15: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master br-8d4dbf9c3f56 state UP group default
  link/ether 76:ba:b8:12:ae:25 brd ff:ff:ff:ff:ff:ff link-netnsid 0
  inet6 fe80::74ba:b8ff:fe12:ae25/64 scope link
    valid_lft forever preferred_lft forever
18: veth076162b@if17: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master br-8d4dbf9c3f56 state UP group default
  link/ether 92:5d:7d:23:d1:de brd ff:ff:ff:ff:ff:ff link-netnsid 1
  inet6 fe80::905d:7dff:fe23:d1de/64 scope link
    valid_lft forever preferred_lft forever
```

4. 查看自定义网络信息

```shell
[root@localhost ~]# docker network inspect my-docker-net
[{
  "Name": "my-docker-net",
  "Id": "8d4dbf9c3f5609303ef88f5e964ee47fd05b2ed27f6a8b2c49ddff0a9602e701",
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

5. 自定义网络内容器通信

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

#### overlay 模式

- docker 运行在 swarm 模式
- 使用键值存储的 docker 主机集群

### 跨网络模式容器通信

#### 自定义 bridge 和 docker0 结合使用

- my-docker-net01 和 my-docker-net02 运行在 my-docker-net 下
- centos01 运行在 docker0 下

```shell
docker network connect my-docker-net centos01 # 使用命令将不同网络模式中的容器加入到当前网络模式中
```

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

2. connect 命令连接容器到自定义网络

```shell
# 连接 centos01 到 自定义网络 my-docker-net
[root@localhost ~]# docker network connect my-docker-net centos01
# 查看自定义网络状态
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

3. 与自定义网络模式中的容器通信

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

4. 断开容器到另一个网络的连接

```shell
# 断开自定义网络和 centos01 的连接
[root@localhost ~]# docker network disconnect my-docker-net centos01
[root@localhost ~]# docker exec -it centos01 ping my-docker-net01
ping: my-docker-net01: Name or service not known
```

### 容器与外部网络互联

- 容器访问外网使用宿主机 `NAT` 转换 IP
- 外网访问容器使用 `docker proxy` 代理监听宿主机容端口映射容器端口

![docker-6](/images/docker-6.jpg)

## 应用

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

Docker Compose 是定义和运行多容器 Docker 应用程序的工具, 运行部分命令时需要在 docker-compose yaml 文件所在目录中, 以应用目录名*服务名*服务名数量编号为规则命名容器, 配置文件使用 yaml 语法, yaml 是一个可读性高，用来表达数据序列化的格式.

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
- create 创建服务, deprecated, Use the `up` command with `--no-start` instead
- events 接收一个来自容器的真实的事件
- exec 在运行的容器中打开命令行
- down 停止并移除资源
- up 创建服务并启动容器

  - -f 指定配置文件
  - \-\-build 启动容器之前构建镜像
  - -d, \-\-detach 后台运行容器
  - \-\-no-build 不构建镜像即使镜像不存在
  - \-\-no-start 创建服务之后不启动它

- kill 关闭容器
- top 显示运行的进程信息
- images 查看所有镜像
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
- scale 服务扩缩容
- pause 暂停服务
- unpasue 取消暂停服务

```shell
[root@localhost ~]# docker-compose up service_id # 启动指定服务
```

### 配置文件

```yaml
version: 3.9   # 版本
services:
  web:   # 服务名称
    build: .
      context: "./web"  # 指定构建 web 服务的镜像的上下文环境目录
      dockerfile: Dockerfile  # 指定构建镜像的配置文件名称
    ports: # 端口映射
      - '5000:5000'
    privileged: true   # 配置容器目录权限
    read_only: true    # 设计容器文件系统模式
    restart: always    # 定义容器重启模式
    container_name: my-web  # 容器名称
    environment:  # 环境变量
      RACK_ENV: development
      SHOW: 'true'
      USER_INPUT:
    env_file: .env # 环境变量配置文件
    command: [ "bundle", "exec", "thin", "-p", "3000" ]   # 覆盖镜像配置文件(Dockerfile)中的 CMD 指令
    entrypoint:    # 覆盖镜像配置文件(Dockerfile)中的 ENTRYPOINT 指令
      - php
      - -d
      - vendor/bin/phpunit
    volumes: # 挂载数据卷
      - type: bind
        source: /home/workspace
        target: /home/workVolume
      - /home/workspace:/var/workspace  # 定义指定路径数据卷
    tmpfs:
      - /run  # 挂载容器内临时文件系统
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
      replicas: 6 # 副本
    dns:
      - 8.8.8.8  # 自定义网络的 DNS 服务器
    extends:
      file: common.yml   # 当前配置中扩展另一个服务
    labels:    # 添加容器元数据
      - "com.example.description=Accounting webapp"
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
    volumes:
      - dbata:/var/lib/mysql   # 定义具名数据卷
    networks:
      - my-web-network
volumes:
  dbData:   # 声明卷名, compose 自动创建该卷名并会添加项目名前缀
    external:   # 使用自定义卷名
      true   # true 确定使用指定卷名, 该卷名需要手动创建, 否则 compose 会报错
networks:
  my-web-network:   # 声明自定义网络模式, compose 自动创建该网络并会添加项目名前缀
    external:
      true   # 作用同上方的数据卷的配置方式
external_links:   # 将服务容器连接到 compose 应用管理以外的服务, 作用同 links
  - database:mysql
  - database:postgresql
extra_hosts:  # 添加主机 ip 映射关系到容器网络接口配置中(/etc/hosts)
  - "somehost:162.242.195.82"
  - "otherhost:50.31.209.229
```

#### version 支持的 Docker 引擎版本

#### volumes

- 使用路径方式挂载数据卷

- 使用卷名方式挂载数据卷, 需要在 `一级配置项` 中声明, compose 会自动创建以项目名为前缀的卷名, 如果不需要卷名前缀, 则使用 `external: true` 指定卷名, 但是需要手动创建该卷名

#### depends_on 服务启动依赖

- 当前服务启动的依赖优先于当前服务启动
- 当前服务关闭优先于当前服务的依赖关闭

## Docker Swarm

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

- \-\-advertise-addr 设置管理节点入口地址
- \-\-listen-addr 监听地址

```shell
[root@localhost ~]# docker swarm init --advertise-addr 192.168.1.2
Swarm initialized: current node (5r1q8c5jaawi9w1wd8yr7w3u2) is now a manager.

To add a worker to this swarm, run the following command:

  docker swarm join --token SWMTKN-1-5uqag7ddbx6jp9l273blxmda6308l5cn23487hbwsnw71w6dsh-eh4h7yhzchi0p6cy2ihg539jh 192.168.1.2:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.
```

### 生成节点令牌 join-token

```shell
# 生成工作节点令牌
[root@localhost ~]# docker swarm join-token worker
To add a worker to this swarm, run the following command:

  docker swarm join --token SWMTKN-1-5uqag7ddbx6jp9l273blxmda6308l5cn23487hbwsnw71w6dsh-eh4h7yhzchi0p6cy2ihg539jh 192.168.1.2:2377

# 生成管理节点令牌
[root@localhost ~]# docker swarm join-token manager
To add a manager to this swarm, run the following command:

  docker swarm join --token SWMTKN-1-5uqag7ddbx6jp9l273blxmda6308l5cn23487hbwsnw71w6dsh-58yur8457jq0ghy45qnvislbi 192.168.1.2:2377
```

### 更新节点令牌

```shell
# 更新管理节点令牌
[root@localhost ~]# docker swarm join-token --rotate manager

# 更新工作节点令牌
[root@localhost ~]# docker swarm join-token --rotate worker
```

### 添加工作/管理节点 join

- \-\-token

```shell
[root@localhost ~]# docker swarm join --token TOKEN HOST:PORT
```

### 移除节点

```shell
[root@localhost ~]# docker swarm leave
```
