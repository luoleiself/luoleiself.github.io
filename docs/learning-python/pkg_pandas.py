import pandas as pd
import numpy as np
import matplotlib.pyplot as plt

"""
pandas: 数据操作和数据分析的 python 库
numpy: 科学计算的 python 库, 核心功能包含 n 维数组, 矩阵运算, 线性代数等

DataFrame:  表格
Series: 列
Index: 索引

df = pd.read_*() # 从文件中读取数据
df.to_*()   # 将数据写入到文件中
pd.melt()   # 宽表转换为长表
pd.pivot()  # 长表转换为宽表
pd.DataFrame()  # 创建表
pd.Series() # 创建列
pd.Timestamp()  # 日期时间

选择器: .str/.dt/.cat

行列转置: .T

索引/列/值: .index/.columns/.values

基本信息预览: .shape/.size/.ndim/len()/.dtypes/.describe()/.info()
预览表格数据: .head()/.tail()/.sample()
类型转换: .astype()
分箱: .cut()
判断: .isin()/.isna()/.isnull()/.between()/.between_time()
计算: .sum()/.mean()/.min()/.max()/.abs()/.var()/.std()/.median()/.replace()/.keys()/.diff()
某列最大/最小的 n 条数据:     .nlargest()/.nsmallest()
统计信息:   .unique()/groupby()/apply()/.agg()/.count()/.value_counts()
索引位置:   .set_index()/.reset_index()
排序: .sort_values()/.sort_index()/.rank()

删除: .drop()/.dropna()/.droplevel()
重复: .duplicated()/.drop_duplicates()
数据填充: .fillna()/.ffill()/.bfill()

滑动: .rolling()

日期序列: .date_range(start, periods='', freq='')
采样: .resample() # 时间序列频率转换和重采样方法, 对象必须具有类似日期时间的索引,
                    或者调用者必须将类似日期时间系列/索引的标签传递给on/level关键字参数.

收益率: .pct_change()

表格拼接: concat    # 根据行列的方向自然拼接
        join    # 根据行的索引 Index 进行拼接
        merge   # 根据列的值为依据进行拼接
        
数据选取: loc[]/iloc[], at[]/iat[]  # 以 i 开头的属性表示下标方式访问
"""

df_1 = pd.DataFrame([
    {'col_1': 1, 'col_2': 2, 'col_3': 3},
    {'col_1': 4, 'col_2': 5, 'col_3': 6},
    {'col_1': 7, 'col_2': 8, 'col_3': 9}
])
print(df_1)
print(f'df_1.index {df_1.index}')
print(f'df_1.columns {df_1.columns}')
print(f'df_1.values {df_1.values}')
print('-' * 4)
nlargest = df_1.nlargest(1, 'col_1')
print(f'df_1.nlargest(1, "col_1") {nlargest}')
nsmallest = df_1.nsmallest(1, 'col_1')
print(f'df_1.nsmallest(1, "col_1") {nsmallest}')
print('-' * 4)
print(f'df_1.T {df_1.T}')
print(f'df_1.T.index {df_1.T.index}')
print(f'df_1.T.columns {df_1.T.columns}')
print(f'df_1.T.values {df_1.T.values}')
print('-' * 10)

df_2 = pd.DataFrame(
    [(1, 2, 3), (4, 5, 6), (7, 8, 9)],
    columns=['col_1', 'col_2', 'col_3'],
    index=['a', 'b', 'c'])
print(df_2)
print(f'df_2.T {df_2.T}')
print('-' * 10)

df_3 = pd.DataFrame({
    'col_1': (1, 4, 7),
    'col_2': [2, 5, 8],
    'col_3': [3, 6, 9]
}, index=['a', 'b', 'c'])
print(df_3)
print(f'df_3.T {df_3.T}')
print('-' * 10)

print(df_3['col_2'])
print(df_3[['col_1', 'col_3']])
print(df_3[0:2:1])
print('-' * 10)

print('head(1)', df_3.head(1))
print('tail(2)', df_3.tail(2))
print('shape', df_3.shape)
print('size', df_3.size)
print('dtypes', df_3.dtypes)
print('len()', len(df_3))
print('describe()', df_3.describe())
df_3.info()
print('-' * 10)

