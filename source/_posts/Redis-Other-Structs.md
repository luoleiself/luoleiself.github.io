---
title: Redis-Other-Structs
date: 2022-11-09 18:03:19
categories:
  - [server, Redis]
tags:
  - Redis
---

### HyperLogLog

HyperLogLog 是用来做基数统计的算法, 优点是在输入元素的数量或者体积非常大时, 计算基数所需的空间总是固定的、并且是很小的. 每个 HyperLogLog 键只需要花费 12KB 内存, 就可以计算接近 2^64 个不同元素的基数, 因为 HyperLogLog 只会根据输入元素来计算基数, 而不会储存输入元素本身

HyperLogLog 返回的基数不精确, 近似于 0.81% 的标准误差

> 比如数据集 {1, 3, 5, 7, 5, 7, 8}, 那么这个数据集的基数集为 {1, 3, 5 ,7, 8}, 基数(不重复元素)为 5. 基数估计就是在误差可接受的范围内，快速计算基数

- PFADD key [element [element ...]] 添加元素

- PFCOUNT key [key ...] 根据 key 计算基数并返回, 0 表示 key 不存在

```shell
127.0.0.1:6379> PFADD hll foo bar zap
(integer) 1
127.0.0.1:6379> PFADD hll zap zap zap
(integer) 0
127.0.0.1:6379> PFADD hll foo bar
(integer) 0
127.0.0.1:6379> PFCOUNT hll
(integer) 3
127.0.0.1:6379> PFADD other-hll 1 2 3
(integer) 1
127.0.0.1:6379> PFCOUNT hll other-hll
(integer) 6
```

<!-- more -->

- PFDEBUG subcommand key 内部命令, 一般用于开发和测试 Redis
- PFSELFTEST 内部命令, 一般用于开发和测试 Redis
- PFMERGE destkey sourcekey [sourcekey ...] 将多个 HyperLogLog 值合并为 1 个唯一值, 该值将近似于源 HyperLogLog 结构的观察集的并集的基数, 如果 destkey 不存在则新建, 如果 destkey 已存在则将期作为源集之一, 其基数将包含在计算的 HyperLogLog 的基数中

```shell
127.0.0.1:6379> PFADD hll foo bar zap a
(integer) 1
127.0.0.1:6379> PFADD other-hll a b c foo
(integer) 1
# 不存在的 destkey
127.0.0.1:6379> PFMERGE res-hll hll other-hll
OK
127.0.0.1:6379> PFCOUNT res-hll
(integer) 6

# 已存在的 destkey
127.0.0.1:6379> PFADD ex-hll gg yy hehe haha
(integer) 1
127.0.0.1:6379> PFMERGE ex-hll hll other-hll
OK
127.0.0.1:6379> PFCOUNT ex-hll  # 将源集也计算在内
(integer) 10
```

### Geospatial

Redis 地理空间, 该类型就是元素的 2 维坐标, 在地图上就是经纬度. Redis 基于该类型, 提供了经纬度设置、查询、范围查询、距离查询、经纬度 Hash 等常见操作
Geospatial 底层实现原理实现为 zset 类型, 可以使用 zset 的方法

> 比如微信的朋友圈查找附近的人, 或者游戏中获取附近的游戏玩家

#### 规定

地球两极不能添加
EPSG:900913 / EPSG:3785 / OSGEO:41001 标准规定

- 有效经度在 -180 度到 180 度
- 有效维度在 -85.05112878 度到 85.05112878 度

#### 成员操作

##### 添加成员

- GEOADD key [NX|XX] [CH] longitude latitude member [longitude latitude member ...] 添加地理位置信息(经度、纬度、名称)到指定集合中, 通常只返回添加的新成员的数量
  - NX 仅添加新成员, 不再更新已存在的成员
  - XX 仅更新已经存在的成员, 不再添加新成员
  - CH 将 `GEOADD` 返回值统计新成员的添加数量修改为更改的成员总数, 包含更新已存在的数量和新添加的数量

