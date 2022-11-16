---
title: Redis-Set-Zset
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
127.0.0.1:6379> KEYS *  # 查看当前数据库中的 key
(empty array)
# 空集合判断是否包含成员
127.0.0.1:6379> SMISMEMBER myset hello hehe
1) (integer) 0
2) (integer) 0
# 向 myset 添加成员
127.0.0.1:6379> SADD myset hello world hehe haha gg
(integer)
# 判断 myset 是否包含成员
127.0.0.1:6379> SMISMEMBER myset hello yy hehe
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
# 向集合添加成员
127.0.0.1:6379> SADD set hello world gg
(integer) 3
# 查看集合的所有成员
127.0.0.1:6379> SMEMBERS set
1) "gg"
2) "world"
3) "hello"
# 移除不存在的集合的成员
127.0.0.1:6379> SREM myset a b
(integer) 0
# 移除集合不存在的成员
127.0.0.1:6379> SREM set a b
(integer) 0
# 移除集合的成员
127.0.0.1:6379> SREM set hello a
(integer) 1
# 查看集合的所有成员
127.0.0.1:6379> SMEMBERS set
1) "gg"
2) "world"
# 移除集合的成员
127.0.0.1:6379> SREM set world gg
(integer) 2
# 查看集合为空的所有成员
127.0.0.1:6379> SMEMBERS key
(empty array)
```

##### 移除成员添加到其他集合

- SMOVE source destination member 将源集合中的指定成员移除并添加到指定集合中, 1 表示源集合指定成员移除成功(目标集合中可能包含该成员也可能不包含), 0 表示源集合为空或者源集合不包含指定成员

```shell
127.0.0.1:6379> FLUSHALL    # 清空所有数据库
OK
# myset 为空集, 返回 0
127.0.0.1:6379> SMOVE myset destset hello
(integer) 0

# 向 myset 添加成员 hello, world
127.0.0.1:6379> SADD myset hello world
(integer) 2
# 移除 myset 包含的成员 hello 并添加到 destset 中, destset 为空, 返回 1
127.0.0.1:6379> SMOVE myset destset hello
(integer) 1
# 查看 myset 成员
127.0.0.1:6379> SMEMBERS myset
1) "world"
# 查看 destset 成员
127.0.0.1:6379> SMEMBERS destset
1) "hello"
# 移除 myset 不包含的成员 hehe, 返回 0
127.0.0.1:6379> SMOVE myset destset hehe
(integer) 0
# 查看 myset 成员
127.0.0.1:6379> SMEMBERS myset
1) "world"
# 查看 destset 成员
127.0.0.1:6379> SMEMBERS destset
1) "hello"

127.0.0.1:6379> FLUSHALL    # 清空所有数据库
OK
# 向 myset 添加成员 hello, world
127.0.0.1:6379> SADD myset hello world
(integer) 2
# 向 destset 添加成员 hello, gg
127.0.0.1:6379> SADD destset hello gg
(integer) 2
# 移除 myset 包含的成员 hello 并添加到 destset 中, destset 包含 hello, 返回 1
127.0.0.1:6379> SMOVE myset destset hello
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

# myset 为空集合, 返回 nil
127.0.0.1:6379> SPOP myset
(nil)
# myset 为空集合, count >= 0 返回 empty array
127.0.0.1:6379> SPOP myset 2
(empty array)
# 向 myset 添加成员 hello, world, hehe, haha, gg, yy
127.0.0.1:6379> SADD myset hello world hehe haha gg yy
(integer) 6
# count = 0, 任意集合都返回 empty array
127.0.0.1:6379> SPOP myset 0
(empty array)
# count = 0, 任意集合都返回 empty array
127.0.0.1:6379> SPOP newset 0
(empty array)
# count < 0, 返回错误
127.0.0.1:6379> SPOP myset -1
(error) ERR value is out of range, must be positive
# 移除 myset 随机的 1 个成员
127.0.0.1:6379> SPOP myset
"world"
# 移除 myset 随机的 3 个成员
127.0.0.1:6379> SPOP myset 3
1) "haha"
2) "yy"
3) "gg"
127.0.0.1:6379> SMEMBERS myset  # 查看 myset 成员
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
127.0.0.1:6379> SMEMBERS myset  # 查看 myset 成员
1) "haha"
2) "yy"
3) "world"
4) "hehe"
5) "hello"
6) "gg"

# newset 为空集合返回 nil
127.0.0.1:6379> SRANDMEMBER newset
(nil)
# newset 为空集合, count = 0, 返回 empty array
127.0.0.1:6379> SRANDMEMBER newset 0
(empty array)
# newset 为空集合, count = 1, 返回 empty array
127.0.0.1:6379> SRANDMEMBER newset 1
(empty array)
# newset 为空集合, count = -1, 返回 empty array
127.0.0.1:6379> SRANDMEMBER newset -1
(empty array)

# 随机返回 myset 集合的 1 个成员
127.0.0.1:6379> SRANDMEMBER myset
"hello"
# myset 不为空集合, count = 0, 返回 empty array
127.0.0.1:6379> SRANDMEMBER myset 0
(empty array)
# myset 不为空集合, count = 2, 随机返回 2 个成员
127.0.0.1:6379> SRANDMEMBER myset 2
1) "gg"
2) "yy"
# myset 不为空集合, count = -2, 随机返回 count 的绝对值个成员
127.0.0.1:6379> SRANDMEMBER myset -2
1) "haha"
2) "hello"
127.0.0.1:6379> SCARD myset # 查看 myset 成员数量
(integer) 6
```

#### 遍历无序集合

- SSCAN key cursor [MATCH pattern] [COUNT count] 迭代集合中的成员, 返回下一次游标开始的位置和结果, 游标 0 表示迭代已结束
  - cursor 游标
  - pattern 匹配的模式
  - count 控制从数据集返回的成员数量, 默认为 10

```shell
127.0.0.1:6379> FLUSHALL  # 清空所有数据库
OK
# 迭代空集合
127.0.0.1:6379> SSCAN myset 0 MATCH a COUNT 10
1) "0"
2) (empty array)
# 向 myset 添加成员
127.0.0.1:6379> SADD myset a b c d aa ab ac ad ba bb bc bd ca cb cc cd da db dc dd
(integer) 20

