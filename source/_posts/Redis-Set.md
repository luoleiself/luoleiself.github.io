---
title: Redis-Set
date: 2022-11-09 18:03:19
categories:
  - [server, Redis]
tags:
  - Redis
---

### Set 命令

Set 是无序不重复的集合, 集合成员是唯一的, 集合对象的编码可以是 intset 或者 hashtable, 集合是通过哈希表实现的, 最大的成员数为 2^32-1(40 多亿)个成员.

#### 成员操作

- SADD key member [member ...] 向集合中添加多个成员并返回添加成功的数量, 0 表示有重复成员

- SCARD key 获取集合成员的数量, 集合为空或者不存在返回 0

- SMEMBERS key 获取集合中所有的成员, 集合为空或者不存在返回 (empty array)

##### 是否包含成员

- SISMEMBER key member 判断 member 是不是集合的成员, 1 是, 0 不是或者集合为空或者不存在
- SMISMEMBER key member [member ...] 批量判断多个 member 是不是集合的成员, 1 是, 0 不是或者集合为空或者不存在, 6.2.0 支持

```shell
127.0.0.1:6379> KEYS *    # 查看当前数据库中的 key
(empty array)
127.0.0.1:6379> SMISMEMBER myset hello hehe    # 空集合判断是否包含成员
1) (integer) 0
2) (integer) 0
127.0.0.1:6379> SADD myset hello world hehe haha gg    # 向 myset 添加成员
(integer)
127.0.0.1:6379> SMISMEMBER myset hello yy hehe    # 判断 myset 是否包含成员
1) (integer) 1
2) (integer) 0
3) (integer) 1
```

<!-- more -->

##### 批量移除成员

- SREM key member [member ...] 批量移除集合中的成员并返回移除成功的数量, 集合为空或者不存在或者不包含移除成员返回 0

```shell
127.0.0.1:6379> KEYS *  # 查看当前数据库所有的 key
(empty array)
127.0.0.1:6379> SADD set hello world gg  # 向集合添加成员
(integer) 3
127.0.0.1:6379> SMEMBERS set  # 查看集合的所有成员
1) "gg"
2) "world"
3) "hello"
127.0.0.1:6379> SREM myset a b  # 移除不存在的集合的成员
(integer) 0
127.0.0.1:6379> SREM set a b  # 移除集合不存在的成员
(integer) 0
127.0.0.1:6379> SREM set hello a  # 移除集合的成员
(integer) 1
127.0.0.1:6379> SMEMBERS set  # 查看集合的所有成员
1) "gg"
2) "world"
127.0.0.1:6379> SREM set world gg  # 移除集合的成员
(integer) 2
127.0.0.1:6379> SMEMBERS key  # 查看集合为空的所有成员
(empty array)
```

##### 移除成员添加到其他集合

- SMOVE source destination member 将源集合中的指定成员移除并添加到指定集合中, 1 表示源集合指定成员移除成功(目标集合中可能包含该成员也可能不包含), 0 表示源集合为空或者源集合不包含指定成员

```shell
127.0.0.1:6379> FLUSHALL    # 清空所有数据库
OK
127.0.0.1:6379> KEYS *    # 查看当前数据库中的 key
(empty array)
127.0.0.1:6379> SMOVE myset destset hello    # myset 为空集, 返回 0
(integer) 0

127.0.0.1:6379> SADD myset hello world    # 向 myset 添加成员 hello, world
(integer) 2
127.0.0.1:6379> SMOVE myset destset hello    # 移除 myset 包含的成员 hello 并添加到 destset 中, destset 为空, 返回 1
(integer) 1
127.0.0.1:6379> SMEMBERS myset    # 查看 myset 成员
1) "world"
127.0.0.1:6379> SMEMBERS destset    # 查看 destset 成员
1) "hello"
127.0.0.1:6379> SMOVE myset destset hehe    # 移除 myset 不包含的成员 hehe, 返回 0
(integer) 0
127.0.0.1:6379> SMEMBERS myset    # 查看 myset 成员
1) "world"
127.0.0.1:6379> SMEMBERS destset    # 查看 destset 成员
1) "hello"

127.0.0.1:6379> FLUSHALL    # 清空所有数据库
OK
127.0.0.1:6379> KEYS *    # 查看当前数据库中的 key
(empty array)
127.0.0.1:6379> SADD myset hello world    # 向 myset 添加成员 hello, world
(integer) 2
127.0.0.1:6379> SADD destset hello gg    # 向 destset 添加成员 hello, gg
(integer) 2
127.0.0.1:6379> SMOVE myset destset hello    # 移除 myset 包含的成员 hello 并添加到 destset 中, destset 包含 hello, 返回 1
(integer) 1
127.0.0.1:6379> SMEMBERS myset
1) "world"
127.0.0.1:6379> SMEMBERS destset
1) "gg"
2) "hello"
```

