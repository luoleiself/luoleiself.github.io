---
title: Redis
date: 2022-11-09 18:03:19
categories:
  - [server, Redis]
tags:
  - Redis
---

## Redis

Remote Dictionary Server 即远程字典服务, 是一个开源的使用 ANSI C 语言编写、支持网络、可基于内存亦可持久化的日志型、Key-Value 数据库，并提供多种语言的 API
Redis 通常被称为数据结构服务器, 因为它的核心数据类型包括字符串、列表、字典（或哈希）、集合和排序集合等大多编程语言都支持的数据类型. 高版本版的 Redis 还添加了计算基数、地理定位和流处理等高级功能

### 数据类型

![redis-1](/images/redis-1.jpg)

<!-- more -->

### 工具命令

- redis-benchmark 压测工具
  - \-h 主机名
  - \-p 端口号
  - \-s socket 连接(覆盖 host 和 port)
  - \-a 认证密码
  - \-\-user 用户名
  - \-c 客户端并发连接数(default 50)
  - \-n 请求数(default 100000)
  - \-d 测试数据的大小(default 3)
  - \-\-dbnum 连接的数据库编号(default 0)
  - \-k 是否保持连接
  - \-r SET/GET/INCR 使用随机 key, SADD 使用随机值
  - \-p 通过管道传输
  - \-q 退出, 仅显示 query/sec 值
  - \-\-csv 以 CSV 格式输出
  - \-l 循环永远运行测试
  - \-t 仅运行以逗号分割的命令列表
  - \-I Idle 模式, 仅打开 N 个 idle 连接并等待
  - \-x 从 STDIN 读取最后一个参数

```shell
127.0.0.1:6379>redis-benchmark -h localhost -p 6379 -c 100 -n 100000 -d 10 -t set,get,hset,hget,lpush,rpush,sadd
```

- redis-check-aof 检查 aof 备份文件
- redis-check-rdb 检查 rdb 备份文件

### CONFIG 命令

- CONFIG GET parameter [parameter...] 获取指定配置项的值
- CONFIG HELP 显示 CONFIG 命令的帮助信息
- CONFIG RESETSTAT 重置 INFO 返回的统计信息, ok 成功
- CONFIG REWRITE 使用内存配置重写配置文件
- CONFIG SET parameter value [parameter value ...] 设置配置项

### Keys 命令

- INFO [section [section ...]] 返回服务的相关信息, 没有参数返回所有

  - server 返回 redis 服务的通用信息
  - clients 返回客户端链接的信息
  - memory 返回内存的信息
  - persistence 返回持久化的信息 RDB 和 AOF
  - stats 返回统计信息
  - replication 返回副本的信息
  - cpu 返回 cpu 的信息
  - commandstats 返回命令统计信息
  - latencystats 返回命令延迟百分比统计信息
  - cluster 返回集群信息
  - modules 返回模块信息
  - keyspace 返回数据库相关统计信息
  - errorstats 返回错误统计信息
  - all 返回所有信息(除了 modules)
  - default 返回默认配置信息
  - everything 返回所有信息(包含 all 和 modules)

- help command 显示命令的帮助信息
- ECHO message 打印信息

- SAVE 保存数据到本地磁盘
- WAIT numreplicas timeout 阻止当前客户端, 直到所有先前的写入命令成功传输并至少由指定数量的副本确认, 如果达到了以毫秒为单位指定的超时, 则即使尚未达到指定的副本数量, 该命令也会返回

- ROLE 返回当前实例的角色是 master、slave、sentinel, 和当前实例上下文副本的信息

- PING [message] 测试连接是否正常, 通常返回 PONG, 如果传入了 message 则会输出 message
- QUIT 关闭退出当前连接
- SHUTDOWN [NOSAVE|SAVE] [NOW] [FORCE] [ABORT] 同步保存数据到硬盘上并关闭服务

- MONITOR 启动监听模式输出服务器执行的每条命令

#### 操作 key

- TYPE key 返回指定 key 的类型, none 表示 key 不存在
- EXISTS key [key ...] 检查指定 key 是否存在, 1 存在, 0 不存在
- KEYS pattern 查找给定模式(pattern)的 key, 返回列表, 未找到返回 (empty array)
- SCAN cursor [MATCH pattern] [COUNT count] [TYPE type] 查找给定模式(pattern)的 key, 返回列表和上次遍历时的游标
- DEL key [key...] 阻塞删除 key 并返回成功删除 key 的数量
- UNLINK key [key ...] 非阻塞从键空间中取消键指定 key 的链接(在其他线程中执行实际的内存回收), 并返回成功取消 key 的数量, 如果 key 不存在则忽略

- RENAME key newKey 修改 key 的名称, 如果指定 key 不存在返回 错误, 如果 newkey 已存在则覆盖
- RENAMENX key newkey 修改 key 的名称, 如果指定 key 不存在返回 错误, 如果 newkey 已存在不执行任何操作返回 0, 否则返回 1

- SORT key [BY pattern] [LIMIT offset count] [GET pattern [GET pattern ...]] [ASC|DESC] [ALPHA] [STORE destination] 对 list、set、zset 集合中的元素进行排序, 默认是按照数字或者元素的双精度浮点数去比较

- MOVE key db 将当前数据库中的 key 移动到指定的数据库(db)中

- DUMP key 序列化指定 key, 并返回被序列化的值, 不存在返回 &lt;nil&gt;

- TOUCH key [key ...] 更改指定 key 的最后一次访问时间并返回修改成功的数量, 如果 key 不存在则忽略

#### 副本

- REPLICAOF host port 将当前服务器设置为指定主机端口上服务器的副本, 通常返回 ok, 5.0.0 开始代替 `SLAVEOF`
  - 如果当前服务器已经是某个服务器的副本, 则取消对旧服务器的连接同步, 并开始对新服务器同步, 丢弃旧有数据集
  - NO ONE 如果当前服务器已经是副本, 此参数将当前服务器变为 master, 并停止与主服务器的连接同步

#### 设置 key 的过期时间

- EXPIRE key seconds [NX|XX|GT|LT] 为指定 key 设置过期时间(单位秒), 1 设置成功, 0 指定 key 不存在或者提供的参数跳过了操作
- EXPIREAT key unix-time-seconds [NX|XX|GT|LT] 为指定 key 设置过期使用 unix 时间戳, 1 设置成功, 0 指定 key 不存在或者提供的参数跳过了操作
- PEXPIRE key milliseconds [NX|XX|GT|LT] 为指定 key 设置过期时间(单位毫秒), 1 设置成功, 0 指定 key 不存在或者提供的参数跳过了操作
- EXPIRETIME key 返回指定 key 将过期的绝对 Unix 时间戳(以秒为单位), -1 表示 key 存在但没有过期时间, -2 表示 key 不存在, 7.0.0 支持
- PEXPIREAT key unix-time-milliseconds [NX|XX|GT|LT] 为指定 key 设置过期时间使用 unix 时间戳, 1 设置成功, 0 指定 key 不存在或者提供的参数跳过了操作
- PEXPIRETIME key 返回指定 key 将过期的绝对 Unix 时间戳(以毫秒为单位), -1 表示 key 存在但没有过期时间, -2 表示 key 不存在, 7.0.0 支持
  - NX 以上命令该参数作用相同, 仅当指定 key 没有过期时间时
  - XX 以上命令该参数作用相同, 仅当指定 key 存在过期时间时
  - GT 以上命令该参数作用相同, 仅当新的过期时间大于当前的过期时间
  - LT 以上命令该参数作用相同, 仅当新的过期时间小于当前的过期时间

