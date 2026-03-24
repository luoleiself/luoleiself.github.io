## 函数

```python
>>> print(locals()) # locals 不在函数内使用时作用同 globals
{'__name__': '__main__', '__doc__': None, '__package__': None, '__loader__': <_frozen_importlib_external.SourceFileLoader object at 0x000002754662BE00>, '__spec__': None, '__annotations__': {}, '__builtins__': <module 'builtins' (built-in)>, 'platform': <module 'platform' from 'D:\\uv\\python_install\\cpython-3.12.11-windows-x86_64-none\\Lib\\platform.py'>, 'sys': <module 'sys' (built-in)>, 'original_ps1': '>>> ', 'is_wsl': False, 'REPLHooks': <class '__main__.REPLHooks'>, 'get_last_command': <function get_last_command at 0x00000275468E36A0>, 'PS1': <class '__main__.PS1'>}

>>> animal = 'fruitbat'
>>> def change_and_print_global():
...     print('inside change_and_print_globa', animal)  # 访问未定义的变量报错
...     animal = 'wombat'
...     print('after the change', animal)
... 
>>> change_and_print_global()
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
  File "<stdin>", line 2, in change_and_print_global
UnboundLocalError: cannot access local variable 'animal' where it is not associated with a value
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
>>> def func_5(name, age, /, city='Beijing'): 
...     print(name, age, city)
...
>>> func_5('Bob', 20) # 位置参数和默认参数
Bob 20 Beijing
>>> func_5('Bob', 20, 'Shanghai') # 位置参数
Bob 20 Shanghai
>>> func_5('Bob', 20, city='Chongqing') # 位置参数和关键字参数
Bob 20 Chongqing

>>> func_5(name='Bob', age=20)  # / 之前必须使用位置参数传入
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: func_5() got some positional-only arguments passed as keyword arguments: 'name, age'
>>> func_5('Bob', age=20) # / 之前必须使用位置参数传入
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: func_5() got some positional-only arguments passed as keyword arguments: 'age'
```

- `*` 权限关键字参数分隔符, 之后的参数必须使用关键字参数传入

```python
>>> def func_6(name, *, age, city='Beijing'):
...     print(name, age, city)
... 
>>> func_6('Tom', age=99) # 位置参数和关键字参数, 默认参数
Tom 99 Beijing
>>> func_6('Tom', age=99, city='London')  # * 位置参数和关键字参数
Tom 99 London
>>> func_6(name='Tom', age=99, city='London') # 关键字参数
Tom 99 London

>>> func_6('Tom', 99, 'London') # * 之后的参数必须使用关键字参数传入
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: func_6() takes 1 positional argument but 3 were given
>>> func_6('Tom', 99, city='London')  # * 之后的参数必须使用关键字参数传入
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: func_6() takes 1 positional argument but 2 positional arguments (and 1 keyword-only argument) were given
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
>>> def func_4(*args, **kwargs):
...     print(f'接收的参数 {args}\t类型为 {type(args)}')
...     print(f'接收的参数 {kwargs}\t类型为 {type(kwargs)}')
...
>>> func_4('hello', name='Tom', age=18, addr='beijing') # str
接收的参数 ('hello',)   类型为 <class 'tuple'>
接收的参数 {'name': 'Tom', 'age': 18, 'addr': 'beijing'}        类型为 <class 'dict'>
>>> func_4(*'hello', name='Tom', age=18, addr='beijing')  # 解构 str
接收的参数 ('h', 'e', 'l', 'l', 'o')    类型为 <class 'tuple'>
接收的参数 {'name': 'Tom', 'age': 18, 'addr': 'beijing'}        类型为 <class 'dict'>
>>> func_4(*range(9, 13), **{'name': 'Jerry', 'age': 8, 'addr': 'shanghai'})  # 解构数字序列
接收的参数 (9, 10, 11, 12)      类型为 <class 'tuple'>
接收的参数 {'name': 'Jerry', 'age': 8, 'addr': 'shanghai'}      类型为 <class 'dict'>

>>> func_4(['a', 'b', 'c']) # list
接收的参数 (['a', 'b', 'c'],)   类型为 <class 'tuple'>
接收的参数 {}   类型为 <class 'dict'>
>>> func_4(*['a', 'b', 'c'])  # 解构 list
接收的参数 ('a', 'b', 'c')      类型为 <class 'tuple'>
接收的参数 {}   类型为 <class 'dict'>
>>> func_4({1, 2}, **dict((('a', 'b'), ('c', 'd'))))  # set, 解构包含双项序列的序列转换后的字典
接收的参数 ({1, 2},)    类型为 <class 'tuple'>
接收的参数 {'a': 'b', 'c': 'd'} 类型为 <class 'dict'>
>>> func_4(*{1, 2}) # 解构 set                  
接收的参数 (1, 2)       类型为 <class 'tuple'>
接收的参数 {}   类型为 <class 'dict'>
>>> func_4({'x': 'X', 'y': 'Y'})  # dict
接收的参数 ({'x': 'X', 'y': 'Y'},)      类型为 <class 'tuple'>
接收的参数 {}   类型为 <class 'dict'>
>>> func_4(*{'x': 'X', 'y': 'Y'}) # 解构 dict
接收的参数 ('x', 'y')   类型为 <class 'tuple'>
接收的参数 {}   类型为 <class 'dict'>

# 单独作为语句使用报错
>>> te = (1, 2, 3, 4)
>>> *te
  File "<stdin>", line 1
SyntaxError: can't use starred expression here
>>> **{'a': 'A'}
  File "<stdin>", line 1
    **dt
    ^^
SyntaxError: invalid syntax
```

