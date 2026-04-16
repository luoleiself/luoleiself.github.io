from datetime import datetime
from typing import Annotated

import sqlalchemy
from sqlalchemy import String, Integer, Index, create_engine, Column, DateTime, ForeignKey, func, select, update, \
    delete, Table
from sqlalchemy.orm import Mapped, mapped_column, DeclarativeBase, sessionmaker, Session, relationship, aliased, \
    declarative_base

"""
多对多关系, 关联表包含除 左表和右表 的外键之外的其他列时, 使用 关联对象(class) 模式.

relationship():
    # lazy=False # 取消懒惰查询, 一次将关联关系的数据也查询出来, 否则在访问关联关系属性时会自动执行多次 select 查询
    # backref='' # 反向引用, 值应为对方类中 relationship 属性的名称, 
                 # 表示在 Department 类中可以通过 employees 属性访问关联的 Employee 对象
    # back_populates='' # 双向关联关系, 值应为对方类中 relationship 属性的名称
    # secondary='' # 多对多关系关联表
CursorResult:   # execute 执行后返回的结果
    rowcount: # 受影响的行数
    is_insert:  # (insert)是否成功执行
    inserted_primary_key:   # (insert)获取最后插入的 id
    returns_rows:   # 是否有返回行
"""

DATABASE_URL = 'mysql://appuser:appuserpassword@172.31.218.169:3306/app'
engine = create_engine(DATABASE_URL, echo=True)

# 类型声明
primary_key = Annotated[int, mapped_column(Integer, autoincrement=True, primary_key=True, nullable=False, comment='id')]
stu_birthday = Annotated[datetime, mapped_column(DateTime, default=datetime.now(), comment='生日')]


# 方式二: 使用类创建表关系映射类
# Base = declarative_base() # 创建映射关系基类实例
class Base(DeclarativeBase):
    create_at: Mapped[datetime] = mapped_column(DateTime, server_default=func.now(), comment='创建时间')
    update_at: Mapped[datetime] = mapped_column(DateTime, server_default=func.now(), onupdate=func.now(),
                                                comment='更新时间')


class Employee(Base):
    """
    员工类
    id: 整型, 主键, 自增
    name: 姓名: 字符串(120)，非空
    age: 年龄: 整型, 非空, 默认值 18
    birthday: 生日: 日期时间类型, 默认值 当前日期时间
    department_id: 部门 id, 整型, 外键, 非空
    """
    __tablename__ = 'employees'
    # 创建索引
    __indexes__ = [
        Index('idx_name', 'name'),
    ]
    # # 定义列, 2.0 之前写法
    # id = Column(Integer, autoincrement=True, primary_key=True, nullable=False, comment='id')
    # name = Column(String(128), nullable=False, comment='姓名')
    # age = Column(Integer, nullable=False, comment='年龄')

    # # 定义列, 2.0 新写法
    id: Mapped[primary_key]
    name: Mapped[str] = mapped_column(String(128), nullable=False, comment='姓名')
    age: Mapped[int] = mapped_column(Integer, nullable=False, default=18, comment='年龄')
    birthday: Mapped[stu_birthday]
    department_id: Mapped[int] = mapped_column(ForeignKey('departments.id'), nullable=False, comment='部门 id')

    # 双向关联
    # lazy=False # 取消懒惰查询
    # back_populates='employees' # 指向 Department 类的 employees 属性
    department: Mapped['Department'] = relationship(lazy=False, back_populates='employees')

    def __repr__(self) -> str:
        return f"<Employee(id={self.id}, name={self.name}, age={self.age}, birthday={self.birthday}, department_id={self.department_id})>"


class Department(Base):
    """
    部门类
    id: 整型, 主键, 自增
    name: 部门名称: 字符串(128)，非空
    """
    __tablename__ = 'departments'
    id: Mapped[primary_key]
    name: Mapped[str] = mapped_column(String(128), nullable=False, comment='部门名称')

    # 双向关联
    # back_populates='department' # 指向 Employee 类的 department 属性
    employees: Mapped[list['Employee']] = relationship(back_populates='department')

    def __repr__(self) -> str:
        return f"<Department(id={self.id}, name={self.name})>"


# 多对多关系
class User(Base):
    """
    用户类
    id: 整型, 主键, 自增
    name: 用户名称: 字符串(128)，非空
    """
    __tablename__ = 'users'
    id: Mapped[primary_key]
    name: Mapped[str] = mapped_column(String(128), nullable=False, comment='用户名称')

    # 双向关联, 取消懒惰查询
    # secondary='user_roles'    # 指定关联表
    roles: Mapped[list['Role']] = relationship(secondary='user_roles', back_populates='users', lazy=False)

    def __repr__(self) -> str:
        return f"<User(id={self.id}, name={self.name})>"


class Role(Base):
    """
    角色类
    id: 整型, 主键, 自增
    name: 角色名称: 字符串(128)，非空
    """
    __tablename__ = 'roles'
    id: Mapped[primary_key]
    name: Mapped[str] = mapped_column(String(128), nullable=False, comment='角色名称')

    # 双向关联, 取消懒惰查询
    # secondary='user_roles'    # 指定关联表
    users: Mapped[list['User']] = relationship(secondary='user_roles', back_populates='roles', lazy=False)

    def __repr__(self) -> str:
        return f"<Role(id={self.id}, name={self.name})>"