print(df_3[df_3['col_3'].isin([6])])
print(df_3[df_3['col_2'].between(2, 8)])
print(df_3[df_3['col_2'].between(2, 8, inclusive='left')])
print(df_3[df_3['col_2'].between(2, 8, inclusive='right')])
print(df_3[df_3['col_2'].between(2, 8, inclusive='neither')])
print('-' * 10)

print(df_3[df_3['col_1'] >= 4])
print(df_3[(df_3['col_1'] >= 4) & (df_3['col_3'] >= 9)])
print('-' * 10)

print('loc\n', df_3.loc['a':'c':2])  # 索引访问
print('iloc\n', df_3.iloc[0:2])  # 下标访问
print('at', df_3.at['b', 'col_2'])  # 通过 row/column 访问单个值
print('iat', df_2.iat[1, 2])  # 通过 row/column 下标访问单个值
print('-' * 10)

print(df_3.isnull())
print(df_3.isna())
print(df_3.fillna('--'))
print(df_3.ffill())  # 使用上一个非空值填充
print(df_3.bfill())  # 使用下一个非空值填充

print(df_3.duplicated())  # 查看重复值
df_3.drop_duplicates(inplace=True)  # 删除重复值
print('-' * 10)

df_3 = df_3.sort_values('col_1', ascending=False)
print(df_3)
df_3 = df_3.sort_values(['col_1', 'col_3'], ascending=[False, True])
print(df_3)
print('-' * 10)

print(df_3.groupby('col_1')['col_1'].sum())
print('-' * 10)

# numpy
"""
矩阵点积: @
np.array()  # 创建数组
np.zeros()  # 预定义形状, 全 0
np.ones()   # 预定义形状, 全 1
np.empty()  # 预定义形状, 值不固定
np.full()   # 预定义形状, 指定值
# 根据给定的数组结构创建形状
np.zeros_like()
np.ones_like()
np.empty_like()
np.full_like()  

np.arange() # 创建序列
np.linspace()   # 创建相同间隔的序列区间
np.logspace()   # 创建相同间隔的对数序列区间

np.random.rand()   # 生成 0 到 1 之间的随机数组(均匀分布)
np.random.randn()   # 生成 0 到 1 之间的随机数组(正态分布(两边概率小, 中间概率大))
np.random.uniform() # 生成指定范围区间的随机浮点数数组
np.random.randint() # 生成指定范围区间的随机整数数组
np.random.seed()    # 设置随机种子

np.copy()   # 数组拷贝
np.reshape()    # 转换数组结构
np.concatenate()    # 拼接数组

np.max()
np.min()
np.mean()   # 平均数
np.median() # 中位数
np.std()    # 标准差
np.sum()

np.greater()
np.less()
np.equal()
np.logical_and()
np.where()
np.cumsum() # 累加
np.count_nonzero()  # 统计
np.unique()
np.sort()
np.argsort()
"""

arr = np.array([[1, 2, 3], [4, 5, 6]])
print(f'np.array() {arr}')
arr_1 = np.zeros((2, 3))
arr_2 = np.ones((2, 3))
arr_3 = np.empty((2, 3))
arr_4 = np.full((2, 3), 5)
print(f'np.zeros() {arr_1}')
print(f'np.ones() {arr_2}')
print(f'np.empty() {arr_3}')
print(f'np.full() {arr_4}')
print('-' * 4)

arr_5 = np.zeros_like(arr_1, dtype=float)
arr_6 = np.ones_like(arr_2, dtype=float)
arr_7 = np.empty_like(arr_3, dtype=float)
arr_8 = np.full_like(arr_4, 2026)
print(f'np.zeros_like() {arr_5}')
print(f'np.ones_like() {arr_6}')
print(f'np.empty_like() {arr_7}')
print(f'np.full_like() {arr_8}')
print('-' * 4)

