---
title: sql
date: 2022-11-09 18:03:19
categories:
  - [server, sql]
tags:
  - sql
---

## 范式

键: 数据库中由一个或多个属性组成
超键: 在关系中能唯一标识记录的属性集称为关系模式的超键
候选键: 不包含多余属性的超键
主键: 被作为记录标识的候选键

```sql
-- 超键: {学号} {身份证号} 或者 {学号} {性别} 或者 {身份证号} {年龄} 的组合属性集可以确定唯一一条记录
-- 候选键: {学号} 或者 {身份证号} 可以是候选键
-- 主键: {学号} 候选键作为主键
```

- 1NF: 数据库表中的每一列的字段都是不可再分的, 原子性
- 2NF: 数据库表中不存在非关键字段对任意候选关键字段的部分函数依赖(指存在着组合关键字中的某一关键字决定非关键字的情况)
  - 部分函数依赖：没有包含在主键中的列必须完全依赖于主键, 而不能只依赖于主键的一部分
- 3NF: 数据库表中不存在非关键字段对任意候选关键字段的传递函数依赖
  - 传递函数依赖: 非主键列必须直接依赖于主键列，不能存在非主键列依赖于非主键列的情况
- BCNF: 数据库表中不存在任何关键字段对任意候选关键字段(组合关键字)的传递函数依赖

```sql
/*
2NF:
例如 订单明细表(orderDetail): orderId, productId, productName, unitPrice, discount, quantity; 主键为 (orderId, productId). 一个订单可以包含多个商品.
discount, quantity 完全依赖主键, 而 unitPrice, productName 只依赖于 productId, 不符合 2NF.
将 unitPrice, productName 和 productId 拆分为单独的表消除原订单明细表 unitPrice, productName 多次重复的情况
3NF:
例如 订单表(order): orderId, orderDate, customerId, customerName, customerAddr, customerCity; 主键为 orderId
customerName, customerAddr, customerCity 直接依赖 customerId(非主键列), 通过传递才依赖主键, 所以不符合 3NF.
将 customerName, customerAddr, customerCity 拆分为单独的表达到满足 3NF.
*/
```

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

## 权限控制

```sql
-- 创建使用 utf8mb4 字符集的数据库
CREATE DATABASE database_name CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 查看当前字符集设置
SHOW VARIABLES LIKE 'character_set%';

-- 设置字符集
SET NAMES utf8mb4;
```

### 用户权限

```sql
-- 创建用户和密码
CREATE USER 'username'@'host' IDENTIFIED BY 'password';

-- 修改用户密码
ALTER USER 'username'@'host' IDENTIFIED BY 'password';

-- 删除用户
DROP USER 'username'@'host';
```

### 操作权限

```sql
-- 完整语法模板
GRANT 权限列表 ON 数据库名.表名 TO 'username'@'host' [IDENTIFIED BY '密码'];

-- 撤销权限
REVOKE 权限列表 ON 数据库名.表名 FROM 'username'@'host';

-- 查看指定用户的权限
SHOW GRANTS FOR 'username'@'host';

-- 刷新权限缓存
FLUSH PRIVILEGES;
```

- 权限列表, 逗号分隔的权限集合（SELECT/INSERT/UPDATE等）或 `ALL PRIVILEGES`
- 数据库名.表名, 支持通配符：*.*（全局）、mydb.*（单库）、mydb.orders（单表）
- username@host, 用户标识与访问来源 localhost，192.168.1.%，% 任意主机
- IDENTIFIED BY, 可选参数，用户不存在时自动创建并设置密码

```sql
-- 授予本地开发用户读写权限
GRANT SELECT, INSERT, UPDATE ON mydb.* TO 'dev_user'@'localhost' IDENTIFIED BY 'secure_password123';

-- 授予内网管理员全部权限（不含GRANT权限）
GRANT ALL PRIVILEGES ON testdb.* TO 'admin_user'@'192.168.1.%' IDENTIFIED BY 'admin@123';

-- 允许任意IP的只读访问
GRANT SELECT ON report_db.* TO 'readonly_user'@'%' IDENTIFIED BY 'readonly_pass';
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

```sql
SELECT 
    [DISTINCT] 
    select_expr, ...
