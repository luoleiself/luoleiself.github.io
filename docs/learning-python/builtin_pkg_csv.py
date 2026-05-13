import csv

"""
csv 模块: 读写 csv 文件
    reader(f)   # 创建一个 csv 读取迭代器
    writer(f)   # 创建一个 csv 写入对象
        csv_out.writerow(row)   # 写入一行
        csv_out.writerows(rows) # 写入多行
        
    DictReader(f)   # 创建一个 csv 字典读取迭代器
        csv_input.fieldnames    # 输出列名
    DictWriter(f, fieldnames)   # 创建一个 csv 字典写入对象
        csv_out.writeheader()   # 写入列名 fieldnames
        csv_out.writerow(dict)  # 写入一行
        csv_out.writerows(dict_list)    # 写入多行
    
    list_dialects() # 列出所有方言
    field_size_limit()  # 输出字段大小限制
    QUOTE_ALL   # 输出字段所有内容
    QUOTE_MINIMAL   # 输出字段最小内容
"""

print(f'csv.list_dialects() 列出所有方言 {csv.list_dialects()}')
print(f'csv.field_size_limit() 输出字段大小限制 {csv.field_size_limit()}')  # 输出字段大小限制
print(f'csv.QUOTE_ALL 输出字段所有内容 {csv.QUOTE_ALL}')  # 输出字段所有内容
print(f'csv.QUOTE_MINIMAL 输出字段最小内容 {csv.QUOTE_MINIMAL}')  # 输出字段最小内容
print('-' * 20)
print('读写多项序列的列表')
csv_data = [['a', 'b', 'c'], (1, 2, 3), (4, 5, 6)]
print(f'csv_data = {csv_data}')
with open('test_1.csv', 'wt', newline='') as csvfile:
    csv_out = csv.writer(csvfile, strict=True)
    csv_out.writerow(['col 1', 'col 2', 'col 3'])  # 写入一行
    csv_out.writerows(csv_data)  # 写入多行

with open('test_1.csv', 'rt', newline='') as csvfile:
    csv_input = csv.reader(csvfile, strict=True)
    # 迭代读取结果
    for row in csv_input:
        print(f'row = {row}')
print('-' * 20)

print('读写字典列表')
with open('test_2.csv', 'wt', newline='') as csvfile:
    # fieldnames 指定列名
    fieldnames = ['col 1', 'col 2', 'col 3']
    csv_out = csv.DictWriter(csvfile, fieldnames=fieldnames, strict=True)
    csv_out.writeheader()  # 写入列名
    csv_out.writerow({'col 1': 'a', 'col 2': 'b', 'col 3': 'c'})  # 写入一行
    csv_out.writerows([
        {'col 1': 1, 'col 2': 2, 'col 3': 3},
        {'col 1': 4, 'col 2': 5, 'col 3': 6},
    ])  # 写入多行
with open('test_2.csv', 'rt', newline='') as csvfile:
    csv_input = csv.DictReader(csvfile, strict=True)
    print(f'csv_input.fieldnames {csv_input.fieldnames}')
    for row in csv_input:
        print(
            f'row["col 1"] {row["col 1"]}, row["col 2"] {row["col 2"]}, row["col 3"] {row["col 3"]}')

print('-' * 40)
