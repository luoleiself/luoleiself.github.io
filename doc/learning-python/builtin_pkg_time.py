from datetime import datetime as dt, timezone, timedelta, MAXYEAR, MINYEAR, UTC
import time
import calendar

'''
datetime 包的主要对象
    |- date: 表示日期的类
        |- datetime: 表示日期和时间的类
    |- time: 表示时间的类
    |- timedelta: 表示时间间隔的类
    |- tzinfo: 表示时区信息的抽象类
        |- timezone: 表示时区的类

zoneinfo 包的 ZoneInfo 类实现了 tzinfo 抽象类,
    不直接提供时区数据, 而是从系统时区数据库或 pypi 包 tzdata 包中提取时区信息

time.strftime() 格式化指令
%a  周工作日的缩写, time.strftime('%a')   'Thu'
%A  周工作的全名, time.strftime('%A')   'Thursday'
%b  月份的缩写, time.strftime('%b') 'Mar'
%B  月份的全名, time.strftime('%B')   'March'

%y  年份的后两位, time.strftime('%y')   '26'
%Y  年份的完整形式, time.strftime('%Y')   '2026'
%z  时区偏移量, time.strftime('%z') '+0800'
%Z  时区名称, time.strftime('%Z') '中国标准时间'
%m  月份,   time.strftime('%m')   '03'
%d  一个月中的第几天, time.strftime('%d') '26'
%H  二十四小时制的小时, time.strftime('%H')   '13'
%I  12小时制的小时, time.strftime('%I')   '01'
%M  分钟,   time.strftime('%M') '44'
%S  秒,   time.strftime('%S')   '08'

%c  本地化的日期和时间表示, time.strftime('%c')   'Thu Mar 26 13:39:26 2026'

%x  本地化的日期表示, time.strftime('%x')   '03/26/26'
%X  本地化的时间表示, time.strftime('%X')   '13:48:48'
%p  上午或下午, time.strftime('%p') 'PM'

%j  一年的第几天, time.strftime('%j')   '085'
%U  一年中第几周, 周日为一周的开始, time.strftime('%U')   '12'
%u  周几, 周一从 1 开始, time.strftime('%u')   '4'
%W  一年中第几周, 周一为一周的开始, time.strftime('%W')   '12'

%G  ISO8601 格式的年份, time.strftime('%G') '2026'
%V  ISO8601 格式的周数, time.strftime('%V') '13'
%%  %字面量,    time.strftime('%%')   '%'
'''

print('time:')
print(f'time.tzname: {time.tzname}')
print(f'time.timezone: {time.timezone}')

t3 = time.monotonic()
print(f'单调时钟秒 time.monotonic() {t3}')

t_monotoic_ns = time.monotonic_ns()
print(f'单调时钟纳秒 time.monotonic_ns() {t_monotoic_ns}')

t_ns = time.time_ns()
print(f'纪元时间整数纳秒: time.time_ns() {t_ns}')

t_thread_time = time.thread_time()
print(f'线程时钟秒 time.thread_time() {t_thread_time}')

t_thread_time_ns = time.thread_time_ns()
print(f'线程时钟纳秒 time.thread_time_ns() {t_thread_time_ns}')

t_process_time = time.process_time()
print(f'进程时钟秒 time.process_time() {t_process_time}')

t_process_time_ns = time.process_time_ns()
print(f'进程时钟纳秒 time.process_time_ns() {t_process_time_ns}')

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
delta = timedelta(days=1, hours=6, minutes=5, seconds=2,
                  milliseconds=4, microseconds=3)
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
            return f"{int(diff / 60)}分钟前"
        elif diff < 86400:
            return f"{int(diff / 3600)}小时前"
        elif diff < 2592000:
            return f"{int(diff / 86400)}天前"
        else:
            return time.strftime("%Y-%m-%d", time.localtime(timestamp))


# 使用示例
tf = TimeFormatter()

# 格式化秒数
print(tf.format_seconds(45))  # 45.0秒
print(tf.format_seconds(125))  # 2.1分钟
print(tf.format_seconds(7200))  # 2.0小时
print(tf.format_seconds(90000))  # 1.0天
print('---------')

# 解析时间字符串
print(tf.parse_time_string("1h30m"))  # 5400
print(tf.parse_time_string("45m"))  # 2700
print(tf.parse_time_string("30s"))  # 30
print('---------')

# 相对时间
past_time = time.time() - 3600
print(tf.get_relative_time(past_time))  # 1小时前
print('------------------------------------')

