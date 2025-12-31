---
title: Lua
date: 2022-11-28 16:45:19
categories:
  - [server, Lua]
tags:
  - Lua
---

Lua 是一门强大、快速、轻量的嵌入式动态类型脚本语言, 使用 ANSI C 语言编写并以源代码形式开放, 其设计目的是为了嵌入应用程序中, 从而为应用程序提供灵活的扩展和定制功能

## 数据类型

Lua 有八种基本数据类型: nil、boolean、number、string、function、userdata、thread、table

- nil 表示一个有意义的值不存在时的状态, nil 和 false 逻辑表达式中都表示假, 其他任何值都表示真
- userdata 表示任意存储在变量中的 C 数据, 完全用户数据: 指一块由 Lua 管理的内存对应的对象; 轻量用户数据: 指一个简单的 C 指针
- table 本质是一个关联数组, 数组的索引可以是数字、字符串或表类型, 下标默认从 1 开始, table 可以包含任何类型的值(nil 除外), 任何键的值若为 nil 就不会被记入表内, table 的创建通过 `构造表达式 {}` 完成

table、function、thread、userdata 在 Lua 中是引用类型, 对其操作都是针对引用而不是针对值的操作

```lua
print("数据类型: nil boolean number string userdata function trhead table")
print("nil 表示一个无效值(在条件表达式中相当于false)")
print("boolean 表示 true 和 false, 除了 nil 和 false 值表示为 false, 其他值(包括0)都为 true, ")
print("number 表示双精度类型的实浮点数, 数字字符串相加将转换成数字相加")
print("string 表示一对双引号或者单引号包含的内容, [[ 内容 ]] 表示块字符串, .. 字符串拼接, # 计算字符串或表的长度")
print("userdata 表示任意存储在变量中的 C 数据结构")
print("function 由 C 或 Lua 编写的函数")
print("thread 表示执行的独立线程, 用于执行协同程序")
print("table 其实是一个关联数组, 数组的索引可以是数字、字符串或者表类型, 下标默认从 1 开始, table 的创建通过'构造表达式'完成, 空表: {}")
print("---------------------------------------")

print("变量的三种类型: 全局变量, 局部变量(local 声明), 表中的域")
print("变量批量赋值时, 多余的变量会赋值为 nil, 多余的值会被忽略")
print("---------------------------------------")
```

<!-- more -->

### 表

table 其实是关联数组, 可以使用任意类型值(除了 nil)作为数组的索引, 不固定大小

```lua
local tb1 = {"banana", "orange", "apple", "grapes"}
local tb2 = {name = "zhangsan", age = 18, "beijing"}
print("排序前: for 遍历 tb1")
for key, val in pairs(tb1) do
    print(val) -- banana orange apple grapes
end
print("声明空表", tb1)
print("声明非空表", tb2)
print("排序后: for 遍历 tb1")
for key, val in pairs(tb1) do
    print(val) -- apple banana grapes hello world orange
end
print("for 遍历 tb2")
for key, val in pairs(tb2) do
    print(key, ":", val) -- 1:beijing  age:18  name:zhangsan
end
print("-------------------")
```

#### 元表

lua 中的每个值都可以有一个元表, 元表就是一个普通的 lua 表, 它用于定义原始值再特定操作下的行为

元表决定了一个对象在数学运算、位运算、比较、连接、取长度、调用、索引时的行为, 元表还可以定义一个函数, 当表对象或用户数据对象在垃圾回收时调用它

元表中的键对应着不同的事件名, 键关联的值被称为 元方法, 通过 `getmetatable` 方法获取任何值的元表, 通过 `setmetatable` 方法设置元表

lua 中不能改变 table 以外其他类型的值的元表, 如果需要使用 C API

- \_\_add '+' 操作, 如果任何值不是数值类型(包括不能转换数值的字符串)做加法, lua 就会尝试调用此方法, lua 查找两个操作数是否定义此元方法, 只要有一个操作数包含, 则将两个操作数作为参数传入元方法, 元方法的结果作为这个操作的结果, 如果找不到元方法, 则抛出一个错误
- \_\_sub, '-' 操作, 行为和 `add` 操作类似
- \_\_mul, '\*' 操作, 行为和 `add` 操作类似
- \_\_div, '/' 操作, 行为和 `add` 操作类似
- \_\_mod, '%' 操作, 行为和 `add` 操作类似
- \_\_pow, '^' 幂操作, 行为和 `add` 操作类似
- \_\_unm, '-' 取负操作, 行为和 `add` 操作类似
- \_\_idiv, '//' 向下取整除法, 行为和 `add` 操作类似
- \_\_band, '&' 按位与运算, 行为和 `add` 操作类似, 不同的是 lua 会在任何一个操作数无法转换为整数时尝试取元方法
- \_\_bor, '|' 按位或运算, 行为和 `band` 操作类似
- \_\_bxor, '~' 按位异或运算, 行为和 `band` 操作类似
- \_\_bnot, '!' 按位非运算, 行为和 `band` 操作类似
- \_\_shl, '<<' 左移操作, 行为和 `band` 操作类似
- \_\_shr, '>>' 右移操作, 行为和 `band` 操作类似
- \_\_concat, '..' 连接操作, 行为和 `add` 操作类似, 不同的是 lua 在任何操作数即不是字符串也不是数字(数字总能转换为对应的字符串)的情况下尝试取元方法
- \_\_len, '#' 取长度操作, 如果对象不是字符串, lua 尝试取元方法, 如果有元方法, 则调用并将对象以参数形式传入, 返回值作为结果, 如果对象是一张表且没有元方法, lua 使用表的取长度操作, 其他情况均抛出错误
- \_\_eq, '==' 操作, 行为和 `add` 操作类似, 不同的是 lua 仅在两个值都是 table 或都是完全用户数据, 且它们不是同一个对象时才尝试取元方法, 调用的结果总是会被转换为布尔值
- \_\_lt, '<' 操作, 行为和 `add` 操作类似, 不同的是 lua 仅在两个值不全为整数也不全为字符串时才尝试取元方法, 调用的结果总是会被转换为布尔值
- \_\_le, '<=' 操作, 和其他操作不同, 此元方法可能用到两个不同的事件, 首先查找两个操作数的 `__le` 元方法, 如果找不到则再次查找 `__lt` 元方法, 它会假设 a <= b 等价于 not(b < a), 调用的结果总是会被转换为布尔值

