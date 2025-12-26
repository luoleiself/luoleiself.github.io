---
title: Redis-Cluster
date: 2022-11-09 18:03:19
categories:
  - [server, Redis]
tags:
  - Redis
---

## 主从复制

将一台 Redis 服务器的数据,复制到其他的 Redis 服务器. 前者称为主节点(Master/Leader),后者称为从节点(Slave/Follower), 数据的复制是单向的！只能由主节点复制到从节点(主节点以写为主、从节点以读为主)—— 读写分离.
===每台 Redis 服务器都是主节点===
一个主节点可以有 0 个或者多个从节点, 但每个从节点只能有一个主节点

```bash
127.0.0.1:6379> INFO replication # 当前副本的信息
# Replication
role:master
connected_slaves:0
master_failover_state:no-failover
master_replid:e49c9c650c72cd6e3f369365808da6de6efd3825
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:0
second_repl_offset:-1
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0
```

### 作用

- 数据冗余: 主从复制实现了数据的热备份, 是持久化之外的一种数据冗余的方式
- 故障恢复: 当主节点故障时, 从节点可以暂时替代主节点提供服务, 是一种服务冗余的方式
- 负载均衡: 在主从复制的基础上, 配合读写分离, 由主节点进行写操作, 从节点进行读操作, 分担节点的负载; 尤其是在多读少写的场景下, 通过多个从节点分担负载, 提高并发量
- 高可用(集群)基石: 主从复制还是哨兵和集群能够实施的基础

### 复制原理

> Redis 2.8 以上使用 PSYNC 命令完成同步

1. 从节点向主节点发送 `PSYNC` 命令, 如果从节点是首次连接主节点时会触发一次全量复制
2. 接到 `PSYNC` 命令的主节点会调用 `BGSAVE` 命令 fork 一个新线程创建 RDB 文件, 并使用缓冲区记录接下来执行的所有写命令
3. 当 RDB 文件生成完毕后, 主节点向所有从节点发送 RDB 文件, 并在发送期间继续记录被执行的写命令
4. 从节点接收到 RDB 文件后丢弃所有旧数据并载入这个文件
5. 主节点将缓冲区记录的所有写命令发送给从节点执行
6. 如果从节点断开连接后重连, 主节点仅将部分缺失的数据同步给从节点

- 全量复制: 从节点接收到数据库文件后, 将其全部加载到内存中
- 增量复制: 主节点将新的所有收集到的修改命令依次传给从节点, 完成同步

### 命令模式

===每台 Redis 服务器都是主节点===, 只用配置从服务器即可

**运行时有效**, 只在`本次服务器运行时有效`, 重启服务器后将会丢失配置信息

- 方式一: **启动** Redis 服务器时使用指定参数 `redis-server --port 6380 --replicaof 127.0.0.1 6379`

- 方式二: **连接** Redis 服务器使用内置命令 `REPLICAOF host port`

```bash
# 设置关联主服务器
127.0.0.1:6380> REPLICAOF 127.0.0.1 6379
OK
127.0.0.1:6381> REPLICAOF 127.0.0.1 6379
OK
# 查看主服务器配置信息
127.0.0.1:6379> INFO replication
# Replication
role:master
connected_slaves:2
slave0:ip=127.0.0.1,port=6380,state=online,offset=153689,lag=0
slave1:ip=127.0.0.1,port=6381,state=online,offset=153557,lag=0
master_failover_state:no-failover
master_replid:749aaed3f58b97f7c01d3732a6f6c55be205c4b2
master_replid2:451a270c3954af29f43878dd9bfeac579d011972
master_repl_offset:153689
second_repl_offset:133525
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:133525
repl_backlog_histlen:20165
```

```bash
# 主节点写入数据
127.0.0.1:6379> SET name helloworld
OK
# 从节点读取数据
127.0.0.1:6380> GET name
"helloworld"
127.0.0.1:6380> SET age 18 # 从机写入数据报错
(error) READONLY You can\'t write against a read only replica.
# 从节点读取数据
127.0.0.1:6381> GET name
"helloworld"
127.0.0.1:6381> SET age 18 # 从机写入数据报错
(error) READONLY You can\'t write against a read only replica.
```

提升从服务器角色

- `REPLICAOF NO ONE` 将从服务器更改为主服务器

### 配置文件模式

**永久有效**, 但是缺少可扩展性, 每次修改主从节点配置都需要重启 Redis 服务