FROM 
    table1
[INNER | LEFT | RIGHT | FULL OUTER | CROSS] JOIN table2
    ON join_condition
[INNER | LEFT | RIGHT | FULL OUTER | CROSS] JOIN table3
    ON join_condition
[WHERE where_condition]
[GROUP BY col_list]
[HAVING having_condition]
[ORDER BY col_list]
[LIMIT [offset,] row_count];

/*
employees 表结构：‌
  id
  name
  department_id
departments 表结构：‌
  id
  name
*/
```

#### 内连接

- 相当于查询 A、B 交集部分数据

```sql
-- 显式内连接
select * from t1 inner join t2 on condition;
-- 隐式内连接
select * from t1, t2 where condition;

-- 查询分数大于80分的学生及其班级
SELECT s.name, s.score, c.class_name
FROM students s
INNER JOIN classes c ON s.class_id = c.class_id
WHERE s.score > 80
ORDER BY s.score DESC;

-- 查询返回所有匹配 部门 id 的员工的名字和部门名称
SELECT employees.name, departments.name
FROM employees
INNER JOIN departments
ON employees.department_id = departments.id;
```

#### 外连接

##### 左外连接

- 查询**左表**的所有记录, 以及右表中关联字段相匹配的记录, 结果中右表没有匹配的部分填充 null

```sql
select * from t1 left join t2 on condition where where_condition;

-- 查询每个班级的平均分
SELECT 
    c.class_name,
    COUNT(s.student_id) AS student_count,
    AVG(s.score) AS avg_score,
    MAX(s.score) AS max_score
FROM classes c
LEFT JOIN students s ON c.class_id = s.class_id
GROUP BY c.class_id, c.class_name
HAVING AVG(s.score) > 80
ORDER BY avg_score DESC;

-- 查询所有员工的信息, 如果某个员工没有部门, 则部门名称为 null
SELECT employees.name, departments.name
FROM employees
LEFT JOIN departments
ON employees.department_id = departments.id;
```

##### 右外连接

- 查询**右表**的所有记录, 以及左表中关联字段相匹配的记录, 结果中左表没有匹配的部分填充 null

```sql
select * from t1 right join t2 on condition where where_condition;

-- 查询所有部门的信息, 如果某个部门没有员工, 则员工名称为 null
SELECT employees.name, departments.name
FROM employees
RIGHT JOIN departments
ON employees.department_id = departments.id;
```

##### 全外连接

- 返回两个表中的所有记录, 当一边没有匹配时, 另一边的记录将填充 null
- mysql 不支持 全外连接, 使用 左连接和右连接的结果做 union 操作

```sql
select * from t1 full outer join t2 on condition where where_condition;

-- 使用 左连接 UNION 右连接
SELECT employees.name, departments.name
FROM employees
LEFT JOIN departments
ON employees.department_id = departments.id
UNION
SELECT employees.name, departments.name
FROM employees
RIGHT JOIN departments
ON employees.department_id = departments.id;
```

#### 交叉连接/笛卡尔积

返回两个表的笛卡尔积(所有可能的组合)

```sql
select * from t1 as a cross join t2 as b on condition;
```

#### 自连接

- 当前表与自身的连接查询, 自连接必须使用表别名
- 可以是内连接查询, 也可以是外连接查询

```sql
select * from t1 as a inner join t2 as b on condition;
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

### 链接字符串

`mongodb://[username:password@]host1[:port1][,host2[:port2],...][/database][?options]`

认证

- authSource=admin 认证数据库，默认使用连接字符串中指定的database
- authMechanism=SCRAM-SHA-1 使用 SCRAM-SHA-256, MONGODB-X509, PLAIN, GSSAPI等

连接

- maxPoolSize=50  连接池最大连接数，默认100
- cminPoolSize=10  连接池最小连接数，默认0
- maxIdleTimeMS=30000 连接最大空闲时间（毫秒），默认无限制
- waitQueueTimeoutMS=120000 连接等待超时时间（毫秒），默认120秒
- connectTimeoutMS=10000  连接建立超时时间，默认30秒

读写操作

