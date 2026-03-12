import functools
import itertools
import operator

# operator 模块
print('operator 模块导出一组与 python 内置运算符相对应的高效函数')
# 10 + 1 == operator.add(10, 1)
print(f'operator.add(1, 10) = {operator.add(1, 10)}')  # 11
print(f'operator.sub(10, 1) = {operator.sub(10, 1)}')   # 9
print(f'operator.lt(1, 10) = {operator.lt(1, 10)}')   # True

print(f'operator.is_not(1, 2) = {operator.is_not(1, 2)}')   # True
# True
print(f'operator.contains([1, 2], 1) = {operator.contains([1, 2], 1)}')

print(f'operator.lshift(5, 1) = {operator.lshift(5, 1)}')  # 10
print(f'operator.rshift(5, 1) = {operator.rshift(5, 1)}')   # 2

# [1, 2, 3, 'a', 'b']
print(
    f'operator.concat([1, 2, 3], ["a", "b"]) = {operator.concat([1, 2, 3], ["a", "b"])}')

print(f'operator.and(True, False) = {operator.and_(True, False)}')  # False
print(f'operator.or_(True, False) = {operator.or_(True, False)}')   # True
print('-----------------------------------------')

# functools 模块用于处理高阶函数, 作用于其他函数或返回其他函数的函数
# 123
print(
    f'累积 返回一个结果: functools.reduce(lambda x, y: x * y, [1, 2, 3, 4, 5]) = {functools.reduce(lambda x, y: x * y, [1, 2, 3, 4, 5])}')
print('-----------------------------------------')

# itertools 模块
# [1, 3, 6, 10, 15]
print(
    f'累加: itertools.accumulate([1, 2, 3, 4, 5]) = {list(itertools.accumulate([1, 2, 3, 4, 5]))}')
# [1, 2, 6, 24, 120]
print(
    f'累积: itertools.accumulate([1, 2, 3, 4, 5], operator.mul) = {list(itertools.accumulate([1, 2, 3, 4, 5], operator.mul))}')
# [1, 2, 3, [4], 5, [6]]
print(
    f'解构: itertools.chain([1, 2, 3], [[4], 5, [6]]) = {list(itertools.chain([1, 2, 3], [[4], 5, [6]]))}')
# ['a', 'b', 'c', 'A', 'B', 'C', 'D']
print(
    f'解构: itertools.chain("abc", "ABCD") = {list(itertools.chain("abc", "ABCD"))}')
dt = {'a': 'A', 'b': 'B', 'c': 'C'}
# ['a', 'b', 'c']
print(fr'解构字典的键: itertools.chain(dt) = {list(itertools.chain(dt))}')
# [1, 1, 1]
print(
    f'重复 返回元素重复次数的列表: itertools.repeat(1, 3) = {list(itertools.repeat(1, 3))}')
# [[1, 2, 3], [1, 2, 3], [1, 2, 3], [1, 2, 3]]
print(
    f'重复 返回元素重复次数的列表: itertools.repeat([1, 2, 3], 4) = {list(itertools.repeat([1, 2, 3], 4))}')
# (<itertools._tee object at 0x0000021F3B9F4140>, <itertools._tee object at 0x0000021F3B9F4780>, <itertools._tee object at 0x0000021F3B9F4240>)
print(
    f'重复 返回指定数量的迭代器对象元组, itertools.tee([1, 2, 3, 4, 5], 3) = {tuple(itertools.tee([1, 2, 3, 4, 5], 3))}')
te = itertools.tee('ABC', 2)
for t in te:
    print(list(t))
# ['A', 'B', 'C']
# ['A', 'B', 'C']
print('----------')
# [(1, 2), (3, 4), (5,)]
print(
    f'分组, 指定分组元素个数: itertools.batched([1, 2, 3, 4, 5], 2) = {list(itertools.batched([1, 2, 3, 4, 5], 2))}')

print(f'itertools.groupby([], len) 按元素的长度分组')
gb = itertools.groupby(
    ['A', 'B', 'C', 'D', 'EFG', 'H', 'IJ', 'KL', 'MN', 'OP'], len)
for k, v in gb:
    print(k, list(v))
# 1 ['A', 'B', 'C', 'D']
# 3 ['EFG']
# 1 ['H']
# 2 ['IJ', 'KL', 'MN', 'OP']
print('----------')
print(f'compress, filterfalse, dropwhile, takewhile')
# [1, 3]
print(
    f'选择 按元素位置选择: itertools.compress([1, 2, 3, 4, 5], [1, 0, 1, 0]) = {list(itertools.compress([1, 2, 3, 4, 5], [1, 0, 1, 0]))}')
