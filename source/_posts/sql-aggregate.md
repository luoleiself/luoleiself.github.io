---
title: sql-aggregate
date: 2022-11-09 18:03:19
categories:
  - [server, sql]
tags:
  - sql
---

## 聚合操作

操作处理多个文档并返回计算结果, 由 一个或多个处理文档的 `阶段` 组成

- 每个阶段对输入文档执行一个操作
- 从一个阶段输出的文档将传递到下一个阶段
- 一个聚合管道可以返回针对文档组的结果

**字段路径**, 使用 `$` 前缀启用字段路径表达式访问输入文档中的字段.

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

### 变量

变量可以保存任何 BSON 类型的数据, 访问变量时需要使用 `$$` 前缀.

#### 用户变量

变量名称可包含 ASCII 字符和任意非 ASCII 字符, 必须以小写 ASCII 字符开头

#### 系统变量

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

<!--more-->

### 聚合命令

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

### 阶段

文档按顺序通过聚合管道阶段

- $addFields 为文档添加新字段
- $bucket 根据指定的表达式和存储桶边界将传入的文档分为多个组, 并为每个组输出一个文档
  - groupBy 对文档进行分组的表达式
  - boundaries 边界,基于指定每个存储桶边界的 groupBy 表达式的值
  - default 指定附加存储桶 _id 的字面量
  - output 一份文档, 指定除 _id 字段之外要包含在输出文档中的字段
- $bucketAuto 根据指定的表达式, 将接收到的文档归类到特定数量的组中, 自动确定存储桶的边界,以尝试将文档均匀地分配到指定数量的存储桶中
  - buckets 整型, 一个 32 位正整数, 用于指定将输入文档按组分到存储桶的数量

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

