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

#### 遍历无序集合

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

#### 获取无序集合差异

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

#### 获取无序集合交集

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

#### 获取无序集合并集

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

### Zset 命令

Zset 和 set 一样也是 string 类型元素的集合, 且不允许重复的成员, 不同的是每个元素都会关联一个 double 类型的分数, Redis 正是通过分数来为集合中的成员进行从小到大的排序, 有序集合的成员是唯一的,但分数(score)却可以重复, 集合是通过哈希表实现的, 最大的成员数为 2^32-1(40 多亿)个成员.

> 大部分 Set 的 API 将首字母 s 换成 z 就可以使用, 这里只列出部分不一致的 API

#### 特殊标识符

- \-inf 负无穷大
- \+inf 正无穷大
- ( 包含当前字符
- [ 包含当前字符

#### 添加成员

- ZADD key [NX|XX] [GT|LT] [CH] [INCR] score member [score member ...] 添加更新成员, 通常只返回添加的新成员的数量

  - NX 仅添加新成员, 不再更新已存在的成员
  - XX 仅更新已经存在的成员, 不再添加新成员
  - GT 仅当新分数大于当前分数才更新已存在的成员, 此标志不阻止添加新成员
  - LT 仅当新分数小于当前分数才更新已存在的成员, 此标志不阻止添加新成员
  - CH 将 `zadd` 返回值统计新成员的添加数量修改为更改的元素总数, 包含更新已存在的数量和新添加的数量
  - INCR 此选项作用类似 `ZINCRBY`, 只能指定一个成员分数对加上指定的增量(可以为负数), 如果成员不存在从 0 开始计算, 并返回当前成员的最终分数, 多个成员分数对会报错

```shell
127.0.0.1:6379> ZADD myz 1 zhangsan 2 lisi  # 添加成员
(integer) 2
127.0.0.1:6379> ZCARD myz
(integer) 2
127.0.0.1:6379> ZRANDMEMBER myz 2 WITHSCORES
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
127.0.0.1:6379> ZRANDMEMBER myz 2 WITHSCORES
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
127.0.0.1:6379> ZRANDMEMBER myz 4 WITHSCORES
1) "zhaoliu"
2) "3.5"
3) "lisi"
4) "2"
5) "zhangsan"
6) "1.5"
7) "wangwu"
8) "1"

# LT  仅当新分数小于当前分数才更新已存在的成员, 不阻止添加新成员
127.0.0.1:6379> ZADD myz LT 0.5 lisi 4 sunqi  # 0.5 < 2 满足条件, 有新成员 4 sunqi 可添加
(integer) 1
127.0.0.1:6379> ZCARD myz
(integer) 5
127.0.0.1:6379> ZRANDMEMBER myz 5 WITHSCORES
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
127.0.0.1:6379> ZADD myz LT 1 lisi 8 sunqi  # 1 > 0.5, 8 > 4 不满足条件, 没有新成员可添加
(integer) 0
127.0.0.1:6379> ZCARD myz
(integer) 5
127.0.0.1:6379> ZRANDMEMBER myz 5 WITHSCORES
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

# GT 仅当新分数大于当前分数才更新已存在的成员, 不阻止添加新成员
127.0.0.1:6379> ZADD myz GT 1 zhangsan 2 qianba # 1 < 1.5 不满足条件, 有新成员 2 qianba 可添加
(integer) 1
127.0.0.1:6379> ZCARD myz
(integer) 6
127.0.0.1:6379> ZRANDMEMBER myz 6 WITHSCORES
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
127.0.0.1:6379> ZADD myz GT 3 zhangsan 2 qianba # 3 > 1.5 满足条件, 2 = 2 不满足条件, 没有新成员可添加
(integer) 0
127.0.0.1:6379> ZCARD myz
(integer) 6
127.0.0.1:6379> ZRANDMEMBER myz 6 WITHSCORES
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
127.0.0.1:6379> ZADD myz CH 1 zhangsan 5 qianba 1 hello # zhangsan 和 qianba 已存在, hello 为新添加成员,
(integer) 3

# INCR 同时只能指定一个成员分数对加上指定的增量(可以为负数), 如果成员不存在从 0 开始计算
# 多个会报错 (error) ERR INCR option supports a single increment-element pair
127.0.0.1:6379> ZADD myz INCR 10 zhangsan
"10"
127.0.0.1:6379> ZRANDMEMBER myz 1 WITHSCORES
1) "zhangsan"
2) "10"
```

- ZCOUNT key min max 统计有序集合中指定分数区间的成员数量, 0 表示找到结果

```shell
127.0.0.1:6379> ZRANDMEMBER myz 6 WITHSCORES
1) "zhaoliu"
2) "3.5"
3) "lisi"
4) "2.5"
5) "zhangsan"
6) "1.5"
7) "wangwu"
8) "1"
127.0.0.1:6379> ZCOUNT myz -inf +inf  # 获取分数 负无穷大 到 正无穷大 之间的数量
(integer) 4
127.0.0.1:6379> ZCOUNT myz 0 2  # 获取分数在 0 到 2 之间的数量
(integer) 2
127.0.0.1:6379> ZCOUNT myz 1.5 3  # 获取分数在 1.5 到 3 之间的数量
(integer) 2
```

- ZINCRBY key increment member 对指定成员的分数加上增量并返回修改后的分数, 如果指定成员不存在则添加新成员, 等同于 `ZADD key increment member`

```shell
127.0.0.1:6379> ZINCRBY myz 2 zhangsan  # 修改 zhangsan 的分数 + 2
"3.5"
127.0.0.1:6379> ZINCRBY myz 2 zhangsan1 # zhangsan1 不存在, 新添加
"2"
127.0.0.1:6379> ZINCRBY myz -2 zhaoliu  # 修改 zhaoliu 的分数减去 - 2
"1.5"
127.0.0.1:6379> ZRANDMEMBER myz 6 WITHSCORES
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

#### 获取有序集合交集

- ZINTERSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] 计算多个有序集合的交集, 并将交集存储到指定的集合中, 返回保存到指定集合的成员数量
  - numkeys 指定集合的数量, 值和 key 的数量不一致时返回语法错误 syntax error
  - WEIGHTS 指定每个排序集合的权重, 每个集合中的成员的分数都会乘以这个权重, 默认为 1, 如果指定此项, 则值的数量必须和 numkeys 一致, 否则报语法错误
  - AGGREATE 指定结果集聚合的条件, 默认 SUM
    - SUM 结果集中保留以相同成员的分数的和
    - MIN 结果集中保留最小分数
    - MAX 结果集中保留最大分数

```shell
127.0.0.1:6379> ZADD myz  1 a 2 b 3 c
(integer) 3
127.0.0.1:6379> ZADD myzz 4 b 5 c 6 d
(integer) 3
127.0.0.1:6379> ZINTERSTORE destst 2 myz myzz WEIGHTS 1 # 如果存在权重, 权重数量必须和 numkeys 保持一致
(error) ERR syntax error
127.0.0.1:6379> ZINTERSTORE destst 2 myz myzz WEIGHTS 1 1 # myz 和 myzz 权重都为 1
(integer) 2
# 结果集计算, b: myz:2 + myzz:4 = 6, c: myz:3 + myzz:5 = 8
127.0.0.1:6379> ZRANDMEMBER destst 2 WITHSCORES
1) "c"
2) "8"
3) "b"
4) "6"