```shell
127.0.0.1:6379> GEOADD citys 116.405285 39.904989 beijing 121.472644 31.231706 shanghai
(integer) 2
127.0.0.1:6379> GEOADD citys 113.280637 23.125178 guangzhou 114.05571 22.52245 shenzhen
(integer) 2
127.0.0.1:6379> GEOADD citys 106.504962 29.533155 chongqing 112.34882 16.831039 sansha
(integer) 2
127.0.0.1:6379> GEOADD citys 113.665412 34.757975 zhengzhou 103.823557 36.058039 lanzhou
(integer) 2
127.0.0.1:6379> GEOADD citys 108.948024 34.263161 xian 115.892151 28.676493 nanchang
(integer) 2
127.0.0.1:6379> ZCARD citys # 使用 zcard 获取集合数量
(integer) 10

# 仅更新成员信息并返回更新的数量
127.0.0.1:6379> GEOADD citys XX CH 113.88311 22.55371 shenzhen
(integer) 1

# 使用 zset 的 ZRANGE 遍历集合
127.0.0.1:6379> ZRANGE citys 0 -1 WITHSCORES
 1) "sansha"
 2) "3974440648358025"
 3) "chongqing"
 4) "4026042117887371"
 5) "shenzhen"
 6) "4046340107214121"
 7) "guangzhou"
 8) "4046533764066819"
 9) "shanghai"
10) "4054803464817068"
11) "beijing"
12) "4069885370671010"
```

##### 经纬度

###### 获取经纬度

- GEOPOS key member [member ...] 返回指定成员的经纬度信息, 如果集合为空或者不存在或者指定成员不存在则返回 &lt;nil&gt;

```shell
127.0.0.1:6379> GEOPOS city1 beijing shanghai
1) (nil)
2) (nil)
127.0.0.1:6379> GEOPOS citys beijing hello
1) 1) "116.40528291463851929"
   2) "39.9049884229125027"
2) (nil)
127.0.0.1:6379> GEOPOS citys beijing shanghai
1) 1) "116.40528291463851929"
   2) "39.9049884229125027"
2) 1) "121.47264629602432251"
   2) "31.23170490709807012"
```

###### 获取经纬度 hash

- GEOHASH key member [member ...] 返回指定成员的经纬度信息的 hash 编码后的字符串表示, 如果集合为空或者不存在或者指定成员不存在则返回 &lt;nil&gt;

```shell
127.0.0.1:6379> GEOHASH citys hello world
1) (nil)
2) (nil)
127.0.0.1:6379> GEOHASH citys beijing shanghai sansha
1) "wx4g0b7xrt0"
2) "wtw3sjt9vg0"
3) "w6zzsxczvm0"
```

##### 获取指定成员间距离

- GEODIST key member1 member2 [M|KM|FT|MI] 返回集合指定成员之间的距离, 默认距离单位 M, 如果集合为空或者不存在或者指定成员不存在则返回 &lt;nil&gt;
  - M 米
  - KM 公里
  - FT 英里
  - MI 英尺

```shell
127.0.0.1:6379> GEODIST citys1 beijing shanghai KM
(nil)
127.0.0.1:6379> GEODIST citys beijing hello KM
(nil)
127.0.0.1:6379> GEODIST citys beijing shanghai KM
"1067.5980"
127.0.0.1:6379> GEODIST citys beijing shanghai
"1067597.9668"
127.0.0.1:6379> GEODIST citys beijing shanghai FT
"3502618.0013"
127.0.0.1:6379> GEODIST citys beijing shanghai MI
"663.3763"
127.0.0.1:6379> GEODIST citys beijing sansha KM
"2596.1770"
```

##### 范围搜索

- GEOSEARCH key <FROMMEMBER member|FROMLONLAT longitude latitude> <BYRADIUS radius <M|KM|FT|MI>|BYBOX width height <M|KM|FT|MI>> [ASC|DESC] [COUNT count [ANY]] [WITHCOORD] [WITHDIST] [WITHHASH] 返回集合中符合属于给定形状区域的边界内的成员的信息, 默认返回所有的匹配项且没有排序, 如果集合为空或者不存在返回 (empty array), 6.2.0 支持

  - FROMMEMBER 使用给定的已存在的成员进行搜索, 如果指定成员不属于非空集合则报错 ERR could not decode requested zset member
  - FROMLONLAT 使用给定的 longitude 和 latitude 搜索, 如果指定的经纬度超出范围则报错 ERR invalid longitude,latitude pair
  - BYRADIUS 行为类似于 `GEORADIUS`, 根据给定半径的圆形区域内搜索
  - BYBOX 根据给定的 width 和 height 在轴对齐的矩形区域内搜索
  - ASC 相对于中心点, 从近到远对返回的匹配项进行排序
  - DESC 相对于中心点, 从远到近对返回的匹配项进行排序
  - COUNT 限制匹配项的数量
  - ANY 找到足够的匹配项就立刻返回, 有可能匹配项不是最接近指定点的位置, 但是服务器生成这些结果所投入的精力要少的多
    未指定此参数时, 搜索命令将执行与指定匹配区域匹配项数量成正比的操作, 并对其进行排序, 因此, 即使返回少量结果, 使用非常小的 count 查询非常大的区域也可能会很慢.
  - WITHCOORD 返回匹配项的经纬度
  - WITHDIST 返回匹配项距离指定中心点的距离, 距离单位和指定圆形区域搜索或者矩形区域搜索的单位相同
  - WITHHASH 返回匹配项 `GEOHASH` 编码的无符号整数

