import argparse

"""
argparse 参数解析
    # prop=''   # 命令的名称, 未指定默认由 __main__ 和 sys.argv[0] 生成
    # decription='' # 命令的描述
    # add_help=False    # 禁用添加 -h 选项
    ArgumentParser(prog='', description='')    # 实例化参数解析
        .add_argument()  # 添加参数
        .add_argument_group()   # 创建参数分组
        .add_mutually_exclusive_group()  # 创建互斥参数分组, 互斥组中参数同时只能出现一个, 否则报错
        .set_defaults()  # 添加附加参数
        .parse_args()    # 解析参数
        .print_help()    # 打印帮助信息

        subparsers = .add_subparsers()    # 创建子命令解析器
        parser_a = subparsers.add_parser()    # 添加子命令解析器
"""
parser = argparse.ArgumentParser(description='参数解析')
parser.add_argument('-a', '--arg1', default='1', help='arg1 help')
parser.add_argument('-b', '--arg2', default='2', help='arg2 help')

# 添加附加参数
parser.set_defaults(arg7='arg7 value')

# 创建参数分组
group1 = parser.add_argument_group('group1')
# 向分组添加参数
group1.add_argument('-c', '--arg3', default='3', help='arg3 help')
group2 = parser.add_argument_group('group2')
group2.add_argument('-d', '--arg4', default='4', help='arg4 help')

# 创建互斥参数分组
group3 = parser.add_mutually_exclusive_group()
group3.add_argument('-e', '--arg5', default='5', help='arg5 help')
group3.add_argument('-f', '--arg6', default='6', help='arg6 help')

# 创建子命令解析器
subparsers = parser.add_subparsers(dest='subcommand help')
# 添加子命令解析器
parser_a = subparsers.add_parser('a', help='a help')
parser_a.add_argument('-a', '--a_arg1', help='a_arg1 help')
# 添加子命令解析器
parser_b = subparsers.add_parser('b', help='b help')
parser_b.add_argument('-b', '--b_arg1', help='b_arg1 help')

# 输出帮助信息
# parser.parse_args(['-h'])
parser.print_help()

# builtin_pkg_argparse.py: error: argument -e/--arg5: expected one argument
# parser.parse_args(['-e', '-f'])   # 互斥组中的参数只能出现一个, 否则报错
print('-' * 6)

# 输出子命令帮助信息
# parser.parse_args(['a', '-h'])
# print('-' * 6)

print('解析参数')
args = parser.parse_args(['--arg1', 'arg1_value', '--arg2', 'arg2_value'])
print(f'args {args}')
print(f'args.arg1 {args.arg1} args.arg2 {args.arg2}')
print('-' * 6)

# 解析子命令参数
print('解析子命令参数')
args = parser.parse_args(['a', '--a_arg1', 'a_arg1_value'])
print(f'args {args}')
print(f'args.a_arg1 {args.a_arg1}')
print('-' * 6)

# 解析子命令参数
print('解析子命令参数')
args = parser.parse_args(['b', '--b_arg1', 'b_arg1_value'])
print(f'args {args}')
print(f'args.b_arg1 {args.b_arg1}')
print('-' * 20)
