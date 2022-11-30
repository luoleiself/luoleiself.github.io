---
title: Lua
date: 2022-11-28 16:45:19
categories:
  - [server, Lua]
tags:
  - Lua
---

### Lua

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
print("元素(metatable), 改变 table 的行为, 每个行为关联的对应的元方法")
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
print("__gt 对应的操作符 >")
print("__ge 对应的操作符 >=")
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
print("---------------------------------------")

print("模块: 封装公用的代码以 API 接口的形式在其他地方调用")
print("简单理解是将变量、常量、函数放在一个table里面，然后 return 返回")
print("使用 require 方法加载模块, require(\"模块名\") 或者 require \"模块名\"")
print("模块的加载机制: require 用于搜索 lua 文件的路径是存放在全局变量 package.path 中, 当 lua 启动后, 会以环境变量 LUA_PATH 的值来初始这个环境变量, 如果没有找到该环境变量, 则使用一个编译时定义的默认路径来初始化, 此环境变量也可以自定义设置, 在搜索过程中, 如果找到该文件, 则使用 pacakge.loadfile 来加载模块, 否则就去找 C 程序库, 搜索的文件路径是从全局变量 package.cpath 获取, 而这个变量则是通过环境变量 LUA_CPATH 来初始, 此时搜索的文件是 so 或 dll 类型的文件, 如果找到了则使用 package.loadlib 来加载")
print("---------------------------------------")
```
