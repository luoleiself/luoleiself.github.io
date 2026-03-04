money = 500000


def main():
    print('---------main---------')
    print('欢迎来到银行系统')
    print('查询余额\t[输入1]\t')
    print('存款\t\t[输入2]\t')
    print('取款\t\t[输入3]\t')
    print('退出\t\t[输入4]\t')
    return input('请输入您的选择: ')

# 查询余额
def query_balance(show_header: bool):
    if show_header:
        print('---------查询余额---------')
    print(f'您的余额为: {money}')

# 存款
def saving_money(num: float):
    print('---------存款---------')
    global money
    money += num
    query_balance(False)

# 取款
def withdraw_money(num: float):
    print('---------取款---------')
    global money
    money -= num
    if money < 0:
        print('余额不足')
        return False  # 取款失败
    query_balance(False)
    return True   # 取款成功


while True:
    choice = main()
    if choice == '1':
        query_balance(True)
        continue
    elif choice == '2':
        num = float(input('想要存多少钱: '))
        saving_money(num)
        continue
    elif choice == '3':
        num = float(input('想要取多少钱: '))
        withdraw_money(num)
        continue
    elif choice == '4':
        print('拜拜，欢迎下次再来')
        break
    else:
        print('无效的选择，请重新输入')
