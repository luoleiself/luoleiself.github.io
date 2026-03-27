import csv
import pathlib
from urllib import parse
import shutil
import glob
import random
import stat
import sys
import os
import logging
import time
import string

print('string 模块')
print(f'string.ascii_letters {string.ascii_letters}')
print(f'string.ascii_lowercase {string.ascii_lowercase}')
print(f'string.ascii_uppercase {string.ascii_uppercase}')
print(f'string.digits {string.digits}')
print(f'string.octdigits {string.octdigits}')
print(f'string.hexdigits {string.hexdigits}')
print(f'string.punctuation {string.punctuation}')
print(f'string.printable {string.printable}')
print(f'string.whitespace {string.whitespace}')
print('---------------')

print('sys 模块')
print(f'sys.version {sys.version}')
print(f'sys.platform {sys.platform}')
print(f'sys.copyright {sys.copyright}')
print(f'sys.argv {sys.argv}')
print(f'sys.path {sys.path}')
print(f'sys.api_version {sys.api_version}')
print(f'sys.getdefaultencoding() {sys.getdefaultencoding()}')
print(f'sys.getprofile() {sys.getprofile()}')
print(f'sys.intern("hello") {sys.intern("hello")}')
# print(f'sys.__dict__ {sys.__dict__}') # 内部字典
print('---------------')

print('os 模块')
uname_res = os.uname_result
print(f'uname_res {uname_res.sysname} {uname_res.release} {uname_res.version} {uname_res.machine} {uname_res.nodename}')
print(f'os.uname_result {os.uname_result}')
print(f'os.getlogin() {os.getlogin()}')
os.system('git status')  # 在 shell 中执行命令
print(f'os.getcwd() {os.getcwd()}')
print(f'os.name {os.name}')
print(f'os.getpid() {os.getpid()}')
print(f'os.sep {os.sep}')
print(f'os.cpu_count() {os.cpu_count()}')
print(f'os.urandom(16) {os.urandom(16)}')
s = 'hello world 中国'
bs = s.encode()
print(f's = {s} {len(s)}')
print(f'encode: {bs} {len(bs)}, decode: {bs.decode()}')
print(f'os.getenv("PATH") {os.getenv("PATH")}')
print(f'os.getlogin() {os.getlogin()}')
# os.kill(os.getpid(), 9)   # 关闭当前进程
# os.putenv(name, value)    # 添加环境变量
print('---------------')

print(f'os.path {os.path}')
print(
    f'os.path.curdir {os.path.curdir} os.path.pardir {os.path.pardir} os.path.extsep {os.path.extsep}')
print(f'os.path.basename(__file__) {os.path.basename(__file__)}')
print(f'os.path.dirname(__file__) {os.path.dirname(__file__)}')
print(f'os.path.abspath(__file__) {os.path.abspath(__file__)}')
print(f'os.path.realpath(__file__) {os.path.realpath(__file__)}')
print(f'os.path.exists(__file__) {os.path.exists(__file__)}')
print(f'os.path.isfile(__file__) {os.path.isfile(__file__)}')
print(f'os.path.isabs("test.txt") {os.path.isabs("test.txt")}')
print(f'os.path.isdir(__file__) {os.path.isdir(__file__)}')
print(f'os.path.islink(__file__) {os.path.islink(__file__)}')
print(f'os.path.ismount(__file__) {os.path.ismount(__file__)}')
print('---------------')

print('pathlib 模块: 基于对象的文件路径')
p = pathlib.Path(__file__)
print(f'p {p} {p.name} {p.stem} {p.suffix} {p.parent} {p.absolute()}')
print(f'p.exists() {p.exists()}')
print(f'p.home() {p.home()}')
print(f'p.is_file() {p.is_file()}')
print(f'p.is_dir() {p.is_dir()}')
print('---------------')

print('stat 模块')
# print(f'stat.__dict__ {stat.__dict__}')
print(f'stat.filemode(0o100755) {stat.filemode(0o100755)}')
f_stat = os.stat("builtin_pkg.py")
print(f'os.stat("builtin_pkg.py") {f_stat}')
print(f'stat.filemode(f_stat.st_mode) {stat.filemode(f_stat.st_mode)}')
print(f'stat.S_ISDIR(f_stat.st_mode) {stat.S_ISDIR(f_stat.st_mode)}')
print(f'stat.S_ISCHR(f_stat.st_mode) {stat.S_ISCHR(f_stat.st_mode)}')
print(f'stat.S_ISBLK(f_stat.st_mode) {stat.S_ISBLK(f_stat.st_mode)}')
print(f'stat.S_ISPORT(f_stat.st_mode) {stat.S_ISPORT(f_stat.st_mode)}')

