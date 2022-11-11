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

### Keys 命令

- HELP command 显示命令的帮助信息
- DEL key [key...] 删除 key 并返回删除 key 的数量
- DUMP key 序列化指定 key, 并返回被序列化的值, 不存在返回 &lt;nil&gt;
- EXISTS key 检查指定 key 是否存在, 1 存在, 0 不存在
- KEYS pattern 查找给定模式(pattern)的 key, 返回列表, 未找到返回 (empty array)

- EXPIRE key seconds 为指定 key 设置过期时间(单位秒), 1 成功, 0 失败
- EXPIREAT key unix-time-seconds 为指定 key 设置过期使用 unix 时间戳, 1 成功, 0 失败
- PEXPIRE key milliseconds 为指定 key 设置过期时间(单位毫秒), 1 成功, 0 失败
- PEXPIREAT key unix-time-milliseconds 为指定 key 设置过期时间使用 unix 时间戳, 1 成功, 0 失败

- MOVE key db 将当前数据库中的 key 移动到指定的数据库(db)中
- PERSIST key 移除指定 key 的过期时间, key 将永久保持, 1 成功, 0 失败
- PTTL key 以毫秒为单位返回指定 key 的剩余的过期时间
  - \-2 key 不存在
  - \-1 key 存在但没有设置剩余生存时间
- TTL key 以秒为单位返回指定 key 的剩余生存时间
  - \-2 key 不存在
  - \-1 key 存在但没有设置剩余生存时间

<!-- more -->

- RANDOMKEY 从当前数据库随机返回一个 key, 如果当前数据库为空则返回 &lt;nil&gt;
- RENAME key newKey 修改 key 的名称, ok 成功, ERR no such key 失败
- TYPE key 返回指定 key 的类型, none 表示 key 不存在
- SWAPDB index1 index2 切换两个数据库
- SELECT index 更改当前连接的选定的数据库
- DBSIZE 返回当前数据库中 key 的数量

### CONFIG 命令

- CONFIG GET parameter [parameter...] 获取指定配置项的值
- CONFIG HELP 显示 CONFIG 命令的帮助信息
- CONFIG RESETSTAT 重置 INFO 返回的统计信息, ok 成功
- CONFIG REWRITE 使用内存配置重写配置文件
- CONFIG SET parameter value [parameter value ...] 设置配置项

### 字符串命令

字符串是基础类型, 存储字节序列, 包括文本、序列化对象和二进制数组, 一个 key 对应一个 value, value 可以是字符串、整数或浮点数, value 最多可以是 512MB

#### 设置值

- SET key value [NX|XX] [GET] [EX seconds|PX milliseconds|EXAT unix-time-seconds|PXAT unix-time-milliseconds|KEEPTTL] 为 key 设置字符串的值

- SETNX key value 设置指定 key 的值且当 key 不存在时, 1 成功, 0 失败

- APPEND key value 在指定 key 末尾(如果为字符串)追加内容, key 不存在同 `SET`

##### 过期时间

- SETEX key seconds value 设置 key 的值并设置过期时间(单位秒), 返回 ok

- PSETEX key milliseconds value 设置 key 的值的值并设置过期时间(单位毫秒), 返回 ok

##### 批量设置值

- MSET key value [key value ...] 批量设置 key 的值

- MSETNX key value [key value ...] 批量设置 key 的值且当 key 不存在时, 1 成功, 0 失败

- SETRANGE key offset value 覆盖指定 key 的从指定偏移量开始的字符串的一部分, 返回修改后字符串长度, key 不存在则新建

#### 获取值

- GET key 获取一个 key 的值, 不存在返回 &lt;nil&gt;

- GETSET key value 设置指定 key 的值并返回原来的值, key 不存在返回 &lt;nil&gt;

- GETEX key [EX seconds|PX milliseconds|EXAT unix-time-seconds|PXAT unix-time-milliseconds|PERSIST] 获取指定 key 的值并设置过期时间, key 不存在返回 &lt;nil&gt;

- GETDEL key 获取指定 key 的值并删除, key 不存在返回 &lt;nil&gt;

- STRLEN key 返回指定 key 的长度, key 不存在返回 0

##### 批量获取值

- MGET key [key ...] 批量获取 key 的值, key 不存在返回 &lt;nil&gt;

- GETRANGE key start end 返回指定 key 的指定范围的子串部分, key 不存在返回 `""`

- SUBSTR key start end 返回指定 key 的指定范围的子串部分, key 不存在返回 `""`

#### 数值操作

##### 增加

