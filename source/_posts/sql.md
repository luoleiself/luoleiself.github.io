---
title: sql
date: 2022-11-09 18:03:19
categories:
  - [server, sql]
tags:
  - sql
---

## 事务隔离级别

mysql 事务隔离级别定义了一个事务在多大程度上能够看到其他并发事务所做的修改

```sql
SET [GLOBAL|SESSION] TRANSACTION ISOLATION LEVEL <level>;

-- 全局设置
SET GLOBAL TRANSACTION ISOLATION LEVEL REPEATABLE READ;

-- 当前会话设置
SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;

-- 事务开始时设置
SET TRANSACTION ISOLATION LEVEL <level>;
```

### 读未提交(Read Uncommitted)

最低级别的隔离, 允许一个事务读取另一个事务尚未提交的数据, 通常不推荐使用此级别, 除非对数据一致性要求非常低的场景

问题:

- 脏读(Dirty Read), 一个事务可以读取到另一个事务未提交的数据, 如果该事务回滚, 则会导致数据不一致
- 不可重复读(Non-repeatable Read), 在一个事务中多次读取同一数据可能会得到不同的结果
- 幻读(Phantom Read), 在一个事务中执行相同的查询可能会返回不同的行集, 因为其他事务插入了新的行

```sql
-- 事务 A
START TRANSACTION;
UPDATE accounts SET balance = balance - 100 WHERE id = 1;
-- 暂时不提交
COMMIT;

-- 事务 B
START TRANSACTION;
SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
SELECT balance FROM accounts WHERE id = 1;
-- 可能读取到 900 (脏读)
COMMIT;
```

<!--more-->

### 读已提交(Read Committed)

一个事务只能读取另一个事务已经提交的数据, 避免的 `脏读` 问题, 适用于大多数需要避免 `脏读` 的应用场景

问题:

- 不可重复读(Non-repeatable Read), 在一个事务中多次读取同一数据可能会得到不同的结果, 因为其他事务可能在这期间更新并提交了数据
- 幻读(Phantom Read), 在一个事务中执行相同的查询可能会返回不同的行集, 因为其他事务插入了新的行

```sql
-- 事务 A
START TRANSACTION;
UPDATE accounts SET balance = balance - 100 WHERE id = 1;
COMMIT;

-- 事务 B
START TRANSACTION;
SET TRANSACTION ISOLATION LEVEL READ COMMITTED;
SELECT balance FROM accounts WHERE id = 1;
-- 读取到 900 (避免脏读)
COMMIT;
```

### 可重复读(Repeatable Read)

mysql 默认隔离级别, 它确保在一个事务中多次读取同一数据时, 结果是一致的, 即使其他事务在此期间修改并提交了数据

避免了 `脏读` 和 `不可重复读` 问题,

实现机制: InnoDB 存储引擎通过多版本并发控制(MVVC) 来实现这一隔离级别, 并通过锁定机制防止幻读

问题:

- 幻读(Phantom Read), 在一个事务中执行相同的查询可能会返回不同的行集, 因为其他事务插入了新的行

```sql
-- 事务 A
START TRANSACTION;
SELECT balance FROM accounts WHERE id = 1; -- 读取到 1000
UPDATE accounts SET balance = balance - 100 WHERE id = 1;
COMMIT;

-- 事务 B
START TRANSACTION;
SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;
SELECT balance FROM accounts WHERE id = 1; -- 读取到 1000（避免不可重复读）
COMMIT;
```

### 串行化(Serializable)

最高隔离级别, 完全避免了 `脏读`, `不可重复读` 和 `幻读` 问题, 它通过强制事务串行执行来实现这一点

适用于对数据一致性要求极高且并发性不是主要问题的应用场景

问题:

- 性能开销较大, 因为它会极大地限制并发性

