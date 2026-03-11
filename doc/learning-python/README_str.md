## 字符串

不可变的字符序列

- [ ] 获取字符, 下标越界会报错

```python
>>> letters = 'abcdefg'
>>> letters[0]
'a'
>>> letters[-1]
'g'
>>> letters[-2]
'f'
>>> letters[10]  # 下标越界报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
IndexError: string index out of range
>>> letters[-100] # 下标越界报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
IndexError: string index out of range
```

- `切片`获取并返回新字符串, [strat:end:step], 含头不含尾, start 和 end 越界不会报错
  - start, 起始索引, 包含, 默认为字符串开始
  - end, 结束索引, 不包含, 默认为字符串结束
  - step, 步长绝对值，默认为 1, 每隔指定数量的`绝对值`的字符取一次
    - 如果为 -1, `反向提取子串`

```python
'''
  a  b  c  d  e  f  g
  0  1  2  3  4  5  6  正索引
 -7 -6 -5 -4 -3 -2 -1  负索引
'''

>>> letters = 'abcdefg'
>>> letters[:]
'abcdefg'
>>> letters[-8:-2]  # start 超出起始位置使用默认值
'abcde'
>>> letters[-8:10]  # end 超出结束位置使用默认值
'abcdefg'
>>> letters[3:-2] # 含头不含尾
'de'
>>> letters[8:12] # 下标越界提取为空串
''
>>> letters[8:-2] # 下标越界提取为空串
''
# 指定步长
>>> letters[::3]  # 每隔 3 个字符取一次
'adg'
>>> letters[-5:-2:2]  # 每隔 2 个字符取一次
'ce'
>>> letters[-5:-5:2]
''

>>> letters[::-1] # 反向提取字符串
'gfedcba'
>>> letters[6:1:-1]
'gfedc'
>>> letters[6:1:-2] # 反向每隔 2 个字符取一次
'gec'
>>> letters[-3:1:-2]
'ec'
```

### 解构赋值

```python
>>> a, b, *k = 'hello'
>>> a
'h'
>>> b
'e'
>>> k
['l', 'l', 'o']
```

### 字符串前缀

- f'', 格式化字符串, 3.6+ 支持
- b'', 字节字符串, 表示字节类型
- r'', 原始字符串, 忽略转义字符
- u'', Unicode 字符串
- fr'', 原始格式化字符串
- rb'', 原始字节字符串

```python
f'保留两位小数: {pi:.2f}'

# 字节字符串
>>> text = 'hello'
>>> bytes_str = b'hello'
>>> print(type(text))
<class 'str'>
>>> print(type(bytes_str))
<class 'bytes'>
>>> text[0] 
'h'
>>> bytes_str[0]
104

# 原始字符串
>>> print(r'hello\nworld')
hello\nworld
>>> len(r'hello\nworld')
12
```

### 函数

- strip() 移除字符串两端的空白字符('', '\t', '\n')
  - chars, 移除两端的指定字符, 默认为空白字符

```python
>>> t1 = '\n\t\na\n\tb\n\t\n'
>>> t1.strip()
'a\n\tb'
>>> t1.lstrip()
'a\n\tb\n\t\n'
>>> t1.rstrip()
'\n\t\na\n\tb'

# 移除字符串两端的指定字符
>>> '12 hello world 21'.strip('12') # 移除两端的字符
' hello world '
>>> '12 hello world 21'.lstrip('12')  # 移除左侧的字符
' hello world 21'
>>> '12 hello world 21'.rstrip('12')  # 移除右侧的字符
'12 hello world '
```

- find() 和 index(), 查找子串出现的下标, find 未找到返回 -1, index 未找到报错

