from dataclasses import dataclass


# 装饰器 @property, 只读属性
class Circle:
    def __init__(self, radius):
        self.radius = radius

    @property
    def diameter(self):
        return 2 * self.radius


c = Circle(5)
print(f'c.radius {c.radius} c.diameter {c.diameter}')
try:
    # AttributeError: property 'diameter' of 'Circle' object has no setter
    c.diameter = 10   # 只读属性
except AttributeError as e:
    print(f'只能属性不能修改: {e}')
print('---------------------------')


# 装饰器 @property, @setter, 可读写属性
class Duck:
    def __init__(self, name):
        self.__name = name  # 内置属性，外部无法直接访问

    @property
    def name(self):
        return self.__name

    @name.setter
    def name(self, value):
        self.__name = value


d = Duck('dck')
print(f'd.name {d.name}')
d.name = 'dk'
print(f'd.name {d.name}')
# AttributeError: 'Duck' object has no attribute '__name'. Did you mean: 'name'?
# print(f'd.__name {d.__name}') # 不能直接访问内置属性
print('---------------------------')


# 实例的内置方法(魔法方法)
class Word:
    def __init__(self, text):
        self.text = text

    def __eq__(self, other):
        return self.text.lower() == other.text.lower()

    def __str__(self):
        return f'python forever: {self.text}'

    def __repr__(self):
        return 'Word("' + self.text + '")'

    def __mul__(self, other):
        return self.text * other.text


first = Word('hello')
first
print(first)
second = Word('HELLO')
print(f'first == second {first == second}')

n1 = Word(4)
n2 = Word(5)
print(f'实例 n1 * 实例 n2 = {n1 * n2}')
print('---------------------------')


# __new__ 内置魔法方法，创建新实例
class Animal:
    _instance = None

    # 先调用
    # 实现单例模式
    def __new__(cls, *args, **kwargs):
        print('Animal __new__...')
        if cls._instance is None:
            print('创建唯一实例')
            cls._instance = super().__new__(cls)
        return cls._instance

    # 初始化实例
    def __init__(self, *args, **kwargs):
        print('Animal __init__...')
        self.args = args
        self.kwargs = kwargs


ani = Animal(1, 2, 3, x=10, y=20)
print(f'ani = {ani}')
# Animal __new__...
# Animal __init__...
# ani = <__main__.Animal object at 0x0000023215769D90>
print('---------------------------')


# __getattr__ 和 __getattribute__
class User:
    def __init__(self, name, info={}):
        self.name = name
        self.info = info

    # 未找到属性时调用此方法
    def __getattr__(self, item):
        return self.info[item]


print('__getattr__: 未找到属性时调用此方法')
u1 = User('Tom', {'age': 18, 'sex': 'male'})
print(f'u1.name {u1.name}')
print(f'u1.age {u1.age}')
print(f'u1.sex {u1.sex}')
print('---------')


class User2:
    def __init__(self, name, info={}):
        self.name = name
        self.info = info

    # 每次访问属性时都调用此方法
    def __getattribute__(self, item):
        return 'hello world'


print('__getattribute__: 每次访问属性时都调用此方法')
u2 = User2('Jerry', {'age': 20, 'sex': 'female'})
print(f'u2.name {u2.name}')
print(f'u2.age {u2.age}')
print(f'u2.sex {u2.sex}')
print('---------------------------')


# 类属性，所有实例继承
class Fruit:
    color = 'red'  # 类属性，所有实例继承


print('类属性, 所有实例继承, 实例同名属性会覆盖类属性')
banana = Fruit()
print(f'Fruit.color {Fruit.color}')  # red
print(f'banana.color {banana.color}')  # red
print('---------')
print(f'修改 Fruit.color = "yellow"')
Fruit.color = 'yellow'
print(f'Fruit.color {Fruit.color}')  # yellow
print(f'banana.color {banana.color}')  # yellow, 实例未修改过的同名属性也会变化
print('---------')
print(f'修改 banana.color = "green"')
banana.color = 'green'
print(f'Fruit.color {Fruit.color}')  # yellow
print(f'banana.color {banana.color}')  # green
print('---------')
print(f'修改 Fruit.color = "blue"')
Fruit.color = 'blue'
print(f'Fruit.color {Fruit.color}')  # blue
print(f'banana.color {banana.color}')  # green
print('---------')
orange = Fruit()
print(f'Fruit.color {Fruit.color}')  # blue
print(f'orange.color {orange.color}')  # blue
print('---------------------------')


