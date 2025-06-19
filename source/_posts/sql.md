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

- ObjectId  对象标识符, 4 个字节的时间戳秒数, 5 个字节的进程随机数, 3 个字节的递增计数器
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
- $in 匹配字段值等于指定数组中任何值的文档, 支持使用`正则表达式`匹配, 与 [聚合操作符 $in](#in-2) 不同 <em id="in-1"></em> <!--markdownlint-disable-line-->

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

### 聚合操作

操作处理多个文档并返回计算结果, 由 一个或多个处理文档的 `阶段` 组成

- 每个阶段对输入文档执行一个操作
- 从一个阶段输出的文档将传递到下一个阶段
- 一个聚合管道可以返回针对文档组的结果

字段路径, 使用 `$` 前缀启用字段路径表达式访问输入文档中的字段.

管道优化, 该阶段会尝试重塑管道以提高性能, 优先使用 $match、$sort、$limit、$skip 阶段限制进入管道的文档

限制

- 结果大小限制, 每个文档均受 MB16 BSON 文档大小的限制的约束
- 阶段数量限制, 单个管道中允许的聚合管道阶段数量最大为 1000 个
- 内存限制, 6.0 开始 allowDiskUseByDefault 参数控制需要 100MB 以上内存容量来执行的管道阶段是否默认会将临时文件写入磁盘

`db.[collectionName].aggregate()` 方法运行的聚合管道不会修改集合中的文档, 除非管道包含 `$merge` 或 `$out` 阶段

```javascript
db.orders.aggregate([
  // stage 1: Filter pizza order document by pizza size
  {$match: {size: 'medium'}},
  // stage 2: Group remaining document by pizza name and calcuate total quantity
  {$group: {_id: '$name', totalQuantity: {$sum: "$quantity"}}},
  // stage 3: Sort document by totalQuantity in descending order
  {$sort: {totalQuantity: -1}}
])
```

#### 变量

变量可以保存任何 BSON 类型的数据, 访问变量时需要使用 `$$` 前缀.

##### 用户变量

变量名称可包含 ASCII 字符和任意非 ASCII 字符, 必须以小写 ASCII 字符开头

##### 系统变量

- NOW 返回当前日期时间值的变量, 为部署的所有成员返回相同的值, 并在聚合管道的所有阶段保持不变.
- CLUSTER_TIME 返回当前时间戳值的变量
  - 仅适用于副本集和分片的集群
  - 为部署的所有节点返回相同的值, 并在管道的所有阶段保持不变
- ROOT  引用根文档, 即当前正在聚合管道阶段处理的顶层文档
  - 引用聚合管道中当前正在处理的完整文档
  - 常用于需要保留原始文档内容的场景
  - 在 $group、$project 等阶段特别有用
- CURRENT 引用聚合管道阶段正在处理的 `字段路径` 的起始位置
- REMOVE  一个求值为缺失的变量, 允许排除 $addFields 和 $project 阶段的字段
- DESCEND $redact 表达式的允许结果之一, 返回当前文档级别的字段, 不包括嵌入式文档
- PRUNE $redact 表达式的允许结果之一, 排除当前文档/嵌入式文档级别的所有字段, 而不进一步检查任何已排除的字段
- KEEP  $redact 表达式的允许结果之一, 返回或保留此当前文档/嵌入式文档级别的所有字段, 而不进一步检查此级别的字段
- SEARCH_META search 查询元数据结果的变量, 在所有支持的聚合管道阶段中, 设立变量 $SEARCH_META 的字段会返回查询的元数据结果
- USER_ROLES  返回分配给当前用户的角色

#### 聚合命令

- aggregate 使用管道执行聚合任务
- count 计算集合或视图中的文档数量
- distinct 显示在集合或视图中为指定键找到的非重复值
- mapReduce 为大型数据集执行 map-reduce 聚合任务

单一目的的聚合方法

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

#### 阶段

文档按顺序通过聚合管道阶段

- $addFields 为文档添加新字段
- $bucket 根据指定的表达式和存储桶边界将传入的文档分为多个组, 并为每个组输出一个文档
  - groupBy 对文档进行分组的表达式
  - boundaries 边界,基于指定每个存储桶边界的 groupBy 表达式的值
  - default 指定附加存储桶 _id 的字面量
  - output 一份文档, 指定除 _id 字段之外要包含在输出文档中的字段

```ts
/*
// artists
{ "_id" : 1, "last_name" : "Bernard", "first_name" : "Emil", "year_born" : 1868, "year_died" : 1941, "nationality" : "France" },
{ "_id" : 2, "last_name" : "Rippl-Ronai", "first_name" : "Joszef", "year_born" : 1861, "year_died" : 1927, "nationality" : "Hungary" },
{ "_id" : 3, "last_name" : "Ostroumova", "first_name" : "Anna", "year_born" : 1871, "year_died" : 1955, "nationality" : "Russia" },
{ "_id" : 4, "last_name" : "Van Gogh", "first_name" : "Vincent", "year_born" : 1853, "year_died" : 1890, "nationality" : "Holland" },
{ "_id" : 5, "last_name" : "Maurer", "first_name" : "Alfred", "year_born" : 1868, "year_died" : 1932, "nationality" : "USA" },
{ "_id" : 6, "last_name" : "Munch", "first_name" : "Edvard", "year_born" : 1863, "year_died" : 1944, "nationality" : "Norway" },
{ "_id" : 7, "last_name" : "Redon", "first_name" : "Odilon", "year_born" : 1840, "year_died" : 1916, "nationality" : "France" },
{ "_id" : 8, "last_name" : "Diriks", "first_name" : "Edvard", "year_born" : 1855, "year_died" : 1930, "nationality" : "Norway" }
*/
db.artists.aggregate([
  { $bucket: {
      groupBy: '$year_born',  // Field to group by
      boundaries: [1840, 1858, 1860, 1870, 1880], // Boundaries for the buckets
      default: 'Other',   // Bucket ID for documents which do not fail into a bucket
      output: {
        'count': {$sum: 1},
        'artists': {
          $push: {
            'name': {$concat: ['$first_name', ' ', '$last_name']},
            'year_born': '$year_born'
          }
        }
      }
    }
  },
  {
    $match: {count: {$gt: 3}}
  }
]);
// output:
// { "_id" : 1860, "count" : 4, "artists" :
//   [
//     { "name" : "Emil Bernard", "year_born" : 1868 },
//     { "name" : "Joszef Rippl-Ronai", "year_born" : 1861 },
//     { "name" : "Alfred Maurer", "year_born" : 1868 },
//     { "name" : "Edvard Munch", "year_born" : 1863 }
//   ]
// }
```

- $collStats  返回有关集合或视图的统计信息, 必须是聚合管道中的 **第一阶段**
- $count 计数
- $documents  从输入表达式返回字面文档
  - 只能在数据库级聚合管道中使用
  - 必须是聚合管道中的 **第一阶段**
- $facet  在单个阶段内处理同一组输入文档上的多个聚合管道, 每个子管道在输出文档中都有自己的字段, 其结果存储为文档数组
- $fill 填充文档中的 null 和缺失的字段值
  - sortBy 指定每个分区内用于对文档进行排序的字段, 使用与 $sort 阶段相同的语法
  - output 指定一个对象, 其中包含要填充缺失值的每个字段
- $geoNear  根据与地理空间点的接近程度返回有序的文档流
- $group  按指定的标识符表达式对输入文档进行分组，并将累加器表达式（如果指定）应用于每个群组, 不会对其输出文档进行排序
  - _id 指定群组标识符表达式, 如果指定的 \_id 值为空值或任何其他常量值, $group 阶段将返回聚合所有输入文档值的单个文档
  - field 使用累加器操作符进行计算

```ts
[
  {
    $group: {
      _id: <expression>， // Group key
      <field1>: { <accumulator1>: <expression1> },
      // ...
    }
  }
]
```

- $indexStats 返回集合中每个索引使用情况的统计信息, 此阶段采用一个空文档, 不需要配置项
- $limit  将未修改的前 n 个文档传递到管道的下一阶段，其中 n 为指定的限制
- $lookup 对同一数据库中的集合执行`左外连接`, 以过滤外部集合中的文档进行处理
  - from 在同一数据库中指定待连接到本地集合的外部集合
  - localField 指定执行等值匹配时本地文档的字段, 如果输入文档不包含 localField 则被视为 null 值
  - foreignField 指定外部文档的 foreignField 对本地文档的 localField 执行等值匹配, 如果外部文档不包含 foreignField 值, 则使用 null 进行匹配
  - let 引用管道阶段中的变量
  - pipeline 不能包含 $merge 或 $out 阶段
  - as 指定要添加到输入文档中的新数组字段的名称, 新数组字段包含来自 from 集合的匹配文档, 如果指定名称已存在将被重写

```ts
/*
// orders
{ "_id" : 1, "item" : "almonds", "price" : 12, "quantity" : 2 }
{ "_id" : 2, "item" : "pecans", "price" : 20, "quantity" : 1 }
{ "_id" : 3  }
// inventory
{ "_id" : 1, "sku" : "almonds", "description": "product 1", "instock" : 120 }
{ "_id" : 2, "sku" : "bread", "description": "product 2", "instock" : 80 }
{ "_id" : 3, "sku" : "cashews", "description": "product 3", "instock" : 60 }
{ "_id" : 4, "sku" : "pecans", "description": "product 4", "instock" : 70 }
{ "_id" : 5, "sku": null, "description": "Incomplete" }
{ "_id" : 6 }
*/
db.orders.aggregate([
  {$lookup: {from: 'inventory', localField: 'item', foreignField: 'sku', as: 'inventory_docs'}}
]);
// output: localField 和 foreignField 不包含的则作为 null 值执行匹配
// {
//   "_id" : 1,
//   "item" : "almonds",
//   "price" : 12,
//   "quantity" : 2,
//   "inventory_docs" : [
//     { "_id" : 1, "sku" : "almonds", "description" : "product 1", "instock" : 120 }
//   ]
// }
// {
//   "_id" : 2,
//   "item" : "pecans",
//   "price" : 20,
//   "quantity" : 1,
//   "inventory_docs" : [
//     { "_id" : 4, "sku" : "pecans", "description" : "product 4", "instock" : 70 }
//   ]
// }
// {
//   "_id" : 3,
//   "inventory_docs" : [
//     { "_id" : 5, "sku" : null, "description" : "Incomplete" },
//     { "_id" : 6 }
//   ]
// }
```  

- $match  筛选文档流以仅允许匹配的文档将未修改的文档传入下一管道阶段
- $merge  将 aggregation pipeline 的结果文档写入集合, 必须是聚合管道中的 **最后一阶段**
- $out  将 aggregation pipeline 的结果文档写入集合, 必须是聚合管道中的 **最后一阶段**
- $planCacheStats 返回集合的计划缓存的信息, 必须是聚合管道中的 **第一阶段**
  - allHosts  配置聚合阶段如何以分片集群中的节点为目标
- $project  重塑流中的文档至管道中的下个阶段，指定的字段可以是文档中已有字段或新计算的字段. 例如添加新的字段或删除现有字段
- $redact 根据存储在文档本身中的信息, 限制整个文档被输出或者文档中的内容被输出

```ts
/*
// forecasts
{
  _id: 1,
  title: "123 Department Report",
  tags: [ "G", "STLW" ],
  year: 2014,
  subsections: [
    {
      subtitle: "Section 1: Overview",
      tags: [ "SI", "G" ],
      content:  "Section 1: This is the content of section 1."
    },
    {
      subtitle: "Section 2: Analysis",
      tags: [ "STLW" ],
      content: "Section 2: This is the content of section 2."
    },
    {
      subtitle: "Section 3: Budgeting",
      tags: [ "TK" ],
      content: {
        text: "Section 3: This is the content of section 3.",
        tags: [ "HCS" ]
      }
    }
  ]
}
*/
var userAccess = [ "STLW", "G" ];
db.forecasts.aggregate([
  { $match: {year: 2014} },
  { $redact: {
      $cond: {
        if: {$gt: [{$size: {$setIntersection: ['$tags', userAccess]}}, 0]},
        then: '$$DESCEND',
        else: '$$PRUNE'
      }
    }
  }
]);
// output: 查看带有标签 STLW 或 G 的信息
// {
//   "_id" : 1,
//   "title" : "123 Department Report",
//   "tags" : [ "G", "STLW" ],
//   "year" : 2014,
//   "subsections" : [
//     {
//       "subtitle" : "Section 1: Overview",
//       "tags" : [ "SI", "G" ],
//       "content" : "Section 1: This is the content of section 1."
//     },
//     {
//       "subtitle" : "Section 2: Analysis",
//       "tags" : [ "STLW" ],
//       "content" : "Section 2: This is the content of section 2."
//     }
//   ]
// }
```

- $replaceRoot 将输入文档替换为指定文档, 该操作会替换输入文档中的所有现有字段, 包括 _id 字段
- $replaceWith 将输入文档替换为指定文档, 该操作会替换输入文档中的所有现有字段, 包括 _id 字段

```ts
{$replaceRoot: { newRoot: <replacementDocument>}} // 如果 replacementDocument 不是文档或者解析为缺失的文档, 则报错并失败
{$replaceWith: <replacementDocument>} // 如果 replacementDocument 不是文档或者解析为缺失的文档, 则报错并失败

/*
// students
{
  "_id" : 1,
  "grades" : [
    { "test": 1, "grade" : 80, "mean" : 75, "std" : 6 },
    { "test": 2, "grade" : 85, "mean" : 90, "std" : 4 },
    { "test": 3, "grade" : 95, "mean" : 85, "std" : 6 }
  ]
},
{
  "_id" : 2,
  "grades" : [
    { "test": 1, "grade" : 90, "mean" : 75, "std" : 6 },
    { "test": 2, "grade" : 87, "mean" : 90, "std" : 3 },
    { "test": 3, "grade" : 91, "mean" : 85, "std" : 4 }
  ]
}
*/
db.students.aggregate([
  {$unwind: '$grades'},
  {$match: {'grades.grade': {$gte: 90}}},
  {$replaceRoot: {newRoot: '$grades'}}
]);
// output: 将 grade 字段的值大于等于 90 的嵌入式文档提升到顶层
{ "test" : 3, "grade" : 95, "mean" : 85, "std" : 6 }
{ "test" : 1, "grade" : 90, "mean" : 75, "std" : 6 }
{ "test" : 3, "grade" : 91, "mean" : 85, "std" : 4 }
```

- $sample 从输入文档中随机选择指定数量的文档
  - size 随机选择文档的数量
- $set  为文档添加新字段, 与 $project 类似
- $skip 跳过前 n 个文档，其中 n 是指定的跳过编号，并将未修改的剩余文档传递到管道的下一阶段
- $sort 按指定的排序键对文档流重新排序。仅顺序会改变，而文档则保持不变
- $sortByCount  根据指定表达式的值对传入文档进行分组，然后计算每个不同群组中的文档数量
- $unionWith 将两个集合合并为一个结果集
- $unset  从文档中删除/排除字段, $unset 是删除字段的 $project 阶段的别名
- $unwind 解构输入文档中的数组字段, 以便为每个元素输出文档, 并用该元素替换该数组字段的值, 忽略无法转换成 `单元素数组` 的文档

```ts
/*
// inventory
{ "_id" : 1, "item" : "ABC1", sizes: [ "S", "M", "L"] }
*/
db.inventory.aggregate([{$unwind: '$sizes'}]);
// output:
// { "_id" : 1, "item" : "ABC1", "sizes" : "S" }
// { "_id" : 1, "item" : "ABC1", "sizes" : "M" }
// { "_id" : 1, "item" : "ABC1", "sizes" : "L" }
```

#### 操作符

- $add 将数字或者日期相加, 如果其中一个参数是日期, 则其他参数被视为添加到日期的毫秒数
- $sum 计算并返回数字值的总和，忽略非数值
- $avg 返回数值的平均值, 忽略非数值
- $multiply 将数字相乘并返回结果
- $divide 返回第一个数字除以第二个数字的结果
- $mod 返回第一个数字除以第二个数字的余数

```ts
/*
// sales
{ "_id" : 1, "item" : "abc", "price" : 10, "fee": 2, "quantity" : 2, "date" : ISODate("2014-01-01T08:00:00Z") }
{ "_id" : 2, "item" : "jkl", "price" : 20, "fee": 1, "quantity" : 1, "date" : ISODate("2014-02-03T09:00:00Z") }
{ "_id" : 3, "item" : "xyz", "price" : 5, "fee": 0, "quantity" : 5, "date" : ISODate("2014-02-03T09:05:00Z") }
{ "_id" : 4, "item" : "abc", "price" : 10, "fee": 3, "quantity" : 10, "date" : ISODate("2014-02-15T08:00:00Z") }
{ "_id" : 5, "item" : "xyz", "price" : 5, "fee": 4, "quantity" : 10, "date" : ISODate("2014-02-15T09:12:00Z") }
*/
db.sales.aggregate([
  {$project: {item: 1, total: {$add: ['$price', '$fee']}, billing_date: {$add: ['$date', 3*24*60*60*1000]}}}
]);
// output: total = price + fee; billing_date = date + 3*24*60*60*1000
// { "_id" : 1, "item" : "abc", "total" : 12, "billing_date" : ISODate("2014-01-04T08:00:00Z") }
// { "_id" : 2, "item" : "jkl", "total" : 21, "billing_date" : ISODate("2014-02-06T09:00:00Z") }
// { "_id" : 3, "item" : "xyz", "total" : 5, "billing_date" : ISODate("2014-02-06T09:05:00Z") }
// { "_id" : 4, "item" : "abc", "total" : 13, "billing_date" : ISODate("2014-02-18T08:00:00Z") }
// { "_id" : 5, "item" : "xyz", "total" : 9, "billing_date" : ISODate("2014-02-18T09:12:00Z") }

db.sales.aggregate([
  {$group: {_id: '$item', avgAmount: {$avg: {$multiply: ['$price', '$quantity']}, avgQuantity: {$avg: '$quantity'}}}}
]);
// 按 item 对文档进行分组, 以检索非重复的项值
// 使用 $avg 累加器来计算每组的平均金额和平均数量
// output: avgAmount = sum(group(price * quantity)) / groupNum; avgQuantity = sum(group(quantity)) / groupNum
// { "_id" : "xyz", "avgAmount" : 37.5, "avgQuantity" : 7.5 }
// { "_id" : "jkl", "avgAmount" : 20, "avgQuantity" : 1 }
// { "_id" : "abc", "avgAmount" : 60, "avgQuantity" : 6 }
```

##### 数组表达式操作符

- $first 返回一组文档中第一个文档的表达式结果, 仅当文档按定义的顺序排列时才有意义
- $last 返回一组文档中最后一个文档的表达式结果, 仅当文档按指定顺序排列时才有意义
- $in 返回一个布尔值, 表示第一个值是否在第二个值为数组中, 不支持`正则表达式`匹配, 与 [查询操作符 $in](#in-1) 不同 <em id="in-2"></em> <!--markdownlint-disable-line-->
- $indexOfArray 搜索数组中出现的指定值, 并返回首次出现的数字索引, 如果未找到返回 -1, 如果不是一个数组或者数组不存在则返回 null
  - array 搜索的数组
  - search value 搜索的值
  - start 起始位置
  - end 结束位置
- $arrayElemAt 返回位于指定数组索引处的元素，如果索引越界不会返回结果, 如果数组不存在返回 null

```ts
// { $in: [ <expression>, <array expression> ] } 聚合操作符 $in
{ $in: [ 2, [ 1, 2, 3 ] ] } // true
{ $in: ['abc', ['xyz', 'abc']]} // true
{ $in: [ [ "a" ], [ "a" ] ] } // false
{ $in: [ [ "a" ], [ [ "a" ] ] ] } // true


{ $indexOfArray: [ [ "a", "abc" ], "a" ] }  // 0
{ $indexOfArray: [ [ "a", "abc", "de", ["de"] ], ["de"] ] } // 3
{ $indexOfArray: [ [ 1, 2 ], 5 ] }  // -1
{ $indexOfArray: [ [ 1, 2, 3 ], [1, 2] ] }  // -1
{ $indexOfArray: [ [ 10, 9, 9, 8, 9 ], 9, 3 ] } // 4
{ $indexOfArray: [ [ "a", "abc", "b" ], "b", 0, 1 ] } // -1
{ $indexOfArray: [ [ "a", "abc", "b" ], "b", 1, 0 ] } // -1
{ $indexOfArray: [ [ "a", "abc", "b" ], "b", 20 ] } // -1
{ $indexOfArray: [ [ null, null, null ], null ] } // 0
{ $indexOfArray: [ null, "foo" ] }  // null
{ $indexOfArray: [ "foo", "foo" ] } // 错误

{ $arrayElemAt: [ [ 1, 2, 3 ], 0 ] }  // 1
{ $arrayElemAt: [ [ 1, 2, 3 ], -2 ] } // 2
{ $arrayElemAt: [ [ 1, 2, 3 ], 15 ] } // 无返回值, 下标越界
{ $arrayElemAt: [ "$undefinedField", 0 ] }  // null, 第一个参数解析为未定义的数组
```

- $isArray 返回操作数是否为数组
- $reverseArray 接受数组表达式作为参数, 并返回其中的元素按倒序排列的数组
- $size 如果数组字段达到指定大小, 则选择文档
- $arrayToObject 将数组转换为单个文档, 数组必须是由两个元素组成的数组或者包含 k 和 v 字段的对象组成的数组

```ts
/*
// inventory
{ "_id" : 1, "item" : "ABC1",  dimensions: [ { "k": "l", "v": 25} , { "k": "w", "v": 10 }, { "k": "uom", "v": "cm" } ] }
{ "_id" : 2, "item" : "ABC2",  dimensions: [ [ "l", 50 ], [ "w",  25 ], [ "uom", "cm" ] ] }
{ "_id" : 3, "item" : "ABC3",  dimensions: [ [ "l", 25 ], [ "l",  "cm" ], [ "l", 50 ] ] }
*/
db.inventory.aggregate([
  {$project: {item: 1, $arrayToObject: '$dimensions'}}
]);
// output:
// { "_id" : 1, "item" : "ABC1", "dimensions" : { "l" : 25, "w" : 10, "uom" : "cm" } }
// { "_id" : 2, "item" : "ABC2", "dimensions" : { "l" : 50, "w" : 25, "uom" : "cm" } }
// { "_id" : 3, "item" : "ABC3", "dimensions" : { "l" : 50 } }
```

- $filter 根据指定条件选择要返回的数组的子集, 返回一个数组，其中仅包含与条件匹配的元素。返回的元素按原始顺序排列
  - input 输入的数组
  - as 每个元素的变量名
  - cond 可解析为布尔值的表达式, 它可用于确定输出数组中是否应包含某一元素.
  - limit 限制 $filter 返回的匹配大量元素的数量

```ts
// 语法
{
  $filter: {
    input: <array>,
    as: <string>,
    cond: <expression>,
    limit: <number expression>
  }
}
```

```ts
// db.sales.insertMany([
//   {
//     _id: 0,
//     items: [
//       { item_id: 43, quantity: 2, price: 10, name: "pen" },
//       { item_id: 2, quantity: 1, price: 240, name: "briefcase" }
//     ]
//   },
//   {
//     _id: 1,
//     items: [
//       { item_id: 23, quantity: 3, price: 110, name: "notebook" },
//       { item_id: 103, quantity: 4, price: 5, name: "pen" },
//       { item_id: 38, quantity: 1, price: 300, name: "printer" }
//     ]
//   },
//   {
//     _id: 2,
//     items: [
//       { item_id: 4, quantity: 1, price: 23, name: "paper" }
//     ]
//   }
// ]);

db.sales.aggregate([
  {
    $project: {
      items: {
        $filter: {
          input: "$items",
          as: "item",
          cond: { $gte: [ "$$item.price", 100 ] }
        }
      }
    }
  }
]);

// output
// [
//   {
//     _id: 0,
//     items: [ { item_id: 2, quantity: 1, price: 240, name: 'briefcase' } ]
//   }, {
//     _id: 1,
//     items: [
//       { item_id: 23, quantity: 3, price: 110, name: 'notebook' },
//       { item_id: 38, quantity: 1, price: 300, name: 'printer' }
//     ]
//   },
//   { _id: 2, items: [] }
// ]
```

- $map 对数组中的每个元素应用子表达式, 并按顺序返回生成值的数组
- $range 根据用户定义的输入，输出一个包含整数序列的数组
- $reduce 遍历数组中的每个元素, 并将它们组合成一个值
- $slice 返回数组的子集
- $sortArray 对数组的元素进行排序
- $zip 将两个数组进行合并

##### 条件表达式操作符

- $cond 一种三目运算符, 它可用于计算一个表达式, 并根据结果返回另外两个表达式之一的值
- $ifNull 返回第一个表达式的非空结果; 或者, 如果第一个表达式生成空结果, 则返回第二个表达式的结果
- $switch 对一系列 case 表达式求值
  - branches  控制分支文档的数组, 每个分支均为一个包含以下字段的文档
    - case  解析为 boolean 的任何有效表达式
    - then  任何有效表达式
  - default 在没有分支 case 表达式评估为 true 的情况下所采用的路径

```ts
{ $switch: {
    branches: [
      {case: {$eq: [0, 5]}, then: 'equals'},
      {case: {$gt: [0, 5]}, then: 'greater than'}
    ],
    default: 'not match'
  }
}
```

##### [日期表达式操作符](https://www.mongodb.com/zh-cn/docs/manual/reference/operator/aggregation/)

- $dateAdd 向日期对象添加多个时间单位
- $dateDiff 返回两个日期之间的差值
- $dateFromString 将日期/时间字符串转换为日期对象
- $dayOfMonth 返回日期中的某月的一天
- $hour 返回日期中的小时部分
- $minute 返回日期的分钟数
- $month 返回日期的月份
- $second 返回日期的秒数
- $week 返回日期的周数
- $year 返回日期的年份
- $toDate 将数值转换为日期

- $literal 返回一个值而不进行解析, 用于聚合管道可解释为表达式的值

```ts
/*
// inventory
{ "_id" : 1, "item" : "napkins", price: "$2.50" },
{ "_id" : 2, "item" : "coffee", price: "1" },
{ "_id" : 3, "item" : "soap", price: "$1" }
*/
db.inventory.aggregate([
  {$project: {costsOneDollar: {$eq: ['$price', {$literal: '$1'}]}}}
]);
// output: 判断 price 字段的值是否等于字符串 $1
// { "_id" : 1, "costsOneDollar" : false }
// { "_id" : 2, "costsOneDollar" : false }
// { "_id" : 3, "costsOneDollar" : true }
```

##### 对象表达式操作符

- $mergeObjects 将多个文档合并为一个文档, 后面的文档会覆盖前面文档的同名字段, 非文档的参数将会忽略
- $setField 添加、更新或删除文档中的指定字段
- $objectToArray 将文档转换为数组, 返回的数组中的每个元素都是一个包含两个字段 k 和 v 的文档

```ts
/*
// orders
{ "_id" : 1, "item" : "abc", "price" : 12, "ordered" : 2 }
{ "_id" : 2, "item" : "jkl", "price" : 20, "ordered" : 1 }
// items
{ "_id" : 1, "item" : "abc", description: "product 1", "instock" : 120 }
{ "_id" : 2, "item" : "def", description: "product 2", "instock" : 80 }
{ "_id" : 3, "item" : "jkl", description: "product 3", "instock" : 60 }
*/
db.orders.aggregate([
  {$lookup: {from: 'items', localField: 'item', foreignField: 'item', as: 'fromItems'}},
  {$replaceRoot: {newRoot: {$mergeObjects: [{$arrayElemAt: ['$fromItems', 0]}, '$$ROOT']}}},
  {$project: {fromItems: 0}}
]);
// output: 按照 item 字段连接两个集合, 使用 replaceRoot 中的 mergeObjects 合并来自 items 和 orders 中的连接文档
// {
//   _id: 1,
//   item: 'abc',
//   description: 'product 1',
//   instock: 120,
//   price: 12,
//   ordered: 2
// },
// {
//   _id: 2,
//   item: 'jkl',
//   description: 'product 3',
//   instock: 60,
//   price: 20,
//   ordered: 1
// }

/* 
// inventory
{ "_id" : 1, "item" : "ABC1",  dimensions: { l: 25, w: 10, uom: "cm" } }
{ "_id" : 2, "item" : "ABC2",  dimensions: { l: 50, w: 25, uom: "cm" } }
{ "_id" : 3, "item" : "XYZ1",  dimensions: { l: 70, w: 75, uom: "cm" } }
*/
db.inventory.aggregate([
  {$project: {item: 1, dimensions: {$objectToArray: '$dimensions'}}}
]);
// output:
// { "_id" : 1, "item" : "ABC1", "dimensions" : [ { "k" : "l", "v" : 25 }, { "k" : "w", "v" : 10 }, { "k" : "uom", "v" : "cm" } ] }
// { "_id" : 2, "item" : "ABC2", "dimensions" : [ { "k" : "l", "v" : 50 }, { "k" : "w", "v" : 25 }, { "k" : "uom", "v" : "cm" } ] }
// { "_id" : 3, "item" : "XYZ1", "dimensions" : [ { "k" : "l", "v" : 70 }, { "k" : "w", "v" : 75 }, { "k" : "uom", "v" : "cm" } ] }
```

##### 集合表达式操作符

- $allElementsTrue  如果集合中没有元素计算结果为 false, 则返回 true 否则返回 false
- $anyElementsTrue  如果集合中的任何元素计算结果为 true, 则返回 true 否则返回 false
- $setDifference  取两个集并返回一个包含仅存在于第一个集中的元素的数组, 即对第二个集取相对于第一个集的相对补集, 计算第二个集合的补集, 忽略重复项和顺序并且不会递归嵌套数组
- $setEquals  比较两个或多个数组, 如果具有相同的不同元素则返回 true, 否则返回 false, 忽略重复项和顺序并且不会递归嵌套数组
- $setIsSubset  比较两个数组, 当第一个数组是第二个数组的子集时返回 true, 否则返回 false, 忽略重复项和顺序并且不会递归嵌套数组
- $setIntersection  接收两个或多个数组, 并返回一个数组, 其中包含出现在每个输入数组中的元素, 计算交集, 忽略重复项和顺序并且不会递归嵌套数组
- $setUnion 接受两个或多个数组, 并返回一个数组, 其中包含出现在每个输入数组中的元素, 计算并集, 忽略重复项和顺序并且不会递归嵌套数组

```ts
{$setDifference: [['a','b','a'], ['b','a']]} // []
{$setDifference: [['a','b','c'], [['b','a']]]} // ['c']

{$setEquals: [['a','b','a'], ['b','a']]} // true
{$setEquals: [['a','b'], [['a','b']]]}  // false

{$setIsSubset: [['a','b','a'], ['b','a']]}  // true
{$setIsSubset: [['a','b'], [['a','b']]]}  // false

{$setIntersection: [['a','b','a'], ['b','a']]} // ['b', 'a']
{$setIntersection: [['a','b','a'], [['a','b']]]} // []

{$setUnion: [['a','b','a'], ['b','a']]} // ['b', 'a']
{$setUnion: [['a','b'], [['b', 'a']]]}  // ['a', 'b', ['a', 'b']]
```

##### 字符串表达式操作符

- $indexOfBytes 在字符串中搜索子串出现的位置, 并返回第一次出现的索引, 未找到则返回 -1
  - string 字符串值
  - substring 子串
  - start 起始位置
  - end 结束位置
- $regexFind  将正则表达式应用于字符串, 并返回第一个匹配子字符串的信息.

```ts
// { $regexFind: { input: <expression> , regex: <expression>, options: <expression> } }

db.restaurants.aggregate([
  {
    $addFields: {
      resultObject: { $regexFind: { input: "$category", regex: /cafe/, options: 'im' } }
    }
  }
]);
```

##### 类型表达式操作符

- $type 返回字段的 BSON 数据类型
- $toUUID 将字符串转换为 UUID
- $toObjectId 将值转换为 ObjectId
- $convert 将值转换为指定类型
  - input 任意有效的表达式
  - to 指定将 input 表达式转换为的类型，可以为字符串/对象格式
    - type 任何有效表达式, double->1, string->2, binData->5, objectId->7, bool->8, date->9, int->16, long->18, decimal->19
    - subType 指定要转换到的 binData 子类型
  - format 指定输入或输出的 binData 格式
    - base64
    - base64Url
    - utf8
    - hex
    - uuid

#### 分组结果合并原始文档

- 使用 `$mergeObjects` 合并字段

```ts
[{
  $group: {
    _id: '$groupField',
    count: {$sum: 1},
    // 合并所有原始文档的字段
    mergedDoc: {$mergeObjects: '$$ROOT'}
  }
}, {
  $project: {
    _id: 1,
    count: 1,
    // 展开合并后的文档
    // otherField1: '$mergedDoc.field1',
    // otherField2: '$mergedDoc.field2',
    // 保留整个合并文档
    originalDoc: '$mergedDoc'
  }
}]
```

- 使用 `$first` 或 `$last` 保留特定字段

```ts
[{
  $group: {
    _id: '$groupField',
    count: {$sum: 1},
    // 保留每个分组中第一个文档的特定字段
    // field1: {$first: '$field1'},
    // field2: {$first: '$field2'},
    // 保留整个文档
    originalDoc: {$first: '$$ROOT'}
  }
}]
```

- 使用 `$addFields` 或 `$project` 重组数据

```ts
[{
  $group: {
    _id: '$groupField',
    count: {$sum: 1},
    docs: {$push: '$$ROOT'}, // 将分组内的所有文档存入数组
  }
}, {
  $unwind: '$docs', // 展开文档数组
}, {
  $addFields: {
    // 将分组统计结果与原文档字段合并
    'docs.count': '$count',
    'docs.groupId': '$_id'
  }
}, {
  $replaceRoot: {
    newRoot: '$docs', // 将 docs 提升为根文档
  }
}]
```

- 使用 `$lookup` 自连接
- 使用 `$set` 简化操作

```ts
[{
  $group: {
    _id: '$groupField',
    count: {$sum: 1},
    firstDoc: {$first: '$$ROOT'}
  }
}, {
  $set: {
    'firstDoc.count': '$count',
    'firstDoc.groupId': '$_id'
  }
}, {
  $replaceRoot: {
    newRoot: '$firstDoc'
  }
}]
```

将一个订单集合, 按客户ID分组统计订单数量和总金额, 同时保留客户信息

```ts
db.orders.aggregate([{
  $group: {
    _id: '$customId',
    orderCount: {$sum: 1},
    totalAmount: {$sum: '$amount'},
    customerInfo: {
      $first: {
        name: '$customerName',
        email: '$customerEmail',
        joinDate: '$customerJoinDate'
      }
    }
  }
}, {
  $project: {
    customerId: '$_id',
    _id: 0,
    orderCount: 1,
    totalAmount: 1,
    name: '$customerInfo.name',
    email: '$customerInfo.email',
    joinDate: '$customerInfo.joinDate'
  }
}])
```

### 副本集

### 分片集
