
项目结构

```python
my_project/
├── .venv/                 # 虚拟环境目录 (通常跟随项目)
│   ├── bin/               # (Linux/macOS) 可执行文件
│   ├── Scripts/           # (Windows) 可执行文件
│   ├── lib/               # (Linux/macOS) 存放安装的包
│   ├── Lib/               # (Windows) 存放安装的包
│   └── pyvenv.cfg         # 虚拟环境配置文件
├── .python-version        # (可选) 指定项目所需的 Python 版本
├── pyproject.toml         # 项目的核心配置文件 (记录依赖声明)
├── uv.lock                # 依赖的精确锁定文件 (由 uv 自动维护)
├── requirements.txt       # (可选) 传统的依赖列表文件
|—— 项目源代码
```

## 数据类型

- int
- float
- bool, 底层使用整数 1, 0 表示 True, False
- str
- tuple
- list
- set
- dict

```python
# 单行注释

'''
多行字符串/注释, 如果多行字符串有变量引用则作为字符串, 否则作为多行注释
'''

"""
多行字符串/注释, 如果多行字符串有变量引用则作为字符串, 否则作为多行注释
"""
```

- `整型和布尔值混合转换为整型`
- `整型和浮点型混合转换为浮点数`
- `布尔值和浮点型混合转换为浮点数`

```python
# 整型和布尔值
>>> 1 + True
2
>>> 1 + False
1

# 整型和浮点型
>>> 1. + 2
3.0
>>> 1 + 1.2
2.2

# 布尔值和浮点型
>>> 1. + True 
2.0
>>> 1.0 + True 
2.0
>>> 1.0 + False
1.0
```

### 整数

range 是`惰性求值`的，不会一次性占用大量内存

```python
# 立即返回，几乎不占内存
big_range = range(1000000)
# 创建包含 100 万个元素的列表，占用大量内存
big_list = list(range(1000000))
```

- range() 生成数字序列，返回一个可迭代对象, 不能直接使用, 需要`配合其他函数或语句`才能获取值
  - 只有 1 个参数时生成 0 到 arg 的数字序列
  - start, 开始，包含
  - end, 结束，不包含
  - step，步长绝对值，默认为 1
    - 如果为 -1, `反向生成数字序列`

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

### 内置函数

#### int()

将任意值转换成整型化值

- 第一个参数为要转换的值
  - 浮点数则截断整数部分
  - 参数如果是字符串, 只能由 `数字, +, -, _` 组成
  - 布尔值 True 转为 1, False 转为 0
- 第二个参数为进制数位

```python
>>> int(1.)
1
>>> int(1.8)
1
>>> int(1_0_0)
100
>>> int("+1_0")
10
>>> int(1e10)
10000000000
>>> int(True)
1
>>> int(False)
0
```

#### float()

将任意值转换为浮点化值

- 参数为整型则转换为浮点数
- 参数为字符串的浮点数表示形式，支持科学计数格式
- 布尔值 True 转为 1.0, False 转为 0.0

```python
>>> float(1)
1.0
>>> float(1.)
1.0
>>> float('1')
1.0
>>> float('1.')
1.0
>>> float('+1')
1.0
>>> float('-1')
-1.0
>>> float('+1.')
1.0
>>> float('-1.')
-1.0
>>> float('+1.0')
1.0
>>> float('-1.0')
-1.0
>>> float('-1.0_0')
-1.0
>>> float('-1_0.0_0')
-10.0
>>> float('1.e10')
10000000000.0
>>> float(True)
1.0
>>> float(False)
0.0
```

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
```

#### chr()/ord()

转换 unicode 码对应的字符

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
