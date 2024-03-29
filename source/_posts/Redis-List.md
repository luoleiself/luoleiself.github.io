---
title: Redis-List
date: 2022-11-09 18:03:19
categories:
  - [server, Redis]
tags:
  - Redis
---

### Lists 命令

> Redis 3.2 之后, List 数据类型底层使用 quickList 代替双向链表和压缩列表

List 是一个有序重复的双向链表, 按照添加的顺序排序, 可以添加一个元素到列表的头部(左边)、尾部(右边), 一个列表最多可以包含 2^32-1(40 多亿) 个元素.

List 类型的底层数据结构是由**双向链表**或**压缩列表**实现的

- 如果列表元素的个数小于 512 个,列表每个元素值都小于 64B 时, Redis 使用**压缩列表**作为底层数据结构
- 如果列表元素不满足上面的条件, Redis 使用**双向链表**作为底层数据结构

```yaml
# 配置底层数据结构存储数量限制
hash-max-listpack-entries 512
hash-max-listpack-value 64
```

#### 添加元素

LPUSH 和 LPUSHX 命令将多个元素**逆序**插入到列表头部

- LPUSH key element [element ...] 批量添加多个元素到列表头部并返回列表的长度, 列表为空或者不存在新建
- LPUSHX key element [element ...] 批量添加多个元素到**已存在**的列表头部并返回列表的长度, 列表为空或者不存在返回 0

```bash
127.0.0.1:6379> LPUSH list:zhang name1 name2 name3
(integer) 3
127.0.0.1:6379> LPUSHX list:zhang name1 name4 name5 name6
(integer) 7
127.0.0.1:6379> LPUSH list:zhang name1 name4 name5 name6
(integer) 11

# LPUSH 和 LPUSHX 命令将多个元素**逆序**插入到列表头部
127.0.0.1:6379> LRANGE list:zhang 0 11
 1) "name6"
 2) "name5"
 3) "name4"
 4) "name1"
 5) "name6"
 6) "name5"
 7) "name4"
 8) "name1"
 9) "name3"
10) "name2"
11) "name1"

# LPUSHX 对空列表或不存在的列表不进行操作
127.0.0.1:6379> LPUSHX list:zhang:1 name1 name2 name3
(integer) 0
127.0.0.1:6379> KEYS *
1) "list:zhang"
2) "age"
3) "name"
4) "hash:zhang"
```

<!-- more -->

RPUSH 和 RPUSHX 命令将多个元素**顺序**追加到列表尾部

- RPUSH key element [element ...] 批量追加多个元素到列表尾部并返回列表的长度, 列表为空或者不存在新建
- RPUSHX key element [element ...] 批量追加多个元素到**已存在**的列表尾部并返回列表的长度, 列表为空或者不存在返回 0

```bash
127.0.0.1:6379> RPUSH list:zhang name1 name2 name3
(integer) 3
127.0.0.1:6379> RPUSHX list:zhang name1 name4 name5 name6
(integer) 7
127.0.0.1:6379> RPUSH list:zhang name1 name4 name5 name6
(integer) 11

# RPUSH 和 RPUSHX 命令将多个元素**顺序**追加到列表尾部
127.0.0.1:6379> LRANGE list:zhang 0 11
 1) "name1"
 2) "name2"
 3) "name3"
 4) "name1"
 5) "name4"
 6) "name5"
 7) "name6"
 8) "name1"
 9) "name4"
10) "name5"
11) "name6"

# RPUSHX 对空列表或不存在的列表不进行操作
127.0.0.1:6379> RPUSHX list:zhang:1 name1 name2 name3
(integer) 0
127.0.0.1:6379> KEYS *
1) "list:zhang"
2) "age"
3) "name"
4) "hash:zhang"
```

##### 指定位置添加

- LSET key index element 设置列表指定索引的值, 操作成功返回 ok, 列表为空或者不存在或者索引参数超出范围返回错误信息

- LINSERT key BEFORE|AFTER pivot element 在列表的指定元素第 1 次出现的前或后添加元素, 列表为空或者不存在不执行任何操作, 否则未找到指定元素返回 -1, 元素添加成功返回列表的长度
  - BEFORE | AFTER 添加的位置
  - pivot 查找的基准元素
  - element 插入的元素

