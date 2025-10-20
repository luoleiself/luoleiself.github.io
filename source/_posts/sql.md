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

## 索引

- 数据库管理系统里面一个经过排序的数据结构, 为了帮助查询变得更快
- 如果定义了主键索引则采用, 没有则使用第一个唯一索引加非空约束作为主键索引, 都没有的话定义一个新的隐藏字段包含 rowid 作为主键索引
- 聚簇索引决定了表中数据行的实际物理存储顺序

- 覆盖索引, 如果查询所需的全部列都在同一个非聚簇索引中(即所谓的覆盖索引), 则不需要再回表去读取完整的数据行.
- 联合索引, mysql 可以有效地利用最左边前缀原则来匹配查询条件, 对于联合索引 (col1, col2), 如果查询只涉及 col1, 仍然可以使用这个索引来加速搜索, 如果查询涉及到 col2 却没有 col1, 则无法有效利用此索引.

```sql
-- 添加索引
create [unique|primary|fulltext] index index_name on t1(filed);

-- 添加索引
alter table t1 add [unique|primary|fulltext] index index_name(filed);

-- 添加索引
create table table_name (
  id int,
  name varchar(255),
  index idx_name(name)
);

-- 删除索引
drop index index_name on table_name;
```

### 聚簇索引(主键索引)

唯一约束的基础上增加非空约束, 只能创建一个

### 唯一索引

同一个表中某个列的数据不能重复, 可以为 null, 可以创建多个唯一索引

```sql
create table table_name (
  id int,
  name varchar(255),
  email varchar(255) unique
)
```

### 普通索引

非唯一索引或简单索引, 主要用于加速对表中的数据的查询操作, 而不对列中的值增加任何唯一约束

### 联合索引

基于多列创建的索引, 能够提高涉及多列查询的性能

```sql
create table table_name (
  order_id int,
  customer_id int,
  order_date date,
  index idx_order_customer (order_id, customer_id)
)
```

### 全文索引

专门用于全文搜索, 适用于大文本(TEXT, CHAR, VARCHAR), 仅支持 myisam 和 innodb 存储引擎

```sql
create table table_name (
  id int,
  name varchar(255),
  email varchar(255) unique,
  title varchar(200)
  body text,
  fulltext (title, body)
)
```

## 多表查询

### 连接查询

#### 内连接

- 相当于查询 A、B 交集部分数据

```sql
-- 显式内连接
select * from t1 inner join t2 on condition;
-- 隐式内连接
select * from t1, t2 where condition;
```

#### 外连接

##### 左外连接

- 查询**左表**所有数据, 以及两张表交集部分数据
- 右表中不满足的数据填充 null

```sql
select * from t1 left join t2 on condition;
```

##### 右外连接

- 查询**右表**所有数据, 以及两张表交集部分数据
- 左表中不满足的数据填充 null

```sql
select * from t1 right join t2 on condition;
```

#### 自连接

- 当前表与自身的连接查询, 自连接必须使用表别名
- 可以是内连接查询, 也可以是外连接查询

```sql
 select * from t1 as a [inner|left|right] join t2 as b on condition;
```

#### 联合查询

- 查询的字段类型必须保持一致
- 把多次查询的结果合并起来, 形成一个新的结果集
- 默认已去重(distinct), union all 不去重

```sql
select * from t1 where condition union [all] select * from t2 where condition;
```

### 子查询

- sql 查询语句内嵌套查询语句, 称为嵌套语句

#### 标量子查询(单值子查询)

- 子查询的结果为单个值
- =, <>, >, >=, <, <= 常用操作符

```sql
select * from t1 where id = (select id from t1 where condition);
```

#### 列子查询

- 子查询的结果为一列
- in, not in, any, some, all 常用操作符

```sql
select * from t1 where id in (select id from t1 where condition);

-- 查询销售部和市场部的所有员工信息
select * from emp where dept_id in (select id from dept where name = '销售部' or name = '市场部');

-- 查询比财务部所有人工资都高的员工信息
select * from emp where salary > all (select salary from emp where dept_id = (select id from dept where name = '财务部'));

-- 查询比研发部其中任意一人工资高的员工信息
select * from emp where salary > any (select salary from emp where dept_id = (select id from dept where name = '研发部')); 
```

#### 行子查询

- 子查询的结果为一行(可以是多列)
- =, <>, in, not in 常用操作符

```sql
-- 查询与张三的薪资及直属领导相同的员工信息
select * from emp where (salary,managerid) = (select salary,managerid from emp where name = '张三');

-- 查询与张三不在同一部门的员工信息
select * from emp where dept_id not in (select dept_id from emp where name = '张三');
```

