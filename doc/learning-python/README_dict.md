## 字典

字典内部实现基于哈希表

可变的键值对的容器, 键只能是不可变类型

- 使用 `{ }` 定义字典
- 使用 `dict()` 内置函数创建或转换其他类型为字典, 参数为空, `具名参数`, `包含双项序列的任意序列`
  - 具名参数必须是合法的变量名称

```python
# {} 定义字典
>>> dt = {} # 定义空字典
>>> dt
{}
>>> type(dt)
<class 'dict'>
>>> acme = {1: 1, 'name': 'python', 'flag': True}
>>> acme
{1: 1, 'name': 'python', 'flag': True}

# dict() 创建或转换
>>> acme = dict() # 定义空字典
>>> acme
{}
# 使用具名参数定义字典, 具名参数必须是合法的变量名称
>>> name_dict = dict(first = 'First', second = 'Second')
>>> name_dict
{'first': 'First', 'second': 'Second'}
# 具名参数不符合规范
>>> name_dict = dict(first = 'First', second = 'Second', 1 = 1)
  File "<stdin>", line 1
    name_dict = dict(first = 'First', second = 'Second', 1 = 1)
                                                         ^^^
SyntaxError: expression cannot contain assignment, perhaps you meant "=="?
# 具名参数不符合规范
>>> name_dict = dict(first = 'First', second = 'Second', 'third' = '1')
  File "<stdin>", line 1
    name_dict = dict(first = 'First', second = 'Second', 'third' = '1')
                                                         ^^^^^^^^^
SyntaxError: expression cannot contain assignment, perhaps you meant "=="?

# 使用包含双项序列的任意序列定义字典
# 列表
>>> lol = [['a', 'b'], ['c', 'd'], ('e', 'f')]
>>> lol
[['a', 'b'], ['c', 'd'], ('e', 'f')]
>>> dict(lol)
{'a': 'b', 'c': 'd', 'e': 'f'}
# 元组
>>> lol = (['a', 'b'], ['c', 'd'], ('e', 'f'))
>>> lol
(['a', 'b'], ['c', 'd'], ('e', 'f'))
>>> dict(lol)
{'a': 'b', 'c': 'd', 'e': 'f'}
# 双字母字符串列表
>>> los = ['ab', 'cd', 'ef']
>>> los 
['ab', 'cd', 'ef']
>>> dict(los)
{'a': 'b', 'c': 'd', 'e': 'f'}
# 双字母字符串元组
>>> los = ('ab', 'cd', 'ef')
>>> los
('ab', 'cd', 'ef')
>>> dict(los)
{'a': 'b', 'c': 'd', 'e': 'f'}

>>> dt = dict({['a', 'b']: 'AB'}) # 键只能是不可变类型
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: unhashable type: 'list'
>>> dict(['ab', 'cd', 'efg']) # 字符串超过数量报错 
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: dictionary update sequence element #2 has length 3; 2 is required
>>> dict(('ab', 'cd', 'efg',))  # 字符串超过数量报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: dictionary update sequence element #2 has length 3; 2 is required
```

### 字典比较

支持 `==`, `!=`

```python
>>> a = {1: 1, 2: 2, 3: 3}
>>> b = {3: 3, 1: 1, 2: 2}
>>> a == b
True
>>> a != b
False
>>> a <= b  # 不支持大小比较
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: '<=' not supported between instances of 'dict' and 'dict'
>>> a > b   # 不支持大小比较
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: '>' not supported between instances of 'dict' and 'dict'
```

### 合并字典

- `**` 浅复制合并多个字典并返回新字典, 同名键按照后面覆盖前面

```python
>>> first = {'a': 'A', 'b': 'B'}
>>> second = {'b': 'BB', 'c': 'C'}
>>> {**first, **second} # 同名键后面优先级高
{'a': 'A', 'b': 'BB', 'c': 'C'}
```

- update() 合并字典, 修改原字典, 返回 None, 同名键按照传入的键覆盖原键
  - 参数必须符合 `dict()` 创建字典的参数格式

```python
>>> first.update(second)
>>> first
{'a': 'A', 'b': 'BB', 'c': 'C'}
>>> dt.update(first='F')  # 具名参数
>>> dt
{'a': 'A', 'b': 'BB', 'c' : 'C', 'first': 'F'}
>>> dt.pop('first') # 删除键
'F'
>>> first.update([('g', 'G')])  # 元组列表
>>> first
{'a': 'A', 'b': 'BB', 'c': 'C', 'g': 'G'}
>>> first.update((['h', 'H'],)) # 列表元组
>>> first
{'a': 'A', 'b': 'BB', 'c': 'C', 'g': 'G', 'h': 'H'}
>>> first.update(['xy'])  # 双字母字符串列表
>>> first
{'a': 'A', 'b': 'BB', 'c': 'C', 'g': 'G', 'h': 'H', 'x': 'y'}
>>> first.update(('uV',)) # 双字母字符串元组
>>> first
{'a': 'A', 'b': 'BB', 'c': 'C', 'g': 'G', 'h': 'H', 'x': 'y', 'u': 'V'}

>>> first.update([1, 2])  # 参数必须是包含双项序列的任意序列
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: cannot convert dictionary update sequence element #0 to a sequence
>>> first.update(123) # 参数必须是包含双项序列的任意序列
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'int' object is not iterable
```

