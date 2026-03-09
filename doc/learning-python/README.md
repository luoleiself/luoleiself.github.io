
项目结构

```python
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

- 序列: 连续的可用下标访问的数据存储结构, 字符串，元组，列表, 可使用`切片`进行操作
- `*` | `**`: 在数据存储结构中使用时, 解构 `可迭代对象`
- `__` 双下划线起止的名称保留给 python 内部使用

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
- rasie, 抛出异常

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

'''
多行字符串/注释, 如果多行字符串有变量引用则作为字符串, 否则作为多行注释
'''

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

range() 是`惰性求值`的，不会一次性占用大量内存

```python
# 立即返回，几乎不占内存
big_range = range(1000000)
# 创建包含 100 万个元素的列表，占用大量内存
big_list = list(range(1000000))
```

- range() 生成数字序列，返回一个可迭代对象, 不能直接使用, 需要`配合其他函数或语句`才能获取值
  - 只有 1 个参数时生成 0 到 arg 的数字序列
  - start, 开始，包含
  - end, 结束，不包含
  - step，步长绝对值，默认为 1
    - 如果为 -1, `反向生成数字序列`

```python
>>> range(4)  # 直接使用不会输出结果
range(0, 4)
>>> print(range(4)) # 直接使用不会输出结果
range(0, 4)

>>> num = range(4)
>>> print(num[2])
2
>>> list(range(4))
[0, 1, 2, 3]
>>> list(range(1, 4)) 
[1, 2, 3]
>>> list(range(1, 4, 2))
[1, 3]
>>> tuple(range(4))
(0, 1, 2, 3)
>>> tuple(range(1, 4))
(1, 2, 3)
>>> tuple(range(1, 4, 2))
(1, 3)

# 反向生成数字序列
>>> list(range(4, -2, -1))
[4, 3, 2, 1, 0, -1]
>>> list(range(4, -2, -2))
[4, 2, 0]
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

## 内置属性

- `__name__` 用于判断当前模块是被直接运行还是被导入, 直接运行时为 `__main__` 否则值为所在的模块名
- `__all__` 是一个列表, 用于定义当使用 `from module_name import *` 时, 哪些名称可以被导入
  - 在 `__init__.py` 文件中使用限制包级的导入行为
  - 在 模块中 使用时限制模块的导入行为
- `__file__` 当前文件的路径
- `__doc__` 文档字符串
- `__package__` 包名
- `__dict__` 对象的属性字典
- `__slots__` 限制属性(内存优化)
- `__version__` 版本号
- `__author__`  作者

## 内置函数

### tuple()

创建或转换其他类型为元组, 参数为空或`可迭代对象`

### list()

创建或转换其他类型为列表, 参数为空或`可迭代对象`

### dict()

创建或转换其他类型为字典, 参数为空, `具名参数`, `包含双项序列的任意序列`

### set()

创建或转换其他类型为集合, 参数为空或`可迭代对象`

### int()

将任意值转换成整型化值, 参数必须是字符串或者真实的数值

- 第一个参数为要转换的值
  - 浮点数则截断整数部分
  - 参数如果是字符串, 只能由 `数字, +, -, _` 组成
  - 布尔值 True 转为 1, False 转为 0
- 第二个参数为进制数位

### float()

将任意值转换为浮点化值, 参数必须是字符串或者真实的数值

- 参数为整型则转换为浮点数
- 参数为字符串的浮点数表示形式，支持科学计数格式
- 布尔值 True 转为 1.0, False 转为 0.0

### bool()

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

>>> bool({})  # 空字典转换为 False
False
>>> bool(set()) # 空集合转换为 False
False
>>> bool(tuple()) # 空元组转换为 False
False
>>> bool([])  # 空列表转换为 False   
False
```

### str()

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

>>> str([1, 2, 3, 4]) # list 转换为字符串
'[1, 2, 3, 4]'
>>> str(('a', 'b',))  # tuple 转换为字符串
"('a', 'b')"
>>> str({1, 2, 3, 4}) # set 转换为字符串
'{1, 2, 3, 4}'
>>> str({'a': 'A', 'b': 'B'}) # dict 转换为字符串
"{'a': 'A', 'b': 'B'}"
```