- readPreference=primary  primary|primaryPreferred|secondary|secondaryPreferred|nearest, 默认primary
- readPreferenceTags=dc:nyc,rack:1  标签集，用于更精细的读取控制
- writeConcern=1  写入确认级别：0(不确认),1(主节点确认),majority(大多数节点确认)或数字表示确认节点数
- wtimeoutMS=5000 写入超时时间（毫秒）
- journal=true  是否等待日志写入，默认false

副本集/分片集群

- replicaSet=myReplicaSet 副本集名称
- readConcernLevel=local  local|majority|available|linearizable|snapshot
- directConnection=true 是否直连单个节点，默认false（自动发现）
- shardOptions=xxx  分片集群相关选项

TSL/SSL 安全

- tls=true  启用TLS，同ssl=true
- tlsCAFile=/path/to/ca.pem CA证书文件路径
- tlsCertificateKeyFile=/path/to/client.pem 客户端证书文件
- tlsAllowInvalidCertificates=true  允许自签名证书（仅测试环境使用）
- tlsAllowInvalidHostnames=true 允许主机名不匹配（仅测试环境使用）

超时和重试

- socketTimeoutMS=0 Socket操作超时，默认0（不超时）
- serverSelectionTimeoutMS=30000  服务器选择超时，默认30秒
- heartbeatFrequencyMS=10000  心跳频率，默认10秒
- retryWrites=true  是否重试写入，默认true
- retryReads=true 是否重试读取，默认true
- maxStalenessSeconds=90  最大陈旧时间（秒），用于readPreference

压缩

- compressors=zlib  snappy|zlib|zstd，可多个用逗号分隔：snappy,zlib
- zlibCompressionLevel=6  zlib压缩级别 1-9，默认6

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

```typescript
// 使用 $ 投影 compass
// filter:  { 'audioList.id': 'abc' } 等价于 { audioList: { $elemMatch: { id: 'audio123' } } }
// project: { 'audioList.$': 1 }


// 使用聚合过滤嵌套数组中多条数据只保留 id 为 audio123 的一条数据
db.inventory.aggregate([
  // 1. 匹配文档
  {
    $match: {
      "audioList.id": "audio123"
    }
  },
  // 2. 筛选数组元素
  {
    $addFields: {
      audioList: {
        $filter: {
          input: "$audioList",
          as: "audio",
          cond: { $eq: ["$$audio.id", "audio123"] }
        }
      }
    }
  }
])
```

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

## Mongoose

### SchemaTypes

- String
- Number
- Boolean
- Array
- Buffer
- Date
- ObjectId
- Mixed, 可以保存任意值
- UUID
- BigInt
- Double
- Int32

直接声明某种类型, 或者赋值一个含有 type 属性的对象

```ts
const schema1 = new Schema({
  test: String
});
const schema2 = new Schema({
  test: { type: String, lowercase: true}
});
```

#### 选项

##### 通用选项

- required, 布尔值或函数, 如果值为真, 为此属性添加 required 验证器
- default, 任何值或函数, 设置此属性的默认值
- select, 布尔值, 指定查询时返回的默认 projections
- validate, 函数, 添加属性自定义验证器
- get, 函数, 使用 Object.defineProperty() 定义自定义 getter
- set, 函数, 使用 Object.defineProperty() 定义自定义 setter
- alias, 字符串, 设置属性别名
- transform, 函数, 当转换为 JSON 字符串时调用, `Document#toJSON()` 和 `JSON.stringify()`

```ts
const numberSchema = new Schema({
  integerOnly: {
    type: Number,
    get: v => Math.round(v),
    set: v => Math.round(v),
    alias: 'i',
  }
});
const Number = mongoose.Model('Number', numberSchema);
const num = new Number();
num.integerOnly = 2.001;
num.integerOnly; // 2
num.i; // 2
num.i = 3.001;
num.integerOnly; // 3
num.i; // 3

// 自定义验证器
new Schema({
  phone: {
    type: String,
    validate: {
      validator: (v) => {
        return /\d{3}-\d{3}-\d{4}/.test(v);
      },
      message: '{VALUE} is not a valid phone number',
    },
    required: [true, 'User phone number required']
  }
});
```

##### 索引选项