# [6, 8]
print(
    f'过滤 所有符合条件的元素: itertools.filterfalse(lambda x: x < 5, [1, 4, 6, 3, 8]) = {list(itertools.filterfalse(lambda x: x < 5, [1, 4, 6, 3, 8]))}')
# [6, 3, 8]
print(
    f'过滤 丢弃符合条件的元素直到 False 结束: itertools.dropwhile(lambda x: x < 5, [1, 4, 6, 3, 8]) = {list(itertools.dropwhile(lambda x: x < 5, [1, 4, 6, 3, 8]))}')
# [1, 4]
print(
    f'过滤 保留符合条件的元素直到 False 结束: itertools.takewhile(lambda x: x < 5, [1, 4, 6, 3, 8]) = {list(itertools.takewhile(lambda x: x < 5, [1, 4, 6, 3, 8]))}')
print('----------')
# [2, 4]
print(
    f'切片: itertools.isslice([1, 2, 3, 4, 5], 1, 5, 2) = {list(itertools.islice([1, 2, 3, 4, 5], 1, 5, 2))}')
# ['c', 'f', 'i', 'l']
print(
    f'切片: itertools.isslice("abcdefghijkl", 2, None, 3) = {list(itertools.islice("abcdefghijkl", 2, None, 3))}')
# [(1, 2), (2, 3), (3, 4), (4, 5)]
print(
    f'组对: 相邻每两个元素组成一个元组, 返回双项元组列表: itertools.pairwise([1, 2, 3, 4, 5]) = {list(itertools.pairwise([1, 2, 3, 4, 5]))}')
print('----------')
# [2, 12, 30]
print(
    f'遍历每个双项元组列表, 对每个元组应用函数: itertools.starmap(mul, [(1, 2), (3, 4), (5, 6)]) = {list(itertools.starmap(operator.mul, [(1, 2), (3, 4), (5, 6)]))}')
print('----------')
# [(1, 4), (2, 5), (0, 6), (0, 7)]
print(
    f'zip 遍历按最长序列对齐, 返回元组列表, 填充缺失值(默认为 None): itertools.zip_longest((1, 2), [4, 5, 6, 7], fillvalue=0) = {list(itertools.zip_longest((1, 2), [4, 5, 6, 7], fillvalue=0))}')
print('----------')
print(
    f'生成多个序列的笛卡尔积元组列表 itertools.product([1, 2, 3], [4, 5, 6, 7], ("A", "B")) = {list(itertools.product([1, 2, 3], [4, 5, 6, 7], ("A", "B")))}')
#  [('A', 'A'), ('A', 'B'), ('A', 'C'), ('A', 'D'), ('B', 'A'), ('B', 'B'), ('B', 'C'), ('B', 'D'), ('C', 'A'), ('C', 'B'), ('C', 'C'), ('C', 'D'), ('D', 'A'), ('D', 'B'), ('D', 'C'), ('D', 'D')]
print(
    f'生成多个序列的笛卡尔积元组列表 itertools.product("ABCD", repeat=2) = {list(itertools.product("ABCD", repeat=2))}')
# [('A', 'B'), ('A', 'C'), ('A', 'D'), ('B', 'A'), ('B', 'C'), ('B', 'D'), ('C', 'A'), ('C', 'B'), ('C', 'D'), ('D', 'A'), ('D', 'B'), ('D', 'C')]
print(
    f'生成多个序列的笛卡尔积元组列表 排除自身笛卡尔积: itertools.permutations("ABCD", 2) = {list(itertools.permutations("ABCD", 2))}')
# [('A', 'B'), ('A', 'C'), ('A', 'D'), ('B', 'C'), ('B', 'D'), ('C', 'D')]
print(
    f'生成多个序列的笛卡尔积元组列表 排除自身和交换位置笛卡尔积: itertools.combinations("ABCD", 2) = {list(itertools.combinations("ABCD", 2))}')
# [('A', 'A'), ('A', 'B'), ('A', 'C'), ('A', 'D'), ('B', 'B'), ('B', 'C'), ('B', 'D'), ('C', 'C'), ('C', 'D'), ('D', 'D')]
print(
    f'生成多个序列的笛卡尔积元组列表 排除交换位置但保留自身笛卡尔积: itertools.combinations_with_replacement("ABCD", 2) = {list(itertools.combinations_with_replacement("ABCD", 2))}')
print('-----------------------------------------')
