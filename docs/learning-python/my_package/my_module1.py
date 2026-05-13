__all__ = ['name']

name = 'my_module1'


def info_print1():
    print("这是 my_module1 的 info_print1 函数...")
    name = 'info_print1'
    print(f'info_print1 locals(): {locals()}')


# 只有在文件直接运行时, __name__ 内置变量值才是 '__main__'
# 否则作为模块导入时, __name__ 内置变量值是模块名
if __name__ == '__main__':
    print(f'my_package 包下的 my_module1 模块 globals(): {globals()}')
    print(f'my_package 包下的 my_module1 模块 locals(): {locals()}')
    info_print1()