```sql
-- 事务 A
START TRANSACTION;
SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
SELECT * FROM accounts WHERE balance > 500 FOR UPDATE; -- 锁定符合条件的行
-- 执行其他操作
COMMIT;

-- 事务 B
START TRANSACTION;
SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
SELECT * FROM accounts WHERE balance > 500 FOR UPDATE; -- 等待 事务 A 完成
-- 执行其他操作
COMMIT;
```

## mongosh

```mongosh
version()  # 显示版本号
.exit    # 退出
quit     # 退出

show logs # 显示日志
show users
show roles

sleep() 程序睡眠时长, 默认单位毫秒
load()  # 加载指定 js 文件代码到当前的终端下
passwordPrompt() # 输入指定用户的密码
print() # 输出指定内容

it # 查看更多结果

cls # 清屏, 类似于 console.log()
```

### 用户

- db.createUser() 创建用户并指定用户的数据库或者集合访问权限

### database

- db.help() 返回当前数据库对象上常用方法的文本信息
- db.hello() 返回当前数据库对象的基础信息
- db.stats() 返回当前数据库对象的使用统计信息
- db.version() 返回当前 mongod 或 mongos 实例的版本号
- db.getName() 返回当前数据库对象的名称

- db.dropDatabase() 删除当前数据库和关联的数据文件
- db.shutdownServer() 安全的关闭当前 mongod 或 mongos 进程

- db.getSiblingDB() 返回当前数据库同级指定名称的数据库

- db.createCollection() 在当前数据库对象上创建集合, 第二个可选参数可以配置 固定集 的大小
- db.getCollection() 获取指定名称的集合

### collection

- db.[collectionName].stats() 返回当前集合的使用统计信息
- db.[collectionName].getName() 返回当前集合的名称
- db.[collectionName].getDB() 返回当前集合所在的数据库名称
- db.[collectionName].drop() 删除当前的集合
- db.[collectionName].renameCollection() 重命名当前集合
- db.[collectionName].createIndex() 在当前集合中创建索引
- db.[collectionName].getIndexes() 获取当前集合的索引

- db.[collectionName].countDocuments() 获取当前集合中的文档的数量

### document

- ObjectId, 12个字节的对象标识符, 4 个字节的时间戳秒数, 5 个字节的进程随机数, 3 个字节的递增计数器
- String, 字符串, utf-8 编码
- Integer, 数值
- Double, 双精度类型, 用于存储浮点数
- Boolean, 布尔类型
- Arrays, 数组类型
- TimeStamp, 时间戳, 用于记录文档的创建和更新时间, 默认 ISODate()
- Date, 日期类型, 以 UNIX 时间格式存储当前日期或时间
- Object, 对象
- Null, 空值
- Binary Data, 二进制数据
- Symbol, 符号
- Code, js 代码
- Regular Expression, 正则表达式

### operator

#### 查询

比较查询

- $eq 匹配等于指定值的值
- $gt 匹配大于指定值的值
- $gte 匹配大于等于指定值的值
- $in 匹配数组中指定的任何值
- $lt 匹配小于指定值的值
- $lte 匹配小于等于指定值的值
- $ne 匹配所有不等于指定值的值
- $nin 不匹配数组中指定值的任何值

逻辑查询

- $and 使用逻辑 AND 连接查询子句将返回与两个子句的条件匹配的所有文档
- $not 反转查询谓词的效果, 并返回与查询谓词 不匹配 的文档
- $nor 使用逻辑 NOR 连接查询子句会返回无法匹配这两个子句的所有文档
- $or 使用逻辑 OR 连接多个查询子句会返回符合任一子句条件的所有文档
- $exists 匹配具有指定字段的文档
- $type 如果字段为指定类型, 则选择文档

