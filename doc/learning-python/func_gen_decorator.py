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

# 生成器推导式
genobj = (pair for pair in zip(('a', 'b'), ('A', 'B')))
print(f'生成器推导式的类型 {type(genobj)}')  # <class 'generator'>
for i in genobj:
    print(f'生成器推导式生成的值 {i}')

print('----------------------------------')
print('----装饰器----')
'''
最接近 def 函数定义的装饰器最先被调用
'''


# 定义装饰器
def my_decorator(func):
    def new_func(*args, **kwargs):
        print('my_decorator before called.')
        print(f'Running func name: {func.__name__}')
        print(f'Positional arguments: {args}')
        print(f'Keyword arguments: {kwargs}')
        result = func(*args, **kwargs)
        print(f'Result: {result}')
        print('my_decorator after called.')
        return result
    return new_func


# 定义装饰器
def my_decorator_2(func):
    def new_func(name, *args, **kwargs):
        print('my_decorator_2 before called.')
        print(f'my_decorator_2 name: {name}')
        print(f'my_decorator_2 Running func name: {func.__name__}')
        print(f'my_decorator_2 Positional arguments: {args}')
        print(f'my_decorator_2 Keyword arguments: {kwargs}')
        result = func(*args, **kwargs)
        print(f'my_decorator_2 Result: {result}')
        print('my_decorator_2 after called.')
        return result * result
    return new_func


@my_decorator
@my_decorator_2
def add_ints(a, b):
    return a + b


print(f'add_ints(3, 5) result: {add_ints("zhangsan",3, 5)}')