#### 获取 key 的过期时间

- TTL key 以秒为单位返回指定 key 的剩余生存时间

  - \-2 key 不存在
  - \-1 key 存在但没有设置剩余生存时间

- PTTL key 以毫秒为单位返回指定 key 的剩余的过期时间

  - \-2 key 不存在
  - \-1 key 存在但没有设置剩余生存时间

- PERSIST key 移除指定 key 的过期时间, key 将永久保持, 1 成功, 0 key 不存在或者未设置过期时间

#### 数据库操作

- RANDOMKEY 从当前数据库随机返回一个 key, 如果当前数据库为空则返回 &lt;nil&gt;
- SWAPDB index1 index2 切换两个数据库
- SELECT index 更改当前连接的选定的数据库
- DBSIZE 返回当前数据库中 key 的数量

- FLUSHALL [ASYNC|SYNC] 清除所有数据库中的 key, 执行成功返回 ok
- FLUSHDB [ASYNC|SYNC] 清除当前数据库中的 key, 执行成功返回 ok

#### 安全认证

- AUTH [username] password 对当前连接的认证

### 发布订阅

Redis 发布/订阅(pub/sub)是一种消息通信模式: 发送者(pub)发送消息, 订阅者(sub)接收消息
它采用事件作为基本的通信机制，提供大规模系统所要求的松散耦合的交互模式: 订阅者(如客户端)以事件订阅的方式表达出它有兴趣接收的一个事件或一类事件;发布者(如服务器)可将订阅者感兴趣的事件随时通知相关订阅者
订阅者对一个或多个频道感兴趣,只需接收感兴趣的消息,不需要知道什么样的发布者发布的. 这种发布者和订阅者的解耦合可以带来更大的扩展性和更加动态的网络拓扑

- 发布者: 无需独占链接, 可以在 publish 发布消息的同时, 使用同一个链接进行其他操作
- 订阅者: 需要独占链接, 在 subscribe 期间, 以阻塞的方式等待消息

#### 普通订阅

- SUBSCRIBE channel [channel ...] 订阅指定频道立即进入阻塞状态等待接收消息
- UNSUBSCRIBE [channel [channel ...]] 根据给定频道取消客户端订阅, 如果未指定则取消所有频道订阅

#### 碎片频道订阅

- SSUBSCRIBE shardchannel [shardchannel ...] 订阅指定的碎片频道, 7.0.0 支持
- SUNSUBSCRIBE [shardchannel [shardchannel ...]] 根据给定碎片频道取消客户端订阅, 如果未指定则取消所有碎片频道订阅, 7.0.0 支持

#### 模式订阅

- PSUBSCRIBE pattern [pattern ...] 根据给定模式订阅频道立即进入阻塞状态等待接收消息
  - pattern 可以使用正则表达式匹配多个频道
- PUNSUBSCRIBE [pattern [pattern ...]] 根据给定模式取消客户端订阅, 如果未指定则取消所有模式订阅

#### 发布消息

- PUBLISH channel message 给指定的频道发送消息并返回接收到消息的订阅者数量, 0 表示没有订阅者
- SPUBLISH shardchannel message 给指定的碎片频道发送消息并返回接收到消息的订阅者数量, 0 表示没有订阅者, 7.0.0 支持

```shell
# 订阅频道
127.0.0.1:6379> PSUBSCRIBE h?llo
Reading messages... (press Ctrl-C to quit)
1) "psubscribe"
2) "h?llo"
3) (integer) 1
# 接收到的消息
1) "pmessage"
2) "h?llo"
3) "hello"
4) "hello,world"
# 接收到的消息
1) "pmessage"
2) "h?llo"
3) "hallo"
4) "hallo,world"
# 发布消息到 hello 和 hallo 频道
127.0.0.1:6379> PUBLISH hello hello,world
(integer) 0
127.0.0.1:6379> PUBLISH hello hello,world
(integer) 2
127.0.0.1:6379> PUBLISH hallo hallo,world
(integer) 2
```

#### 统计订阅信息

- PUBSUB CHANNELS [pattern] 返回当前活跃频道列表(不包含使用模式订阅的频道)

```shell
127.0.0.1:6379> PUBSUB CHANNELS
1) "conn"
```

- PUBSUB NUMSUB [channel [channel ...]] 返回订阅者的数量(不包含使用模式订阅的频道)
  - 如果不指定 channel 将返回 (empty array)

```shell
127.0.0.1:6379> PUBSUB NUMSUB hello conn
1) "hello"
2) (integer) 1
3) "conn"
4) (integer) 1
```

- PUBSUB NUMPAT 返回订阅者通过模式订阅的频道的数量

```shell
127.0.0.1:6379> PUBSUB NUMPAT
(integer) 0
127.0.0.1:6379> PUBSUB NUMPAT
(integer) 1
```

- PUBSUB SHARDCHANNELS [pattern] 返回当前活动的碎片频道, 未找到返回 empty array, 7.0.0 支持
- PUBSUB SHARDNUMSUB [shardchannel [shardchannel ...]] 返回指定的碎片频道的订阅者数量, 未找到返回 empty arryay, 7.0.0 支持

```shell
127.0.0.1:6379> PUBSUB SHARDNUMSUB conn
1) "conn"
2) (integer) 0
```

### 配置文件

- include /path/to/\*.conf # 导入其他 redis 配置文件

- bind 127.0.0.1 -::1 # 默认绑定本地 127.0.0.1
- protected-mode yes # 保护模式, 默认开启
- port 6379 # 默认端口号

- tcp-backlog 511 # tcp 连接数
- timeout 0 # 关闭客户端连接的延迟, 0 表示禁用, 单位秒
- tcp-keepalive 300 # 保持长连接的时间, 单位秒

#### TLS/SSL

安全连接配置项, 默认未开启

- tls-port 6379 # 安全连接端口
- tls-cert-file redis.cert # 安全连接证书
- tls-key-file redis.key # 安全连接 key
- tls-key-file-pass secret # key 文件加密摘要
- tls-client-cert-file client.crt # 客户端安全连接证书
- tls-client-key-file client.key # 客户端安全连接 key
- tls-client-key-file-pass secret # 客户端安全连接 key 文件加密摘要
- tls-ca-cert-file ca.crt # CA 证书
- tls-ca-cert-dir /etc/ssl/certs # CA 证书目录
- tls-auth-clients no # no 不需要也不接受客户端证书连接, optional 证书不必需, 如果提供证书则必须验证有效
- tls-session-caching no # 默认启用 TLS 会话缓存, no 表示禁用缓存
- tls-session-cache-size 5000 # TLS 缓存大小, 默认 20480
- tls-session-cache-timeout 60 # TLS 缓存有效期, 默认 300 秒

#### 通用设置

- daemonize no 是否后台模式启动服务
- pidfile /var/run/redis_6379.pid # 进程 id 文件
- loglevel notice # 设置日志级别, 默认 notice
  - debug (a lot of information, useful for development/testing)
  - verbose (many rarely useful info, but not a mess like the debug level)
  - notice (moderately verbose, what you want in production probably)
  - warning (only very important / critical messages are logged)
