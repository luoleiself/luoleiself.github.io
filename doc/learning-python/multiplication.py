print('-----9 * 9-----')
print('-----while-----')
i = 1
while i <= 9:
    j = 1
    while j <= i:
        print(f'{j} * {i} = {j * i}', end='\t')
        j += 1
    print()
    i += 1

print('---for-range---')
for i in range(1, 10):
    for j in range(1, i + 1):
        print(f'{j} * {i} = {j * i}', end='\t')
    print()

print('----------------')

"""
猴子第一天摘了若干个桃子, 当即吃了一半儿零一个, 以后每一天都吃掉前一天剩下的一半儿零一个.
到第10天早上还想吃时, 只剩下一个桃子, 求第一天共摘了多少个桃子?
"""
peaches = 1 # 第10天早上剩下桃子的数量
for day in range(9, 0, -1):
    # 逆向计算前一天开始时的桃子数量
    peaches = (peaches + 1) * 2
    print(day, peaches)

print(f'第一天共摘了 {peaches} 个桃子')
print('----------------')


# 计算日期是一年中的第几天
def day_in_year(year, month, day):
    days_in_month = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    if (year % 4 == 0 and year % 100 != 0) or year % 400 == 0:
        days_in_month[1] = 29

    days_count = sum(days_in_month[:month - 1]) + day
    return days_count


day1 = day_in_year(2023, 3, 20)
day2 = day_in_year(2024, 3, 20)
print(f'2023年3月20日是第 {day1} 天')
print(f'2024年3月20日是第 {day2} 天')
print('----------------')
