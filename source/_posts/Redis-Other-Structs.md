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

#### 规定

EPSG:900913 / EPSG:3785 / OSGEO:41001 标准规定

- 有效经度在 -180 度到 180 度
- 有效维度在 -85.05112878 度到 85.05112878 度

#### 成员操作

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
```

- GEOSEARCH key <FROMMEMBER member|FROMLONLAT longitude latitude&> <BYRADIUS radius <M|KM|FT|MI>|BYBOX width height <M|KM|FT|MI>> [ASC|DESC] [COUNT count [ANY]] [WITHCOORD] [WITHDIST] [WITHHASH] 返回集合中符合属于给定形状区域的边界内的成员的信息, 默认返回的匹配项没有排序, 使用 COUNT 限制返回匹配项的数量

  - FROMMEMBER 使用给定的已存在的成员进行搜索
  - FROMLONLAT 使用给定的 longitude 和 latitude 搜索
  - BYRADIUS 行为类似于 `GEORADIUS`, 根据给定半径的圆形区域内搜索
  - BYBOX 根据给定的 width 和 height 在轴对齐的矩形区域内搜索
  - ASC 相对于中心点, 从近到远对返回的成员进行排序
  - DESC 相对于中心点, 从远到近对返回的成员的进行排序
  - COUNT 限制匹配项的数量
  - ANY 找到足够的匹配项就立刻返回, 有可能匹配项不是最接近指定点的位置, 但是服务器生成这些结果所投入的精力要少的多
    未指定此参数时, 搜索命令将执行与指定匹配区域匹配项数量成正比的操作, 并对其进行排序, 因此, 即使返回少量结果, 使用非常小的 count 查询非常大的区域也可能会很慢.
  - WITHCOORD 返回匹配项的经纬度
  - WITHDIST 返回匹配项距离指定中心点的距离, 距离单位和指定圆形区域搜索或者矩形区域搜索的单位相同
  - WITHHASH 返回以 52 位无符号整数的形式表示匹配项 `GEOHASH` 编码

```shell

```