# 迭代 myset, 匹配模式 a*
127.0.0.1:6379> SSCAN myset 0 MATCH a* COUNT 20
1) "0"
2) 1) "ad"
   2) "a"
   3) "ac"
   4) "aa"
   5) "ab"
# 迭代 myset, 匹配模式 *c, 游标位置
127.0.0.1:6379> SSCAN myset 0 MATCH *c COUNT 8
1) "10"
2) 1) "dc"
   2) "cc"
   3) "ac"
   4) "c"
# 从第 1 次返回的游标位置开始继续迭代 myset
127.0.0.1:6379> SSCAN myset 10 MATCH *c COUNT 8
1) "11"
2) 1) "bc"
```

#### 获取无序集合差异

- SDIFF key [key ...] 比较第一个集合和其他集合之间的差异并返回差异的结果, 第一个集合为空或者第一个集合的所有成员在出现在其他集合中返回(empty array)

```shell
127.0.0.1:6379> KEYS *  # 查看当前数据库中的 key
1) "myset"
127.0.0.1:6379> SMEMBERS myset  # 查看集合 myset 的成员
1) "world"
2) "hello"
# 比较集合的差异, newset 和 newset2 为空, 返回 myset 的所有成员
127.0.0.1:6379> SDIFF myset newset newset2
1) "world"
2) "hello"

# 向集合 newset 添加成员 hello, gg
127.0.0.1:6379> SADD newset hello gg
(integer) 2
# 比较集合的差异, newset2 为空, 返回其他集合不包含的成员 world
127.0.0.1:6379> SDIFF myset newset newset2
1) "world"
# 向集合 newset2 添加成员 hehe
127.0.0.1:6379> SADD newset2 hehe
(integer) 1
# 比较集合的差异, 返回其他集合不包含的成员 world
127.0.0.1:6379> SDIFF myset newset newset2
1) "world"
# 向集合 newset2 添加成员 world
127.0.0.1:6379> SADD newset2 world
(integer) 1
# 比较集合的差异, myset 的成员包含在 newset 和 newset2 集合中, 返回 empty array
127.0.0.1:6379> SDIFF myset newset newset2
(empty array)

127.0.0.1:6379> DEL myset  # 删除集合 myset
(integer) 1
127.0.0.1:6379> KEYS *  # 查看当前数据库中的 key
1) "newset2"
2) "newset"
# 比较集合的差异, myset 集合为空, newset 和 newset2 集合有成员, 返回 empty array
127.0.0.1:6379> SDIFF myset newset newset2
(empty array)
```

- SDIFFSTORE destination key [key ...] 比较第一个集合和其他集合之间的差异把差异结果存储到指定集合并返回指定集合的数量, 如果指定集合不存在则新建, 如果指定集合存在则覆盖指定集合, 0 表示未找到结果

```shell
127.0.0.1:6379> KEYS *  # 查看当前数据库中的 key
(empty array)
# 向 myset 中添加 hello, world
127.0.0.1:6379> SADD myset hello world
(integer) 2
# 向 newset 中添加 world, gg
127.0.0.1:6379> SADD newset world gg
(integer) 2
# 向 destset 中添加 hehe, haha
127.0.0.1:6379> SADD destset hehe haha
(integer) 2
127.0.0.1:6379> KEYS *   # 查看当前数据库中的 key
1) "destset"
2) "myset"
3) "newset"
# 比较 myset 和 newset 的差异, 将差异结果覆盖存储到 destset 中
127.0.0.1:6379> SDIFFSTORE destset myset newset
(integer) 1
# 查看 destset 集合的成员
127.0.0.1:6379> SMEMBERS destset
1) "hello"
```

#### 获取无序集合交集

- SINTER key [key ...] 返回所有给定集合之间的交集, key 不存在被当作空集合, 当给定集合中有一个空集合时返回结果也为空集合(empty array)

```shell
127.0.0.1:6379> KEYS *  # 查看当前数据库中的 key
(empty array)
# 向 myset 中添加成员 hello, world
127.0.0.1:6379> SADD myset hello world
(integer) 2
# 获取 myset 和 newset 的交集, newset 集合为空集, 返回空集
127.0.0.1:6379> SINTER myset newset
(empty array)

# 向 newset 中添加成员 hehe, hello
127.0.0.1:6379> SADD newset hehe hello
(integer) 2
# 获取 myset 和 newset 的交集, 返回相同的成员 hello
127.0.0.1:6379> SINTER myset newset
1) "hello"

# 向 newset2 中添加成员 world, gg
127.0.0.1:6379> SADD newset2 world gg
(integer) 2
# 获取 myset, newset, newset2 的交集, 没有找到相同的成员返回空集
127.0.0.1:6379> SINTER myset newset newset2
(empty array)
```

- SINTERSTORE destination key [key ...] 将所有给定集合之间的交集存储到指定的集合中并返回指定集合的数量, 如果指定集合不存在则新建, 如果指定集合已存在则覆盖指定集合, 0 表示未找到结果

```shell
127.0.0.1:6379> KEYS *    # 查看当前数据库中的 key
(empty array)
# 向 myset 中添加成员 hello, world
127.0.0.1:6379> SADD myset hello world
(integer) 2
# 向 newset 中添加成员 hello, hehe
127.0.0.1:6379> SADD newset hello hehe
(integer) 2
# 向 destset 中添加成员 gg, yy
127.0.0.1:6379> SADD destset gg yy
(integer) 2
# 获取 myset, newset 的交集覆盖存储到 destset 中
127.0.0.1:6379> SINTERSTORE destset myset newset
(integer) 1
127.0.0.1:6379> SMEMBERS destset
1) "hello"
```

- SINTERCARD numkeys key [key ...] [LIMIT limit] 返回给定多个集合的交集的数量, 0 表示未找到结果, 7.0.0 支持
  - numkeys 指定集合的数量, 值和 key 的数量不一致时返回语法错误 syntax error
  - limit 指定返回结果的偏移量, 默认为 0, limit < 0 时, 报错不能为负数

```shell
127.0.0.1:6379> KEYS *    # 查看当前数据库中的 key
(empty array)
# 语法错误
127.0.0.1:6379> SINTERCARD 1 myset newset
(error) ERR syntax error
# 返回交集的数量
127.0.0.1:6379> SINTERCARD 2 myset newset
(integer) 0

