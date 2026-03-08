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

### 生成器

生成器是一个 python 序列生成对象, 无需一次性在内存中创建可能会很长的序列

- 生成器是动态生成值, 迭代器只能遍历一次

### 装饰器

装饰器是一种函数，接受一个函数作为输入并返回另一个函数

### 文档注释

函数体的顶部定义的多行字符串

## 类

- `__init__` 非必需的初始化方法, 第一个参数为 self 实例本身
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


# 实例化未定义初始化方法的子类时, 参数必须符合父类初始化方法参数的要求
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
hunter3 = SubQuestionQuote('Daffy Duck', 'It\'s rabbit season', 'good')
who_says(hunter3)
```

### 内置方法(魔术方法)

比较魔术方法

- `__str__` 自定义返回实例的结果, 默认输出实例的内存地址
- `__lt__` 自定义实例的 < 比较, 返回 bool
- `__le__` 自定义实例的 <= 比较, 返回 bool
- `__gt__` 自定义实例的 >比较, 返回 bool
- `__ge__` 自定义实例的 >= 比较, 返回 bool
- `__eq__`  自定义实例的 == 比较, 返回 bool
- `__ne__`  自定义实例的 != 比较, 返回 bool

数学运算魔术方法

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

- 使用 `__` 起始定义内置属性, python 将内置属性增加 `__类名` 的前缀绑定到实例上
  - 单下划线开始的属性和方法名约定私有, 但实际上是公开的, 子类可以重写
  - 双下划线开始的属性和方法名私有实现, 不想被子类意外重写

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

# 实例化后内置属性 __name 将变为 __Duck__name 绑定到实例上
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
print(f'banana.color {banana.color}')  # yellow, 实例会修改过的同名属性也会变化
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

### 多继承

- 访问自己不存在的属性或方法时, 优先使用最先继承的父类的属性和方法

### 多态