- logfile "" # 日志文件, 守护进程模式将指定 /dev/null
- syslog-enabled no # 是否允许指向 系统 日志
- syslog-ident redis # 日志标识符

- databases 16 # 默认数据库数量

- always-show-logo no # 是否总是显示 logo
- set-proc-title yes # 设置进程标题

#### REPLICATION

- replicaof \<masterip\> \<masterport\> 设置主服务器的 IP 和 Port
- masterauth \<master-password\> 设置主服务器的认证密码
- masteruser \<username\> 设置主服务器的用户名
- replica-read-only yes 只读模式, 默认开启
- repl-diskless-sync-delay 5 服务器同步延迟, 默认 5 秒
- replica-priority 100 哨兵模式下被选为主服务器的优先级, 值越小优先级越高
- replica-ignore-maxmemory yes 副本忽略最大内存限制

#### CLUSTER

- cluster-enabled yes 打开集群模式
- cluster-config-file nodes-6379.conf 设置节点配置文件
- cluster-node-timeout 15000 设置节点失联时间, 超过该时间集群自动主从切换, 默认毫秒
- cluster-allow-replica-migration yes 允许集群副本迁移
- cluster-require-full-coverage yes 当某一段插槽主从服务器都宕机, 设置 yes 则整个集群都挂掉, 设置 no 则只是该插槽不可用
- cluster-allow-pubsubshard-when-down yes 允许集群服务器宕机时发布/订阅

#### MEMORY

- maxmemory \<bytes\> # 设置内存容量
- maxmemory-policy noeviction # 内存管理策略
  - volatile-lru 使用 LRU 算法移除 key, 只对设置了过期时间的 key
  - allkeys-lru 在所有集合 key 中, 使用 LRU 算法移除 key
  - volatile-lfu 使用 LFU 算法移除 key, 只对设置了过期时间的 key
  - allkeys-lfu 在所有集合 key 中, 使用 LFU 算法移除 key
  - volatile-random 在过期集合 key 中, 移除随机的 key, 只对设置了过期时间的 key
  - allkeys-random 在所有集合 key 中, 移除随机的 key
  - volatile-ttl 移除那些 TTL 值最小的 key, 即那些最近要过期的 key
  - noeviction 不进行移除, 针对写操作, 只是返回错误信息
- maxmemory-samples 5 # 设置 Redis 移除 key 时的样本数量, 10 接近 LRU 算法但非常消耗内存, 3 最快却不是精确的

#### SNAPSHOTTING

- save 3600 1 300 100 60 10000 # 快照执行机制, 3600 秒后如果超过 1 次更改, 300 秒后超过 100 次更改, 60 秒后超过 10000 次更改

```shell
save <seconds> <changes> [<seconds> <changes> ...]
```

- stop-writes-on-bgsave-error yes # 是否开启停止在保存快照发生错误的时的写操作
- rdbcompression yes # 开启 rdb 文件压缩
- rdbchecksum yes # 开启 rdb 文件的校验检查
- dbfilename dump.rdb # rdb 文件名称
- dir ./ # rdb 文件存储目录

- appendonly no # 是否启动 aof 备份
- appendfilename "appendonly.aof" # aof 备份文件名
- appenddirname "appendonlydir" # aof 备份目录
- appendfsync everysec # aof 备份模式, 每秒中执行
  - always 只要 key 发生改变就要备份
  - no 不备份
- auto-aof-rewrite-percentage 100 # 代表当前 AOF 文件大小和上一次重写后 AOF 文件大小的比值
- auto-aof-rewrite-min-size 64mb # 触发 AOF 重写的文件最小值, 默认 64MB

#### SECURITY

- acllog-max-len 128 # ACL 日志在内存中时的最大条目数
- aclfile /etc/redis/users.acl # 默认 ACL 配置文件
- requirepass foobared # 认证密码

- maxclients 10000 # 客户端最大连接数

- io-threads 4 # I/O 线程

### ACL

ACL(access control list)访问控制列表的简称, 是为了控制某些 Redis 客户端在访问 Redis 服务器时, 能够执行的命令和能够获取的 key, 提高操作安全性, 避免对数据造成损坏

#### 规则分类

|         参数         | 说明                                                                                    |
| :------------------: | --------------------------------------------------------------------------------------- |
|          on          | 表示启动该用户, 默认为 off                                                              |
|        nopass        | 删除所有与用户关联的密码                                                                |
|        reset         | 移除用户的所有功能, 并关闭用户                                                          |
|      +[command]      | 将命令添加到用户可以调用的命令列表中                                                    |
|      -[command]      | 将命令从用户可以调用的命令列表中移除                                                    |
|  +[command]\|subcmd  | 允许使用已禁用命令的特定子命令                                                          |
|     +@[category]     | 允许用户调用 category 类别中的所有命令, 可以使用 `ACL CAT` 命令查看所有类别             |
|     -@[category]     | 禁止用户调用 category 类别中的所有命令                                                  |
|     allcommands      | +@all 的别名                                                                            |
|      nocommands      | -@all 的别名                                                                            |
|     ~\<pattern\>     | 允许用户可以访问的 key(正则匹配), 例如: ~foo:\* 只允许访问 foo:\* 的 key                |
|    %R~\<pattern\>    | 添加指定的只读 key(正则匹配), 例如: %R~app:\* 只允许读 app:\* 的 key, 7.0 支持          |
|    %W~\<pattern\>    | 添加指定的只写 key(正则匹配), 例如: %W~app:\* 只允许写 app:\* 的 key, 7.0 支持          |
|   %RW~\<pattern\>    | 添加指定的可读可写的 key(正则匹配), 例如: %RW~app:\* 只允许读写 app:\* 的 key, 7.0 支持 |
|       allkeys        | ~\* 的别名                                                                              |
|      resetkeys       | 移除所有的 key 匹配模式                                                                 |
|     &\<pattern\>     | 允许用户可使用的 Pub/Sub 通道(正则匹配)                                                 |
|     allchannels      | &\* 的别名                                                                              |
|    resetchannels     | 移除所有的通道匹配模式                                                                  |
|    \>\<password\>    | 为用户添加明文密码, 服务器自动转换成 hash 存储, 例如: >123456                           |
|    \<\<password\>    | 从有效密码列表中删除密码                                                                |
| #\<hashedpassword\>  | 为用户添加 hash 密码, 例如: #cab3...c4f2                                                |
| \!\<hashedpassword\> | 从有效密码列表中删除密码                                                                |

- ACL HELP 显示 ACL 的帮助信息

```shell
127.0.0.1:6379> ACL HELP
 1) ACL <subcommand> [<arg> [value] [opt] ...]. Subcommands are:
 2) CAT [<category>]
 3)     List all commands that belong to <category>, or all command categories
 4)     when no category is specified.
 5) DELUSER <username> [<username> ...]
 6)     Delete a list of users.
 7) DRYRUN <username> <command> [<arg> ...]
 8)     Returns whether the user can execute the given command without executing the command.
 9) GETUSER <username>
10)     Get the user\'s details.
11) GENPASS [<bits>]
12)     Generate a secure 256-bit user password. The optional `bits` argument can
13)     be used to specify a different size.
14) LIST
15)     Show users details in config file format.
16) LOAD
17)     Reload users from the ACL file.
18) LOG [<count> | RESET]
19)     Show the ACL log entries.
20) SAVE
21)     Save the current config to the ACL file.
22) SETUSER <username> <attribute> [<attribute> ...]
23)     Create or modify a user with the specified attributes.
24) USERS
25)     List all the registered usernames.
26) WHOAMI
27)     Return the current connection username.
28) HELP
29)     Prints this help.
```