# 向 myset 添加成员
127.0.0.1:6379> SADD myset hello world hehe haha gg yy
(integer) 6
# 向 newset 添加成员
127.0.0.1:6379> SADD newset hello hehe haha yy
(integer) 4
# 返回交集的数量(hello, hehe, haha, yy)
127.0.0.1:6379> SINTERCARD 2 myset newset
(integer) 4
 # 作用同上一条命令
127.0.0.1:6379> SINTERCARD 2 myset newset LIMIT 0
(integer) 4
# 指定返回结果的偏移量
127.0.0.1:6379> SINTERCARD 2 myset newset LIMIT 1
(integer) 1
# 指定返回结果的偏移量
127.0.0.1:6379> SINTERCARD 2 myset newset LIMIT 3
(integer) 3
# 偏移量不能为负数
127.0.0.1:6379> SINTERCARD 2 myset newset LIMIT -1
(error) ERR LIMIT can't be negative
```

#### 获取无序集合并集

- SUNION key [key ...] 返回所有给定集合的并集并移除相同的成员只保留一个, 不存在的 key 被当作空集合, 集合都为空返回 (empty array)

```shell
127.0.0.1:6379> FLUSHALL    # 清空所有数据库
OK
# 返回空集合的并集
127.0.0.1:6379> SUNION myset newset
(empty array)
# 向 myset 添加 hello, world
127.0.0.1:6379> SADD myset hello world
(integer) 2
# 返回 myset, newset 的并集
127.0.0.1:6379> SUNION myset newset
1) "world"
2) "hello"
# 向 newset 添加成员 hello, hehe
127.0.0.1:6379> SADD newset hello hehe
(integer) 2
# 返回 myset, newset 的并集, 只保留相同一个成员 hello
127.0.0.1:6379> SUNION myset newset
1) "world"
2) "hehe"
3) "hello"
```

- SUNIONSTORE destination key [key ...] 将所有给定集合之间的并集存储到指定的集合中并返回指定集合的数量, 如果指定集合不存在则新建, 如果指定集合已存在则覆盖指定集合, 0 表示未找到结果

```shell
127.0.0.1:6379> FLUSHALL  # 清空所有数据库
OK
# myset, newset 都为空集合, 返回 0
127.0.0.1:6379> SUNIONSTORE destset myset newset
(integer) 0
# 向 myset 添加成员 hello, world
127.0.0.1:6379> SADD myset hello world
(integer) 2
# 获取 myset, newset 的并集新建存储到 destset 中
127.0.0.1:6379> SUNIONSTORE destset myset newset
(integer) 2
# 查看 destset 成员
127.0.0.1:6379> SMEMBERS destset
1) "world"
2) "hello"
# 向 newset 添加成员 hello, haha
127.0.0.1:6379> SADD newset hello hehe haha
(integer) 3
# 获取 myset, newset 的并集覆盖存储到 destset 中
127.0.0.1:6379> SUNIONSTORE destset myset newset
(integer) 4
127.0.0.1:6379> SMEMBERS destset  # 查看 destset 成员
1) "world"
2) "hehe"
3) "haha"
4) "hello"
```

### Zset 命令

Zset 和 set 一样也是 string 类型元素的集合, 且不允许重复的成员, 不同的是每个元素都会关联一个 double 类型的分值, Redis 正是通过分值来为集合中的成员进行从小到大的排序, 有序集合的成员是唯一的,但分值(score)却可以重复, 集合是通过哈希表实现的, 最大的成员数为 2^32-1(40 多亿)个成员.

> 大部分 Set 的 API 将首字母 s 换成 z 就可以使用, 这里只列出部分不一致的 API

#### 特殊标识符

- \- 负
- \+ 正
- inf 无穷大
- ({val} 不包含 val
- [{val} 包含 val

#### 添加成员

- ZADD key [NX|XX] [GT|LT] [CH] [INCR] score member [score member ...] 添加更新成员, 通常只返回添加的新成员的数量

  - NX 仅添加新成员, 不再更新已存在的成员
  - XX 仅更新已经存在的成员, 不再添加新成员
  - GT 仅当新分值大于当前分值才更新已存在的成员, 此标志不阻止添加新成员
  - LT 仅当新分值小于当前分值才更新已存在的成员, 此标志不阻止添加新成员
  - CH 将 `ZADD` 返回值统计新成员的添加数量修改为更改的成员总数, 包含更新已存在的数量和新添加的数量
  - INCR 此选项作用类似 `ZINCRBY`, 只能指定一个成员分值对加上指定的增量(可以为负数), 如果成员不存在从 0 开始计算, 并返回当前成员的最终分值, 多个成员分值对会报错

```shell
127.0.0.1:6379> ZADD myz 1 zhangsan 2 lisi  # 添加成员
(integer) 2
127.0.0.1:6379> ZCARD myz
(integer) 2
127.0.0.1:6379> ZRANGE myz 0 -1 WITHSCORES
1) "lisi"
2) "2"
3) "zhangsan"
4) "1"

# XX  仅更新已存在的成员，不再添加新成员
# 仅更新 zhangsan, 忽略添加新成员 1 wangwu 3.5 zhaoliu
127.0.0.1:6379> ZADD myz XX 1.5 zhangsan 1 wangwu 3.5 zhaoliu
(integer) 0
127.0.0.1:6379> ZCARD myz
(integer) 2
127.0.0.1:6379> ZRANGE myz 0 -1 WITHSCORES
1) "lisi"
2) "2"
3) "zhangsan"
4) "1.5"

