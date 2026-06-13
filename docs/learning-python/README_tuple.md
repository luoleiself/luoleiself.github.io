## 元组

不可变的有序重复的任意数据类型的序列

- 使用 `,` 定义元组
- 使用 `( )` 定义元组, 当只有一个元素时, `尾逗号不能省略`
- 使用 `tuple()` 内置函数创建或转换其他类型为元组, 参数为空或`可迭代对象`
- 使用 `(* )` 解构`可迭代对象`为元组

```python
# , 定义元组, 只有一个元素时不能省略尾逗号
marx_tuple = 'Gloves',
print(type(marx_tuple))
# <class 'tuple'>
marx_tuple = 'Gloves', True, False, 3.1415, True
print(marx_tuple)
# ('Gloves', True, False, 3.1415, True)

marx_tuple = ('Gloves')  # 省略尾逗号变成普通的赋值操作
print(marx_tuple)
'Gloves'
print(type(marx_tuple))
# <class 'str'>

# () 定义元组, 只有一个元素时不能省略尾逗号
te = ()  # 定义空元组
print(te)
# ()
print(type(te))
# <class 'tuple'>
marx_tuple = ('Gloves', 'Sheep', 'Wolves')
print(type(marx_tuple))
# <class 'tuple'>
marx_tuple = ('Gloves',)  # 尾逗号不能省略
print(type(marx_tuple))
# <class 'tuple'>
print(marx_tuple)
# ('Gloves',)
len(marx_tuple)
# 1

# tuple() 创建或转换
tuple()  # 定义空元组
# ()
print(tuple('hello'))  # 将 str 转换为元组
# ('h', 'e', 'l', 'l', 'o')
print(tuple(range(4)))  # 将数字序列转换为元组
# (0, 1, 2, 3)
print(tuple(['hello', 'world']))  # 将 list 转换为元组
# ('hello', 'world')
print(tuple({'hello', 'world'}))  # 将 set 转换为元组
# ('hello', 'world')
print(tuple({'name': 'python', 'age': 18}))  # 将 dict 的键转换为元组
# ('name', 'age')
print(tuple(123))  # 参数必须是可迭代对象     
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: 'int' object is not iterable

# * 解构为元组
print((*'hello',))  # 解构 str
# ('h', 'e', 'l', 'l', 'o')
print((*range(4, 10, 2),))  # 解构数字序列        
# (4, 6, 8)
print((*[('a', 'A'), ('b', 'B')],))  # 解构 list
# (('a', 'A'), ('b', 'B'))
print((*{'a', 'b', 'c'},))  # 解构 set
# (True, False, 3.14)
print((*{'a', 'b', 'c'},))  # 解构 set
# ('a', 'b', 'c')
print((*{'a': 'A', 'b': 'B', 'c': 'C'},))  # 解构 dict 的键
# ('a', 'b', 'c')
print((*123,))  # 必须是可迭代对象
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: Value after * must be an iterable, not int
```

- [ ] 获取元素, 下标越界会报错

```python
marx_tuple = 'Gloves', True, False, 3.1415, True
print(marx_tuple[0])
'Gloves'
print(marx_tuple[-3])
# False
print(marx_tuple[-100])  # 下标越界报错
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# IndexError: tuple index out of range
print(marx_tuple[10])  # 下标越界报错
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# IndexError: tuple index out of range
```

- `切片`截取并返回新元组，用法同字符串切片截取子串一致

```python
marx_tuple = tuple(('gloves', 'sheep', 'wolves', True, 1, False, 1.1))
print(marx_tuple)
# ('gloves', 'sheep', 'wolves', True, 1, False, 1.1)
print(marx_tuple[1:])
# ('sheep', 'wolves', True, 1, False, 1.1)
print(marx_tuple[::-1])  # 反向截取, 步长为 1
# (1.1, False, 1, True, 'wolves', 'sheep', 'gloves')
print(marx_tuple[::-2])  # 反向截取, 步长为 2
# (1.1, 1, 'wolves', 'gloves')
print(marx_tuple[6:-3:-1])  # 反向截取
# (1.1, False)
```

