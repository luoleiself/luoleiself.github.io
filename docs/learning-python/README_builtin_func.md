## 内置函数

由 python 解释器提供, 不需要导入直接使用

### 类型转换

#### int()

将任意值转换成整型化值, 参数必须是字符串或者真实的数值

- 第一个参数为要转换的值
    - 浮点数则截断整数部分
    - 参数如果是字符串, 只能由 `数字, +, -, _` 组成
    - 布尔值 True 转为 1, False 转为 0
    - 可以为其他进制数字符串形式
- 第二个参数为当前值的进制数位

#### float()

将任意值转换为浮点化值, 参数必须是字符串或者真实的数值

- 参数为整型则转换为浮点数
- 参数为字符串的浮点数表示形式，支持科学计数格式
- 布尔值 True 转为 1.0, False 转为 0.0

- inf, Inf, INFINITY, iNfINity 都表示为无穷大, 左侧使用 '+' 和 '-' 表示无穷大的方向
- NaN 表示为 Not a Number

```python
float('inf')
# inf
float('-iNf')
# -inf
float('+iNfiNiTy')
# inf

float('nan')
# nan
f1 = float('nan')
f2 = float('nan')
print(f1 == f2)
# False
```

#### complex()

#### bool()

将任意值转为布尔值

- 任意非零值转为 True
- 任意零值转为 False

```python
bool(1)
# True
bool(1.0)
# True
bool('1.0')
# True
bool('0.0')
# True
bool('True')
# True
bool('False')  # 任意非零值转换为 True
# True
bool('')  # 任意零值转换为 False
# False
bool(-1.0)
# True
bool(+0.0)
# False
bool(-0.0)
# False
bool(0.0)
# False
bool(0)
# False
bool(1e0)
# True

bool('')  # 空字符串转为 False
# False
bool({})  # 空字典转换为 False
# False
bool(set())  # 空集合转换为 False
# False
bool(tuple())  # 空元组转换为 False
# False
bool([])  # 空列表转换为 False   
# False
```

#### str()

将任意值转换为字符串

- 将 bytes 对象传递给 str() 函数时, 传入 encoding 参数表示转换为字符串

```python
str(1)
'1'
str(1.)
'1.0'
str(1.0)
'1.0'
str(1.01)
'1.01'
str(1.0_1)
'1.01'
str(-1)
'-1'
str(-1.)
'-1.0'
str(-1.0)
'-1.0'
str(True)
'True'
str(False)
'False'

str([1, 2, 3, 4])  # list 转换为字符串
'[1, 2, 3, 4]'
str(('a', 'b',))  # tuple 转换为字符串
"('a', 'b')"
str({1, 2, 3, 4})  # set 转换为字符串
'{1, 2, 3, 4}'
str({'a': 'A', 'b': 'B'})  # dict 转换为字符串
"{'a': 'A', 'b': 'B'}"

b = bytes('中国')
print(b)
b'\xe4\xb8\xad\xe5\x9b\xbd'
type(b)
# <class 'bytes'>
str(b)
"b'\\xe4\\xb8\\xad\\xe5\\x9b\\xbd'"
str(b, 'utf-8')  # bytes 转换为字符串指定编码格式
'中国'
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
ord('a')
# 97
ord('一')
# 19968
chr(19968)
'一'
ord('棒')
# 26834

chr(97)
'a'
chr(65)
'A'
chr(48)
'0'
chr(26834)
'棒'
```

#### bytearray()

创建或转换为一个新的字节数组, 表示 0 <= x < 256 范围内的可变整数序列

- 如果传入字符串, 需要指定编码格式
- 如果传入整数, 创建空字节填充的指定大小的字节数组
- 如果传入符合缓冲区接口的对象, 将使用该对象的只读缓冲区来初始化字节数组
- 如果传入可迭代对象, 则元素必须符合 0 <= x < 256 范围内的整数
- 如果没有参数, 将创建一个大小为 0 的数组

