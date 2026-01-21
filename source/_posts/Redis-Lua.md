---
title: Redis-Lua
date: 2022-11-09 18:03:19
categories:
  - [server, Redis]
tags:
  - Redis
---


## Redis 函数

> Redis 7.0 支持

Redis 函数是临时脚本的进化步骤, 函数提供与脚本相同的核心功能但却是数据库的一流软件工件

Redis 将函数作为数据库的一个组成部分进行管理, 并通过数据持久化和复制确保它们的可用性. 由于函数是数据库的一部分, 因此在使用前声明, 因此应用程序不需要在运行时加载它们, 也不必冒事务中止的风险. 使用函数的应用程序仅依赖于它们的 API, 而不依赖于数据库中的嵌入式脚本逻辑

Redis 函数可以将 Lua 的所有可用功能用于临时脚本, 唯一例外的是 Redis Lua 脚本调试器

Redis 函数还通过启用代码共享来简化开发, 每个函数都属于一个库, 任何给定的库都可以包含多个函数, 库的内容是不可变的, 并且不允许对其功能进行选择性更新. 取而代之的是, 库作为一个整体进行更新, 在一个操作中将它们的所有功能一起更新. 这允许从同一库中的其他函数调用函数, 或者通过使用库内部方法中的公共代码在函数之间共享代码, 这些函数也可以采用语言本机参数

Redis 函数也被持久化到 AOF 文件中, 并从主服务器复制到副本服务器, 因此它们与数据一样可以持久化

Redis 函数的执行是原子的, 函数的执行在其整个时间内阻止所有服务器活动, 类似于事务的语义, 已执行函数的阻塞语义始终适用于所有连接的客户端, 因为运行一个函数会阻塞 Redis 服务器

- 函数都属于一个库, 任何给定的库都可以包含多个函数
- 库的内容是不可变的, 并且不允许选择性地更新其函数, 只能将库作为一个整体进行更新

### 函数命令

- FUNCTION help 显示 FUNCTION 的帮助信息

```bash
127.0.0.1:6379> FUNCTION help
 1) FUNCTION <subcommand> [<arg> [value] [opt] ...]. Subcommands are:
 2) LOAD [REPLACE] <FUNCTION CODE>
 3)     Create a new library with the given library name and code.
 4) DELETE <LIBRARY NAME>
 5)     Delete the given library.
 6) LIST [LIBRARYNAME PATTERN] [WITHCODE]
 7)     Return general information on all the libraries:
 8)     * Library name
 9)     * The engine used to run the Library
10)     * Library description
11)     * Functions list
12)     * Library code (if WITHCODE is given)
13)     It also possible to get only function that matches a pattern using LIBRARYNAME argument.
14) STATS
15)     Return information about the current function running:
16)     * Function name
17)     * Command used to run the function
18)     * Duration in MS that the function is running
19)     If no function is running, return nil
20)     In addition, returns a list of available engines.
21) KILL
22)     Kill the current running function.
23) FLUSH [ASYNC|SYNC]
24)     Delete all the libraries.
25)     When called without the optional mode argument, the behavior is determined by the
26)     lazyfree-lazy-user-flush configuration directive. Valid modes are:
27)     * ASYNC: Asynchronously flush the libraries.
28)     * SYNC: Synchronously flush the libraries.
29) DUMP
30)     Return a serialized payload representing the current libraries, can be restored using FUNCTION RESTORE command
31) RESTORE <PAYLOAD> [FLUSH|APPEND|REPLACE]
32)     Restore the libraries represented by the given payload, it is possible to give a restore policy to
33)     control how to handle existing libraries (default APPEND):
34)     * FLUSH: delete all existing libraries.
35)     * APPEND: appends the restored libraries to the existing libraries. On collision, abort.
36)     * REPLACE: appends the restored libraries to the existing libraries, On collision, replace the old
37)       libraries with the new libraries (notice that even on this option there is a chance of failure
38)       in case of functions name collision with another library).
39) HELP
40)     Prints this help.
```

<!-- more -->

- FUNCTION DELETE 删除指定的库
- FUNCTION LIST 查看所有库和函数