```lua
-- metatable.lua
myTable = {k1 = 1, k2 = 2}
newTable = {k1 = 3, k2 = 4}
setmetatable(myTable, {
    __add = function(t1, t2)
        print("__add was called...")
        for k, v in pairs(t1) do
                print(k, v)
        end
        for k, v in pairs(t2) do
                print(k, v)
        end
        return -1
    end,
    __sub = function(t1, t2)
        print("__sub was called...")
        print(t2.k1 - t1.k1, t2.k2 - t1.k2)
        return 1000
    end,
    __pow = function(t1, t2)
        print("__pow was called...")
        return 10 ^ 2
    end,
    __mod = function(t1, t2)
        print("__mod was called...")
        return 10 % 3
    end,
    __band = function(t1, t2)
        print("__band was called...")
        return 10 & 5
    end,
    __shl = function(t1, t2)
        print("__shl was called...")
        return t1.k2 << 1
    end,
    __lt = function(t1, t2)
        print("__lt was called...")
        return t1.k2 < t2.k2
    end
})
print(myTable + newTable) -- __add 操作
print(myTable - newTable) -- __sub 操作
print(myTable ^ newTable) -- __pow 操作
print(myTable % newTable) -- __mod 操作
print(myTable & newTable) -- __band 操作
print(myTable << newTable) -- __shl 操作
print(myTable < newTable) -- __lt 操作
[root@centos7 workspace]# lua metatable.lua
__add was called...	-- __add 操作	<!-- markdownlint-disable-line -->
k1      1
k2      2
k1      3
k2      4
-1
__sub was called...	-- __sub 操作	<!-- markdownlint-disable-line -->
2       2
1000
__pow was called...	-- __pow 操作 	<!-- markdownlint-disable-line -->
100.0
__mod was called...	-- __mod 操作	<!-- markdownlint-disable-line -->
1
__band was called...	-- __band 操作	<!-- markdownlint-disable-line -->
0
__shl was called...	-- __shl 操作	<!-- markdownlint-disable-line -->
4
__lt was called...	-- __lt 操作	<!-- markdownlint-disable-line -->
true
```

- \_\_tostring, 元方法用于修改表的输出行为(自定义输出内容)

```lua
local mtstringstr = [[
mtstring = setmetatable({ 10, 20, 30}, {
    __tostring = function (t)
        local sum = 0
        for k, v in pairs(t) do
            sum = sum + v
        end
        return "表中所有元素的和为 "..sum
    end
})
print(mtstring) -- 表中所有元素的和为 60
]]
print(mtstringstr)
mtstring = setmetatable({ 10, 20, 30}, {
    __tostring = function (t)
        for k, v in pairs(t) do
            sum = sum + v
        end
        return "表中所有元素的和为 "..sum
    end
})
print(mtstring) -- 表中所有元素的和为 60
```

- \_\_call, 函数调用操作 func(args), 当 lua 尝试调用一个非函数的值时会尝试取元方法, 如果存在元方法则调用该方法, func 作为第一个参数传入, 原来调用的参数一次排在后面

```lua
myTable = {k1 = 1, k2 = 2, 5}
newTable = {k1 = 3, k2 = 4}
setmetatable(myTable, {
    __call = function(t1, t2)
        local num = 0
        for k, v in pairs(t1) do
            num = num + v
        end
        for k, v in pairs(t2) do
            num = num + v
        end
        return num
    end
})
print(myTable(newTable))
[root@centos7 workspace]# lua metatable.lua
15
```

- \_\_index, table[key] 查找操作, 当 table 不是表或者 table 中不存在 key 这个键时, lua 会尝试取元方法
  - 如果 \_\_index 键包含一个 table 时, lua 则会在这个 table 中查找相应的 key
  - 如果 \_\_index 键包含一个函数时, lua 则会调用这个函数, table 和 key 作为参数传递给函数并接收函数的返回值作为结果

查找顺序:

1. 在 table 中查找, 如果找到则返回该元素, 否则继续
2. 判断该 table 是否有元表, 如果没有则返回 nil, 否则继续
3. 判断元表是否有 \_\_index 键, 如果没有则返回 nil, 如果 \_\_index 键包含一个 table, 则重复 1. 2. 3, 如果 \_\_index 键包含一个函数, 则返回调用该函数的返回值

```lua
print("-------__index是table---------")
t = setmetatable({}, {__index = { name = "hello world" }})
print("t = setmetatable({}, {__index = { name = \"hello world\" }})")
print("print(t.name)", t.name) -- hello world
print("print(t.foo)", t.foo) -- nil
print("-------__index是函数---------")
mytable = setmetatable({foo = "bar"},{
    __index = function(t, k)
        if k == "baz" then
            return "baz = baz"
        else
            return nil
        end
    end
})
local mytableStr = [[
mytable = setmetatable({foo = "bar"},{
    __index = function(t, k)
        if k == "baz" then
            return "baz = baz"
        else
            return nil
        end
    end
})]]
print(mytableStr)
print("print(mytable.foo, mytable.baz)", mytable.foo, mytable.baz) -- bar     baz = baz
```

- \_\_newindex, table[key] = value 索引赋值操作, 发生在 table 不是表或者 table 中不存在 key 这个键时, lua 会尝试取元方法
  - 如果 \_\_newindex 键包含一个 table 时, lua 则会对这个 table 做索引赋值操作, 索引过程有可能会引发另一次元方法
  - 如果 \_\_newindex 键包含一个 函数时, lua 会调用这个函数而不进行赋值操作, table、key、value 将作为函数的参数传入

```lua
print("-------__newindex是table---------")
local mtnewmtstr = [[
mtnewmt = {}
mtnew = setmetatable({ name = "hello world" }, { __newindex = mtnewmt})
print(mtnew.name) -- hello world
mtnew.name = "hello lua"
print(mtnew.name, mtnewmt.name) -- hello lua  nil
mtnew.addr = "beijing"
print(mtnew.addr, mtnewmt.addr) -- nil  beijing
]]
print(mtnewmtstr)
mtnewmt = {}
mtnew = setmetatable({ name = "hello world" }, { __newindex = mtnewmt})
print(mtnew.name) -- hello world
mtnew.name = "hello lua"
print(mtnew.name, mtnewmt.name) -- hello lua  nil
mtnew.addr = "beijing"
print(mtnew.addr, mtnewmt.addr) -- nil beijing
print("-------__newindex是函数---------")
local mtnewmt2str = [[
mtnewmt2 = setmetatable({ name = "hello world" }, {
    __newindex = function(t, k, v)
        rawset(t, k, "gg_".."\""..v.."\"".."_gg")
    end
})
mtnewmt2.age = 18
mtnewmt2.addr = "beijing"
print(mtnewmt2.name, mtnewmt2.age, mtnewmt2.addr) -- hello world     gg_"18"_gg      gg_"beijing"_gg
]]
mtnewmt2 = setmetatable({ name = "hello world" }, {
    __newindex = function(t, k, v)
        rawset(t, k, "gg_".."\""..v.."\"".."_gg")
    end
})
mtnewmt2.age = 18
mtnewmt2.addr = "beijing"
print(mtnewmt2.name, mtnewmt2.age, mtnewmt2.addr) -- hello world     gg_"18"_gg      gg_"beijing"_gg
```

