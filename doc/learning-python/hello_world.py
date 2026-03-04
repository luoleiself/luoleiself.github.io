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

import random
print('hello world')

name = "luolei"
age = 12
print(f'my name is {name}, and I am {age} years old')

print('----------')

# 字符串拼接, 加号 +
message = "hello " + name
print(message)

# 字符串拼接, 占位符 %,  %s 是字符串占位符，%d 是整数占位符，%f 是浮点数占位符
# 需要考虑浮点数的精度问题
name = "luolei"
message = "hello %s" % name
print(message)

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

# print('----------')
# age = int(input("请输入你的年龄: "))
# if age >= 18:
#     print("你已经成年了，可以独立生活了。")
# else:
#     print("你还未成年")

# print('----------')
# num = random.randint(1, 10)
# guess_num = int(input("请输入一个 1-10 以内的数字: "))
# if guess_num == num:
#     print('恭喜你，猜对了!!')
# else:
#     if guess_num > num:
#         print('你猜的数字太大了')
#     else:
#         print('你猜的数字太小了')
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

# 定义加法函数
def add(a: int, b: int) -> int:
    '''
    @method add 函数求和
    @param a: int
    @param b: int
    @return int
    '''
    return a + b


result = add(1, 3)
print(f'result: {result}')
