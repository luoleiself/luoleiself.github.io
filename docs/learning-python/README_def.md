## 函数

```python
print(locals())  # locals 不在函数内使用时作用同 globals
# {'__name__': '__main__', '__doc__': None, '__package__': None, '__loader__': <_frozen_importlib_external.SourceFileLoader object at 0x000002754662BE00>, '__spec__': None, '__annotations__': {}, '__builtins__': <module 'builtins' (built-in)>, 'platform': <module 'platform' from 'D:\\uv\\python_install\\cpython-3.12.11-windows-x86_64-none\\Lib\\platform.py'>, 'sys': <module 'sys' (built-in)>, 'original_ps1': '>>> ', 'is_wsl': False, 'REPLHooks': <class '__main__.REPLHooks'>, 'get_last_command': <function get_last_command at 0x00000275468E36A0>, 'PS1': <class '__main__.PS1'>}

animal = 'fruitbat'


def change_and_print_global():
    print('inside change_and_print_globa', animal)  # 变量先定义后使用
    animal = 'wombat'
    print('after the change', animal)


change_and_print_global()
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
#   File "<stdin>", line 2, in change_and_print_global
# UnboundLocalError: cannot access local variable 'animal' where it is not associated with a value
```

### 参数

参数定义顺序: 位置参数 > 默认参数 > *args > 关键字参数 > **kwargs

```python
# 系统函数的定义
# print(*objects, sep=' ', end='\n', file=None, flush=False)

# sorted(iterable, /, *, key=None, reverse=False)
```

- `/` 权限位置参数分隔符, 之前的参数必须使用位置参数传入, python 3.8 支持

```python
def func_5(name, age, /, city='Beijing'):
    print(name, age, city)


func_5('Bob', 20)  # 位置参数和默认参数
# Bob 20 Beijing
func_5('Bob', 20, 'Shanghai')  # 位置参数
# Bob 20 Shanghai
func_5('Bob', 20, city='Chongqing')  # 位置参数和关键字参数
# Bob 20 Chongqing

func_5('Bob', age=20)  # / 之前必须使用位置参数传入
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: func_5() got some positional-only arguments passed as keyword arguments: 'age'
```

- `*` 权限关键字参数分隔符, 之后的参数必须使用关键字参数传入

```python
def func_6(name, *, age, city='Beijing'):
    print(name, age, city)


func_6('Tom', age=99)  # 位置参数和关键字参数, 默认参数
# Tom 99 Beijing
func_6('Tom', age=99, city='London')  # * 位置参数和关键字参数
# Tom 99 London
func_6(name='Tom', age=99, city='London')  # 关键字参数
# Tom 99 London

func_6('Tom', 99, 'London')  # * 之后的参数必须使用关键字参数传入
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: func_6() takes 1 positional argument but 3 were given
func_6('Tom', 99, city='London')  # * 之后的参数必须使用关键字参数传入
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: func_6() takes 1 positional argument but 2 positional arguments (and 1 keyword-only argument) were given
```

#### 可变参数

只能用在 `函数定义` 或 `函数调用`, 单独作为语句使用时报错

- `*` 可变位置参数
    - 函数定义, 接收位置参数之后剩余的参数保存为元组
    - 函数调用, 解构 `可迭代对象` 作为参数传入函数

- `**` 可变关键字参数
    - 函数定义, 接收关键字参数之后剩余的关键字参数保存为字典
    - 函数调用, 解构字典转换为 `key=value` 格式作为参数传入函数

```python
def func_4(*args, **kwargs):
    print(f'接收的参数 {args}\t类型为 {type(args)}')
    print(f'接收的参数 {kwargs}\t类型为 {type(kwargs)}')


func_4('hello', name='Tom', age=18, addr='beijing')  # str
# 接收的参数 ('hello',)   类型为 <class 'tuple'>
# 接收的参数 {'name': 'Tom', 'age': 18, 'addr': 'beijing'}        类型为 <class 'dict'>
func_4(*'hello', name='Tom', age=18, addr='beijing')  # 解构 str
# 接收的参数 ('h', 'e', 'l', 'l', 'o')    类型为 <class 'tuple'>
# 接收的参数 {'name': 'Tom', 'age': 18, 'addr': 'beijing'}        类型为 <class 'dict'>
func_4(*range(9, 13), **{'name': 'Jerry', 'age': 8, 'addr': 'shanghai'})  # 解构数字序列
# 接收的参数 (9, 10, 11, 12)      类型为 <class 'tuple'>
# 接收的参数 {'name': 'Jerry', 'age': 8, 'addr': 'shanghai'}      类型为 <class 'dict'>

func_4(['a', 'b', 'c'])  # list
# 接收的参数 (['a', 'b', 'c'],)   类型为 <class 'tuple'>
# 接收的参数 {}   类型为 <class 'dict'>
func_4(*['a', 'b', 'c'])  # 解构 list
# 接收的参数 ('a', 'b', 'c')      类型为 <class 'tuple'>
# 接收的参数 {}   类型为 <class 'dict'>
func_4({1, 2}, **dict((('a', 'b'), ('c', 'd'))))  # set, 解构包含双项序列的序列转换后的字典
# 接收的参数 ({1, 2},)    类型为 <class 'tuple'>
# 接收的参数 {'a': 'b', 'c': 'd'} 类型为 <class 'dict'>
func_4(*{1, 2})  # 解构 set                  
# 接收的参数 (1, 2)       类型为 <class 'tuple'>
# 接收的参数 {}   类型为 <class 'dict'>
func_4({'x': 'X', 'y': 'Y'})  # dict
# 接收的参数 ({'x': 'X', 'y': 'Y'},)      类型为 <class 'tuple'>
# 接收的参数 {}   类型为 <class 'dict'>
func_4(*{'x': 'X', 'y': 'Y'})  # 解构 dict
# 接收的参数 ('x', 'y')   类型为 <class 'tuple'>
# 接收的参数 {}   类型为 <class 'dict'>

# 单独作为语句使用报错
te = (1, 2, 3, 4)
# *te
#   File "<stdin>", line 1
# SyntaxError: can't use starred expression here
# **{'a': 'A'}
#   File "<stdin>", line 1
#     **dt
#     ^^
# SyntaxError: invalid syntax
```

### 装饰器

装饰器是一种函数，接受一个函数作为输入并返回另一个函数

- @functools.wraps(func) 在闭包函数上添加装饰器, 保留被装饰函数的元信息

```python
import functools


def my_decorator(func):
    @functools.wraps(func)  # 保留被装饰函数的元信息
    def new_func(*args, **kwargs):
        result = func(*args, **kwargs)
        return result

    return new_func


@my_decorator
def add_int(a, b):
    return a + b


# 输出 add_int 的函数名, 没有 wraps 装饰器将输出 new_func
print(f'add_int.__name__ {add_int.__name__}')
```

### 文档注释

函数体的顶部定义的多行字符串
