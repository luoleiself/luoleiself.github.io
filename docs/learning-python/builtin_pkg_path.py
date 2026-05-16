import os
from pathlib import Path, PurePath

"""
pathlib 模块提供表示文件系统路径的类, 其语义适用于不同的操作系统.
    路径类分为纯路径和具体路径, 纯路径提供没有 I/O 的纯计算操作.
    具体路径继承纯路径但也提供 I/O 操作.
    PurePath
    |-  PurePosixPath(PurePath)
    |-  PureWindowsPath(PurePath)
    |-  Path(PurePath)
        |- PosixPath(Path, PurePosixPath)
        |- WindowsPath(Path, PureWindowsPath)
"""

print(f'os.path {os.path}')
print(
    f'os.path.curdir {os.path.curdir} os.path.pardir {os.path.pardir} os.path.extsep {os.path.extsep}')
print(f'os.path.basename(__file__) {os.path.basename(__file__)}')
print(f'os.path.dirname(__file__) {os.path.dirname(__file__)}')
print(f'os.path.abspath(__file__) {os.path.abspath(__file__)}')
print(f'os.path.realpath(__file__) {os.path.realpath(__file__)}')
print(f'os.path.exists(__file__) {os.path.exists(__file__)}')
print(f'os.path.isfile(__file__) {os.path.isfile(__file__)}')
print(f'os.path.isabs("test.txt") {os.path.isabs("test.txt")}')
print(f'os.path.isdir(__file__) {os.path.isdir(__file__)}')
print(f'os.path.islink(__file__) {os.path.islink(__file__)}')
print(f'os.path.ismount(__file__) {os.path.ismount(__file__)}')
print('-' * 20)

print('pathlib 模块: 基于对象的文件路径')
print('PurePath 类')
path_init_1 = PurePath(__file__)
print(f'PurePath(__file__) {path_init_1}')
path_init_2 = PurePath('foo', 'some/path', 'bar')
print(f'PurePath("foo", "some/path", "bar") {path_init_2}')
print(f'PurePath("/etc, "/usr", "lib64") {PurePath("/etc", "/usr", "lib64")}')
print('-' * 4)

print('/ 拼接返回新路径, 如果拼接路径为绝对路径则忽略前一个路径')
print(f'path_init_2 / baz => {path_init_2 / "baz"}')
print(f'path_init_2 / /baz => {path_init_2 / "/baz"}')
print(f'path_init_2.joinpath("boo") {path_init_2.joinpath("boo")}')
print(path_init_2)
print('-' * 4)

print(f'path_init_1.parts {path_init_1.parts}')
print(f'path_init_1.driver {path_init_1.drive}')
print(f'path_init_1.root {path_init_1.root}')
# 驱动器和根的连接
print(f'path_init_1.anchor {path_init_1.anchor}')
for part in path_init_1.parents:
    print('parents', part)

print(f'path_init_1.parent {path_init_1.parent}')
print(f'path_init_1.name {path_init_1.name}')
print(f'path_init_1.suffix {path_init_1.suffix}')
print(f'path_init_1.suffixes {path_init_1.suffixes}')
print(f'排除后缀名 path_init_1.stem {path_init_1.stem}')
print(f'path_init_1.as_posix() {path_init_1.as_posix()}')
print(f'返回相对于指定路径的路径: path_init_2.relative_to("foo") {path_init_2.relative_to("foo")}')
print('-' * 4)

print(f'返回新文件名的路径: path_init_1.with_name("bar.txt") {path_init_1.with_name("bar.txt")}')
print(f'返回新文件名(保留原扩展名)的路径: path_init_1.with_stem("baz") {path_init_1.with_stem("baz")}')
print(f'返回新扩展名的路径: path_init_1.with_suffix(".baz") {path_init_1.with_suffix(".baz")}')
print('-' * 4)

print(f'path_init_1.is_absolute() {path_init_1.is_absolute()}')
print(f'path_init_2.is_absolute() {path_init_2.is_absolute()}')
print(f'判断是否相对于指定路径: path_init_2.is_relative_to("foo") {path_init_2.is_relative_to("foo")}')
print(f'判断是否相对于指定路径: path_init_2.is_relative_to("faz") {path_init_2.is_relative_to("faz")}')
print(f'path_init_2.full_match("foo/") {path_init_2.match("foo/")}')
print(f'path_init_2.match("some/**/*") {path_init_2.match("some/**/*")}')
print('-' * 4)

print('Path 类, 继承 PurePath')
path_init_3 = Path()
print(f'Path() {path_init_3}')
print('类方法 =>')
# 3.13 新增
# print(f'Path.from_uri("file:///etc/passwd") {Path.from_uri("file:///etc/passwd")}')
print(f'返回用户家目录的新路径对象: Path.home() {Path.home()}')
print(f'返回当前工作目录的新路径对象: Path.cwd() {Path.cwd()}')
print('实例方法 =>')
print(f'返回绝对路径, 不进行规范化和符号链接转换: path_init_3.absolute() {path_init_3.absolute()}')
print(f'返回绝对路径, 解析任何符号链接: path_init_3.resolve() {path_init_3.resolve()}')
print(f'path_init_3.parent {path_init_3.parent}')
print('-' * 20)