# NX  仅添加新成员, 不再更新已存在的成员
# 仅添加新成员 1 wangwu 3.5 zhaoliu, 忽略更新 lisi
127.0.0.1:6379> ZADD myz NX 2.5 lisi 1 wangwu 3.5 zhaoliu
(integer) 2
127.0.0.1:6379> ZCARD myz
(integer) 4
127.0.0.1:6379> ZRANGE myz 0 -1 WITHSCORES
1) "zhaoliu"
2) "3.5"
3) "lisi"
4) "2"
5) "zhangsan"
6) "1.5"
7) "wangwu"
8) "1"

# LT  仅当新分值小于当前分值才更新已存在的成员, 不阻止添加新成员
# 0.5 < 2 满足条件, 有新成员 4 sunqi 可添加
127.0.0.1:6379> ZADD myz LT 0.5 lisi 4 sunqi
(integer) 1
127.0.0.1:6379> ZCARD myz
(integer) 5
127.0.0.1:6379> ZRANGE myz 0 -1 WITHSCORES
 1) "sunqi"
 2) "4"
 3) "zhaoliu"
 4) "3.5"
 5) "zhangsan"
 6) "1.5"
 7) "wangwu"
 8) "1"
 9) "lisi"
10) "0.5"
# 1 > 0.5, 8 > 4 不满足条件, 没有新成员可添加
127.0.0.1:6379> ZADD myz LT 1 lisi 8 sunqi
(integer) 0
127.0.0.1:6379> ZCARD myz
(integer) 5
127.0.0.1:6379> ZRANGE myz 0 -1 WITHSCORES
 1) "sunqi"
 2) "4"
 3) "zhaoliu"
 4) "3.5"
 5) "zhangsan"
 6) "1.5"
 7) "wangwu"
 8) "1"
 9) "lisi"
10) "0.5"

# GT 仅当新分值大于当前分值才更新已存在的成员, 不阻止添加新成员
# 1 < 1.5 不满足条件, 有新成员 2 qianba 可添加
127.0.0.1:6379> ZADD myz GT 1 zhangsan 2 qianba
(integer) 1
127.0.0.1:6379> ZCARD myz
(integer) 6
127.0.0.1:6379> ZRANGE myz 0 -1 WITHSCORES
 1) "sunqi"
 2) "4"
 3) "zhaoliu"
 4) "3.5"
 5) "qianba"
 6) "2"
 7) "zhangsan"
 8) "1.5"
 9) "wangwu"
10) "1"
11) "lisi"
12) "0.5"
# 3 > 1.5 满足条件, 2 = 2 不满足条件, 没有新成员可添加
127.0.0.1:6379> ZADD myz GT 3 zhangsan 2 qianba
(integer) 0
127.0.0.1:6379> ZCARD myz
(integer) 6
127.0.0.1:6379> ZRANGE myz 0 -1 WITHSCORES
 1) "sunqi"
 2) "4"
 3) "zhaoliu"
 4) "3.5"
 5) "zhangsan"
 6) "3"
 7) "qianba"
 8) "2"
 9) "wangwu"
10) "1"
11) "lisi"
12) "0.5"

# CH 统计集合所有受影响的成员的数量, 包含更新已存在的数量和新添加的数量
# zhangsan 和 qianba 已存在, hello 为新添加成员,
127.0.0.1:6379> ZADD myz CH 1 zhangsan 5 qianba 1 hello
(integer) 3

# INCR 同时只能指定一个成员分值对加上指定的增量(可以为负数), 如果成员不存在从 0 开始计算
# 多个会报错 (error) ERR INCR option supports a single increment-element pair
127.0.0.1:6379> ZADD myz INCR 10 zhangsan
"10"
127.0.0.1:6379> ZRANGE myz 0 -1 WITHSCORES
1) "zhangsan"
2) "10"
```

#### 指定成员的分值增量

- ZINCRBY key increment member 对指定成员的分值加上增量并返回修改后的分值, 如果指定成员不存在则添加新成员, 等同于 `ZADD key increment member`

```shell
# 修改 zhangsan 的分值 + 2
127.0.0.1:6379> ZINCRBY myz 2 zhangsan
"3.5"
# zhangsan1 不存在, 新添加
127.0.0.1:6379> ZINCRBY myz 2 zhangsan1
"2"
# 修改 zhaoliu 的分值减去 - 2
127.0.0.1:6379> ZINCRBY myz -2 zhaoliu
"1.5"
127.0.0.1:6379> ZRANGE myz 0 -1 WITHSCORES
 1) "zhangsan"
 2) "3.5"
 3) "lisi"
 4) "2.5"
 5) "zhangsan1"
 6) "2"
 7) "zhaoliu"
 8) "1.5"
 9) "wangwu"
10) "1"
```

#### 移除成员

##### 分值最大或最小

- ZPOPMAX key [count] 移除指定集合的指定数量的最高分值成员并返回移除的成员和分值, count 默认为 1, 如果集合为空或者不存在返回 (empty array), 5.0.0 支持

  - count < 0 时, 报错 ERR value is out of range, must be positive
  - count = 0 时, 不做任何操作, 返回 empty array
  - count > 1 时, 指定移除数量

- BZPOPMAX key [key ...] timeout 阻塞版 `ZPOPMAX`, 从多个集合中第 1 个非空集合中移除并返回 1 个最高分值成员, 如果集合为空会阻塞集合直到等待超时或发现可移除成员为止, 如果集合为空或者超时返回 &lt;nil&gt;, 否则, 返回 1 个含有 3 个元素的列表, 第 1 个为被移除成员所属的集合, 第 2 个为被移除的成员, 第 3 个为移除成员的分值, 5.0.0 支持

- ZPOPMIN key [count] 移除指定集合的指定数量的最低分值成员并返回移除的成员和分值, count 默认为 1, 如果集合为空或者不存在返回 (empty array), 5.0.0 支持

  - 参数同 `ZPOPMAX`

- BZPOPMIN key [key ...] timeout 阻塞版 `ZPOPMIN`, 从多个集合中第 1 个非空集合中移除并返回 1 个最低分值成员, 如果集合为空会阻塞集合直到等待超时或发现可移除成员为止, 如果集合为空或者超时返回 &lt;nil&gt;, 否则, 返回 1 个含有 3 个元素的列表, 第 1 个为被移除成员所属的集合, 第 2 个为被移除的成员, 第 3 个为移除成员的分值, 5.0.0 支持

```shell
127.0.0.1:6379> ZPOPMAX myz -1
(error) ERR value is out of range, must be positive
127.0.0.1:6379> ZADD myz 1 a 2 b 3 c 3 d 2 e 1 f
(integer) 6
127.0.0.1:6379> ZPOPMAX myzz
(empty array)
127.0.0.1:6379> ZPOPMAX myz 0
(empty array)
127.0.0.1:6379> ZPOPMAX myz
1) "d"
2) "3"
127.0.0.1:6379> ZPOPMAX myz 3 # 移除 3 个最高分值成员
1) "d"
2) "3"
3) "c"
4) "3"
5) "e"
6) "2"

