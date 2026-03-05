## 集合

> 可哈希: 对象有一个固定的哈希值(通过 hash() 函数获取), 在整个生命周期中不会改变, python 用这个 hash 值来快速查找、比较对象

集合内部实现基于哈希表

可变的无序不重复的`可哈希`的任意类型的容器

- 使用 `{ }` 定义集合
- 使用 `set()` 内置函数创建或转换其他类型为集合, 参数为空或`可迭代对象`

```python
# 使用 { } 定义集合, 元素不能为空, 否则定义为字典
>>> s = {1, 2, 3, '1', 2, 'A'}
>>> s
{1, 2, 3, 'A', '1'}
>>> type(s)
<class 'set'>
>>> s = {}  # 空元素定义为字典
>>> type(s)
<class 'dict'>

# 创建或转换
>>> s = set() # 定义空集合, 只能使用 set() 函数定义空集合, 因为 { } 被字典占用了
>>> s   
set()
>>> type(s)
<class 'set'>
>>> set('hellworld')  # 将 str 转换为集合
{'h', 'e', 'l', 'd', 'o', 'r', 'w'}
>>> set(range(6)) # 将数字序列转换为集合
{0, 1, 2, 3, 4, 5}
>>> set(['a', 'b', 'c', 'a', 'b', 'B']) # 将 list 转换为集合
{'B', 'b', 'c', 'a'}
>>> set(('a', 'b', 'c', 'a', 'b', 'B')) # 将 tuple 转换为集合
{'B', 'b', 'c', 'a'}
>>> set({'a': 'A', 'b': 'B', 'c': 'C'}) # 将 dict 的键转换为集合
{'b', 'c', 'a'}

>>> s = set(['a', 'b', ['c', 'd']]) # 嵌套 list 不可哈希, 报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: unhashable type: 'list'
>>> set(123)  # 参数必须是可迭代对象
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'int' object is not iterable
```

- 使用 `frozenset()` 创建不可变集合, 参数为空或`可迭代对象`

```python
>>> fs = frozenset()
>>> type(fs)
<class 'frozenset'>
>>> len(fs)
0
>>> frozenset(['a', 'b', ('c', 'd')]) # 可迭代对象
frozenset({('c', 'd'), 'b', 'a'})

>>> fs = frozenset({1, 2, 3})
frozenset({1, 2, 3})
>>> fs.add(4)
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
AttributeError: 'frozenset' object has no attribute 'add'
```

### 集合比较

支持 `==`, `!=`, `<`, `<=`, `>`, `>=`

```python
>>> s1 = {1, 2}
>>> s2 = {3, 4}
>>> s1 == s2
False
>>> s1 > s2
False
>>> s1 >= s2
False
>>> s1 <= s2
False
>>> s1 != s2
True
```

### 修改集合

- add() 添加元素, 修改原集合, 返回 None
  - 参数必须符合`可哈希`

```python
>>> s = {1, 2, 3}
>>> s
{1, 2, 3}
>>> s.add(1)  # 忽略重复元素
>>> s
{1, 2, 3}
>>> s.add('A')
>>> s
{1, 2, 3, 'A'}
>>> s.add(4)   
>>> s
{1, 2, 3, 4, 'A'}
>>> s.add(('hello', 'world')) # 添加 tuple
>>> s
{1, 2, 3, 4, ('hello', 'world'), 'A'}

>>> s.add(['python', 'god'])  # 参数不能为 list, 不可哈希
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: unhashable type: 'list'
>>> s.add({'a': 'A'}) # 参数不能为 dict, 不可哈希
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: unhashable type: 'dict'
>>> s.add({True, False})  # 参数不能为 set, 不可哈希
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: unhashable type: 'set'
```

- update() 添加元素, 修改原集合, 返回 None
  - 参数必须符合 `set()` 创建集合的参数格式, 忽略重复值

