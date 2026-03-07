## 列表

可变的有序重复的任意数据类型的序列

- 使用 `[ ]` 定义列表
- 使用 `list()` 内置函数创建或转换其他类型为列表, 参数为空或`可迭代对象`
- 使用 `[* ]` 解构`可迭代对象`为列表

```python
# [ ] 定义列表
>>> lt = [] # 定义空列表
>>> lt
[]
>>> type(lt)
<class 'list'>
>>> weekdays = ['Monday', 'TuesDay', 'Wednesday', 'Thursday', 'Friday']
>>> weekdays
['Monday', 'TuesDay', 'Wednesday', 'Thursday', 'Friday']
>>> len(weekdays)
5

# list() 创建或转换
>>> empty_list = list() # 定义空列表
>>> empty_list
[]
>>> list('hello')   # 将 str 转换为列表  
['h', 'e', 'l', 'l', 'o']
>>> list(range(4))  # 将数字序列转换为列表
[0, 1, 2, 3]
>>> list(('hello', 'world'))  # 将 tuple 转换为列表
['hello', 'world']
>>> list({'hello', 'python'}) # 将 set 转换为列表
['python', 'hello']
>>> list({'name': 'python', 'age': 18}) # 将 dict 的键转换为列表
['name', 'age']
>>> list(123) # 参数必须是可迭代对象
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'int' object is not iterable

# * 解构为列表
>>> [*'hello']  # 解构 str
['h', 'e', 'l', 'l', 'o']
>>> [*range(1, 9, 2)] # 解构数字序列
[1, 3, 5, 7]
>>> [*[1, 2, 3]]  # 解构 list
[1, 2, 3]
>>> [*([1, 2], [3, 4])] # 解构 tuple
[[1, 2], [3, 4]]
>>> [*{'a', 'b', 'c'}]  # 解构 set
['a', 'b', 'c']
>>> [*{'a': 'A', 'b': 'B', 'c': 'C'}] # 解构 dict 的键
['a', 'b', 'c']
>>> [*123]  # 必须是可迭代对象
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: Value after * must be an iterable, not int
```

- [ ] 获取元素, 下标越界会报错

```python
>>> weekdays = ['Monday', 'TuesDay', 'Wednesday', 'Thursday', 'Friday']
>>> weekdays[2]  
'Wednesday'
>>> weekdays[-3]
'Wednesday'
>>> weekdays[-4]
'TuesDay'

>>> weekdays[10]  # 下标越界报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
IndexError: list index out of range
>>> weekdays[-10] # 下标越界报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
IndexError: list index out of range
```

- `切片`截取并返回新列表, 用法同字符串切片截取子串一致

```python
>>> weekdays = ['Monday', 'TuesDay', 'Wednesday', 'Thursday', 'Friday']
>>> weekdays[::2]    
['Monday', 'Wednesday', 'Friday']
>>> weekdays[4:-4:-1]
['Friday', 'Thursday', 'Wednesday']
>>> weekdays
['Monday', 'TuesDay', 'Wednesday', 'Thursday', 'Friday']
```

### 解构赋值

```python
>>> a = 1
>>> b = 2
>>> a, b = [b, a] # 交换变量的值
>>> a
2
>>> b
1

# 解构赋值
>>> a, b, *k = [1, 2, 3, 4, 5]
>>> a
1
>>> b
2
>>> k
[3, 4, 5]
```

### 列表比较

支持 `==`, `!=`, `<`, `<=`, `>`, `>=`

```python
>>> a = [1, 2]
>>> b = [3, 4]
>>> a == b
False
>>> a >= b
False
>>> a <= b
True
```

### 切片修改列表

- 值必须为`可迭代对象`
- `=` 右侧的元素数量不必和切片包含的元素数量相同

```python
>>> a = [1, 2, 3, 4]
>>> a
[1, 2, 3, 4]
>>> a[1:3] = (98, 99, 100, 120) # 使用 tuple 修改列表, 元素数量大于切片包含的元素数量
>>> a
[1, 98, 99, 100, 120, 4]
>>> a[1:4] = 'hello'  # 使用 str 修改列表
>>> a
[1, 'h', 'e', 'l', 'l', 'o', 120, 4]
>>> a[1:7] = ['hello', 'woorld']  # 使用 list 修改列表
>>> a
[1, 'hello', 'woorld', 4]
>>> a[1:3] = {'name': 'python'}   # 使用 dict 的键修改列表
>>> a
[1, 'name', 4]
>>> a[1:2] = {'dog','cat'}  # 使用 set 修改列表
>>> a
[1, 'dog', 'cat', 4]
>>> a[1:3] = ['h']  # 元素数量小于切片包含的元素数量
>>> a
[1, 'h', 4]

>>> a[1:7] = 1  # 值必须为可迭代对象
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: must assign iterable to extended slice
```

### 列表拼接

`*` 和 `+` 拼接列表返回新的列表

#### * 重复列表

重复数量只能是整型, 不区分书写顺序

```python
>>> a = [1, 2]
>>> a * 3
[1, 2, 1, 2, 1, 2]
>>> 3 * a
[1, 2, 1, 2, 1, 2]
>>> a
[1, 2]
```

#### + 拼接列表

- `+` 必须是列表类型

```python
>>> a = [1, 2]
>>> a + [3]
[1, 2, 3]
>>> [3] + a # 不区分书写顺序
[3, 1, 2]
>>> a
[1, 2]

>>> a + 1 # + 必须是列表类型
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: can only concatenate list (not "int") to list
>>> a + 'False' # + 必须是列表类型
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: can only concatenate list (not "str") to list
```

#### extend() 拼接列表

- 修改原列表, 返回 None
- 参数必须符合 `list()` 创建列表的参数格式