127.0.0.1:6379> ZPOPMIN myz 0
(empty array)
127.0.0.1:6379> ZPOPMIN myz
1) "a"
2) "1"
127.0.0.1:6379> ZPOPMIN myz 3
1) "a"
2) "1"
3) "f"
4) "1"
5) "b"
6) "2"

# 阻塞移除多个集合中第 1 个非空集合的最高值成员
127.0.0.1:6379> BZPOPMAX myz1 myz 0
1) "myz"
2) "d"
3) "4"

# 阻塞移除多个集合中第 1 个非空集合的最低值成员
127.0.0.1:6379> BZPOPMIN myz1 myz 0
1) "myz"
2) "a"
3) "1"
```

##### 指定区间

- ZREMRANGEBYLEX key min max 移除指定字典区间的所有成员并返回成功移除的数量, min 和 max 需要使用 `(` 或 `[` 前导符, 0 表示集合为空或者未找到结果

```shell
127.0.0.1:6379> ZREMRANGEBYLEX myz a f
(error) ERR min or max not valid string range item
127.0.0.1:6379> ZADD myz 0 a 0 b 0 c 0 d 0 e 0 alpha 0 zip
(integer) 7
# 匹配范围含尾不含头, 删除 b c d e zip
127.0.0.1:6379> ZREMRANGEBYLEX myz (alpha [zip
(integer) 5
127.0.0.1:6379> ZRANGE myz 0 -1
1) "a"
2) "alpha"

# 匹配范围含头不含尾, 删除 alpha b c d e
127.0.0.1:6379> ZREMRANGEBYLEX myz [alpha (zip
(integer) 5
127.0.0.1:6379> ZRANGE myz 0 -1
1) "a"
2) "zip"
```

- ZREMRANGEBYSCORE key min max 移除指定分值区间的所有成员并返回成功移除的数量, 0 表示集合为空或者未找到结果

```shell
127.0.0.1:6379> ZREMRANGEBYSCORE myz1 100 500
(integer) 0
127.0.0.1:6379> ZADD myz 1 a 2 b 3 c 4 d 5 e 6 f 2 beta 3 cipher 4 delete
(integer) 9
# 匹配范围含头含尾
127.0.0.1:6379> ZREMRANGEBYSCORE myz 2 5
(integer) 7
127.0.0.1:6379> ZRANGE myz 0 -1
1) "a"
2) "f"

# 匹配范围不含头不含尾
127.0.0.1:6379> ZREMRANGEBYSCORE myz (2 (5
(integer) 4
127.0.0.1:6379> zrange myz 0 -1
1) "a"
2) "b"
3) "beta"
4) "e"
5) "f"
```

- ZREMRANGEBYRANK key start stop 移除指定排名区间的所有成员并返回成功移除的数量, 0 表示集合为空或者未找到结果

```shell
127.0.0.1:6379> ZREMRANGEBYRANK myz1 1 4
(integer) 0
127.0.0.1:6379> ZADD myz 1 a 2 b 3 c 4 d 5 e 6 f 2 beta 3 cipher 4 delete
(integer) 9
# 移除范围下标从 1 到 4, b beta c cipher
127.0.0.1:6379> ZREMRANGEBYRANK myz 1 4
(integer) 4
127.0.0.1:6379> ZRANGE myz 0 -1
1) "a"
2) "d"
3) "delete"
4) "e"
5) "f"

# 移除范围下标从 1 到 5, b beta c cipher d
127.0.0.1:6379> ZREMRANGEBYRANK myz 1 5
(integer) 5
127.0.0.1:6379> ZRANGE myz 0 -1
1) "a"
2) "delete"
3) "e"
4) "f"
```

##### 批量移除相同成员

- ZREM key member [member ...] 批量移除指定的成员并返回成功移除的数量, 如果指定成员不存在则被忽略, 0 表示集合为空或者成员不存在

```shell
127.0.0.1:6379> ZADD myz 1 hello 2 world 3 gg 2 yy 1 hehe 4 haha
(integer) 6
127.0.0.1:6379> ZREM myz YY GG
(integer) 0
127.0.0.1:6379> ZREM myz gg GG
(integer) 1
```

##### 批量移除相邻成员

- ZMPOP numkeys key [key ...] MIN|MAX [COUNT count] 从多个集合中第 1 个非空集合中移除指定数量的最高或最低分值的成员并返回移除的成员和分值及成员所属的集合名称, count 默认为 1, 集合为空或者不存在返回 &lt;nil&gt;, 7.0.0 支持
  - COUNT count 移除成员的数量, 默认为 1
- BZMPOP timeout numkeys key [key ...] MIN|MAX [COUNT count] 阻塞版 `ZMPOP`, 如果集合为空会阻塞集合直到等待超时或发现可移除成员为止, 如果集合为空或者超时返回 &lt;nil&gt;, 7.0.0 支持

```shell
# 如果集合都为空, 阻塞指定时间后返回 nil
127.0.0.1:6379> BZMPOP 5 2 myzz myz MAX COUNt 2
(nil)
(5.09s)