- ACL SETUSER 设置用户访问权限
- ACL GETUSER username 获取指定用户的权限

```shell
# 添加 lisi 账号, 明文密码 123456, 添加所有分类的命令
127.0.0.1:6379> ACL SETUSER lisi >123456 off +@all
OK
127.0.0.1:6379> ACL GETUSER lisi
 1) "flags"
 2) 1) "off"
 3) "passwords"
 4) 1) "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92"
 5) "commands"
 6) "+@all"
 7) "keys"
 8) ""
 9) "channels"
10) ""
11) "selectors"
12) (empty array)
# 添加禁用账号 zhangsan
# 只包含 string, hash, list, set分类下的命令
# 只能操作以 zhang 开头匹配模式的 key 和通道
127.0.0.1:6379> ACL SETUSER zhangsan off +@string +@hash +@list +@set ~zhang:* &zhang:*
OK
127.0.0.1:6379> ACL GETUSER zhangsan
 1) "flags"
 2) 1) "off"
 3) "passwords"
 4) 1) "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92"
 5) "commands"
 6) "-@all +@list +@string +@hash +@set"
 7) "keys"
 8) "~zhang:*"
 9) "channels"
10) "&zhang:*"
11) "selectors"
12) (empty array)
```

- ACL LIST 显示 Redis 服务器当前活动的 ACL 规则

```shell
127.0.0.1:6379> ACL LIST
1) "user default on nopass ~* &* +@all"
2) "user zhangsan off ~zhang:* resetchannels &zhang:* -@all +@list +@string +@hash +@set"
```

- ACL DRYRUN username command [arg [arg ...]] 模拟指定用户对给定命令的执行, 此命令可以用来测试用户的权限而无需启用用户, 7.0.0 支持

```shell
127.0.0.1:6379> ACL DRYRUN zhangsan SET name zhangsan
"This user has no permissions to access the 'name' key"
# 只能操作以 zhang 开头匹配模式的 key
127.0.0.1:6379> ACL DRYRUN zhangsan SET zhang:name zhangsan
OK
```

### 事务

=== Redis 单条命令是保证原子性的, 但是 Redis 事务不保证原子性 ===
=== Redis 事务没有隔离级别的概念 ===
所有的命令在事务中, 并没有被执行而是加入队列, 只有发起执行命令的时候才会执行! EXEC

Redis 事务允许在一个步骤中执行一组命令, 事务中的所有命令都是串行化的, 并按顺序执行. 在 Redis 事务执行过程中, 另一个客户端发送的请求将永远不会被处理. 这保证了命令作为一个单独的操作执行
Redis 事务执行的三个重要保证:

- 批量操作在发送 EXEC 命令前被放入队列缓存
- 收到 EXEC 命令后进入事务执行, 事务中任意命令执行失败, 其余的命令依然被执行
- 事务执行过程, 其他客户端提交的命令请求不会插入到事务执行命令序列中

事务开始到执行的三个阶段

- 开始事务(multi)
- 命令入队
- 执行事务(exec)

#### 命令

- MULTI 开启事务, 通常返回 ok
- DISCARD 丢弃事务, 通常返回 ok
  - 必须在 `MULTI` 命令之后才能调用, 否则报错 ERR DISCARD without MULTI
- EXEC 执行事务, 通常返回 ok

  - 必须在 `MULTI` 命令之后才能调用, 否则报错 ERR EXEC without MULTI
  - 如果 `WATCH` 观察的 key 在当前的事务执行时已被修改, 则返回 &lt;nil&gt;

- WATCH key [key ...] 观察指定 key, 通常返回 ok, 如果在事务执行之前观察的 key 被修改, 则事务将被打断
  - 如果在 `MULTI` 命令后调用, 则会报错 ERR WATCH inside MULTI is not allowed
- UNWATCH 取消所有观察的 key, 通常返回 ok, 如果调用了 `EXEC` 或 `DISCARD` 命令, 通常不再需要调用此命令

```shell
127.0.0.1:6379> GET money
"250"
127.0.0.1:6379> WATCH money # 观察 money
OK
127.0.0.1:6379> MULTI # 开启事务
OK
# 修改 money 加上增量 1000 命令入事务队列
127.0.0.1:6379(TX)> INCRBY money 1000
QUEUED
###################################
# 另一个客户端连接修改 money 加上增量 50
127.0.0.1:6379> INCRBY money 50
(integer) 300
###################################
127.0.0.1:6379(TX)> EXEC # 执行事务
(nil)
127.0.0.1:6379> GET money
"300"
```

#### 编译时错误

```shell
127.0.0.1:6379> set key1 hello
OK
127.0.0.1:6379> MULTI # 开启事务
OK
# 此处的错误在命令加入事务队列时发现, 直接报告, 导致整个事务的执行失败
127.0.0.1:6379(TX)> INCR key1 10
(error) ERR wrong number of arguments for 'incr' command
127.0.0.1:6379(TX)> set key2 key2 # 命令入队列
QUEUED
127.0.0.1:6379(TX)> get key2
QUEUED
127.0.0.1:6379(TX)> exec  # 执行事务
(error) EXECABORT Transaction discarded because of previous errors.
127.0.0.1:6379> get key2
(nil)
```

#### 运行时错误

```shell
127.0.0.1:6379> MULTI # 开启事务
OK
127.0.0.1:6379(TX)> set key1 hello  # 命令入队列
QUEUED
# 此处的错误在命令运行时才能发现, 但不影响下面的命令的执行
127.0.0.1:6379(TX)> INCR key1
QUEUED
127.0.0.1:6379(TX)> set key2 key2
QUEUED
127.0.0.1:6379(TX)> get key2  # 命令执行成功
QUEUED
127.0.0.1:6379(TX)> EXEC  # 执行事务
1) OK
2) (error) ERR value is not an integer or out of range
3) OK
4) "key2"
127.0.0.1:6379> get key2
"key2"
```

### 持久化

Redis 是基于内存的数据库, 遇到断电就会丢失数据, 持久化就是将内存中的数据保存到磁盘中便于以后使用, Redis 提供了 RDB 和 AOF 两种持久化方式, 默认使用 RDB 方式持久化数据
Redis 在持久化的过程中, 会先将数据写入到一个临时的文件中, 待持久化过程结束后, 才会用这个临时文件替换赏赐持久化生成的文件

#### 触发方式

- 通过配置文件配置项会触发持久化操作

```shell
# 3600秒至少有 1 次修改, 300秒至少有 100 次修改, 60秒至少有 10000 次修改
save 3600 1 300 100 60 10000
```

- 通过 `FLUSHALL` 命令主动触发
- 通过 `SAVE` 命令主动触发

#### RDB