### chr()/ord()

转换 unicode 码对应的字符

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

### max()/min()

获取容器中的最大值或最小值

- 字典获取所有键中的最大/最小值

```python
>>> max('abcdefg')
'g'
>>> max([1, 2, 3, 4, 5])
5
>>> max((1, 2, 3, 4, 5，))
5
>>> max({1, 2, 3, 4, 5})
5
>>> max({'a': 'A', 'b': 'B', 'c': 'C', 'd': 'D'})
'd'

>>> min('abcdefg')
'a'
>>> min([1, 2, 3, 4, 5])
1
>>> min((1, 2, 3, 4, 5,))
1
>>> min({1, 2, 3, 4, 5}) 
1
>>> min({'a': 'A', 'b': 'B', 'c': 'C', 'd': 'D'})
'a'
```

### id() 获取变量引用的对象的 id

```python
>>> n = 1
>>> id(n)
140726623005112
>>> m = n
>>> id(m)
140726623005112
>>> n = 2
>>> id(n)
140726623005144
>>> id(m)
140726623005112

>>> names = ['alice', 'tom', 'jerry']
>>> id(names)
2970721128832
>>> names_lst = names
>>> id(names_lst)
2970721128832
>>> names.append('bob')
>>> names
['alice', 'tom', 'jerry', 'bob']
>>> names_lst
['alice', 'tom', 'jerry', 'bob']
```

### locals() 返回函数内部的内容目录

- 在全局作用域内使用时获取的结果和 `globals()` 相同

### globals() 返回当前全局作用域的内容目录

- `__name__` 属性在只有当前文件直接运行时, `__name__` 属性值才会是 `__main__`，否则被作为模块导入时, 属性值为模块名

```python
>>> print(globals()) 
{'__name__': '__main__', '__doc__': None, '__package__': None, '__loader__': <_frozen_importlib_external.SourceFileLoader object at 0x0000014B4D70BE00>, '__spec__': None, '__annotations__': {}, '__builtins__': <module 'builtins' (built-in)>, 'platform': <module 'platform' from 'D:\\uv\\python_install\\cpython-3.12.11-windows-x86_64-none\\Lib\\platform.py'>, 'sys': <module 'sys' (built-in)>, 'original_ps1': '>>> ', 'is_wsl': False, 'REPLHooks': <class '__main__.REPLHooks'>, 'get_last_command': <function get_last_command at 0x0000014B4D9C36A0>, 'PS1': <class '__main__.PS1'>, 'name': 'global variable', 'new_func': <function new_func at 0x0000014B4D9E20C0>, 'animal': 'wom', 'change_and_print_global': <function change_and_print_global at 0x0000014B4DA05EE0>, 'change_global': <function change_global at 0x0000014B4DA05D00>}
```

### property()

用于创建和管理类的属性, 允许将方法当作属性访问, 从而实现对属性的控制

`property(fget=None, fset=None, fdel=None, doc=None)`

```python
class Person:
    def __init__(self, name, age):
        # 如果知道了 hidden_name 和 hidden_age, 仍然可以直接操作它们
        self.hidden_name = name
        self.hidden_age = age

    # def 定义 name 和 age 的 getter 和 setter 方法

    # property 内置函数将方法当作属性访问
    name = property(get_name, set_name)
    age = property(get_age, set_age)
    

p = Person('xiaoming', 19)
print(f'p.name {p.name} p.age {p.age}')
# 仍然可以访问到隐藏的属性
print(f'p.hidden_name {p.hidden_name} p.hidden_age {p.hidden_age}')
```