# 指定 myz 集合成员分数权重为 1, myzz 集合成员分数权重为 2
127.0.0.1:6379> ZINTERSTORE destst 2 myz myzz WEIGHTS 1 2  # myz 权重为 1, myzz 权重为 2
(integer) 2
# 结果集计算, b: myz:2 + myzz:4 * 2 = 10, c: myz:3 + myzz:5 * 2 = 13
127.0.0.1:6379> ZRANDMEMBER destst 10 WITHSCORES
1) "c"
2) "13"
3) "b"
4) "10"

# 指定 myz 集合成员分数权重为 3, myzz 集合成员分数权重为 2
127.0.0.1:6379> ZINTERSTORE destst 2 myz myzz WEIGHTS 3 2 # myz 权重为 3, myzz 权重为 2
(integer) 2
# 结果集计算, b: myz:2 * 3 + myzz:4 * 2 = 14, c: myz:3 * 3 + myzz:5 * 2 = 19
127.0.0.1:6379> ZRANDMEMBER destst 10 WITHSCORES
1) "c"
2) "19"
3) "b"
4) "14"

# 指定 myz 集合成员分数权重为 3, myzz 集合成员分数权重为 2, 结果集保存最小值
127.0.0.1:6379> ZINTERSTORE destst 2 myz myzz WEIGHTS 3 2 AGGREGATE MIN
(integer) 2
# 结果集计算, b: myz:2 * 3 = 6, myz:3 * 3 = 9
127.0.0.1:6379> ZRANDMEMBER destst 10 WITHSCORES
1) "c"
2) "9"
3) "b"
4) "6"

