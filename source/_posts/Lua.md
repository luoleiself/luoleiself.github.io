---
title: Lua
date: 2022-11-28 16:45:19
categories:
  - [server, Lua]
tags:
  - Lua
---

### Lua

Lua 是一门强大、快速、轻量的嵌入式动态类型脚本语言, 使用 ANSI C 语言编写并以源代码形式开放, 其设计目的是为了嵌入应用程序中, 从而为应用程序提供灵活的扩展和定制功能

Lua 有八种基本数据类型: nil、boolean、number、string、function、userdata、 thread、table

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

```lua
#!/usr/local/bin/lua
print("hello world")

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

print("关键字: and break do else elseif end false for function if in local nil not or repeat return then true until while goto")

print("默认声明的变量都是全局变量, 给一个变量赋值即创建了一个全局变量, 访问一个未初始化的全局变量返回 nil, 删除全局变量赋值为 nil 即可")
a = 100
print(a)
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
print("string.upper 转换大写", string.upper("AbCdE"))
print("string.lower 转换小写", string.lower("AbCdE"))
print("string.gsub 字符串查找替换", string.gsub("hello world", "l", 'r'))
print("string.find 查找子串位置", string.find("hello lua user", "lua"))
print("string.reverse 字符串反转", string.reverse("hello world"))
print("string.format 根据字符串模板返回格式化的字符串", string.format("the value is %d", 4))
print("string.char 返回数值表示的字符, byte 返回字符的数值表示", string.char(96, 99, 100), string.byte('ABCD'))
print("string.len 返回字符串的长度", string.len("abc"))
print("string.rep 返回字符串的 n 个拷贝", string.rep("abc", 3))
print("string.gmatch 返回一个迭代器函数, 每次调用函数返回一个字符串中找到的下一个符合 pattern 模式的子串, 可以结合 for 循环查找")
print("string.match 返回在字符串中查找符合匹配模式的第一个子串")
print("string.sub 截取字符串", string.sub('hello world', 1, 6))
print("---------------------------------------")

print("% 作为特殊字符的转义字符, 就像正则表达式中的转义字符 \\")
print("匹配模式: %a 表示任意字符, %c 表示任意控制字符, %d 表示任意数字, %l 表示任意小写字母, %u 表示任意大写字母, %s 表示任意空白字符, %w 表示任意字母/数字, %x 表示任意十六进制数, %p 表示任意标点, %z 表示任意代表0的字符")
print("当上述字符类用大写表示时, 表示与非此字符类的任意字符匹配, %S 表示任意非空白字符")
print("---------------------------------------")

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
print("泛型 for 迭代器: for k, v in pairs(t) do ... end")
print("无状态的迭代器: 不保留任何状态的迭代器, 在循环中可以利用无状态迭代器避免创建闭包花费额外的代价, ipairs 无状态迭代函数")
print("多状态的迭代器: 使用闭包, 或者将所有的状态信息封装到 table 内")
print("---------------------------------------")

local tb1 = {"banana", "orange", "apple", "grapes"}
local tb2 = {name = "zhangsan", age = 18, "beijing"}
print("排序前: for 遍历 tb1")
for key, val in pairs(tb1) do
        print(val) -- banana orange apple grapes
end
print("table 其实是关联数组, 可以使用任意类型值(除了 nil)作为数组的索引, 不固定大小")
print("table.concat 列出表中指定区间的所有元素", table.concat({"hello", "world", "lua", 2022}, "-"))
print("table.insert 在表中指定位置插入一个元素", table.insert(tb1, 3, "hello world"))
print("table.remove 移除并返回表中指定位置的元素", table.remove(tb2, 2))
print("table.sort 对指定的表进行升序排序", table.sort(tb1))
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
print("元表(metatable), 改变 table 的行为, 每个行为关联的对应的元方法")
print("setmetatable(table, metatable) 对 table 设置元表")
print("getmetatable(table) 返回对象的元表")
print("__add 对应的操作符 +")
print("__sub 对应的操作符 -")
print("__mul 对应的操作符 *")
print("__div 对应的操作符 /")
print("__mod 对应的操作符 %")
print("__unm 对应的操作符 -")
print("__concat 对应的操作符 ..")
print("__eq 对应的操作符 ==")
print("__lt 对应的操作符 <")
print("__le 对应的操作符 <=")
print("__len 对应的操作符 #")
print("__tostring 元方法用于修改表的输出行为(自定义输出内容)")
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
            local sum = 0
            for k, v in pairs(t) do
                sum = sum + v
            end
            return "表中所有元素的和为 "..sum
        end
})
print(mtstring) -- 表中所有元素的和为 60
print("__call 元方法在 lua 调用一个值时调用")
print("__newindex 元方法用来对 table 更新")
print("\t", "如果给 table 不存在的 key 赋值, 解释器就会查找  __newindex 元方法如果存在则会调用这个函数而不进行赋值操作")
print("\t", "如果给 table 已存在的 key 赋值, 则不会调用 __newindex 元方法")
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

print("__index 元方法用来对 table 访问, 当通过 key 访问 table 时, 如果 key 不存在, lua 则会寻找该 metatable 中的 __index 键")
print("\t", "如果 __index 键包含一个 table, lua 则会在这个 table 中查找相应的 key")
print("\t", "如果 __index 键包含一个函数时, lua 则会调用这个函数, table 和 key 作为参数传递给函数, 并接收函数的返回值作为结果")
print("查找顺序:")
print("\t", "1. 在 table 中查找, 如果找到则返回该元素, 否则继续")
print("\t", "2. 判断该 table 是否有元表, 如果没有则返回 nil, 否则继续")
print("\t", "3. 判断元表是否有 __index 键, 如果没有则返回 nil, 如果 __index 是一个 table, 则重复 1. 2. 3, 如果 __index 是一个函数, 则返回调用该函数的返回值")
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
})
]]
print(mytableStr)
print("print(mytable.foo, mytable.baz)", mytable.foo, mytable.baz) -- bar     baz = baz
print("获取元表 getmetatable(mytable)", getmetatable(mytable))
print("---------------------------------------")

print("模块: 封装公用的代码以 API 接口的形式在其他地方调用")
print("简单理解是将变量、常量、函数放在一个table里面，然后 return 返回")
print("使用 require 方法加载模块, require(\"模块名\") 或者 require \"模块名\"")
print("模块的加载机制: require 用于搜索 lua 文件的路径是存放在全局变量 package.path 中, 当 lua 启动后, 会以环境变量 LUA_PATH 的值来初始这个环境变量, 如果没有找到该环境变量, 则使用一个编译时定义的默认路径来初始化, 此环境变量也可以自定义设置, 在搜索过程中, 如果找到该文件, 则使用 pacakge.loadfile 来加载模块, 否则就去找 C 程序库, 搜索的文件路径是从全局变量 package.cpath 获取, 而这个变量则是通过环境变量 LUA_CPATH 来初始, 此时搜索的文件是 so 或 dll 类型的文件, 如果找到了则使用 package.loadlib 来加载")
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
print("---------------------------------------")

print("协程(coroutine): 拥有独立的堆栈, 独立的局部变量, 独立的指令指针, 同时又与其他协程共享全局变量和其他大部分东西")
print("与线程的主要区别: 一个具有多个线程的程序可以同时运行多个线程, 协程却需要彼此协作的运行, 在任一指定时刻只有一个协程在运行, 并且这个正在运行的协程只有在明确的被要求挂起时才会被挂起")
print("coroutine.create() 创建并返回一个 coroutine, 参数为一个函数, 当和 resume 配合使用时就唤醒函数调用")
print("coroutine.resume() 重启 coroutine, 和 create 配合使用")
print("coroutine.yield() 挂起 coroutine, 如果有参数, 则将参数返回给调用线程")
print("coroutine.status() 查看 coroutine 的状态, 通常返回 dead, suspended, running")
print("coroutine.wrap() 创建 coroutine, 并返回一个函数, 一旦手动调用这个函数, 就进入 coroutine, 和 create 功能重复")
print("coroutine.running() 返回正在运行的 coroutine")
local costr = [[
co = coroutine.create(
        function (i)
                print(i)
        end
)
coroutine.resume(co, 100) -- 100
print(coroutine.status(co)) -- dead
print(coroutine.running()) -- thread: 0x149b018  true
print("---------------")
co = coroutine.wrap(
        function (i)
                print(i)
        end
)
co(250) -- 250
print("---------------")
co = coroutine.create(
        function ()
                for i = 1, 10 do
                        print(i) -- 1 2 3
                        if i == 3 then
                                print(coroutine.status(co)) -- running
                                print(coroutine.running()) -- 返回正在运行的 coroutine thread: 0x14ab788       false
                         end
                         coroutine.yield() -- 挂起 coroutine
                end
        end
)
coroutine.resume(co)
coroutine.resume(co)
coroutine.resume(co)
]]
print(costr)
print("---------------")
co = coroutine.create(
        function (i)
                print(i)
        end
)
coroutine.resume(co, 100) -- 100
print(coroutine.status(co)) -- dead
print(coroutine.running()) -- thread: 0x149b018  true
print("---------------")
co = coroutine.wrap(
        function (i)
                print(i)
        end
)
co(250) -- 250
print("---------------")
co = coroutine.create(
        function ()
                for i = 1, 10 do
                        print(i) -- 1 2 3
                        if i == 3 then
                                print(coroutine.status(co)) -- running
                                print(coroutine.running()) -- 返回正在运行的 coroutine thread: 0x14ab788       false
                         end
                         coroutine.yield() -- 挂起 coroutine
                end
        end
)
coroutine.resume(co)
coroutine.resume(co)
coroutine.resume(co)
print("---------------------------------------")
print("生产者和消费者")
local prodstr = [[
local newProductor
function productor()
        local i = 0
        while true do
                i = i + 1
                send(i) -- 生产者发送数据
        end
end
function consumer()
        while true do
                local i = receive() -- 消费者接收数据
                print(i)
        end
end
function receive()
        local status, value = coroutine.resume(newProductor) -- 唤起 coroutine
        return value
end
function send(v)
        coroutine.yield(v) -- 发送数据后就挂起 coroutine
end
newProductor = coroutine.create(productor)
consumer()
]]
print(prodstr)
print("---------------")
local newProductor
function productor()
        local i = 0
        while true do
                i = i + 1
                send(i) -- 生产者发送数据
        end
end
function consumer()
        while true do
                local i = receive() -- 消费者接收数据
                print(i)
        end
end
function receive()
        local status, value = coroutine.resume(newProductor) -- 唤起 coroutine
        return value
end
function send(v)
        coroutine.yield(v) -- 发送数据后就挂起 coroutine
end
--newProductor = coroutine.create(productor)
--consumer()
print("---------------------------------------")

print("文件 I/O: 用于读取和处理文件")
print("\t", "简单模式(simple mode): 拥有一个当前输入文件和一个当前输入文件, 并且提供针对这些文件相关的操作")
print("\t", "完全模式(complete mode): 使用外部的文件句柄来实现, 它以一种面向对象的方式, 将所有的文件操作定义为文件句柄的方法")
print("io.seek(where, offset) 设置和获取当前文件的位置")
print("\t", "set 从文件头开始")
print("\t", "cur 从当前位置开始, 默认")
print("\t", "end 从文件尾开始")
print("\t", "offset 默认 0, 偏移量")
print("io.read() 读取文件, 默认读取一行")
print("\t", "*n 从当前位置读取数字直到行尾或者非数字字符结束并返回结果, 否则返回 nil")
print("\t", "*a 从当前位置开始读取所有内容")
print("\t", "*| 默认, 从当前位置开始读取一行, 遇到文件末尾(EOF)返回 nil")
print("\t", "number 从当前位置读取指定数量 number 个字符并返回")
print("简单模式:")
local iostr = [[
-- 以追加的方式打开可读可写文件
file = io.open("test", "a+")
-- 设置默认输入文件为 test
io.input(file)
-- 读取文件
print("使用 io.read(\"*n\") 从当前位置读取数字直到行尾或者非数字字符结束并返回结果, 否则返回 nil", io.read("*n"))
print("使用 io.read 从当前位置读取一行",io.read())
--print("使用 io.read('*|') 从当前位置读取一行, 遇到文件末尾(EOF)返回 nil", io.read("*|"))
print("使用 io.read(number) 从当前位置读取指定 number 个数的字符并返回", io.read(10))
-- 写入文件
io.write("-- 在当前位置追加内容, 追加的内容是注释\n")
-- 再次读取内容
print("使用 io.read(\"*a\") 从当前位置读取整个文件", io.read("*a"))
-- 关闭文件
io.close(file)
]]
print(iostr)
print("---------------")
print("完全模式: 使用 file:function_name 代替 io:function_name")
local iocmstr = [[
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
]]
print(iocmstr)
print("---------------------------------------")

print("错误: 语法错误 和 运行时错误")
print("错误处理:")
print("assert(arg1, arg2) 类型断言, 如果第一个参数为真, assert 不做任何处理, 否则将第二个作为错误信息输出:  assert(type(a) == 'number', 'a 不是一个数字')")
print("error(message [, level]) 终止正在执行的函数, 并返回 message 的内容作为错误信息, level 指示获取错误的位置: 1 默认, 为调用 error 的位置(文件+行), 2 指出调用 error 函数的函数, 0 不添加错误位置信息")
print("pcall() 保护模式调用, 接收一个函数和要传递给函数的参数, 以保护模式执行第一个参数, 可以捕获函数执行中的任何错误, 无错误返回 true, 有错误返回 false 和 errorinfo")
print("xpcall() 第一个参数和第二个后面的参数作用同 pcall, 第二个参数为一个错误处理函数, 当错误发生时, lua 会在调用栈展开前调用错误处理函数, debug.debug() 提示一个 lua 提示符, 让用户来检查错误的原因, debug.traceback() 根据调用栈来构建一个扩展的错误消息")
print("---------------------------------------")

print("调试: lua 提供了 debug 库用于提供创建自定义调试器的功能")
print("debug() 进入一个用户交互模式, 运行用户输入的每个字符串, 使用简单的命令以及其他调试设置, 用户可以检阅全局变量和局部变量, 改变变量的值, 计算一些表达式等等")
print("getfenv(object) 返回对象的环境变量")
print("gethook(optional thread) 返回三个表示线程钩子设置的值: 当前钩子函数, 当前钩子掩码, 当前钩子计数")
print("getinfo([thread,] f [, where]) 返回一个关于函数信息的表, 也可以提供一个数字 f 表示的函数, 数字 f 表示运行在指定线程的调用栈对应层次上的函数, 0 层表示当前函数(getinfo自身), 1 层表示调用 getinfo 的函数")
print("debug.getlocal([thread,] f, local) 返回在栈的 f 层处函数的索引为 local 的局部变量的名字和值, 此函数不仅用于访问显式定义的局部变量, 还包括形参, 临时变量等")
print("getmetatable(value) 把给定索引指向的元表压入堆栈")
print("getregistry() 返回注册表表, 这是一个预定以的表, 可以用来保存任何 C 代码想保存的 lua 值")
print("getupvalue(f, up) 返回函数 f 的第 up 个上值的名字和值, 如果没有则返回 nil")
print("sethook([thread,] hook, mask [, count]) 将一个函数作为钩子函数设入, 字符串 mask 以及数字 count 决定了钩子将在何时调用, mask: c 每当 lua 调用一个函数, 调用此钩子, r 每当 lua 从一个函数内返回时, 调用钩子, l 每当 lua 进入新的一行时, 调用钩子")
print("setlocal([thread,] level, local, value) 将 value 赋值给栈上第 level 层函数的第 local 个局部变量, 如果没有那个变量返回 nil, 如果 level 越界则抛出一个错误")
print("setmetatable(value, table) 设置元表")
print("setupvalue(f, up, value) 将 value 设置为函数 f 第 up 个上值, 如果函数没有那个上值返回 nil, 否则返回 up 上值的名字")
print("traceback([thread,] [message [, level]]) 追踪堆栈信息, message 被添加到栈回朔信息的头部, level 指定从栈的哪一层开始回朔(默认: 1)")
print("---------------------------------------")

print("垃圾回收: lua 采用了自动内存管理, collectgarbage([opt] [, arg])")
print("collectgarbage('collect') 做一次完整的垃圾回收循环")
print("collectgarbage('count') 以 K 字节为单位返回 lua 使用的总内存数")
print("collectgarbage('restart') 重启垃圾回收器的自动运行")
print("collectgarbage('setpause') 将 arg 设置为收集器的间歇率, 返回间歇率的前一个值")
print("collectgarbage('setstepmul') 返回步进倍率的前一个值")
print("collectgarbage('step') 单步运行垃圾收集器, 步长大小由 arg 决定")
print("collectgarbage('stop') 停止垃圾收集器的运行")
print("---------------------------------------")

print("面向对象: lua 使用 table 描述对象的属性, 使用 function 描述方法, 使用 table + function 模拟面向对象")
local oopstr = [[
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
]]
print(oopstr)
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