```python
s = 'hello 中国, 你好!'
# bytearray() 创建可变整数序列
bytearray()  # 创建空字节数组
bytearray(b'')
bytearray(s)  # 字符串必须指定编码格式
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: string argument without an encoding

ba = bytearray(s, 'utf-8')
bytearray(b'hello \xe4\xb8\xad\xe5\x9b\xbd, \xe4\xbd\xa0\xe5\xa5\xbd!')
bytearray(['a', 2, 3, 4])  # 可迭代对象必须是符合条件的整数
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: 'str' object cannot be interpreted as an integer
bytearray([1, 2, 3, 4])
bytearray(b'\x01\x02\x03\x04')

bay = bytearray(10)  # 可变整数序列
bytearray(b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00')
bay[1] = 250
print(bay)
bytearray(b'\x00\xfa\x00\x00\x00\x00\x00\x00\x00\x00')

ba.decode('utf-8')  # 字节数组解码为字符串
'hello 中国, 你好!'
```

#### bytes()

创建或转换为一个新的表示 0 <= x < 256 范围内的不可变整数序列

- 参数和 bytearray() 相同

```python
s = 'hello 中国, 你好!'
# bytes() 创建不可变整数序列
bytes(s, 'utf-8')  # 字符串必须指定编码格式
b'hello \xe4\xb8\xad\xe5\x9b\xbd, \xe4\xbd\xa0\xe5\xa5\xbd!'
bytes(10)  # 创建指定大小的不可变整数序列
b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
bytes([1, 2, 3, 4])
b'\x01\x02\x03\x04'
bytes(10)[1] = 12  # 不可变整数序列     
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: 'bytes' object does not support item assignment

bytes(s, 'utf-8').decode('utf-8')
'hello 中国, 你好!'
```

#### 其他转换

##### bin()/oct()/hex()

将数字转换二进制/八进制/十六进制字符串, 带进制前置符

```python
# 数字转换为二进制, 带进制前置符
bin(100)
'0b1100100'
bin(1)
'0b1'

# 数字转换为八进制, 带进制前置符
oct(73)
'0o111'

# 数字转换为十六进制, 带进制前置符
hex(255)
'0xff'
```

##### ascii()

替代 repr(), 返回一个包含对象可打印表示的字符串, 但使用 \x、\u或\U转义符对 repr() 返回的字符串中的非 ascii 字符进行转义

```python
ascii('中')
"'\\u4e2d'"
ascii('hello中')
"'hello\\u4e2d'"

repr('hello 中')
"'hello 中'"
```

##### format()

> 内部调用 type(vale).__format__(value, format_spec)

将一个值转换为由第二个参数控制的 `字符串格式` 表示,
如果 type(value) 不支持 __format__ 方法, format_spec 为空返回 str(value)

- value, 转换的值
- format_spec, 格式化标识, 默认转换为字符串

```python
# 格式化为字符串, 等价于调用 str(value)
format(100)
'100'
format(1.2)
'1.2'
format(True)
'True'
format([])
'[]'

# 数字进制格式化, 没有进制前置符
format(20013, 'b')  # 十进制格式化为二进制, 等价于调用 bin(20013)
'100111000101101'
format(20013, 'o')  # 十进制格式化为八进制, 等价于调用 oct(20013)
'47055'
format(20013, 'x')  # 十进制格式化为十六进制, 等价于调用 hex(20013)
'4e2d'
format(0b100111000101101, 'd')  # 二进制格式化为十进制, 等价于调用 int('0b100111000101101', 2)
'20013'
format(0b100111000101101, 'x')  # 二进制格式化为十六进制, 等价于调用 hex('0b100111000101101')
'4e2d'
format(0b100111000101101, 'o')
'47055'
format(0x4e2d, 'd')  # 十六进制格式化为十进制, 等价于调用 int('0x4e2d', 16)
'20013'
format(0x4e2d, 'o')  # 十六进制格式化为八进制, 等价于调用 oct('0x4e2d')
'47055'
format(0x4e2d, 'b')  # 十六进制格式化为二进制, 等价于调用 bin('0x4e2d')
'100111000101101'
format(0o47055, 'd')  # 八进制格式化为十进制, 等价于调用 int('0o47055', 8)
'20013'
format(0o47055, 'b')  # 八进制格式化为二进制, 等价于调用 bin('0o47055')
'100111000101101'
format(0o47055, 'x')
'4e2d'
```

### 数学运算

#### 基本数学运算

##### abs()

##### pow()