# 类方法，所有实例共享
class A():
    count = 0

    def __init__(self):
        A.count += 1

    # 类方法, 所有实例共享
    @classmethod
    def count_func(cls):
        print(f'class A has {cls.count} instances...')


class AA(A):
    pass


print('类方法, 实例和类都可以调用')
a = A()
a.count_func()
b = A()
b.count_func()
aa = AA()
aa.count_func()
AA.count_func()
print('---------------------------')


# 静态方法, 不需要实例化直接调用
class Animal:
    @classmethod
    def class_func(cls):
        print(f'1. @classmethod of class {cls.__name__}')

    @staticmethod
    def static_func():
        print(f'1. @staticmethod of class Animal')

    def breathe(self):
        print(f'Animal instancemethod breathe...')


class Dog(Animal):
    @classmethod
    def class_func(cls):
        super().static_func()
        Animal.static_func()

        super().class_func()
        Animal.class_func()
        print(f'2. @classmethod of class {cls.__name__}')

    def breathe(self):
        # 使用 super() 或 类名直接调用类静态方法
        super().static_func()
        Animal.static_func()

        super().class_func()
        Animal.class_func()
        print(f'Dog instancemethod breathe...')


print('静态方法, 实例方法内调用静态方法和类方法')
d = Dog()
d.breathe()
print('--------')
print('静态方法, 实例调用类方法, 类方法调用静态方法')
d = Dog()
d.class_func()
print('--------')
print('静态方法, 实例直接调用静态方法')
d = Dog()
d.static_func()
print('---------------------------')

print(r'''
多继承, 访问自己没有的属性或方法时, 优先使用最先继承的父类的属性和方法
按 MRO(method resolution order) 方法解析顺序查找方法或属性的规则
''')
print('--------')
print('菱形继承:')


class D:
    name = 'cls_D'


class B(D):
    # name = 'cls_B'
    pass


class C(D):
    name = 'cls_C'


class A(B, C):
    # name = 'cls_A'
    pass


a = A()
print(f'a.name {a.name}\nA.__mro__ {A.__mro__}')
# a.name cls_C
# A.__mro__ (<class '__main__.A'>, <class '__main__.B'>, <class '__main__.C'>, <class '__main__.D'>, <class 'object'>)
print('--------')
print('树形继承:')


class E:
    name = 'cls_E'


class D:
    name = 'cls_D'


class B(D):
    # name = 'cls_B'
    pass


class C(E):
    name = 'cls_C'


class A(B, C):
    # name = 'cls_A'
    pass


a = A()
print(f'a.name {a.name}\nA.__mro__ {A.__mro__}')
# a.name cls_D
# A.__mro__ (<class '__main__.A'>, <class '__main__.B'>, <class '__main__.D'>, <class '__main__.C'>, <class '__main__.E'>, <class 'object'>)
print('---------------------------')


# Duck typing
class Quote:
    def __init__(self, person, words):
        self.person = person
        self.words = words

    def who(self):
        return self.person

    def says(self):
        return self.words+'.'


class QuestionQuote(Quote):
    def says(self):
        return self.words+'?'


class ExclamationQuote(Quote):
    def says(self):
        return self.words+'!'


class SubQuestionQuote(QuestionQuote):
    def says(self):
        return self.words+'??'


def who_says(obj):
    print(f'{obj.who()} says: {obj.says()}')