print('权限位:')
s_perm = (stat.S_ISUID, stat.S_ISGID, stat.S_ISVTX)
u_perm = (stat.S_IRWXU, stat.S_IRUSR, stat.S_IWUSR, stat.S_IXUSR)
g_perm = (stat.S_IRWXG, stat.S_IRGRP, stat.S_IWGRP, stat.S_IXGRP)
o_perm = (stat.S_IRWXO, stat.S_IROTH, stat.S_IWOTH, stat.S_IXOTH)

s_perm_keys = ('S_ISUID', 'S_ISGID', 'S_ISVTX')
u_perm_keys = ('S_IRWXU', 'S_IRUSR', 'S_IWUSR', 'S_IXUSR')
g_perm_keys = ('S_IRWXG', 'S_IRGRP', 'S_IWGRP', 'S_IXGRP')
o_perm_keys = ('S_IRWXO', 'S_IROTH', 'S_IWOTH', 'S_IXOTH')

s_perm_oct_int = [(f'{x:#06o}', x) for x in s_perm]
u_perm_oct_int = [(f'{x:#06o}', x) for x in u_perm]
g_perm_oct_int = [(f'{x:#06o}', x) for x in g_perm]
o_perm_oct_int = [(f'{x:#06o}', x) for x in o_perm]

s_perm_dict = {k: v for k, v in zip(s_perm_keys, s_perm_oct_int)}
print(f'特殊权限位: {s_perm_dict}')
u_perm_dict = {k: v for k, v in zip(u_perm_keys, u_perm_oct_int)}
print(f'所有者权限位: {u_perm_dict}')
g_perm_dict = {k: v for k, v in zip(g_perm_keys, g_perm_oct_int)}
print(f'组权限位: {g_perm_dict}')
o_perm_dict = {k: v for k, v in zip(o_perm_keys, o_perm_oct_int)}
print(f'其他用户权限位: {o_perm_dict}')
print('----------------------------------------------')

print('glob 模块: 按模式匹配文件')
print(f'glob.glob("*.py") {glob.glob("*.py")}')
print('----------------------------------------------')

print('shutil 模块: 高级文件操作')
# shutil.make_archive('test', 'zip', '.')  # 创建归档文件
# shutil.copy(__file__, __file__ + '.copy') # 拷贝文件
# shutil.move(__file__, __file__ + '.move') # 移动文件
# print(
#     f'shutil.copy(__file__, __file__ + \'.copy\') {shutil.copy(__file__, __file__ + ".copy")}')
print('----------------------------------------------')

print('csv 模块: 读写 csv 文件')
print(f'csv.reader(f) 创建一个 csv 读取迭代器')
print(f'csv.writer(f) 创建一个 csv 写入对象')
print(f'    csv_out.writerow(row) 写入一行')
print(f'    csv_out.writerows(rows) 写入多行')
print(f'csv.list_dialects() 列出所有方言 {csv.list_dialects()}')
print(f'csv.field_size_limit() 输出字段大小限制 {csv.field_size_limit()}')  # 输出字段大小限制
print(f'csv.QUOTE_ALL 输出字段所有内容 {csv.QUOTE_ALL}')  # 输出字段所有内容
print(f'csv.QUOTE_MINIMAL 输出字段最小内容 {csv.QUOTE_MINIMAL}')  # 输出字段最小内容
print('---------------')
print('读写多项序列的列表')
csv_data = [['a', 'b', 'c'], (1, 2, 3), (4, 5, 6)]
print(f'csv_data = {csv_data}')
with open('test_1.csv', 'wt', newline='') as csvfile:
    csv_out = csv.writer(csvfile, strict=True)
    csv_out.writerow(['col 1', 'col 2', 'col 3'])  # 写入一行
    csv_out.writerows(csv_data)  # 写入多行

with open('test_1.csv', 'rt', newline='') as csvfile:
    csv_input = csv.reader(csvfile, strict=True)
    # 迭代读取结果
    for row in csv_input:
        print(f'row = {row}')
print('---------------')
print('读写字典列表')
print(f'csv.DictReader(f) 创建一个 csv 字典读取迭代器')
print(f'    csv_input.fieldnames 输出列名')
print(f'csv.DictWriter(f, fieldnames) 创建一个 csv 字典写入对象')
print(f'    csv_out.writeheader() 写入列名 fieldnames')
print(f'    csv_out.writerow(dict) 写入一行')
print(f'    csv_out.writerows(dict_list) 写入多行')
with open('test_2.csv', 'wt', newline='') as csvfile:
    # fieldnames 指定列名
    fieldnames = ['col 1', 'col 2', 'col 3']
    csv_out = csv.DictWriter(csvfile, fieldnames=fieldnames, strict=True)
    csv_out.writeheader()   # 写入列名
    csv_out.writerow({'col 1': 'a', 'col 2': 'b', 'col 3': 'c'})  # 写入一行
    csv_out.writerows([
        {'col 1': 1, 'col 2': 2, 'col 3': 3},
        {'col 1': 4, 'col 2': 5, 'col 3': 6},
    ])  # 写入多行
