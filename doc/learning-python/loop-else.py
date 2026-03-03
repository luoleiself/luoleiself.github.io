print('使用 else 检查 while 和 for 的 break')
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
print('-----')
print('for 循环没有执行 break, 则执行 else 块')
word = 'thud'
for letter in word:
    if letter == 'x':
        print('Found an x!')
        break
    print(letter)
else:  # for 循环没有执行 break, 则执行 else 块
    print('No x found')