##### round()

四舍五入

- 0.5 和 -0.5 都为 0

```python
round(0.5)
# 0
round(-0.5)
# 0
round(-0.6)
# -1
round(0.6)
# 1
```

##### divmod()

返回商和余数的元组

#### 数值计算

##### sum()

计算`可迭代对象`中的元素的累和, start 不能为字符串

##### max()/min()

获取容器中的最大值或最小值

- key=None, 可选参数, 函数的返回值作为比较的键
- 字典获取所有键中的最大/最小值

```python
max('abcdefg')
'g'
max([1, 2, 3, 4, 5])
# 5
max((1, 2, 3, 4, 5))
# 5
max({1, 2, 3, 4, 5})
# 5
max({'a': 'A', 'b': 'B', 'c': 'C', 'd': 'D'})
'd'

min('abcdefg')
'a'
min([1, 2, 3, 4, 5])
# 1
min((1, 2, 3, 4, 5,))
# 1
min({1, 2, 3, 4, 5})
# 1
min({'a': 'A', 'b': 'B', 'c': 'C', 'd': 'D'})
'a'
```

### 迭代器和序列处理

#### range()

生成数字序列, 返回一个`可迭代对象`, 不能直接使用, 需要`配合其他函数或语句`才能获取值

- range() 是`惰性求值`的，不会一次性占用大量内存
    - 只有 1 个参数时生成 0 到 arg 的数字序列
    - start, 开始，包含
    - stop, 结束，不包含
    - step，步长绝对值，默认为 1
        - 如果为 -1, `反向生成数字序列`

```python
# 立即返回，几乎不占内存
big_range = range(1000000)
# 创建包含 100 万个元素的列表，占用大量内存
big_list = list(range(1000000))
```

```python
print(range(4)) # 直接使用不会输出结果

num = range(4)
print(num[2])
# 2
list(range(4))
# [0, 1, 2, 3]
list(range(1, 4)) 
# [1, 2, 3]
list(range(1, 4, 2))
# [1, 3]
tuple(range(4))
# (0, 1, 2, 3)
tuple(range(1, 4))
# (1, 2, 3)
tuple(range(1, 4, 2))
# (1, 3)

# 反向生成数字序列
list(range(4, -2, -1))
# [4, 3, 2, 1, 0, -1]
list(range(4, -2, -2))
# [4, 2, 0]
```

#### enumerate()

枚举索引和值

```python
fruits = ['apple', 'banana', 'cherry']
for i, fruit in enumerate(fruits, start=1):
    print(f'{i}: {fruit}')

# 1: apple
# 2: banana
# 3: cherry
```

#### zip()

并行迭代, 以最短的为准

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
squared = list(map(lambda x: x ** 2, [1, 2, 3, 4]))
print(squared)  # [1, 4, 9, 16]

# 多个序列
result = list(map(lambda x, y: x + y, [1, 2, 3], [10, 20, 30]))
print(result)  # [11, 22, 33]
```

#### filter()

> filter(function, iterable, /)

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

#### reversed()

返回结果为反向迭代器对象, 需要使用 list() 转换为列表

- 参数为可反转对象或支持 len 和 getitem的对象

```python
print(list(reversed([1, 2, 3, 4])))  # [4, 3, 2, 1]
```

#### sorted()

> sorted(iterable, /, *, key=None, reverse=False)

- key 指定一个参数的函数, 用于从 iterable 中的每个元素中提取比较键, 默认值为 None

排序`可迭代对象`返回一个新列表

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

#### slice()

返回一个切片对象, 表示由范围(start, stop, step)指定的索引集, start 和 step 参数默认为 None

- 使用切片语法时也会生成切片对象

```python
# 传入 stop
s = slice(5)
print(s.start)
print(s.stop)
# 5
print(s.step)
print(s)
slice(None, 5, None)
s.indices(3)  # 根据给定的长度, 计算切片边界并返回索引集
# (0, 3, 1)

