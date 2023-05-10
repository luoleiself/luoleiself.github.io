---
title: Redis-String-Hash
date: 2022-11-09 18:03:19
categories:
  - [server, Redis]
tags:
  - Redis
---

### Strings 命令

字符串是基础类型, 存储字节序列, 包括文本、序列化对象和二进制数组, 一个 key 对应一个 value, value 可以是字符串、整数或浮点数, value 最多可以是 512MB.

Strings: 键名: key, 键类型: string, 键值: string

#### 设置值

- SET key value [NX|XX] [GET] [EX seconds|PX milliseconds|EXAT unix-time-seconds|PXAT unix-time-milliseconds|KEEPTTL]

  为 key 设置字符串的值, 执行成功返回 ok

  - NX 仅当 key 不存在时设置
  - XX 仅当 key 存在时设置
  - EX 过期时间, 单位秒
  - PX 过期时间, 单位毫秒
  - EXAT 过期时间戳, 单位秒
  - PXAT 过期时间戳, 单位毫秒
  - KEEPTTL 保留 key 关联的生存时间

- SETNX key value 当 key 不存在时设置指定 key 的值, 1 成功, 0 失败

```shell
127.0.0.1:6379> KEYS *
1) "xiaoming"
2) "name"
127.0.0.1:6379> SETNX age 18
(integer) 1
127.0.0.1:6379> SETNX age 18
(integer) 0
127.0.0.1:6379> KEYS *
1) "xiaoming"
2) "age"
3) "name"
```

- APPEND key value 在指定 key 末尾(如果为字符串)追加内容, key 不存在同 `SET` 并返回追加内容的长度

```shell
127.0.0.1:6379> APPEND age 1
(integer) 3
127.0.0.1:6379> GET age
"181"
127.0.0.1:6379> APPEND addr beijing
(integer) 7
127.0.0.1:6379> KEYS *
1) "xiaoming"
2) "age"
3) "addr"
4) "name"
127.0.0.1:6379> APPEND a hello
(integer) 5
127.0.0.1:6379> APPEND b gg
(integer) 2
```

##### 批量设置值

- MSET key value [key value ...] 批量设置 key 的值
- MSETNX key value [key value ...] 批量设置 key 的值且当所有的 key 不存在时, 1 成功, 0 失败

- SETRANGE key offset value

  覆盖指定 key 的从指定偏移量开始的字符串的一部分, 返回修改后字符串长度, key 不存在则新建

##### 过期时间

- SETEX key seconds value 设置 key 的值并设置过期时间(单位秒), 返回 ok
- PSETEX key milliseconds value 设置 key 的值的值并设置过期时间(单位毫秒), 返回 ok

<!-- more -->

#### 获取值

- GET key 获取一个 key 的值, 不存在返回 \<nil\>

- GETSET key value 设置指定 key 的值并返回原来的值, key 不存在返回 \<nil\>

```shell
127.0.0.1:6379> GETSET age 18
(nil)
127.0.0.1:6379> KEYS *
1) "name"
2) "age"
```

- GETEX key [EX seconds|PX milliseconds|EXAT unix-time-seconds|PXAT unix-time-milliseconds|PERSIST]

  获取指定 key 的值并设置过期时间, key 不存在返回 \<nil\>

  - PERSIST 移除 key 关联的生存时间

- GETDEL key 获取指定 key 的值并删除, key 不存在返回 \<nil\>

```shell
127.0.0.1:6379> GETDEL age
"18"
127.0.0.1:6379> KEYS *
1) "name"
```

- STRLEN key 返回指定 key 的值的长度, key 不存在返回 0

##### 批量获取值

- MGET key [key ...] 批量获取 key 的值, key 不存在返回 \<nil\>

- GETRANGE key start end 返回指定 key 的指定范围的子串部分, key 不存在返回 `""`
  - start, end 只支持整数, 其他类型会报错

```shell
127.0.0.1:6379> GETRANGE name -inf +inf
(error) ERR value is not an integer or out of range
127.0.0.1:6379> GETRANGE name - +
(error) ERR value is not an integer or out of range
127.0.0.1:6379> GETRANGE name (1 (4
(error) ERR value is not an integer or out of range
127.0.0.1:6379> GETRANGE name [1 [4
(error) ERR value is not an integer or out of range
127.0.0.1:6379> GETRANGE name 1 4
"uole"
```