```python
>>> t1 = '\n\t\na\n\tb\n\t\n'
>>> t1.find('a'， 3)  # 指定开始查找位置
3
>>> t1.rfind('a', 3)  # 指定开始查找位置  
3
>>> t1.index('a', 3)  # 指定开始查找位置
3
>>> t1.rindex('a', 3) # 指定开始查找位置
3

>>> t1.find('aa')
-1
>>> t1.rfind('aa')
-1
>>> t1.index('aa')  # 未找到报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: substring not found
>>> t1.rindex('aa') # 未找到报错
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: substring not found
```

- 对齐字符串(center, ljust, rjust), 在指定的空间范围内对齐, 如果空间范围小于字符串长度则返回原值

```python
>>> s = 'a duck goes into bar...'
>>> len(s)
23
>>> s.center(5) # 对齐空间范围小于字符串长度返回原值
'a duck goes into bar...'
>>> s2 = s.center(40) # 居中对齐
>>> s2
'        a duck goes into bar...         '
>>> len(s2)
40
>>> s.ljust(40) # 左对齐
'a duck goes into bar...                 '
>>> s.rjust(40) # 右对齐
'                 a duck goes into bar...'
```

### 字符串拼接

`*` 和 `+` 拼接字符串返回新的字符串

#### 字符串字面量拼接

只能是`字符串字面量`

```python
>>> name = 'python'
>>> 'hello' name  # 字符串拼接变量报错
  File "<stdin>", line 1
    'hello' name
            ^^^^
SyntaxError: invalid syntax

>>> 'hello' "world"
'helloworld'
>>> 'hello' "world" '''你好''' """python"""
'helloworld你好python'
>>> '''hello'''"world" """你好""" 'python'
'helloworld你好python'
```

#### * 重复字符串

重复数量只能是整型, 不区分书写顺序

```python
>>> """hello"""*2
'hellohello'
>>> num = 3
>>> 'hello' * num
'hellohellohello'

>>> 2 * 'hello'
'hellohello'

>>> 'hello' * 2.
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: can't multiply sequence by non-int of type 'float'
```

#### + 拼接字符串

只能是`字符串类型`, 其他类型需要先转换为字符串

```python
>>> 'hello' + 1
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: can only concatenate str (not "int") to str
>>> 'hello' + 1.2
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: can only concatenate str (not "float") to str
>>> 'hello' + True
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: can only concatenate str (not "bool") to str

name = "luolei"
message = "hello " + name
print(message)
```

#### % 占位符

> 老版本语法, python 3 兼容

% 出现在字符串内的数量和字符串末尾的 % 之后的数据项的数量必须相同

- 需要考虑浮点数的精度问题
- 多个数据项时必须使用元组形式

```python
>>> "hello %s" % 'python' # 单个数据项
'hello python'

# 多个数据项必须使用元组形式
>>> 'name: %s, age: %d' % 'python' 18 
  File "<stdin>", line 1
    'name: %s, age: %d' % 'python' 18
                                   ^^
SyntaxError: invalid syntax
>>> 'name: %s, age: %d' % 'python'   
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: not enough arguments for format string

salary = 5000
tel = 13112345678
>>> "my salary is %.2f, my tel is %d" % (salary, tel)  # 元组形式        
'my salary is 5000.00, my tel is 13112345678'
```

**占位符规则** <!--markdownlint-disable-line-->

- 起始字符 %
- 对齐字符: + 表示右对齐, - 表示左对齐
- 最小字段宽度
- . 字符, 分隔最小字段宽度和最大字符数
- 最大字符数
  - 类型为字符串时表示打印的字符数
  - 类型为浮点数时表示精度
- 转换类型字符