```bash
# 向 list 中第 1 次出现 a 的前面添加 hello
127.0.0.1:6379> LINSERT list BEFORE a hello
(integer) 5
127.0.0.1:6379> LRANGE list 0 -1 # 遍历 list
1) "hello"
2) "a"
3) "d"
4) "c"
5) "b"
```

#### 删除元素

##### 修剪列表

- LTRIM key start stop 对列表不包含在 start 到 stop 区间的元素进行删除, 执行成功返回 ok
  - start, end 只支持整数, 其他类型会报错

```bash
127.0.0.1:6379> LPUSH mylist hello world gg yy hehe haha
(integer) 6
127.0.0.1:6379> LTRIM mylist - +
(error) ERR value is not an integer or out of range
127.0.0.1:6379> LTRIM mylist -inf +inf
(error) ERR value is not an integer or out of range
127.0.0.1:6379> LTRIM mylist [1 [4
(error) ERR value is not an integer or out of range
127.0.0.1:6379> LTRIM mylist 1 4.5
(error) ERR value is not an integer or out of range
127.0.0.1:6379> LTRIM mylist 1 4
OK
127.0.0.1:6379> LRANGE mylist 0 -1
1) "hehe"
2) "yy"
3) "gg"
4) "world"
```

##### 批量移除相同元素

- LREM key count element 移除列表指定数量的元素并返回移除元素的数量, 列表为空或者不存在返回 0
  - count > 0 从列表头部开始向尾部搜索, 移除与 element 相等的元素, 数量为 count
  - count < 0 从列表尾部开始向头部搜索, 移除与 element 相等的元素, 数量为 count 的绝对值
  - count = 0 移除列表中与 element 相等的所有的元素

```bash
127.0.0.1:6379> KEYS *  # 获取当前数据库的 key
1) "sex"
2) "addr"
3) "name"
4) "runoob"
5) "age"
127.0.0.1:6379> LPUSH list a b a d a c d a b d  # 创建列表 list
(integer) 10
127.0.0.1:6379> LREM list 3 a   # 移除列表中 3 个元素 a
(integer) 3
127.0.0.1:6379> LRANGE list 0 7
1) "d"
2) "b"
3) "d"
4) "c"
5) "d"
6) "b"
7) "a"
```

##### 批量移除相邻元素

- LPOP key [count] 移除并返回列表头部指定数量的元素, count 默认为 1, 列表为空或者不存在返回 \<nil\>
- BLPOP key [key...] timeout 阻塞版的 `LPOP`, 从多个列表中第 1 个非空列表中的头部移除并返回 1 个元素, 如果列表为空会阻塞列表直到等待超时或发现可弹出元素为止, 如果列表为空或者超时返回 \<nil\>
  否则, 返回 1 个含有两个元素的列表, 第 1 个元素是被弹出元素所属的列表, 第 2 个元素是被弹出的元素

- RPOP key [count] 移除并返回列表尾部指定数量的元素, count 默认为 1, 列表为空或者不存在返回 \<nil\>
- BRPOP key [key ...] timeout 阻塞版的 `RPOP`, 从多个列表中第 1 个非空列表中的尾部移除并返回 1 个元素, 返回值 `BLPOP`

```bash
# 向 list 尾部添加元素
127.0.0.1:6379> RPUSH list c d e
(integer) 3
# newlist 为空或者不存在, 会删除 list 的尾部的元素
127.0.0.1:6379> BRPOP newlist list 0
1) "list"
2) "e"
# 向 newlist 尾部添加元素
127.0.0.1:6379> RPUSH newlist g a b f
(integer) 4
# list 和 newlist 都是非空的列表, 会删除返回 list 的头部的元素
127.0.0.1:6379> BLPOP list newlist 0
1) "list"
2) "c"
127.0.0.1:6379> del list  # 删除 list
(integer) 1
# list 为空或者不存在, 会删除返回 newlist 的头部的元素
127.0.0.1:6379> BLPOP list newlist 0
1) "newlist"
2) "g"
```