- SUBSTR key start end 返回指定 key 的指定范围的子串部分, key 不存在返回 `""`
  - start, end 只支持整数, 其他类型会报错

```shell
127.0.0.1:6379> SUBSTR name -inf +inf
(error) ERR value is not an integer or out of range
127.0.0.1:6379> SUBSTR name (1 (4
(error) ERR value is not an integer or out of range
```

#### 数值操作

##### 增加

- INCR key 将 key 中存储的数字值增加 1 并返回修改后的值, 非数字值或值为浮点数报错, key 不存在从 0 开始计算
- INCRBY key increment 将 key 中存储的数字值加上给定的增量值(increment), 返回值同 `INCR`
- INCRBYFLOAT key increment 将 key 中存储的数字值加上给定的浮点增量值(increment), 返回值同 `INCR`

##### 减少

- DECR key 将 key 中存储的数字值减 1 并返回修改后的值, 非数字值或者值为浮点数会报错, key 不存在从 0 开始计算
- DECRBY key decrement 将 key 中存储的数字值减去给定的增量值(decrement), 返回值同 `DECR`

### Hashes 命令

hash 是一个 string 类型的 field(字段) 和 value(值)的映射表, hash 适合用于存储对象, 每个 hash 可以存储 2^32-1(40 多亿)键值对

Hashes: 键名: key, 键类型: hash, 键值: string {field => value}

#### 哈希表存取

- HEXISTS key field 查看哈希表中是否存在 field, 1 表示存在, 0 表示不存在或者哈希表不存在

- HDEL key field [field ...] 批量删除多个 field 并返回删除字段成功的数量, 0 表示 field 未找到或者哈希表不存在

- HSET key field value [field value ...] 同时将多个键值对存入到哈希表中并返回新添加的数量, 如果 field 已存在则修改 field 的值

```shell
# 成功添加一个 field 并修改已存在的 field
127.0.0.1:6379> HSET runoob name "new-redis" age 19 addr "beijing" sex "男"
(integer) 3
```

- HMSET key field value [field value ...] 批量向哈希表中存入多个键值对, 如果 field 存在则修改 field 的值, 执行成功返回 ok

```shell
127.0.0.1:6379> HMSET runoob name redis newname "new-redis" age 18 addr "beijing" sex "男"
OK
```

- HSETNX key field value 将键值对存入到哈希表中且当指定 field 不存在时, 1 成功, 0 失败(字段已存在)

- HGET key field 获取哈希表指定 field 的值, field 或者 哈希表不存在返回 \<nil\>
- HMGET key field [field ...] 批量获取哈希表中指定 field 的值, 哈希表或者指定字段不存在返回 \<nil\>
- HGETALL key 获取哈希表中所有的字段和值, 未找到或者哈希表不存在返回 (empty array)

```shell
127.0.0.1:6379> HGET xiaoming name
"xiaoming"
127.0.0.1:6379> HMGET xiaoming name age addr
1) "xiaoming"
2) "1"
3) "beijing"
127.0.0.1:6379> HGETALL xiaoming
1) "name"
2) "xiaoming"
3) "age"
4) "1"
5) "addr"
6) "beijing"
```

#### 获取哈希表的键、值、长度

- HLEN key 获取哈希表中字段的数量, 0 表示哈希表为空或者不存在
- HSTRLEN key field 返回哈希表中指定 field 的值的字符串长度, 哈希表或者指定字段不存在返回 0

```shell
127.0.0.1:6379> HLEN xiaoming
(integer) 3
127.0.0.1:6379> HSTRLEN xiaoming name
(integer) 8
```

- HKEYS key 获取哈希表中所有的字段, 哈希表为空或者不存在返回 (empty array)
- HVALS key 获取哈希表中所有的值, 哈希表为空或者不存在返回 (empty array)

```shell
127.0.0.1:6379> HKEYS xiaoming
1) "name"
2) "age"
3) "addr"
127.0.0.1:6379> HVALS xiaoming
1) "xiaoming"
2) "1"
3) "beijing"
```

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

- HRANDFIELD key [count [WITHVALUES]] 从哈希表中获取一个或多个随机字段, 哈希表为空返回 \<nil\>
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

#### 使用案例

##### 存储用户的基本信息

以用户 uid 作为 key, 用户的基本信息组成的 field => value 作为值, 使用 hash 存储

key 的格式: {uid}
value 的格式: {age => 18}, {addr => beijing}