```python
# 格式化为字符串
>>> thing = 'woodchuck'
>>> '%s' % thing
'woodchuck'
>>> '%12s' % thing
'   woodchuck'
>>> '%2s' % thing 
'woodchuck'
>>> '%+12s' % thing # 右对齐
'   woodchuck'
>>> '%-12s' % thing # 左对齐
'woodchuck   '
>>> '%.3s' % thing 
'woo'
>>> '%.3s' % thing 
'woo'
>>> '%12.3s' % thing
'         woo'
>>> '%-12.3s' % thing
'woo         '
# 格式化为浮点数
>>> n = 98.6       
>>> '%f' % n
'98.600000'
>>> '%12f' % n
'   98.600000'
>>> '%-12f' % n 
'98.600000   '
>>> '%+12f' % n
'  +98.600000'
>>> '%.3f' % n 
'98.600'
>>> '%12.3f' % n
'      98.600'
>>> '%-12.3f' % n
'98.600      '
# 格式化为整数
>>> m = 9527    
>>> '%d' % m 
'9527'
>>> '%12d' % m
'        9527'
>>> '%-12d' % m
'9527        '
>>> '%.3d' % m # 精度对整数无效
'9527'
>>> '%12.3d' % m
'        9527'
>>> '%+12.3d' % m
'       +9527'
>>> '%-12.3d' % m
'9527        '
```

#### '{}'.format

> python 2.7 和 python 3 支持

```python
>>> thing = 'woodchuck'
>>> '{}'.format(thing)
'woodchuck'
```

- 下标指定参数，0 引用第一个参数，1 引用第二个参数

```python
>>> 'The {1} is in the {0}'.format('lake', 'woodchuck')
'The woodchuck is in the lake'
```

- 具名参数

```python
>>> 'The {thing} is in the {place}'.format(place='lake', thing='woodchuck')
'The woodchuck is in the lake
```

- 字典, 0 是 format 的第一个参数

```python
>>> d = {'thing': 'wookchuck', 'place': 'lake'}
>>> 'The {0[thing]} is in the {0[place]}'.format(d)                        
'The wookchuck is in the lake'
```

**占位符规则** <!--markdownlint-disable-line-->

- 起始冒号
- 填充字符，默认为 ''
- 对齐字符: < 表示左对齐, > 表示右对齐(默认), ^ 表示居中对齐
- 数字符号, 表示只在负数前面添加负号
- 进制符号: #, 输出进制前缀
  - 如果使用对齐字符, 进制符号必须在对齐字符之后, 在对齐字符之前则为填充字符
  - 如果使用填充符号，进制符号必须在填充符号之前, 否则填充字符不生效
- 最小字段宽度
- . 字符, 分隔最小字段宽度和最大字符数
- 最大字符数
- 分组符号 , 或 _, 仅支持数字整数部分
- 转换类型字符

