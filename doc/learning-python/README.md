
项目结构

my_project/
├── .venv/                 # 虚拟环境目录 (通常跟随项目)
│   ├── bin/               # (Linux/macOS) 可执行文件
│   ├── Scripts/           # (Windows) 可执行文件
│   ├── lib/               # (Linux/macOS) 存放安装的包
│   ├── Lib/               # (Windows) 存放安装的包
│   └── pyvenv.cfg         # 虚拟环境配置文件
├── .python-version        # (可选) 指定项目所需的 Python 版本
├── pyproject.toml         # 项目的核心配置文件 (记录依赖声明)
├── uv.lock                # 依赖的精确锁定文件 (由 uv 自动维护)
├── requirements.txt       # (可选) 传统的依赖列表文件
|—— 项目源代码

## 数据类型

```python
# 单行注释

'''
多行字符串/注释, 如果多行字符串有变量引用则作为字符串, 否则作为多行注释
'''

"""
多行字符串/注释, 如果多行字符串有变量引用则作为字符串, 否则作为多行注释
"""
```

- int
- float
- bool, 底层使用整数 1, 0 表示 True, False
- str
- tuple
- list
- set
- dict

```python
print("hello world")
```

- `整型和布尔值混合转换为整型`
- `整型和浮点型混合转换为浮点数`
- `布尔值和浮点型混合转换为浮点数`

```python
# 整型和布尔值
>>> 1 + True
2
>>> 1 + False
1

# 整型和浮点型
>>> 1. + 2
3.0
>>> 1 + 1.2
2.2

# 布尔值和浮点型
>>> 1. + True 
2.0
>>> 1.0 + True 
2.0
>>> 1.0 + False
1.0
```

### 字符串

`只读的字符序列`

- [] 获取字符, 下标越界将报错

```python
>>> letters = 'abcdefg' 
>>> letters[0]
'a'
>>> letters[-1]
'g'
>>> letters[-2]
'f'
>>> letters[10]  
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
IndexError: string index out of range
>>> letters[-100]
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
IndexError: string index out of range
```

- `分片`获取子串, [strat:end:step], 含头不含尾, start 和 end 越界不会报错
  - start, 起始索引, 包含, 默认为字符串开始
  - end, 结束索引, 不包含, 默认为字符串结束
  - step, 步长绝对值，默认为 1, 每隔指定数量的`绝对值`的字符取一次
    - 如果为 -1, `反向提取子串`

```python
'''
  a  b  c  d  e  f  g
  0  1  2  3  4  5  6  正索引
 -7 -6 -5 -4 -3 -2 -1  负索引
'''

>>> letters = 'abcdefg'
>>> letters[:]
'abcdefg'
>>> letters[-8:-2]  # start 超出起始位置使用默认值
'abcde'
>>> letters[-8:10]  # end 超出结束位置使用默认值
'abcdefg'
>>> letters[3:-2] # 含头不含尾
'de'
>>> letters[8:12] # 下标越界提取为空串
''
>>> letters[8:-2] # 下标越界提取为空串
''
# 指定步长
>>> letters[::3]  # 每隔 3 个字符取一次
'adg'
>>> letters[-5:-2:2]  # 每隔 2 个字符取一次
'ce'
>>> letters[-5:-5:2]
''

>>> letters[::-1] # 反向提取字符串
'gfedcba'
>>> letters[6:1:-1]
'gfedc'
>>> letters[6:1:-2] # 反向每隔 2 个字符取一次
'gec'
>>> letters[-3:1:-2]
'ec'
```

- strip() 移除字符串两端的空白字符('', '\t', '\n')

```python
>>> t1 = '\n\t\na\n\tb\n\t\n'
>>> t1.strip()
'a\n\tb'
>>> t1.lstrip()
'a\n\tb\n\t\n'
>>> t1.rstrip()
'\n\t\na\n\tb'
```

- find() 和 index(), 查找子串出现的下标, find 未找到返回 -1, index 未找到报错

```python
>>> t1.find('a'， 3)  # 指定开始查找位置
3
>>> t1.rfind('a', 3)  # 指定开始查找位置  
3
>>> t1.index('a', 3)  # 指定开始查找位置
3
>>> t1.rindex('a', 3) # 指定开始查找位置
3

>>> t1.find('aa')
-1
>>> t1.rfind('aa')
-1
>>> t1.index('aa')
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: substring not found
>>> t1.rindex('aa')
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: substring not found
```

- 对齐字符串(center, ljust, rjust), 在指定的空间范围内对齐, 如果空间范围小于字符串长度则返回原值

```python
>>> s = 'a duck goes into bar...'
>>> len(s)
23
>>> s.center(5) # 对齐空间范围小于字符串长度返回原值
'a duck goes into bar...'
>>> s2 = s.center(40) # 居中对齐
>>> s2
'        a duck goes into bar...         '
>>> len(s2)
40
>>> s.ljust(40) # 左对齐
'a duck goes into bar...                 '
>>> s.rjust(40) # 右对齐
'                 a duck goes into bar...'
```

#### 字符串拼接

- 字符串字面量拼接，`只能是字符串字面量`

```python
>>> name = 'python'
>>> 'hello' name
  File "<stdin>", line 1
    'hello' name
            ^^^^
SyntaxError: invalid syntax

>>> 'hello' "world"
'helloworld'
>>> 'hello' "world" '''你好''' """python"""
'helloworld你好python'
>>> '''hello'''"world" """你好""" 'python'
'helloworld你好python'
```