> GEORADIUS 6.2.0 开始废弃, 使用 `GEOSEARCH BYRADIUS` 代替
> GEORADIUS_RO 6.2.0 开始废弃, 使用 `GEOSEARCH BYRADIUS` 代替
> GEORADIUSBYMEMBER 6.2.0 开始废弃, 使用 `GEOSEARCH FROMMEMBER member BYRADIUS radius` 代替
> GEORADIUSBYMEMBER_RO 6.2.0 开始废弃, 使用 `GEOSEARCH FROMMEMBER member BYRADIUS radius` 代替

```shell
# 集合为空或者不存在返回 empty array
127.0.0.1:6379> GEOSEARCH citys1 FROMMEMBER zhengzhou BYRADIUS 500 KM
(empty array)
# 指定成员不属于非空集合报错
127.0.0.1:6379> GEOSEARCH citys FROMMEMBER wuhan BYRADIUS 500 KM
(error) ERR could not decode requested zset member
# 经纬度超出范围限制报错
127.0.0.1:6379> GEOSEARCH citys FROMLONLAT -220.10 26.31 BYRADIUS 500 KM
(error) ERR invalid longitude,latitude pair -220.100000,26.310000

# 以指定成员位置为中心点在半径 700KM 的圆形区域内搜索符合条件的成员
# 返回匹配项的经纬度和与中心点的距离的的信息
# 并按相对于中心点, 从远到近对返回的匹配项进行排序
127.0.0.1:6379> GEOSEARCH citys FROMMEMBER nanchang BYRADIUS 700 KM DESC WITHCOORD WITHDIST
1) 1) "guangzhou"
   2) "670.3933"
   3) 1) "113.28063815832138062"
      2) "23.12517743834835215"
2) 1) "shanghai"
   2) "608.1426"
   3) 1) "121.47264629602432251"
      2) "31.23170490709807012"
3) 1) "nanchang"
   2) "0.0000"
   3) 1) "115.89214950799942017"
      2) "28.67649306190701708"

# 以指定经纬度为中心点(wuhan)在 宽 1500KM, 高 1500KM的轴对齐的矩形区域内搜索符合条件的成员
# 返回匹配项的经纬度和与中心点的距离的信息和对 GEOHASH 编码的无符号整数
# 并按相对于中心点, 从近到远对返回的匹配项进行排序后最多返回 3 条
127.0.0.1:6379> GEOSEARCH citys FROMLONLAT 114.298572 30.584355 BYBOX 1500 1500 KM ASC COUNT 3 WITHCOORD WITHDIST WITHHASH
1) 1) "nanchang"
   2) "262.2278"
   3) (integer) 4051506205099900
   4) 1) "115.89214950799942017"
      2) "28.67649306190701708"
2) 1) "zhengzhou"
   2) "467.9826"
   3) (integer) 4064942392187921
   4) 1) "113.66541177034378052"
      2) "34.75797603259534441"
3) 1) "xian"
   2) "647.7152"
   3) (integer) 4040115616141630
   4) 1) "108.94802302122116089"
      2) "34.2631604414749944"
# COUNT 限制匹配项的数量, 第 4 条不返回
# 4) 1) "shanghai"
#    2) "688.2826"
#    3) (integer) 4054803464817068
#    4) 1) "121.47264629602432251"
#       2) "31.23170490709807012"
```

##### 范围搜索存储

