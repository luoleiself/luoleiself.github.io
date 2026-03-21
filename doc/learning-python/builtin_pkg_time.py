from datetime import datetime as dt, timezone, timedelta, MAXYEAR, MINYEAR, UTC
import time

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
