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
#!/usr/local/bin/lua  -- shebang
print("hello world")

--这是单行注释, 不会被执行

--[[这是多行注释, 不会被执行
这是多行注释, 不会被执行
这是多行注释, 不会被执行
这是多行注释, 不会被执行]]--

print('单行注释: \n-- 单行注释内容')
print("多行注释: \n--[[\n 多行注释内容\n 多行注释内容\n 多行注释内容\n]]--")
print("---------------------------------------")

print("关键字: and break do else elseif end false for function if in local nil not or repeat return then true until while goto")

print("默认声明的变量都是全局变量, 给一个变量赋值即创建了一个全局变量, 访问一个未初始化的全局变量返回 nil, 删除全局变量赋值为 nil 即可")
a = 100
print(a)
print("---------------------------------------")

print("数据类型: nil boolean number string userdata function trhead table")
print("nil 表示一个无效值(在条件表达式中相当于 false)")
print("boolean 表示 true 和 false, 除了 nil 和 false 值表示为 false, 其他值(包括 0)都为 true, ")
print("number 表示双精度类型的实浮点数, 数字字符串相加将转换成数字相加")
print("string 表示一对双引号或者单引号包含的内容, [[内容]] 表示块字符串, 字符串拼接使用 .., 计算字符串的长度使用 #")
print("userdata 表示任意存储在变量中的 C 数据结构")
print("function 由 C 或 Lua 编写的函数")
print("thread 表示执行的独立线程, 用于执行协同程序")
print("table 关联数组, 数组的索引可以是数字、字符串或者表类型, 下标默认从 1 开始, table 的创建通过'构造表达式'完成, 空表: {}")
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
print("---------------------------------------")

print("使用 type 函数获取类型: type()")
print("a = 100 type(a)", type(a))
print("type(nil)", type(nil))
print("---------------------------------------")

print("for(condition) do 循环, 泛型 for 循环通过一个迭代器函数(pairs)来遍历所有值")
print("while(condition) do 循环")
print("repeat .. until(condition) 循环, 先执行一次语句再判断条件, 像 while .. do")
local ru = [[
repeat
    statements
until(conditiion)
]]
print(ru)
print("循环控制语句: break goto")
print("---------------------------------------")

for var=1, 10, 2 do -- 第一个值为初始值, 第二个值为终止条件，第三个值为步长,省略默认为 1
    print(var) -- 1 3 5 7 9
end
print("---------------------------------------")

local tb1 = {}
local tb2 = {name = "zhangsan", age = 18, "beijing"}
print("声明空表", tb1)
print("声明非空表", tb2)
print("for 遍历 tb2")
for key, val in pairs(tb2) do
    print(key, ":", val) -- 1:beijing  age:18  name:zhangsan
end
print("---------------------------------------")

print("函数声明 function")
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
```
