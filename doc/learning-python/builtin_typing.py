from typing import NewType, Annotated, Any, Literal, LiteralString, Never, Union, Optional, final, Final, Unpack
from collections.abc import Callable

print('''
类型缩写: type t_name = o_type
... 字面量, 作为参数时, 表示可接收任意参数
Any: 任意类型
Final: 锁定类型和值, 不能再次修改, 等价于 Literal[type]
Literal: 字面量值, Literal['r', 'w', 'a']
LiteralString: 字符串字面量, 只能使用字符串字面量
Union: 联合类型, 等价于 x | y
Optional: 可选类型, 等价于 x | None 或 Union[x, None]
Never:  表示函数没有返回值
NeverReturn: 表示函数没有返回值
Self: 表示当前的类本身
Required: 表示必需
NotRequired: 表示非必需
ReadOnly: 表示只读
Unpack: 表示对象可解构
''')
type Vector = list[str]
lt_str: Vector = ['hello', 'world', 1]
print(f'类型缩写: {lt_str} {type(lt_str)}')
print('-----------')
MAX_SIZE: Final = 100
MAX_SIZE += 1   # check error
MAX_SIZE = 'hello world'    # check error
print(f'Final MAX_SIZE: {MAX_SIZE} {type(MAX_SIZE)}')

type Mode = Literal['r', 'w', 'a']
mode: Mode = 'a'
mode = 'x'  # check error
print(f'Literal mode: {mode} {type(mode)}')

s: LiteralString = 'hello world'
s = lt_str  # check error
s = f'{lt_str}'  # check error
print(f'LiteralString s {s} {type(s)}')
print('-----------')
print(
    f'Union[int, str, bool] == int | str | bool {Union[int, str, bool] == int | str | bool}')
print(
    f'Optional[int] == Union[int, None] == int | None {Optional[int] == Union[int, None] == int | None}')
print('---------------------------')

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
print('---------------------------')

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
print('---------------------------')

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
print('---------------------------')
