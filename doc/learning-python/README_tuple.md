## 元组

不可变的有序重复的任意数据类型的序列

- 使用 `,` 定义元组
- 使用 `( )` 定义元组, 当只有一个元素时, `尾逗号不能省略`
- 使用 `tuple()` 内置函数创建或转换其他类型为元组, 参数为空或`可迭代对象`
- 使用 `(* )` 解构`可迭代对象`为元组

```python
# , 定义元组, 只有一个元素时不能省略尾逗号
>>> marx_tulpe = 'Gloves',
>>> type(marx_tuple)      
<class 'tuple'>
>>> marx_tuple = 'Gloves', True, False, 3.1415, True
>>> marx_tuple
('Gloves', True, False, 3.1415, True)
# 省略尾逗号变成普通的赋值操作
>>> marx_tuple = ('Gloves')
>>> marx_tuple
'Gloves'
>>> type(marx_tuple)
<class 'str'>

# () 定义元组, 只有一个元素时不能省略尾逗号
>>> te = () # 定义空元组
>>> te
()
>>> type(te)
<class 'tuple'>
>>> marx_tuple = ('Gloves', 'Sheep', 'Wolves')
>>> type(marx_tuple)
<class 'tuple'>
>>> marx_tuple = ('Gloves',)  # 尾逗号不能省略
>>> type(marx_tuple)
<class 'tuple'>
>>> marx_tuple
('Gloves',)
>>> len(marx_tuple)
1

# tuple() 创建或转换
>>> tuple() # 定义空元组
()
>>> tuple('hello')  # 将 str 转换为元组
('h', 'e', 'l', 'l', 'o')
>>> tuple(range(4)) # 将数字序列转换为元组
(0, 1, 2, 3)
>>> tuple(['hello', 'world']) # 将 list 转换为元组
('hello', 'world')
>>> tuple({'hello', 'world'}) # 将 set 转换为元组
('hello', 'world')
>>> tuple({'name': 'python', 'age': 18})  # 将 dict 的键转换为元组
('name', 'age')
>>> tuple(123)  # 参数必须是可迭代对象     
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'int' object is not iterable

# * 解构为元组
>>> (*'hello',) # 解构 str
('h', 'e', 'l', 'l', 'o')
>>> (*range(4, 10, 2),) # 解构数字序列        
(4, 6, 8)
>>> (*[('a', 'A'), ('b', 'B')],)  # 解构 list
(('a', 'A'), ('b', 'B'))
>>> (*(True, False, 3.14),) # 解构 tuple
(True, False, 3.14)
>>> (*{'a', 'b', 'c'},) # 解构 set
('a', 'b', 'c')
>>> (*{'a': 'A', 'b': 'B', 'c': 'C'},)  # 解构 dict 的键
('a', 'b', 'c')
>>> (*123,) # 必须是可迭代对象
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: Value after * must be an iterable, not int
```

- [ ] 获取元素, 下标越界会报错

```python
>>> marx_tuple = 'Gloves', True, False, 3.1415, True
>>> marx_tuple[0]
'Gloves'
>>> marx_tuple[-3]
False
>>> marx_tuple[-100]  # 下标越界报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
IndexError: tuple index out of range
>>> marx_tuple[10]  # 下标越界报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
IndexError: tuple index out of range
```

- `切片`截取并返回新元组，用法同字符串切片截取子串一致

```python
>>> marx_tuple = tuple(('gloves', 'sheep', 'wolves', True, 1, False, 1.1))
>>> marx_tuple      
('gloves', 'sheep', 'wolves', True, 1, False, 1.1)
>>> marx_tuple[1:]  
('sheep', 'wolves', True, 1, False, 1.1)
>>> marx_tuple[::-1]
(1.1, False, 1, True, 'wolves', 'sheep', 'gloves')
>>> marx_tuple[::-2]
(1.1, 1, 'wolves', 'gloves')
>>> marx_tuple[6:-3:-1]
(1.1, False)
```

### 借助元组交换变量的值

```python
>>> first = 'one'
>>> second = 'two'
>>> first, second = second, first # 借助元组交换变量的值
>>> first
'two'
>>> second
'one'
```

### 元组比较

支持 `==`, `!=`, `<`, `<=`, `>`, `>=`

```python
>>> a = (1,2) 
>>> b = (3,4) 
>>> a == b
False
>>> a < b 
True
>>> a >= b
False
```

### 元组拼接

`*` 和 `+` 拼接元组返回新的元组

#### * 重复元组

重复数量只能是整型, 不区分书写顺序

```python
>>> ('gloves', 'sheep') * 3      
('gloves', 'sheep', 'gloves', 'sheep', 'gloves', 'sheep')
>>> 3 * ('gloves', 'sheep')
('gloves', 'sheep', 'gloves', 'sheep', 'gloves', 'sheep')
```

#### + 拼接元组

- `+` 左右的表示形式尽量保持一致

```python
# 元组类型必须在 + 前面，后面可以跟任意类型
>>> 1 + 'gloves',                  
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: unsupported operand type(s) for +: 'int' and 'str'
>>> 1 + ('gloves',)
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: unsupported operand type(s) for +: 'int' and 'tuple'

>>> a = (1, 2)
>>> b = (3, 4)
>>> a += b
>>> a
(1, 2, 3, 4)

>>> 'gloves', + 1
('gloves', 1)
>>> 'gloves', 'sheep', + 1, 2, 3  # + 左右表示形式一致
('gloves', 'sheep', 1, 2, 3)
>>> ('gloves', 'sheep') + (1, True) # + 左右表示形式一致
('gloves', 'sheep', 1, True)

# 连接元组时, + 左右表示形式不一致报错
>>> 'gloves', 'sheep', + (1, 2, 3)
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: bad operand type for unary +: 'tuple'
>>> ('gloves', 'sheep') + 1, 2, 3 
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: can only concatenate tuple (not "int") to tuple
```

### 查找

- index() 查找第一个符合的偏移并返回下标, 未找到会报错
  - val, 要查找的值
  - start, 开始位置, 默认为列表开始位置
  - end, 结束位置, 默认为列表结束位置

```python
>>> te = 1, 2, 3, 4
>>> te
(1, 2, 3, 4)
>>> te.index(2, 1, 3)
1
>>> te.index(2, 2, 3)
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: tuple.index(x): x not in tuple
```

### 删除元组

- del 删除元组

> del 是 python 语句

```python
>>> te = (1, 2,3,4,)
>>> te
(1, 2, 3, 4)
>>> del te  # 删除元组
>>> te
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
NameError: name 'te' is not defined
```