#### 表子查询

- 子查询的结果为多行多列
- in 常用操作符

```sql
-- 查询所有的部门信息, 并统计部门的员工人数
select d.id,d.name,(select count(*) from emp as e where e.dept_id = d.id) as '人数' from dept as d;

-- 查询与张三, 李四的职位和薪资相同的员工信息
select * from emp where (salary,jobs) in (select salary,jobs from emp where name = '张三' or name = '李四');

-- 查询入职日期是 2006-01-01 之后的员工信息, 及其部门信息
select e.*,d.* from (select * from emp where entrydate > '2006-01-01') as e left join dept as d on e.dept_id = d.id;
```

### 聚合查询

- 单一聚合
- 聚合管道
- mapReduce  MongoDB 5.0 开始废弃, 使用聚合管道代替

通常使用 group by 子句和聚合函数

- sum()
- avg()
- max()
- min()
- count()

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

```ts
db.createUser({
  user: "<name>", 
  pwd: '', 
  customData: {},
  roles: [
    {role: "dbAdmin", db: 'test'},
    {role: 'dbOwner', db: 'admin'}
  ]
}, writeConcernConf);
```

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
- db.serverStatus() 查看数据库状态

- db.auth() 用户认证

#### 备份和恢复

- -h, \-\-host  连接的数据库(端口号)
- \-\-port 数据库端口号
- -u, \-\-username  使用用户名认证
- -p, \-\-password  使用密码认证
- -d, \-\-db  选择数据库
- -c，\-\-collection  选择集合
- \-\-uri   数据库连接字符串格式

mongodump

- -o, \-\-out  存储目录, 默认使用 dump
- \-\-gzip  使用 gzip 压缩归档文件
- \-\-archive  设置归档文件名
- \-\-excludeCollection  排除指定的集合
- \-\-excludeCollectionsWithPrefix  排除指定前缀的集合
- -j, \-\-numParallelCollections  指定并行执行备份的集合数量

mongorestore

- \-\-archive  使用归档文件
- \-\-gzip  使用 gzip 解压归档文件
- \-\-dir 恢复备份文件的目录
- \-\-drop  恢复备份之前是否删除集合数据
- \-\-noIndexRestore  不恢复索引
- \-\-noOptionsRestore  不恢复集合的配置项
- \-\-stopOnError   插入文档失败时停止, 默认情况下, mongorestore 尝试文档验证和重复 key 的错误
- \-\-preserveUUID  保持原集合中的 UUIDs, 默认为 false

mongoexport

- -f, \-\-fields=\<fields\>[,\<fields\>]*   导出的数据字段
- \-\-fieldFile  导出的数据字段的文件, 每个字段一行
- \-\-type  指定导出文件的格式, json(默认) | csv | tsv
- -o, \-\-out  导出的文件名
- \-\-jsonArray  导出一个 JSON 数组每个对象一行
- \-\-pretty  导出一个人类阅读友好的 JSON 格式
- \-\-noHeaderLine  导出 CSV 文件不使用字段名作为首行
- -q, \-\-query  查询条件, JSON 字符串格式
- \-\-queryFile  查询条件的文件
- \-\-skip  指定导出文档的起始位置
- \-\-limit  指定导出文档的数量
- \-\-sort  排序条件, JSON 字符串格式
- \-\-assertExists  如果集合不存在则导出失败  

mongoimport

- -f, \-\-fields=\<fields\>[,\<fields\>]*   导出的数据字段
- \-\-fieldFile  导入的数据字段的文件, 每个字段一行
- \-\-file  导入的数据文件
- \-\-jsonArray  将导入文件作为 JSON 数组
- \-\-type  导入文件的格式, json(默认) | csv | tsv
- \-\-headerline  使用第一行作为列字段, 仅 CSV, TSV 支持
- \-\-ignoreBlanks   忽略空行, 仅 CSV, TSV 支持
- \-\-drop  插入文档之前是否删除集合数据
- \-\-mode  导入模式, insert(默认)|upsert|merge|delete,

```bash
# 备份
mongodump -h 127.0.0.1:27017 -d test -o ~/test
# 恢复备份
mongorestore -h 127.0.0.1:27017 -d test --dir ~/test --drop
# 导出数据
mongoexport -h 127.0.0.1:27017 -d test -o ~/test.json --type json --jsonArray --query '{age: {$gt: 18}}' --limit 100
# 导入数据
mongoimport -h 127.0.0.1:27017 -d test --file ~/test.json --type json --drop --mode upsert
```