### 元组解构

- 交换变量的值

```python
first = 'one'
second = 'two'
first, second = second, first  # 交换变量的值
print(first)
'two'
print(second)
'one'
```

- 元组解构赋值给变量, 未匹配到变量的剩余值保存为一个列表, 剩余值变量优先级最低
    - 使用 *_ 或者 *var_name 收集剩余值, 在同一层级两者不能同时使用

```python
# 嵌套结构需要对齐格式, 否则将嵌套结构作为整体赋值
a, b, c, *d, f = ('a', 'b', ['c', 'd', 'e'], 'f')
print(a)
'a'
print(b)
'b'
print(c)  # 嵌套结构作为整体赋值
# ['c', 'd', 'e']
print(d)
# []
print(f)
# 'f'
# 嵌套结构对齐格式
a, b, (c, *d), *e, f = (1, 2, (3, 4, 5), 6, 7, 8)
print(a)
# 1
print(b)
# 2
print(c)
# 3
print(d)  # 收集嵌套结构剩余值
# [4, 5]
print(e)
# [6, 7]  # 收集剩余值
print(f)
# 8
```

- 元组解构为其它类型

```python
te = ('a', 'b', 'c')
print((*te, 1, 2))  # 解构为元组
# ('a', 'b', 'c', 1, 2)
print([*te, 1, 2])  # 解构为列表
# ['a', 'b', 'c', 1, 2]
print({*te, 1, 2})  # 解构为集合
# {'c', 1, 2, 'b', 'a'}
```

### 元组比较

支持 `==`, `!=`, `<`, `<=`, `>`, `>=`

```python
a = (1, 2)
b = (3, 4)
print(a == b)
# False
print(a < b)
# True
print(a >= b)
# False
```

### 元组拼接

`*` 和 `+` 拼接元组返回新的元组

####      * 重复元组

重复数量只能是整型, 不区分书写顺序

```python
print(('gloves', 'sheep') * 3)
# ('gloves', 'sheep', 'gloves', 'sheep', 'gloves', 'sheep')
print(3 * ('gloves', 'sheep'))
# ('gloves', 'sheep', 'gloves', 'sheep', 'gloves', 'sheep')
```

#### + 拼接元组

- `+` 后面支持任意元素类型的元组

```python
# 元组类型必须在 + 前面，后面可以跟任意类型
print(1 + 'gloves', )
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: unsupported operand type(s) for +: 'int' and 'str'
print(1 + ('gloves',))
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: unsupported operand type(s) for +: 'int' and 'tuple'

a = (1, 2)
b = (3, 4)
a += b
print(a)
# (1, 2, 3, 4)

print('gloves', + 1)
# ('gloves', 1)
print('gloves', 'sheep', + 1, 2, 3)  # + 左右类型一致
# ('gloves', 'sheep', 1, 2, 3)
print(('gloves', 'sheep') + (1, True))  # + 左右类型一致
# ('gloves', 'sheep', 1, True)

# 连接元组时, + 左右表示形式不一致报错
print('gloves', 'sheep', + (1, 2, 3))
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: bad operand type for unary +: 'tuple'
print(('gloves', 'sheep') + 1, 2, 3)
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# TypeError: can only concatenate tuple (not "int") to tuple
```

### 查找

- count(val) 统计元素出现的次数
- index(val[, start[, stop]]) 查找第一个符合的偏移并返回下标, 未找到会报错
    - val, 要查找的值
    - start, 开始位置, 默认为列表开始位置
    - stop, 结束位置, 默认为列表结束位置

```python
te = 1, 2, 3, 4
print(te)
# (1, 2, 3, 4)
te.index(2, 1, 3)
# 1
te.index(2, 2, 3)
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# ValueError: tuple.index(x): x not in tuple

te.count(3)
# 1
```

### 删除元组

- del 删除元组

> del 是 python 语句

```python
te = (1, 2, 3, 4,)
print(te)
# (1, 2, 3, 4)
del te  # 删除元组
print(te)
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# NameError: name 'te' is not defined
```