```javascript
// 查询 price 不等于 1.99 并且 qty 不小于 20 并且 sale 不等于 true 的所有文档
db.inventory.find({$nor: [{price: 1.99}, {qty: {$lt: 20}}, {sale: true}]});
// 查询 quantity 小于 20 或者 price 等于 10 的所有文档
db.inventory.find({$or: [{quantity: {$gt: 20} }, {price: 10}]});

// 查询 qty 字段存在并且值 不等于 5 或 15 的所有文档
db.inventory.find({qty: {$exists: true, $nin: [5, 15]}});
// 查询 readings 字段为数组(空或非空)的所有文档
db.sensorReading.find({readings: {$type: 'array'}});
```

评估查询

- $expr 允许在查询语言中使用聚合表达式
- $jsonSchema 根据给定的 JSON 模式验证文档
- $mod  对字段值执行模运算, 并选择具有指定结果的文档, 当传递的数组元素不是两个时将会报错
- $regex  选择值匹配指定正则表达式的文档
- $text  执行文档搜索
- $where  匹配满足 javascript 表达式的文档

```javascript
// 查询 spent 金额超过 budget 的所有文档
db.monthlyBudget.find({$expr: { $gt: [ "$spent" , "$budget" ] }});
// 查询 qty 字段的值 模 4 等于 0 的所有文档
db.inventory.find({qty: { $mod: [ 4, 0 ] }});
// 查询 name 的值等于 MD5 哈希值的所有文档
db.players.find({$where: function() {
  return (hex_md5(this.name) == "9b53e667f30cd329dca1ec9e6a83e994")
}});
```

地理空间

- $geoIntersects 选择与 GeoJSON 几何图形相交的几何图形
- $geoWithin  选择在边界 GeoJSON 几何图形内的几何图形
- $near 返回接近某个点的地理空间对象, 需要地理空间索引
- $nearSphere  返回与球面上的某个点接近的地理空间对象, 需要地理空间索引

数组查询

- $all  匹配包含查询中指定的所有元素的数组
- $elemMatch 如果数组字段中的元素与所有指定的 $elemMatch 条件均匹配, 则选择文档
- $size 如果数组字段达到指定大小, 则选择文档

```javascript
// 查询 tags 字段值至少包含 appliance, school, book 的文档
db.inventory.find({tags: { $all: [ "appliance", "school", "book" ]}});
// 查询 results 数组中至少包含一个 product 等于 xyz 且 score 大于或等于 8 的元素的文档
db.survey.find({results: { $elemMatch: { product: "xyz", score: { $gte: 8 } }}});
```

按位查询

- $bitsAllClear 匹配数字或二进制值，其中一组片段位置均包含值 0
- $bitsAllSet 匹配数字或二进制值，其中一组片段位置均包含值 1
- $bitsAnyClear 匹配数字或二进制值，其中一组位位置中的任何 位的值为 0
- $bitsAnySet 匹配数字或二进制值，其中一组位位置中的任何 位的值为 1

投影操作符

- $ 对数组中与查询条件匹配的第一个元素进行投影
- $elemMatch  对数组中与指定 $elemMatch 条件匹配的第一个元素进行投影
- $meta 预测在 $text 操作中分配的文件分数
- $slice  限制从数组中投影的元素数量。支持跳过切片和对切片进行数量限制

其他操作符

- $rand 生成介于 0 和 1 之间的随机浮点数
- $natural  可通过 sort() 或 hint() 方法提供的特殊提示，可用于强制执行正向或反向集合扫描

#### update

字段

- $currentDate  将字段的值设置为当前日期，可以是日期或时间戳。
- $inc  将字段的值按指定量递增。
- $min  仅当指定值小于现有字段值时才更新字段。
- $max  仅当指定值大于现有字段值时才更新字段。
- $mul  将字段的值乘以指定量。
- $rename 重命名字段。
- $set  设置文档中字段的值。
- $setOnInsert  如果某一更新操作导致插入文档，则设置字段的值。对修改现有文档的更新操作没有影响。
- $unset  从文档中删除指定的字段。

其他操作符

- $bit 对整数值执行按位 AND、OR、XOR 更新

##### 数组

操作符