```python
>>> a = [1,2]
[1, 2]
>>> a.extend([3.1, True]) # 拼接 list
>>> a
[1, 2, 3.1, True]
>>> a.extend(('hello', 'world'))  # 拼接 tuple
>>> a
[1, 2, 3.1, True, 'hello', 'world']
>>> a.extend({'python'})  # 拼接 set
>>> a
[1, 2, 3.1, True, 'hello', 'world', 'python']
>>> a.extend({'age': 18}) # 拼接 dict 的键
>>> a
[1, 2, 3.1, True, 'hello', 'world', 'python', 'age']
>>> a.extend(range(8,10)) # 拼接数字序列
>>> a
[1, 2, 3.1, True, 'hello', 'world', 'python', 'age', 8, 9]
>>> a.extend('good')  # 拼接 str
>>> a
[1, 2, 3.1, True, 'hello', 'world', 'python', 'age', 8, 9, 11, 15, 'g', 'o', 'o', 'd']
>>> print(a.extend([11, 15])) # 返回 None
None

>>> a.extend(3.1) # 参数必须是可迭代对象
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'float' object is not iterable
>>> a.extend(True)  # 参数必须是可迭代对象
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'bool' object is not iterable
```

#### append() 追加元素

- 修改原列表, 返回 None
- 参数被当作整体添加到列表末尾

```python
>>> a = [1, 2]
>>> a
[1, 2]
>>> a.append('h')
>>> a
[1, 2, 'h']
>>> a.append('world')
>>> a
[1, 2, 'h', 'world']
>>> a.append(['python', 'god']) # 参数列表作为整体添加到列表末尾
>>> a
[1, 2, 'h', 'world', ['python', 'god']]
>>> print(a.append(10)) # 返回 None
None
>>> a.append(('hello', 'world')) # 追加 tuple
>>> a
[1, 2, 'h', 'world', ['python', 'god'], ('hello', 'world')]
>>> a.append({'a': 'A'})  # 追加 dict
>>> a
[1, 2, 'h', 'world', ['python', 'god'], ('hello', 'world'), {'a' : 'A'}]
```

#### insert() 插入元素

- 修改原列表, 返回 None
  - index, 指定偏移索引, 超出列表范围则按列表起始或结束位置
  - val，插入的值, 可以为任意类型

```python
>>> a = [1, 2, 3, 4, 5]
>>> a
[1, 2, 3, 4, 5]
>>> a.insert(1, 'a')
>>> a
[1, 'a', 2, 3, 4, 5]
>>> a.insert(3, 'b')
>>> a
[1, 'a', 2, 'b', 3, 4, 5]
>>> a.insert(10, 'c') # 偏移索引超过结束位置按照结束位置
>>> a
[1, 'a', 2, 'b', 3, 4, 5, 'c']
>>> a.insert(-10, 'd')  # 偏移索引超过起始位置按照起始位置
>>> a
['d', 1, 'a', 2, 'b', 3, 4, 5, 'c']
>>> a.insert(1, ['f', 'g']) # 插入 list
>>> a
['d', ['f', 'g'], 1, 'a', 2, 'b', 3, 4, 5, 'c']

>>> print(a.insert(1, ('hello', 'python'))) # 插入 tuple
None
>>> a
['d', ('hello', 'python'), ['f', 'g'], 1, 'a', 2, 'b', 3, 4, 5, 'c']
```

### 查找

- index() 查找第一个符合的偏移并返回下标, 未找到会报错
  - val, 要查找的值
  - start, 开始位置, 默认为列表开始位置
  - end, 结束位置, 默认为列表结束位置

```python
>>> a = [1, 2, 3, 4, 5]
>>> a.index(1)
0
>>> a.index(4)
3
>>> a.index(2, 0, 3)
1
>>> a.index(2, 0, 1)  # 在区间内未找到会报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: 2 is not in list

>>> a.index(10) # 未找到报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: 10 is not in list
>>> a.index(-10)  # 未找到报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: -10 is not in list
```

### 删除

- del 删除列表或列表中元素, 下标越界会报错

> del 是 python 语句, 并非列表方法

```python
>>> a = [1, 2, 3, 4, 5] 
>>> a
[1, 2, 3, 4, 5]
>>> del a[-1] # 删除最后一个
>>> a
[1, 2, 3, 4]
>>> del a[-4] # 删除第一个
>>> a
[2, 3, 4]

>>> del a[10] # 下标越界报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
IndexError: list assignment index out of range
>>> del a[-10]  # 下标越界报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
IndexError: list assignment index out of range

>>> del a # 删除列表
>>> a
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
NameError: name 'a' is not defined
```

- remove() 按值删除第一个符合条件的列表项，返回 None
  - val, 如果值不存在会报错

```python
>>> a = [1, 3, 2, 3, 4, 3, 5]   
>>> a.remove(3)
>>> a
[1, 2, 3, 4, 3, 5]

>>> print(a.remove(3))  # 返回 None
None

>>> a.remove(10)  # 值不存在会报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: list.remove(x): x not in list
```

- pop() 按下标删除并返回删除的列表项
  - index, 元素偏移索引, 下标越界会报错

```python
>>> a = [1, 2, 3, 4, 5] 
>>> a
[1, 2, 3, 4, 5]
>>> a.pop()
5
>>> a
[1, 2, 3, 4]
>>> a.pop(-2)
3
>>> a
[1, 2, 4]
>>> a.pop(0)
1
>>> a
[2, 4]

>>> a.pop(10) # 下标越界报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
IndexError: pop index out of range
>>> a.pop(-10)  # 下标越界报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
IndexError: pop index out of range
```

- clear() 删除所有列表项, 返回 None, python 3.3 支持