# 传入 start, stop, step
s = slice(3, 10, 4)
print(s)
# slice(3, 10, 4)
print(s.indices(20))  # 根据给定的长度, 计算切片边界并返回索引集
# (3, 10, 4)
print(s.indices(2))  # 根据给定的长度, 计算切片边界并返回索引集
# (2, 2, 4)
print(s.indices(8))  # 根据给定的长度, 计算切片边界并返回索引集
# (3, 8, 4)
```

### 逻辑判断

#### all()/any()

逻辑判断`可迭代对象`

all()

- 如果可迭代对象的所有元素调用 `bool()` 都为 True, 则返回 True
- 如果可迭代对象为空, 则返回 True

any()

- 如果可迭代对象的任意元素调用 `bool()` 为 True, 则返回 True
- 如果可迭代对象为空, 则返回 False

```python
# all() 和 any() 接收可迭代对象
all(123)
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: 'int' object is not iterable
any(123)
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: 'int' object is not iterable

# all() - 所有元素都为真
print(all([True, True, True]))  # True
print(all([True, False, True]))  # False
print(all([1, 2, 3, 4]))  # True
print(all([1, 0, 3, 4]))  # False
# 可迭代对象为空, 返回 True
print(all(''))  # True
print(all(list()))  # True  
print(all(tuple()))  # True
print(all(set()))  # True
print(all(dict()))  # True
print(all(bytes()))  # True
print(all(bytearray()))  # True

# any() - 至少一个元素为真
print(any([False, False, True]))  # True
print(any([False, False, False]))  # False
print(any([0, 0, 1, 0]))  # True
print(any([]))  # False（空列表返回 False）
# 可迭代对象为空, 返回 False
print(any(''))  # False
print(any(list()))  # False
print(any(tuple()))  # False
print(any(set()))  # False
print(any(dict()))  # False
print(any(bytearray()))  # False
print(any(bytes()))  # False
```

### 输入输出

#### print()/input()

在标准输出/输入中写入或读取一行

### 对象和属性操作

#### object()

所有类的最终基类, 创建并返回一个新的无任何属性的对象, 不接收任何参数

- 没有 `__dict__` 魔法属性, 无法为新的实例添加属性

```python
no = object(1)
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: object() takes no arguments
o2 = object()
o2.name = 'hello'
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# AttributeError: 'object' object has no attribute 'name'
```

#### type()

返回对象的类型, 返回值是一个类型对象

#### super()

返回一个代理对象, 该对象将方法调用委托给类型的父类或同级类

#### isinstance()/issubclass()

```python
# isinstance() - 检查类型
print(isinstance(123, int))  # True
print(isinstance("Hello", str))  # True
print(isinstance([], (list, tuple)))  # True


# issubclass() - 检查子类
class Animal: pass


class Dog(Animal): pass


class Cat(Animal): pass


print(issubclass(Dog, Animal))  # True
print(issubclass(Dog, Cat))  # False
```

#### id() 获取变量引用的对象的 id

- 返回对象的"标识". 这是一个整数, 保证此对象在其生命周期内是唯一且恒定的.

```python
n = 1
id(n)
# 140726623005112
m = n
id(m)
# 140726623005112
n = 2
id(n)
# 140726623005144
id(m)
# 140726623005112

names = ['alice', 'tom', 'jerry']
id(names)
# 2970721128832
names_lst = names
id(names_lst)
# 2970721128832
names.append('bob')
print(names)
# ['alice', 'tom', 'jerry', 'bob']
print(names_lst)
# ['alice', 'tom', 'jerry', 'bob']
```

#### hash()

获取哈希值, 哈希值是整数. 它们用于在字典查找过程中快速比较字典键, 比较相等的数值具有相同的哈希值(即使它们是不同类型的,
如1和1.0的情况)

#### dir()

列出属性和方法

```python
lt = [1, 2, 3, 4]
dir(lt)
# ['__add__', '__class__', '__class_getitem__', '__contains__', '__delattr__', '__delitem__', '__dir__', '__doc__', '__eq__', '__format__', '__ge__', '__getattribute__', '__getitem__', '__getstate__', '__gt__', '__hash__', '__iadd__', '__imul__', '__init__', '__init_subclass__', '__iter__', '__le__', '__len__', '__lt__', '__mul__', '__ne__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__reversed__', '__rmul__', '__setattr__', '__setitem__', '__sizeof__', '__str__', '__subclasshook__', 'append', 'clear', 'copy', 'count', 'extend', 'index', 'insert', 'pop', 'remove', 'reverse', 'sort']
```

#### help()

帮助

#### 属性操作

##### hasattr()

```python
class Person:
    def __init__(self, name, avg):
        self.name = name
        self.avg = avg