### collection

- db.[collectionName].stats() 返回当前集合的使用统计信息
- db.[collectionName].getName() 返回当前集合的名称
- db.[collectionName].getDB() 返回当前集合所在的数据库名称
- db.[collectionName].drop() 删除当前的集合
- db.[collectionName].renameCollection() 重命名当前集合
- db.[collectionName].createIndex() 在当前集合中创建索引
- db.[collectionName].getIndexes() 获取当前集合的索引

- db.[collectionName].countDocuments() 获取当前集合中的文档的数量

#### indexs

一个经过排序的特殊的数据结构, 提高查询效率

- getIndexs() 查看索引
- createIndex() 创建索引, 索引名称必须是唯一的, 不能重命名已有的索引名称

```ts
db.<collection>.createIndex({ <field>: <value> }, { name: "<indexName>", unique: true });
```

- dropIndex() 删除指定索引
- dropIndexes() 批量删除索引, 不指定索引名则删除 _id 索引之外的所有索引

### document

- ObjectId  12个字节组成的对象标识符, 4 个字节的时间戳秒数, 5 个字节的进程随机数, 3 个字节的递增计数器
  - hexadecimal 可选。新对象标识符的 24 个字符十六进制字符串值。
  - integer 可选。整数值（以秒为单位）被添加到 Unix 纪元 以创建新的时间戳。
  - ObjectId().getTimestamp() 以日期形式返回对象的时间戳部分
  - ObjectId().toString() 以十六进制字符串形式返回对象标识符
  - ObjectId.createFromBase64(\<base64String\> [, \<subType\>]) 根据 base64 值创建对象标识符
  - ObjectId.createFromHexString(\<hexadecimalString\>) 根据十六进制值创建对象标识符
- String  字符串, utf-8 编码
- Integer 数值
- Double  双精度类型, 用于存储浮点数
- Boolean 布尔类型
- Arrays  数组类型
- TimeStamp 时间戳, 用于记录文档的创建和更新时间, 默认 ISODate()
- Date  日期类型, 以 UNIX 时间格式存储当前日期或时间
- Object  对象
- Null  空值
- Binary Data 二进制数据
- Symbol  符号
- Code  js 代码
- Regular Expression  正则表达式

### operator

#### 查询

比较查询

- $eq 匹配等于指定值的值
- $gt 匹配大于指定值的值
- $gte 匹配大于等于指定值的值
- $lt 匹配小于指定值的值
- $lte 匹配小于等于指定值的值
- $ne 匹配所有不等于指定值的值
- $nin 不匹配数组中指定值的任何值
- $in 匹配字段值等于指定数组中任何值的文档, 支持使用`正则表达式`匹配, 与 聚合操作符 $in 不同

```ts
// { field: { $in: [<value1>, <value2>, ... <valueN> ] } } // 查询操作符 $in
db.inventory.find( { quantity: { $in: [ 5, 15 ] } }, { _id: 0 } );
// output: 查询 inventory 集合中 quantity 字段的值为 5 或 15 的所有文档
```

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
// { name: { $regex: /acme.*corp/i, $nin: [ 'acmeblahcorp' ] } }

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

```ts
db.users.updateOne({name: 'laoli'}， {$inc: {age: 10}}); // 将名字为 laoli 的文档的 age 增加 10

db.users.updateMany({city: {$in: ['beijing']}}, {$set: {city: 'shanghai'}}); // 将城市为 beijing 的文档的 city 设置为 shanghai

db.users.updateMany({name: 'laoli'}， {$rename: ['lve', 'love']}); // 将名字为 laoli 的文档的 lve 重命名为 love
```

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

```ts
db.users.updateOne({name: 'laoli'}, {$push: {lve: 'TV'}}); // 将名字为 laoli 的文档的 lve 数组添加 TV

db.users.updateMany({age: {$gte: 35}}, {$addToSet: {lve: {$each: ['code', 'football']}}}); // 将年龄 大于等于 35 岁的文档的 lve 增加 code 和 football
```

### 批量写入

通过更少的数据库调用来执行多个写入操作

```ts
[
  [ 'deleteMany' => [ $filter ] ],
  [ 'deleteOne'  => [ $filter ] ],
  [ 'insertOne'  => [ $document ] ],
  [ 'replaceOne' => [ $filter, $replacement, $options ] ],
  [ 'updateMany' => [ $filter, $update, $options ] ],
  [ 'updateOne'  => [ $filter, $update, $options ] ],
]
```

## 副本集

## 分片集