```bash
127.0.0.1:6379> FUNCTION LIST
1) 1) "library_name"
   2) "mylib"
   3) "engine"
   4) "LUA"
   5) "functions"
   6) 1) 1) "name"
         2) "knockknock"
         3) "description"
         4) (nil)
         5) "flags"
         6) (empty array)
```

- 加载库和函数

每个 Redis 函数都属于一个加载到 Redis 的库, 使用命令 `FUNCTION LOAD` 将库加载到数据库, 库必须以 shebang 语句开头 `#!<engine name> name=<library name>`

```bash
# 加载一个空库
127.0.0.1:6379> FUNCTION LOAD "#!lua name=mylib\n"
(error) ERR No functions registered
```

### 注册调用

<em id="redis.register_function"></em> <!-- markdownlint-disable-line -->

注册

- redis.register_function(name, callback, flags, description) 注册函数
  - name 注册的函数名
  - callback 注册的函数
  - flags
    - no-writes 标识脚本只能读取但不能写入
    - allow-oom 标识允许脚本在服务器内存不足(OOM)时执行
    - allow-stable
    - no-cluster 标识脚本在 Redis 集群模式下返回错误, 防止对集群中的节点执行脚本
    - allow-cross-slot-keys 允许脚本从多个 slot 访问密钥
  - description 函数描述

调用

- FCALL function numkeys [key [key ...]] [arg [arg ...]] 调用注册的函数
- FCALL_RO function numkeys [key [key ...]] [arg [arg ...]] 调用注册的只读函数

#### Redis 命令行注册调用

```bash
# 方式1
127.0.0.1:6379> FUNCTION LOAD "#!lua name=mylib\nredis.register_function{function_name='noop', callback=function() end, flags={ 'no-writes' }, description='Does nothing'}"
# 方式2
127.0.0.1:6379> FUNCTION LOAD "#!lua name=mylib\nredis.register_function('knockknock', function() return 'Who\\'s there?' end)"
"mylib"
# 调用
127.0.0.1:6379> FCALL knockknock 0
"Who's there?"
```

#### Lua 脚本文件注册调用

```lua
#!lua name=mylib
--方式1
--[[redis.register_function{
  function_name='knockknock',
  callback=function() return 'Who\'s there?' end,
  flags={ },
  description='Does nothing'
}]]--
--方式2
--[[redis.register_function(
   'knockknock',
   function() return 'Who\'s there?' end
)]]--
local function knockknock()
   return 'Who\'s there?'
end

local function my_hset(keys, args)
   local key = keys[1]
   -- 调用 redis 命令 TIME 获取当前时间戳
   local time = redis.call('TIME')[1]
   return redis.call('HSET',key, '_last_modified_', time, unpack(args))
end

local function my_hgetall(keys, args)
   -- 使用 resp3 协议进行请求应答
   redis.setresp(3)
   local key = keys[1]
   local res = redis.call('HGETALL', key)
   res['map']['_last_modified_'] = nil
   return res
end

redis.register_function('knockknock', knockknock)
redis.register_function('my_hset', my_hset)
redis.register_function('my_hgetall', my_hgetall)

-- 注册 FCALL_RO 执行的函数
redis.register_function{
   function_name='my_hgetall_ro',
   callback=my_hgetall,
   flags={'no-writes'}
   description='read-only hash getall command'
}
```

```bash
[root@centos7 workspace]# cat mylib.lua | redis-cli -x FUNCTION LOAD REPLACE
"mylib"
# 调用注册函数
127.0.0.1:6379> FCALL my_hset 1 hash:zhang name "zhangsan" age 18 addr "beijing"
(integer) 4
127.0.0.1:6379> KEYS *
1) "hash:zhang"
2) "bit:zhang"
3) "xiaoming"
4) "name"

127.0.0.1:6379> FCALL my_hgetall 1 hash:zhang
1) "age"
2) "18"
3) "addr"
4) "beijing"
5) "name"
6) "zhangsan"

# FCALL 调用只读函数 my_hgetall_ro
127.0.0.1:6379> FCALL my_hgetall_ro 1 hash:zhang
1) "age"
2) "18"
3) "addr"
4) "beijing"
5) "name"
6) "zhangsan"

# FCALL_RO 调用普通函数 my_hgetall
127.0.0.1:6379> FCALL_RO my_hgetall 1 hash:zhang
(error) ERR Can not execute a script with write flag using *_ro command.
# FCALL_RO 调用只读函数 my_hgetall_ro
127.0.0.1:6379> FCALL_RO my_hgetall_ro 1 hash:zhang
1) "age"
2) "18"
3) "addr"
4) "beijing"
5) "name"
6) "zhangsan"
```