print('多态性:')
hunter = Quote('Elmer Fudd', 'I\'m hunting wabbits')
# 如果子类没有定义初始化方法, python 自动调用父类的初始化方法完成属性绑定
hunter1 = QuestionQuote('Bugs Bunny', 'What\'s up, doc')
hunter2 = ExclamationQuote('Daffy Duck', 'It\'s rabbit season')
who_says(hunter)
who_says(hunter1)
who_says(hunter2)
hunter3 = SubQuestionQuote('Jerry', 'It\'s Tom')
who_says(hunter3)
print('-----------')
# 实例化时, 参数必须符合初始化方法的参数要求
# TypeError: Quote.__init__() takes 3 positional arguments but 4 were given
# hunter3 = SubQuestionQuote('Jerry', 'It\'s Tom', 'good')
# who_says(hunter3)
# TypeError: Quote.__init__() missing 2 required positional arguments: 'person' and 'words'
# hunter3 = SubQuestionQuote()
# who_says(hunter3)
print('---------------------------')


# 数据类
@dataclass
class DataClass:
    name: str
    age: int
    height: float


dc = DataClass('Tom', 18, 1.75)
print(f'数据类: {dc} 类型为 {type(dc)}')
print('---------------------------')


# 定义一个元类（继承自 type）
class MyMeta(type):
    """自定义元类"""

    def __new__(cls, name, bases, attrs):
        print(f"元类 __new__ 被调用")
        print(f"  创建类: {name}")
        print(f"  基类: {bases}")
        print(f"  属性: {list(attrs.keys())}")

        # 可以修改类的属性
        attrs['created_by'] = 'MyMeta'
        attrs['version'] = '1.0'

        return super().__new__(cls, name, bases, attrs)

    def __init__(cls, name, bases, attrs):
        print(f"元类 __init__ 被调用")
        print(f"  初始化类: {name}")
        super().__init__(name, bases, attrs)


# 使用元类
class MyClass(metaclass=MyMeta):
    x = 10

    def hello(self):
        return "Hello"


print(f"MyClass.created_by: {MyClass.created_by}")  # MyMeta
print(f"MyClass.version: {MyClass.version}")
print('-----------')


class SingletonMeta(type):
    """单例元类"""
    _instances = {}

    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            print(f"创建 {cls.__name__} 的第一个实例")
            cls._instances[cls] = super().__call__(*args, **kwargs)
        else:
            print(f"返回 {cls.__name__} 的已有实例")
        return cls._instances[cls]


class Database(metaclass=SingletonMeta):
    def __init__(self):
        self.connection = None
        print("Database.__init__ 被调用")

    def connect(self):
        if not self.connection:
            self.connection = "Connected"
        return self.connection


# 测试
db1 = Database()
db2 = Database()
print(f"db1 is db2: {db1 is db2}")  # True
print(f"db1.connect(): {db1.connect()}")
print('---------------------------')

print('函数创建类')


# 函数创建类
def UpperMetaClass(class_name, class_bases, class_attrs):
    # 字典推导式将类属性非下划线开头的转换为大写
    new_attrs = dict((key, value) if key.startswith('_') else (
        key.upper(), value) for key, value in class_attrs.items())
    return type(class_name, class_bases, new_attrs)


class Person(object, metaclass=UpperMetaClass):
    name = 'Tom'
    _age = 18


print(f'hasattr(Person, name): {hasattr(Person, "name")}')  # False
print(f'hasattr(Person, "NAME"): {hasattr(Person, "NAME")}')    # True
print(f'hasattr(Person, "_AGE"): {hasattr(Person, "_AGE")}')    # False
print(f'hasattr(Person, "_age"): {hasattr(Person, "_age")}')  # True
print('----------')

print('元类创建类')


# 元类创建类
class UpperMetaClass(type):
    def __new__(cls, class_name, class_bases, class_attrs):
        # 字典推导式将类属性非下划线开头的转换为大写
        new_attrs = dict((key, value) if key.startswith('_') else (
            key.upper(), value) for key, value in class_attrs.items())
        return super().__new__(cls, class_name, class_bases, new_attrs)


class Person1(metaclass=UpperMetaClass):
    name = 'Tom'
    _age = 18


print(f'hasattr(Person1, name): {hasattr(Person1, "name")}')  # False
print(f'hasattr(Person1, "NAME"): {hasattr(Person1, "NAME")}')    # True
print(f'hasattr(Person1, "_AGE"): {hasattr(Person1, "_AGE")}')    # False
print(f'hasattr(Person1, "_age"): {hasattr(Person1, "_age")}')  # True
print('---------------------------')