# 多对多关系关联表
user_roles = Table(
    'user_roles',
    Base.metadata,
    # 联合主键
    Column('user_id', ForeignKey('users.id'), primary_key=True, comment='用户 id'),
    Column('role_id', ForeignKey('roles.id'), primary_key=True, comment='角色 id')
)
# 当多对多关系, 关联表包含除 左表和右表 的外键之外的其他列时, 使用 关联对象(class) 模式. 不使用 relationship.secondary 参数; 相反, 类直接映射到关联表.
# 然后, 在关联类中创建两个单独的 relationship() 构造, 首先通过一对多将左表链接到映射的关联类, 然后通过多对一将映射的关联类链接到右表, 形成从左表到关联表再到右表的单向关联对象关系.
# 对于双向关联关系，使用四个 relationship() 构造将映射的关联类双向链接到左表和右表.
# class UserRole(Base):
#     __tablename__ = 'user_roles'
#

# 创建表, 原理同 meta_data 元表
Base.metadata.create_all(engine)

# 事务


# 使用 session 会话管理数据库操作
session_maker = sessionmaker(engine, class_=Session)
with session_maker() as session:
    # 批量添加
    # dep1 = Department(name='部门1')
    # dep2 = Department(name='部门2')
    # session.add_all([dep1, dep2])
    # print(dep1, dep2)
    # session.add_all([
    #     Employee(name='张三', age=18, birthday=datetime(2020, 1, 1), department=dep1),
    #     Employee(name='张三1', age=19, birthday=datetime(2021, 1, 1), department=dep1),
    #     Employee(name='张三2', age=20, birthday=datetime(2022, 1, 1), department=dep2)
    # ])

    # 按年龄查询大于等于 19 岁的员工
    result = session.query(Employee).filter(Employee.age >= 19).all()
    print(f'按年龄查询大于等于 19 岁的员工: {result}')

    # 按年龄降序查询
    result = session.execute(select(Employee).order_by(Employee.age.desc())).all()
    print(f'按年龄降序查询: {result}')

    # 使用 select() 创建查询语句
    result = session.execute(select(Employee).where(Employee.age < 19)).all()
    print(f'按年龄查询小于 19 岁的员工: {result}')

    # 使用 join() 创建查询语句
    join_statement = select(Employee).join(Department).where(Employee.department_id == Department.id).filter(
        Department.name == '部门1').order_by(Employee.age.desc()).limit(10)
    result = session.execute(join_statement).all()
    print(f'查询 部门1 下的员工, 按员工年龄降序排列, 获取前 10 条数据: {result}')

    # 根据 id 更新员工年龄
    result = session.execute(update(Employee).where(Employee.id == 6).values(age=36))
    print(
        f'根据 id 更新员工年龄: result.rowcount {result.rowcount}, result.is_insert {result.is_insert}, result.returns_rows {result.returns_rows}')

    # 删除员工年龄大于等于 65 的员工信息
    result = session.execute(delete(Employee).where(Employee.age >= 65))
    print(
        f'删除员工年龄大于等于 65 的员工信息: result.rowcount {result.rowcount}, result.is_insert {result.is_insert}, result.returns_rows {result.returns_rows}')

    # 使用聚合函数按部门统计人数
    result = session.query(func.count(Employee.id), Department.name).join(Department).group_by(Department.name).all()
    print(f'query 使用聚合函数按部门统计人数: {result}')
    agg_statement = select(func.count(Employee.id), Department.name).join(Department).group_by(Department.name)
    result_execute = session.execute(agg_statement).all()
    print(f'execute 使用聚合函数按部门统计人数: {result_execute}')
    # 平均年龄
    avg_statement = select(func.avg(Employee.age), Department.name).join(Department).group_by(Department.name)
    result = session.execute(avg_statement).all()
    print(f'平均年龄: {result}')
    # 最大年龄
    max_statement = select(func.max(Employee.age), Department.name).join(Department).group_by(Department.name)
    result = session.execute(max_statement).all()
    print(f'最大年龄: {result}')
    # 最小年龄
    min_statement = select(func.min(Employee.age), Department.name).join(Department).group_by(Department.name)
    result = session.execute(min_statement).all()
    print(f'最小年龄: {result}')
    # 年龄总和
    sum_statement = select(func.sum(Employee.age), Department.name).join(Department).group_by(Department.name)
    result = session.execute(sum_statement).all()
    print(f'年龄总和: {result}')

    # 创建别名
    emp = aliased(Employee, name='emp')
    dpt = aliased(Department, name='dpt')

    print('------------------------------------------')
    # 创建 用户 - 角色 多对多关系
    # user1 = User(name='user1')
    # user2 = User(name='user2')
    # user3 = User(name='user3')
    #
    # role1 = Role(name='role1')
    # role2 = Role(name='role2')
    #
    # user1.roles.extend([role1, role2])
    # user2.roles.extend([role2])
    # user3.roles.extend([role1])
    # # 保存用户
    # session.add_all([user1, user2, user3])

    # 查询所有用户
    result = session.query(User).all()
    print(f'查询所有用户: {result}')
    for user in result:
        print(f'用户: {user.name}, 角色: {user.roles}')

    # 查询所有角色
    result = session.query(Role).all()
    print(f'查询所有角色: {result}')
    for role in result:
        print(f'角色: {role.name}, 用户: {role.users}')

    # 查询角色为 role1 的所有用户, query
    result = session.query(User).join(user_roles).join(Role).filter(Role.name == 'role1').all()
    print(f'query 查询角色为 role1 的所有用户: {result}')
    # 查询角色为 role1 的所有用户, execute
    result_execute = session.execute(
        select(User).join(user_roles).join(Role).where(Role.name == 'role1')).unique().all()
    print(f'execute 查询角色为 role1 的所有用户: {result_execute}')

    session.commit()

if __name__ == '__main__':
    print(f'sqlalchemy.__version__: {sqlalchemy.__version__}')