RDB(Redis Database), 在某一时刻将 Redis 中的数据生成一份快照保存到磁盘中
Redis 会单独 fork 一个子进程进行持久化, 而主进程不会进行任何 I/O 操作, 这样就保证了 Redis 极高的性能, 如果需要进行大规模数据的恢复,且对于数据恢复的完整性不是非常敏感, 此方式比 AOF 方式更加的高效

- `dbfilename dump.rdb` 默认文件名
- `dir ./` 默认存储目录

检查 RDB 文件的命令 `redis-check-rdb`

##### RDB 优点

- 每隔一段时间完全备份一次
- 容灾简单, 可远程传输
- RDB 最大限度地提高了 Redis 的性能
- 文件较大时重启和恢复速度要快

##### RDB 缺点

- 如果发生故障, 最后备份的数据可能会丢失, 无法处理实时备份
- RDB 需要经常 fork() 以便使用子进程在磁盘上持久化, 增加 CPU 的负担

#### AOF

AOF(Append Only File), 将执行过的写命令全部记录下来, 在数据恢复时按照从前往后的顺序再将指令都执行一遍, AOF 的持久化策略在开启后默认是每秒钟 sync 一次

- `appendonly yes` 启动 AOF 模式
- `appendfilename appendonly.aof` 默认文件名
- `appenddirname appendonlydir` 默认存储目录
- `appendfsync everysec` 持久化策略, 每秒钟执行一次, 可以修改为 `always` 和 `no`
  - `always` 每次将新命令附加到 AOF 时, 速度慢, 但是最安全
  - `no` 将写入策略权交给操作系统, 速度快, 但是不安全
- `auto-aof-rewrite-percentage 100` aof 重写的基准值, 当达到 100% 时重写
- `auto-aof-rewrite-min-size 64mb` 当文件大小达到 64mb 的 100% 时重写

如果 appendonly.aof 文件有错误, Redis 服务将会启动失败

- redis-check-aof 检查 AOF 文件, \-\-fix 参数修复文件的错误, 通常会丢弃文件中无法识别的命令

##### AOF 优点

- AOF 更耐用, 可以在几秒钟内完成备份
- 当数据过大时, Redis 可以在后台自动重写 AOF, 节省空间
- AOF 实时性比较好, 并且支持配置写入策略

##### AOF 缺点

- AOF 文件一半大于 RDB 的文件
- 即使 AOF 有写入策略, 但是本质上速度还是慢于 RDB

#### RDB 和 AOF 组合

`aof-use-rdb-preamble yes`

### 主从复制

将一台 Redis 服务器的数据,复制到其他的 Redis 服务器. 前者称为主节点(Master/Leader),后者称为从节点(Slave/Follower), 数据的复制是单向的！只能由主节点复制到从节点(主节点以写为主、从节点以读为主)—— 读写分离.
===每台 Redis 服务器都是主节点===
一个主节点可以有 0 个或者多个从节点, 但每个从节点只能由一个主节点

```shell
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

#### 作用

- 数据冗余: 主从复制实现了数据的热备份, 是持久化之外的一种数据冗余的方式
- 故障恢复: 当主节点故障时, 从节点可以暂时替代主节点提供服务, 是一种服务冗余的方式
- 负载均衡: 在主从复制的基础上, 配合读写分离, 由主节点进行写操作, 从节点进行读操作, 分担服务器的负载; 尤其是在多读少写的场景下, 通过多个从节点分担负载, 提高并发量
- 高可用(集群)基石: 主从复制还是哨兵和集群能够实施的基础

#### 复制原理

- 从服务器向主服务器发送 `SYNC` 命令
- 接到 `SYNC` 命令的主服务器会调用 `BGSAVE` 命令, 创建一个 RDB 文件, 并使用缓冲区记录接下来执行的所有写命令
- 当主服务器执行完 `BGSAVE` 命令时, 它会向从服务器发送 RDB 文件, 而从服务器则会接收并载入这个文件
- 主服务器将缓冲区储存的所有写命令发送给从服务器执行

##### 全量复制

从服务器接收到数据库文件后, 将其全部加载到内存中

##### 增量复制

主服务器将新的所有收集到的修改命令依次传给从服务器, 完成同步

#### 命令模式配置

===每台 Redis 服务器都是主节点===, 只用配置从服务器即可
使用`命令配置`只能在`本次服务器运行时有效`, 重启服务器后将会丢失配置信息, 使用配置文件永久生效

##### 运行时有效

- 方式一: 启动 Redis 服务器时参数指定 `redis-server --port 6380 --replicaof 127.0.0.1 6379`
- 方式二: 连接 Redis 服务器使用内置命令 `REPLICAOF host port`

提升从服务器角色

- REPLICAOF NO ONE 将从服务器更改为主服务器

##### 永久有效

单机主从配置: 新建多个 Redis 服务器配置文件并修改其中关键项

- `bind 127.0.0.1` 修改绑定的 ip
- `port 6379` 修改绑定的端口号
- `daemonize yes` 开启后台运行, 默认为 no
- `pidfile /var/run/redis_6379.pid` 修改进程文件, 默认为 redis_6379.pid
- `logfile "6379.log"` 修改日志文件名, 默认为空
- `dbfilename dump6379.rdb` 修改持久化文件名, 默认为 dump.rdb

```shell
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

- 读写数据

```shell
# 主机写入数据
127.0.0.1:6379> SET name helloworld
OK
# 从机读取数据
127.0.0.1:6380> GET name
"helloworld"
127.0.0.1:6380> SET age 18 # 从机写入数据报错
(error) READONLY You can't write against a read only replica.
# 从机读取数据
127.0.0.1:6381> GET name
"helloworld"
127.0.0.1:6381> SET age 18 # 从机写入数据报错
(error) READONLY You can't write against a read only replica.
```

#### 配置文件配置

redis.conf

- replicaof &lt;masterip&gt; &lt;masterport&gt; 配置主服务器 ip 和 port
- masterauth &lt;master-password&gt; 主服务器认证密码, 如果需要
- masteruser &lt;username&gt; 主服务器用户
- replica-read-only yes 只读模式, 默认开启
- repl-diskless-sync yes 不使用向磁盘写 rdb 文件通信的方式直接通过新建进程 socket 同步 rdb 文件
- repl-diskless-sync-delay 5 同步延迟, 默认 5 秒
- replica-priority 100 哨兵模式下被选为主服务器的优先级, 值越小优先级越高

#### 哨兵模式配置

哨兵模式是一种特殊的模式, 首先 Redis 提供了启动哨兵的工具命令, 哨兵是一个独立的进程, 作为进程, 它会独立运行
哨兵模式选举: 首先, 领头 sentinel 根据从服务器的信号反馈将从服务器列表中失联的服务器剔除, 按照从服务器的优先级(replica-priority)进行排序并选择优先级最高的从服务器, 如果有多个具有相同最高优先级的从服务器, 那么, 领头 sentinel 将多个具有相同最高优先级的从服务器按照复制偏移量(复制积压缓冲区中存储的写操作的字节占用累加, 主从服务器进行 PSYNC 使用)进行排序并选择其中偏移量最大(偏移量最大保存的数据最新)的从服务器, 如果有多个优先级最高, 复制偏移量最大的从服务器, 那么 领头 sentinel 将按照从服务器的运行 ID 进行排序并选择其中 ID 最小的从服务器

