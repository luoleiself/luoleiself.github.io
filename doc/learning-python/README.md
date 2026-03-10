
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

## 迭代器

可迭代对象实现了 `__iter__` 方法, 但不是迭代器

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
- 生成器, 是一种特殊的迭代器, 使用 yield 关键字创建, 可以暂停和恢复执行

```python
# 自定义迭代器
class CountDown:
    """倒计时迭代器"""
    def __init__(self, start):
        self.current = start
    
    def __iter__(self):
        return self  # 返回迭代器本身
    
    def __next__(self):
        if self.current <= 0:
            raise StopIteration  # 结束迭代
        value = self.current
        self.current -= 1
        return value

# 使用自定义迭代器
counter = CountDown(5)
for num in counter:
    print(num, end=' ')  # 输出: 5 4 3 2 1

# 生成器
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

## 内置属性(魔术属性)

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