{
  $bucketAuto: {
    groupBy: <expression>,
    buckets: <number>,
    output: {
      <output1>: { <$accumulator expression> },
      ...
    }
    granularity: <string>
  }
}
```

- $changeStream  返回集合、数据库或整个集群上的变更流游标, 必须是聚合管道中的 **第一阶段**
- $collStats  返回有关集合或视图的统计信息, 必须是聚合管道中的 **第一阶段**
- $count  将文档传递到下一阶段, 该阶段包含输入到该阶段的文档数的计数

```ts
db.scores.insertMany([
  { "_id" : 1, "subject" : "History", "score" : 88 },
  { "_id" : 2, "subject" : "History", "score" : 92 },
  { "_id" : 3, "subject" : "History", "score" : 97 },
  { "_id" : 4, "subject" : "History", "score" : 71 },
  { "_id" : 5, "subject" : "History", "score" : 79 },
  { "_id" : 6, "subject" : "History", "score" : 83 }
]);
// 使用 $match 阶段排除 score 值小于或等于 80 的文档, 以便将 score 大于 80 的文档传递到下一个阶段
// $count 阶段回返回聚合管道中剩余文档的计数, 并将改值分配给名为 passing_scores 的字段
db.scores.aggregate([
  { $match: { score: { $gt: 80 } } },
  { $count: "passing_scores" }
]);
// output: 输出结果
{ 'passing_scores': 4 }
```

- $densify  在文档序列中创建新文档, 其中缺少字段中的某些值
- $documents  从输入表达式返回字面文档
  - 只能在数据库级聚合管道中使用
  - 必须是聚合管道中的 **第一阶段**
- $facet  在单个阶段内处理同一组输入文档上的多个聚合管道, 每个子管道在输出文档中都有自己的字段, 其结果存储为文档数组
- $fill 填充文档中的 null 和缺失的字段值
  - sortBy 指定每个分区内用于对文档进行排序的字段, 使用与 $sort 阶段相同的语法
  - output 指定一个对象, 其中包含要填充缺失值的每个字段

```ts
// dailySales
// [
//   {
//     "date": ISODate("2022-02-02"),
//     "bootsSold": 10,
//     "sandalsSold": 20,
//     "sneakersSold": 12
//   },
//   {
//     "date": ISODate("2022-02-03"),
//     "bootsSold": 7,
//     "sneakersSold": 18
//   },
//   {
//     "date": ISODate("2022-02-04"),
//     "sneakersSold": 5
//   }
// ]
// 将每天销售中缺少的鞋款的销售数量设置为 0
db.dailySales.aggregate( [{
  $fill: {
    output: {
      "bootsSold": { value: 0 },
      "sandalsSold": { value: 0 },
      "sneakersSold": { value: 0 }
    }
  }
}]);
```

- $geoNear  根据与地理空间点的接近程度返回有序的文档流
- $graphLookup  对集合执行递归搜索, 并提供按照递归深度和查询筛选器限制搜索的选项
- $group  按指定的标识符表达式对输入文档进行分组，并将累加器表达式（如果指定）应用于每个群组, 不会对其输出文档进行排序
  - _id 指定群组标识符表达式, 如果指定的 \_id 值为空值或任何其他常量值, $group 阶段将返回聚合所有输入文档值的单个文档
  - field 可选, 使用累加器操作符进行计算

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
  - let 指定在管道阶段中使用的变量, 使用变量表达式访问本地集合文档中的字段, 这些文档输入到 pipeline
  - pipeline 指定在 `外部集合` 上运行的管道, 返回外部集合的文档, 如果返回所有文档, 指定一个空的管道: []
    - 不能包含 $merge 或 $out 阶段.
    - 无法访问输入文档中的字段, 可以使用 let 选项定义文档字段的变量, 然后在 pipeline 阶段引用这些变量.
  - as 指定要添加到输入文档中的新数组字段的名称, 新数组字段包含来自 from 集合的匹配文档, 如果指定名称已存在将被重写

```ts
/*
// orders
[
  { _id: 1, item: "almonds", price: 12, ordered: 2 },
  { _id: 2, item: "pecans", price: 20, ordered: 1 },
  { _id: 3, item: "cookies", price: 10, ordered: 60 }
] 
// warehouses
[
  { _id: 1, stock_item: "almonds", warehouse: "A", instock: 120 },
  { _id: 2, stock_item: "pecans", warehouse: "A", instock: 80 },
  { _id: 3, stock_item: "almonds", warehouse: "B", instock: 60 },
  { _id: 4, stock_item: "cookies", warehouse: "B", instock: 40 },
  { _id: 5, stock_item: "cookies", warehouse: "A", instock: 80 }
]
*/
db.orders.aggregate([
  {
    $lookup: {
      from: 'warehouses',
      localField: 'item',
      foreignField: 'stock_item',
      let: { order_qty: '$ordered' },
      pipeline: [
        {
          $match: {
            $expr: [ { $gte: ['$instock', '$$order_qty'] } ]
          } 
        },
        { $project: { stock_item: 0， _id: 0 } }
      ],
      as: 'stockdata'
    }
  }
]);
// output: localField 和 foreignField 不包含的则作为 null 值执行匹配
// {
//   _id: 1,
//   item: 'almonds',
//   price: 12,
//   ordered: 2,
//   stockdata: [
//     { warehouse: 'A', instock: 120 },
//     { warehouse: 'B', instock: 60 }
//   ]
// },
// {
//   _id: 2,
//   item: 'pecans',
//   price: 20,
//   ordered: 1,
//   stockdata: [ { warehouse: 'A', instock: 80 } ]
// },
// {
//   _id: 3,
//   item: 'cookies',
//   price: 10,
//   ordered: 60,
//   stockdata: [ { warehouse: 'A', instock: 80 } ]
// }
```  

- $match  筛选文档流以仅允许匹配的文档将未修改的文档传入下一管道阶段
- $merge  将 aggregation pipeline 的结果文档写入集合, 必须是聚合管道中的 **最后一阶段**
- $out  将 aggregation pipeline 的结果文档写入集合, 必须是聚合管道中的 **最后一阶段**
- $planCacheStats 返回集合的计划缓存的信息, 必须是聚合管道中的 **第一阶段**
  - allHosts  配置聚合阶段如何以分片集群中的节点为目标
- $project  重塑流中的文档至管道中的下个阶段，指定的字段可以是文档中已有字段或新计算的字段. 例如添加新的字段或删除现有字段

- $redact 根据存储在文档本身中的信息, 限制整个文档被输出或者文档中的内容被输出
  - $$DESCEND  返回当前文档级别的字段, 不包括嵌入式文档
  - $$PRUNE  排除当前文档/嵌入式文档级别的所有字段, 而不进一步检查任何已排除的字段
  - $$KEEP  返回或保留此当前文档/嵌入式文档级别的所有字段, 而不进一步检查此级别的字段

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

```ts
{ $sort: { <field1>: <sort order>, <field2>: <sort order> ... } }
```

- $unionWith 将两个集合合并为一个结果集
  - coll 如果省略 pipeline 则为必传
  - pipeline  应用于输入文档聚合管道

```ts
{ $unionWith: { coll: "<collection>", pipeline: [ <stage1>, ... ] } }
```

- $unset  从文档中删除/排除字段, $unset 是删除字段的 $project 阶段的别名
- $unwind 解构输入文档中的数组字段, 以便为数组中的每个元素输出文档, 并用该元素替换该数组字段的值, 忽略无法转换成 `单元素数组` 的文档
  - path  数组字段的字段路径
  - includeArrayIndex  可选, 新字段的名称, 用于保存该元素的数组索引, 名称不能以 `$` 开头
  - preserveNullAndEmptyArrays  可选, 默认为 false
    - 如果为 true, 如果 path 为 null、缺失或空, $unwind 会输出文档
    - 如果为 false, 如果 path 为 null、缺失或空, $unwind 不会输出文档

```ts
{
  $unwind: {
    path: <field path>,
    includeArrayIndex: <string>,
    preserveNullAndEmptyArrays: <boolean>
  }
}