### 获取

- 使用 [key] 获取字典项, `key 不存在时会报错` , 可使用 in 检查 key 是否存在
- 使用 get() 获取字典项
  - key, 要获取的键
  - default, 可选, 默认为 None, 如果键不存在则返回该值

```python
>>> name_dict = dict(first = 'First', second = 'Second', third = 'Third')
>>> name_dict
{'first': 'First', 'second': 'Second', 'third': 'Third'}
>>> name_dict['forth']  # key 不存在时会报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
KeyError: 'forth'
>>> 'forth' in name_dict  # in 检查 key 是否存在
False

# 使用 get() 获取字典项
>>> name_dict.get('forth') 
>>> print(name_dict.get('forth'))
None
>>> print(name_dict.get('forth', 'Not found'))  # 未找到返回指定的值
Not found
```

python 3 中 `keys()`, `values()`, `items()` 方法返回的是可迭代视图, 使用 `list()` 转换为列表

- keys() 获取所有键

```python
>>> type(name_dict.keys())
<class 'dict_keys'>
>>> list(name_dict.keys())  # 使用 list() 转换为普通列表
['first', 'second', 'third']
```

- values() 获取所有值

```python
>>> type(name_dict.values())
<class 'dict_values'>
>>> list(name_dict.values())  # 使用 list() 转换为普通列表
['First', 'Second', 'Second']
```

- items() 获取所有键值对

```python
>>> type(name_dict.items()) 
<class 'dict_items'>
>>> list(name_dict.items()) # 使用 list() 转换为普通元组列表
[('first', 'First'), ('second', 'Second'), ('third', 'Third')]
```

- fromkeys() 根据传入的迭代器和值创建并返回一个新的字典
  - key, 键迭代器
  - val, 不传默认为 None

```python
>>> dict.fromkeys(('a', 'b', 1))  # 不传值默认为 None     
{'a': None, 'b': None, 1: None}
>>> dict.fromkeys(['a', 'b', 1], 'good')  # 根据 list 的值生成 dict
{'a': 'good', 'b': 'good', 1: 'good'}
>>> dict.fromkeys(('a', 'b', 1), 'good')  # 根据 tuple 的值生成 dict
{'a': 'good', 'b': 'good', 1: 'good'}
>>> dict.fromkeys({'a', 'b', 1}, 'good')  # 根据 list 的值生成 dict
{1: 'good', 'a': 'good', 'b': 'good'}
>>> dict.fromkeys({'a': 'A', 'b': 'B', 1: '1'}, 'good') # 根据 dict 的键生成 dict
{'a': 'good', 'b': 'good', 1: 'good'}

>>> dict.fromkeys(123, True)  # 必须是可迭代对象
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'int' object is not iterable
```

### 删除

- del 删除字典或字典键, `key 不存在时会报错`

> del 是 python 语句, 并非集合方法

```python
>>> del first['c']
>>> first
{'a': 'A', 'b': 'BB'}
>>> del first['d']  # key 不存在时会报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
KeyError: 'd'

>>> dt = {1: 'a', 2: 'b', 3: 'c'}
>>> type(dt)
<class 'dict'>
>>> del dt  # 删除字典
>>> dt
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
NameError: name 'dt' is not defined
```

- pop() 按 key 删除字典并返回删除的字典项
  - key, 要删除的键, 如果键不存在且没有传入默认值会报错
  - default, 可选, 默认为 None, 如果键不存在则返回该值

```python
>>> first
{'a': 'A', 'b': 'BB'}
>>> first.pop('d', 'Not found') # 键不存在时返回传入的默认值
'Not found'
>>> first.pop('b', 'Not found') # 返回删除的字典项
'BB'
>>> first
{'a': 'A'}

>>> first.pop('d')  # 键不存在时且没有传入默认值, 报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
KeyError: 'd'
```

- popitem() 随机删除字典中的键值对并返回 元组 类型
  - 字典为空会报错

```python
>>> first = {'a':'A', 'b':'B', ('a', 'b', 'c'): 'abc'}          
>>> first
{'a': 'A', 'b': 'B', ('a', 'b', 'c'): 'abc'}
>>> first.popitem()
(('a', 'b', 'c'), 'abc')
>>> first
{'a': 'A', 'b': 'B'}
>>> first.popitem()
('b', 'B')
>>> first
{'a': 'A'}
>>> first.popitem()
('a', 'A')
>>> first
{}
>>> first.popitem() # 字典为空会报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
KeyError: 'popitem(): dictionary is empty'
```

- clear() 清空字典, 返回 None

```python
>>> first
{'a': 'A'}
>>> first.clear()
```