## Lua 脚本

Redis 允许在服务器上上传和执行 Lua 脚本, 脚本可以采用编程控制结构并在执行时使用大部分命令来访问数据库, 因为脚本在服务器中执行, 所以从脚本中读取和写入数据非常高效.

Redis 保证脚本的原子执行, 在执行脚本时, 所有服务器活动在其整个运行期间都被阻止.

Lua 允许在 Redis 中运行部分应用程序逻辑, 这样的脚本可以跨多个键执行条件更新, 可能以原子方式组合几种不同的数据类型

Lua 脚本由嵌入式执行引擎在 Redis 中执行, 尽管服务器执行它们, 但 EVAL 脚本被视为客户端应用程序的一部分, 这就是它们没有命名、版本化或持久化的原因. 因此, 如果所有脚本丢失, 应用程序可能需要随时重新加载

- 不能在 lua 脚本中声明全局变量, 只能使用 local 声明局部变量, 否则将执行报错

```bash
127.0.0.1:6379> EVAL "name=10; return name" 0   # 声明全局变量执行报错
(error) ERR user_script:1: Attempt to modify a readonly table script: be32aff3876f2170bc8f0e19998edc94924341e1, on @user_script:1.
127.0.0.1:6379> EVAL "local name=10; return name" 0   # local 声明局部变量正常执行
(integer) 10
```

### 脚本命令

> **脚本参数化** 为了确保在独立部署和集群部署中正确执行脚本, 脚本访问的所有键名都必须作为输入键参数显式提供

- EVAL script numkeys key [key ...] arg [arg ...] 执行 Lua 脚本
  - script 要执行的脚本语句
  - numkeys 指定后续的参数有几个 key
  - key 要操作的键的数量, 在 Lua 脚本中通过 `KEYS[1]`, `KEYS[2]` 获取
  - arg 参数, 在 Lua 脚本中通过 `ARGV[1]`, `ARGV[2]` 获取
- EVALSHA sha1 numkeys key [key ...] arg [arg ...] 使用缓存 Lua 脚本的 sha 执行脚本(SCRIPT LOAD 命令缓存脚本)

- EVAL_RO script numkeys [key [key ...]] [arg [arg ...]] 只读版本的 EVAL 命令, Redis 7.0 支持
- EVALSHA_RO sha1 numkeys [key [key ...]] [arg [arg ...]] 只读版本的 EVALSHA 命令, Redis 7.0 支持

```bash
127.0.0.1:6379> EVAL "return 10" 0
(integer) 10
127.0.0.1:6379> EVAL "return ARGV[1]" 0 100
"100"
127.0.0.1:6379> EVAL "return {ARGV[1], ARGV[2]}" 0 100 101
1) "100"
2) "101"
127.0.0.1:6379> EVAL "return {KEYS[1], KEYS[2], ARGV[1], ARGV[2], ARGV[3]}" 2 name age v1 v2
1) "name"
2) "age"
3) "v1"
4) "v2"
127.0.0.1:6379> EVAL "return {1, 2, { 3, 'hello world' } }" 0
1) (integer) 1
2) (integer) 2
3) 1) (integer) 3
   2) "hello world"
```