- index, 布尔值, 是否对这个属性创建索引
- unique, 布尔值, 是否对这个属性创建唯一索引
- sparse, 布尔值, 是否对这个属性创建稀疏索引

```ts
const schema2 = new Schema({
  test: {
    type: String,
    index: true,
    unique: true,
  }
});
```

##### 字符串选项

- lowercase, 布尔值, 是否在保存前对此值调用 .toLowerCase()
- uppercase, 布尔值, 是否在保存前对此值调用 .toUpperCase()
- trim, 布尔值, 是否在保存前对此值调用 .trim()
- match, 正则表达式, 创建验证器检查属性是否匹配给定正则表达式
- enum, 数组, 创建验证器检查属性是否包含于给定数组
- minlength, 数值, 创建验证器检查属性是否大于该值
- maxlength, 数值, 创建验证器检查属性是否小于该值

##### 数值选项

- min, 数值, 创建验证器检查属性是否大于或等于该值
- max, 数值, 创建验证器检查属性是否小于或等于该值
- enum, 数组, 创建验证器检查属性是否包含于给定数组

##### 日期选项

- min, Date, 创建验证器检查属性是否大于该值
- max, Date, 创建验证器检查属性是否小于该值
- expires, 数值或字符串, 创建以秒为单位的生存时间

#### 验证器

验证器定义于 SchemaType, 是一个中间件, 默认作为 pre('save') 钩子注册在 schema 上

- 手动验证 `doc.validate(callback)` 或 `doc.validateSync()`
- 不对未定义的值进行验证, 唯一例外是 required 验证器
- 验证是异步递归的

```ts
const schema = new Schema({
  name: { type: String, required: true }
});
const Cat = mongoose.model('Cat', schema);

const cat = new Cat();
cat.save(function(error) {
  assert.equal(error.errors['name'].message, 'Path \'name\' is required');
  error = cat.validateSync();
  assert.equal(error.errors['name'].message, 'Path \'name\' is required');
});
```

##### 内建验证器

- required 验证器
- 字符串选项的 enum, match, maxlength, minlength 验证器
- 数值选项的 min 和 max 验证器

### SchemaOptions

- audoIndex: boolean, 默认为 null, 使用数据连接自带的 autoIndex 选项
- autoCreate: boolean, 默认为 null, 使用数据库连接自带的 autoCreate 选项
- capped: boolean/number/Object, 默认为 false, 设置集合的大小上限
- collection: string, 没有默认值, 指定集合名称, 默认使用模型名称作为集合名称
- id: boolean, 默认为 true, 虚拟 id 指向文档的 _id
- _id: boolean, 默认为 true, 缺失无法更新文档
- minimize: boolean, 默认为 true 不保存空对象, 当手动调用时控制 `document#toObject` 的行为
- readConcern: Object, 默认为 null, 设置所有查询的读关注
- writeConcern: Object, 默认为 null, 设置写关注
- strict: boolean, 默认为 true, 不能保存 Schema 中没有声明的属性
- typeKey: string, 默认为 `type`, 设置属性类型的键
- validateBeforeSave: boolean, 默认为 true, 保存之前验证
- versionKey: string/Object, 默认为 '__v'
- timestamps: boolean/Object, 默认为 false, 如果为 true，Mongoose 将添加 `createAt` 和 `updateAt` 属性到 Schema 中并维护它们
- lean: boolean, 默认为 false, 如果为 true, 则所有的查询都使用 lean 将结果转换为普通 JavaScript 对象
- selectPopulatedPaths: boolean, 默认为 true, 设置在查询时是否返回此属性

```ts
const schema = new Schema({
  name: { $type: String, select: false }
}, {
  autoIndex: false, // 禁用自动创建索引
  autoCreate: false, // 禁用自动创建集合
  collaction: 'data', // 指定集合名称
  typeKey: '$type'
})
```

### 查询方法

#### select

指定查询时返回的结果中的字段, 参数接收 string/array/object

```ts
// include a and b, exclude other fields
query.select('a b');
// Equivalent syntaxes:
query.select(['a', 'b']);
query.select({ a: 1, b: 1 });

// exclude c and d, include other fields
query.select('-c -d');

// Use `+` to override schema-level `select: false` without making the
query.select('+name');
```