- \_\_gc, 垃圾收集元方法, 当垃圾收集循环时触发

```lua
-- metatable.lua
myTable = {k1 = 1, k2 = 2, 5}
setmetatable(myTable, {
    __gc = function()
        print("__gc was called...")
    end
})
myTable = nil
[root@centos7 workspace]# lua metatable.lua
__gc was called...
```

### 弱表

弱表指内部元素为 `弱引用` 的表, 垃圾收集器会忽略弱引用计数, 如果一个对象只被弱引用引用时, 垃圾收集器就会回收这个对象
一张弱表可以有弱键或者弱值, 也可以键值都是弱引用, 仅包含弱键的表允许垃圾收集器回收它的键, 但会阻止对值所指的对象回收, 若一张表的键值均为弱引用, 那么收集器可以回收其中的任意键和值
在任何情况下, 只要键或值的任意一项被回收, 相关联的键值对都会从表中移除
一张表的元表中的 \_\_mode 域控制着这张表的弱属性

- 当 \_\_mode 域是一个包含字符 k 的字符串时, 这张表的所有键都为弱引用
- 当 \_\_mode 域是一个包含字符 v 的字符串时, 这张表的所有值都为弱引用

属性为 `弱键强值` 的表也被称为 `暂时表`, 对于一张暂时表, 它的值是否可达仅取决于其对应键是否可达
对一张表的弱属性的修改仅在下次手机循环才生效, 只有那些有显式构造过程的对象才会从弱表中移除, 值, 例如数字和轻量 C 函数, 不受垃圾收集器管辖, 因此不会从弱表中移除(除非它们的关联项被回收)

## 协程 <em id="coroutine"></em> <!-- markdownlint-disable-line -->

lua 支持协程(协同式多线程), 一个协程在 lua 中代表了一段独立的执行线程, 协程拥有独立的堆栈, 独立的局部变量, 同时又与其他协程共享全局变量和其他大部分东西,

协程和线程的主要区别：

一个具有多个线程的程序可以同时运行多个线程, 协程却需要彼此协作的运行, 在任一指定时刻只有一个协程在运行, 并且这个正在运行的协程只有在明确的被要求挂起时才会被挂起

协程的运行可能被两种方式终止, 正常途径是主函数返回(显式返回或者执行完最后一条指令), 非正常途径是发生了一个未捕获的错误, 对于正常结束, `coroutine.resume()` 将返回 true 和协程主函数的返回值, 当错误发生时, `coroutine.resume()` 将返回 false 和错误信息

- coroutine.create(func) 创建 coroutine 并返回协程句柄, 当和 resume 配合使用时就唤醒函数调用
- coroutine.resume(co [, val1, ...]) 重启 coroutine 并将参数传入, 协程正常运行返回 true 和 传给 yield 的所有值(当协程让出)或者主体函数的所有返回值(当协程中止), 有错误发生时返回 false 和错误信息
- coroutine.isyieldable() 判断正在运行的协程是否可以让出, 可以返回 1, 否则返回 0
- coroutine.yield(args) 挂起 coroutine, 如果有参数则将参数返回给调用线程
- coroutine.status(co) 查看 coroutine 的状态, 通常返回 dead, suspended, running
- coroutine.wrap(f) 创建 coroutine 并返回一个函数, 启动协程需要手动调用这个函数
- coroutine.running() 返回当前正在运行的 coroutine 和一个布尔值, 如果当前运行的协程是主线程, 布尔值为 true, 否则为 false

```lua
-- coroutine.lua
co = coroutine.create(
    function(i)
        print(i)
    end
)
coroutine.resume(co, 120) -- 120
print(coroutine.status(co)) -- dead
print(coroutine.running()) -- thread:0xef8018 true
print("---------------")
co = coroutine.wrap(
    function(i)
        print(i)
    end
)
co(110) -- 110
print(coroutine.running()) -- thread:0xef8018 true
print("---------------")
[root@centos7 workspace]# lua coroutine.lua
120
dead
thread: 0xef8018        true
---------------
110
thread: 0xef8018        true
("---------------")
```

```lua
-- coroutine.lua
function foo (a)
    -- 4. 执行打印 foo 2
    print("foo", a)
    -- 5. 挂起协程返回结果 4
    return coroutine.yield(2*a)
end
co = coroutine.create(
    function (a,b)
        -- 2. 执行打印 co-body 1 10
        print("co-body", a, b)
        -- 3. 调用 foo 函数传入参数 2
        -- 8. 执行赋值操作将 resume 的参数 r 赋值给局部变量 r
        local r = foo(a+1)
        -- 9. 执行打印 co-body r
        print("co-body", r)
        -- 10. 挂起协程返回结果 11 -9
        -- 13. 执行赋值操作将 resume 的参数 x y 赋值给局部变量 r s
        local r, s = coroutine.yield(a+b, a-b)
        -- 14. 执行打印 co-body x y
        print("co-body", r, s)
        -- 15. 返回结果 10 end, 协程执行完成退出
        return b, "end"
    end
)
-- 1. 唤起协程传入参数 1 10  -- 6. 输出结果 main true 4
print("main", coroutine.resume(co, 1, 10))
-- 7. 唤起协程传入参数 r  -- 11. 输出结果 main true 11 -9
print("main", coroutine.resume(co, "r"))
-- 12. 唤起协程传入参数 x y  -- 16. 输出结果 main true 10 end
print("main", coroutine.resume(co, "x", "y"))
-- 17. 唤起协程传入参数 x y, 协程执行完毕输出结果 main false cannot resume dead coroutine
print("main", coroutine.resume(co, "x", "y"))
[root@centos7 workspace]# lua coroutine.lua
co-body 1       10
foo     2
main    true    4
co-body r
main    true    11      -9
co-body x       y
main    true    10      end
main    false   cannot resume dead coroutine
```