p = Person('Tom', 18)
# hasattr() - 检查是否有属性
print(hasattr(p, 'name'))  # True
print(hasattr(p, 'age'))  # False
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
print(eval("3 + 5 * 2"))  # 13

# exec() - 执行字符串代码
exec("x = 10; y = 20; print(x + y)")  # 30
```

#### compile()

将源代码编译为代码或 AST 对象, 代码对象可以通过 `exec()` 或 `eval()` 执行, 源可以是字符串、字节字符串或 AST 对象

### 高级函数

#### len()

返回对象的长度, 参数支持序列(str, list, tuple, range, bytes)或集合(dict, set, forzenset)

#### open()

打开文件

#### callable()

返回对象是否可调用, 判断对象是否有 `__call__` 方法

```python
class A:
    pass


a = A()
callable(a)  # 判断 a 是否可调用
# False
a()  # 不能调用, 报错


# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: 'A' object is not callable

# 定义 __call__ 方法
class B:
    def __call__(self):
        print(f'has __call__ method...')


b = B()
callable(b)
# True
b()
# has __call__ method...
```

#### vars()

返回模块、类、实例或任何其他具有 `__dict__` 属性的对象的 `__dict__` 属性

```python
class Dog:
    name = 'DD'

    def __init__(self, n, a):
        self.n = n
        self.a = a


dog = Dog('Tom', 19)
vars(dog)
# {'n': 'Tom', 'a': 19}
```

#### iter()

从`可迭代对象`返回一个迭代器对象

```python
g = iter([1, 2, 3, 4])
print(g)
# <list_iterator object at 0x0000023107888BE0>
next(g)  # 使用 next 获取迭代器对象的下一个值
# 1
for i in g:  # 使用 for 遍历迭代器
    print(i)

# 2
# 3
# 4
next(g)  # 迭代器结束触发异常
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# StopIteration
```

#### next()

获取迭代器的下一个值, 如果传入了默认值则在迭代器结束之后返回该默认值, 否则将触发 StopIteration 异常

```python
gen = (x for x in range(10))  # 生成器推导式
next(gen)
# 0
next(gen)
# 1
...  # 更多操作
next(gen)  # 迭代器结束后再次获取没有默认值则触发异常
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# StopIteration
next(gen, 'default value')  # 返回默认值
'default value'
```

#### locals() 返回函数内部的内容目录

- 在全局作用域内使用时获取的结果和 `globals()` 相同

#### globals() 返回当前全局作用域的内容目录

- `__name__` 属性在只有当前文件直接运行时, `__name__` 属性值才会是 `__main__`，否则被作为模块导入时, 属性值为模块名

```python
print(globals())
# {'__name__': '__main__', '__doc__': None, '__package__': None, '__loader__': <_frozen_importlib_external.SourceFileLoader object at 0x0000014B4D70BE00>, '__spec__': None, '__annotations__': {}, '__builtins__': <module 'builtins' (built-in)>, 'platform': <module 'platform' from 'D:\\uv\\python_install\\cpython-3.12.11-windows-x86_64-none\\Lib\\platform.py'>, 'sys': <module 'sys' (built-in)>, 'original_ps1': '>>> ', 'is_wsl': False, 'REPLHooks': <class '__main__.REPLHooks'>, 'get_last_command': <function get_last_command at 0x0000014B4D9C36A0>, 'PS1': <class '__main__.PS1'>, 'name': 'global variable', 'new_func': <function new_func at 0x0000014B4D9E20C0>, 'animal': 'wom', 'change_and_print_global': <function change_and_print_global at 0x0000014B4DA05EE0>, 'change_global': <function change_global at 0x0000014B4DA05D00>}
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
    # name = property(get_name, set_name)
    # age = property(get_age, set_age)


p = Person('xiaoming', 19)
print(f'p.name {p.name} p.age {p.age}')
# 仍然可以访问到隐藏的属性
print(f'p.hidden_name {p.hidden_name} p.hidden_age {p.hidden_age}')
```