redis.conf 基础配置，[集群配置](#redisclusterconfigure) <em id="redisbaseconfigure"></em> <!-- markdownlint-disable-line-->

```yaml
# 引入 redis 默认配置文件
include /root/redis-cluster/redis.conf
# 修改绑定 ip, 此处演示全为本机
bind 127.0.0.1
# 保护模式, 默认 yes, 只能允许本机连接
protected-mode no
# 修改 redis 端口号, 本机演示需要修改, 多机器时可以不用
port 6379
# 开启后台运行, 默认 no
daemonize yes
# 修改 redis 进程文件名
pidfile /var/run/redis_6379.pid

loglevel notice
# 修改日志文件名, 默认为空
# 守护进程模式将指定 /dev/null
logfile ""
# 修改持久化文件名, 默认为 dump.rdb
dbfilename dump.rdb

# 是否在未开启持久化模式下删除复制中使用的 RDB 文件, 默认 no
# rdb-del-sync-files no

dir "" # 工作目录, dbfilename, logfile, appenddirname 目录相对于此配置项

appendonly no # 是否开启 AOF
appendfilename "appendonly.aof"  # AOF 文件名
appenddirname "appendonlydir" # AOF 存储目录

# 配置主服务器 ip 和 port
replicaof <masterip> <masterport>

# 副本和主服务器同步时的认证密码, 如果主服务器开启验证
# masterauth <master-password>
# 副本和主服务器同步时的认证用户
# 如果主服务器使用 requirepass 配置项, 则必须配置此项
masteruser <username>

# 从节点只读模式, 默认 yes
# replica-read-only yes
# 不使用向磁盘写 rdb 文件通信的方式直接通过新建进程 socket 同步 rdb 文件, 默认 yes
# repl-diskless-sync yes
# 同步延迟, 默认 5 sec
# repl-diskless-sync-delay 5

# 主服务器发送 PING 指令到副本的平均时间间隔, 默认 10 sec
# repl-ping-replica-period 10

# 副本的超时时间, 默认 60 sec
# 确保此项值大于 repl-ping-replica-period 的值, 否则
# 每当主服务器和副本之间的流量较低时，都会检测到超时
# repl-timeout 60

# 如果设置为 yes, Redis 将使用更小的 tcp 包和更少的带宽同步数据
# 主从同步延迟取决于Linux 内核的配置默认 40 毫秒一次
# 是否关闭主从节点同步的无延迟, 默认 no
# repl-disable-tcp-nodelay no

# 哨兵模式下被选为主服务器的优先级, 值越小优先级越高, 默认 100
# replica-priority 100

# 默认情况哨兵模式下所有副本被包含在报告中
# 设置为 no 表示报告中不包含副本
# 但不影响被选举为 master 的优先级
# replica-announced yes

# 副本用于监听 master 连接副本的 ip 和 端口
# 可以被 master 自动检测到
# replica-announce-ip 5.5.5.5
# replica-announce-port 1234

# 支持存储最多的 key 的数量, 默认 1000000
# tracking-table-max-keys 1000000
# 支持同时最多连接的客户端数量, 默认 10000
# maxclients 10000

# 最大内容
# maxmemory <bytes>

# 副本忽略最大内存限制
# replica-ignore-maxmemory yes

# 从 Redis 6.0 开始作为新 ACL 系统之上的一个兼容配置
# 该选项将只是为默认用户设置密码
# 客户端仍需要使用 AUTH [username] password 进行身份认证
# requirepass foobared

# 主服务器关机时副本的最大等待时间, 默认 10 sec
# shutdown-timeout 10

# 高级配置
# 哈希类型元素的个数小于 512 个, 每个元素值都小于 64B 时
# 使用 压缩列表 作为底层数据结构
# 否则使用 哈希表 作为底层数据结构
# hash-max-listpack-entries 512
# hash-max-listpack-value 64

# -2: max size: 8kb
# list-max-listpack-size -2
# 0 表示禁用所有的列表压缩
# list-compress-depth 0

# 集合元素都是整数且元素个数小于 512 个使用 整数集合 作为底层数据结构
# 否则使用 哈希表 作为底层数据结构
# set-max-intset-entries 512
# set-max-listpack-entries 128
# set-max-listpack-value 64

# 有序集合的元素个数小于 128 个并且每个元素的值小于 64B 时
# 使用 压缩列表 作为底层数据结构
# 否则使用 跳表 作为数据结构
# zset-max-listpack-entries 128
# zset-max-listpack-value 64

# HyperLogLog
# hll-sparse-max-bytes 3000

# Streams
# stream-node-max-bytes 4096
# stream-node-max-entries 100

# 绑定 cpu
# server_cpulist 0,2,4,6
```

- requirepass 认证

```bash
# 第一种方式: 连接 redis 后使用内置命令 AUTH 命令认证
[root@centos7 workspace]# redis-cli
127.0.0.1:6379> KEYS *
(error) NOAUTH Authentication required.
127.0.0.1:6379> ACL WHOAMI
(error) NOAUTH Authentication required.
127.0.0.1:6379> AUTH 1006611
OK
127.0.0.1:6379> ACL WHOAMI
"default"

# 第二种方式: 连接 redis 时使用参数认证
[root@centos7 workspace]# redis-cli --user default --pass 1006611
Warning: Using a password with '-a' or '-u' option on the command line interface may not be safe.
127.0.0.1:6379> ACL WHOAMI
"default"
127.0.0.1:6379> KEYS *
1) "hash:zhang"
2) "age"
3) "name"
```

## 哨兵模式

哨兵模式是一种特殊的模式, Redis 提供了启动哨兵的工具命令, 哨兵是一个独立的进程运行

- 哨兵节点通过发送 `PING` 命令, 监控所有的主(从)节点的反馈运行状态
- 当哨兵节点监控到 master 掉线并且其它多个哨兵节点确认 master 掉线后, 开始选取 leader 启动故障转移操作执行切换 master, 然后通过发布订阅模式通知其他的从节点, 修改配置文件并关联新的主节点
- 当 master 重连之后, 哨兵节点自动将 master 节点修改为 slave 模式

- 不能水平扩容, 不能动态的增、删节点
- 高可用特性会受到主节点的内存的限制

### 执行任务

- 监控: 定期检查主节点和从节点的健康状态, 包括发送 `PING` 命令、检查返回结果和检测通信故障
- 自动故障转移: 当一个主节点不能正常工作时, Sentinel 会开始一次自动故障迁移操作, 它会将失效主节点的其中一个从节点升级为新的主节点, 并让失败的主节点的其他从节点改为复制新的主节点. 当客户端试图连接失效的主节点时, 集群也会向客户端返回新的主节点的地址, 使得集群可以使用新主节点代替失效节点
- 高可用性切换: 选举新的主节点后, 哨兵节点会自动将从节点切换为新的主节点, 并通知其它从节点更新复制目标
- 配置提供者: 当客户端连接到哨兵节点时, 哨兵节点可以根据 Redis 集群的配置信息, 将其重定向到正确的主节点

### 选举算法

首先, 领头 sentinel 根据从节点的信号反馈将从节点列表中失联的节点剔除, 按照从节点的优先级(replica-priority)进行排序并选择优先级最高的从节点, 如果有多个具有相同最高优先级的从节点, 那么, 领头 sentinel 将多个具有相同最高优先级的从节点按照复制偏移量(复制积压缓冲区中存储的写操作的字节占用累加, 主从节点进行 `PSYNC` 使用)进行排序并选择其中偏移量最大(偏移量最大保存的数据最新)的从节点, 如果有多个优先级最高, 复制偏移量最大的从节点, 那么 领头 sentinel 将按照从节点的运行 ID 进行排序并选择其中 ID 最小的从节点

replica-priority > replica-offset > run-ID

### 配置方式

- 方式一: **启动** Redis 服务器时使用指定参数 `redis-server /path/to/sentinel.conf --sentinel`

- 方式二: **配置文件** 指定配置项 `redis-sentinel /path/to/sentinel.conf`

sentinel.conf 配置文件

```yaml
protected-mode no # 保护模式, 默认 yes, 只能允许本机连接
port 26379 # 服务端口号
daemonize no # 是否后台运行模式
pidfile /var/run/redis-sentinel-26379.pid # 进程文件
loglevel notice   # 日志等级
logfile "" # 日志文件
dir /tmp # 工作目录

# 监测服务器配置, 数字表示确认主服务器宕机的票数
# sentinel monitor <master-name> <ip> <redis-port> <quorum>
sentinel monitor mymaster 127.0.0.1 6379 2

# 通过本地地址监听外部网络中的 redis，通常用在 NAT 环境中
# sentinel announce-ip <ip>
# sentinel announce-port <port>

# 设置用于与主节点和副本进行身份验证的密码
# sentinel auth-pass <master-name> <password>

# 配置哨兵节点的认证密码
# requirepass <password>

# 配置其他 sentinel 认证的用户, 如果没有配置 sentinel-user
# 将使用 default 用户 和 sentinel-pass 进行认证
# sentinel sentinel-user <username>
# sentinel sentinel-pass <password>

# 不可触达的超时时间, 默认 30 sec
# sentinel down-after-milliseconds <master-name> <milliseconds>
sentinel down-after-milliseconds mymaster 30000

# ACL 日志内存大小
acllog-max-len 128
# ACL 日志存储文件
# aclfile /etc/redis/sentinel-users.acl

# 当主服务器宕机时支持最大同时重配服务器的数量, 默认 1
# sentinel parallel-syncs <master-name> <numreplicas>
sentinel parallel-syncs mymaster 1

# 当服务器宕机后等待再次重启的时间, 默认 3 min
# sentinel failover-timeout <master-name> <milliseconds>
sentinel failover-timeout mymaster 180000

# 服务器唤起脚本文件
# sentinel notification-script <master-name> <script-path>
# sentinel notification-script mymaster /var/redis/notify.sh

# 拒绝脚本配置, 默认 yes
sentinel deny-scripts-reconfig yes
```

![redis-2](/images/redis-2.png)
![redis-3](/images/redis-3.png)

### 一主三从哨兵配置

- 3 个哨兵配置

```yaml
# sentinel_26379.conf
port 26379
pidfile /var/run/redis-sentinel-26379.pid
logfile "redis_26379.log"
dir /tmp
sentinel monitor myredis 127.0.0.1 6379 2

# sentinel_36379.conf
port 36379
pidfile /var/run/redis-sentinel-36379.pid
logfile "redis_36379.log"
dir /tmp
sentinel monitor myredis 127.0.0.1 6379 2

# sentinel_46379.conf
port 46379
pidfile /var/run/redis-sentinel-46379.pid
logfile "redis_46379.log"
dir /tmp
sentinel monitor myredis 127.0.0.1 6379 2
```

- 3 台 redis 服务器配置

```yaml
# redis_6379.conf
bind 127.0.0.1
port 6379
daemonize yes
pidfile /var/run/redis_6379.pid
logfile "redis_6379.log"   #  文件目录相对于 dir 配置项
dir /tmp
dbfilename dump6379.rdb

# redis_6380.conf
bind 127.0.0.1
port 6380
daemonize yes
pidfile /var/run/redis_6380.pid
logfile "redis_6380.log"   #  文件目录相对于 dir 配置项
dir /tmp
dbfilename dump6380.rdb
# 配置主服务器 ip 和 port
replicaof 127.0.0.1 6379

# redis_6381.conf
bind 127.0.0.1
port 6381
daemonize yes
pidfile /var/run/redis_6381.pid
logfile "redis_6381.log"   #  文件目录相对于 dir 配置项
dir /tmp
dbfilename dump6381.rdb
# 配置主服务器 ip 和 port
replicaof 127.0.0.1 6379
```

- 根据配置文件启动所有服务

```bash
[root@centos7 ~]# redis-server .config/redis_6379.conf # 启动 redis 服务器
[root@centos7 ~]# redis-server .config/redis_6380.conf # 启动 redis 服务器
[root@centos7 ~]# redis-server .config/redis_6381.conf # 启动 redis 服务器
[root@centos7 ~]# redis-sentinel .config/sentinel_26379.conf # 启动哨兵
[root@centos7 ~]# redis-sentinel .config/sentinel_36379.conf # 启动哨兵
[root@centos7 ~]# redis-sentinel .config/sentinel_46379.conf # 启动哨兵
```

## 集群模式

> Redis 3.0 支持

Redis Cluster 是一种服务器 Sharding(分片) 技术, Sharding 采用 slot 的概念, 一共分成 16384 个 slot, 对于每个进入 Redis 的键值对, 对 key 执行 CRC16(key) mod 16384 操作, 得到的结果就是对应的 slot.

Redis 集群中的每个 node 负责分摊这 16384 个 slot 中的一部分, 当动态增减 node 时, 需要将 16384 个 slot 再分配, slot 中的键值对也要迁移, 这一过程目前还处于半自动状态仍需要人工介入, 如果某个 node 发生故障, 则此 node 负责的 slot 也就失效, 整个集群将不能工作

官方推荐的方案是将 node 配置成主从结构, 即 1:n, 如果主节点失效, Redis Cluster 根据选举算法从 slave 节点中选择一个升级为主节点继续提供服务, 如果失效的主节点恢复正常后则作为新的主节点的从节点

### Cluster Slot

> 从心跳包的大小、网络带宽、心跳并发、压缩率鞥维度考虑, 16384 个插槽更具有优势且能满足业务需求

- 正常的心跳数据包携带节点的完整配置, 它能以幂等方式来更新配置. 如果采用 16384 个插槽, 占空间 2KB(16384/8); 如果采用 65536 个插槽,占空间 8KB(65536/8). 8KB 的心跳包看似不大, 比起 16384 个插槽, 头大小增加了 4 倍,ping 消息的消息头太大, 浪费带宽
- Redis Cluster 不太可能扩展到超过 1000 个主节点, 太多可能导致网络拥堵
- 16384 个插槽比较合适, 当集群扩展到 1000 个节点时, 也能确保每个主节点有足够的插槽

### 集群特点

- 数据自动分片: 集群自动将数据分布到不同的节点上, 实现数据的均衡存储和负载均衡
- 自动故障转移: 当主节点发生故障时, 集群会自动进行故障检测, 并将从节点升级为新的主节点, 以保证系统的可用性
- 内部通信协议: 集集群使用 `Gossip` 协议进行节点之间的通信和状态更新, 确保集群的一致性和高效性
- 客户端路由: 客户端可以通过集群提供的路由机制, 自动将请求发送到正确的节点上, 实现透明访问
- 负载均衡: 在 Redis 集群中, 数据和请求会自动分布到不同的节点上, 实现负载均衡, 这样可以避免单个节点过载, 提高系统的稳定性和性能
- 扩展性好: 通过使用 Redis 集群, 可以便利地扩展系统的容量和性能, 将数据和请求分布到多个节点上, 提高整体系统的吞吐量和承载能力
- 高可用性: 通过 Redis 集群, 可以将数据分布到多个节点上, 实现数据的冗余备份和容错能力, 当部分节点不可用时, 集群仍然可以继续提供服务, 保证系统的可用性

### 命令

使用 `redis-cli -c -p port` 命令接入集群节点

- redis-cli \-\-cluster help # 查看集群命令帮助信息
- redis-cli \-\-cluster create \<host1:port1\> ... \<hostN:portN\> # 创建指定 IP 和 Port 的服务器作为集群
  - \-\-cluster-config-file  \<file\>   # 集群配置文件
  - \-\-cluster-replicas \<num\> # 指定集群中主节点和从节点数量的比例, 1 表示 1:1
  - \-\-cluster-timeout \<ms\>   # 节点超时时间
  - \-\-cluster-yes  # 自动确认配置
  - \-a \<password\> # 设置密码
  - \-\-askpass   # 交互式输入密码
- redis-cli \-\-cluster add-node \<new_host:new_port\> \<existing_host:existing_port\> # 添加集群节点
  - \-\-cluster-slave # 添加为从节点
  - \-\-cluster-master-id \<id\> # 添加到指定节点 ID
- redis-cli \-\-cluster del-node \<host:port\> \<node_id\> # 删除集群节点
  - \-\-cluster-yes # 自动确认

- redis-cli \-\-cluster replicate \<host:port\> \<node_id\>  # 设置主节点的副本

- redis-cli \-\-cluster reshard \<host:port\> # 手动重新分配节点槽位
  - \-\-cluster-from \<arg\> # 已有节点 id, 多个 id 之间使用半角逗号分隔
  - \-\-cluster-to \<arg\> # 新节点 id
  - \-\-cluster-slots \<arg\> # 新节点的 hash 槽数量

- redis-cli \-\-cluster rebalance \<host:port\> # 自动重新分配节点
  - \-\-cluster-weight \<node1=w1...nodeN=wN\> # 分配节点权重
  - \-\-cluster-timeout \<arg\> # 节点超时时间
  - \-\-cluster-threshold \<arg\> # 节点阈值
  - \-\-cluster-use-empty-masters

- redis-cli \-\-cluster failover    # 手动故障转移

- redis-cli \-\-cluster import host:port # 导入指定节点
  - \-\-cluster-from \<arg\> # 从指定 id
  - \-\-cluster-from-user \<arg\> # 指定用户名
  - \-\-cluster-from-pass \<arg\> # 指定密码
- redis-cli \-\-cluster info \<host:port\> # 查看指定节点信息
- redis-cli \-\-cluster check \<host:port\> # 检查指定节点
- redis-cli \-\-cluster call \<host:port\> \<command\> [args...] # 集群节点执行指定命令
  - \-\-cluster-only-masters 所有主节点
  - \-\-cluster-only-replicas 所有副本节点

  ```bash
  # 在所有主节点上执行加载的命令
  [root@centos7 workspace]# redis-cli --cluster --cluster-only-masters call host:port FUNCTION LOAD ...
  ```

- redis-cli \-\-cluster set-timeout host:port milliseconds # 设置节点的超时时间
- redis-cli \-\-cluster backup host:port backup_directory # 备份节点数据到指定目录

### 集群部署

#### 集群命令

- CLUSTER HELP 在 Redis 命令行中查看所有集群操作命令

```bash
127.0.0.1:6380> CLUSTER HELP
```

- CLUSTER INFO
- CLUSTER SLOTS 返回集群中 hash 槽的详细信息, redis 7.0 开始使用 CLUSTER SHARDS 命令代替
- CLUSTER REPLICAS \<node-id\> 列出指定节点的所有副本节点的信息, 功能和 CLUSTER NODES 类似
- CLUSTER NODES
- CLUSTER REPLICATE \<node-id\> 配置当前节点为指定主节点的从节点
- CLUSTER KEYSLOT \<somekey\> 计算指定 key 所在的 hash 槽
- CLUSTER COUNTKEYSINSLOT \<slot\> 统计集群中 hash 槽中存储的 key 的数量
- CLUSTER FAILOVER 手动启动集群故障转移操作, 此命令只能发送给集群从节点
- CLUSTER FLUSHSLOTS 清空当前节点的所有插槽

#### 编辑配置文件 <em id="bjpzwj"></em> <!-- markdownlint-disable-line -->

创建 Redis 服务器配置文件, 引入默认配置文件并覆盖配置项, 开启集群模式
创建 `cluster6379.conf`, `cluster6380.conf`, `cluster6381.conf`, `cluster6382.conf`, `cluster6383.conf`, `cluster6384.conf` 6 个文件
修改其中的 bind, port, pidfile, cluster-enabled, cluster-config-file

集群配置, [基础配置](#redisbaseconfigure) <em id="redisclusterconfigure"></em> <!-- markdownlint-disable-line -->

```yaml
# # 引入 redis 默认配置文件
# include /root/redis-cluster/redis.conf

port  6379
appendonly yes
daemonize yes

# 开启集群模式
cluster-enabled yes
# 修改集群节点文件名, 默认在存储在当前目录下
cluster-config-file nodes-6379.conf
# 设置节点失联时间, 超过该时间集群自动切换主从节点, 默认 15000 milsec
cluster-node-timeout 15000

# 集群总线监听 TCP 连接的端口, 默认端口为客户端命令端口+10000(eg: 6379+10000)
# 每个集群节点都需要开放两个端口, 一个(6379)用于客户端的 TCP 连接,
# 另一个(16379)用于节点使用集群总线进行故障监测、配置更新、故障转移等
# cluster-port 0

# (cluster-node-timeout * cluster-replica-validity-factor) + repl-ping-replica-period
# 例如, 集群节点超时为 30 sec, 并且集群副本有效性因子为 10 sec,
# 并且假设默认的 repl-ping-replica-period 为 10 sec,
# 则如果副本无法与主机进行超过 310 sec的通话, 则副本将不会尝试故障转移
# 集群副本有效性因子, 默认 10 sec
# cluster-replica-validity-factor 10

# 允许使用较少的自动集群配置, 默认 yes
# cluster-allow-replica-migration yes

# 默认当某一插槽不可用时, 整个集群都挂掉
# no 表示仅该插槽不可用, 默认 yes
cluster-require-full-coverage yes

# 允许集群失效的情况下依然可以从节点中读取数据, 保证了高可用性
# 默认 no, 不允许
# cluster-allow-reads-when-down no
# 允许集群服务器宕机时发布/订阅, 默认 yes
# cluster-allow-pubsubshard-when-down yes

# 集群模式使用 hostname 进行节点间通信
# 设置为空字符串表示将删除 hostname 并同步给其他节点
# cluster-announce-hostname ""

# 指定集群模式连接节点的方式是使用 ip、hostname或unknown-endpoint 方式
# 如果设置为 hostname 但未设置 cluster-announce-hostname 将返回 ?
# cluster-preferred-endpoint-type ip

# cluster-announce-ip ip # 集群模式下节点的 ip
# cluster-announce-port 6379 # 集群模式下节点的端口
# cluster-announce-tls-port 0 # 集群模式下节点的安全端口
# cluster-announce-bus-port 16379
```

#### 启动 Redis 服务器

启动所有的 redis 服务器, 使用 `ps -ef | grep redis` 命令查看 redis 服务器进程
redis 进程后中括号中的 cluster 表示 redis 工作在集群模式下, 需要进一步配置 redis 的集群关系

```bash
[root@centos7 redis-cluster]# redis-server cluster_6379.conf
[root@centos7 redis-cluster]# redis-server cluster_6380.conf
[root@centos7 redis-cluster]# redis-server cluster_6381.conf
[root@centos7 redis-cluster]# redis-server cluster_6382.conf
[root@centos7 redis-cluster]# redis-server cluster_6383.conf
[root@centos7 redis-cluster]# redis-server cluster_6384.conf
[root@centos7 redis-cluster]# ps -ef | grep redis
root      3731     1  0 05:49 ?        00:00:00 redis-server 127.0.0.1:6379 [cluster]
root      3737     1  0 05:49 ?        00:00:00 redis-server 127.0.0.1:6380 [cluster]
root      3743     1  0 05:49 ?        00:00:00 redis-server 127.0.0.1:6381 [cluster]
root      3749     1  0 05:49 ?        00:00:00 redis-server 127.0.0.1:6382 [cluster]
root      3755     1  0 05:49 ?        00:00:00 redis-server 127.0.0.1:6383 [cluster]
root      3761     1  0 05:49 ?        00:00:00 redis-server 127.0.0.1:6384 [cluster]
```

#### 创建集群节点

使用 `redis-cli --cluster create --cluster-replicas arg hostN:portN` 命令创建集群节点, arg 参数表示集群主从节点的数量比例, 1 表示 1:1
创建过程中提示输入 `yes` 表示接受当前配置信息并写入指定文件中, 最后输出 `[OK] All 16384 slots covered.` 表示集群创建完成

```bash
[root@centos7 redis-cluster]# redis-cli --cluster create --cluster-replicas 1 \
> 127.0.0.1:6379 127.0.0.1:6380 127.0.0.1:6381 \
> 127.0.0.1:6382 127.0.0.1:6383 127.0.0.1:6384
>>> Performing hash slots allocation on 6 nodes...
Master[0] -> Slots 0 - 5460
Master[1] -> Slots 5461 - 10922
Master[2] -> Slots 10923 - 16383
Adding replica 127.0.0.1:6383 to 127.0.0.1:6379
Adding replica 127.0.0.1:6384 to 127.0.0.1:6380
Adding replica 127.0.0.1:6382 to 127.0.0.1:6381
>>> Trying to optimize slaves allocation for anti-affinity
[WARNING] Some slaves are in the same host as their master
M: 2b144f1d7bdb31000a519492be980c6634576462 127.0.0.1:6379
   slots:[0-5460] (5461 slots) master
M: a770892444fbbe4b7d9391b458ac04d6bcba26f0 127.0.0.1:6380
   slots:[5461-10922] (5462 slots) master
M: 76cb8ea9a5d6ba0fa43d31cfa4c33cea8442e07d 127.0.0.1:6381
   slots:[10923-16383] (5461 slots) master
S: 6c9823906baa11aba873a798cce3a3b3c95465f2 127.0.0.1:6382
   replicates 2b144f1d7bdb31000a519492be980c6634576462
S: 4a56b76a379da615b606a499ae475e986eda3efd 127.0.0.1:6383
   replicates a770892444fbbe4b7d9391b458ac04d6bcba26f0
S: eaf9833aa105e36b22f6330585a972239bab9f50 127.0.0.1:6384
   replicates 76cb8ea9a5d6ba0fa43d31cfa4c33cea8442e07d
Can I set the above configuration? (type 'yes' to accept): yes
>>> Nodes configuration updated
>>> Assign a different config epoch to each node
>>> Sending CLUSTER MEET messages to join the cluster
Waiting for the cluster to join
.
>>> Performing Cluster Check (using node 127.0.0.1:6379)
M: 2b144f1d7bdb31000a519492be980c6634576462 127.0.0.1:6379
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
S: 6c9823906baa11aba873a798cce3a3b3c95465f2 127.0.0.1:6382
   slots: (0 slots) slave
   replicates 2b144f1d7bdb31000a519492be980c6634576462
M: a770892444fbbe4b7d9391b458ac04d6bcba26f0 127.0.0.1:6380
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
S: eaf9833aa105e36b22f6330585a972239bab9f50 127.0.0.1:6384
   slots: (0 slots) slave
   replicates 76cb8ea9a5d6ba0fa43d31cfa4c33cea8442e07d
S: 4a56b76a379da615b606a499ae475e986eda3efd 127.0.0.1:6383
   slots: (0 slots) slave
   replicates a770892444fbbe4b7d9391b458ac04d6bcba26f0
M: 76cb8ea9a5d6ba0fa43d31cfa4c33cea8442e07d 127.0.0.1:6381
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
```

#### 连接节点服务器

使用 `redis-cli -c -p port` 命令接入集群节点

- -c 以集群模式接入

```bash
[root@centos7 redis-cluster]# redis-cli -c -p 6379
```

#### 查看节点信息

- 方式一: 命令行中使用 `redis-cli --cluster info host:port` 命令查看指定节点的信息

```bash
[root@centos7 redis-cluster]# redis-cli --cluster info 127.0.0.1:6380
127.0.0.1:6380 (a7708924...) -> 3 keys | 5462 slots | 1 slaves.
127.0.0.1:6382 (6c982390...) -> 0 keys | 5461 slots | 0 slaves.
127.0.0.1:6381 (76cb8ea9...) -> 0 keys | 5461 slots | 1 slaves.
[OK] 3 keys in 3 masters.
0.00 keys per slot on average.
```

- 方式二: 在 Redis 命令行中使用 `CLUSTER INFO|SLOTS|NODES` 查看节点信息

```bash
# 查看当前节点信息
127.0.0.1:6380> CLUSTER INFO
cluster_state:ok
cluster_slots_assigned:16384
cluster_slots_ok:16384
cluster_slots_pfail:0
cluster_slots_fail:0
cluster_known_nodes:6
cluster_size:3
cluster_current_epoch:7
cluster_my_epoch:2
cluster_stats_messages_ping_sent:468
cluster_stats_messages_pong_sent:473
cluster_stats_messages_meet_sent:1
cluster_stats_messages_auth-ack_sent:1
cluster_stats_messages_sent:943
cluster_stats_messages_ping_received:473
cluster_stats_messages_pong_received:469
cluster_stats_messages_fail_received:1
cluster_stats_messages_auth-req_received:1
cluster_stats_messages_received:944
total_cluster_links_buffer_limit_exceeded:0

# 查看所有插槽信息
127.0.0.1:6380> CLUSTER SLOTS
1) 1) (integer) 0
   2) (integer) 5460
   3) 1) "127.0.0.1"
      2) (integer) 6382
      3) "6c9823906baa11aba873a798cce3a3b3c95465f2"
      4) (empty array)
   4) 1) "127.0.0.1"
      2) (integer) 6379
      3) "2b144f1d7bdb31000a519492be980c6634576462"
      4) (empty array)
2) 1) (integer) 5461
   2) (integer) 10922
   3) 1) "127.0.0.1"
      2) (integer) 6380
      3) "a770892444fbbe4b7d9391b458ac04d6bcba26f0"
      4) (empty array)
   4) 1) "127.0.0.1"
      2) (integer) 6383
      3) "4a56b76a379da615b606a499ae475e986eda3efd"
      4) (empty array)
3) 1) (integer) 10923
   2) (integer) 16383
   3) 1) "127.0.0.1"
      2) (integer) 6381
      3) "76cb8ea9a5d6ba0fa43d31cfa4c33cea8442e07d"
      4) (empty array)
   4) 1) "127.0.0.1"
      2) (integer) 6384
      3) "eaf9833aa105e36b22f6330585a972239bab9f50"
      4) (empty array)

# 查看所有节点信息
127.0.0.1:6379> CLUSTER NODES
2b144f1d7bdb31000a519492be980c6634576462 127.0.0.1:6379@16379 myself,slave 6c9823906baa11aba873a798cce3a3b3c95465f2 0 1669529164000 7 connected
eaf9833aa105e36b22f6330585a972239bab9f50 127.0.0.1:6384@16384 slave 76cb8ea9a5d6ba0fa43d31cfa4c33cea8442e07d 0 1669529164744 3 connected
a770892444fbbe4b7d9391b458ac04d6bcba26f0 127.0.0.1:6380@16380 master - 0 1669529165776 2 connected 5461-10922
6c9823906baa11aba873a798cce3a3b3c95465f2 127.0.0.1:6382@16382 master - 0 1669529166000 7 connected 0-5460
76cb8ea9a5d6ba0fa43d31cfa4c33cea8442e07d 127.0.0.1:6381@16381 master - 0 1669529166792 3 connected 10923-16383
4a56b76a379da615b606a499ae475e986eda3efd 127.0.0.1:6383@16383 slave a770892444fbbe4b7d9391b458ac04d6bcba26f0 0 1669529165000 2 connected
```

#### 数据操作

设置键时, 根据键散列后的值所在的插槽位置自动切换到插槽所在的节点上

```bash
127.0.0.1:6379> KEYS *
(empty array)
127.0.0.1:6379> SET name zhangsan
-> Redirected to slot [5798] located at 127.0.0.1:6380
OK
127.0.0.1:6380> GET name
"zhangsan"

# 集群不支持设置多个键, 需要使用分组的方式
127.0.0.1:6380> MSET age 18 addr beijing
(error) CROSSSLOT Keys in request don't hash to the same slot
127.0.0.1:6380> MSET age{y} 18 addr{y} beijing
OK
127.0.0.1:6380> KEYS *
1) "age{y}"
2) "addr{y}"
3) "name"
127.0.0.1:6379> GET age{y}
-> Redirected to slot [5474] located at 127.0.0.1:6380
"18"
127.0.0.1:6380> GET addr{y}
"beijing"
```

#### 测试节点

使用 `kill` 命令停止端口号为 6381 的 redis 进程时, 集群切换 6381 的状态为失联, 同时将从节点 6384 升级为主节点, 等到 6381 恢复后变为 6384 的从节点

```bash
127.0.0.1:6379> CLUSTER NODES
2b144f1d7bdb31000a519492be980c6634576462 127.0.0.1:6379@16379 myself,slave 6c9823906baa11aba873a798cce3a3b3c95465f2 0 1669529208000 7 connected
eaf9833aa105e36b22f6330585a972239bab9f50 127.0.0.1:6384@16384 slave 76cb8ea9a5d6ba0fa43d31cfa4c33cea8442e07d 0 1669529208000 3 connected
a770892444fbbe4b7d9391b458ac04d6bcba26f0 127.0.0.1:6380@16380 master - 0 1669529207000 2 connected 5461-10922
6c9823906baa11aba873a798cce3a3b3c95465f2 127.0.0.1:6382@16382 master - 0 1669529206000 7 connected 0-5460
76cb8ea9a5d6ba0fa43d31cfa4c33cea8442e07d 127.0.0.1:6381@16381 master - 1669529208334 1669529202178 3 disconnected 10923-16383
4a56b76a379da615b606a499ae475e986eda3efd 127.0.0.1:6383@16383 slave a770892444fbbe4b7d9391b458ac04d6bcba26f0 0 1669529209355 2 connected
```

- kill 停止 6381 端口的进程后再启动进程

```bash
# 6381 恢复后变为 6384 的从节点
127.0.0.1:6379> CLUSTER NODES
2b144f1d7bdb31000a519492be980c6634576462 127.0.0.1:6379@16379 myself,slave 6c9823906baa11aba873a798cce3a3b3c95465f2 0 1669529245000 7 connected
eaf9833aa105e36b22f6330585a972239bab9f50 127.0.0.1:6384@16384 master - 0 1669529243086 8 connected 10923-16383
a770892444fbbe4b7d9391b458ac04d6bcba26f0 127.0.0.1:6380@16380 master - 0 1669529244194 2 connected 5461-10922
6c9823906baa11aba873a798cce3a3b3c95465f2 127.0.0.1:6382@16382 master - 0 1669529245227 7 connected 0-5460
76cb8ea9a5d6ba0fa43d31cfa4c33cea8442e07d 127.0.0.1:6381@16381 slave eaf9833aa105e36b22f6330585a972239bab9f50 0 1669529243552 8 connected
4a56b76a379da615b606a499ae475e986eda3efd 127.0.0.1:6383@16383 slave a770892444fbbe4b7d9391b458ac04d6bcba26f0 0 1669529244000 2 connected
```

#### 添加主节点

`redis-cli --cluster add-node 127.0.0.1:6386 127.0.0.1:6379`   # 添加主节点 6386

##### 重新分配槽位

- 自动分配槽位

`redis-cli --cluster rebalance 127.0.0.1:6379 --cluster-use-empty-masters`

- 带权重重新分片

`redis-cli --cluster rebalance 127.0.0.1:6379 --cluster-weight 127.0.0.1:6380=1.5 --cluster-weight 127.0.0.1:6381=1.0`

##### 手动重新分片

`redis-cli --cluster reshard 127.0.0.1:6379` # 交互式重新分片

#### 添加从节点

按照 [编辑配置文件](#bjpzwj) 创建并修改 `cluster6385.conf` 文件
启动服务器 `redis-server cluster6385.conf`, 同时查看服务器是否正常启动

- `redis-cli --cluster add-node --cluster-slave 127.0.0.1:6385 127.0.0.1:6379` 将 6385 添加为 6379 的从节点

```bash
# 向 6379 节点添加新的从节点
[root@centos7 redis-cluster]# redis-cli --cluster add-node --cluster-slave \
> 127.0.0.1:6385 127.0.0.1:6379
```

- 查看节点 6379 的信息, 显示 2 个从节点

```bash
# 查看节点信息
[root@centos7 redis-cluster]# redis-cli --cluster info 127.0.0.1:6379
127.0.0.1:6379 (8e20e97a...) -> 0 keys | 5461 slots | 2 slaves.
127.0.0.1:6381 (fba8c2ae...) -> 0 keys | 5461 slots | 1 slaves.
127.0.0.1:6380 (57c53adc...) -> 0 keys | 5462 slots | 1 slaves.
[OK] 0 keys in 3 masters.
0.00 keys per slot on average.
```

#### 删除节点

迁移主节点所有槽位

`redis-cli --cluster rebalance 127.0.0.1:6379 --cluster-weight 127.0.0.1:6386=0` # 清空 6386 节点的槽位

`redis-cli --cluster del-node 127.0.0.1:6379 <6386-node-id>`   # 删除节点

### 集群优点

- 实现扩容
- 分担压力
- 无中心配置相对简单

### 集群缺点

- 多建操作不支持, 例如 `MSET` 命令设置多个键不支持, 需要使用分组方式 `MSET name{user} zhangsan age{user} 20 addr{user} beijing`
- 多键的 Redis 事务不支持
- Lua 脚本不支持
- 迁移方案需要整体迁移而不是逐步过渡, 复杂度较大