每次执行脚本都需要重新加载一遍脚本代码, 浪费资源, 使用 [脚本缓存](#scriptcache)

### lua 脚本中的 redis 实例

redis 单例实例, 使脚本能够与运行它的 Redis 服务器进行交互

全局变量

- KEYS 获取脚本声明的键参数, 下标从 1 开始
- ARGV 获取脚本声明的键参数剩余的参数, 下标从 1 开始

<em id="redis.call"></em> <!-- markdownlint-disable-line -->

- redis.call(command [, arg...]) 执行 redis 命令并返回结果, 如果遇到错误时直接返回给客户端
- redis.pcall(command [, arg...]) 执行 redis 命令并返回结果, 如果遇到错误时将返回给脚本的执行上下文

```bash
127.0.0.1:6379> GET name
"hello world"
127.0.0.1:6379> EVAL "return redis.call('SET', KEYS[1], ARGV[1])" 1 name "hello redis"
OK
127.0.0.1:6379> GET name
"hello redis"

127.0.0.1:6379> get user:100:score
"60"
127.0.0.1:6379> EVAL "if redis.call('EXISTS', KEYS[1]) == 1 then return redis.call('INCRBY', KEYS[1], ARGV[1]) else return nil end" 1 user:100:score 10
(integer) 70
127.0.0.1:6379> get user:100:score
"70"
```

分布式接口限流，每秒钟限制接口的最大请求次数

```lua
local key = KEYS[1]
local limit = tonumber(ARGV[1])
local expire_time = tonumber(ARGV[2])

local current = redis.call('GET', key)

if current and tonumber(current) >= limit then
  -- 超出限制
  return 0
end

current = redis.call('INCR', key)

if tonumber(current) == 1 then
  -- 第一次访问, 设置过期时间
  redis.call('EXPIRE', key, expire_time)
end

-- 允许访问
return 1
```

```bash
# key: rate:limit:192.168.1.1
# limit: 5 次
# expire: 100 秒
[root@centos7 workspace]# cat rate_limit.lua | redis-cli -x SCRIPT LOAD
"688faba01b26f6e54969ba7b9d0ce340b4c298c0"
127.0.0.1:6379> SCRIPT EXISTS 688faba01b26f6e54969ba7b9d0ce340b4c298c0
1) (integer) 1
127.0.0.1:6379> EVALSHA 688faba01b26f6e54969ba7b9d0ce340b4c298c0 1 rate:limit:192.168.1.1 5 100
(integer) 1
127.0.0.1:6379> EVALSHA 688faba01b26f6e54969ba7b9d0ce340b4c298c0 1 rate:limit:192.168.1.1 5 100
(integer) 1
127.0.0.1:6379> EVALSHA 688faba01b26f6e54969ba7b9d0ce340b4c298c0 1 rate:limit:192.168.1.1 5 100
(integer) 1
127.0.0.1:6379> EVALSHA 688faba01b26f6e54969ba7b9d0ce340b4c298c0 1 rate:limit:192.168.1.1 5 100
(integer) 1
127.0.0.1:6379> EVALSHA 688faba01b26f6e54969ba7b9d0ce340b4c298c0 1 rate:limit:192.168.1.1 5 100
(integer) 1
127.0.0.1:6379> EVALSHA 688faba01b26f6e54969ba7b9d0ce340b4c298c0 1 rate:limit:192.168.1.1 5 100
(integer) 0
127.0.0.1:6379> GET rate:limit:192.168.1.1
"5"
127.0.0.1:6379> TTL rate:limit:192.168.1.1
(integer) 66
```

- redis.error_reply(x) 辅助函数, 返回一个错误信息
- redis.status_reply(x) 辅助函数, 可以修改 Redis 命令的默认返回值 OK

```bash
# 返回错误信息
127.0.0.1:6379> EVAL "return redis.error_reply('ERR This is a special error')" 0
(error) ERR This is a special error

# 修改默认返回值
127.0.0.1:6379> EVAL "return { ok = 'TICK' }" 0
"TICK"
127.0.0.1:6379> EVAL "return redis.status_reply('TOCK')" 0
"TOCK"
```

- redis.sha1hex(x) 返回单个字符串参数的 SHA1 十六进制摘要信息
- redis.log(level, message) 写入 Redis 服务器日志
  - redis.LOG_DEBUG 日志级别
  - redis.LOG_VERBOSE 日志级别
  - redis.LOG_NOTICE 日志级别
  - redis.LOG_WARNING 日志级别

```bash
127.0.0.1:6379> EVAL "return redis.sha1hex('')" 0
"da39a3ee5e6b40d3255bfef95601890afd80709"
127.0.0.1:6379> EVAL "return redis.log(redis.LOG_WARNING, 'Something is terribly wrong')" 0
```

- redis.setresp(x) 设置执行脚本和服务器之间的请求应答协议, 默认值 2. Redis 6.0 支持
- redis.set_repl(x)
- redis.replicate_commands()
- redis.breakpoint() 在使用 Redis Lua 调试器时触发断点
- redis.debug(x) 在 Redis Lua 调试器控制台中打印其参数
- redis.acl_check_cmd(command [,arg...]) 用于检查运行脚本的当前用户是否具有使用给定参数执行给定命令的 ACL 权限, 返回值布尔类型. Redis 7.0 支持
- [redis.register_function(name, callback, flags, description)](#redis.register_function) Redis 7.0 支持

- redis.REDIS_VERSION 以字符串形式返回当前 Redis 服务器版本, 格式 MM.mm.PP. Redis 7.0 支持
- redis.REDIS_VERSION_NUM 以数字形式返回当前 Redis 服务器版本, 格式为十进制值. Redis 7.0 支持

```bash
127.0.0.1:6379> EVAL "return redis.REDIS_VERSION" 0
"7.0.5"
127.0.0.1:6379> EVAL "return redis.REDIS_VERSION_NUM" 0
(integer) 458757
```

### **脚本缓存** <em id="scriptcache"></em> <!-- markdownlint-disable-line -->

存储在服务器的脚本专用缓存中, 缓存内容由脚本的 SHA1 摘要作为缓存中的唯一标识

```bash
127.0.0.1:6379> SCRIPT help # 脚本帮助命令
 1) SCRIPT <subcommand> [<arg> [value] [opt] ...]. Subcommands are:
 2) DEBUG (YES|SYNC|NO)
 3)     Set the debug mode for subsequent scripts executed.
 4) EXISTS <sha1> [<sha1> ...]
 5)     Return information about the existence of the scripts in the script cache.
 6) FLUSH [ASYNC|SYNC]
 7)     Flush the Lua scripts cache. Very dangerous on replicas.
 8)     When called without the optional mode argument, the behavior is determined by the
 9)     lazyfree-lazy-user-flush configuration directive. Valid modes are:
10)     * ASYNC: Asynchronously flush the scripts cache.
11)     * SYNC: Synchronously flush the scripts cache.
12) KILL
13)     Kill the currently executing Lua script.
14) LOAD <script>
15)     Load a script into the scripts cache without executing it.
16) HELP
17)     Prints this help.
```

- SCRIPT FLUSH 从脚本缓存中移除所有脚本, 返回 ok
- SCRIPT KILL 杀死系统当前正在运行的 Lua 脚本(又名慢脚本)
- SCRIPT DEBUG 设置脚本内执行时的模式
- SCRIPT LOAD \<script\> 将脚本加载到服务器缓存中, 并不立即执行

```bash
# 添加 Lua 缓存脚本
127.0.0.1:6379> SCRIPT LOAD "return redis.call('GET', KEYS[1])"
"d3c21d0c2b9ca22f82737626a27bcaf5d288f99f"
# 使用 EVALSHA 执行缓存脚本
127.0.0.1:6379> EVALSHA d3c21d0c2b9ca22f82737626a27bcaf5d288f99f 1 name
"hello redis"
```

- SCRIPT EXISTS \<script\> [script ...] 查看缓存中是否存在 sha 对应的脚本, 1 表示存在, 0 表示不存在

```bash
127.0.0.1:6379> SCRIPT EXISTS d3c21d0c2b9ca22f82737626a27bcaf5d288f99f
1) (integer) 1
127.0.0.1:6379> SCRIPT EXISTS d3c21d0c2b9ca22f82737626a27bcaf5d288f99g
1) (integer) 0
```

### 脚本复制

一般在集群部署环境下, Redis 确保脚本执行的所有写操作也被发送到副本以保持一致性, 脚本复制有两种概念

- 逐字复制: master 将脚本的源代码发送到 slave, 然后 slave 执行脚本并写入效果.
  - 在短脚本生成许多命令的情况下, 可以节省资源, 但意味着 slave 会重做 master 完成的相同工作而浪费资源
- 效果复制: 仅复制脚本的数据修改命令, slave 然后执行命令而不执行任何脚本, 从 redis 5.0 开始为默认模式

脚本效果复制 —— 复制命令

在这种模式下，在执行 Lua 脚本的同时, Redis 会收集 Lua 脚本引擎执行的所有实际修改数据集的命令, 当脚本执行完成时, 脚本生成的命令序列被包装到一个 **事务** 中并发送到副本和 AOF

### Lua API

- 使用未声明为本地的变量和函数会引起 Redis 的报错
- 沙盒执行上下文不支持使用导入的 Lua 模块

### 数据类型转换

#### RESP2

- RESP2 -> Lua
  - RESP2 整数 -> Lua 数
  - RESP2 批量字符串 -> Lua 字符串
  - RESP2 数组 -> Lua 表(可能嵌套额其他 Redis 数据类型)
  - RESP2 状态 -> 包含状态字符串的单个 ok 字段的 Lua 表
  - RESP2 错误 -> 包含错误字符串的单个 err 字段的 Lua 表
  - RESP 空批量|空多批量 -> Lua false 布尔类型
- Lua -> RESP2
  - Lua 数字 -> RESP2 整数(数字转为整数, 舍去小数部分)
  - Lua 字符串 -> RESP 批量字符串
  - Lua 表(索引, 非关联数组) -> RESP2 数组(在表中遇到第一个 nil 时截断)
  - 带有单个 ok 字段的 Lua 表 -> RESP2 状态
  - 带有单个 err 字段的 Lua 表 -> RESP2 错误
  - Lua false 布尔类型 -> RESP2 空批量
  - Lua true 布尔类型 -> RESP2 整数 1

```bash
127.0.0.1:6379> EVAL "return {1, 2, {3, 'hello world'}, 'bar'}" 0
1) (integer) 1
2) (integer) 2
3) 1) (integer) 3
   2) "hello world"
4) "bar"
# 忽略表中的 键、数值的小数部分，nil 处截断
127.0.0.1:6379> EVAL "return {1, 2, 3.33, somekey = 'somevalue', 'foo', nil, 'bar'}" 0
1) (integer) 1
2) (integer) 2
3) (integer) 3
4) "foo"
```

#### RESP3

> 一旦 Redis 的回复采用 RESP3 协议, 所有 RESP2 到 Lua 的转换规则都适用, 并添加以下内容

- RESP3 -> Lua
  - RESP3 map -> 带有单个映射字段的 Lua 表, 其中包含表示映射字段和值的 Lua 表
  - RESP3 set -> 具有单个集合字段的 Lua 表
  - RESP3 null -> Lua nil
  - RESP3 true -> Lua true 布尔类型
  - RESP3 false -> Lua false 布尔类型
  - RESP3 浮点数 -> 带有一个浮点数字段的 Lua 表
  - RESP3 大数字 -> 带有单个大数字字段的 Lua 表. Redis 7.0 支持
  - RESP3 逐句逐字字符串 -> Lua 表, 其中包含单个 verbatim_string 字段的 Lua 表, 其中包含两个字段 string 和 format,分别表示 verbatim string 和它的格式. Redis 7.0 支持
- Lua -> RESP3
  - Lua Boolean -> RESP3 Boolean
  - 将单个映射字段设置为关联 Lua 表的 Lua 表 -> RESP3 map
  - 将单个集合字段设置为关联 Lua 表的 Lua 表 -> RESP3 set, 值可以为任何值, 都会被丢弃
  - 带有单个浮点数字段的 Lua 表到关联的 Lua 表 -> RESP3 浮点数
  - Lua nil -> RESP3 null

### 外部库

#### struct

- struct.pack(x) 返回一个结构编码的字符串, 接收一个结构格式字符串作为第一个参数, 后面是要编码的值

```bash
127.0.0.1:6379> EVAL "return struct.pack('bb', 1, 2)" 0
"\x01\x02"
127.0.0.1:6379> EVAL "return struct.pack('BB', 1, 2)" 0
"\x01\x02"
127.0.0.1:6379> EVAL "return struct.pack('B', 1, 2)" 0
"\x01"
127.0.0.1:6379> EVAL "return struct.pack('xB', 1, 2)" 0
"\x00\x01"
127.0.0.1:6379> EVAL "return struct.pack('xBx', 1, 2)" 0
"\x00\x01\x00"
127.0.0.1:6379> EVAL "return struct.pack('xBxx', 1, 2)" 0
"\x00\x01\x00\x00"
127.0.0.1:6379> EVAL "return struct.pack('xBxxH', 1, 2)" 0
"\x00\x01\x00\x00\x02\x00"
127.0.0.1:6379> EVAL "return struct.pack('BxxH', 1, 2)" 0
"\x01\x00\x00\x02\x00"
127.0.0.1:6379> EVAL "return struct.pack('Bxxh', 1, 2)" 0
"\x01\x00\x00\x02\x00"
127.0.0.1:6379> EVAL "return struct.pack('BxxB', 1, 2)" 0
"\x01\x00\x00\x02"
127.0.0.1:6379> EVAL "return struct.pack('Bxxl', 1, 2)" 0
"\x01\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00"
```

- struct.unpack(x) 返回结构的解码值, 接收一个结构格式字符串作为第一个参数, 然后是编码结构的字符串

```bash
127.0.0.1:6379> EVAL "return {struct.unpack('BxxH', ARGV[1])}" 0 "\x01\x00\x00\x02\x00"
1) (integer) 1
2) (integer) 2
3) (integer) 6
127.0.0.1:6379> EVAL "return {struct.unpack('BB', ARGV[1])}" 0 "\x01\x02"
1) (integer) 1
2) (integer) 2
3) (integer)
```

- struct.size(x) 返回结构的大小(以字节为单位), 接收结构格式字符串作为唯一参数

```bash
127.0.0.1:6379> EVAL "return struct.size('b')" 0
(integer) 1
127.0.0.1:6379> EVAL "return struct.size('B')" 0
(integer) 1
127.0.0.1:6379> EVAL "return struct.size('h')" 0
(integer) 2
127.0.0.1:6379> EVAL "return struct.size('H')" 0
(integer) 2
127.0.0.1:6379> EVAL "return struct.size('l')" 0
(integer) 8
127.0.0.1:6379> EVAL "return struct.size('L')" 0
(integer) 8
```

##### 结构格式

- \> 大端
- < 小端
- ![num] 结盟
- x 填充
- b/B 有/无符号字节
- h/H 有/无符号短
- l/L 有/无符号长
- T 大小
- i/In 大小为 n 的有/无符号整数(默认为 int 的大小)
- cn n 个字符的序列, 打包时, n ==0 表示整个字符串, 解包时, n == 0 表示使用先前读取的数字作为字符串的长度
- s 零终止字符串
- f float
- d double
- (space) 忽略

#### cjson

cjson 库提供了来自 Lua 的快速 JSON 编码和解码

- cjson.encode(x) 返回作为其参数提供的 Lua 数据类型的 JSON 编码字符串
- cjson.decode(x) 从作为其参数提供的 JSON 编码字符串返回 Lua 数据类型

```bash
127.0.0.1:6379> EVAL "return cjson.encode({ 1, 2, 'foo', 'bar' })" 0
"[1,2,\"foo\",\"bar\"]"
127.0.0.1:6379> EVAL "return cjson.encode({ 1, 2, 3.33, 'foo', 'bar' })" 0
"[1,2,3.33,\"foo\",\"bar\"]"
127.0.0.1:6379> EVAL "return cjson.encode({ ['foo'] = 'bar' })" 0
"{\"foo\":\"bar\"}"
127.0.0.1:6379> EVAL "return cjson.encode({ ['foo'] = 'bar', ['fov'] = 'baz' })" 0
"{\"fov\":\"baz\",\"foo\":\"bar\"}"

127.0.0.1:6379> EVAL "return cjson.decode(ARGV[1])[4]" 0 "[1,2,3.33,\"foo\",\"bar\"]"
"foo"
127.0.0.1:6379> EVAL "return cjson.decode(ARGV[1])['fov']" 0 "{\"fov\":\"baz\",\"foo\":\"bar\"}"
"baz"
```

#### cmsgpack

cmsgpack 库提供了来自 Lua 的快速 MessagePack 编码和解码

- cmsgpack.pack(x) 返回作为参数给出的 Lua 数据类型的压缩字符串编码
- cmsgpack.unpack(x) 返回解码其输入字符串参数的解压缩值

```bash
127.0.0.1:6379> EVAL "return cmsgpack.pack({'foo', 'bar', 'baz', 'hello'})" 0
"\x94\xa3foo\xa3bar\xa3baz\xa5hello"
127.0.0.1:6379> EVAL "return cmsgpack.unpack(ARGV[1])" 0 "\x94\xa3foo\xa3bar\xa3baz\xa5hello"
1) "foo"
2) "bar"
3) "baz"
4) "hello"
```

#### bit

bit 提供对数字的按位运算

- bit.tobit(x)` 将数字格式化为位运算的数值范围并返回
- bit.tohex(x [, n]) 将第一个参数转换为十六进制并返回, 第二个参数的绝对值控制返回值的数量

```bash
127.0.0.1:6379> EVAL "return bit.tobit(1)" 0
(integer) 1

127.0.0.1:6379> EVAL "return bit.tohex(422342)" 0
"000671cd"
```

- bit.bnot(x) 返回其参数的按位非运算
- bit.bor(x1 [, x2...]) 返回其所有参数的按位或运算
- bit.band(x1 [, x2...]) 返回其所有参数的按位与运算
- bit.bxor(x1 [, x2...]) 返回其所有参数的按位异或运算

```bash
# 0000 1100 12
#         !
# 1111 0011 -13
127.0.0.1:6379> EVAL "return bit.bnot(12)" 0
(integer) -13
# 0010 0000 32
#         !
# 1101 1111 -33
127.0.0.1:6379> EVAL "return bit.bnot(32)" 0
(integer) -33

127.0.0.1:6379> EVAL "return bit.bor(1,2,4,8,16,32,64)" 0
(integer) 127

# 0100 1010 74
# 0000 1100 12
#         &
# 0000 1000 8
127.0.0.1:6379> EVAL "return bit.band(12, 74)" 0
(integer) 8

# 0100 1010 74
# 0000 1100 12
#         ^
# 0100 0110 70
127.0.0.1:6379> EVAL "return bit.bxor(12, 74)" 0
(integer) 70
```

- bit.lshift(x, n) 返回第一个参数按位左移 n 位的结果
- bit.rshift(x, n) 返回第一个参数按位右移 n 位的结果
- bit.arshift(x, n) 返回第一个参数按位**算术右移** n 位的结果, 不改变符号位的移位操作

```bash
127.0.0.1:6379> EVAL "return bit.lshift(1, 3)" 0
(integer) 8
127.0.0.1:6379> EVAL "return bit.lshift(2, 1)" 0
(integer) 4
127.0.0.1:6379> EVAL "return bit.lshift(3, 2)" 0
(integer) 12
127.0.0.1:6379> EVAL "return bit.rshift(1, 1)" 0
(integer) 0
127.0.0.1:6379> EVAL "return bit.rshift(2, 1)" 0
(integer) 1
127.0.0.1:6379> EVAL "return bit.rshift(3, 1)" 0
(integer) 1
127.0.0.1:6379> EVAL "return bit.arshift(10, 1)" 0
(integer) 5
127.0.0.1:6379> EVAL "return bit.arshift(128, 1)" 0
(integer) 64

127.0.0.1:6379> EVAL "return bit.rshift(-12, 1)" 0
(integer) 2147483642
127.0.0.1:6379> EVAL "return bit.arshift(-12, 1)" 0
(integer) -6
```

- bit.rol(x, n) 按第二个参数给定的位数返回其第一个参数的按位左旋转
- bit.ror(x, n) 按第二个参数给定的位数返回其第一个参数的按位右旋转

```bash
127.0.0.1:6379> EVAL "return bit.rol(12, 1)" 0
(integer) 24
127.0.0.1:6379> EVAL "return bit.rol(12, 2)" 0
(integer) 48
127.0.0.1:6379> EVAL "return bit.rol(12, 6)" 0
(integer) 768

127.0.0.1:6379> EVAL "return bit.ror(12, 1)" 0
(integer) 6
127.0.0.1:6379> EVAL "return bit.ror(12, 4)" 0
(integer) -1073741824
127.0.0.1:6379> EVAL "return bit.ror(12, 6)" 0
(integer) 805306368
```

- bit.bswap(x) 交换其参数的字节并返回它, 可用于将小端 32 位数字转换位大端 32 位数字, 反之亦然

```bash
127.0.0.1:6379> EVAL "return bit.bswap(1)" 0
(integer) 16777216
127.0.0.1:6379> EVAL "return bit.bswap(2)" 0
(integer) 33554432
127.0.0.1:6379> EVAL "return bit.bswap(12)" 0
(integer) 201326592
```