- INCR key 将 key 中存储的数字值增加 1 并返回修改后的值, 非数字值报错, key 不存在从 0 开始计算
- INCRBY key increment 将 key 中存储的数字值加上给定的增量值(increment), 返回值同 `INCR`
- INCRBYFLOAT key increment 将 key 中存储的数字值加上给定的浮点增量值(increment), 返回值同 `INCR`

##### 减少

- DECR key 将 key 中存储的数字值减 1 并返回修改后的值, 非数字值或者值为浮点数会报错, key 不存在从 0 开始计算
- DECRBY key decrement 将 key 中存储的数字值减去给定的增量值(decrement), 返回值同 `DECR`

### Hash 命令

hash 是一个 string 类型的 field(字段) 和 value(值)的映射表, hash 适合用于存储对象, 每个 hash 可以存储 2^32-1(40 多亿)键值对

#### 哈希表存取

- HEXISTS key field 查看哈希表中是否存在 field, 1 表示存在, 0 表示不存在或者哈希表不存在

- HDEL key field [field ...] 批量删除多个 field 并返回删除字段成功的数量, 0 表示 field 未找到或者哈希表不存在

- HSET key field value [field value ...] 同时将多个键值对存入到哈希表中并返回新添加的数量, 如果 field 已存在则修改 field 的值

```shell
127.0.0.1:6379> HSET runoob name "new-redis" age 19 addr "beijing" sex "男"
(integer) 3 # 成功添加一个 field 并修改已存在的 field
```

- HMSET key field value [field value ...] 批量向哈希表中存入多个键值对, 如果 field 存在则修改 field 的值, 执行成功返回 ok

```shell
127.0.0.1:6379> HMSET runoob name redis newname "new-redis" age 18 addr "beijing" sex "男"
OK
```

- HSETNX key field value 将键值对存入到哈希表中且当指定 field 不存在时, 1 成功, 0 失败(字段已存在)

- HGET key field 获取哈希表指定 field 的值, field 或者 哈希表不存在返回 &lt;nil&gt;
- HMGET key field [field ...] 批量获取哈希表中指定 field 的值, 哈希表或者指定字段不存在返回 &lt;nil&gt;
- HGETALL key 获取哈希表中所有的字段和值, 未找到或者哈希表不存在返回 (empty array)

#### 获取哈希表的键、值、长度

- HLEN key 获取哈希表中字段的数量, 0 表示哈希表为空或者不存在
- HSTRLEN key field 返回哈希表中指定 field 的值的字符串长度, 哈希表或者指定字段不存在返回 0
- HKEYS key 获取哈希表中所有的字段, 哈希表为空或者不存在返回 (empty array)
- HVALS key 获取哈希表中所有的值, 哈希表为空或者不存在返回 (empty array)

#### 哈希表字段数值操作

- HINCRBY key field increment 为哈希表中指定的 field 的数字值加上给定的增量值(increment)并返回修改后的值, 非数字值报错, 哈希表不存在新建, 字段不存在从 0 开始计算
- HINCRBYFLOAT key field increment 为哈希表中指定的 field 的数字值加上给定的浮点数增量值(increment)并返回修改后的值, 非数字值报错, 哈希表不存在新建, 字段不存在从 0 开始计算

#### 迭代哈希表

- HSCAN key cursor [MATCH pattern] [COUNT count] 使用模式(pattern)匹配迭代哈希表中的键值对
  - cursor 游标
  - pattern 匹配的模式
  - count 指定从数据集里返回多少元素, 默认为 10

```shell
127.0.0.1:6379> HSCAN runoob 0 MATCH *name COUNT 100
1) "0"
2) 1) "name"
   2) "redis"
   3) "newname"
   4) "new-redis"
```

#### 获取哈希表随机字段

- HRANDFIELD key [count [WITHVALUES]] 从哈希表中获取一个或多个随机字段, 哈希表为空返回 &lt;nil&gt;
  - count 指定返回随机的字段的数量, 默认为 1
  - WITHVALUES 指定返回随机的字段和值

```shell
127.0.0.1:6379> HRANDFIELD runoob
"age"
127.0.0.1:6379> HRANDFIELD runoob 3
1) "name"
2) "addr"
3) "sex"
127.0.0.1:6379> HRANDFIELD runoob 3 WITHVALUES
1) "newname"
2) "new-redis"
3) "addr"
4) "beijing"
5) "sex"
6) "\xe7\x94\xb7"
127.0.0.1:6379> HRANDFIELD newrunoob
(nil)
```