### 装饰器

装饰器是一种函数，接受一个函数作为输入并返回另一个函数

- @functools.wraps(func) 在闭包函数上添加装饰器, 保留被装饰函数的元信息

```python
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

## 类

- `__new__` 创建并返回一个新的实例, 第一个参数为 cls 类本身
  - 必须返回实例, 如果不返回实例则不会调用 `__init__` 方法
  - 在 `__init__` 方法执行之前调用

```python
class Animal:
    _instance = None

    # 先调用
    # 实现单例模式
    def __new__(cls, *args, **kwargs):
        print('Animal __new__...')
        if cls._instance is None:
            print('创建唯一实例')
            cls._instance = super().__new__(cls)
        return cls._instance

    # 初始化实例
    def __init__(self, *args, **kwargs):
        print('Animal __init__...')
        self.args = args
        self.kwargs = kwargs


ani = Animal(1, 2, 3, x=10, y=20)
print(f'ani = {ani}')
# Animal __new__...
# Animal __init__...
# ani = <__main__.Animal object at 0x0000023215769D90>
```

- `__init__` 非必需的初始化实例属性的方法, 第一个参数为 self 实例本身
  - 没有返回值或者返回 None
  - 如果定义了初始化方法和参数, 则必须传入参数否则报错
  - 如果子类没有定义初始化方法, python 自动调用父类的初始化方法完成属性绑定
  - 如果实例化未定义初始化方法的子类时, 参数必须符合父类初始化方法参数的要求

```python
>>> class Cat():
...     def __init__(self, name):
...             self.name = name
... 
>>> another_cat = Cat('green miao') # 如果定义了初始化方法和参数则必须传入
>>> another_cat.name
'green miao'
>>> a_cat = Cat()
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: Cat.__init__() missing 1 required positional argument: 'name'

# 继承 - 初始化
class Quote:
    def __init__(self, person, words):
        self.person = person
        self.words = words

    def who(self):
        return self.person

    def says(self):
        return self.words+'.'


class QuestionQuote(Quote):
    def says(self):
        return self.words+'?'


class ExclamationQuote(Quote):
    def says(self):
        return self.words+'!'


class SubQuestionQuote(QuestionQuote):
    def says(self):
        return self.words+'??'


def who_says(obj):
    print(f'{obj.who()} says: {obj.says()}')


