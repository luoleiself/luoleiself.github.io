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

## 内置函数

由 python 解释器提供, 不需要导入直接使用

### 类型转换

#### int()

将任意值转换成整型化值, 参数必须是字符串或者真实的数值

- 第一个参数为要转换的值
  - 浮点数则截断整数部分
  - 参数如果是字符串, 只能由 `数字, +, -, _` 组成
  - 布尔值 True 转为 1, False 转为 0
- 第二个参数为进制数位

#### float()

将任意值转换为浮点化值, 参数必须是字符串或者真实的数值

- 参数为整型则转换为浮点数
- 参数为字符串的浮点数表示形式，支持科学计数格式
- 布尔值 True 转为 1.0, False 转为 0.0

#### complex()

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

>>> bool({})  # 空字典转换为 False
False
>>> bool(set()) # 空集合转换为 False
False
>>> bool(tuple()) # 空元组转换为 False
False
>>> bool([])  # 空列表转换为 False   
False
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

>>> str([1, 2, 3, 4]) # list 转换为字符串
'[1, 2, 3, 4]'
>>> str(('a', 'b',))  # tuple 转换为字符串
"('a', 'b')"
>>> str({1, 2, 3, 4}) # set 转换为字符串
'{1, 2, 3, 4}'
>>> str({'a': 'A', 'b': 'B'}) # dict 转换为字符串
"{'a': 'A', 'b': 'B'}"
```

#### tuple()

创建或转换其他类型为元组, 参数为空或`可迭代对象`

#### list()

创建或转换其他类型为列表, 参数为空或`可迭代对象`

#### dict()

创建或转换其他类型为字典, 参数为空, `具名参数`, `包含双项序列的任意序列`

#### set()

创建或转换其他类型为集合, 参数为空或`可迭代对象`

#### frozenset()

创建或转换其他类型为不可变的集合, 参数为空或`可迭代对象`

#### chr()/ord()

转换单个 unicode 码对应的字符

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

#### 其他转换

##### bin()

转换二进制字符串

```python
>>> bin(100)
'0b1100100'
>>> bin(1)  
'0b1'
```

##### oct()

转换八进制字符串

```python
>>> oct(73)
'0o111'
```

##### hex()

转换十六进制字符串

```python
>>> hex(255)
'0xff'
```

##### ascii()

替代 repr(), 返回一个包含对象可打印表示的字符串, 但使用 \x、\u或\U转义符对 repr() 返回的字符串中的非 ascii 字符进行转义

### 数学运算

#### 基本数学运算

##### abs()

##### pow()

##### round()

四舍五入

##### divmod()

返回商和余数的元组

#### 数值计算

##### sum()

##### max()/min()

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

### 迭代器和序列处理

#### range()

生成数字序列, 返回一个`可迭代对象`, 不能直接使用, 需要`配合其他函数或语句`才能获取值

- range() 是`惰性求值`的，不会一次性占用大量内存
  - 只有 1 个参数时生成 0 到 arg 的数字序列
  - start, 开始，包含
  - end, 结束，不包含
  - step，步长绝对值，默认为 1
    - 如果为 -1, `反向生成数字序列`

```python
# 立即返回，几乎不占内存
big_range = range(1000000)
# 创建包含 100 万个元素的列表，占用大量内存
big_list = list(range(1000000))
```

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

#### enumerate()

枚举索引和值

```python
>>> fruits = ['apple', 'banana', 'cherry']
>>> for i, fruit in enumerate(fruits, start=1): 
...     print(f'{i}: {fruit}')
... 
1: apple
2: banana
3: cherry
```

#### zip()

并行迭代

```python
names = ['Alice', 'Bob', 'Charlie']
ages = [25, 30, 35]
cities = ['Beijing', 'Shanghai', 'Guangzhou']

# 合并多个序列
for name, age in zip(names, ages):
    print(f"{name} is {age} years old")

# 创建字典
person_dict = dict(zip(names, ages))
print(person_dict)  # {'Alice': 25, 'Bob': 30, 'Charlie': 35}

# 处理不等长序列
print(list(zip(names, ages, cities)))  # 以最短的为准
```

#### map()

映射函数返回结果为迭代器对象, 需要使用 list() 转换为列表

```python
squared = list(map(lambda x: x**2, [1, 2, 3, 4]))
print(squared)  # [1, 4, 9, 16]

# 多个序列
result = list(map(lambda x, y: x + y, [1, 2, 3], [10, 20, 30]))
print(result)  # [11, 22, 33]
```

#### filter()

过滤序列返回结果为迭代器对象, 需要使用 list() 转换为列表

```python
# 过滤出偶数
evens = list(filter(lambda x: x % 2 == 0, [1, 2, 3, 4, 5, 6]))
print(evens)  # [2, 4, 6]

# 过滤掉 None 和 False 值
values = [0, 1, False, 2, '', 3, None, 4]
filtered = list(filter(None, values))  # None 表示过滤掉假值
print(filtered)  # [1, 2, 3, 4]
```

#### sorted()

排序

```python
# 基本排序
numbers = [3, 1, 4, 1, 5, 9, 2]
print(sorted(numbers))  # [1, 1, 2, 3, 4, 5, 9]
print(sorted(numbers, reverse=True))  # [9, 5, 4, 3, 2, 1, 1]

