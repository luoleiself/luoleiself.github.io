from datetime import datetime

import sqlalchemy
from sqlalchemy import create_engine, text, Column, Integer, MetaData, Table, String, DateTime, ForeignKey, select

"""
mysqldb: 仅支持 python 2.x
mysqlclient: mysqldb 的分支, 支持 python 3, 使用 C 语言编写的库, 直接绑定到 mysql 的 C API, 性能比下面的快.
pymysql: 使用纯 python 语言编写的 mysql 数据库驱动, 在纯 python 环境中更容易部署.

sqlalchemy:
    # 创建并返回 sql 查询语句, 使用 :name 绑定参数, 在执行语句时传入参数
    text("SELECT * FROM users WHERE id=:user_id")
    func:   # 聚合函数
        func.now()  # 当前日期时间
        func.max()  # 最大值
    declarative_base()  # 创建映射关系基类实例, 2.0 开始使用 DeclarativeBase 代替
    
    create_engine() # 创建数据库引擎
    sessionmaker() # 创建 session 会话连接
    create_async_engine() # 创建异步数据库引擎
    async_sessionmaker()  # 创建异步 session 会话连接
    
    方式一: 使用元表
    meta_data = MetaData(): 创建元表, 创建所有表时传入元表实例
        meta_data.create_all(engine)  # 创建所有表
    Table(table_name, meta_data, *[Column()]): 定义表结构
    Column(): 定义列, 2.0 之前写法
    ForeignKey(): 定义关系外键
    
    方式二: 创建表关系映射类
    DeclarativeBase: 所有映射关系类的基类, 必须继承此类的子类, 否则报错
        sqlalchemy.exc.InvalidRequestError: Cannot use 'DeclarativeBase' directly as a declarative base class. Create a Base by creating a subclass of it.
        [subClass].metadata.create_all(engine)  # 创建所有表
    Mapped: 表示类的 ORM 映射关系, 定义类属性类型
    mapped_column(): 定义类属性到表列的 ORM 关系映射, 2.0 新写法
    relationship(): 定义类之间的关系以及查询数据的关联关系的处理方式
        lazy,   # 懒惰查询
        backref,    # 反向引用, 值应为对方类中 relationship 属性的名称
        back_populates, # 双向关联关系, 值应为对方类中 relationship 属性的名称
        secondary:  # 多对多关系关联表
    aliased(): 创建别名
    scalars(): 过滤返回结果的 ScalarResult 对象

    unique(): 去除 join 产生的重复行
    select(): 创建基于 ORM 的 select 语句
    update(): 创建基于 ORM 的 update 语句
    delete(): 创建基于 ORM 的 delete 语句
"""

DATABASE_URL = 'mysql+mysqldb://appuser:appuserpassword@172.31.218.169:3306/app?charset=utf8mb4'
engine = create_engine(DATABASE_URL, echo=True)
# conn = engine.connect() # 创建连接

# 方式一: 使用 MetaData 元表, Table 类创建表结构
# 创建元表 MetaData 实例
meta_data = MetaData()
# 创建表结构 Table 实例
employee = Table('employees', meta_data,
                 Column('id', Integer, autoincrement=True, primary_key=True, nullable=False, comment='id'),
                 Column('name', String(128), nullable=False, comment='姓名'),
                 Column('age', Integer, comment='年龄'),
                 Column('birthday', DateTime, comment='出生日期'),
                 # 外键
                 Column('department_id', Integer, ForeignKey('departments.id'), nullable=False, comment='部门ID'))

department = Table('departments', meta_data,
                   Column('id', Integer, autoincrement=True, primary_key=True, nullable=False, comment='部门ID'),
                   Column('name', String(128), nullable=False, unique=True, comment='部门名称'))
# 使用元表实例创建表
meta_data.create_all(engine)

# 创建 sql 语句
t = text("SELECT * FROM employee WHERE id=:user_id")
print(f't = {t}')

# 使用 engine 连接执行数据库操作
with engine.connect() as conn:
    # 创建插入语句
    employee_length = conn.execute(employee.select()).rowcount
    print(f'employee_length = {employee_length}')
    if employee_length == 0:
        insert_one = employee.insert().values([
            {'name': '张三', 'age': 18, 'birthday': datetime.now()},
            {'name': '张三1', 'age': 19, 'birthday': datetime.now()}
        ])
        result = conn.execute(insert_one)

    # # 创建更新语句
    update_one = employee.update().values(age=20).where(employee.c.id == 3)
    result = conn.execute(update_one)
    print(f'result = {result.rowcount} {result.is_insert} {result.returns_rows}')
    print('-' * 10)
    # # 创建查询语句
    query_stmt = employee.select()
    result = conn.execute(query_stmt)
    print(f'result = {result.fetchall()}')
    print('-' * 10)
    # # 创建删除语句
    # delete_one = employee.delete().where(employee.c.id == 1)

    # 1
    join_query1 = select(employee).join(department, employee.c.department_id == department.c.id).where(
        department.c.id == 1)
    print(f'join_query1 = {join_query1}')
    result = conn.execute(join_query1)
    print(f'result = {result.fetchall()}')
    print('-' * 10)
    # 2
    # # 创建连接查询语句
    join_statement = employee.join(department, employee.c.department_id == department.c.id)
    join_query2 = select(employee).select_from(join_statement).where(employee.c.id == 1)
    result = conn.execute(join_query2)
    print(f'result = {result.fetchall()}')
    print('-' * 10)
    conn.commit()

if __name__ == '__main__':
    print(f'sqlalchemy.__version__: {sqlalchemy.__version__}')