with open('test_2.csv', 'rt', newline='') as csvfile:
    csv_input = csv.DictReader(csvfile, strict=True)
    print(f'csv_input.fieldnames {csv_input.fieldnames}')
    for row in csv_input:
        print(
            f'row["col 1"] {row["col 1"]}, row["col 2"] {row["col 2"]}, row["col 3"] {row["col 3"]}')

print('----------------------------------------------')

print('logging 模块')
# 日志格式；asctime 默认用 localtime，通过 Formatter.converter=time.gmtime 改为 GMT（与 UTC 等效）
fmt = '%(asctime)s %(levelname)s %(filename)s Line: %(lineno)d %(message)s'

# 创建日志处理程序
_log_handler = logging.StreamHandler()
# 自定义日期时间格式化
_log_formatter = logging.Formatter(fmt=fmt, datefmt='%Y-%m-%dT%H:%M:%S.000Z')
# 改为 GMT（与 UTC 等效）
_log_formatter.converter = time.gmtime
_log_handler.setFormatter(_log_formatter)

logging.basicConfig(level=logging.NOTSET, handlers=[_log_handler])

logger = logging.getLogger(__name__)

logger.info('This is a info message')
logger.debug('This is a debug message')
logger.warning('This is a warning message')
logger.error('This is a error message')
logger.fatal('This is a fatal message')
print('----------------------------------------------')

print('random 模块')
print(f'random.random() {random.random():.4f}')
print(f'random.gauss(0, 1) {random.gauss(0, 1):.4f}')
print(f'random.choice([1, 2, 3, 4, 5]) {random.choice([1, 2, 3, 4, 5])}')
print(f'random.randint(1, 10) {random.randint(1, 10)}')
print(f'random.randrange(3, 9, 2) {random.randrange(3, 9, 2)}')
print(f'random.sample(range(100_000), 30) {random.sample(range(100_000), 30)}')
print(f'random.randbytes(10) {random.randbytes(10)}')
print(
    f'随机返回序列中任意两个元素: random.sample([1, 2, 3, 4, 5], 2) {random.sample([1, 2, 3, 4, 5], 2)}')
print('----------------------------------------------')

print('urllib 包')
print('编码和解码')
encode_param = parse.quote("中国")
print(f'parse.quote(中国) {encode_param}')
print(f'parse.unquote({encode_param}) {parse.unquote(encode_param)}')
print('------------')

url = 'https://www.bilibili.com/video/BV1rpWjevEip?spm_id_from=333.788.player.switch&vd_source=f70e6b2f23318ff77f33a2e2ba2f1224&p=89'
# unwrap_url = parse.unwrap('https://www.bilibili.com/video/BV1rpWjevEip?spm_id_from=333.788.player.switch&vd_source=f70e6b2f23318ff77f33a2e2ba2f1224&p=89')
# print(f'parse.unwrap(https://www.bilibili.com/video/BV1rpWjevEip?spm_id_from=333.788.player.switch&vd_source=f70e6b2f23318ff77f33a2e2ba2f1224&p=89) {unwrap_url}')

encode_url = parse.urlencode({'a': ('A', 'AA', 'AAA'), 'b': 'B', 'c': 'D'})
print(
    f'parse.urlencode({{"a": ("A", "AA", "AAA"), "b": "B", "c": "D"}}) {encode_url}')

print('------------')
parsed_url = parse.urlparse(url)
print(f'解析 url 返回 ParseResult: parse.urlparse(url) {parsed_url}')
parsed_query_dict = parse.parse_qs(parsed_url.query)
print(f'parse.parse_qs(query) {parsed_query_dict}')
parsed_query_list = parse.parse_qsl(parsed_url.query)
print(f'parse.parse_qsl(query) {parsed_query_list}')
print(f'回解析: parse.urlunparse(parsed_url) {parse.urlunparse(parsed_url)}')
print('------------')

splited_url = parse.urlsplit(url)
print(f'解析 url 返回 SplitResult: parse.urlsplit(url) {splited_url}')
print(f'splited_url.query {splited_url.query}')
print(f'回解析: parse.urlunsplit(splited_url) {parse.urlunsplit(splited_url)}')
print('------------------------------------')
