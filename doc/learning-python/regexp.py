import re

text = "Hello, my email addresses are foo@example.com and bar@example.com."

print(f'text = {text}')
# 匹配开头是否符合, 不在内容中查找
print(f're.match("com", text) {re.match("com", text)}')  # 未匹配到
print(f're.match("he", text) {re.match("he", text)}')  # 未匹配到
print(f're.match("He", text) {re.match("He", text)}')  # 匹配成功
print('---------')

# 匹配整个字符串
# 匹配成功
print(f're.fullmatch("Hello", "Hello") {re.fullmatch("Hello", "Hello")}')
print('---------')

# 全局查找第一个符合条件的内容
match = re.search("com", text)
# <re.Match object; span=(42, 45), match='com'>
print(f're.search("com", text) {match}')
print(f'match.span() {match.span()}')  # (42, 45)
print(f'match.start() {match.start()}')  # 42
print(f'match.end() {match.end()}')  # 45
print(f'match.group() {match.group()}')  # com
print('---------')

# compile 编译正则表达式
reg = re.compile("com")
print(f'reg {reg} {type(reg)}')
# 全局查找所有符合条件的内容
print(f're.findall(reg, text) {re.findall(reg, text)}')  # ['com', 'com']
print('---------')

# 返回一个匹配结果的迭代器
iter = re.finditer("com", text)
print(f'iter {iter}')
for m in iter:
    print(f'for m in iter {m}')
print('---------')

# 分割字符串
print(f're.split("com", text) {re.split("com", text)}')
print('---------')

# 替换字符串, count 替换次数
result = re.sub("com", "xxx", text, count=1)
print(f're.sub("com", "xxx", text) {result}')
print('---------')

# 替换字符串, 返回元组包含替换后的字符串和替换的次数
result = re.subn("com", "xxx", text)
print(f're.subn("com", "xxx", text) {result}')
print('---------')
