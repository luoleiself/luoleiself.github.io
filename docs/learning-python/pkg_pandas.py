# pandas
import pandas as pd
import matplotlib.pyplot as plt

"""
DataFrame:  表格
Series: 列

df = pd.read_*() # 从文件中读取数据
df.to_*()   # 将数据写入到文件中
"""

df_1 = pd.DataFrame([
    {'col_1': 1, 'col_2': 2, 'col_3': 3},
    {'col_1': 4, 'col_2': 5, 'col_3': 6},
    {'col_1': 7, 'col_2': 8, 'col_3': 9}
])
print(df_1)
print('-' * 10)

df_2 = pd.DataFrame(
    [(1, 2, 3), (4, 5, 6), (7, 8, 9)],
    columns=['col_1', 'col_2', 'col_3'],
    index=['a', 'b', 'c'])
print(df_2)
print('-' * 10)

df_3 = pd.DataFrame({
    'col_1': [1, 4, 7],
    'col_2': [2, 5, 8],
    'col_3': [3, 6, 9]
}, index=['a', 'b', 'c'])
print(df_3)
print('-' * 10)

print(df_3['col_2'])
print(df_3[['col_1', 'col_3']])
print(df_3[0:2:1])
print('-' * 10)

print(df_3.head(1))
print(df_3.tail(2))
print(df_3.describe())
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

print(df_3.iloc[0:2])  # 下标访问
print(df_3.loc['a':'c':2])  # 索引访问
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

# matplotlib

# 创建 画布 和 轴
fig, ax = plt.subplots()  # nrows=1, ncols=1, figsize=(8, 6)
fig.suptitle()  # 画布标题
fig.subplots_adjust()  # 调整子图间距
ax.plot([1, 2, 3, 4], [0, 0.5, 1, 0.2])  # 绘制折线图
ax.set_title("Sample plot")
plt.show()

# 直接绘制折线图
plt.plot([1, 2, 3, 4], [0, 0.5, 1, 0.2])
plt.title("Sample plot")
plt.show()

# 标题
ax.set_title("Sample plot", fontsize=12, fontweight='bold')

ax.set_xticks()  # 设置 x 轴值
ax.set_xticklabels(['a', 'b', 'c', 'd'])  # 设置 x 轴显示值
ax.set_yticks([0, 0.5, 1, 0.2])  # 设置 y 轴值
ax.set_yticklabels(['0.0', '0.5', '1.0', '0.2'])  # 设置 y 轴显示值
ax.set_xlabel()  # x 轴名称
ax.set_ylabel()  # y 轴名称

ax.grid(True, linestyle='--')  # 网格
ax.legend()  # 图例

plt.rcParams['font.sans-serif'] = ['SimHei']  # 设置字体, 解决中文乱码
