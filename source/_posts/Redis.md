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

### 配置文件

- daemonize no 是否后台模式启动服务

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

<!-- more -->

### Keys 命令

- HELP command 显示命令的帮助信息
- TYPE key 返回指定 key 的类型, none 表示 key 不存在
- DEL key [key...] 删除 key 并返回删除 key 的数量
- DUMP key 序列化指定 key, 并返回被序列化的值, 不存在返回 &lt;nil&gt;
- RENAME key newKey 修改 key 的名称, ok 成功, ERR no such key 失败
- MOVE key db 将当前数据库中的 key 移动到指定的数据库(db)中

- EXISTS key 检查指定 key 是否存在, 1 存在, 0 不存在
- KEYS pattern 查找给定模式(pattern)的 key, 返回列表, 未找到返回 (empty array)

- SHUTDOWN [NOSAVE|SAVE] [NOW] [FORCE] [ABORT] 同步保存数据到硬盘上并关闭服务

#### 设置 key 的过期时间

- EXPIRE key seconds 为指定 key 设置过期时间(单位秒), 1 成功, 0 失败
- EXPIREAT key unix-time-seconds 为指定 key 设置过期使用 unix 时间戳, 1 成功, 0 失败
- PEXPIRE key milliseconds 为指定 key 设置过期时间(单位毫秒), 1 成功, 0 失败
- PEXPIREAT key unix-time-milliseconds 为指定 key 设置过期时间使用 unix 时间戳, 1 成功, 0 失败

#### 获取 key 的过期时间

- PTTL key 以毫秒为单位返回指定 key 的剩余的过期时间
  - \-2 key 不存在
  - \-1 key 存在但没有设置剩余生存时间
- TTL key 以秒为单位返回指定 key 的剩余生存时间

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

### CONFIG 命令

- CONFIG GET parameter [parameter...] 获取指定配置项的值
- CONFIG HELP 显示 CONFIG 命令的帮助信息
- CONFIG RESETSTAT 重置 INFO 返回的统计信息, ok 成功
- CONFIG REWRITE 使用内存配置重写配置文件
- CONFIG SET parameter value [parameter value ...] 设置配置项