- 哨兵通过发送命令, 让 Redis 服务器返回监控其运行状态, 包括主服务器和从服务器
- 当哨兵监测到 master 宕机, 会自动将 slave 切换成 master, 然后通过发布订阅模式通知其他的从服务器, 修改配置文件并关联新的主服务器

默认配置文件 `sentinel.conf`

```shell
protected-mode no # 保护模式, 默认不开启
port 26379 # 服务端口号
daemonize no # 是否后台运行模式
pidfile /var/run/redis-sentinel.pid # 进程文件
# sentinel announce-ip <ip> # 广播地址
# sentinel announce-port <port> # 广播端口
logfile "" # 日志文件
dir /tmp # 工作目录
sentinel monitor mymaster 127.0.0.1 6379 2 # 监测服务器配置, 数字表示确认主服务器宕机的票数
# sentinel auth-pass <master-name> <password> # 认证配置
sentinel down-after-milliseconds mymaster 30000 # 不可触达的超时时间, 默认 30 s
sentinel parallel-syncs mymaster 1 # 当主服务器宕机时支持最大同时重配服务器的数量, 默认 1
sentinel failover-timeout mymaster 180000 # 当服务器宕机后等待再次重启的时间, 默认 3 min
# sentinel notification-script <master-name> <script-path> # 服务器唤起脚本文件
sentinel deny-scripts-reconfig yes # 拒绝脚本配置, 默认拒绝
```

- 方式一: 使用命令 `redis-server /path/to/sentinel.conf --sentinel` 开启哨兵模式
- 方式二: 使用命令 `redis-sentinel /path/to/sentinel.conf` 开启哨兵模式

```shell
# sentinel.conf
sentinel monitor myredis 127.0.0.1 6379 1
```

![redis-2](/images/redis-2.png)
![redis-3](/images/redis-3.png)

##### 一主三从哨兵配置

3 个哨兵配置文件

```shell
# sentinel26379.conf
port 26379
pidfile /var/run/redis-sentinel-26379.pid
logfile "26379.log"
dir /tmp
sentinel monitor myredis 127.0.0.1 6379 2

# sentinel36379.conf
port 36379
pidfile /var/run/redis-sentinel-36379.pid
logfile "36379.log"
dir /tmp
sentinel monitor myredis 127.0.0.1 6379 2

# sentinel46379.conf
port 46379
pidfile /var/run/redis-sentinel-46379.pid
logfile "46379.log"
dir /tmp
sentinel monitor myredis 127.0.0.1 6379 2
```

3 台 redis 服务器配置文件

```shell
# redis6379.conf
bind 127.0.0.1
port 6379
daemonize yes
pidfile /var/run/redis_6379.pid
logfile "6379.log"
dbfilename dump6379.rdb

# redis6380.conf
bind 127.0.0.1
port 6380
daemonize yes
pidfile /var/run/redis_6380.pid
logfile "6380.log"
dbfilename dump6380.rdb
replicaof 127.0.0.1 6379 # 配置主服务器 ip 和 port

# redis6381.conf
bind 127.0.0.1
port 6381
daemonize yes
pidfile /var/run/redis_6381.pid
logfile "6381.log"
dbfilename dump6381.rdb
replicaof 127.0.0.1 6379 # 配置主服务器 ip 和 port
```

```shell
[root@centos7 ~]# redis-server .config/redis6379.conf # 启动 redis 服务器
[root@centos7 ~]# redis-server .config/redis6380.conf # 启动 redis 服务器
[root@centos7 ~]# redis-server .config/redis6381.conf # 启动 redis 服务器
[root@centos7 ~]# redis-sential .config/sentinel26379.conf # 启动哨兵
[root@centos7 ~]# redis-sential .config/sentinel36379.conf # 启动哨兵
[root@centos7 ~]# redis-sential .config/sentinel46379.conf # 启动哨兵
```

### 集群

Redis Cluster 是一种服务器 Sharding 技术, Redis 3.0 版本开始支持
在 Redis Cluster 中, Sharding 采用 slot 的概念, 一共分成 16384 个 slot, 对于每个进入 Redis 的键值对, 对 key 执行 CRC16 算法然后再对 16384 取模, 得到的结果就是对应的 slot.
Redis 集群中的每个 node 负责分摊这 16384 个 slot 中的一部分, 当动态添加或减少 node 时, 需要将 16384 个 slot 再分配, slot 中的键值对也要迁移, 这一过程目前还处于半自动状态仍需要人工介入, 如果某个 node 发生故障, 则此 node 负责的 slot 也就失效, 整个集群将不能工作
官方推荐的方案是将 node 配置成主从结构, 即 1:n, 如果主节点失效, Redis Cluster 根据选举算法从 slave 节点中选择一个升级为主节点继续提供服务, 如果失效的主节点恢复正常后则作为新的主节点的从节点

- redis-cli \-\-cluster help # 查看集群命令帮助信息
- redis-cli \-\-cluster create host1:port1 ... hostN:portN # 创建指定 IP 和 Port 的服务器作为集群
  - \-\-cluster-replicas \<arg\> # 指定集群中主节点和从节点数量的比例, 1 表示 1:1
- redis-cli \-\-cluster add-node new_host:new_port existing_host:existing_port # 添加集群节点
  - \-\-cluster-slave # 添加集群节点从服务器
  - \-\-cluster-master-id \<arg\> # 添加到指定主服务器下
- redis-cli \-\-cluster reshard \<host:port\> # 重新分配节点的 hash 插槽
  - \-\-cluster-from \<arg\> # 已有节点 id, 多个 id 之间使用半角逗号分隔
  - \-\-cluster-to \<arg\> # 新节点 id
  - \-\-cluster-slots \<arg\> # 新节点的 hash 槽数量
- redis-clis \-\-cluster rebalance \<host:port\> # 重新分配节点
  - \-\-cluster-weight \<node1=w1...nodeN=wN\> # 分配节点权重
  - \-\-cluster-timeout \<arg\> # 节点超时时间
  - \-\-cluster-threshold \<arg\> # 节点阈值
- redis-cli \-\-cluster import host:port # 导入指定节点
  - \-\-cluster-from \<arg\> # 从指定 id
  - \-\-cluster-from-user \<arg\> # 指定用户名
  - \-\-cluster-from-pass \<arg\> # 指定密码
- redis-cli \-\-cluster info \<host:port\> # 查看指定节点信息
- redis-cli \-\-cluster check \<host:port\> # 检查指定节点
- redis-cli \-\-cluster del-node host:port node_id # 删除集群节点
- redis-cli \-\-cluster call host:port command arg arg ... arg # 集群节点执行指定命令
- redis-cli \-\-cluster set-timeout host:port milliseconds # 设置节点的超时时间
- redis-cli \-\-cluster backup host:port backup_directory # 备份节点数据到指定目录

使用 `redis-cli -c -p port` 命令接入集群节点

#### 集群部署

##### 编辑配置文件 <em id="bjpzwj"></em> <!-- markdownlint-disable-line -->

创建 Redis 服务器配置文件, 引入默认配置文件并覆盖配置项, 开启集群模式
创建 `redis6379.conf`, `redis6380.conf`, `redis6381.conf`, `redis6382.conf`, `redis6383.conf`, `redis6384.conf` 6 个文件
修改其中的 ip, port, pidfile, cluster-enabled, cluster-config-file

