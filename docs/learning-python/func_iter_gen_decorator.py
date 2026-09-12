from collections.abc import Iterator
import functools
from typing import Any

print('迭代器: 自定义迭代器')


# 方式1:
# 可迭代对象: 实现 __iter__ 方法
class Company:
    def __init__(self, employee_list):
        self.employee = employee_list

    def __iter__(self):
        return MyIterator(self.employee)


# 迭代器: 实现 __iter__ 和 __next__ 方法
# 将记录 next 下一个值的索引保存在自定义迭代器对象中, 不要保存在可迭代对象中
class MyIterator(Iterator):
    def __init__(self, employee_list):
        self.employee = employee_list
        self.index = 0

    def __iter__(self):
        return self

    def __next__(self):
        try:
            emp = self.employee[self.index]
        except IndexError:
            raise StopIteration  # 抛出 StopIteration 异常
        self.index += 1
        return emp


for emp in Company(['tom', 'bob', 'jane']):
    print(emp)

print('-' * 10)

print('迭代器: 在类内部实现迭代器协议')


# 方式2: 在类内部实现迭代器协议
class Person:
    def __init__(self, name, age, gender, address):
        self.name = name
        self.age = age
        self.gender = gender
        self.address = address
        self.__index = 0  # 私有属性, 记录初始化状态
        self.__attrs = ['name', 'age', 'gender', 'address']

    def __iter__(self):
        self.__index = 0  # 每次迭代对象时, 重置状态
        return self

    def __next__(self):
        try:
            attr = self.__attrs[self.__index]
        except IndexError:
            raise StopIteration
        self.__index += 1
        return getattr(self, attr)


for emp in Person('tom', 18, 'male', 'beijing'):
    print(emp)

print('-' * 10)


# 自定义迭代器
class CountDown:
    """倒计时迭代器"""

    def __init__(self, start):
        self.current = start

    def __iter__(self):
        return self  # 返回迭代器对象

    def __next__(self):
        if self.current <= 0:
            raise StopIteration  # 结束迭代
        value = self.current
        self.current -= 1
        return value


# 使用自定义迭代器
counter = CountDown(5)
for num in counter:
    print(num, end=' ')  # 输出: 5 4 3 2 1
print()
print('-' * 30)


# 生成器函数，返回一个生成器对象
def my_gen(start=0, end=100, step=1):
    while start < end:
        yield start
        start += step


ranger = my_gen(1, 10, 2)

# 生成器是动态生成值, 迭代器只能遍历一次
for i in ranger:
    print(f'1 for {i}')

for i in ranger:
    print(f'2 for {i}')

print('-' * 10)


def my_gen2():
    inp = yield 'hehe'
    print(f'yield input {inp}')
    yield 250


gen2 = my_gen2()
print(next(gen2))
gen2.send(10)
print('-' * 10)


# 生成器生成质数
def prime_gen():
    i = 2
    yield i
    while True:
        i += 1
        for j in range(2, i):
            if i % j == 0:
                break
        else:  # 循环正常结束, 没有 break
            yield i


for i in prime_gen():
    print(f'质数: {i}')
    if i > 100:
        break
print('-' * 30)

# 生成器推导式
gen_obj = (pair for pair in zip(('a', 'b'), ('A', 'B')))
type_gen_obj: Any = type(gen_obj)
print(f'生成器推导式的类型 {type_gen_obj}')  # <class 'generator'>
for i in gen_obj:
    print(f'生成器推导式生成的值 {i}')

print('-' * 30)
print(f'{"-" * 6}装饰器{"-" * 6}')
'''
最接近 def 函数定义的装饰器最先被调用
'''


# 定义装饰器
def my_decorator(func):
    @functools.wraps(func)  # 保留被装饰函数的元信息
    def new_func(*args, **kwargs):
        print('my_decorator before called.')  # 1.
        print(f'Running func name: {func.__name__}')  # 2.
        print(f'Positional arguments: {args}')  # 3.
        print(f'Keyword arguments: {kwargs}')  # 4.
        result = func(*args, **kwargs)
        print(f'Result: {result}')  # 13.
        print('my_decorator after called.')  # 14.
        return result  # 15.

    return new_func


# 定义装饰器
def my_decorator_2(func):
    @functools.wraps(func)  # 保留被装饰函数的元信息
    def new_func(name, *args, **kwargs):
        print('my_decorator_2 before called.')  # 5.
        print(f'my_decorator_2 name: {name}')  # 6.
        print(f'my_decorator_2 Running func name: {func.__name__}')  # 7.
        print(f'my_decorator_2 Positional arguments: {args}')  # 8.
        print(f'my_decorator_2 Keyword arguments: {kwargs}')  # 9.
        result = func(*args, **kwargs)
        print(f'my_decorator_2 Result: {result}')  # 10.
        print('my_decorator_2 after called.')  # 11.
        return result * result  # 12.

    return new_func


@my_decorator
@my_decorator_2
def add_ints(a, b):
    return a + b


print(f'add_ints(3, 5) result: {add_ints("zhangsan", 3, 5)}')
print(
    f'@functools.wraps(func) 保留被装饰函数的元信息 add_ints.__name__ {add_ints.__name__}')


# 带参数的装饰器：外层吃参数，中层吃函数，内层吃调用参数
def repeat(n):
    """重复执行 n 次的装饰器"""

    def decorator(func):
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            result = None
            for _ in range(n):
                result = func(*args, **kwargs)
            return result  # 返回最后一次的结果

        return wrapper

    return decorator


@repeat(3)
def say_hello(name):
    print(f"Hello, {name}!")


say_hello("Python")
# Hello, Python!
# Hello, Python!
# Hello, Python!