127.0.0.1:6379> ZADD myz 1 a 2 b 3 c 3 d 3 e 2 f 2 g 1 h 4 i
(integer) 9
# 移除多个集合中第 1 个非空集合的 1 个最大值成员
127.0.0.1:6379> ZMPOP 2 myzz myz MAX
1) "myz"
2) 1) 1) "i"
      2) "4"
# 移除多个集合中第 1 个非空集合的 2 个最大值成员
127.0.0.1:6379> ZMPOP 2 myzz myz MAX COUNT 2
1) "myz"
2) 1) 1) "i"
      2) "4"
   2) 1) "e"
      2) "3"
# 移除多个集合中第 1 个非空集合的 3 个最小值成员
127.0.0.1:6379> ZMPOP 2 myzz myz MIN COUNT 3
1) "myz"
2) 1) 1) "a"
      2) "1"
   2) 1) "h"
      2) "1"
   3) 1) "b"
      2) "2"
# 阻塞移除多个集合中第 1 个非空集合的 4 个最大值成员
127.0.0.1:6379> BZMPOP 5 2 myzz myz MAX COUNt 4
1) "myz"
2) 1) 1) "i"
      2) "4"
   2) 1) "e"
      2) "3"
   3) 1) "d"
      2) "3"
   4) 1) "c"
      2) "3"
```

#### 获取成员数量

- ZCARD key 返回指定集合的成员数量, 集合为空或者不存在返回 0

##### 指定分值区间

- ZCOUNT key min max 统计指定分值区间的成员数量, 0 表示未找到结果

```shell
127.0.0.1:6379> ZRANGE myz 0 -1 WITHSCORES
1) "zhaoliu"
2) "3.5"
3) "lisi"
4) "2.5"
5) "zhangsan"
6) "1.5"
7) "wangwu"
8) "1"
# 获取分值 负无穷大 到 正无穷大 之间的数量
127.0.0.1:6379> ZCOUNT myz -inf +inf
(integer) 4
# 获取分值在 0 到 2 之间的数量
127.0.0.1:6379> ZCOUNT myz 0 2
(integer) 2
# 获取分值在 1.5 到 3 之间的数量
127.0.0.1:6379> ZCOUNT myz 1.5 3
(integer) 2
```

##### 指定字典区间

- ZLEXCOUNT key min max 计算指定字典区间内成员数量, 0 表示未找到结果

```shell
127.0.0.1:6379> ZADD myz 1 a 2 b 3 c 4 d 5 e 6 f
(integer) 6
127.0.0.1:6379> ZLEXCOUNT myz - +
(integer) 6
127.0.0.1:6379> ZLEXCOUNT myz (b [f # 不包含 b
(integer) 4
127.0.0.1:6379> ZLEXCOUNT myz [b [f # 包含 b
(integer) 5
127.0.0.1:6379> ZLEXCOUNT myz [d [f
(integer) 3
127.0.0.1:6379> ZLEXCOUNT myz [a [c
(integer) 3
127.0.0.1:6379> ZLEXCOUNT myz [a [b
(integer) 0
```

#### 随机获取指定数量成员

- ZRANDMEMBER key [count [WITHSCORES]] 返回指定集合随机的多个成员, 不改变原集合, 如果集合为空或者不存在返回 &lt;nil&gt;, 否则返回 (empty array), 6.2.0 支持
  - count 指定随机返回的数量, 默认为 1
    - count >= 1 时, 空集合返回 (empty array)
    - count = 0 时, 任意集合都返回 (empty array)
    - count < 0 时, 空集合返回 (empty array), 非空集合返回 count 的绝对值数量
  - WITHSCORES 返回结果的分值

```shell
127.0.0.1:6379> ZRANDMEMBER myz1
(nil)
127.0.0.1:6379> ZRANDMEMBER myz1 -1
(empty array)
127.0.0.1:6379> ZADD myz 1 hello 2 world 3 gg 2 yy 4 hehe 5 haha
(integer) 6
127.0.0.1:6379> ZRANDMEMBER myz
"world"
127.0.0.1:6379> ZRANDMEMBER myz 0
(empty array)
127.0.0.1:6379> ZRANDMEMBER myz 3 WITHSCORES
1) "hehe"
2) "4"
3) "gg"
4) "3"
5) "yy"
6) "2"
```

#### 获取指定成员排名

```shell
127.0.0.1:6379> ZRANGE myz 0 -1 WITHSCORES
 1) "hello"
 2) "1"
 3) "world"
 4) "2"
 5) "yy"
 6) "2"
 7) "gg"
 8) "3"
 9) "hehe"
10) "4"
11) "haha"
12) "5"
```

- ZRANK key member 返回指定成员的按递增顺序的排名, 从 0 开始计算, 如果指定成员不属于指定集合则返回 &lt;nil&gt;

```shell
127.0.0.1:6379> ZRANK myz yy
(integer) 2
127.0.0.1:6379> ZRANK myz world
(integer) 1
127.0.0.1:6379> ZRANK myz HAHA
(nil)
```

- ZREVRANK key member 返回指定成员的按递减顺序的排名, 从 0 开始计算, 如果指定成员不属于指定集合则返回 &lt;nil&gt;

```shell
127.0.0.1:6379> ZREVRANK myz yy
(integer) 3
127.0.0.1:6379> ZREVRANK myz HAHA
(nil)
127.0.0.1:6379> ZREVRANK myz hehe
(integer) 1
```

#### 获取指定成员分值

```shell
127.0.0.1:6379> ZRANGE myz 0 -1 WITHSCORES
 1) "hello"
 2) "1"
 3) "world"
 4) "2"
 5) "yy"
 6) "2"
 7) "gg"
 8) "3"
 9) "hehe"