```shell
# 引入 redis 默认配置文件
include /root/redis-cluster/redis.conf
# 修改绑定 ip, 此处演示全为本机
bind 127.0.0.1
# 修改 redis 端口号, 本机演示需要修改, 多机器时可以不用
port 6379
# 关闭保护模式, 默认 yes
protected-mode no
# 开启后台运行, 默认 no
daemonize yes
# 修改 redis 进程文件名
pidfile /var/run/redis_6379.pid
# 开启集群模式
cluster-enabled yes
# 修改集群节点文件名, 默认在存储在当前目录下
cluster-config-file nodes-6379.conf
# 设置节点失联时间, 超过该时间集群自动切换主从节点, 默认毫秒
cluster-node-timeout 15000
# 关闭当某一插槽主从服务器挂掉时, 整个集群都挂掉, no 只是该插槽不可用, 默认 yes
cluster-require-full-coverage no
```

##### 启动 Redis 服务器

启动所有的 redis 服务器, 使用 `ps -ef | grep redis` 命令查看 redis 服务器进程

```shell
[root@centos7 redis-cluster]# redis-server redis6379.conf
[root@centos7 redis-cluster]# redis-server redis6380.conf
[root@centos7 redis-cluster]# redis-server redis6381.conf
[root@centos7 redis-cluster]# redis-server redis6382.conf
[root@centos7 redis-cluster]# redis-server redis6383.conf
[root@centos7 redis-cluster]# redis-server redis6384.conf
[root@centos7 redis-cluster]# ps -ef | grep redis
root      3731     1  0 05:49 ?        00:00:00 redis-server 127.0.0.1:6379 [cluster]
root      3737     1  0 05:49 ?        00:00:00 redis-server 127.0.0.1:6380 [cluster]
root      3743     1  0 05:49 ?        00:00:00 redis-server 127.0.0.1:6381 [cluster]
root      3749     1  0 05:49 ?        00:00:00 redis-server 127.0.0.1:6382 [cluster]
root      3755     1  0 05:49 ?        00:00:00 redis-server 127.0.0.1:6383 [cluster]
root      3761     1  0 05:49 ?        00:00:00 redis-server 127.0.0.1:6384 [cluster]
```

##### 创建集群节点

使用 `redis-cli --cluster create --cluster-replicas arg hostN:portN` 命令创建集群节点, arg 参数表示集群主从节点的数量比例, 1 表示 1:1
创建过程中提示输入 `yes` 表示接受当前配置信息并写入指定文件中, 最后输出 `[OK] All 16384 slots covered.` 表示集群创建完成

```shell
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

##### 连接 Redis 服务器

使用 `redis-cli -c -p port` 命令接入集群节点

- -c 允许集群模式接入

```shell
[root@centos7 redis-cluster]# redis-cli -c -p 6379
```

##### 集群命令

- CLUSTER HELP # 在 Redis 命令行中查看所有集群操作命令

```shell
127.0.0.1:6380> CLUSTER HELP
```

##### 查看节点信息

方式一: 命令行中使用 `redis-cli --cluster info host:port` 命令查看指定节点的信息

```shell
[root@centos7 redis-cluster]# redis-cli --cluster info 127.0.0.1:6380
127.0.0.1:6380 (a7708924...) -> 3 keys | 5462 slots | 1 slaves.
127.0.0.1:6382 (6c982390...) -> 0 keys | 5461 slots | 0 slaves.
127.0.0.1:6381 (76cb8ea9...) -> 0 keys | 5461 slots | 1 slaves.
[OK] 3 keys in 3 masters.
0.00 keys per slot on average.
```

方式二: 在 Redis 命令行中使用 `CLUSTER INFO` 查看节点信息

```shell
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

##### 数据操作

设置键时, 根据键散列后的值所在的插槽位置自动切换到插槽所在的节点上

```shell
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
127.0.0.1:6380> MSET age{user} 18 addr{user} beijing
OK
127.0.0.1:6380> KEYS *
1) "age{user}"
2) "addr{user}"
3) "name"
127.0.0.1:6379> GET age{user}
-> Redirected to slot [5474] located at 127.0.0.1:6380
"18"
127.0.0.1:6380> GET addr{user}
"beijing"
```

##### 测试节点

使用 `kill` 命令停止端口号为 6381 的 redis 进程时, 集群切换 6381 的状态为失联, 同时将从节点 6384 升级为主节点, 等到 6381 恢复后变为 6384 的从节点

```shell
127.0.0.1:6379> CLUSTER NODES
2b144f1d7bdb31000a519492be980c6634576462 127.0.0.1:6379@16379 myself,slave 6c9823906baa11aba873a798cce3a3b3c95465f2 0 1669529208000 7 connected
eaf9833aa105e36b22f6330585a972239bab9f50 127.0.0.1:6384@16384 slave 76cb8ea9a5d6ba0fa43d31cfa4c33cea8442e07d 0 1669529208000 3 connected
a770892444fbbe4b7d9391b458ac04d6bcba26f0 127.0.0.1:6380@16380 master - 0 1669529207000 2 connected 5461-10922
6c9823906baa11aba873a798cce3a3b3c95465f2 127.0.0.1:6382@16382 master - 0 1669529206000 7 connected 0-5460
76cb8ea9a5d6ba0fa43d31cfa4c33cea8442e07d 127.0.0.1:6381@16381 master - 1669529208334 1669529202178 3 disconnected 10923-16383
4a56b76a379da615b606a499ae475e986eda3efd 127.0.0.1:6383@16383 slave a770892444fbbe4b7d9391b458ac04d6bcba26f0 0 1669529209355 2 connected

# 6381 恢复后变为 6384 的从节点
127.0.0.1:6379> CLUSTER NODES
2b144f1d7bdb31000a519492be980c6634576462 127.0.0.1:6379@16379 myself,slave 6c9823906baa11aba873a798cce3a3b3c95465f2 0 1669529245000 7 connected
eaf9833aa105e36b22f6330585a972239bab9f50 127.0.0.1:6384@16384 master - 0 1669529243086 8 connected 10923-16383
a770892444fbbe4b7d9391b458ac04d6bcba26f0 127.0.0.1:6380@16380 master - 0 1669529244194 2 connected 5461-10922
6c9823906baa11aba873a798cce3a3b3c95465f2 127.0.0.1:6382@16382 master - 0 1669529245227 7 connected 0-5460
76cb8ea9a5d6ba0fa43d31cfa4c33cea8442e07d 127.0.0.1:6381@16381 slave eaf9833aa105e36b22f6330585a972239bab9f50 0 1669529243552 8 connected
4a56b76a379da615b606a499ae475e986eda3efd 127.0.0.1:6383@16383 slave a770892444fbbe4b7d9391b458ac04d6bcba26f0 0 1669529244000 2 connected
```

##### 查看节点配置文件

