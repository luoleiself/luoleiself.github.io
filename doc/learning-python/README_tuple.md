## 元组

只读的任意数据类型的序列

- 使用 `,` 定义元组
- 使用 `( )` 定义元组, 当小括号只有一个元素时, `尾逗号不能省略`
- 使用 `tuple()` 内置函数创建或转换其他类型为元组, 参数为空或`可迭代对象`

```python
# 使用 , 定义元组, 只有一个元素时不能省略尾逗号
>>> marx_tulpe = 'Gloves',
>>> type(marx_tuple)      
<class 'tuple'>
>>> marx_tuple = 'Gloves', True, False, 3.1415, True
>>> marx_tuple
('Gloves', True, False, 3.1415, True)
>>> marx_tuple = ('Gloves') # 省略尾逗号变成普通的赋值操作
>>> marx_tuple
'Gloves'
>>> type(marx_tuple)
<class 'str'>

# 使用 () 定义元组
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

# 创建空的元组
>>> tuple()
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

- 借助元组交换变量的值

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

```python
>>> a = (1,2) 
>>> b = (3,4) 
>>> a == b
False
>>> a < b 
True
>>> a >= b
False
>>> a += b
>>> a
(1, 2, 3, 4)
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
