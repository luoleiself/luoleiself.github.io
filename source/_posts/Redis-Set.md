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

- SCARD key 获取集合成员的数量, 集合为空或者不存在返回 0
- SADD key member [member ...] 向集合中添加多个成员并返回添加成功的数量, 0 表示有重复成员
- SMEMBERS key 获取集合中所有的成员, 集合为空或者不存在返回 (empty array)
- SISMEMBER key member 判断 member 是不是集合的成员, 1 是, 0 不是或者集合为空或者不存在
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