```shell
[root@centos7 redis-cluster]# cat nodes-6381.conf
76cb8ea9a5d6ba0fa43d31cfa4c33cea8442e07d 127.0.0.1:6381@16381 myself,slave eaf9833aa105e36b22f6330585a972239bab9f50 0 1669529243521 8 connected
2b144f1d7bdb31000a519492be980c6634576462 127.0.0.1:6379@16379 slave 6c9823906baa11aba873a798cce3a3b3c95465f2 0 1669529243524 7 connected
4a56b76a379da615b606a499ae475e986eda3efd 127.0.0.1:6383@16383 slave a770892444fbbe4b7d9391b458ac04d6bcba26f0 0 1669529243524 2 connected
eaf9833aa105e36b22f6330585a972239bab9f50 127.0.0.1:6384@16384 master - 0 1669529243529 8 connected 10923-16383
6c9823906baa11aba873a798cce3a3b3c95465f2 127.0.0.1:6382@16382 master - 0 1669529243529 7 connected 0-5460
a770892444fbbe4b7d9391b458ac04d6bcba26f0 127.0.0.1:6380@16380 master - 0 1669529243524 2 connected 5461-10922
vars currentEpoch 8 lastVoteEpoch 7
```

##### 添加新节点

按照 <a href="#bjpzwj">编辑配置文件</a> 创建并修改 `redis6385.conf` 文件 <!-- markdownlint-disable-line -->
启动服务器 `redis-server redis6385.conf`, 同时查看服务器是否正常启动

使用命令 `redis-cli --cluster add-node --cluster-slave 127.0.0.1:6385 127.0.0.1:6379` 将 6385 添加为 6379 的从节点

```shell
# 向 6379 节点添加新的从节点
[root@centos7 redis-cluster]# redis-cli --cluster add-node --cluster-slave \
> 127.0.0.1:6385 127.0.0.1:6379
>>> Adding node 127.0.0.1:6385 to cluster 127.0.0.1:6379
>>> Performing Cluster Check (using node 127.0.0.1:6379)
M: 8e20e97a99e1abed4eb568079d63538439d39382 127.0.0.1:6379
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
M: fba8c2aefddd1cb4b694f7f29ade77b9309f0359 127.0.0.1:6381
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
S: c1a6f2b05266dd2c99a21e8ec715b5760828bd05 127.0.0.1:6384
   slots: (0 slots) slave
   replicates fba8c2aefddd1cb4b694f7f29ade77b9309f0359
S: 6750e7cf044954b21381c074d0a3a89f8d9e211b 127.0.0.1:6383
   slots: (0 slots) slave
   replicates 57c53adcc0a7d8ddcc79487a7386837b364044fb
M: 57c53adcc0a7d8ddcc79487a7386837b364044fb 127.0.0.1:6380
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
S: ae1e820f675956a281404de626a7e2194bede899 127.0.0.1:6382
   slots: (0 slots) slave
   replicates 8e20e97a99e1abed4eb568079d63538439d39382
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
Automatically selected master 127.0.0.1:6379
>>> Send CLUSTER MEET to node 127.0.0.1:6385 to make it join the cluster.
Waiting for the cluster to join

>>> Configure node as replica of 127.0.0.1:6379.
[OK] New node added correctly.
```

查看节点 6379 的信息, 显示 2 个从节点

```shell
# 查看节点信息
[root@centos7 redis-cluster]# redis-cli --cluster info 127.0.0.1:6379
127.0.0.1:6379 (8e20e97a...) -> 0 keys | 5461 slots | 2 slaves.
127.0.0.1:6381 (fba8c2ae...) -> 0 keys | 5461 slots | 1 slaves.
127.0.0.1:6380 (57c53adc...) -> 0 keys | 5462 slots | 1 slaves.
[OK] 0 keys in 3 masters.
0.00 keys per slot on average.
```

#### 集群优点

- 实现扩容
- 分担压力
- 无中心配置相对简单

#### 集群缺点

- 多建操作不支持, 例如 `MSET` 命令设置多个键不支持, 需要使用分组方式 `MSET name{user} zhangsan age{user} 20 addr{user} beijing`
- 多键的 Redis 事务不支持
- Lua 脚本不支持
- 迁移方案需要整体迁移而不是逐步过渡, 复杂度较大

### 慢查询

Redis 慢查询和 Redis 定义慢查询的 `阈值` 有关

`slowlog-log-slower-than 10000` 单位微秒, 当 Redis 命令的执行时间超过该值时, Redis 将其记录在 Redis 的慢查询日志中
`slowlog-max-len 128` 记录的条数超过时会只存储最新的 slowlog-max-len 条

#### 使用复杂度过高的命令

复杂的命令一般指 O(N)以上的命令, 如 sort、sunion、zunionstore 聚合类的命令, 或是 O(N)类的命令, 对于 O(N)以上的命令, Redis 在操作内存数据时耗时过高, 会耗费更多的 CPU 资源, 导致查询变慢
Redis 是单线程处理客户端请求的, 如果遇到处理上面的请求时, 就会导致后面的请求发生排队, 对于客户端来说响应时间就会变长

解决问题的原则

- 尽量不使用 O(N)以上的命令, 某些数据需要排序或者聚合操作时, 可以放在客户端处理
- 执行 O(N)命令时, 保证 N 尽量的小(推荐 N <= 300), 每次获取尽量少的数据, 让 Redis 可以及时处理返回

#### 大 Key 问题

通常是 key 对应的 value 值过大, 此类问题在 SET/DEL 这类命令中也会出现慢查询
SET/DEL 的过程

- 写入数据: 为该数据分配内存空间
- 删除数据: 释放该数据对应的内存空间

当数据值较大时, Redis 分配数据内存和释放内存空间都比较耗时

解决问题的原则

- 尽量避免写入大 Key(不要写入无关的数据, 数据实在过大进行拆分, 通过多 key 存储)
- 如果 Redis 是 4.0 以上版本, 尽量使用 `UNLINK`代替 `DEL`命令, 此命令将删除 key 和内存回收放到其他线程执行, 从而降低对 Redis 的影响
- 如果 Redis 是 6.0 以上版本, 可以开启 lazy-free, 在执行 DEL 命令时、释放内存也会放到其他线程中执行

`lazyfree-lazy-user-del no` 修改 `DEL` 默认命令的行为使其更接近于 `UNLINK`命令, 默认不开启

#### 集中过期

Redis 过期策略

- 被动过期: 只有当访问某个 key 时, 才会检测该 key 是否已经过期, 如果已经过期则从实例删除该 key
- 主动过期: Redis 内部存在一个定时任务, 默认每间隔 100 毫秒就会从全局的过期哈希表中随机取出 20 个 key, 然后删除其中过期的 key, 如果过期 key 的比例超过了 25%, 则继续重复此过程, 直到过期 key 的比例下降到 25% 以下, 或者这次任务的执行耗时超过了 25 毫秒, 才会退出循环

主动过期 key 的定时任务是在 Redis 主线程中执行的, 如果在执行主动过期的过程中, 出现了集中过期, 就需要大量删除过期 key, 如果此时应用程序在访问 Redis 时, 必须等待这个过期任务执行结束, 此时 Redis 就有可能产生慢查询

解决问题的原则

- 避免集中过期, 比如将过期时间随机化, 添加一个随机的值, 分散集中过期 key 的过期时间, 降低 Redis 清理过期 key 的压力
- 如果 Redis 是 4.0 以上版本, 可以开启 lazy-free, 当删除过期 key 时, 把释放内存的操作放到其他线程中执行, 避免阻塞主线程