/*
// inventory
{ "_id" : 1, "item" : "ABC1", sizes: [ "S", "M", "L"] }
*/
db.inventory.aggregate([{$unwind: {path: '$sizes', includeArrayIndex: 'arrayIndex'} }]);
// output:
// { "_id" : 1, "item" : "ABC1", "sizes" : "S", arrayIndex: NumberLong(0) }
// { "_id" : 1, "item" : "ABC1", "sizes" : "M", arrayIndex: NumberLong(1) }
// { "_id" : 1, "item" : "ABC1", "sizes" : "L", arrayIndex: NumberLong(2) }

// 数据结构
/* 
[{
  _id: string,
  name: string,
  modules: [{
    name: string,
    nodeId: string,
    flowNodeType: string,
    inputs: [{
      key: string,
      value: string,
    }]
  }]
}]
*/
[
  { $match: { type: "simple" } }, // 过滤 type 为 simple 的文档
  {
    $unwind: {
      path: "$modules", // 解构 modules 数组, 为数组中的每个元素输出文档,并用该元素替换该字段的值, 
      includeArrayIndex: "m", // 添加 m 字段标识当前元素所在的下标
      preserveNullAndEmptyArrays: true
    }
  },
  { $match: { 'modules.flowNodeType': 'chatNode' } }, // 过滤 modules.flowNodeType 为 chatNode 的文档
  {
    $addFields: { // 添加字段
      inputsSize: { $size: "$modules.inputs" }, // 值为 modules.inputs 数组的大小
      inputsValues: {
        $arrayToObject: { // 将键值对数组转换为对象
          $map: { // 遍历 modules.inputs 的每一项, 并返回每一项的 key 属性的值和 value 属性的值组成的数组
            input: "$modules.inputs",
            as: "item",
            in: ["$$item.key", "$$item.value"]
          }
        }
      }
    }
  }
]
```

### 操作符

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

#### 数组表达式操作符

- $first 返回一组文档中第一个文档的表达式结果, 仅当文档按定义的顺序排列时才有意义
- $last 返回一组文档中最后一个文档的表达式结果, 仅当文档按指定顺序排列时才有意义
- $in 返回一个布尔值, 表示第一个值是否在第二个值为数组中, 不支持`正则表达式`匹配, 与 查询操作符 $in 不同
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

- $isArray 判断操作数是否为数组, 返回一个布尔值
- $reverseArray 接受数组表达式作为参数, 并返回其中的元素按倒序排列的数组
- $size 如果数组字段达到指定大小, 则选择文档
- $arrayToObject 将数组转换为单个文档
  - 数组必须是由两个元素组成的数组
  - 包含 k 和 v 字段的对象组成的数组

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

- $map 对数组中的每个项目应用表达式，并返回包含已应用结果的数组
  - input 解析为数组的表达式
  - as  可选, 代表 input 数组中每个元素的变量名称
  - in  应用于 input 数组中每个元素的表达式, 该表达式单独引用带有 as 中指定的变量名称的每个元素

```ts
{ $map: { input: <expression>, as: <string>, in: <expression> } }
{
  $map: {
    input: '$pubStat',
    as: 'item',
    in: {
      $and: [
        {
          $or: [
            { $eq: ['$$item.auditStatus', 'success'] },
            { $eq: ['$$item.auditStatus', 'faile'] }
          ]
        }, {
          $eq: ['$$item.visibilityType', 2]
        },
      ]
    }
  }
}
```

- $range 返回一个数组，其元素是生成的数字序列
  - start 开始
  - end 结束
  - step  步长

```ts
{ $range: [ <start>, <end>, <non-zero step> ] }
{ $range: [ 0, 10, 2 ] }  // [ 0, 2, 4, 6, 8 ]
{ $range: [ 10, 0, -2 ] } // [ 10, 8, 6, 4, 2 ]
{ $range: [ 0, 10, -2 ] } // [ ]
{ $range: [ 0, 5 ] }  // [ 0, 1, 2, 3, 4 ]
```

- $reduce 将表达式应用于数组中的每个元素，并将它们组合成一个值
  - input 解析为数组的表达式
  - initialValue  在 in 之前设置的初始累积 value, 将应用于 input 数组的第一个元素
  - in  一个有效的表达式, $reduce 将按从左到右的顺序应用于 input 数组中的每个元素

```ts
{
  $reduce: {
    input: <array>,
    initialValue: <expression>,
    in: <expression>
  }
}
{
  $reduce: {
    input: ["a", "b", "c"],
    initialValue: "",
    in: { $concat : ["$$value", "$$this"] }
  }
}
// "abc"

