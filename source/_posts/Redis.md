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

- HELP command 显示命令的帮助信息
- TYPE key 返回指定 key 的类型, none 表示 key 不存在
- EXISTS key [key ...] 检查指定 key 是否存在, 1 存在, 0 不存在
- KEYS pattern 查找给定模式(pattern)的 key, 返回列表, 未找到返回 (empty array)
- DEL key [key...] 阻塞删除 key 并返回成功删除 key 的数量
- UNLINK key [key ...] 非阻塞从键空间中取消键指定 key 的链接, 并返回成功取消 key 的数量, 如果 key 不存在则忽略
- DUMP key 序列化指定 key, 并返回被序列化的值, 不存在返回 &lt;nil&gt;
- RENAME key newKey 修改 key 的名称, 如果指定 key 不存在返回 错误, 如果 newkey 已存在则覆盖
- RENAMENX key newkey 修改 key 的名称, 如果指定 key 不存在返回 错误, 如果 newkey 已存在不执行任何操作返回 0, 否则返回 1

- TOUCH key [key ...] 更改指定 key 的最后一次访问时间并返回修改成功的数量, 如果 key 不存在则忽略

- WAIT numreplicas timeout 阻止当前客户端, 直到所有先前的写入命令成功传输并至少由指定数量的副本确认, 如果达到了以毫秒为单位指定的超时, 则即使尚未达到指定的副本数量, 该命令也会返回

- MOVE key db 将当前数据库中的 key 移动到指定的数据库(db)中
- ECHO message 打印信息
- PING [message] 测试连接是否正常, 通常返回 PONG, 如果传入了 message 则会输出 message
- QUIT 关闭退出当前连接
- SHUTDOWN [NOSAVE|SAVE] [NOW] [FORCE] [ABORT] 同步保存数据到硬盘上并关闭服务

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

##### ACL

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
- aclfile /etc/redis/users.acl # ACL 日志文件
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
Redis 客户端可以订阅任意数量的频道