### 生产者与消费者

```lua
local newProducer
function producer()
    local i = 0
    while true do
        i = i + 1
        coroutine.yield(v) -- 发送数据后就挂起 coroutine
    end
end
function consumer()
    while true do
        -- 唤起 coroutine, 接收 coroutine 挂起时返回的结果
        local status, value = coroutine.resume(newProducer)
        print(value)
    end
end
newProducer = coroutine.create(producer)
consumer()
```

## 词法约定

### 关键字

and break do else elseif end false for function goto if in local nil not or repeat return then true until while

### 字符串

```lua
--这是单行注释, 不会被执行

--[[
这是多行注释, 不会被执行
这是多行注释, 不会被执行
这是多行注释, 不会被执行
这是多行注释, 不会被执行
]]--

print('单行注释: \n-- 单行注释内容')
print("多行注释: \n--[[ \n 多行注释内容\n 多行注释内容\n 多行注释内容\n]]--")
print("---------------------------------------")


html = [[
<html>
<head></head>
<body>
        <a href="//www.baidu.com">baidu</a>
</body>
</html>
]]
print("块字符串[[ ]]:\n", html)
print("字符串拼接 .. :", "hello".."world")
print("计算字符串的长度 # :", #html)

print("% 作为特殊字符的转义字符, 就像正则表达式中的转义字符 \\")
print("匹配模式: %a 表示任意字符, %c 表示任意控制字符, %d 表示任意数字, %l 表示任意小写字母, %u 表示任意大写字母, %s 表示任意空白字符, %w 表示任意字母/数字, %x 表示任意十六进制数, %p 表示任意标点, %z 表示任意代表0的字符")
print("当上述字符类用大写表示时, 表示与非此字符类的任意字符匹配, %S 表示任意非空白字符")
```

```lua
a = 'alo\n123"' --  字符串中包含另一种字符串的引号不需要转义
print(a)
--[[
alo
123"
]]
a = "alo\n123\"" -- 字符串中包含相同的引号需要进行转义
print(a)
--[[
alo
123"
]]
a = '\97lo\10\04923"' -- 使用 \XX 的十六进制形式表示字符
print(a)
--[[
alo
123"
]]
a = '\u{3b1} \u{3b2} \u{3b3}' -- 使用 \u{XXX} 十六进制表示用 UTF-8 编码的 Unicode 字符
print(a) -- α β γ
-- 使用 [[ ]] 表示多行字符串
a = [[alo
123"]]
print(a)
--[[
alo
123"
]]
-- 使用任意数量的 = 间隔分隔表示对其它多行字符串的引用
a = [==[
alo
123"
]==]
print(a)
--[[
alo
123"
]]
```

### 变量

Lua 有三种变量: 全局变量, 局部变量和表的域
所有没有显式声明的局部变量名都被当做全局变量, 在变量的首次赋值之前, 默认值都为 nil
变量的作用范围开始于声明它们之后的第一个语句段

```lua
x = 10 -- 全局变量
do
    local x = x -- 新的本地变量 x, 并赋值 10
    print(x) -- 10 第一层作用域的本地变量
    x = x + 1
    do
        local x = x + 1 -- 取最近的变量 x 的值加 1 并赋值给新的本地变量 x
        print(x) -- 12 第二层作用域的本地变量
    end
    print(x) -- 11 第一层作用域的本地变量
end
print(x) -- 10 全局变量
```

### 控制结构

- if, if exp then block {elseif exp then block} [else block] end
- while, while exp do block end
- for, for exp do block end
- repeat, repeat block until exp

### 运算符

- 算术运算符: + - \* / // % ^ -
- 关系运算符: == ~= < <= > >=
- 逻辑运算符: and or not
- 位运算符: & | ~ << >>
- 字符串拼接: ..
- 取长度操作符(元方法\_\_len): #

### 函数定义

- 冒号语法可以用来定义方法, 使函数有一个隐形的形参 self, 代表函数自己

```lua
print("面向对象: lua 使用 table 描述对象的属性, 使用 function 描述方法, 使用 table + function 模拟面向对象")
print("---------------------------------------")
-- 冒号语法可以用来定义方法, 使函数有一个隐形的形参 self, 代表函数自己
-- Meta class
Shape = {area = 0}
-- 基础类方法 new
function Shape:new (o,side)
  o = o or {}
  setmetatable(o, self)
  self.__index = self
  side = side or 0
  self.area = side*side;
  return o
end
-- 基础类方法 printArea
function Shape:printArea ()
  print("面积为 ",self.area)
end
-- 创建对象
myshape = Shape:new(nil,10)
myshape:printArea()
print("---------------")
Square = Shape:new()
-- 派生类方法 new
function Square:new (o,side)
  o = o or Shape:new(o,side)
  setmetatable(o, self)
  self.__index = self
  return o
end
-- 派生类方法 printArea
function Square:printArea ()
  print("正方形面积为 ",self.area)
end
-- 创建对象
mysquare = Square:new(nil,10)
mysquare:printArea()
print("---------------")
Rectangle = Shape:new()
-- 派生类方法 new
function Rectangle:new (o,length,breadth)
  o = o or Shape:new(o)
  setmetatable(o, self)
  self.__index = self
  self.area = length * breadth
  return o
end
-- 派生类方法 printArea
function Rectangle:printArea ()
  print("矩形面积为 ",self.area)
end
-- 创建对象
myrectangle = Rectangle:new(nil,10,20)
myrectangle:printArea()
print("---------------------------------------")
```

## C API 中的函数和类型

- lua_status(lua_State \*L) 返回线程 L 的状态, 正常状态为 0(LUA_OK), 当线程用 lua_resume 执行完毕并抛出了一个错误时, 状态值时错误码, 如果线程被挂起, 状态为 LUA_YIELD
- lua_version(lua_State \*L) 返回在 Lua 内核中保存的版本数字的地址

## 辅助库

辅助库提供了一些便捷函数, 方便在 C 中为 Lua 编程, 基础 API 提供了 C 和 Lua 交互用的主要函数, 而辅助库则为一些常见的任务提供了高阶函数
所有辅助库中的函数和类型都定义在头文件 lauxlib.h 中, 它们均带有前缀 luaL\_
辅助库中的所有的函数都基于基础 API 实现
一些辅助库函数会在内部使用一些额外的栈空间, 当辅助库使用的栈空间少于 5 个时, 它们不会取检查栈大小, 而是简单的假设栈够用
一些辅助库看中的函数用于检查 C 函数的参数, 因为错误信息格式化为指代参数