- $ 充当占位符，用于更新与查询条件匹配的第一个元素
- $[] 充当占位符，以更新数组中与查询条件匹配的文档中的所有元素
- $[\<identifier\>] 充当占位符，以更新与查询条件匹配的文档中所有符合 arrayFilters 条件的元素
- $addToSet 仅向数组中添加尚不存在于该数组的元素, 配合 $each 批量添加元素
- $pop  删除数组的第一项或最后一项，-1 第一个，1 最后一个
- $pull 删除与指定查询匹配的所有数组元素
- $push 向数组添加一项, 配合 $each 批量添加元素
- $pullAll  从数组中删除所有匹配值

修饰符

- $each 修改 $push 和 $addToSet 运算符，以在数组更新时追加多个项目
- $position 修改 $push 运算符，以指定在数组中添加元素的位置
- $slice  修改 $push 运算符以限制更新后数组的大小, 负数表示从数组尾部开始截取, 正数表示从数组头部开始截取
- $sort 修改 $push 运算符，以对存储在数组中的文档重新排序

### 聚合操作

处理多个文档并返回计算结果

#### 聚合管道

由 一个或多个处理文档的 阶段 组成

- 每个阶段对输入文档执行一个操作
- 从一个阶段输出的文档将传递到下一个阶段
- 一个聚合管道可以返回针对文档组的结果

`db.[collectionName].aggregate()` 方法运行的聚合管道不会修改集合中的文档, 除非管道包含 `$merge` 或 `$out` 阶段

```javascript
db.orders.aggregate([
  // stage 1: Filter pizza order document by pizza size
  {
    $match: {size: 'medium'}
  },
  // stage 2: Group remaining document by pizza name and calcuate total quantity
  {
    $group: {_id: '$name', totalQuantity: {$sum: "$quantity"}}
  },
  // stage 3: Sort document by totalQuantity in descending order
  {
    $sort: {totalQuantity: -1}
  }
])
```

阶段

- $addFields 为文档添加新字段
- $bucket 根据指定的表达式和存储桶边界将传入的文档分为多个组
- $collStats  返回有关集合或视图的统计信息
- $count 计数
- $documents  从输入表达式返回字面文档
- $fill 填充文档中的 null 和缺失的字段值
- $geoNear  根据与地理空间点的接近程度返回有序的文档流
- $group  按指定的标识符表达式对输入文档进行分组，并将累加器表达式（如果指定）应用于每个群组
- $limit  将未修改的前 n 个文档传递到管道，其中 n 为指定的限制
- $match  筛选文档流以仅允许匹配的文档将未修改的文档传入下一管道阶段
- $merge  将 aggregation pipeline 的结果文档写入集合
- $out  将 aggregation pipeline 的结果文档写入集合
- $project  重塑流中的每个文档，例如添加新的字段或删除现有字段
- $set  为文档添加新字段, 与 $project 类似
- $skip 跳过前 n 个文档，其中 n 是指定的跳过编号，并将未修改的剩余文档传递到管道
- $sort 按指定的排序键对文档流重新排序。仅顺序会改变，而文档则保持不变
- $sortByCount  根据指定表达式的值对传入文档进行分组，然后计算每个不同群组中的文档数量
- $unset  从文档中删除/排除字段, $unset 是删除字段的 $project 阶段的别名

#### 单一目的的聚合方法

- db.[collectionName].estimatedDocumentCount()  返回集合或视图中文档的近似数量
- db.[collectionName].count() 返回集合或视图中文档的数量
- db.[collectionName].distinct()  返回具有指定字段的不同值的文档数组

```javascript
// 统计 ord_dt 晚于指定日期的文档数量
db.orders.count( { ord_dt: { $gt: new Date('01/01/2012') } } );
// 功能同上, 使用索引返回计数
db.orders.find( { ord_dt: { $gt: new Date('01/01/2012') } } ).count();

// 返回所有文档中字段 dept 不同的值
db.inventory.distinct( "dept" );
```