```python
>>> s = set()
>>> s.update('hello') # 添加 str
>>> s
{'o', 'h', 'e', 'l'}
>>> s.update(range(5))  # 添加数字序列
>>> s
{0, 1, 'h', 2, 3, 4, 'e', 'l', 'o'}
>>> s.update(['world', 'python']) # 添加 list
>>> s
{0, 1, 'h', 2, 3, 4, 'e', 'l', 'python', 'world', 'o'}
>>> s.update(('good', 'New')) # 添加 tuple
>>> s
{0, 1, 'h', 2, 3, 4, 'e', 'l', 'python', 'good', 'New', 'world', 'o'}
>>> s.update({'a': 'A', 'b': 'B'})  # 添加 dict 的键
>>> s
{0, 1, 'h', 2, 3, 4, 'e', 'l', 'python', 'b', 'good', 'New', 'world', 'a', 'o'}
>>> s.update({'Cat', 'Tome'}) # 添加 set
>>> s
{0, 1, 'h', 2, 3, 4, 'e', 'l', 'python', 'b', 'good', 'New', 'world', 'Tome', 'a', 'Cat', 'o'}

>>> s.update(123) # 参数必须是可迭代对象
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'int' object is not iterable
```

### 删除

- del 删除集合

> del 是 python 语句

```python
>>> s = {1,2,3,4}
>>> del s # 删除集合
>>> s
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
NameError: name 's' is not defined
```

- discard() 删除集合中的元素, 返回 None
  - val, 值不存在也不会报错

```python
>>> s = {1, 2, 3, 'A'}
>>> s 
{1, 2, 3, 'A'}
>>> s.discard(5) 
>>> s
{1, 2, 3, 'A'}
>>> s.discard(2)
>>> s
{1, 3, 'A'}

# 空集合删除不存在的元素不会报错
>>> set().discard(5)
```

- remove() 删除集合中的元素, 返回 None
  - val，如果值不存在会报错

```python
>>> s = {1, 2, 3, 4, 5}
>>> type(s)
<class 'set'>
>>> s
{1, 2, 3, 4, 5}
>>> s.remove(2)
>>> s
{1, 3, 4, 5}

>>> s.remove(10)  # 值不存在报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
KeyError: 10
```

- pop() 随机删除集合中的元素并返回删除的元素
  - 集合为空会报错

```python
>>> s = {1, 2, 3, 4, 5}
>>> s
{1, 2, 3, 4, 5}
>>> s.pop() # 删除集合元素
1
>>> s
{2, 3, 4, 5}

>>> len(s)  # 空集合
0
>>> s.pop() # 集合为空报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
KeyError: 'pop from an empty set'
```

- clear() 清空集合, 返回 None

```python
>>> s.clear()
>>> s
set()
```

### 集合运算

**集合运算的方法参数必须是`可迭代对象`**

```python
>>> s1 = {1, 2}
>>> s2 = {2, 3}
```

#### 交集

同时出现在两个集合中的元素

- `&` 运算符返回新集合
- intersection() 返回新集合

```python
# & 取交集
>>> s1 & s2
{2}
>>> s1
{1, 2}
>>> s2
{2, 3}

# intersection() 取交集
>>> s1.intersection(s2)
{2}
>>> s1
{1, 2}
>>> s2
{2, 3}
>>> s1.intersection({3, 4})
set()

>>> s1.intersection(123)  # 参数必须是可迭代对象       
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'int' object is not iterable
```

- `&=` 运算符修改原集合
- intersection_update() 修改原集合, 只保留交集元素

```python
# &= 修改原集合
>>> s1 = {1, 2}
>>> s2 = {2, 3}
>>> s1 &= s2
>>> s1
{2}
>>> s2
{2, 3}

# intersection_update() 修改原集合
>>> s1 = {1, 2}
>>> s2 = {2, 3}
>>> s1.intersection_update(s2)
>>> s1
{2}
>>> s2
{2, 3}
```

#### 并集

存在于任意集合中的元素

- `|` 运算符返回新集合
- union() 返回新集合

```python
# | 取并集
>>> s1 | s2
{1, 2, 3}
>>> s1
{1, 2}
>>> s2
{2, 3}

# union() 取并集
>>> s1.union(s2)
{1, 2, 3}
>>> s1
{1, 2}
>>> s2
{2, 3}

>>> s1.union(123)  # 参数必须是可迭代对象       
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'int' object is not iterable
```