print('''
calendar: 日历包
    Calendar 日历基类
        TextCalendar 文本格式日历类
            LocalTextCalendar   本地文本格式日历类
        HTMLCalendar   html 格式日历类
            LocaleHTMLCalendar  本地 html 格式日历类
''')
# 一周完整格式
print(f'calendar.day_name {list(calendar.day_name)}')
# 一周的缩写格式
print(f'calendar.day_abbr {list(calendar.day_abbr)}')
# 月份的完成格式
print(f'calendar.month_name {list(calendar.month_name)}')
# 月份的缩写格式
print(f'calendar.month_abbr {list(calendar.month_abbr)}')
print('---------')
print(f'''calendar.MONDAY {calendar.MONDAY}
calendar.TUESDAY {calendar.TUESDAY}
calendar.WEDNESDAY {calendar.WEDNESDAY}
calendar.THURSDAY {calendar.THURSDAY}
calendar.FRIDAY {calendar.FRIDAY}
calendar.SATURDAY {calendar.SATURDAY}
calendar.SUNDAY {calendar.SUNDAY}''')
print('---------')
print(f'''calendar.JANUARY {calendar.JANUARY}
calendar.FEBRUARY {calendar.FEBRUARY}
calendar.MARCH {calendar.MARCH}
calendar.APRIL {calendar.APRIL}
calendar.MAY {calendar.MAY}
calendar.JUNE {calendar.JUNE}
calendar.JULY {calendar.JULY}
calendar.AUGUST {calendar.AUGUST}
calendar.SEPTEMBER {calendar.SEPTEMBER}
calendar.OCTOBER {calendar.OCTOBER}
calendar.NOVEMBER {calendar.NOVEMBER}
calendar.DECEMBER {calendar.DECEMBER}''')
print('---------')

# 日历基类
c = calendar.Calendar()
print(f'calandar.Calandar() {c}')  # 0
print(f'c.firstweekday {c.firstweekday}')
# 返回周迭代器 [0, 1, 2, 3, 4, 5, 6]
print(f'c.iterweekdays() {list(c.iterweekdays())}')
# 返回指定年月中的所有天数迭代器
# [0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 0, 0, 0]
print(f'c.itermonthdays(2026, 4) {list(c.itermonthdays(2026, 4))}')
# 返回指定年月的 datetime.date 日期对象组成的迭代器
# [datetime.date(2026, 3, 30), datetime.date(2026, 3, 31), datetime.date(2026, 4, 1), datetime.date(2026, 4, 2)...]
print(f'c.itermonthdates(2026, 4) {list(c.itermonthdates(2026, 4))}')
# 返回指定年月的所有天数和周几的元组的迭代器
# [(0, 0), (0, 1), (1, 2), (2, 3), (3, 4), (4, 5), (5, 6), (6, 0), (7, 1), (8, 2), (9, 3), (10, 4), (11, 5), (12, 6), (13, 0), (14, 1), (15, 2), (16, 3), (17, 4), (18, 5), (19, 6), (20, 0), (21, 1), (22, 2), (23, 3), (24, 4), (25, 5), (26, 6), (27, 0), (28, 1), (29, 2), (30, 3), (0, 4), (0, 5), (0, 6)]
print(f'c.itermonthdays2(2026, 4) {list(c.itermonthdays2(2026, 4))}')
# 返回指定年月的所有 (年, 月, 日)元组的迭代器
# [(2026, 3, 30), (2026, 3, 31), (2026, 4, 1), (2026, 4, 2), (2026, 4, 3), (2026, 4, 4), (2026, 4, 5), (2026, 4, 6), (2026, 4, 7), (2026, 4, 8), (2026, 4, 9), (2026, 4, 10), (2026, 4, 11), (2026, 4, 12), (2026, 4, 13), (2026, 4, 14), (2026, 4, 15), (2026, 4, 16), (2026, 4, 17), (2026, 4, 18), (2026, 4, 19), (2026, 4, 20), (2026, 4, 21), (2026, 4, 22), (2026, 4, 23), (2026, 4, 24), (2026, 4, 25), (2026, 4, 26), (2026, 4, 27), (2026, 4, 28), (2026, 4, 29), (2026, 4, 30), (2026, 5, 1), (2026, 5, 2), (2026, 5, 3)]
print(f'c.itermonthdays3(2026, 4) {list(c.itermonthdays3(2026, 4))}')
# 返回指定年月的所有 (年, 月, 日, 周)元组的迭代器
# [(2026, 3, 30, 0), (2026, 3, 31, 1), (2026, 4, 1, 2), (2026, 4, 2, 3), (2026, 4, 3, 4), (2026, 4, 4, 5), (2026, 4, 5, 6), (2026, 4, 6, 0), (2026, 4, 7, 1), (2026, 4, 8, 2), (2026, 4, 9, 3), (2026, 4, 10, 4), (2026, 4, 11, 5), (2026, 4, 12, 6), (2026, 4, 13, 0), (2026, 4, 14, 1), (2026, 4, 15, 2), (2026, 4, 16, 3), (2026, 4, 17, 4), (2026, 4, 18, 5), (2026, 4, 19, 6), (2026, 4, 20, 0), (2026, 4, 21, 1), (2026, 4, 22, 2), (2026, 4, 23, 3), (2026, 4, 24, 4), (2026, 4, 25, 5), (2026, 4, 26, 6), (2026, 4, 27, 0), (2026, 4, 28, 1), (2026, 4, 29, 2), (2026, 4, 30, 3), (2026, 5, 1, 4), (2026, 5, 2, 5), (2026, 5, 3, 6)]
print(f'c.itermonthdays4(2026, 4) {list(c.itermonthdays4(2026, 4))}')