10) "4"
11) "haha"
12) "5"
```

- ZSCORE key member 获取指定成员的分值, 如果集合为空或者指定成员不属于集合返回 &lt;nil&gt;

```shell
127.0.0.1:6379> ZSCORE myz1 hehe
(nil)
127.0.0.1:6379> ZSCORE myz HAHA
(nil)
127.0.0.1:6379> ZSCORE myz gg
"3"
```

- ZMSCORE key member [member ...] 批量获取指定成员的分值, 如果集合为空或者指定成员不属于集合返回 &lt;nil&gt;, 6.2.0 支持

```shell
127.0.0.1:6379> ZMSCORE myz HAHA hehe
1) (nil)
2) "4"
127.0.0.1:6379> ZMSCORE myz1 hehe haha
1) (nil)
2) (nil)
127.0.0.1:6379> ZMSCORE myz hello gg
1) "1"
2) "3"
```

#### 遍历有序集合

- ZRANGE key start stop [BYSCORE|BYLEX] [REV] [LIMIT offset count] [WITHSCORES] 遍历有序集合的指定区间, 并返回默认递增顺序的成员

  - start 起始指针, 默认 start 表示数字, 0 表示第一个成员
  - stop 结束指针, 默认 stop 表示数字, -1 表示最后 1 个成员, -2 表示倒数第 2 个成员
  - BYSCORE 成员按分值排序, 此模式 start 和 stop 代表分值, start 指针要考量最高分值, stop 指针考量最低分值, start 必须大于等于 stop 才能返回内容
  - BYLEX 成员按字典排序, 此模式 start 和 stop 代表字典, 行为类似于 `ZRANGEBYLEX`, 返回 start 和 stop 字典闭合区间的成员, start 和 stop 需要使用 `(` 或 `[` 前导符
  - REV 反向排序结果
  - LIMIT 指定返回结果的偏移量, 需要结合 `BYSCORE` `BYLEX` 使用
  - WITHSCORES 返回结果的分值

> ZRANGEBYSCORE 6.2.0 开始废弃, 使用 `ZRANGE BYSCORE` 代替
> ZRANGEBYLEX 6.2.0 开始废弃, 使用 `ZRANGE BYLEX` 代替

> ZREVRANGE 6.2.0 开始废弃, 使用 `ZRANGE REV` 代替
> ZREVRANGEBYSCORE 6.2.0 开始废弃, 使用 `ZRANGE BYSCORE REV` 代替
> ZREVRANGEBYLEX 6.2.0 开始废弃, 使用 `ZRANGE BYLEX REV` 代替

```shell
127.0.0.1:6379> ZADD myz 10 a 9 b 13 c 7 d 11 e 6 f
(integer) 6
# 匹配范围  8 < score < 13
127.0.0.1:6379> ZRANGE myz (8 (13 BYSCORE WITHSCORES
1) "b"
2) "9"
3) "a"
4) "10"
5) "e"
6) "11"
# 反向排序结果
127.0.0.1:6379> ZRANGE myz 0 -1 REV WITHSCORES
 1) "c"
 2) "13"
 3) "e"
 4) "11"
 5) "a"
 6) "10"
 7) "b"
 8) "9"
 9) "d"
10) "7"
11) "f"
12) "6"
# 匹配范围 0 到 正无穷大, 偏移 1 最多取 3 条
127.0.0.1:6379> ZRANGE myz 0 +inf BYSCORE LIMIT 1 3 WITHSCORES
1) "d"
2) "7"
3) "b"
4) "9"
5) "a"
6) "10"
```

- ZRANGESTORE dst src min max [BYSCORE|BYLEX] [REV] [LIMIT offset count] 命令行为和 `ZRANGE` 相似, 区别是把结果集存储到指定的集合 `dst` 中

  - dst 存储结果集的指定集合
  - src 被操作的集合

- ZSCAN key cursor [MATCH pattern] [COUNT count] 迭代集合中的成员(包括成员和分值), 返回下一次游标开始的位置和结果, 游标 0 表示迭代已结束
  - cursor 游标
  - pattern 匹配的模式
  - count 控制从数据集返回的成员数量, 默认为 10

```shell
127.0.0.1:6379> ZSCAN myz1 0
1) "0"
2) (empty array)
127.0.0.1:6379> ZADD myz 1 a 2 b 3 c 4 d 5 e 6 f
(integer) 6
127.0.0.1:6379> ZSCAN myz 0 MATCH *e*
1) "0"
2) 1) "e"
   2) "5"
```

#### 获取有序集合差异

```shell
127.0.0.1:6379> ZADD myz 1 a 2 b 3 c 4 d
(integer) 4
127.0.0.1:6379> ZADD myzz 5 e 6 b 7 c 8 f
(integer) 4
```

- ZDIFF numkeys key [key ...] [WITHSCORES] 比较第一个集合和其他集合之间的差异并返回差异的结果, 第一个集合为空或者第一个集合的所有成员在出现在其他集合中返回(empty array), 6.2.0 支持

  - 部分参数同 `ZDIFFSTORE`

- ZDIFFSTORE destination numkeys key [key ...] 比较第一个集合和其他集合之间的差异把差异结果存储到指定的集合中并返回指定集合的数量, 如果指定集合不存在则新建, 如果指定集合存在则覆盖指定集合, 0 表示未找到结果, 6.2.0 支持
  - numkeys 指定集合的数量, 值和 key 的数量不一致时返回语法错误 syntax error

```shell
127.0.0.1:6379> ZDIFF 2 myz myzz WITHSCORES
1) "a"
2) "1"
3) "d"
4) "4"

