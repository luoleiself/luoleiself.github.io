---
title: Docker-Swarm
date: 2022-04-30 11:11:58
categories:
  - [tools]
  - [linux, Docker]
tags:
  - Docker
---

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

```bash
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

```bash
[root@localhost ~]# docker swarm init --advertise-addr 192.168.1.2 --listen-addr 192.168.1.2
Swarm initialized: current node (5r1q8c5jaawi9w1wd8yr7w3u2) is now a manager.

To add a worker to this swarm, run the following command:

  docker swarm join --token SWMTKN-1-5uqag7ddbx6jp9l273blxmda6308l5cn23487hbwsnw71w6dsh-eh4h7yhzchi0p6cy2ihg539jh 192.168.1.2:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.
```

<!--more-->

### 生成节点令牌

- join-token

#### 生成管理节点令牌

```bash
[root@localhost ~]# docker swarm join-token manager
To add a manager to this swarm, run the following command:

  docker swarm join --token SWMTKN-1-5uqag7ddbx6jp9l273blxmda6308l5cn23487hbwsnw71w6dsh-58yur8457jq0ghy45qnvislbi 192.168.1.2:2377
```

#### 生成工作节点令牌

```bash
[root@localhost ~]# docker swarm join-token worker
To add a worker to this swarm, run the following command:

  docker swarm join --token SWMTKN-1-5uqag7ddbx6jp9l273blxmda6308l5cn23487hbwsnw71w6dsh-eh4h7yhzchi0p6cy2ihg539jh 192.168.1.2:2377
```

### 更新节点令牌

- \-\-rotate

#### 更新管理节点令牌

```bash
[root@localhost ~]# docker swarm join-token --rotate manager
```

#### 更新工作节点令牌

```bash
[root@localhost ~]# docker swarm join-token --rotate worker
```

### 添加节点 join

- \-\-token

```bash
[root@localhost ~]# docker swarm join --token TOKEN HOST:PORT
```

#### 添加管理节点

```bash
[root@localhost ~]# docker swarm join --token SWMTKN-1-5uqag7ddbx6jp9l273blxmda6308l5cn23487hbwsnw71w6dsh-58yur8457jq0ghy45qnvislbi 192.168.1.2:2377
```

#### 添加工作节点

```bash
[root@localhost ~]# docker swarm join --token SWMTKN-1-5uqag7ddbx6jp9l273blxmda6308l5cn23487hbwsnw71w6dsh-eh4h7yhzchi0p6cy2ihg539jh 192.168.1.2:2377
```

### 移除节点

```bash
[root@localhost ~]# docker swarm leave
```
