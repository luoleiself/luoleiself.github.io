name = 'my_module2'


def info_print2():
    print("这是 my_module2 的 info_print2 函数...")
    name = 'info_print2'
    print(f'info_print2 locals(): {locals()}')


# 只有在文件直接运行时, __name__ 内置变量值才是 '__main__'
# 否则作为模块导入时, __name__ 内置变量值是模块名
if __name__ == '__main__':
    print(f'my_package 包下的 my_module2 模块 globals(): {globals()}')
    print(f'my_package 包下的 my_module2 模块 locals(): {locals()}')
    info_print2()