- LMPOP numkeys key [key ...] LEFT|RIGHT [COUNT count] 从多个列表中第 1 个非空列表中指定位置批量移除指定数量的元素并返回操作成功的 key 和移除的元素, 如果列表都为空或者不存在返回 \<nil\>, Redis 7.0.0 支持
  - numkeys 指定列表名的数量, 值和 key 的数量不一致时返回语法错误
  - LEFT | RIGHT 移除元素的位置
  - COUNT count 移除元素的数量, 默认为 1
- BLMPOP timeout numkeys key [key ...] LEFT|RIGHT [COUNT count] 阻塞版的 `LMPOP`, 列表为空时会阻塞直到等待超时或发现可弹出元素为止, Redis 7.0.0 支持

```bash
127.0.0.1:6379> RPUSH list a b c d e f  # 创建列表 list
(integer) 6
127.0.0.1:6379> RPUSH newlist 1 2 3 4 5 6 # 创建列表 newlist
(integer) 6
# list 不为空, 从 list 的头部删除返回 3 个 元素
127.0.0.1:6379> LMPOP 2 list newlist LEFT COUNT 3
1) "list"
2) 1) "a"
   2) "b"
   3) "c"
127.0.0.1:6379> del list  # 删除列表 list
(integer) 1
# list 为空, 从 newlist 的尾部删除返回 3 个元素
127.0.0.1:6379> LMPOP 2 list newlist right COUNT 3
1) "newlist"
2) 1) "6"
   2) "5"
   3) "4"
```

#### 删除元素添加到其他列表

- BLMOVE source destination LEFT|RIGHT LEFT|RIGHT timeout 阻塞版的 `LMOVE`, 如果列表为空会阻塞直到等待超时或发现可弹出元素为止, 6.2 开始可用代替 `BRPOPLPUSH`
  - LEFT|RIGHT 分别表示源列表移除元素的位置, 目标列表添加元素的位置
  - timeout 超时时间(单位秒)

> BRPOPLPUSH source destination timeout 移除列表的最后一个元素添加到另一个列表的头部并返回操作的元素, 如果列表为空会阻塞列表直到等待超时或发现可弹出元素为止, 6.2 开始废弃

```bash
# 向 list 尾部添加元素
127.0.0.1:6379> RPUSH list a b c
(integer) 3
# 向 newlist 尾部添加元素
127.0.0.1:6379> RPUSH newlist 1 2 3
(integer) 3
# 从 list 头部移除一个元素添加到 newlist 的尾部
127.0.0.1:6379> BLMOVE list newlist LEFT RIGHT 0
"a"
# 遍历 newlist
127.0.0.1:6379> LRANGE newlist 0 -1
1) "1"
2) "2"
3) "3"
4) "a"

# 阻塞执行等待超时或者有可操作元素
127.0.0.1:6379> BRPOPLPUSH list newlist 10
(nil)
(10.05s)
# 向 list 尾部添加元素
127.0.0.1:6379> RPUSH list a b c d
(integer) 4
# 从 list 尾部移除一个元素添加到 newlist 的头部
127.0.0.1:6379> BRPOPLPUSH list newlist
"d"
```

- LMOVE source destination LEFT|RIGHT LEFT|RIGHT 移除源列表中的头部或者尾部的元素添加到目标列表的头部或者尾部并返回移除的元素, 源列表为空或者不存在返回 \<nil\> 6.2 开始可用代替 `RPOPLPUSH`
  - LEFT|RIGHT 分别表示源列表移除元素的位置, 目标列表添加元素的位置

> RPOPLPUSH source destination 移除列表的最后一个元素添加到另一个列表的头部并返回操作的元素, 源列表为空或者不存在返回 \<nil\> 6.2 开始废弃

```bash
# 向 list 尾部添加元素
127.0.0.1:6379> RPUSH list a b c d e f g
(integer) 7
# 从 list 头部移除元素 a 添加到 newlist 的头部
127.0.0.1:6379> LMOVE list newlist LEFT LEFT
"a"
127.0.0.1:6379> LRANGE newlist 0 -1 # 遍历 newlist
1) "a"
# 从 list 头部移除元素 b 添加到 newlist 的尾部
127.0.0.1:6379> LMOVE list newlist LEFT RIGHT
"b"
127.0.0.1:6379> LRANGE newlist 0 -1  # 遍历 newlist
1) "a"
2) "b"
# 从 list 尾部移除元素 g 添加到 newlist 的头部
127.0.0.1:6379> LMOVE list newlist RIGHT LEFT
"g"
127.0.0.1:6379> LRANGE newlist 0 -1  # 遍历 newlist
1) "g"
2) "a"
3) "b"
# 从 list 尾部移除元素 f 添加到 newlist 的头部
127.0.0.1:6379> LMOVE list newlist RIGHT RIGHT
"f"
127.0.0.1:6379> LRANGE newlist 0 -1  # 遍历 newlist
1) "g"
2) "a"
3) "b"
4) "f"
```