hunter = Quote('Elmer Fudd', 'I\'m hunting wabbits')
# 如果子类没有定义初始化方法, python 自动调用父类的初始化方法完成属性绑定
hunter1 = QuestionQuote('Bugs Bunny', 'What\'s up, doc')
hunter2 = ExclamationQuote('Daffy Duck', 'It\'s rabbit season')
who_says(hunter)
who_says(hunter1)
who_says(hunter2)
print('-----------')
# TypeError: Quote.__init__() takes 3 positional arguments but 4 were given
hunter3 = SubQuestionQuote('Jerry', 'It\'s Tom', 'good')
who_says(hunter3)
```

### 自省

通过一定的机制查询到对象的内部解构

### 内置属性(魔法属性)

- `__name__` 标识函数、模块、包的名称
  - 直接运行文件时为 `__main__` 否则值为所在的模块名
- `__all__` 是一个列表, 用于定义当使用 `from module_name import *` 时, 哪些名称可以被导入
  - 在 `__init__.py` 文件中使用限制包级的导入行为
  - 在 模块中 使用时限制模块的导入行为
- `__file__` 当前文件的路径
- `__doc__` 文档字符串
- `__package__` 包名
- `__slots__` 限制属性(内存优化)
- `__version__` 版本号
- `__author__`  作者

- `__dict__` 存储对象的命名空间
  - 实例: 实例对象自身的属性, 不包含类属性、方法，可读写
  - 类: 类属性、方法、装饰器描述符等, 不能直接修改
  - 模块: 模块中定义的所有名称(变量, 函数， 类, 导入的模块等)

```python
# __dict__ 存储对象的命名空间
class Person:
    species = 'Human'

    def __init__(self, name):
        self.name = name

    @property
    def new_name(self):
        return self.name

    @classmethod
    def cls_method(cls):
        print(f'class Person method...')

    @staticmethod
    def static_method():
        print(f'static method...')

    def say_hello(self):
        print(f'{self.name} say hello!')


p = Person('Jerry')
print(f'p.__dict__ {p.__dict__}') # 实例自身的属性
p.say_hello()
print(f'Person.__dict__ {Person.__dict__}') # 类属性、方法、装饰器描述符等
print(f'moduleName.__dict__ {select.__dict__}') # 模块中定义的所有名称
print('------------------')

p.__dict__ {'name': 'Jerry'}
Jerry say hello!
Person.__dict__ {'__module__': '__main__', 'species': 'Human', '__init__': <function Person.__init__ at 0x000001BC44D95120>, 'new_name': <property object at 0x000001BC44D98540>, 'cls_method': <classmethod(<function Person.cls_method at 0x000001BC44D95260>)>, 'static_method': <staticmethod(<function Person.static_method at 0x000001BC44D95300>)>, 'say_hello': <function Person.say_hello at 0x000001BC44D953A0>, '__dict__': <attribute '__dict__' of 'Person' objects>, '__weakref__': <attribute '__weakref__' of 'Person' objects>, '__doc__': None}
moduleName.__dict__ {'__name__': 'select', '__doc__': 'This module supports asynchronous I/O on multiple file descriptors.\n\n*** IMPORTANT NOTICE ***\nOn Windows, only sockets are supported; on Unix, all file descriptors.', '__package__': '', '__loader__': <_frozen_importlib_external.ExtensionFileLoader object at 0x000001989F2D7260>, '__spec__': ModuleSpec(name='select', loader=<_frozen_importlib_external.ExtensionFileLoader object at 0x000001989F2D7260>, origin='D:\\uv\\python_install\\cpython-3.12.11-windows-x86_64-none\\DLLs\\select.pyd'), 'select': <built-in function select>, '__file__': 'D:\\uv\\python_install\\cpython-3.12.11-windows-x86_64-none\\DLLs\\select.pyd', 'error': <class 'OSError'>}
```

- `__class__` 标识创建实例对象的类
- `__bases__` 标识类对象继承的父类

```python
>>> class Person:
...     pass
...
>>> p = Person()
>>> p.__class__
<class '__main__.Person'>
>>> p.__bases__ # 实例对象没有 __bases__ 属性
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
AttributeError: 'Person' object has no attribute '__bases__'. Did you mean: '__class__'?
>>> Person.__class__
<class 'type'>
>>> Person.__bases__
(<class 'object'>,)
```

### 内置方法(魔法方法)

比较方法

- `__str__` 自定义返回实例的结果, 默认输出实例的内存地址
- `__lt__` 自定义实例的 < 比较, 返回 bool
- `__le__` 自定义实例的 <= 比较, 返回 bool
- `__gt__` 自定义实例的 >比较, 返回 bool
- `__ge__` 自定义实例的 >= 比较, 返回 bool
- `__eq__`  自定义实例的 == 比较, 返回 bool
- `__ne__`  自定义实例的 != 比较, 返回 bool

数学运算方法

- `__add__` 自定义实例的加法运算
- `__sub__` 自定义实例的减法运算
- `__mul__` 自定义实例的乘法运算
- `__floordiv__`  自定义实例的整除运算
- `__truediv__` 自定义实例的除法运算
- `__mod__` 自定义实例的取模运算
- `__pow__` 自定义实例的幂运算

- `__str__` 自定义实例的 print 函数的输出结果
- `__repr__`  自定义实例的回显
- `__len__` 自定义实例的获取长度

```python
class Word:
    def __init__(self, text):
        self.text = text

    def __eq__(self, other):
        return self.text.lower() == other.text.lower()

    def __str__(self):
        return f'python forever: {self.text}'

    def __repr__(self):
        return 'Word("' + self.text + '")'

    def __mul__(self, other):
        return self.text * other.text