arr_9 = np.arange(1, 10, 3)
print(f'np.arange() {arr_9}')
arr_10 = np.linspace(1, 10, 4)
print(f'np.linspace(1, 10, 4) {arr_10}')
arr_11 = np.linspace(60, 100, 5)
print(f'np.linspace(60, 100, 5) {arr_11}')
arr_12 = np.logspace(1, 10, 4, base=2)
print(f'np.logspace(1, 10, 4, base=2) {arr_12}')
print('-' * 4)

arr_13 = np.random.rand(2, 3)
print(f'np.random.rand(2, 3) {arr_13}')
arr_14 = np.random.randn(2, 3)
print(f'np.random.randn(2, 3) {arr_14}')
arr_15 = np.random.uniform(1, 10, (2, 3))
print(f'np.random.uniform(1, 10, (2, 3)) {arr_15}')
arr_16 = np.random.randint(1, 10, (2, 3))
print(f'np.random.randint(1, 10, (2, 3)) {arr_16}')
print('-' * 4)

np.random.seed(1024)
arr_17 = np.random.randint(1, 10, size=(2, 5))
print(f'np.random.seed(1024) np.random.randint(1, 10, size=(2, 5)) {arr_17}')
print('-' * 4)

# 数组的算术运算: 相同维度的元素数量必须一致
# operands could not be broadcast together with shapes (4,) (3,)
arr_18 = np.array([1, 2, 3])
arr_19 = np.array([4, 5, 6])
print(f'arr_18 + arr_19 {arr_18 + arr_19}')
print(f'arr_18 - arr_19 {arr_18 - arr_19}')
print(f'arr_18 * arr_19 {arr_18 * arr_19}')
print(f'arr_18 / arr_19 {arr_18 / arr_19}')
print(f'arr_18 ** arr_19 {arr_18 ** arr_19}')

print('-' * 10)

# matplotlib
"""
plt.figure()    # 设置画布
plt.subplots()  # 创建子图和轴
plt.show()  # 显示图

plt.rcParams    # 设置

plot()  # 折线图
bar()   # 柱状图
pie()   # 饼图
scatter()   # 散点图
boxplot()   # 箱线图

.grid()  # 设置网格
.legend()   # 设置图例
.text() # 文字
.tight_layout() # 自适应排版
"""
# 直接绘制折线图
# plt.plot([1, 2, 3, 4], [0, 0.5, 1, 0.2])
# plt.title("Sample plot")
# plt.show()

# 创建 画布 和 轴
fig, axs = plt.subplots(nrows=2, ncols=2, figsize=(8, 6))  # nrows=1, ncols=1, figsize=(8, 6)
fig.suptitle('hello world', fontsize=13)  # 画布标题
fig.subplots_adjust()  # 调整子图间距

plt.rcParams['font.sans-serif'] = ['SimHei']  # 设置字体, 解决中文乱码

# 标题
axs[0, 0].set_title("Sample plot", fontsize=12, fontweight='bold')
axs[0, 0].plot([1, 2, 3, 4], [0, 0.5, 1, 0.2], label='plot-1', color='red', marker='o', markersize=5)  # 绘制折线图
axs[0, 0].plot([1, 2, 3, 4], [0.8, 1.2, 0.1, 1.3], label='plot-2', marker='*', markersize=5)
# axs[0, 0].set_xticks()  # 设置 x 轴值
# axs[0, 0].set_xticklabels(['a', 'b', 'c', 'd'])  # 设置 x 轴显示值
# axs[0, 0].set_yticks([0, 0.5, 1, 0.2], rotation=45)  # 设置 y 轴值
# axs[0, 0].set_yticklabels(['0.0', '0.5', '1.0', '0.2'])  # 设置 y 轴显示值
axs[0, 0].set_xlabel('x 轴')  # x 轴名称
axs[0, 0].set_ylabel('y 轴')  # y 轴名称

axs[0, 0].grid(True, linestyle='--')  # 网格
axs[0, 0].legend(loc='upper right', fontsize=10)  # 图例

# 柱状图
axs[0, 1].bar([1, 2, 3, 4], [0, 0.5, 1, 0.2])
axs[0, 1].set_title("Sample bar")

plt.show()
