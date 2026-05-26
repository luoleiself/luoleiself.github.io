python 解释器配置文件:
> %APPDATA%\JetBrains\PyCharm2026.x\options\jdk.table.xml

> \ / : * ? " < > |

项目结构

```conf
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
```

- GIL(Global Interpreter Lock) 全局解释器锁, 是 Python 解释器中的互斥锁机制, 确保在同一进程内同一时刻只有一个线程执行 python 字节码, 
  它简化了内存管理但是限制了多线程的并行计算能力.
  - 设计目的: python 使用引用计数作为内存管理方式, 确保同一时刻只有一个线程修改引用计数, 避免多线程内存冲突.
  - 释放条件: 时间片耗尽、遇到I/O操作, python3.2 后采用优先级平衡策略, 防止单个线程长时间持有 GIL 导致其他线程阻塞.
  - 多进程方案: 创建多个独立进程, 每个进程有独立的 GIL 和内存空间.
  - 异步编程: 使用 asyncio 实现高效的 I/O 并发处理, 不受 GIL 限制.
  - 替代解释器: IronPython, Jython, PyPy 等.

- 序列: 连续的可用下标访问的数据存储结构, 字符串, 元组, 列表, 数组, 可使用`切片`进行操作
- `*` | `**`: 在数据存储结构中使用时, 解构 `可迭代对象`
- `_` 匿名变量
  - 在 REPL 中保存上一表达式结果
  - 在 match-case 中匹配任意值但不绑定

- `==`, is, isinstance 运算符和函数,
  - `==` 比较值是否相等, 列表和元组必须有序, 集合无序
  - is 比较内存地址是否是同一个对象, 常用 None 检查
  - isinstance 类型检查, 检查对象是否是某个类或其子类的实例

```python
# == 比较值是否相等, 小整数缓存
# python 会缓存 -5 到 256 的整数
>>> a = 5
>>> b = 5
>>> a == b
True
>>> a is b  # 小整数缓存
True
>>> a = 1000
>>> b = 1000
>>> a == b
True
>>> a is b
False
```

match-case 结构化模式匹配语法, 用来替代冗长的 if-elif-else 链, 支持值匹配、结构解包和类型校验.

> python 3.10 支持

- match 定义表达式
- case 定义匹配模式, 按顺序匹配, 命中第一个就跳出, 不再检查后续 case 块
  - `_` 通配符, 匹配任意值
  - 字符串不支持 `f''` 和 `t''`

```python
match event:
    case r'\nhello world':  # 原始字符串匹配
      print("hello world")
    case 400 | 404: # 匹配 400 或 404
        print("Bad request")
    case int(): # 类型匹配
        print('整型')
    case list() | tuple():  # 类型匹配
        print('集合类型')
    case [x, *y]:   # 列表解包
        print("value is", x, y)
    case {'status': code, 'data': data}:    # 匹配包含 status 和 data 的字典
        print("status:", code, "data:", data)
    case [x, y] if x >= 0 and y >= 0:  # 先匹配列表, 再检查条件
        print("x and y are positive")
    case (x, y) if x == y:  # 先匹配元组, 再检查添加
        print("x is equal to y")
    case Point(x=0, y=3):   # 匹配 Point 类实例
        print("x is 0, y is 3")
    case _: # 匹配任意值
        print("No match")
```

## 关键字

- global, 用于在任何函数内部声明变量为全局变量
- nonlocal, 只能用于在嵌套函数中声明一个来自外层函数的变量, 既不是当前函数的变量也不是全局变量

