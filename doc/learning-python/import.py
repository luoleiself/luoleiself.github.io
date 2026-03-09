import sys
for path in sys.path:
    print(f' {path}')

# # 导入方式1: import 导入包下的模块
# import my_package.my_module1
# import my_package.my_module2
# # 使用时: 包名.模块名.函数名
# my_package.my_module1.info_print1()
# my_package.my_module2.info_print2()

# # 导入方式2: import 导入包下的模块并指定别名
# import my_package.my_module1 as my_m1
# import my_package.my_module2 as my_m2
# # 使用时: 别名.函数名
# my_m1.info_print1()
# my_m2.info_print2()

# 导入方式3: from ... import ... 导入包下的模块
from design_mode.singleton import *  # 模块内的 __all__ 控制模块 import * 的行为
from my_package import my_module1, my_module2
# 使用时: 模块名.函数名
my_module1.info_print1()
my_module2.info_print2()

# # 导入方式4: from ... import ... 导入包下的模块并指定别名
# from my_package import my_module1 as my_m1, my_module2 as my_m2
# # 使用时: 别名.函数名
# my_m1.info_print1()
# my_m2.info_print2()

# # 导入方式5: from ... import * 导入包下的模块的所有内容
# from my_package.my_module1 import *
# from my_package.my_module2 import *
# # 使用时: 函数名
# info_print1()
# info_print2()

# # 在 __init__.py 中使用 __all__ 控制包级 import * 的行为
# from my_package import *
# # 只能访问 __all__ 中列出的名称
# print(func_1())
# print(func_2())
# # print(func_3()) # NameError: name 'func_3' is not defined.
# 在 my_module1 模块中添加 __all__ 内置变量控制模块 import * 的行为
# from my_package.my_module1 import *
# info_print1()   # NameError: name 'info_print1' is not defined

print('---------------------------------------------------')
# 模块中 __all__ 列表中没有添加 StrTools 类
# s = StrTools()    # NameError: name 'StrTools' is not defined. Did you mean: 'str_tools'?
s1 = str_tools
s2 = str_tools
print(f's1 {s1}')
print(f's2 {s2}')
