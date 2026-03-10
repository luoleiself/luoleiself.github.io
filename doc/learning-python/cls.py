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


# 实例的内置方法(魔术方法)
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
print(f'banana.color {banana.color}')  # yellow, 实例会修改过的同名属性也会变化
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


# 多继承, 访问自己没有的属性或方法时, 优先使用最先继承的父类的属性和方法
class NFCPhone:
    name = 'NFCPhone'

    def call_phone(self):
        print(f'NFCPhone {self.name} is calling...')


class Phone:
    name = 'phone'

    def call_phone(self):
        print(f'Phone {self.name} is calling...')


class IPhone(Phone, NFCPhone):
    pass


print('多继承, 访问自己没有的属性或方法时, 优先使用最先继承的父类的属性和方法')
p = IPhone()
p.call_phone()  # Phone phone is calling...
print(f'p name {p.name}')   # p name phone
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