print(f'c.monthdatescalendar(2026, 4) {c.monthdatescalendar(2026, 4)}')
print(f'c.monthdayscalendar(2026, 4) {c.monthdayscalendar(2026, 4)}')
print(f'c.monthdays2calendar(2026, 4) {c.monthdays2calendar(2026, 4)}')
print(f'c.yeardayscalendar(2026, 4)  {c.yeardayscalendar(2026, 4)}')
print(f'c.yeardatescalendar(2026, 4) {c.yeardatescalendar(2026, 4)}')
print('---------')

print('文本格式日历, 继承 Calendar')
txt_calendar = calendar.TextCalendar()
print(f'txt_calendar {txt_calendar}')
print(f'txt_calendar.getfirstweekday() {txt_calendar.getfirstweekday()}')
print(f'txt_calendar.iterweekdays() {list(txt_calendar.iterweekdays())}')
print(f'''
文本格式输出:
指定年日历
txt_calendar.formatyear(2026)
指定月份日历
txt_calendar.formatmonth(2026, 8)
{txt_calendar.formatmonth(2026, 8)}

# 右对齐日期
# txt_calendar.formatday(12, 6, 38)

# 一周的头部标识
# txt_calendar.formatweekheader(9)
{txt_calendar.formatweekheader(9)}

# 一周的日期和对应的周几
# txt_calendar.formatweek([(0,0), (0,1), (1,2), (2,3), (3, 4), (4, 5), (5, 6)], 2)
{txt_calendar.formatweekheader(2)}
{txt_calendar.formatweek([(0, 0), (0, 1), (1, 2), (2, 3), (3, 4), (4, 5), (5, 6)], 2)}
''')
print('---------')

print('本地文本格式日历, 继承 TextCalendar')
local_txt_calendar = calendar.LocaleTextCalendar()
print(f'local_txt_calendar {local_txt_calendar}')
print(f'''
本地文本格式输出:
指定月份日历
local_txt_calendar = calendar.LocaleTextCalendar()
local_txt_calendar.formatmonth(2026, 8)
{local_txt_calendar.formatmonth(2026, 8)}
''')
print('---------')

print('html 格式日历, 继承 Calendar')
html_calendar = calendar.HTMLCalendar()
print(f'html_calendar {html_calendar}')
print(f'html_calendar.getfirstweekday() {html_calendar.getfirstweekday()}')
print(f'html_calendar.cssclass_month {html_calendar.cssclass_month}')
print(f'html_calendar.monthdayscalendar(2026, 4) {html_calendar.monthdayscalendar(2026, 4)}')
print(f'''
html 格式输出:
指定年日历
# html_calendar.formatyear(2026)
指定月日历
html_calendar.formatmonth(2026, 8)
{html_calendar.formatmonth(2026, 8)}

一周日历头部标识
# html_calendar.formatweekheader()
{html_calendar.formatweekheader()}

一天
# html_calendar.formatday(12, 6)
{html_calendar.formatday(12, 6)}
''')
print('---------')

print('本地 html 格式日历, 继承 HTMLCalendar')
local_html_calendar = calendar.LocaleHTMLCalendar()
print(f'local_html_calendar {local_html_calendar}')
print(f'''
本地 html 格式输出:
指定月份日历
local_html_calendar.formatmonth(2026, 8)
{local_html_calendar.formatmonth(2026, 8)}
''')