#### 查找遍历列表

- LLEN key 返回列表长度, 0 表示列表为空或者不存在

- LRANGE key start stop 遍历列表指定区间的元素, 列表为空或者不存在返回 (empty array)
  - start, end 支持整数, 其他类型会报错

```bash
127.0.0.1:6379> LPUSH mylist hello world gg yy hehe haha
(integer) 6
127.0.0.1:6379> LRANGE mylist - +
(error) ERR value is not an integer or out of range
127.0.0.1:6379> LRANGE mylist -inf +inf
(error) ERR value is not an integer or out of range
127.0.0.1:6379> LRANGE mylist 1 4.5
(error) ERR value is not an integer or out of range
127.0.0.1:6379> LRANGE mylist (1 (4
(error) ERR value is not an integer or out of range
127.0.0.1:6379> LRANGE mylist [1 [4
(error) ERR value is not an integer or out of range
127.0.0.1:6379> LRANGE mylist 1 4
1) "hehe"
2) "yy"
3) "gg"
4) "world"
```

##### 查找指定索引元素

- LINDEX key index 获取列表中指定索引的元素, 如果列表或者索引不存在返回 \<nil\>

```bash
127.0.0.1:6379> LRANGE list:zhang 0 11
 1) "li10"
 2) "li9"
 3) "li8"
 4) "li7"
 5) "li6"
 6) "li5"
 7) "li4"
 8) "li3"
 9) "li2"
10) "li1"
127.0.0.1:6379> LINDEX list:zhang 11
(nil)
127.0.0.1:6379> LINDEX list:zhang 5
"li5"
127.0.0.1:6379> LINDEX list:zhang 8
"li2"
```

##### 查找匹配项的下标

- LPOS key element [RANK rank] [COUNT num-matches] [MAXLEN len] 返回列表中匹配元素的下标, 列表为空或者不存在返回 \<nil\>

  - RANK 指定匹配项第几次出现的下标, 默认为 1
  - COUNT 指定匹配项的下标的个数, 默认为 1
  - MAXLEN 指定命令对列表项进行比较的次数

```bash
# 创建列表 list
127.0.0.1:6379> RPUSH list a b c a b c b d c a b a c d b c a d
(integer) 18
# 查找元素 b 第 1 次出现的下标
127.0.0.1:6379> LPOS list b
(integer) 1
# 查找元素 b 第 3 次出现的下标
127.0.0.1:6379> LPOS list b RANK 3
(integer) 6
# 查找元素 b 出现 3 次的下标
127.0.0.1:6379> LPOS list b COUNT 3
1) (integer) 1
2) (integer) 4
3) (integer) 6
# 查找元素 b 第 2 次开始出现 3 次的下标
127.0.0.1:6379> LPOS list b RANK 2 COUNT 3
1) (integer) 4
2) (integer) 6
3) (integer) 10
# 查找元素 b 第 3 次开始出现 4 次的下标
127.0.0.1:6379> LPOS list b RANK 3 COUNT 4
1) (integer) 6
2) (integer) 10
3) (integer) 14
# 查找 10 次元素 b 第 2 次开始出现 4 次的下标
127.0.0.1:6379> LPOS list b RANK 2 COUNT 4 MAXLEN 10
1) (integer) 4
2) (integer) 6
# 查找 15 次元素 b 第 3 次开始出现 2 次的下标
127.0.0.1:6379> LPOS list b RANK 3 COUNT 2 MAXLEN 15
1) (integer) 6
2) (integer) 10
```

#### 使用案例

##### 抢红包

将抢红包的用户信息用 list 存储作为队列的形式保证先进先出的规则