#### 随机获取成员

- SPOP key [count] 移除指定集合随机的多个成员并返回移除的成员, 改变原集合, 不带 count 如果集合为空或者不存在返回 &lt;nil&gt;, 否则返回 (empty array), count 不能为负数
  - count 指定随机移除的数量, 默认为 1
    - count >= 1 时, 空集合返回 (empty array)
    - count = 0 时, 任何集合都返回 (empty array)
    - count < 0 时, 返回错误, count must be positive

```shell
127.0.0.1:6379> FLUSHALL    # 清空所有数据库
OK
127.0.0.1:6379> KEYS *    # 查看当前数据库中的 key
(empty array)

127.0.0.1:6379> SPOP myset    # myset 为空集合, 返回 nil
(nil)
127.0.0.1:6379> SPOP myset 2  # myset 为空集合, count >= 0 返回 empty array
(empty array)
127.0.0.1:6379> SADD myset hello world hehe haha gg yy    # 向 myset 添加成员 hello, world, hehe, haha, gg, yy
(integer) 6
127.0.0.1:6379> SPOP myset 0    # count = 0, 任意集合都返回 empty array
(empty array)
127.0.0.1:6379> SPOP newset 0    # count = 0, 任意集合都返回 empty array
(empty array)
127.0.0.1:6379> SPOP myset -1    # count < 0, 返回错误
(error) ERR value is out of range, must be positive
127.0.0.1:6379> SPOP myset    # 移除 myset 随机的 1 个成员
"world"
127.0.0.1:6379> SPOP myset 3    # 移除 myset 随机的 3 个成员
1) "haha"
2) "yy"
3) "gg"
127.0.0.1:6379> SMEMBERS myset    # 查看 myset 成员
1) "hehe"
2) "hello"
```

- SRANDMEMBER key [count] 返回指定集合随机的多个成员, 不改变原集合, 不带 count 如果集合为空或者不存在返回 &lt;nil&gt;, 否则返回 (empty array), count 为负数将会取绝对值
  - count 指定随机返回的数量, 默认为 1
    - count >= 1 时, 空集合返回 (empty array)
    - count = 0 时, 任意集合都返回 (empty array)
    - count < 0 时, 空集合返回 (empty array), 非空集合返回 count 的绝对值数量

```shell
127.0.0.1:6379> KEYS *    # 查看当前数据库中的 key
1) "myset"
127.0.0.1:6379> SMEMBERS myset    # 查看 myset 成员
1) "haha"
2) "yy"
3) "world"
4) "hehe"
5) "hello"
6) "gg"

127.0.0.1:6379> SRANDMEMBER newset    # newset 为空集合返回 nil
(nil)
127.0.0.1:6379> SRANDMEMBER newset 0    # newset 为空集合, count = 0, 返回 empty array
(empty array)
127.0.0.1:6379> SRANDMEMBER newset 1    # newset 为空集合, count = 1, 返回 empty array
(empty array)
127.0.0.1:6379> SRANDMEMBER newset -1    # newset 为空集合, count = -1, 返回 empty array
(empty array)

127.0.0.1:6379> SRANDMEMBER myset    # 随机返回 myset 集合的 1 个成员
"hello"
127.0.0.1:6379> SRANDMEMBER myset 0    # myset 不为空集合, count = 0, 返回 empty array
(empty array)
127.0.0.1:6379> SRANDMEMBER myset 2    # myset 不为空集合, count = 2, 随机返回 2 个成员
1) "gg"
2) "yy"
127.0.0.1:6379> SRANDMEMBER myset -2  # myset 不为空集合, count = -2, 随机返回 count 的绝对值个成员
1) "haha"
2) "hello"
127.0.0.1:6379> SCARD myset    # 查看 myset 成员数量
(integer) 6
```

#### 遍历成员

- SSCAN key cursor [MATCH pattern] [COUNT count] 迭代集合中的元素, 返回游标的位置和结果
  - cursor 游标
  - pattern 匹配的模式
  - count 控制从数据集返回的成员数量, 默认为 10

```shell
127.0.0.1:6379> FLUSHALL    # 清空所有数据库
OK
127.0.0.1:6379> KEYS *    # 查看当前数据库中的 key
(empty array)
127.0.0.1:6379> SSCAN myset 0 MATCH a COUNT 10    # 迭代空集合
1) "0"
2) (empty array)
127.0.0.1:6379> SADD myset a b c d aa ab ac ad ba bb bc bd ca cb cc cd da db dc dd    # 向 myset 添加成员
(integer) 20

127.0.0.1:6379> SSCAN myset 0 MATCH a* COUNT 20    # 迭代 myset, 匹配模式 a*
1) "0"
2) 1) "ad"
   2) "a"
   3) "ac"
   4) "aa"
   5) "ab"
127.0.0.1:6379> SSCAN myset 0 MATCH *c COUNT 8    # 迭代 myset, 匹配模式 *c, 游标位置
1) "10"
2) 1) "dc"
   2) "cc"
   3) "ac"
   4) "c"
127.0.0.1:6379> SSCAN myset 10 MATCH *c COUNT 8    # 从第 1 次返回的游标位置开始继续迭代 myset
1) "11"
2) 1) "bc"
```