# 自定义排序键
words = ['python', 'Java', 'C++', 'JavaScript']
print(sorted(words, key=len))  # ['C++', 'Java', 'python', 'JavaScript']
print(sorted(words, key=str.lower))  # 忽略大小写
```

#### reversed()

反转返回结果为反向迭代器对象, 需要使用 list() 转换为列表

```python
print(list(reversed([1, 2, 3, 4])))  # [4, 3, 2, 1]
```

### 逻辑判断

#### all()/any()

逻辑判断

```python
# all() - 所有元素都为真
print(all([True, True, True]))    # True
print(all([True, False, True]))   # False
print(all([1, 2, 3, 4]))          # True
print(all([1, 0, 3, 4]))          # False

# any() - 至少一个元素为真
print(any([False, False, True]))  # True
print(any([False, False, False])) # False
print(any([0, 0, 1, 0]))          # True
print(any([]))                    # False（空列表返回 False）
```

### 输入输出

#### print()/input()

### 对象和属性操作

#### type()

#### isinstance()/issubclass()

```python
# isinstance() - 检查类型
print(isinstance(123, int))        # True
print(isinstance("Hello", str))    # True
print(isinstance([], (list, tuple))) # True

# issubclass() - 检查子类
class Animal: pass
class Dog(Animal): pass
print(issubclass(Dog, Animal))     # True
```

#### id() 获取变量引用的对象的 id

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

#### hash()

获取哈希值

#### dir()

列出属性和方法

```python
>>> lt = [1, 2, 3, 4]
>>> dir(lt)
['__add__', '__class__', '__class_getitem__', '__contains__', '__delattr__', '__delitem__', '__dir__', '__doc__', '__eq__', '__format__', '__ge__', '__getattribute__', '__getitem__', '__getstate__', '__gt__', '__hash__', '__iadd__', '__imul__', '__init__', '__init_subclass__', '__iter__', '__le__', '__len__', '__lt__', '__mul__', '__ne__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__reversed__', '__rmul__', '__setattr__', '__setitem__', '__sizeof__', '__str__', '__subclasshook__', 'append', 'clear', 'copy', 'count', 'extend', 'index', 'insert', 'pop', 'remove', 'reverse', 'sort']
```

#### help()

帮助

#### 属性操作

##### hasattr()

```python
# hasattr() - 检查是否有属性
print(hasattr(p, 'name'))      # True
print(hasattr(p, 'age'))       # False
```

##### getattr()

获取属性

##### setattr()

设置属性

##### delattr()

删除属性

### 代码运行

#### eval()/exec()

```python
# eval() - 执行字符串表达式
print(eval("3 + 5 * 2"))    # 13

# exec() - 执行字符串代码
exec("x = 10; y = 20; print(x + y)")  # 30
```

#### compile()

### 高级函数

#### len()

#### open()

打开文件

#### callable()

返回对象是否可调用, 判断对象是否有 `__call__` 方法

```python
>>> class A:
...     pass
... 
>>> a = A()    
>>> callable(a) # 判断 a 是否可调用
False
>>> a() # 不能调用, 报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'A' object is not callable

# 定义 __call__ 方法
>>> class B:
...     def __call__(self):
...             print(f'has __call__ method...')
... 
>>> b = B();   
>>> callable(b)
True
>>> b()
has __call__ method...
```

#### vars()

返回模块、类、实例或任何其他具有 `__dict__` 属性的对象的 `__dict__` 属性

#### iter()

从`可迭代对象`返回一个迭代器对象

```python
>>> g = iter([1, 2, 3, 4])
>>> g
<list_iterator object at 0x0000023107888BE0>
>>> next(g) # 使用 next 获取迭代器对象的下一个值
1
>>> for i in g: # 使用 for 遍历迭代器
...     print(i)
... 
2
3
4
>>> next(g) # 迭代器结束触发异常
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
StopIteration
```

#### next()

获取迭代器的下一个值, 如果传入了默认值则在迭代器结束之后返回该默认值, 否则将触发 StopIteration 异常

```python
>>> gen = (x for x in range(10)) # 生成器推导式
>>> next(gen)
0
>>> next(gen)
1
... # 更多操作
>>> next(gen)   # 迭代器结束后再次获取没有默认值则触发异常
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
StopIteration
>>> next(gen, 'default value')  # 返回默认值
'default value'
```

#### locals() 返回函数内部的内容目录

- 在全局作用域内使用时获取的结果和 `globals()` 相同

#### globals() 返回当前全局作用域的内容目录

- `__name__` 属性在只有当前文件直接运行时, `__name__` 属性值才会是 `__main__`，否则被作为模块导入时, 属性值为模块名

```python
>>> print(globals()) 
{'__name__': '__main__', '__doc__': None, '__package__': None, '__loader__': <_frozen_importlib_external.SourceFileLoader object at 0x0000014B4D70BE00>, '__spec__': None, '__annotations__': {}, '__builtins__': <module 'builtins' (built-in)>, 'platform': <module 'platform' from 'D:\\uv\\python_install\\cpython-3.12.11-windows-x86_64-none\\Lib\\platform.py'>, 'sys': <module 'sys' (built-in)>, 'original_ps1': '>>> ', 'is_wsl': False, 'REPLHooks': <class '__main__.REPLHooks'>, 'get_last_command': <function get_last_command at 0x0000014B4D9C36A0>, 'PS1': <class '__main__.PS1'>, 'name': 'global variable', 'new_func': <function new_func at 0x0000014B4D9E20C0>, 'animal': 'wom', 'change_and_print_global': <function change_and_print_global at 0x0000014B4DA05EE0>, 'change_global': <function change_global at 0x0000014B4DA05D00>}
```

#### property()

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
