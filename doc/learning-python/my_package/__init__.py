# 控制 import * 的行为
__all__ = ['func_1', 'func_2']

from . import my_module1, my_module2

def func_1():
    return 'function 1'


def func_2():
    return 'function 2'


def func_3():
    return 'function 3'
