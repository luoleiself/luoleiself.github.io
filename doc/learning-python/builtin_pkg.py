import shutil
from datetime import datetime as dt, timezone, timedelta, MAXYEAR, MINYEAR, UTC
import glob
import random
import stat
import sys
import time
import os
import logging

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

print('logging 模块')
# logging.basicConfig(level=logging.NOTSET) # 日志级别
logging.info('This is a info message')
logging.debug('This is a debug message')
logging.warning('This is a warning message')
logging.error('This is a error message')
logging.fatal('This is a fatal message')
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

'''
datetime 包的主要对象
    |- date: 表示日期的类
        |- datetime: 表示日期和时间的类
    |- time: 表示时间的类
    |- timedelta: 表示时间间隔的类
    |- tzinfo: 表示时区信息的抽象类
        |- timezone: 表示时区的类
'''

print('time:')
print(f'time.tzname: {time.tzname}')
print(f'time.timezone: {time.timezone}')
print(f'time.time_ns(): {time.time_ns()}')

t3 = time.monotonic()
print(f'秒 time.monotonic() {t3}')

# 1. 当前时间秒
ts = time.time()
print(f"1. 秒 time.time(): {ts}")

# 2 秒 -> 字符串
c_time = time.ctime(ts)
print(f"2. 秒 -> 字符串 time.ctime(ts): {c_time}")

# 2. 秒 → 结构化时间
local_time = time.localtime(ts)
print(
    f"2. 秒 -> local 结构化: time.localtime(ts): {local_time} 时区 {local_time.tm_zone}")
utc_time = time.gmtime(ts)
print(f"2. 秒 -> UTC 结构化: time.gmtime(ts): {utc_time} 时区 {utc_time.tm_zone}")

# 3. 结构化时间 → 字符串
asc_time = time.asctime(local_time)
print(f"3. 结构化 -> 字符串 time.asctime(local_time): {asc_time}")
local_str_ftime = time.strftime("%Y-%m-%d %H:%M:%S", local_time)
print(
    f"3. 结构化 -> 自定义格式字符串 time.strftime('%Y-%m-%d %H:%M:%S', local_time): {local_str_ftime}")

# 3.1 结构化时间 -> UTC 字符串
utc_str_time = time.strftime('%A %B %Y-%m-%dT%H:%M:%S.000Z', utc_time)
print(
    f"3.1 结构化 -> UTC 字符串 time.strftime('%A %B %Y-%m-%dT%H:%M:%S.000Z', utc_time): {utc_str_time}")
utc_str_ptime = time.strptime(utc_str_time, "%A %B %Y-%m-%dT%H:%M:%S.000Z")
print(
    f'3.1 UTC 字符串 -> 结构化 time.strptime(utc_str_time, "%A %B %Y-%m-%dT%H:%M:%S.000Z") {utc_str_ptime} {time.mktime(utc_str_ptime)}')

# 4. 字符串 → 结构化时间
str_ptime = time.strptime(local_str_ftime, "%Y-%m-%d %H:%M:%S")
print(f"4. 回结构化 time.strptime(str_time, '%Y-%m-%d %H:%M:%S'): {str_ptime}")

# 5. 结构化时间 → 时间戳
ts2 = time.mktime(str_ptime)
print(f"5. 转回时间戳 time.mktime(str_ptime): {ts2}")

# 验证一致性
print(f"时间戳一致: {abs(ts - ts2) < 0.001}")
print('---------')

# datetime 模块
print('datetime 包')
# timezone 类
print(f'timezone.max {timezone.max} timezone.min {timezone.min}')
print(f'timezone.tzname {timezone.tzname}')
print(f'timezone.utc {timezone.utc} timezone.type {type(timezone.utc)}')
print('----------')
d = dt.now()
print(f'dt.now() {d.ctime()} {d.strftime("%Y-%m-%d %H:%M:%S")}')
print(
    f'd.isoformat() {d.isoformat()} d.astimezone() {d.astimezone(timezone.utc)} d.utcoffset() {d.utcoffset()}')