{
  $reduce: {
    input: [ 1, 2, 3, 4 ],
    initialValue: { sum: 5, product: 2 },
    in: {
      sum: { $add : ["$$value.sum", "$$this"] },
      product: { $multiply: [ "$$value.product", "$$this" ] }
    }
  }
}
// { "sum" : 15, "product" : 48 }

{
  $reduce: {
    input: [ [ 3, 4 ], [ 5, 6 ] ],
    initialValue: [ 1, 2 ],
    in: { $concatArrays : ["$$value", "$$this"] }
  }
}
// [ 1, 2, 3, 4, 5, 6 ]
```

- $slice 返回数组的子集
  - array 解析为数组的表达式
  - position  解析为整数的有效表达式, 指定截取的位置
  - n 解析为整数的有效表达式, 指定返回元素的个数
- $sortArray 对数组的元素进行排序
  - input 待排序的数组
  - sortBy  指定排序的文档或布尔值

```ts
{ $slice: [ <array>, <position>, <n> ] }
{ $slice: [ [ 1, 2, 3 ], 1, 1 ] }  //[ 2 ]
{ $slice: [ [ 1, 2, 3 ], -2 ] } // [ 2, 3 ]
{ $slice: [ [ 1, 2, 3 ], 15, 2 ] }  // [  ]
{ $slice: [ [ 1, 2, 3 ], -15, 2 ] } // [ 1, 2 ]

{
  $sortArray: {
    input: <array>,
    sortBy: <sort spec>
  }
}
```

- $zip 将两个数组进行合并, 将 `[ [ 1, 2, 3 ], [ "a", "b", "c" ] ]` 转化为 `[ [ 1, "a" ], [ 2, "b" ], [ 3, "c" ] ]`

#### 条件表达式操作符

- $cond 一种三目运算符, 它可用于计算一个表达式, 并根据结果返回另外两个表达式之一的值

```ts
{ $cond: { if: <boolean-expression>, then: <true-case>, else: <false-case> } }
{ $cond: [ <boolean-expression>, <true-case>, <false-case> ] }
```

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

#### 自定义聚合表达式操作

- $accumulator 定义自定义累加器
- $function 定义自定义函数

```ts
{
  $accumulator: {
    init: <code>,
    initArgs: <array expression>,        // Optional
    accumulate: <code>,
    accumulateArgs: <array expression>,
    merge: <code>,
    finalize: <code>,                    // Optional
    lang: <string>
  }
}