- 使用 * 重复字符串，`只能是字符串类型和整型`

```python
>>> """hello"""*2
'hellohello'
>>> num = 3
>>> 'hello' * num
'hellohellohello'

>>> 2 * 'hello'
'hellohello'

>>> 'hello' * 2.
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: can't multiply sequence by non-int of type 'float'
```

- 使用 + 拼接的值 `只能是字符串类型`, 其他类型需要先转换为字符串

```python
>>> 'hello' + 1
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: can only concatenate str (not "int") to str
>>> 'hello' + 1.2
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: can only concatenate str (not "float") to str
>>> 'hello' + True
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: can only concatenate str (not "bool") to str

name = "luolei"
message = "hello " + name
print(message)
```

- 使用 % 占位符, 需要考虑浮点数的精度问题

```python
# 字符串拼接, 占位符 %,  %s 是字符串占位符，%d 是整数占位符，%f 是浮点数占位符
name = "luolei"
message = "hello %s" % name
print(message)

salary = 5000
tel = 13112345678
message = "my salary is %.2f, my tel is %d" % (salary, tel)
print(message)
```

- 使用 format 方法

```python
# 字符串拼接, format方法
message = "my salary is {:.2f}, my tel is {}".format(salary, tel)
print(message)
```

- 使用 f-string, f"{占位符}", 不需要考虑浮点数的精度问题

```python
# 字符串拼接, f-string, 精度控制可选
message = f"my salary is {salary:.2f}, my tel is {tel}"
print(message)
print(F"""my salary is {salary:.2f}, my tel is {tel}""")

age = 18
>>> print(f'''age: {age}''')
age: 18
>>> print(f"""age: {age}""")
age: 18
>>> print(F"""age: {age}""")
age: 18
>>> print(F'''age: {age}''')
age: 18
```

- 使用 join 方法拼接字符串, 拼接的元素只能是字符串类型

```python
>>> t = ['get gloves,get mask.give cat vitamins,call ambulance', 'get glove', 'get mak']
>>> t2 = '--gg,gg--'
>>> t2.join(t)
'get gloves,get mask.give cat vitamins,call ambulance--gg,gg--get glove--gg,gg--get mak'
>>> t3 = t2.join(t)
>>> t3.split(t2)
['get gloves,get mask.give cat vitamins,call ambulance', 'get glove', 'get mak']

>>> t4 = [1, 2, 3, True]
>>> type(t4)
<class 'list'>
>>> ','.join(t4)
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: sequence item 0: expected str instance, int found
```

### 内置函数

#### int()

将任意值转换成整型化值

- 第一个参数为要转换的值
  - 浮点数则截断整数部分
  - 参数如果是字符串, 只能由 `数字, +, -, _` 组成
  - 布尔值 True 转为 1, False 转为 0
- 第二个参数为进制数位

```python
>>> int(1.)
1
>>> int(1.8)
1
>>> int(1_0_0)
100
>>> int("+1_0")
10
>>> int(1e10)
10000000000
>>> int(True)
1
>>> int(False)
0
```

#### float()

将任意值转换为浮点化值

- 参数为整型则转换为浮点数
- 参数为字符串的浮点数表示形式，支持科学计数格式
- 布尔值 True 转为 1.0, False 转为 0.0

```python
>>> float(1)
1.0
>>> float(1.)
1.0
>>> float('1')
1.0
>>> float('1.')
1.0
>>> float('+1')
1.0
>>> float('-1')
-1.0
>>> float('+1.')
1.0
>>> float('-1.')
-1.0
>>> float('+1.0')
1.0
>>> float('-1.0')
-1.0
>>> float('-1.0_0')
-1.0
>>> float('-1_0.0_0')
-10.0
>>> float('1.e10')
10000000000.0
>>> float(True)
1.0
>>> float(False)
0.0
```

#### bool()

将任意值转为布尔值

- 任意非零值转为 True
- 任意零值转为 False

```python
>>> bool(1)
True
>>> bool(1.0)
True
>>> bool('1.0')
True
>>> bool('0.0')
True
>>> bool('True')
True
>>> bool('False') # 任意非零值转换为 True
True
>>> bool('')  # 任意零值转换为 False
False
>>> bool(-1.0)
True
>>> bool(+0.0)
False
>>> bool(-0.0)
False
>>> bool(0.0)
False
>>> bool(0)
False
>>> bool(1e0)
True
```

#### str()

将任意值转换为字符串

```python
>>> str(1)
'1'
>>> str(1.)
'1.0'
>>> str(1.0)
'1.0'
>>> str(1.01)
'1.01'
>>> str(1.0_1)
'1.01'
>>> str(-1)
'-1'
>>> str(-1.)
'-1.0'
>>> str(-1.0)
'-1.0'
>>> str(True)
'True'
>>> str(False)
'False'
```

#### chr()/ord()

转换 unicode 码对应的字符串

```python
>>> chr(97)
'a'
>>> chr(65)
'A'
>>> chr(48)
'0'
>>> ord('a')
97
>>> ord('一')
19968
>>> chr(19968)
'一'
>>> ord('棒')
26834
>>> chr(26834)
'棒'
```