### 函数和类型

- luaL_addchar(luaL_Buffer \*B, char c) 向缓存 B 添加一个字节 c
- luaL_addlstring(luaL_Buffer *B, const char *s, size_t l) 向缓存 B 添加一个长度为 l 的字符串 s, 这个字符串可以包含零
- luaL_addstring(luaL_Buffer *B, const char *s) 向緩存 B 添加一个零结尾的字符串 s
- luaL_callmeta(lua_State *L, int obj, const char *e) 调用一个元方法, 如果在索引 obj 处的对象有元表, 且元表有域 e, 这个函数会以该对象为参数调用这个域, 这种情况下, 函数返回真并将调用返回值压栈, 如果这个位置没有元表, 或没有对应的元方法, 此函数返回假(并不会将任何东西压栈)
- luaL_dostring(lua_State *L, const char *str) 加载并运行指定的字符串(使用 luaL_loadstring 或者 luaL_pcall 定义), 如果没有错误返回假, 有错误返回真
- luaL_getmetatable(lua_State *L, const char *tname) 将注册表中 tname 对应的元表压栈, 如果没有 tname 对应的元表, 则将 nil 压栈并返回假
- luaL_len(lua_State \*L, int index) 以数字形式返回给定索引处值的 长度, 等价于在 lua 中使用 # 的操作, 如果结果不是一个整数, 则抛出一个错误
- luaL_loadstring(lua_State \*L, const char \*s) 将一个字符串加载为 lua 代码块, 这个函数使用 lua_load 加载一个零结尾的字符串 s, 返回值和 lua_load 相同

## 标准库

标准库提供了一些有用的函数, 它们都是直接用 C API 实现的, 其中一些函数提供了原本语言就有的服务(type/getmetatable), 另一些提供和 `外部` 打交道的服务(I/O)
还有些本可以用 lua 本身来实现, 但在 C 中实现可以满足关键点上的性能需求(例如 table.sort)
所有的库都是直接用 C API 实现的, 并以分离的 C 模块形式提供

### 基础库

- assert(v[,message]) 如果参数 v 的值为假(nil 或 false)就会调用 error, message 为错误对象, 否则返回所有的参数
- error (message [, level]) 终止正在执行的函数, 并返回 message 的内容作为错误信息, level 指示获取错误的位置: 1 默认, 为调用 error 的位置(文件+行), 2 指出调用 error 函数的函数, 0 不添加错误位置信息
- pcall(f [, arg1, ...]) 传入参数, 以保护模式调用函数 f

```lua
print("错误: 语法错误 和 运行时错误")
print("assert(arg1, arg2) 类型断言, 如果第一个参数为真, assert 不做任何处理, 否则将第二个作为错误信息输出:  assert(type(a) == 'number', 'a 不是一个数字')")
print("pcall() 保护模式调用, 接收一个函数和要传递给函数的参数, 以保护模式执行第一个参数, 可以捕获函数执行中的任何错误, 无错误返回 true, 有错误返回 false 和 errorinfo")
print("xpcall() 第一个参数和第二个后面的参数作用同 pcall, 第二个参数为一个错误处理函数, 当错误发生时, lua 会在调用栈展开前调用错误处理函数, debug.debug() 提示一个 lua 提示符, 让用户来检查错误的原因, debug.traceback() 根据调用栈来构建一个扩展的错误消息")
print("---------------------------------------")
```

- collectgarbage ([opt [, arg]]) 垃圾收集器的通用接口, opt 提供了一组不同的功能

```lua
print("垃圾回收: lua 采用了自动内存管理, collectgarbage([opt] [, arg])")
print("collectgarbage('collect') 做一次完整的垃圾回收循环")
print("collectgarbage('count') 以 K 字节为单位返回 lua 使用的总内存数")
print("collectgarbage('restart') 重启垃圾回收器的自动运行")
print("collectgarbage('setpause') 将 arg 设置为收集器的间歇率, 返回间歇率的前一个值")
print("collectgarbage('setstepmul') 返回步进倍率的前一个值")
print("collectgarbage('step') 单步运行垃圾收集器, 步长大小由 arg 决定")
print("collectgarbage('stop') 停止垃圾收集器的运行")
print("collectgarbage('isrunning') 返回表示收集器是否在工作的布尔值")
print("---------------------------------------")
```

- getmetatable(object) 返回 object 的元表, 如果不包含元表则返回 nil
- ipairs(t) 返回 3 个值(迭代函数, 表 t, 以及 0)
- pairs(t) 如果 t 有元方法 \_\_pairs, 以 t 为参数调用它并返回其返回的前 3 个值,否则, 返回 3 个值(迭代函数, 表 t, 以及 nil)

```lua
for i,v in ipairs(t) do -- 将迭代键值对(1, t[1]), (2, t[2]) ... 知道第一个空值
    body
end
for k,v in pairs(t) do
    body
end
```

- print(...) 接收任意数量的参数, 并将它们的值打印到 stdout
- rawequal(v1, v2) 在不触发任何元方法的情况检查 v1 和 v2 是否相等, 返回一个布尔值
- rawlen(v) 在不触任何元方法的情况下返回对象 v 的长度
- rawset(table, index, value) 在不触发任何元方法的情况将 table[index] 设置为 value, table 必须是一张表
- tonumber(e [, base]) 尝试将 e 转换为一个指定 base 进制的数字
- tostring(v) 将参数 v 转换为可阅读的字符串形式
- type(v) 返回指定参数的类型编码的字符串形式

### [协程库](#coroutine)

### 包管理库

模块: 封装公用的代码以 API 接口的形式在其他地方调用, 简单理解是将变量、常量、函数放在一个 table 里面，然后 return 返回