- GEOSEARCHSTORE destination source <FROMMEMBER member|FROMLONLAT longitude latitude> <BYRADIUS radius <M|KM|FT|MI>|BYBOX width height <M|KM|FT|MI>> [ASC|DESC] [COUNT count [ANY]] [STOREDIST] 命令同 `GEOSEARCH`, 区别是将结果存储到指定排序集合并返回指定集合的数量, 默认存储匹配项的名称和匹配项 `GEOHASH` 的无符号整数, 如果指定集合为空或者不存在则新建, 如果指定集合已存在则覆盖指定集合, 6.2.0 支持
  - 部分参数同 `GEOSEARCH`
  - STOREDIST 返回匹配项距离中心点的距离

```shell
# 存储匹配项的 GEOHASH 的无符号整数到指定排序集合中, 返回指定集合的数量
127.0.0.1:6379> GEOSEARCHSTORE destst citys FROMLONLAT 114.298572 30.584355 BYBOX 1500 1500 KM
(integer) 4
127.0.0.1:6379> ZRANGE dest 0 -1 WITHSCORES
(empty array)
127.0.0.1:6379> ZRANGE destst 0 -1 WITHSCORES
1) "xian"
2) "4040115616141630"
3) "nanchang"
4) "4051506205099900"
5) "shanghai"
6) "4054803464817068"
7) "zhengzhou"
8) "4064942392187921"

# 按照匹配项距离中心点的距离排序后存储到指定集合中, 返回指定集合的数量
127.0.0.1:6379> GEOSEARCHSTORE destst citys FROMLONLAT 114.298572 30.584355 BYBOX 1500 1500 KM ASC STOREDIST
(integer) 4
127.0.0.1:6379> ZRANGE destst 0 -1 WITHSCORES
1) "nanchang"
2) "262.22782004072326"
3) "zhengzhou"
4) "467.98262044388139"
5) "xian"
6) "647.71515052090376"
7) "shanghai"
8) "688.28260742289308"
127.0.0.1:6379> GEOSEARCHSTORE destst citys FROMLONLAT 114.298572 30.584355 BYBOX 1500 1500 KM DESC STOREDIST
(integer) 4
127.0.0.1:6379> ZRANGE destst 0 -1 WITHSCORES
1) "nanchang"
2) "262.22782004072326"
3) "zhengzhou"
4) "467.98262044388139"
5) "xian"
6) "647.71515052090376"
7) "shanghai"
8) "688.28260742289308"
```

### Bitmaps

Redis Bitmaps 是字符串数据类型的一种扩展, 可以将字符串视为位向量, 对 1 个或者多个字符串进行逐位操作.
可以把 Bitmaps 想象成一个以位为单位的数组, 数组中的每个单元只能存储 0 或 1, 数组的下标在 Bitmaps 中叫做偏移量, 单个 bitmaps 的最大长度是 512MB,即 2^32 个比特位.

> 集合成员对应于整数 0-N 的情况的有效集合表示
> 对象权限, 其中每一位表示一个特定的权限, 类似于文件系统存储权限的方式

#### 操作值

- SETBIT key offset value 设置指定偏移量的值 0 或 1, 当 key 的指定偏移量的值从 1 修改为 0 时返回 1, 其他情况返回 0

  - offset 下标
  - value 值

- GETBIT key offset 获取指定便宜连的值, 如果 key 不存在或者偏移量不存在返回 0
- BITCOUNT key [start end [BYTE|BIT]] 统计 key 中状态为 1 的下标数量, 默认以 BYTE 为单位, 如果 key 不存在或者
  - BYTE 以字节(8bit)为单位, 默认
  - BIT 以 bit 为单位

```shell
# 当指定偏移量的值从 1 修改为 0 则返回 1, 其他情况返回 0
127.0.0.1:6379> SETBIT user 0 0
(integer) 0
127.0.0.1:6379> SETBIT user 0 1
(integer) 0
127.0.0.1:6379> SETBIT user 1 1
(integer) 0
127.0.0.1:6379> SETBIT user 1 0
(integer) 1

# 获取指定偏移量的值
127.0.0.1:6379> GETBIT user 6
(integer) 0
127.0.0.1:6379> GETBIT user 0
(integer) 1
127.0.0.1:6379> GETBIT user 1
(integer) 0

# 统计状态为 1 的下标数量
127.0.0.1:6379> BITCOUNT user 0 0
(integer) 5
127.0.0.1:6379> BITCOUNT user 0 0 BYTE
(integer) 5
127.0.0.1:6379> BITCOUNT user 0 0 BIT
(integer) 0
```