# 获取差异并存储到指定集合
127.0.0.1:6379> ZDIFFSTORE destst 2 myz myzz
(integer) 2
127.0.0.1:6379> ZRANGE destst 0 -1 WITHSCORES
1) "a"
2) "1"
3) "d"
4) "4"
```

#### 获取有序集合交集

```shell
127.0.0.1:6379> ZADD myz  1 a 2 b 3 c
(integer) 3
127.0.0.1:6379> ZADD myzz 4 b 5 c 6 d
(integer) 3
```

- ZINTER numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES] 返回多个有序集合的交集, key 不存在被当作空集合, 当给定集合中有一个空集合时返回结果也为空集合(empty array), 6.2.0 支持

  - 部分参数同 `ZINTERSTORE`

```shell
# 获取交集相同成员分值乘以权重后的求和
127.0.0.1:6379> ZINTER 2 myz myzz WEIGHTS 2 3 WITHSCORES
1) "b"
2) "16"
3) "c"
4) "21"
# 获取交集相同成员分值乘以权重后的最大值
127.0.0.1:6379> ZINTER 2 myz myzz WEIGHTS 2 3 AGGREGATE MAX WITHSCORES
1) "b"
2) "12"
3) "c"
4) "15"
```

- ZINTERCARD numkeys key [key ...] [LIMIT limit] 返回多个有序集合的交集的数量, 0 表示未找到结果, 7.0.0 支持

- ZINTERSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] 计算多个有序集合的交集将结果存储到指定集合并返回保存到指定集合的成员数量, 0 表示未找到结果
  - destination 指定存储结果集的集合名字
  - numkeys 指定集合的数量, 值和 key 的数量不一致时返回语法错误 syntax error
  - WEIGHTS 指定每个排序集合的权重, 每个集合中的成员的分值都会乘以这个权重, 默认为 1, 如果指定此项, 则值的数量必须和 numkeys 一致, 否则报语法错误
  - AGGREATE 指定结果集聚合的条件, 默认 SUM
    - SUM 结果集中保留相同成员和所有相同成员分值的求和
    - MIN 结果集中保留最小分值
    - MAX 结果集中保留最大分值

```shell
# 如果存在权重, 权重数量必须和 numkeys 保持一致
127.0.0.1:6379> ZINTERSTORE destst 2 myz myzz WEIGHTS 1
(error) ERR syntax error
# myz 和 myzz 权重都为 1
127.0.0.1:6379> ZINTERSTORE destst 2 myz myzz WEIGHTS 1 1
(integer) 2
# 结果集计算, b: myz:2 + myzz:4 = 6, c: myz:3 + myzz:5 = 8
127.0.0.1:6379> ZRANGE destst 0 -1 WITHSCORES
1) "c"
2) "8"
3) "b"
4) "6"

# 指定 myz 集合成员分值权重为 1, myzz 集合成员分值权重为 2
# myz 权重为 1, myzz 权重为 2
127.0.0.1:6379> ZINTERSTORE destst 2 myz myzz WEIGHTS 1 2
(integer) 2
# 结果集计算, b: myz:2 + myzz:4 * 2 = 10, c: myz:3 + myzz:5 * 2 = 13
127.0.0.1:6379> ZRANGE destst 0 -1 WITHSCORES
1) "c"
2) "13"
3) "b"
4) "10"

# 指定 myz 集合成员分值权重为 3, myzz 集合成员分值权重为 2
# myz 权重为 3, myzz 权重为 2
127.0.0.1:6379> ZINTERSTORE destst 2 myz myzz WEIGHTS 3 2
(integer) 2
# 结果集计算, b: myz:2 * 3 + myzz:4 * 2 = 14, c: myz:3 * 3 + myzz:5 * 2 = 19
127.0.0.1:6379> ZRANGE destst 0 -1 WITHSCORES
1) "c"
2) "19"
3) "b"
4) "14"

# 指定 myz 集合成员分值权重为 3, myzz 集合成员分值权重为 2, 结果集保存最小值
127.0.0.1:6379> ZINTERSTORE destst 2 myz myzz WEIGHTS 3 2 AGGREGATE MIN
(integer) 2
# 结果集计算, b: myz:2 * 3 = 6, myz:3 * 3 = 9
127.0.0.1:6379> ZRANGE destst 0 -1 WITHSCORES
1) "c"
2) "9"
3) "b"
4) "6"

# 指定 myz 集合成员分值权重为 3, myzz 集合成员分值权重为 2, 结果集保存最大值
127.0.0.1:6379> ZINTERSTORE destst 2 myz myzz WEIGHTS 3 2 AGGREGATE MAX
(integer) 2
# 结果集计算, b: myzz:4 * 2 = 8, myzz:5 * 2 = 10
127.0.0.1:6379> ZRANGE destst 0 -1 WITHSCORES
1) "c"
2) "10"
3) "b"
4) "8"
```

#### 获取有序集合并集

```shell
127.0.0.1:6379> ZADD myz 1 a 2 b 3 c 4 d
(integer) 4
127.0.0.1:6379> ZADD myz1 5 e 6 b 7 c 8 f
(integer) 4
```

- ZUNION numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES] 返回多个有序集合的并集并移除相同的成员只保留一个, 不存在的 key 被当作空集合, 集合都为空返回 (empty array), 6.2.0 支持

  - 部分参数同 `ZUNIONSTORE`

```shell
# 获取并集其他成员分值乘以权重和相同成员分值乘以权重后的求和
127.0.0.1:6379> ZUNION 2 myz myz1 WEIGHTS 2 3 WITHSCORES
 1) "a"
 2) "2"
 3) "d"
 4) "8"
 5) "e"
 6) "15"
 7) "b"
 8) "22"
 9) "f"
10) "24"
11) "c"
12) "27"
```

- ZUNIONSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] 计算多个有序集合的并集将结果存储到指定集合并返回保存到指定集合的成员数量, 如果指定集合不存在则新建, 如果指定集合已存在则覆盖指定集合, 0 表示未找到结果
  - 参数同 `ZINTERSTORE`

```shell
# 获取并集其他成员分值乘以权重和相同成员分值乘以权重后的最小值保存
127.0.0.1:6379> ZUNIONSTORE destst 2 myz myz1 WEIGHTS 3 4 AGGREGATE MIN
(integer) 6
127.0.0.1:6379> ZRANGE destst 0 -1 WITHSCORES
 1) "a"
 2) "3"
 3) "b"
 4) "6"
 5) "c"
 6) "9"
 7) "d"
 8) "12"
 9) "e"
10) "20"
11) "f"
12) "32"
```
