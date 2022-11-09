---
title: Redis
date: 2022-11-09 18:03:19
categories:
  - [server, Redis]
tags:
  - Redis
---

### Redis

Remote Dictionary Server 即远程字典服务, 是一个开源的使用 ANSI C 语言编写、支持网络、可基于内存亦可持久化的日志型、Key-Value 数据库，并提供多种语言的 API

#### Keys 命令

- DEL key [key...] 删除 key, 1 成功, 0 失败
- DUMP key 序列化指定 key, 并返回被序列化的值, 不存在返回 &lt;nil&gt;
- EXISTS key 检查指定 key 是否存在, 1 存在, 0 不存在
- EXPIRE key seconds 为指定 key 设置过期时间(单位秒), 1 成功, 0 失败
- EXPIREAT key unix-time-seconds 为指定 key 设置过期使用 unix 时间戳, 1 成功, 0 失败
- PEXPIRE key milliseconds 为指定 key 设置过期时间(单位毫秒), 1 成功, 0 失败
- PEXPIREAT key unix-time-milliseconds 为指定 key 设置过期时间使用 unix 时间戳, 1 成功, 0 失败
- KEYS pattern 查找给定模式(pattern)的 key, 返回列表
- MOVE key db 将当前数据库中的 key 移动到指定的数据库(db)中
- PERSIST key 移除指定 key 的过期时间, key 将永久保持, 1 成功, 0 失败
- PTTL key 以毫秒为单位返回指定 key 的剩余的过期时间
  - \-2 key 不存在
  - \-1 key 存在但没有设置剩余生存时间
- TTL key 以秒为单位返回指定 key 的剩余生存时间
  - \-2 key 不存在
  - \-1 key 存在但没有设置剩余生存时间
- RANDOMKEY 从当前数据库随机返回一个 key, 如果当前数据库为空则返回 &lt;nil&gt;
- RENAME key newKey 修改 key 的名称, ok 成功, ERR no such key 失败
- TYPE key 返回指定 key 的类型, none 表示 key 不存在

#### CONFIG 命令

- CONFIG GET parameter [parameter...] 获取指定配置项的值
- CONFIG HELP 显示 CONFIG 命令的帮助信息
- CONFIG RESETSTAT 重置 INFO 返回的统计信息, ok 成功
- CONFIG REWRITE 使用内存配置重写配置文件
- CONFIG SET parameter value [parameter value ...] 设置配置项
