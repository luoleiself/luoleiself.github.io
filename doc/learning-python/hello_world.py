"""
This is a multi-line comment in python, which can be used to write longer comments or documentation.

Author: luolei
Date: 2024-06-01
"""

'''
This is a multi-line comment in python, which can be used to write longer comments or documentation.

Author: luolei
Date: 2024-06-01
'''

import os
import random
import sys

# 需要考虑浮点数的精度问题
salary = 5000
tel = 13112345678
message = "my salary is %.2f, my tel is %d" % (salary, tel)
print(message)

# 字符串拼接, format方法
message = "my salary is {:.2f}, my tel is {}".format(salary, tel)
print(message)

# 字符串拼接, f-string, 不需要考虑浮点数的精度问题
message = f"my salary is {salary:.2f}, my tel is {tel}"
print(message)
print(F"""my salary is {salary:.2f}, my tel is {tel}""")
print('--------------')

letters = 'abcdefghijklmnopqrstuvwxyz'
for x in letters:
    print(f'{x}', end='\t')
print()
for x in range(97, 123):    # 临时变量, 外层可以访问到, 不建议在外层访问
    print(f'{x}={chr(x)}', end='\t')    # chr 获取 unicode 码对应的字符
print()
# print(x)    # 能访问到循环语句内声明的变量
for x in range(65, 91):
    print(f'{x}={chr(x)}', end='\t')
# print(x)    # 能访问到循环语句内声明的变量
print(ord('一'))    # ord 获取字符的 unicode 码
print('--------------')

print('-----random guess-----')
random_num = random.randint(1, 100)
loop = 1
while loop <= 100:
    print(f'Python 天下无敌：{loop}')
    if loop == random_num:
        print(f'Python 克星找到了, 它是 {random_num}')
        break
    loop += 1
print('--------------')

# 获取环境变量和命令行参数
# uv run --env-file .env .\hello_world.py 'helloworld' 'hellochina'
print(f'os.environ: {os.environ}')
if os.environ.get('TOKEN'):
    print(f'os.environ["TOKEN"]: {os.environ["TOKEN"]}')
print(f'sys.argv: {sys.argv}')
print(f'sys.orig_argv: {sys.orig_argv}')
print(f'sys.api_version: {sys.api_version}')
print(f'sys.version: {sys.version}')

lt = [i for i in range(10000)]
print(f'列表推导式内存大小: sys.getsizeof(lt): {sys.getsizeof(lt)}')
lt_gen = (i for i in range(10000))
print(f'生成器内存大小: sys.getsizeof(lt_gen): {sys.getsizeof(lt_gen)}')
print('--------------')