- BITPOS key bit [start [end [BYTE|BIT]]] 返回 key 中指定区间内的第 1 个符合指定值的偏移量
  - BYTE|BIT 参数同 `BITCOUNT`
  - 当 bit 为 1, 如果 key 不存在或者全为 0 值则返回 -1
  - 当 bit 为 0, 如果 key 不存在或者全为 0 值则返回 0

```shell
127.0.0.1:6379> BITPOS user1 1
(integer) -1
127.0.0.1:6379> BITPOS user1 0
(integer) 0

127.0.0.1:6379> BITPOS user 1
(integer) 1
# 以字节为单位查找第 1 个值为 0 的偏移量
127.0.0.1:6379> BITPOS user 0 1 -1 BYTE
(integer) -1
# 以bit为单位查找第 1 个值为 0 的偏移量
127.0.0.1:6379> BITPOS user 0 1 -1 BIT
(integer) 6
```

#### 位运算

- BITOP operation destkey key [key ...] 在多个 key 之间按照 操作符 逐位操作, 将结果追加存储到指定 key 中, key 为空或者不存在返回 0

  - operation 操作符

    - AND 与
    - OR 或
    - XOR 异或
    - NOT 非, 此参数只能传入 1 个 key, 多个 key 报错 ERR BITOP NOT must be called with a single source key.

  - destst 结果集存储的名称

```shell
BITOP AND destst key1 key2 ....keyN
127.0.0.1:6379> SETBIT key1 0 1
(integer) 0
127.0.0.1:6379> SETBIT key1 1 1
(integer) 0
127.0.0.1:6379> SETBIT key1 2 1
(integer) 0
127.0.0.1:6379> SETBIT key2 0 0
(integer) 0
127.0.0.1:6379> SETBIT key2 1 0
(integer) 0
127.0.0.1:6379> SETBIT key2 2 0
(integer) 0

# BITCOUNT 统计指定区间内的值为 1 的下标数量
127.0.0.1:6379> BITCOUNT key1 0 -1 BIT
(integer) 3
127.0.0.1:6379> BITCOUNT key2 0 -1 BIT
(integer) 0
127.0.0.1:6379> BITOP AND dest key1 key2
(integer) 1
127.0.0.1:6379> BITCOUNT dest 0 -1 BIT
(integer) 0
127.0.0.1:6379> BITOP OR dest key1 key2
(integer) 1
127.0.0.1:6379> BITCOUNT dest 0 -1 BIT
(integer) 3
127.0.0.1:6379> BITOP XOR dest key1 key2
(integer) 1
127.0.0.1:6379> BITCOUNT dest 0 -1 BIT
(integer) 3
127.0.0.1:6379> BITOP NOT dest key1
(integer) 1
127.0.0.1:6379> BITCOUNT dest 0 -1 BIT
(integer) 5
127.0.0.1:6379> BITOP NOT dest key2
(integer) 1
127.0.0.1:6379> BITCOUNT dest 0 -1 BIT
(integer) 8
```

#### 使用案例

##### 记录用户是否在线

记录用户是否在线时, 将用户 uid 作为偏移量, 如果已登录设置为 1, 退出设置为 0

> SETBIT login_stat {uid} 1|0

```shell
# 使用 SETBIT 以每个 id 作为偏移量记录一个状态
127.0.0.1:6379> SETBIT login_stat 101 1
(integer) 0
127.0.0.1:6379> SETBIT login_stat 102 1
(integer) 0
127.0.0.1:6379> SETBIT login_stat 103 0
(integer) 0
127.0.0.1:6379> SETBIT login_stat 104 0
(integer) 0
127.0.0.1:6379> SETBIT login_stat 105 1
(integer) 0
127.0.0.1:6379> SETBIT login_stat 106 0
(integer) 0
127.0.0.1:6379> BITCOUNT login_stat 0 -1 BIT
(integer) 3
# 登录状态修改为 1
127.0.0.1:6379> SETBIT login_stat 103 1
(integer) 0
127.0.0.1:6379> SETBIT login_stat 106 1
(integer) 0
127.0.0.1:6379> BITCOUNT login_stat 0 -1 BIT
(integer) 5
# 退出状态修改为 0
127.0.0.1:6379> SETBIT login_stat 105 0
(integer) 1
127.0.0.1:6379> BITCOUNT login_stat 0 -1 BIT
(integer) 4
```