print('----------')
print(
    f'datetime.MAXYEAR {MAXYEAR} datetime.MINYEAR {MINYEAR} datetime.UTC {UTC}')
utc_timestamp = dt.now(timezone.utc).timestamp()
print(f'utc_timestamp: {utc_timestamp}')
iso_string = dt.now(timezone.utc).isoformat()
print(f'iso_string: {iso_string}')
local_from_utc_timestamp = dt.fromtimestamp(
    utc_timestamp, timezone.utc).astimezone()
print(f'local_from_utc_timestamp: {local_from_utc_timestamp}')
local_from_iso_string = dt.fromisoformat(iso_string).astimezone()
print(f'local_from_iso_string: {local_from_iso_string}')
print('----------')
print('timedelta 模块')
delta = timedelta(days=1, seconds=2, microseconds=3,
                  milliseconds=4, minutes=5, hours=6, weeks=7)
print(f'delta {delta} delta.max {delta.max} delta.min {delta.min}')
print(
    f'delta.days {delta.days} delta.seconds {delta.seconds} delta.microseconds {delta.microseconds}')
print(f'delta.total_seconds() {delta.total_seconds()}')
delta += timedelta(days=1, seconds=2200,
                   microseconds=3000, minutes=50, hours=10)
print(f'delta.__add__(timedelta(days=1)) {delta}')
print('----------------------------------------------')


class TimeFormatter:
    """时间格式化工具"""

    @staticmethod
    def format_seconds(seconds):
        """将秒数格式化为可读字符串"""
        if seconds < 60:
            return f"{seconds:.1f}秒"
        elif seconds < 3600:
            minutes = seconds / 60
            return f"{minutes:.1f}分钟"
        elif seconds < 86400:
            hours = seconds / 3600
            return f"{hours:.1f}小时"
        else:
            days = seconds / 86400
            return f"{days:.1f}天"

    @staticmethod
    def parse_time_string(time_str):
        """解析时间字符串"""
        # 支持格式: "1h30m", "2h", "45m", "30s"
        import re

        total_seconds = 0

        # 匹配小时
        hour_match = re.search(r'(\d+)h', time_str)
        if hour_match:
            total_seconds += int(hour_match.group(1)) * 3600

        # 匹配分钟
        minute_match = re.search(r'(\d+)m', time_str)
        if minute_match:
            total_seconds += int(minute_match.group(1)) * 60

        # 匹配秒
        second_match = re.search(r'(\d+)s', time_str)
        if second_match:
            total_seconds += int(second_match.group(1))

        return total_seconds if total_seconds > 0 else None

    @staticmethod
    def get_relative_time(timestamp):
        """获取相对时间描述"""
        now = time.time()
        diff = now - timestamp

        if diff < 0:
            return "未来时间"
        elif diff < 60:
            return f"{int(diff)}秒前"
        elif diff < 3600:
            return f"{int(diff/60)}分钟前"
        elif diff < 86400:
            return f"{int(diff/3600)}小时前"
        elif diff < 2592000:
            return f"{int(diff/86400)}天前"
        else:
            return time.strftime("%Y-%m-%d", time.localtime(timestamp))


# 使用示例
tf = TimeFormatter()

# 格式化秒数
print(tf.format_seconds(45))        # 45.0秒
print(tf.format_seconds(125))       # 2.1分钟
print(tf.format_seconds(7200))      # 2.0小时
print(tf.format_seconds(90000))     # 1.0天
print('---------')

# 解析时间字符串
print(tf.parse_time_string("1h30m"))  # 5400
print(tf.parse_time_string("45m"))    # 2700
print(tf.parse_time_string("30s"))    # 30
print('---------')

# 相对时间
past_time = time.time() - 3600
print(tf.get_relative_time(past_time))  # 1小时前
print('------------------------------------')
