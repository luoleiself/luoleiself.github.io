print('------位运算符------')
x = 5
y = 1
print(f'x = {x} {x:#06b}, y = {y} {y:#06b}')
print(f'x & y = {(x & y)} {(x & y):#6b}')
print(f'x | y = {x | y} {(x | y):#6b}')
print(f'~x = {~x} {(~x):#6b}')
print(f'x ^ y = {x ^ y} {(x ^ y):#6b}')
print(f'x << 1 = {x << 1} {(x << 1):#6b}')
print(f'x >> 1 = {x >> 1} {(x >> 1):#6b}')

print('------转换为字符串------')
print(f'int 转换为字符串: {str(123)}')
print(f'float 转换为字符串: {str(123.456)}')
print(f'list 转换为字符串: {str([1, 2, 3, 4, 'a', 'b', 'c', 'd'])}')
print(f'tuple 转换为字符串: {str((1, 2, 3, 4, 'a', 'b', 'c', 'd'))}')
print(f'set 转换为字符串: {str({1, 2, 3, 4, 'a', 'b', 'c', 'd'})}')
print(f'dict 转换为字符串: {str({"a": 'A', "b": 'B', "c": 'C', "d": 'D'})}')

print('------转换为列表------')
print(f'str 转换为列表: {list("1234abcde")}')
print(f'tuple 转换为列表: {list((1, 2, 3, 4, 'a', 'b', 'c', 'd'))}')
print(f'set 转换为列表: {list({1, 2, 3, 4, 'a', 'b', 'c', 'd'})}')
# dict 转换为列表时, 只转换为键的列表
print(f'dict 转换为列表: {list({"a": 'A', "b": 'B', "c": 'C', "d": 'D'})}')

print('------转换为元组------')
print(f'str 转换为元组: {tuple("1234abcde")}')
print(f'list 转换为元组: {tuple([1, 2, 3, 4, 'a', 'b', 'c', 'd'])}')
print(f'set 转换为元组: {tuple({1, 2, 3, 4, 'a', 'b', 'c', 'd'})}')
# dict 转换为元组时, 只转换为键的元组
print(f'dict 转换为元组: {tuple({"a": 'A', "b": 'B', "c": 'C', "d": 'D'})}')

print('------转换为集合------')
print(f'str 转换为集合: {set("1234abcde")}')
print(f'list 转换为集合: {set([1, 2, 3, 4, 'a', 'b', 'c', 'd'])}')
print(f'tuple 转换为集合: {set((1, 2, 3, 4, 'a', 'b', 'c', 'd'))}')
# dict 转换为集合时, 只转换为键的集合
print(f'dict 转换为集合: {set({"a": 'A', "b": 'B', "c": 'C', "d": 'D'})}')

print('------转换为字典------')
print('列表, 元组, 集合 转换为字典时, 必须符合包含双项序列的任意序列')
print(f'list 转换为字典: {dict([("a", 1), ("b", 2), ("c", 3), ("d", 4)])}')
print(f'tuple 转换为字典: {dict((("a", 1), ("b", 2), ("c", 3), ("d", 4)))}')
print(f'set 转换为字典: {dict({("a", 1), ("b", 2), ("c", 3), ("d", 4)})}')