first = Word('hello')
first # 调用 __repr__ 内置方法
print(first)  # 调用 __str__ 内置方法
second = Word('HELLO')
print(f'first == second {first == second}') # 比较两个实例

n1 = Word(4)
n2 = Word(5)
print(f'n1 * n2 {n1 * n2}') # 计算两个实例的乘法
```

数据描述符的优先级最高, 给实例的不存在的属性赋值时不会直接存入实例的 `__dict__` 中

- `__get__`
- `__set__`
- `__delete__`

```python
class IntField:
    def __get__(self, instance, owner):
        return self.value
    def __set__(self, instance, value):
        if not isinstance(value, numbers.Integral):
            raise ValueError('int value need')
        if value < 0:
            raise ValueError('positive value need')
        self.value = value
    def __delete__(self, instance):
        pass

class User:
    age = IntField()  # 数据描述符
```

上下文管理器

- `__enter__` with 语句绑定这个方法返回的结果到 as 子句中指定的目标
- `__exit__`  退出与此对象相关的运行时上下文, 可以有返回值
  - 如果返回值可以被 `bool()` 转换为 True, 表示异常已被处理, 程序可以继续执行
  - 如果返回值可以被 `bool()` 转换为 False, 表示异常没有被处理, 程序将会中断

```python
>>> class Sample:
...     def __enter__(self):
...             print('__enter__...')
...             return self
...     def __exit__(self, exec_type, exec_val, traceback):
...             print('__exit__...')
...     def do_something(self):
...             print('doing')
... 
>>> with Sample() as s:
...     s.do_something()
... 
__enter__...
doing
__exit__...

# 使用装饰器和生成器
>>> import contextlib
>>> @contextlib.contextmanager
... def f_open(file_name):
...     print('start...')   # acquire resource
...     yield {}
...     print('end...') # release resource
... 
>>> with f_open('hello world') as f_o:
...     print('with start...')
...                           
start...
with start...
end..
```

迭代器方法

- `__iter__` 自定义实例返回迭代器对象
- `__next__`  自定义实例返回下一个值
- `__getitem__` 用于支持对象通过索引或键进行访问, obj\[key\] 或 obj\[index\]
  - 如果未实现 `__iter__` 方法时, 解释器将尝试通过此方法从索引 0 开始依次获取元素进行迭代
  - 使用切片时, 需要判断 key 是否为 slice 类型

```python
# 支持切片
class MyList:
    def __init__(self, data):
        self.data = data
    
    def __getitem__(self, key):
        # 判断 key 是否是切片类型
        if isinstance(key, slice):
            # indices 根据给定的长度, 计算切片边界并返回 start, end, step 元组
            return [self.data[i] for i in range(*key.indices(len(self.data)))]
        return self.data[key]