```python
# 全局变量声明 global
count = 0
total = 100
name = "Python"

def update_stats():
    global count, total  # 可以同时声明多个全局变量
    count += 1
    total -= 10
    name = "Java"  # 如果没有 global name，这是局部变量

update_stats()
print(f"count: {count}")   # 输出: count: 1
print(f"total: {total}")   # 输出: total: 90
print(f"name: {name}")     # 输出: name: Python（没变）

# 外层变量声明 nonlocal
def level1():
    x = 1
    
    def level2():
        x = 2  # level2 的局部变量
        
        def level3():
            nonlocal x  # 指向 level2 的 x，不是 level1 的 x
            x = 3
            print(f"level3: {x}")
        
        level3()
        print(f"level2: {x}")
    
    level2()
    print(f"level1: {x}")

level1()
# 输出:
# level3: 3
# level2: 3  (被 level3 修改了)
# level1: 1  (没变)
```

- lambda, 定义匿名函数, 不能复用, `lambda 执行参数: 函数体(一行代码)`
- pass, 空操作占位符, 避免语法错误
- raise, 抛出异常

## 数据类型

- int, 不可变类型
- float, 不可变类型
- complex, 不可变类型
- bool, 不可变类型, 底层使用整数 1, 0 表示 True, False
- str, 不可变类型
- tuple, 不可变类型
- list
- set
- dict

```python
# 单行注释

"""
多行字符串/注释, 如果多行字符串有变量引用则作为字符串, 否则作为多行注释
"""
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

### 整数

- 小整数缓存（Python 会缓存 -5 到 256 的整数）

```python
>>> a = 5
>>> b = 5
>>> a is b
True
>>> c = 1000
>>> d = 1000
>>> c is d
False
```

### 枚举

- 一组绑定到唯一值的符号名称(成员)
- 可以迭代以按定义顺序返回其规范(非别名)成员
- 使用调用语法按值返回成员
- 使用索引语法按名称返回成员

```python
from enum import Enum

# class syntax
class Color(Enum):
    RED = 1
    GREEN = 2
    BLUE = 3

# functional syntax
Color = Enum('Color', [('RED', 1), ('GREEN', 2), ('BLUE', 3)])
print(Color['RED']) # <Color.RED: 1>
print(Color.RED.name)   # 'RED'
print(Color.RED.value)  # 1

print(list(Color))  # [<Color.RED: 1>, <Color.GREEN: 2>, <Color.BLUE: 3>]
print(tuple(Color)) # (<Color.RED: 1>, <Color.GREEN: 2>, <Color.BLUE: 3>)
print(set(Color))   # {<Color.GREEN: 2>, <Color.RED: 1>, <Color.BLUE: 3>}

print(Color.RED in Color)   # True
print(Color.RED.value in Color) # True
```

## 迭代器

- 可迭代对象实现了 `__iter__` 方法, 但不是迭代器

```python
my_list = [1, 2, 3]  # 列表是可迭代对象，但不是迭代器
print(hasattr(my_list, '__iter__'))    # True
print(hasattr(my_list, '__next__'))    # False
# 获取迭代器
list_iterator = iter(my_list)
print(hasattr(list_iterator, '__iter__'))  # True
print(hasattr(list_iterator, '__next__'))  # True
```

- 迭代器, 是一个可记忆遍历位置的对象, 实现了迭代器协议(`__iter__` 和 `__next__` 方法)
- 生成器, 是一种特殊的迭代器, 使用 `yield` 关键字创建, 可以暂停和恢复执行

```python
# 生成器函数
def countdown_gen(n):
    """生成器版本的倒计时"""
    print(f"开始倒计时 from {n}")
    while n > 0:
        yield n  # 暂停，返回 n
        n -= 1
    print("倒计时结束！")

# 创建生成器对象
gen = countdown_gen(5)
print(type(gen))  # <class 'generator'>
# 使用
for num in gen:
    print(num, end=' ')  # 5 4 3 2 1
```

## 推导式

推导式(Comprehensions) 是 python 中的一种简洁、高效的创建序列的语法结构，允许用一行代码从一个`可迭代对象`生成新的列表、字典、集合或生成器, 通常比传统的 for 循环更简洁

- 列表推导式

基本语法: [expression for item in iterable if condition]

```python
# 生成 0-9 的平方
>>> squares = [x**2 for x in range(10)]
>>> squares
[0, 1, 4, 9, 16, 25, 36, 49, 64, 81]
# 只保留偶数
>>> evens = [x for x in range(20) if x % 2 == 0]
>>> evens
[0, 2, 4, 6, 8, 10, 12, 14, 16, 18]