使用 require 方法加载模块, require(\"模块名\") 或者 require \"模块名\"

模块的加载机制: require 用于搜索 lua 文件的路径是存放在全局变量 package.path 中, 当 lua 启动后, 会以环境变量 LUA_PATH 的值来初始这个环境变量, 如果没有找到该环境变量, 则使用一个编译时定义的默认路径来初始化, 此环境变量也可以自定义设置, 在搜索过程中, 如果找到该文件, 则使用 pacakge.loadfile 来加载模块, 否则就去找 C 程序库, 搜索的文件路径是从全局变量 package.cpath 获取, 而这个变量则是通过环境变量 LUA_CPATH 来初始, 此时搜索的文件是 so 或 dll 类型的文件, 如果找到了则使用 package.loadlib 来加载

- require(modename) 加载一个模块
- package.config 描述一些为包管理准备的编译期配置信息的串
- package.cpath 模块在 C 加载器中加载时的搜索路径
- package.loaded 控制哪些模块已经被加载的表
- package.loadlib(libname, funcname)
- package.path 模块在 lua 加载器中加载时搜索路径
- package.preload 保存一些特殊模块的加载器
- package.searchers 控制如何加载模块的表
- package.searchpath(name, path [, sep [, rep]]) 在指定 path 中搜索指定的 name

```lua
local modulestr = [[
local tst = {}
local name = "hello world"
function tst.getName()
  return name
end
return tst

local tst = require("tst")
print(tst.getName())
]]
print(modulestr)
```

### 字符串控制

这个库提供了字符串处理的通用函数

- string.upper 转换大写", string.upper("AbCdE"))
- string.lower 转换小写", string.lower("AbCdE"))
- string.gsub 字符串查找替换", string.gsub("hello world", "l", 'r'))
- string.find 查找子串位置", string.find("hello lua user", "lua"))
- string.reverse 字符串反转", string.reverse("hello world"))
- string.format 根据字符串模板返回格式化的字符串", string.format("the value is %d", 4))
- string.char 返回数值表示的字符, byte 返回字符的数值表示", string.char(96, 99, 100), string.byte('ABCD'))
- string.len 返回字符串的长度", string.len("abc"))
- string.rep 返回字符串的 n 个拷贝", string.rep("abc", 3))
- string.gmatch 返回一个迭代器函数, 每次调用函数返回一个字符串中找到的下一个符合 pattern 模式的子串, 可以结合 for 循环查找")
- string.match 返回在字符串中查找符合匹配模式的第一个子串")
- string.sub 截取字符串", string.sub('hello world', 1, 6))

### 基础 UTF-8 支持

这个库提供了怼 UTF-8 编码的基础支持, 所有的函数都放在表 utf8 中, 此库不提供除编码处理之外的任何 unicode 支持

- utf8.char(...) 接收零个或多个整数, 将每个整数转换成对应的 UTF-8 字节序列, 并返回这些序列连接到一起的字符串
- utf8.codes(s) 返回一系列的值, 迭代出字符串 s 中所有的字符
- utf8.codepoint(s [, i [, j]]) 以整数形式返回 s 中从位置 i 到 j 间(包括两端啊)所有字符的编码, 默认 i 为 1, j 为 i
- uft8.len(s [, i [, j]]) 返回字符串 s 中从位置 i 到 j 间(包含两端) UTF-8 字符的个数, 默认 i 为 1, j 为 -1
- uft8.offset(s, n [, i]) 返回编码在 s 中的第 n 个字符的开始位置(按字节数)(从位置 i 开始统计), 如果指定的字符不在其中或在结束点之后, 函数返回 nil

### 表控制

- table.insert(list, [pos, ] value) 在 list 的位置 pos 处插入元素 value, 并向后移动元素, table.insert(tb1, 3, "hello world")
- table.pack(...) 返回用所有参数乘以键 1, 2, 等填充的新表, 并将 n 这个域设为参数的总数
- table.remove(list [, pos]) 移除 list 中 pos 位置的元素并返回移除的元素, pos 默认为 #list, table.remove(tb2, 2)
- table.sort(list [, comp]]) 对 list 进行排序, 如果提供了参数 comp, 则 comp 必须是一个可以接收两个列表内元素为参数的函数, table.sort(tb1)
- table.unpack(list [, i [, j]]) 返回 list 中的元素, 默认 i 为 1, j 为 #list
- table.concat 列出表中指定区间的所有元素, table.concat({"hello", "world", "lua", 2022}, "-")

### 数学函数

这个库提供了基本的数学函数, 所有函数都放在表 math 中, 注解有 integer/float 的函数会对整数参数返回整数结果, 对浮点(或混合)参数返回浮点结果, 圆整函数(math.ceil, math.floor, math.modf)的结果在整数范围内是返回整数, 否则返回浮点数

### 输入输出

I/O 库提供了两套不同风格的文件处理接口

- 简单模式(simple mode), 它提供设置默认输入文件及默认输出文件的操作, 所有的输入输出操作都针对这些默认文件
- 完全模式(complete mode), 当使用隐式文件句柄时, 所有的操作都是由表 io 提供, 若使用显式文件句柄, io.open 会返回一个文件句柄, 且所有的操作都由该文件句柄的方法来提供

表 io 中也提供了三个和 C 中含义相同的预定义文件句柄: io.stdin, io.stdout, 以及 io.stderr, I/O 库永远不会关闭这些文件

- io.close([file]) 等价于 file:close(), 不指定 file 时关闭默认输出文件
- io.flush() 等价于 io.output():flush()
- io.input([file]) 当文件名调用它时, (以文本模式)来打开该名字的文件, 并将文件句柄设为默认输入文件, 如果用文件句柄调用它时, 就简单的将该句柄设为默认输入文件, 如果调用时不传参数, 则返回当前的默认输入文件
- io.lines([filename ...]) 以读模式打开指定的文件名并返回一个迭代函数, 此迭代函数的工作方式和用一个已打开的文件去调用 file:lines(...) 得到的迭代器相同, 当迭代函数检测到文件结束, 它不返回值(让循环结束)并自动关闭文件

  调用 io.lines() (不传文件名) 等价于 io.input():lines('\*|'), 按行迭代标准输入文件, 在此情况下, 循环结束后它不会关闭文件

- io.open(filename [, mode]) 用字符串 mode 指定的模式打开一个文件并返回新的文件句柄, 当出错时, 返回 nil 和错误信息
  - r 读模式, 默认
  - w 写模式
  - a 追加模式
  - r+ 更新模式, 所有之前的数据都保留
  - w+ 更新模式, 所有之前的数据都删除
  - a+ 追加更新模式, 所有之前的数据都保留, 只允许在文件尾部写入
