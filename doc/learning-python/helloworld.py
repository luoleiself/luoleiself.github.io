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

print('hello world')

name="luolei"
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
import random
# num = random.randint(1, 10)
# guess_num = int(input("请输入一个 1-10 以内的数字: "))
# if guess_num == num:
#     print('恭喜你，猜对了!!')
# else:
#     if guess_num > num:
#         print('你猜的数字太大了')
#     else:
#         print('你猜的数字太小了')

print('-----random guess-----')
random_num = random.randint(1, 100)
loop = 1
while loop <= 100:
    print(f'Python 天下无敌：{loop}')
    if loop == random_num:
        print(f'Python 克星找到了, 它是 {random_num}')
        break
    loop += 1

print('-----9 * 9-----')
i = 1
while i <= 9:
    j = 1
    while j <= i:
        print(f'{j} * {i} = {j * i}', end='\t')
        j += 1
    print()
    i += 1

print('--------------')
letters = 'abcdefghijklmnopqrstuvwxyz'
for x in letters:
    print(f'{x}', end='\t')
print()
for x in range(97, 123):
    print(f'{x}={chr(x)}', end='\t')
print()
for x in range(65, 91):
    print(f'{x}={chr(x)}', end='\t')
print()

print('--------------')
print(ord('一'))    # 获取字符的 unicode 码