a = MyList([1, 2, 3, 4])
print(a[1:3])   # [2, 3]
```

- `__getattr__` 未找到属性时调用此方法
- `__getattribute__` 每次访问属性时都调用此方法, 优先级高于 `__getattr__`

```python
# __getattr__ 和 __getattribute__
class User:
    def __init__(self, name, info={}):
        self.name = name
        self.info = info

    # 未找到属性时调用此方法
    def __getattr__(self, item):
        return self.info[item]


print('__getattr__: 未找到属性时调用此方法')
u1 = User('Tom', {'age': 18, 'sex': 'male'})
print(f'u1.name {u1.name}') # u1.name Tom
print(f'u1.age {u1.age}') # u1.age 18
print(f'u1.sex {u1.sex}') # u1.sex male
print('---------')


class User2:
    def __init__(self, name, info={}):
        self.name = name
        self.info = info

    # 每次访问属性时都调用此方法
    def __getattribute__(self, item):
        return 'hello world'


print('__getattribute__: 每次访问属性时都调用此方法')
u2 = User2('Jerry', {'age': 20, 'sex': 'female'})
print(f'u2.name {u2.name}') # u2.name hello world
print(f'u2.age {u2.age}') # u2.age  hello world
print(f'u2.sex {u2.sex}') # u2.sex  hello world
```

协程方法

- `__aiter__` 异步可迭代对象
- `__anext__` 异步迭代器对象  
- `__aenter__` 绑定异步上下文中返回的结果到 as 子句中指定的目标
- `__aexit__` 退出与此对象相关的运行时异步上下文

其他方法

- `__call__` 自定义实例支持函数调用方式

### 私有属性

python 没有私有属性

- 使用 getter 和 setter, 仍然可以操作隐藏属性

```python
def get_name(self):
    return self.hidden_name
  
def set_name(self, value):
    self.hidden_name = value
```

- 使用内置函数 property() 允许将方法当作属性访问
- 使用装饰器 @property

```python
# 使用 property 内置函数
class Person:
    # ...
    name = property(get_name, set_name)

# 使用装饰器
@property
def name(self):
    return self.hidden_name

@name.setter
def name(self, value):
    self.hidden_name = value

# 只读属性
@property
def diameter(self):
    return 2 * self.radius
```

- 使用 `__` 起始定义内置属性, python 将内置属性增加 `_类名` 的前缀绑定到实例上
  - 单下划线开始的属性和方法约定私有, 但实际上是公开的, 子类可以重写
  - 双下划线开始的属性和方法名称修饰, 真正私有, 不想被子类意外重写
  - 双下划线起止的属性和方法内置(魔法), 不要自己创建

```python
class Duck:
    def __init__(self, name):
        self.__name = name
    
    @property
    def name(self):
      return self.__name

    @name.setter
    def name(self, value)
      self.__name = value

# 实例化后内置属性 __name 将变为 _Duck__name 绑定到实例上
```

### 类属性, 类方法和静态方法

- 类属性(成员属性)可以被继承, 修改类属性后只会影响实例未修改过的同名属性和新创建的实例
  - 实例修改和类同名的属性将覆盖类属性

```python
class Fruit:
    color = 'red'  # 类属性，所有实例共享