- `|=` 运算符修改原集合

```python
# |= 修改原集合
>>> s1 = {1, 2}
>>> s2 = {2, 3}
>>> s1 |= s2
>>> s1
{1, 2, 3}
>>> s2
{2, 3}
```

#### 差集

存在于第一个集合不存在于第二个集合中的元素

- `-` 运算符返回新集合
- difference() 返回新集合

```python
# - 取差集
>>> s1 - s2
{1}
>>> s1
{1, 2}
>>> s2
{2, 3}

# difference() 取差集
>>> s1.difference(s2)
{1}
>>> s1
{1, 2}
>>> s2
{2, 3}

>>> s1.difference(123)  # 参数必须是可迭代对象      
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'int' object is not iterable
```

- `-=` 运算符修改原集合
- difference_update() 修改原集合, 只保留差集元素

```python
# -= 修改原集合
>>> s1 = {1, 2}
>>> s2 = {2, 3} 
>>> s1 -= s2
>>> s1
{1}
>>> s2
{2, 3}

# difference_update() 修改原集合
>>> s1 = {1, 2}
>>> s2 = {2, 3}
>>> s1.difference_update(s2)
>>> s1
{1}
>>> s2
{2, 3}
```

#### 异或集

元素仅在两个集合中出现一次, 并集减去交集

- `^` 运算符返回新集合
- symmetric_difference() 返回新集合

```python
# ^ 取异或集
>>> s1 ^ s2
{1, 3}
>>> s1
{1, 2}
>>> s2
{2, 3}

# symmetric_difference() 取异或集
>>> s1.symmetric_difference(s2)
{1, 3}
>>> s1
{1, 2}
>>> s2
{2, 3}

>>> s1.symmetric_difference(123)  # 参数必须是可迭代对象
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'int' object is not iterable
```

- `^=` 运算符修改原集合
- symmetric_difference_update() 修改原集合, 只保留异或集

```python
# ^= 修改原集合
>>> s1 = {1, 2}
>>> s2 = {2, 3}
>>> s1 ^= s2
>>> s1
{1, 3}
>>> s2
{2, 3}

# symmetric_difference_update() 修改原集合
>>> s1 = {1, 2}   
>>> s2 = {2, 3}
>>> s1.symmetric_difference_update(s2)
>>> s1
{1, 3}
>>> s2
{2, 3}
```

#### 子集

判断第二个集合是否包含第一个集合中的所有元素

- `<=` 运算符返回 bool
- issubset() 返回 bool

```python
# <= 判断子集
>>> s1 <= s2
False
>>> s1.issubset(s2)
False

# 参数传入可迭代对象
>>> s1.issubset({1, 2, 3, 'a'})
True

# 集合是本身的子集
>>> s1 <= s1
True
>>> s1.issubset(s1)
True
```

#### 真子集

判断第二个集合是否包含第一个集合中的所有元素，并且还包含其他元素

- `<` 运算符返回 bool

```python
# < 判断真子集
>>> s1 < s2
False
>>> s1 < {1, 2, 3, 4, 'a'}
True

# 集合不是本身的真子集
>>> s1 < s1
False
```

#### 超集

判断第一个集合是否包含第二个集合中的所有元素, 和 子集 相反

- `>=` 运算符返回 bool
- issuperset() 返回 bool

```python
# >= 判断超集
>>> s1 >= s2
False
>>> s1.issuperset(s2)
False

# 参数传入可迭代对象
>>> s1.issuperset({1})
True

# 集合是本身的超集
>>> s1 >= s1
True
>>> s1.issuperset(s1)
True
```

#### 真超集

判断第一个集合是否包含第二个集合中的所有元素, 并且还包含其他元素

- `>` 运算符返回 bool

```python
# > 判断真超集
>>> s1 > s2
False
>>> s1 > {1}
True

# 集合不是本身的真超集
>>> s1 > s1
False
```