- io.output([file]) 类似于 io.input(), 不过都针对默认输出文件操作
- io.popen(prog [, mode]) 跟系统有关, 不是所有平台都提供, 用一个分离进程开启程序 prog, 返回的文件句柄可用于从这个程序中读取数据
- io.read(...) 等价于 io.input():read(...)
- io.tmpfile() 返回一个临时文件的句柄, 这个文件以更新模式打开, 在程序结束时自动删除
- io.type(obj) 检查 obj 是否是合法的文件句柄, 如果是返回 'file', 如果是关闭的文件句柄返回 'closed file', 不是则返回 nil
- io.write(...) 等价于 io.output():write(...)
- file:close() 关闭文件
- file:flush() 将写入的数据保存到 file 中
- file:lines(...) 返回一个迭代器函数, 每次调用迭代器时, 都从文件中按指定格式读取数据, 默认 '|'
- file:read(...) 按照指定格式读取文件, 默认读取一行
  - \*|l 默认, 从当前位置开始读取一行, 遇到文件末尾(EOF)返回 nil
  - \*L 读取一行并保留行结束标记(如果有的话), 当在文件末尾时，返回 nil
  - \*n 从当前位置读取数字直到行尾或者非数字字符结束并返回结果, 否则返回 nil
  - \*i 读取一个整数并返回
  - \*a 从当前位置开始读取所有内容
  - number 从当前位置读取指定数量 number 个字符并返回
- file:seek([where [, offset]]) 设置及获取当前文件的位置
  - cur 从当前位置开始, 默认
  - set 从文件头开始
  - end 从文件尾开始
  - offset 默认 0, 偏移量
- file:setvbuf(mode [, size]) 设置文件的缓冲模式
  - no 不缓冲, 输出操作立刻生效
  - full 完全缓冲, 只有在缓存满或当显式的对文件调用 flush 时才真正做输出操作
  - line 行缓冲, 缓冲有效将到每次换行前, 对于某些特殊文件(例如终端设备)缓冲到任何输入前
  - size 以字节为单位指定缓冲区大小
- file:write(...) 将参数的值逐个写入 file, 参数必须是字符串或数字, 成功返回 file, 否则返回 nil 和错误信息

```lua
print("文件 I/O: 用于读取和处理文件")
print("\t", "简单模式(simple mode): 拥有一个当前输入文件和一个当前输入文件, 并且提供针对这些文件相关的操作")
print("\t", "完全模式(complete mode): 使用外部的文件句柄来实现, 它以一种面向对象的方式, 将所有的文件操作定义为文件句柄的方法")
print("简单模式:")
-- 以追加的方式打开可读可写文件
file = io.open("test", "a+")
-- 设置默认输入文件为 test
io.input(file)
-- 读取文件
print("使用 io.read(\"*n\") 从当前位置读取数字直到行尾或者非数字字符结束并返回结果, 否则返回 nil", io.read("*n"))
print("使用 io.read 从当前位置读取一行",io.read())
print("使用 io.read('*|') 从当前位置读取一行, 遇到文件末尾(EOF)返回 nil", io.read("*|"))
print("使用 io.read(number) 从当前位置读取指定 number 个数的字符并返回", io.read(10))
-- 写入文件
io.write("-- 在当前位置追加内容, 追加的内容是注释\n")
-- 再次读取内容
print("使用 io.read(\"*a\") 从当前位置读取整个文件", io.read("*a"))
-- 关闭文件
io.close(file)
print("---------------")
print("完全模式: 使用 file:function_name 代替 io:function_name")
-- 打开文件
file = io.open('test', 'a+')
-- 输出文件第一行
print("读取一行", file:read())
print("读取一行", file:read())
print("读取 10 个字符", file:read(10))
print("file:seek(where, offset) 设置和获取当前文件位置", file:seek())
-- 写入文件
file:write("-- 完全模式插入的内容\n")
file:seek('set') -- 'set' 从文件头开始, 'cur' 从当前位置, 'end' 从文件尾开始, offset: 0 偏移量
print("读取所有内容", file:read("*a"))
-- 关闭文件
file:close()
```

### 操作系统库

- os.clock() 返回程序使用的按秒计 CPU 时间的近似值
- os.date([format [, time]]) 返回一个包含日期及时刻的字符串或表, 格式化方法取决于所给字符串 format
- os.difftime(t2, t1) 返回以秒计算的时刻 t1 到 t2 的差值
- os.execute([command]) 调用系统解释器执行 command, 执行成功返回 true, 否则返回 nil, 在第一个返回值之后, 函数返回一个字符串加一个数字
  - exit 命令正常结束, 接下来的数字是命令的退出状态码
  - signal 命令被信号打断, 接下来的数字是打断该命令的信号
- os.exit([code [, close]]) 终止宿主程序, 如果 close 为真, 在退出前关闭 lua 状态机
  - 如果 code 为 true, 返回的状态码是 EXIT_SUCCESS
  - 如果 code 为 false, 返回的状态码是 EXIT_FAILURE
  - 如果 code 是一个数字, 返回的状态码就是这个数字, code 默认值为 true
- os.getenv(varname) 返回进程环境变量 varname 的值, 如果未定义则返回 nil
- os.remove(filename) 删除指定的文件, 如果函数失败则返回 nil 和错误信息
- os.rename(oldname, newname) 重命名文件, 如果函数失败则返回 nil 和错误信息
- os.setlocale(locale [, category]) 设置程序的当前区域, locale 是一个区域设置的系统相关字符串, category 是一个描述由改变哪个分类的可选字符串: all, collate, ctype, monetary, numeric, time, 默认为 all
- os.time([table]) 当不传参数时, 返回当前时刻, 如果传入一张表则返回由这张表表示的时刻, 这张表必须包含域 year, month, day, 可以包含 hour(默认为 12), min(默认为 0), sec(默认为 0), 以及 isdst(默认为 nil)
- os.tmpname() 返回一个可用于临时文件的文件名字符串, 这个文件在使用前必须显式打开, 不再使用时需要显式删除

### 调试库

lua 提供了 debug 库用于提供创建自定义调试器的功能