banana = Fruit()
print(f'Fruit.color {Fruit.color}')  # red
print(f'banana.color {banana.color}')  # red
print('---------')
print(f'修改 Fruit.color = "yellow"')
Fruit.color = 'yellow'
print(f'Fruit.color {Fruit.color}')  # yellow
print(f'banana.color {banana.color}')  # yellow, 实例未修改过的同名属性也会变化
print('---------')
print(f'修改 banana.color = "green"')
banana.color = 'green'
print(f'Fruit.color {Fruit.color}')  # yellow
print(f'banana.color {banana.color}')  # green
print('---------')
print(f'修改 Fruit.color = "blue"')
Fruit.color = 'blue'
print(f'Fruit.color {Fruit.color}')  # blue
print(f'banana.color {banana.color}')  # green
print('---------')
orange = Fruit()
print(f'Fruit.color {Fruit.color}')  # blue
print(f'orange.color {orange.color}')  # blue
```

- 类方法(成员方法)，使用 @classmethod 装饰器, 方法的第一个参数为类本身 cls, 所有实例共享
- 静态方法，使用 @staticmethod 装饰器, 不需要 self 和 cls 参数, 不需要实例化可以直接调用

```python
# 类方法
class A():
    count = 0
    def __init__(self):
        A.count += 1

    # 类方法, 所有实例共享
    @classmethod
    def toString(cls):
        print(f'class A has {cls.count} instances...')

class AA(A):
    pass

a = A()
b = A()
aa = AA()
AA.toString()

# 静态方法
class C:
    @staticmethod
    def go_home():
        print(f'class c static method go_home...')


C.go_home() # 不需要实例化直接调用
```

## 多继承

- 访问自己不存在的属性或方法时, 优先使用最先继承的父类的属性和方法
  - 搜索过程从左往右搜索 `__mro__` 列表, 按顺序优先匹配

菱形继承: A -> B -> C -> D

```python
class D:
    name = 'cls_D'


class B(D):
    # name = 'cls_B'
    pass


class C(D):
    name = 'cls_C'


class A(B, C):
    # name = 'cls_A'
    pass


a = A()
print(f'a.name {a.name}\nA.__mro__ {A.__mro__}')
# a.name cls_C
# A.__mro__ (<class '__main__.A'>, <class '__main__.B'>, <class '__main__.C'>, <class '__main__.D'>, <class 'object'>)
```

树形继承: A -> B -> D
           -> C -> E

```python
class E:
    name = 'cls_E'


class D:
    name = 'cls_D'


class B(D):
    # name = 'cls_B'
    pass


class C(E):
    name = 'cls_C'


class A(B, C):
    # name = 'cls_A'
    pass


a = A()
print(f'a.name {a.name}\nA.__mro__ {A.__mro__}')
# a.name cls_D
# A.__mro__ (<class '__main__.A'>, <class '__main__.B'>, <class '__main__.D'>, <class '__main__.C'>, <class '__main__.E'>, <class 'object'>)
```

### 多态

## 元类

metaclass 是创建类的类, 元类定义了类的行为

```python
def __init__(self, name):
    self.name = name

def greet(self):
    return f'Hello, I\'m {self.name}'

# type 创建 Person 类
Person = type('Person', (), {
    'species': 'human',
    '__init__': __init__,
    'greet': greet
})

p2 = Person('Tom')
p2.greet()
print('---------------')
# 函数创建类
def UpperMetaClass(class_name, class_bases, class_attrs):
    # 字典推导式将类属性非下划线开头的转换为大写
    new_attrs = dict((key, value) if key.startswith('_') else (
        key.upper(), value) for key, value in class_attrs.items())
    return type(class_name, class_bases, new_attrs)

# 元类创建类
class UpperMetaClass(type):
    def __new__(cls, class_name, class_bases, class_attrs):
        # 字典推导式将类属性非下划线开头的转换为大写
        new_attrs = dict((key, value) if key.startswith('_') else (
            key.upper(), value) for key, value in class_attrs.items())
        return super().__new__(cls, class_name, class_bases, new_attrs)

class Person(object, metaclass=UpperMetaClass):
    name = 'Tom'
    _age = 18


print(f'hasattr(Person, name): {hasattr(Person, "name")}')  # False
print(f'hasattr(Person, "NAME"): {hasattr(Person, "NAME")}')    # True
print(f'hasattr(Person, "_AGE"): {hasattr(Person, "_AGE")}')    # False
print(f'hasattr(Person, "_age"): {hasattr(Person, "_age")}')  # True
```