##### 用户每个月签到情况

统计签到时, 每个用户每天的签到用 1 个 bit 位表示, 将用户 id 和月份组合作为 key, 日期作为偏移量, 签到记作 1, 未签到记作 0

key 的格式: sign:{uid}:{yyyyMM}
offset 的格式: {dd}

> SETBIT sign:{uid}:{yyyyMM} {dd} 1|0

```shell
# 记录指定用户每天的签到情况
127.0.0.1:6379> SETBIT sign:101:202211 1 1
(integer) 0
127.0.0.1:6379> SETBIT sign:101:202211 2 1
(integer) 0
# 获取指定用户在某天是否签到
127.0.0.1:6379> GETBIT sign:101:202211 1
(integer) 0
# 记录指定用户每天的签到情况
127.0.0.1:6379> SETBIT sign:102:202211 1 0
(integer) 0
127.0.0.1:6379> SETBIT sign:102:202211 2 0
(integer) 0
# 获取指定用户在某天是否签到
127.0.0.1:6379> GETBIT sign:102:202211 2
(integer) 0
# 统计某个用户某个月的签到情况
127.0.0.1:6379> BITCOUNT sign:102:202211 0 -1 BIT
(integer) 30
```

##### 统计连续 7 天签到

统计连续签到时, 将日期作为 key, 总共有 7 个 BitMaps, 用户 uid 作为偏移量, 签到记作 1, 未签到记作 0
用 `BITOP` 对 7 个 BitMaps 做 `AND` 运算, 结果中包含连续签到的位

key 的格式: bitmap:{dd}
offset 的格式: {uid}

> SETBIT bitmap:{dd} {uid} 1|0

```shell
# 记录每天每个用户的签到情况
127.0.0.1:6379> SETBIT bitmap:01 101 1
(integer) 0
127.0.0.1:6379> SETBIT bitmap:01 102 0
(integer) 0
127.0.0.1:6379> SETBIT bitmap:01 103 1
(integer) 0
127.0.0.1:6379> SETBIT bitmap:02 101 1
(integer) 0
127.0.0.1:6379> SETBIT bitmap:02 102 1
(integer) 0
127.0.0.1:6379> SETBIT bitmap:02 103 0
(integer) 0
127.0.0.1:6379> SETBIT bitmap:03 101 1
(integer) 0
127.0.0.1:6379> SETBIT bitmap:03 102 1
(integer) 0
127.0.0.1:6379> SETBIT bitmap:03 103 1
(integer) 0
# BITOP AND 位运算
127.0.0.1:6379> BITOP AND dest bitmap:01 bitmap:02 bitmap:03
(integer) 13
# 获取第一个值为 1 的偏移量
127.0.0.1:6379> BITPOS dest 1 0 -1 BIT
(integer) 101
```

### Bitfields

Redis Bitfields 是一种可以自定义设置、递增和获取任意位长度的整数值, 这些值使用二进制编码的 Redis 字符串存储, 位字段支持原子读取、写入和增量操作, 使其成为管理计数器和类似数值的好选择.

- BITFIELD key <GET encoding offset | [OVERFLOW <WRAP | SAT | FAIL>] <SET encoding offset value | INCRBY encoding offset increment> [GET encoding offset | [OVERFLOW <WRAP | SAT | FAIL>] <SET encoding offset value | INCRBY encoding offset increment> ...]> 自定义存储位数

```shell
127.0.0.1:6379> BITFIELD mykey incrby u2 100 1 OVERFLOW SAT incrby u2 102 1
1) (integer) 1
2) (integer) 1
127.0.0.1:6379> BITFIELD mykey incrby u2 100 1 OVERFLOW SAT incrby u2 102 1
1) (integer) 2
2) (integer) 2
127.0.0.1:6379> BITFIELD mykey incrby u2 100 1 OVERFLOW SAT incrby u2 102 1
1) (integer) 3
2) (integer) 3
127.0.0.1:6379> BITFIELD mykey incrby u2 100 1 OVERFLOW SAT incrby u2 102 1
1) (integer) 0
2) (integer) 3
127.0.0.1:6379> BITFIELD mykey OVERFLOW FAIL incrby u2 102 1
1) (nil)
```