```python
>>> 'The {:10s} is at the {:10s}'.format('wraith', 'window')  
'The wraith     is at the window    '
# 右对齐
>>> 'The {:>10s} is at the {:>10s}'.format('wraith', 'window') 
'The     wraith is at the     window'
# 居中对齐
>>> 'The {:^10s} is at the {:^10s}'.format('wraith', 'window')
'The   wraith   is at the   window  '
# 填充字符
>>> 'The {:!^10s} is at the {:^10s}'.format('wraith', 'window') 
'The !!wraith!! is at the   window  '
>>> 'The {:##^10s} is at the {:^10s}'.format('wraith', 'window')
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: Invalid format specifier '##^10s' for object of type 'str'

# 分组符号仅支持数字整数部分
>>> x = 1000000000
>>> f'{x = :15}'
'x =      1000000000'
>>> f'{x = :15_}'
'x =   1_000_000_000'
>>> f'{x = :15,}'
'x =   1,000,000,000'
>>> f'{x = :15b}'
'x = 111011100110101100101000000000'
>>> f'{x = :15_b}'
'x = 11_1011_1001_1010_1100_1010_0000_0000'
>>> f'{x = :15_o}'
'x =    73_4654_5000'
>>> f'{x = :15_x}'
'x =       3b9a_ca00'
>>> f'{x = :15,x}'
>>> f'{x = :15,b}'  # 分组符号 , 支持数字输出
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: Cannot specify ',' with 'b'.
>>> f'{x = :15,o}'
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: Cannot specify ',' with 'o'.
>>> f'{x = :15,x}'
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: Cannot specify ',' with 'x'.
# 仅支持整数部分分组
>>> x = 10000000.12398913
>>> f'{x = :15,}'
'x = 10,000,000.12398913'
>>> f'{x = :15_}'
'x = 10_000_000.12398913'

# 进制转换
# 在对齐字符之前为填充字符
>>> 'x = {:#^10b}'.format(10)
'x = ###1010###'
>>> 'x = {:#^10o}'.format(10)
'x = ####12####'
>>> 'x = {:#^10x}'.format(10)
'x = ####a#####'
# 对齐字符控制填充字符
>>> 'x = {:^#010x}'.format(10)
'x = 0000xa0000'
# 对齐字符之前的填充字符
>>> 'x = {:9^#014b}'.format(10)
'x = 99990b10109999'
# 没有对齐字符的填充字符
>>> 'x = {:#014b}'.format(10)
'x = 0b000000001010'

# 显示进制前缀
>>> 'x = {:10b}'.format(100)
'x =    1100100'
>>> 'x = {:#10b}'.format(100) # 显示二进制带前缀 0b
'x =  0b1100100'
>>> 'x = {:+#10b}'.format(100)
'x = +0b1100100'
>>> 'x = {:#10o}'.format(100) # 显示八进制带前缀 0o
'x =      0o144'
>>> 'x = {:#10x}'.format(100) # 显示十六进制带前缀 0x
'x =       0x64'
```

#### f-string

> python 3.6 以上支持

不需要考虑浮点数的精度问题

```python
# 字符串拼接, f-string, 精度控制可选
message = f"my salary is {salary:.2f}, my tel is {tel}"
print(message)
print(F"""my salary is {salary:.2f}, my tel is {tel}""")

age = 18
>>> print(f'''age: {age}''')
age: 18
>>> print(f"""age: {age}""")
age: 18
>>> print(F"""age: {age}""")
age: 18
>>> print(F'''age: {age}''')
age: 18
```

- `变量名=` 输出变量名及值, python 3.8 支持

```python
>>> thing = 'woodchuck'
>>> place = 'lake'
>>> f'The {thing=:>10s} is in the {place=:!^10s}'
'The thing= woodchuck is in the place=!!!lake!!!'
```

**占位符规则** <!--markdownlint-disable-line-->

规则和 '{}'.format 一致

```python
>>> thing = 'woodchuck'
>>> place = 'lake'
>>> f'The {thing:>10s} is in the {place:!^10s}'
'The  woodchuck is in the !!!lake!!!'

# 显示进制前缀
>>> f'{x = :^+10}'
'x =    +100   '
>>> f'{x = :^+10b}' 
'x =  +1100100 '
>>> f'{x = :^+#10b}'  # 居中显示二进制带符号带前缀 0b
'x = +0b1100100'
>>> f'{x = :^+#10o}'  # 居中显示八进制带符号带前缀 0o
'x =   +0o144  '
>>> f'{x = :^+#10x}'  # 居中显示十六进制带符号带前缀 0x
'x =   +0x64   '
```

#### join 方法拼接字符串

拼接的元素只能是`字符串类型`

```python
>>> t = ['get gloves,get mask.give cat vitamins,call ambulance', 'get glove', 'get mak']
>>> t2 = '--gg,gg--'
>>> t2.join(t)
'get gloves,get mask.give cat vitamins,call ambulance--gg,gg--get glove--gg,gg--get mak'
>>> t3 = t2.join(t)
>>> t3.split(t2)
['get gloves,get mask.give cat vitamins,call ambulance', 'get glove', 'get mak']

>>> t4 = [1, 2, 3, True]
>>> type(t4)
<class 'list'>
>>> ','.join(t4)
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: sequence item 0: expected str instance, int found
```