{
  $function: {
    body: <code>,
    args: <array expression>,
    lang: "js"
  }
}
/* 
{ $function:
  {
    body: function(name) {
      return hex_md5(name) == "15b0a220baa16331e8d80e15367677ad"
    },
    args: [ "$name" ],
    lang: "js"
  }
}
*/
```

#### [日期表达式操作符](https://www.mongodb.com/zh-cn/docs/manual/reference/operator/aggregation/)

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

#### 对象表达式操作符

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

#### 集合表达式操作符

- $allElementsTrue  如果集合中没有元素计算结果为 false, 则返回 true 否则返回 false
- $anyElementTrue  如果集合中的任何元素计算结果为 true, 则返回 true 否则返回 false
- $setDifference  取两个集并返回一个包含仅存在于第一个集中的元素的数组, 即对第二个集取相对于第一个集的相对补集, 计算第二个集合的补集, 忽略重复项和顺序并且不会递归嵌套数组
- $setEquals  比较两个或多个数组, 如果具有相同的不同元素则返回 true, 否则返回 false, 忽略重复项和顺序并且不会递归嵌套数组
- $setIsSubset  比较两个数组, 当第一个数组是第二个数组的子集时返回 true, 否则返回 false, 忽略重复项和顺序并且不会递归嵌套数组
- $setIntersection  接收两个或多个数组, 并返回一个数组, 其中包含出现在每个输入数组中的元素, 计算交集, 忽略重复项和顺序并且不会递归嵌套数组
- $setUnion 接受两个或多个数组, 并返回一个数组, 其中包含出现在每个输入数组中的元素, 计算并集, 忽略重复项和顺序并且不会递归嵌套数组

```ts
// { $allElementsTrue: [ <expression> ] }
{ $allElementsTrue: [ [ true, 1, "someString" ] ] } // true
{ $allElementsTrue: [ [ [ false ] ] ] } // true
{ $allElementsTrue: [ [ ] ] } // true
{ $allElementsTrue: [ [ null, false, 0 ] ] } // false

// { $anyElementTrue: [ <expression> ] }
{ $anyElementTrue: [ [ true, false ] ] } // true
{ $anyElementTrue: [ [ [ false ] ] ] } // true
{ $anyElementTrue: [ [ null, false, 0 ] ] } // false
{ $anyElementTrue: [ [ ] ] } // false

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

db.collection.aggregate([
  {
    $match: {
      $expr: {
        $anyElementTrue: {
          $map: {
            input: "$arrayField",
            as: "item",
            in: { // 条件表达式
              $and: [
                { $gt: ["$$item.price", 100] }, 
                { $lt: ["$$item.price", 1000] }
              ] 
            }
          }
        }
      }
    }
  }
]);
```

#### 字符串表达式操作符

- $concat 连接任意数量的字符串
- $dateToString 以格式化字符串的形式返回日期
- $indexOfBytes 在字符串中搜索子串出现的位置, 并返回第一次出现的索引, 未找到则返回 -1
  - string 字符串值
  - substring 子串
  - start 起始位置
  - end 结束位置
- $indexOfCP  在字符串中搜索子串出现的位置, 并返回第一次出现时的 UTF-8 码位索引, 未找到则返回 -1
- $ltrim  删除字符串开头和结尾的空白或指定字符
- $rtrim
- $trim
- $regexFind  将正则表达式应用于字符串, 并返回第一个匹配子字符串的信息.
- $regexFindAll 将正则表达式应用于字符串, 并返回有关所有匹配字符串的信息
- $regexMatch 将正则表达式应用于字符串并返回一个布尔值, 表示是否已找到匹配项

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

- $replaceOne 替换给定输入中匹配字符串的第一个实例
- $replaceAll 替换给定输入中匹配字符串的所有实例
- $split  根据给定字符串拆分子字符串并返回子字符串数组
- $toLower  将字符串转换成小写
- $toUpper  将字符串转换成大写

```ts
{ $split: [ "June-15-2013", "-" ] } // [ "June", "15", "2013" ]
```

#### 类型表达式操作符

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