#### 获取成员差异

- SDIFF key [key ...] 比较第一个集合和其他集合之间的差异并返回差异的结果, 第一个集合为空或者第一个集合的所有成员在出现在其他集合中返回(empty array)

```shell
127.0.0.1:6379> KEYS *  # 查看当前数据库中的 key
1) "myset"
127.0.0.1:6379> SMEMBERS myset  # 查看集合 myset 的成员
1) "world"
2) "hello"
127.0.0.1:6379> SDIFF myset newset newset2  # 比较集合的差异, newset 和 newset2 为空, 返回 myset 的所有成员
1) "world"
2) "hello"

127.0.0.1:6379> SADD newset hello gg  # 向集合 newset 添加成员 hello, gg
(integer) 2
127.0.0.1:6379> SDIFF myset newset newset2  # 比较集合的差异, newset2 为空, 返回其他集合不包含的成员 world
1) "world"
127.0.0.1:6379> SADD newset2 hehe  # 向集合 newset2 添加成员 hehe
(integer) 1
127.0.0.1:6379> SDIFF myset newset newset2  # 比较集合的差异, 返回其他集合不包含的成员 world
1) "world"
127.0.0.1:6379> SADD newset2 world   # 向集合 newset2 添加成员 world
(integer) 1
127.0.0.1:6379> SDIFF myset newset newset2  # 比较集合的差异, myset 的成员包含在 newset 和 newset2 集合中, 返回 empty array
(empty array)

127.0.0.1:6379> DEL myset  # 删除集合 myset
(integer) 1
127.0.0.1:6379> KEYS *  # 查看当前数据库中的 key
1) "newset2"
2) "newset"
127.0.0.1:6379> SDIFF myset newset newset2  # 比较集合的差异, myset 集合为空, newset 和 newset2 集合有成员, 返回 empty array
(empty array)
```

- SDIFFSTORE destination key [key ...] 比较第一个集合和其他集合之间的差异把差异结果存储到指定的集合中并返回指定集合的数量, 如果指定集合不存在则新建, 如果指定集合存在则覆盖指定集合, 0 表示未找到结果

```shell
127.0.0.1:6379> KEYS *  # 查看当前数据库中的 key
(empty array)
127.0.0.1:6379> SADD myset hello world  # 向 myset 中添加 hello, world
(integer) 2
127.0.0.1:6379> SADD newset world gg   # 向 newset 中添加 world, gg
(integer) 2
127.0.0.1:6379> SADD destset hehe haha  # 向 destset 中添加 hehe, haha
(integer) 2
127.0.0.1:6379> KEYS *   # 查看当前数据库中的 key
1) "destset"
2) "myset"
3) "newset"
127.0.0.1:6379> SDIFFSTORE destset myset newset   # 比较 myset 和 newset 的差异, 将差异结果覆盖存储到 destset 中
(integer) 1
127.0.0.1:6379> SMEMBERS destset   # 查看 destset 集合的成员
1) "hello"
```

#### 获取成员交集

- SINTER key [key ...] 返回所有给定集合之间的交集, key 不存在被当作空集合, 当给定集合中有一个空集合时返回结果也为空集合(empty array)

```shell
127.0.0.1:6379> KEYS *  # 查看当前数据库中的 key
(empty array)
127.0.0.1:6379> SADD myset hello world   # 向 myset 中添加成员 hello, world
(integer) 2
127.0.0.1:6379> SINTER myset newset   # 获取 myset 和 newset 的交集, newset 集合为空集, 返回空集
(empty array)

127.0.0.1:6379> SADD newset hehe hello   # 向 newset 中添加成员 hehe, hello
(integer) 2
127.0.0.1:6379> SINTER myset newset    # 获取 myset 和 newset 的交集, 返回相同的成员 hello
1) "hello"

127.0.0.1:6379> SADD newset2 world gg   # 向 newset2 中添加成员 world, gg
(integer) 2
127.0.0.1:6379> SINTER myset newset newset2    # 获取 myset, newset, newset2 的交集, 没有找到相同的成员返回空集
(empty array)
```

- SINTERSTORE destination key [key ...] 将所有给定集合之间的交集存储到指定的集合中并返回指定集合的数量, 如果指定集合不存在则新建, 如果指定集合已存在则覆盖指定集合, 0 表示未找到结果