# 指定 myz 集合成员分数权重为 3, myzz 集合成员分数权重为 2, 结果集保存最大值
127.0.0.1:6379> ZINTERSTORE destst 2 myz myzz WEIGHTS 3 2 AGGREGATE MAX
(integer) 2
# 结果集计算, b: myzz:4 * 2 = 8, myzz:5 * 2 = 10
127.0.0.1:6379> ZRANDMEMBER destst 10 WITHSCORES
1) "c"
2) "10"
3) "b"
4) "8"
```

- ZLEXCOUNT key min max 计算有序集合中指定字典区间内成员数量

```shell
127.0.0.1:6379> ZADD myz 1 a 2 b 3 c 4 d 5 e 6 f
(integer) 6
127.0.0.1:6379> ZLEXCOUNT myz - +
(integer) 6
127.0.0.1:6379> ZLEXCOUNT myz (b [f # 包含 b
(integer) 4
127.0.0.1:6379> ZLEXCOUNT myz [b [f # 包含 b
(integer) 5
127.0.0.1:6379> ZLEXCOUNT myz [d [f
(integer) 3
127.0.0.1:6379> ZLEXCOUNT myz [a [c
(integer) 3
```

#### 遍历有序集合

- ZRANGE key start stop [BYSCORE|BYLEX] [REV] [LIMIT offset count] [WITHSCORES] 遍历有序集合的指定区间, 并返回递增顺序的成员

  - start 起始指针, 0 表示第一个成员
  - stop 结束指针, -1 表示最后 1 个成员, -2 表示倒数第 2 个成员
  - BYSCORE start 指针要考量最高分数, stop 指针考量最低分数, start 必须大于等于 stop 才能返回内容
  - BYLEX 行为类似于 `ZRANGEBYLEX`, 返回 start 和 stop 字典闭合区间的成员
  - REV 反向排序结果
  - LIMIT 指定返回结果的偏移量, ERR syntax error, LIMIT is only supported in combination with either BYSCORE or BYLEX
  - WITHSCORES 返回结果的分数

> ZREVRANGE 6.2.0 开始废弃, 使用 `ZRANGE REV` 代替
> ZRANGEBYSCORE 6.2.0 开始废弃, 使用 `ZRANGE BYSCORE` 代替
> ZRANGEBYLEX 6.2.0 开始废弃, 使用 `ZRANGE BYLEX` 代替

```shell
127.0.0.1:6379> ZADD myz 10 a 9 b 13 c 7 d 11 e 6 f
(integer) 6
127.0.0.1:6379> ZRANGE myz 0 -1 WITHSCORES
 1) "f"
 2) "6"
 3) "d"
 4) "7"
 5) "b"
 6) "9"
 7) "a"
 8) "10"
 9) "e"
10) "11"
11) "c"
12) "13"
127.0.0.1:6379> ZRANGE myz (8 (13 BYSCORE WITHSCORES
1) "b"
2) "9"
3) "a"
4) "10"
5) "e"
6) "11"
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
127.0.0.1:6379> ZRANGE myz 0 +inf BYSCORE LIMIT 1 3 WITHSCORES
1) "d"
2) "7"
3) "b"
4) "9"
5) "a"
6) "10"
```

- ZRANK key member 返回有序集中指定成员的按递增顺序的排名, 从 0 开始计算, 如果指定成员不属于集合则返回 &lt;nil&gt;

```shell
127.0.0.1:6379> ZRANK myz h
(nil)
127.0.0.1:6379> zrank myz f
(integer) 0
127.0.0.1:6379> ZRANK myz a
(integer) 3
```
