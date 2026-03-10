'''使用 else 检查 while 和 for 的 break, try 的错误'''
print('while 循环没有执行 break, 则执行 else 块')
numbers = [1, 3, 5, 7, 9]
position = 0
while position < len(numbers):
    number = numbers[position]
    if number % 2 == 0:
        print(f'Found even number {number}')
        break
    position += 1
else:  # while 循环没有执行 break, 则执行 else 块
    print('No even number found')
print('----------------------------')

print('for 循环没有执行 break, 则执行 else 块')
word = 'thud'
for letter in word:
    if letter == 'x':
        print('Found an x!')
        break
    print(letter)
else:  # for 循环没有执行 break, 则执行 else 块
    print('No x found')
print('----------------------------')


# fibonacci 数列 0 1 1 2 3 5 8 13 21 34 55 89 ...
def fibonacci(n):
    a, b = 0, 1
    while n > 0:
        a, b = b, a+b
        print(f'{a} ', end=' ')
        n -= 1
    print()
    return a


print(f'fibonacci(10) {fibonacci(10)}')
print('----------------------------')

print('try 如果没有发生异常，则执行 else 块')
try:
    print(1 / 0)
except ZeroDivisionError as e:
    print('除以 0 的错误:', e)
except Exception as e:  # 捕获所有异常
    print('发生其他错误:', e)
else:
    print('没有发生错误执行的代码')
finally:
    print('总是会执行')
print('---------')
try:
    fr = open('./uv.lock', 'r', encoding='utf-8')
except FileNotFoundError as e:
    print('文件未找到:', e)
else:
    print('else: 文件打开成功')
    txt = fr.readline()
    print(f'读取的内容：\n{txt}')
finally:
    print('总是会执行')
    fr.close()
print('----------------------------')


# 自定义异常类
class MyError(Exception):
    def __init__(self, value):
        self.value = value

    def __str__(self):
        return f'MyError: {self.value}'


try:
    raise MyError('my error')
except MyError as e:
    print(f'自定义异常 {e}')
else:
    print('没有发生错误执行的代码')
print('----------------------------')
