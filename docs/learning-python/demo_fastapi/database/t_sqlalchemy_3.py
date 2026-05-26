import asyncio
import sqlalchemy
from sqlalchemy import Integer, String, Float, select, func
from sqlalchemy.ext.asyncio import create_async_engine, AsyncSession, async_sessionmaker
from sqlalchemy.orm import declarative_base, DeclarativeBase, Mapped, mapped_column

"""
create_async_engine(): # 创建异步数据库引擎
async_sessionmaker():  # 创建异步 session 会话连接
"""

# 异步操作上数据库
ASYNC_DATABASE_URL = 'mysql+aiomysql://appuser:appuserpassword@172.31.218.169:3306/app'
async_engine = create_async_engine(
    ASYNC_DATABASE_URL,
    pool_size=20,
    max_overflow=10,
    pool_timeout=30,
    pool_recycle=3600,
    pool_pre_ping=True,
    echo=True,
)

# 创建映射关系基类实例
Base = declarative_base()


class Book(Base):
    """
    书籍类
    id: 主键, 整型
    name: 书名, 字符串类型, 非空
    description: 描述, 字符串类型, 非空
    author: 作者, 字符串类型, 非空
    price: 价格, 数值类型
    """
    __tablename__ = 'books'
    id: Mapped[int] = mapped_column(Integer, primary_key=True, autoincrement=True, nullable=False, comment='id')
    name: Mapped[str] = mapped_column(String(255), nullable=False, comment='name')
    description: Mapped[str] = mapped_column(String(255), nullable=True, comment='description')
    author: Mapped[str] = mapped_column(String(255), nullable=True, comment='author')
    price: Mapped[float] = mapped_column(Float, comment='price')

    def __repr__(self):
        return f"<Book(id={self.id}, name={self.name}, description={self.description}, author={self.author}, price={self.price})>"


def create_tables(conn):
    """创建所有表的包装函数"""
    Base.metadata.create_all(conn)


async def init_database():
    """初始化数据库, 创建所有表"""
    async with async_engine.begin() as conn:
        # 直接使用 create_all 将引起类型签名不匹配
        # 或者传入 lambda c: Base.metadata.create_all(c) 函数
        await conn.run_sync(create_tables)


# 创建异步会话连接
AsyncSessionLocal = async_sessionmaker(bind=async_engine, class_=AsyncSession, expire_on_commit=False)


async def main():
    await init_database()
    async with AsyncSessionLocal() as session:
        # 添加数据
        result = await session.execute(select(func.count(Book.id)))
        length = result.scalar()
        print(f'result length {length}')
        if length == 0:
            session.add_all([
                Book(name='book1', description='description1', price=1.1),
                Book(name='book2', description='description2', price=2.2),
                Book(name='book3', description='description3', price=3.3),
                Book(name='book4', description='description4', price=4.4),
                Book(name='book5', description='description5', price=5.5),
                Book(name='book6', description='description6', price=6.6),
                Book(name='book7', description='description7', price=7.7),
                Book(name='book8', description='description8', price=8.8),
            ])
            await session.commit()

        # 查询所有数据
        result = await session.execute(select(Book))
        books = result.scalars().all()
        print(f'查询所有数据: {books}')

        # 按名称查询
        result = await session.execute(select(Book).where(Book.name == 'book1'))
        book = result.scalars().all()
        print(f'查询书名为 book1 的书: {book}')

        # 按价格查询
        result = await session.execute(select(Book).where(Book.price >= 2))
        books = result.scalars().all()
        print(f'查询价格 大于等于 2 的书: {books}')

        # 按价格降序查询
        result = await session.execute(select(Book).where(Book.price >= 5).order_by(Book.price.desc()))
        books = result.scalars().all()
        print(f'查询价格大于等于 5 且降序排列的书: {books}')

        # 按价格降序查询限制数量
        result = await session.execute(select(Book).where(Book.price >= 5).order_by(Book.price.desc()).limit(2))
        books = result.scalars().all()
        for book in books:
            print(f'for in books: {book} {book.id} {book.name} {book.price}')
        print(f'查询价格大于等于 5 且降序排列的书, 限制数量为 2: {books}')

    await async_engine.dispose()


asyncio.run(main())

if __name__ == '__main__':
    print(f'sqlalchemy.__version__: {sqlalchemy.__version__}')