```shell
127.0.0.1:6379> KEYS *    # 查看当前数据库中的 key
(empty array)
127.0.0.1:6379> SADD myset hello world    # 向 myset 中添加成员 hello, world
(integer) 2
127.0.0.1:6379> SADD newset hello hehe    # 向 newset 中添加成员 hello, hehe
(integer) 2
127.0.0.1:6379> SADD destset gg yy       # 向 destset 中添加成员 gg, yy
(integer) 2
127.0.0.1:6379> SINTERSTORE destset myset newset    # 获取 myset, newset 的交集覆盖存储到 destset 中
(integer) 1
127.0.0.1:6379> SMEMBERS destset
1) "hello"
```

#### 获取成员并集

- SUNION key [key ...] 返回所有给定集合的并集并移除相同的成员只保留一个, 不存在的 key 被当作空集合, 集合都为空返回 (empty array)

```shell
127.0.0.1:6379> FLUSHALL    # 清空所有数据库
OK
127.0.0.1:6379> KEYS *    # 查看当前数据库中的 key
(empty array)
127.0.0.1:6379> SUNION myset newset    # 返回空集合的并集
(empty array)
127.0.0.1:6379> SADD myset hello world    # 向 myset 添加 hello, world
(integer) 2
127.0.0.1:6379> SUNION myset newset    # 返回 myset, newset 的并集
1) "world"
2) "hello"
127.0.0.1:6379> SADD newset hello hehe    # 向 newset 添加成员 hello, hehe
(integer) 2
127.0.0.1:6379> SUNION myset newset    # 返回 myset, newset 的并集, 只保留相同一个成员 hello
1) "world"
2) "hehe"
3) "hello"
```

- SUNIONSTORE destination key [key ...] 将所有给定集合之间的并集存储到指定的集合中并返回指定集合的数量, 如果指定集合不存在则新建, 如果指定集合已存在则覆盖指定集合, 0 表示未找到结果

```shell
127.0.0.1:6379> FLUSHALL    # 清空所有数据库
OK
127.0.0.1:6379> KEYS *    # 查看当前数据库中的 key
(empty array)
127.0.0.1:6379> SUNIONSTORE destset myset newset    # myset, newset 都为空集合, 返回 0
(integer) 0
127.0.0.1:6379> SADD myset hello world    # 向 myset 添加成员 hello, world
(integer) 2
127.0.0.1:6379> SUNIONSTORE destset myset newset    # 获取 myset, newset 的并集新建存储到 destset 中
(integer) 2
127.0.0.1:6379> SMEMBERS destset    # 查看 destset 成员
1) "world"
2) "hello"
127.0.0.1:6379> SADD newset hello hehe haha    # 向 newset 添加成员 hello, haha
(integer) 3
127.0.0.1:6379> SUNIONSTORE destset myset newset    # 获取 myset, newset 的并集覆盖存储到 destset 中
(integer) 4
127.0.0.1:6379> SMEMBERS destset    # 查看 destset 成员
1) "world"
2) "hehe"
3) "haha"
4) "hello"
```

- SINTERCARD numkeys key [key ...] [LIMIT limit] 返回给定多个集合的交集的数量, 0 表示未找到结果, 7.0.0 支持
  - numkeys 指定集合的数量, 值和 key 的数量不一致时返回语法错误 syntax error
  - limit 指定返回结果的偏移量, 默认为 0, limit < 0 时, 报错不能为负数

```shell
127.0.0.1:6379> KEYS *    # 查看当前数据库中的 key
(empty array)
127.0.0.1:6379> SINTERCARD 1 myset newset    # 语法错误
(error) ERR syntax error
127.0.0.1:6379> SINTERCARD 2 myset newset    # 返回交集的数量
(integer) 0

127.0.0.1:6379> SADD myset hello world hehe haha gg yy    # 向 myset 添加成员
(integer) 6
127.0.0.1:6379> SADD newset hello hehe haha yy    # 向 newset 添加成员
(integer) 4
127.0.0.1:6379> SINTERCARD 2 myset newset     # 返回交集的数量(hello, hehe, haha, yy)
(integer) 4
127.0.0.1:6379> SINTERCARD 2 myset newset LIMIT 0    # 作用同上一条命令
(integer) 4
127.0.0.1:6379> SINTERCARD 2 myset newset LIMIT 1    # 指定返回结果的偏移量
(integer) 1
127.0.0.1:6379> SINTERCARD 2 myset newset LIMIT 3    # 指定返回结果的偏移量
(integer) 3
127.0.0.1:6379> SINTERCARD 2 myset newset LIMIT -1    # 偏移量不能为负数
(error) ERR LIMIT can't be negative
```