- debug.debug() 进入一个用户交互模式,运行用户输入的每个字符串, 使用简单命令以及其他调试设置, 用户可以检阅全局变量和局部变量, 改变变量的值, 计算一些表达式等等, 输入一行仅包含 count 的字符串将结束这个函数继续向下运行
- debug.gethook([thread]) 返回三个表示线程钩子设置的值: 当前钩子函数, 当前钩子掩码, 当前钩子计数
- debug.getinfo([thread,] f [, what]) 返回一个关于函数信息的表, 也可以提供一个数字 f 表示的函数, 数字 f 表示运行在指定线程的调用栈对应层次上的函数, 0 层表示当前函数(getinfo 自身), 1 层表示调用 getinfo 的函数
- debug.getlocal([thread, ] f, local) 返回在栈的 f 层处函数的索引为 local 的局部变量的名字和值, 此函数不仅用于访问显式定义的局部变量, 还包括形参, 临时变量等
- debug.getmetatable(value) 返回给定 value 的元表, 如果没有则返回 nil
- debug.getregistry() 返回注册表表, 这是一个预定以的表, 可以用来保存任何 C 代码想保存的 lua 值
- debug.getupvalue(f, up) 返回函数 f 的第 up 个上值的名字和值, 如果没有则返回 nil
- debug.getuservalue(u) 返回关联在 u 上的 lua 值, 如果 u 并非用户数据, 返回 nil
- debug.sethook([thread,] hook, mask [, count]) 将一个函数作为钩子函数设入, 字符串 mask 以及数字 count 决定了钩子将在何时调用, mask: c 每当 lua 调用一个函数, 调用此钩子, r 每当 lua 从一个函数内返回时, 调用钩子, l 每当 lua 进入新的一行时, 调用钩子
- debug.setlocal([thread, ] level, local, value) 将 value 赋值给栈上第 level 层函数的第 local 个局部变量, 如果没有那个变量返回 nil, 如果 level 越界则抛出一个错误
- debug.setmetatable(value, table) 将 value 的元表设置为 table(可以是 nil), 返回 value
- debug.setupvalue(f, up, value) 将 value 设置为函数 f 第 up 个上值, 如果函数没有那个上值返回 nil, 否则返回 up 上值的名字
- debug.setuservalue(udata, value) 将 value 设置为 udata 的关联值, udata 必须是一个完全用户数据, 返回 udata
- debug.traceback([thread,] [message [, level]]) 追踪堆栈信息, message 被添加到栈回朔信息的头部, level 指定从栈的哪一层开始回朔(默认: 1)
- debug.upvalueid(f, n) 返回指定函数第 n 个上值的唯一标识符(一个轻量用户数据), 这个唯一标识符可以让程序检查两个不同的闭包是否共享了上值, 如果是则返回相同的标识符
- debug.upvaluejoin(f1, n1, f2, n2) 让 lua 闭包 f1 的第 n1 个值引用 lua 闭包 f2 的第 n2 个值

## 其他语法

```lua
#!/usr/local/bin/lua

print("使用 type函数获取类型: type()")
print("a = 100 type(a)", type(a))
print("type(nil)", type(nil))
print("---------------------------------------")

print("for(condition) do ... end 循环, 泛型 for 循环通过一个迭代器函数(pairs)来遍历所有值")
print("while(condition) do ... end 循环")
print("repeat ... until(condition) 循环, 先执行一次语句再判断条件, 像 while do")
local ru = [[
repeat
        statements
until(conditiion)
]]
print(ru)
print("循环控制语句: break goto")
print("goto 将控制流程转到被标记的语句处, ::label::")
print("---------------------------------------")

for var=1, 10, 2 do -- 第一个值为初始值, 第二个值为终止条件，第三个值为步长,省略默认为 1
        print(var) -- 1 3 5 7 9
end
print("---------------------------------------")

print("函数声明 function, 变长参数 ...")
local defineFunc = [[
optional_function_scope function function_name(argument1, ..., argumentN)
        function_body
        return result_params_comma_separated
end
]]
print(defineFunc, "\n")

function test(n)
        if n == 1 then
                return true
        elseif n > 1 then
                return  '> 1'
        else
                return false
        end
end
print("test(1)", test(1)) -- true
print("test(0)", test(0)) -- false
print("test(2)", test(2)) -- > 1
print("---------------------------------------")

print("迭代器(iterator): 是一种对象, 能够用来遍历标准模板库容器中的部分或全部元素, 每个迭代器对象代表容器中的确定的地址")
print("\t", "使用闭包模拟实现迭代器, 需要创建一个闭包的代价")
local function values(t)
    local i = 0
    return function() i = i + 1; return t[i] end
end
tb = {3, 4, 5, 6}
for v in values(tb) do
    print(v)
end
print("泛型 for: 自身保存迭代状态, 不必付出闭包的代价")
print("\t", "for k, v in pairs(t) do ... end")
print("无状态的迭代器: 不保留任何状态的迭代器, 在循环中可以利用无状态迭代器避免创建闭包花费额外的代价")
print("\t", "pairs() 会无序输出所有数据, 遇到 nil 不会停止输出")
print("\t", "ipairs() 按照 key 的顺序输出数据, 跳过字符串的 key, 遇到不连续的数据、nil会停止输出")
print("多状态的迭代器: 使用闭包, 或者将所有的状态信息封装到 table 内")
print("---------------------------------------")
```

## 包管理工具

LuaRocks 是一个 Lua 包管理器，基于 Lua 语言开发，提供一个命令行的方式来管理 Lua 包依赖、安装第三方 Lua 包等，社区比较流行的包管理器之一

### 安装

```bash
[root@centos7 workspace]# wget https://luarocks.org/releases/luarocks-3.9.2.tar.gz
[root@centos7 workspace]# tar zxpf luarocks-3.9.2.tar.gz
[root@centos7 workspace]# cd luarocks-3.9.2
[root@centos7 workspace]# ./configure && make && sudo make install
[root@centos7 workspace]# luarocks --version
```

### 常用命令

- \-\-help 查看命令帮助信息
- show \<rock\> 显示指定包的详细信息
- doc \<rock\> 查看指定包的文档
- search \<rock\> 查找服务器中的指定包
- list 列出所有安装的包
- install \<rock\> 安装指定的包
  - \-\-check-lua-versions 检查包的兼容性
- remove \<rock\> 移除指定的包
- purge 移除所有的包
- path 显示当前配置的包的目录
- build 编译安装当前目录的 rock
- make 在当前目录中使用 rockspec 文件编译安装 rock
- pack 在当前目录下创建一个 rock
- unpack \<rock\> 解包一个 rock
- upload \<rockspec\> 创建一个 rock 并上传到公共服务器

### 连接 redis

```lua
local redis = require "redis"
local client = redis.connect('127.0.0.1', 6379)

local res = client:ping()

print(res)

local name = client:get('name')
print(name)

local xiaoming = client:hgetall('xiaoming')
for key, val in pairs(xiaoming) do
    print(key, " => ", val)
end
```
