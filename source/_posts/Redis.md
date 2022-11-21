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

#### 操作 key

- TYPE key 返回指定 key 的类型, none 表示 key 不存在
- EXISTS key [key ...] 检查指定 key 是否存在, 1 存在, 0 不存在
- KEYS pattern 查找给定模式(pattern)的 key, 返回列表, 未找到返回 (empty array)
- DEL key [key...] 阻塞删除 key 并返回成功删除 key 的数量
- UNLINK key [key ...] 非阻塞从键空间中取消键指定 key 的链接, 并返回成功取消 key 的数量, 如果 key 不存在则忽略

- RENAME key newKey 修改 key 的名称, 如果指定 key 不存在返回 错误, 如果 newkey 已存在则覆盖
- RENAMENX key newkey 修改 key 的名称, 如果指定 key 不存在返回 错误, 如果 newkey 已存在不执行任何操作返回 0, 否则返回 1

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

### ACL

ACL(access control list)访问控制列表的简称, 是为了控制某些 Redis 客户端在访问 Redis 服务器时, 能够执行的命令和能够获取的 key, 提高操作安全性, 避免对数据造成损坏

#### 规则分类

|          参数          | 说明                                                                                    |
| :--------------------: | --------------------------------------------------------------------------------------- |
|           on           | 表示启动该用户, 默认为 off                                                              |
|         nopass         | 删除所有与用户关联的密码                                                                |
|         reset          | 移除用户的所有功能, 并关闭用户                                                          |
|       +[command]       | 将命令添加到用户可以调用的命令列表中                                                    |
|       -[command]       | 将命令从用户可以调用的命令列表中移除                                                    |
| +[command]\|subcommand | 允许使用已禁用命令的特定子命令                                                          |
|      +@[category]      | 允许用户调用 category 类别中的所有命令, 可以使用 `ACL CAT` 命令查看所有类别             |
|      -@[category]      | 禁止用户调用 category 类别中的所有命令                                                  |
|      allcommands       | +@all 的别名                                                                            |
|       nocommands       | -@all 的别名                                                                            |
|      ~\<pattern\>      | 允许用户可以访问的 key(正则匹配), 例如: ~foo:\* 只允许访问 foo:\* 的 key                |
|     %R~\<pattern\>     | 添加指定的只读 key(正则匹配), 例如: %R~app:\* 只允许读 app:\* 的 key, 7.0 支持          |
|     %W~\<pattern\>     | 添加指定的只写 key(正则匹配), 例如: %W~app:\* 只允许写 app:\* 的 key, 7.0 支持          |
|    %RW~\<pattern\>     | 添加指定的可读可写的 key(正则匹配), 例如: %RW~app:\* 只允许读写 app:\* 的 key, 7.0 支持 |
|        allkeys         | ~\* 的别名                                                                              |
|       resetkeys        | 移除所有的 key 匹配模式                                                                 |
|      &\<pattern\>      | 允许用户可使用的 Pub/Sub 通道(正则匹配)                                                 |
|      allchannels       | &\* 的别名                                                                              |
|     resetchannels      | 移除所有的通道匹配模式                                                                  |
|     \>\<password\>     | 为用户添加明文密码, 服务器自动转换成 hash 存储, 例如: >123456                           |
|     \<\<password\>     | 从有效密码列表中删除密码                                                                |
|  #\<hashedpassword\>   | 为用户添加 hash 密码, 例如: #cab3...c4f2                                                |
|  \!\<hashedpassword\>  | 从有效密码列表中删除密码                                                                |

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

#### SECURITY

- acllog-max-len 128 # ACL 日志在内存中时的最大条目数
- aclfile /etc/redis/users.acl # 默认 ACL 配置文件
- requirepass foobared # 认证密码

- maxclients 10000 # 客户端最大连接数

- io-threads 4 # I/O 线程

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

`dbfilename dump.rdb` 默认文件名
`dir ./` 默认存储目录

- redis-check-rdb 检查 RDB 文件

#### AOF

AOF(Append Only File), 将执行过的写命令全部记录下来, 在数据恢复时按照从前往后的顺序再将指令都执行一遍, 默认 AOF 的持久化策略在开启后是每秒钟 sync 一次, 可以修改配置文件的 `appendfsync` 项

`appendonly yes` 启动 AOF 模式
`appendfilename appendonly.aof` 默认文件名
`appenddirname appendonlydir` 默认存储目录
`appendfsync everysec` 持久化策略, 每秒钟执行一次, 可以修改为 `always` 和 `no`

如果 appendonly.aof 文件有错误, Redis 服务将会启动失败

- redis-check-aof 检查 AOF 文件, --fix 参数修复文件的错误, 通常会丢弃文件中无法识别的命令

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
使用命令配置只能在本次服务器运行时有效, 重启服务器后将会丢失配置信息, 使用配置文件永久生效

- 启动 Redis 服务器时参数指定 `redis-server --port 6380 --replicaof 127.0.0.1 6379`
- 客户端连接服务器后使用内置命令 `REPLICAOF host port`

单机集群: 新建多个 Redis 服务器配置文件并修改其中关键项

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

#### 哨兵模式配置

哨兵模式是一种特殊的模式, 首先 Redis 提供了哨兵的命令, 哨兵是一个独立的进程, 作为进程, 它会独立运行

- 通过发送命令, 让 Redis 服务器返回监控其运行状态, 包括主服务器和从服务器
- 当哨兵监测到 master 宕机, 会自动将 slave 切换成 master, 然后通过发布订阅模式通知其他的从服务器, 修改配置文件, 让它们切换主机

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

- 使用命令 `redis-server /path/to/sentinel.conf --sentinel` 开启哨兵模式
- 使用命令 `redis-sentinel /path/to/sentinel.conf` 开启哨兵模式

```shell
# sentinel.conf
sentinel monitor myredis 127.0.0.1 6379 1
```

一主三从哨兵配置

```shell
# sentinel.conf
port 26379
pidfile /var/run/redis-sentinel-26379.pid
logfile "26379.log"
dir /tmp
sentinel monitor myredis 127.0.0.1 6379 2

# sentinel.conf
port 36379
pidfile /var/run/redis-sentinel-36379.pid
logfile "36379.log"
dir /tmp
sentinel monitor myredis 127.0.0.1 6379 2

# sentinel.conf
port 46379
pidfile /var/run/redis-sentinel-46379.pid
logfile "46379.log"
dir /tmp
sentinel monitor myredis 127.0.0.1 6379 2
```

![redis-2](/images/redis-2.png)
![redis-3](/images/redis-3.png)