>>> numbers = [number-1 for number in range(1, 6)]
>>> numbers
[0, 1, 2, 3, 4]
# 带判断条件推导式
>>> numbers = [number for number in range(1, 6) if number % 2 == 1]  
>>> numbers
[1, 3, 5]

# 嵌套推导式
>>> rows = range(1, 4)                                             
>>> cols = range(1, 3) 
>>> cells = [(row, col) for row in rows for col in cols]
>>> cells
[(1, 1), (1, 2), (2, 1), (2, 2), (3, 1), (3, 2)]
```

- 字典推导式

基本语法: {key_expression: value_expression for item in iterable if condition}

```python
# 创建数字及平方的映射
>>> squares_dict = {x: x**2 for x in range(5)}
>>> squares_dict
{0: 0, 1: 1, 2: 4, 3: 9, 4: 16}
# 统计名字长度
>>> names = ['alice', 'tom', 'jerry', 'bob']     
>>> name_length = {name: len(name) for name in names}
>>> name_length
{'alice': 5, 'tom': 3, 'jerry': 5, 'bob': 3}
```

- 集合推导式

基本语法：{expression for item in iterable if condition}

```python
>>> numbers = [1, 2, 3, 4, 5, 6]
>>> numbers
[1, 2, 3, 4, 5, 6]
>>> square_set = {x**2 for x in numbers}
>>> square_set
{1, 4, 36, 9, 16, 25}
```

- 生成器推导式

基本语法：(expression for item in iterable if condition)

```python
>>> gen = (x**2 for x in range(5))
>>> next(gen)
0
>>> next(gen)
1
>>> next(gen)
4
>>> next(gen)
9
>>> next(gen)
16
```
## 上下文管理器

- 实现 `__enter__` 和 `__exit__` 魔法方法
- 使用 contextmanager 装饰器和 yield

### 异步上下文管理器

- 实现 `__aenter__` 和 `__aexit__` 魔法方法
- 使用 asynccontextmanager 装饰器 和 yield

- async for 用于迭代异步可迭代对象 (实现了 `__aiter__` 的对象), 只能用在 async 协程函数中, 否则报错
- async with 用于获取异步上下文管理器, 只能用在 async 协程函数中, 否则报错

## 异常

所有的异常类都是 Exception 的子类

- 使用 raise 抛出异常

```python
try:
# 可能发生错误的代码
except (NameError, ZeroDivisionError) e:
# 捕获未定义的变量或除以 0 的错误 
# except: # 等价于下一行的写法
except Exception as e:
# 捕获所有错误后执行的代码
else:
# 没有发生错误执行的代码
finally:
# 最后总是会执行的代码
```


## 模块

一个 python 文件, 可以包含 类，函数, 变量等

- `*` 导入全部 和 as 别名, 指定内容导入不能混合同时使用

`[from 模块名] import [模块 | 类 | 变量 | 函数 | *] [as 别名]`

## 包

分类管理多个模块

- `__init__.py` 文件定义目录为 python 包而不是普通的目录
- `__all__` 内置变量控制 `import *` 的导入行为

```python
# __init__.py
# 控制包级别 from package_name import * 的导入行为
__all__ = ['func_1']

def func_1():
    pass

def func_2():   # 不在 __all__ 中, 不允许被导入
    pass

# my_module1.py
# 控制模块级别的 import * 导入行为
__all__ = ['name']
```

### __init__.py

> python3.3 开始, 非必需 __init__.py 文件声明目录为 python 包

- 老版本声明目录为 python 包
- 定义初始化代码, 被导入时里面的代码会先运行
- 控制导入内容, 决定导入包时可以访问的内容
- 整理复杂项目, 帮助管理大项目中各个部分的关系
