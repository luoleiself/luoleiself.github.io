# 装饰器 @property, 只读属性
class Circle:
    def __init__(self, radius):
        self.radius = radius

    @property
    def diameter(self):
        return 2 * self.radius


c = Circle(5)
print(f'c.radius {c.radius} c.diameter {c.diameter}')
# AttributeError: property 'diameter' of 'Circle' object has no setter
# c.diameter = 10   # 只读属性
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
print(f'n1 * n2 {n1 * n2}')
print('---------------------------')


# 类属性，所有实例继承
class Fruit:
    color = 'red'  # 类属性，所有实例继承


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
    def toString(cls):
        print(f'class A has {cls.count} instances...')


class AA(A):
    pass


a = A()
a.toString()
b = A()
b.toString()
aa = AA()
aa.toString()
AA.toString()
print('---------------------------')


# 静态方法, 不需要实例化直接调用
class C:
    @staticmethod
    def go_home():
        print(f'class c static method go_home...')


c = C()
c.go_home()
C.go_home()
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


# 实例化未定义初始化方法的子类时, 参数必须符合父类初始化方法参数的要求
class SubQuestionQuote(QuestionQuote):
    def says(self):
        return self.words+'??'


def who_says(obj):
    print(f'{obj.who()} says: {obj.says()}')


hunter = Quote('Elmer Fudd', 'I\'m hunting wabbits')
# 如果子类没有定义初始化方法, python 自动调用父类的初始化方法完成属性绑定
hunter1 = QuestionQuote('Bugs Bunny', 'What\'s up, doc')
hunter2 = ExclamationQuote('Daffy Duck', 'It\'s rabbit season')
who_says(hunter)
who_says(hunter1)
who_says(hunter2)
print('-----------')
# TypeError: Quote.__init__() takes 3 positional arguments but 4 were given
hunter3 = SubQuestionQuote('Daffy Duck', 'It\'s rabbit season', 'good')
who_says(hunter3)
print('---------------------------')
