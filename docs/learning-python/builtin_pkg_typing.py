from typing import TypeVar, TypeAlias, NewType, Annotated, Any, Literal, LiteralString, Never, Union, Optional, \
    Final, Unpack, Mapping, NamedTuple, TypedDict
from collections.abc import Callable

print('''
类型缩写: type t_name = o_type
    TypeAlias: Predicate: TypeAlias = Callable[..., bool] 定义类型缩写, 3.12 开始废弃, 使用 type 定义

... 字面量, 作为参数时, 表示可接收任意参数
Any: 任意类型
Final: 锁定类型和值, 不能再次修改, 等价于 Literal[type]
Literal: 字面量值, Literal['r', 'w', 'a']
LiteralString: 字符串字面量, 只能使用字符串字面量
Mapping: 键值对的通用容器, 3.9 开启废弃, 使用 [key, value] 代替
Union: 联合类型, 等价于 x | y
Optional: 可选类型, 等价于 x | None 或 Union[x, None]
Never:  表示函数没有返回值
NoReturn: 表示函数没有返回值
Self: 表示当前的类本身
Required: 表示必需
NotRequired: 表示非必需
ReadOnly: 表示只读
Unpack: 表示对象可解构
Annotated: 类型元数据上下文
TypeVar: 用于定义泛型类型变量, 表示某个未知但一致的类型, 常用于函数或类中支持多种具体类型但保持类型一致性
NewType: 用于创建名义类型的别名, 即使底层类型相同, 也视为不同类型, 用于增强类型安全
TypeVarTuple: 用于定义泛型元组
NamedTuple: 具名元组类型
TypedDict: 字典类型, 期望所有实例都有一组特定的键, 其中每个键都与一个一致类型的值相关联
''')
type Vector = list[str]
lt_str: Vector = ['hello', 'world', 1]
print(f'类型缩写: {lt_str} {type(lt_str)}')
print('-' * 10)
MAX_SIZE: Final = 100
MAX_SIZE += 1  # check error
MAX_SIZE = 'hello world'  # check error
print(f'Final MAX_SIZE: {MAX_SIZE} {type(MAX_SIZE)}')

type Mode = Literal['r', 'w', 'a']
mode: Mode = 'a'
mode = 'x'  # check error
print(f'Literal mode: {mode} {type(mode)}')

s: LiteralString = 'hello world'
s = lt_str  # check error
s = f'{lt_str}'  # check error
print(f'LiteralString s {s} {type(s)}')
print('-' * 10)
print(
    f'Union[int, str, bool] == int | str | bool {Union[int, str, bool] == int | str | bool}')
print(
    f'Optional[int] == Union[int, None] == int | None {Optional[int] == Union[int, None] == int | None}')
print('-' * 30)

m: Mapping[str, int] = {'hello': 1, 'world': 2, 'name': 'zhangsan'}
print(f'键值对通用容器: Mapping[str, int] {m} {type(m)}')
print('-' * 30)

print('''
NewType 创建新类型, 静态类型检查器把新类型视为原始类型的子类
    type_name = NewType(t_name, o_type)
    variable_name = type_name(v)
''')
UserId = NewType('UserId', int)
user_id1 = int(1)
user_id2 = UserId(1)
print(f'user_id2 {user_id2} 类型为 {type(user_id2)}')
print(f'int(1) == UserId(1) {user_id1 == user_id2}')
print(f'int(1) + UserId(1) {user_id1 + user_id2}')
print('-' * 30)

print('''
TypeVar: 创建类型
T = TypeVar('T')  # Can be anything
S = TypeVar('S', bound=str)  # Can be any subtype of str
A = TypeVar('A', str, bytes)  # Must be exactly str or bytes
''')
T = TypeVar('T')  # Can be anything
S = TypeVar('S', bound=str)  # Can be any subtype of str
A = TypeVar('A', str, bytes)  # Must be exactly str or bytes

t: T = True
t = 'hello'
print(f't: {t} {type(t)}')
s: S = 'hello python'
s = True  # check error
print(f's: {s} {type(s)}')
a: A = b'hello'
a = 250  # check error
print(f'a: {a} {type(a)}')
print('-' * 30)

print('''
TypeVarTuple: 创建元组类型
''')
print('-' * 30)

print('''
NamedTuple: 具名元组

class Employee(NamedTuple):
    name: str
    id: int

Employee = NamedTuple('Employee', [('name', str), ('id', int)])
''')
print('-' * 30)

print('''
TypedDict: 字典类型, 期望所有实例都有一组特定的键, 其中每个键都与一个一致类型的值相关联
class Point2D(TypedDict):
    x: int
    y: int
    label: str

a: Point2D = {'x': 1, 'y': 2, 'label': 'hello world'}
b: Point2D = {'x': 1, 'y': 2}   # check error
''')
print('-' * 30)

print('''
Callable: 可调用类型, 定义可调用对象接收一个 int 类型参数, 返回字符串类型
def feeder(get_next_item: Callable[[int], str]) -> None:
    pass
    
... 字面量作为参数时, 定义可调用对象接收任意参数
te: tuple[int, ...] = (1, 2, 3)
te = (1, 'hello') # check error
''')


def feeder(get_next_item: Callable[[], str]) -> None:
    get_next_item()
    pass


te: tuple[int, ...] = (1, 2, 3)
te = (1, 'hello')  # check error
print(f'te: {te} {type(te)}')
print('-' * 30)

print('''
元数据类型提示: Annotated[<type>, <metadata>]
Annotated[Any, \'\'\'类型描述\'\'\']
''')

var_any: Annotated[Any, '''类型描述'''] = True
print(f'Annotated var_any: {var_any} {type(var_any)}')
var_any = 'hello world'
print(f'Annotated var_any: {var_any} {type(var_any)}')

var_int: Annotated[int, '''整数的描述'''] = 125
var_int = 'hello world'  # check error
print(f'Annotated var_int: {var_int} {type(var_int)}')
print('-' * 30)
