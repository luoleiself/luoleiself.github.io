import uvicorn
from fastapi import FastAPI, Path, Query, Request
from fastapi.params import Depends
from pydantic import BaseModel, Field
from starlette.responses import FileResponse, HTMLResponse, JSONResponse
import time

"""
直接安装 fastapi 不包含 standard 标准包的中扩展
FastAPI: annotated-doc, pydantic, starlette, typing-extensions, typing-inspection
    [standard] 包含 email-validator, fastapi-cli, httpx, jinja2, pydantic-settings, python-multipart, uvicorn
    
依赖注入: Depends

中间件: @app.middleware('http') 装饰器定义 http 中间件, 按定义顺序自下而上执行
    request: 请求实例
    call_next: 

pydantic: 数据校验库, 定义包含属性的继承 BaseModel 的类, 实例化这个类会自动校验这些属性值,
    并在需要时把它们转换为合适的类型, 返回一个包含所有数据的对象
    from datatime import datatime
    class User(BaseModel):
        id: int
        name: str = 'Tom'
        signed_at: datetime | None = None
        friends: list[int] = []
        
    user = User(**{'id': '123', 'name': 'Jerry', signed_at: '2025-03-29 16:27:30', 'friends': [1, '2', b'3']})
    print(user)
    # User id=123 name='John Doe' signup_ts=datetime.datetime(2025, 3, 29, 16, 27, 30) friends=[1, 2, 3]
    # 校验数据结构
    user.model_validate(request.json())
    
    BaseModel: 所有模型校验和序列化操作的基类
        # 模型校验装饰器
        @model_validator(mode='before' | 'after' | 'wrap')
        # 字段校验装饰器
        @field_validator(field_name, mode='before' | 'after' | 'wrap' | ‘plain’) 
    RootModel: 用于校验没有字段名的顶层 json 值, 比如纯列表 [1, 2, 3], 纯字典{'a': 1}, 或任意单一层级的原始值,
        不是用来替代 BaseModel 的通用容器, 而是解决 整个输入就是一个值, 而非键值对集合 的场景
"""


app = FastAPI()

@app.get('/')
async def index():
    return 'hello world'


@app.get('/fruit/{id}', response_class=JSONResponse)
async def get_book(
    id: int = Path(gt=0, lt=101, title='ID',
                   description='must be between 0 and 101'),
    skip: int = Query(default=0, ge=0, description='must be greater than 0'),
    limit: int = Query(default=20, ge=0, description='must be greater than 0'),
):
    """
    docstring
    """
    return {"id": id, "skip": skip, 'limit': limit}

class Address(BaseModel):
    street: str
    city: str
    zip_code: str


# 注册参数验证
class User(BaseModel):
    username: str = Field(min_length=5, max_length=20,
                          description='username minlength 5 and maxlength 20')
    password: str
    email: str
    address: Address | None = None


# 装饰器定义响应模型, 返回结果必须符合响应模型定义的值
@app.post('/user', response_model=User)
async def register(user: User):
    print(f'received: {user}')
    return {'id': '001', 'username': 'zhangsan'}


# 装饰器定义响应类
@app.get('/html', response_class=HTMLResponse)
async def get_html():
    return '<h1>HTMLResponse<h1>'


# 返回响应类实例
@app.get('/file/name')
async def get_file():
    return FileResponse('./file.html')


# 依赖注入
async def depend_function(
    skip: int = Query(default=0, ge=0, description='must be greater than 0'),
    limit: int = Query(default=20, ge=0, description='must be greater than 0'),
):
    return {'skip': skip, 'limit': limit}


@app.get('/news/list')
async def get_news(commons=Depends(depend_function)):
    return commons


@app.get('/blog/list')
async def get_blog(commons=Depends(depend_function)):
    return commons


# 中间件按定义顺序自下而上执行
@app.middleware('http')
async def print_request_info(request: Request, call_next):
    print('start2...')
    print(f'request.method: {request.method}')
    print(f'request.url: {request.url}')
    print(f'request.query_params: {request.query_params}')
    print(f'request.headers: {request.headers}')
    print(f'request.cookies: {request.cookies}')
    response = await call_next(request)
    print('end2...')
    return response


@app.middleware('http')
async def add_process_time_header(request: Request, call_next):
    print('start1...')
    start_time = time.time()
    response = await call_next(request)
    process_time = time.time() - start_time
    response.headers['X-Process-Time'] = str(process_time)
    response.set_cookie('process_time', str(process_time), max_age=7 * 24 * 60 * 60, path="/")
    print('end1...')
    return response


# from sqlalchemy.ext.asyncio import create_async_pool_from_url, create_async_engine, async_sessionmaker, AsyncSession
#
# ASYNC_DATABASE_URL = 'mysql+aiomysql://root:123456@localhost:3306/test?charset=utf-8'
#
# async_engine = create_async_engine(
#     ASYNC_DATABASE_URL,
#     echo=True,
#     pool_size=10,
#     max_overflow=5,
# )
#
# from sqlalchemy import func, DateTime, Integer, String, select
# from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column
# from datetime import datetime
#
#
# class Base(DeclarativeBase):
#     create_at: Mapped[datetime] = mapped_column(DateTime, insert_default=func.now(), default=func.now, comment='创建时间')
#     update_at: Mapped[datetime] = mapped_column(DateTime, insert_default=func.now(), default=None, comment='更新时间')
#
#
# class Book(Base):
#     __tablename__ = 'book'
#     id: Mapped[int] = mapped_column(Integer, primary_key=True, autoincrement=True, comment='id')
#     name: Mapped[str] = mapped_column(String(20), nullable=False, comment='name')
#     author: Mapped[str] = mapped_column(String(20), nullable=False, comment='author')
#
#
# @app.on_event('startup')
# async def startup_event():
#     async with async_engine.begin() as conn:
#         await conn.run_sync(Base.metadata.create_all)
#
#
# # 创建会话连接
# AsyncSessionLocal = async_sessionmaker(bind=async_engine, class_=AsyncSession, expire_on_commit=False)
# # 获取会话连接
# async def get_database():
#     async with AsyncSessionLocal() as session:
#         try:
#             yield session
#             await session.commit()
#         except Exception:
#             await session.rollback()
#             raise
#         finally:
#             await session.close()
#
#
# @app.get('/book')
# async def get_book(session: AsyncSession = Depends(get_database)):
#     result = await session.execute(select(Book))
#     books = result.all()
#     print(f'result {result} books {books}')
#
#     print(f'result.scalars() {result.scalars()}')
#     print(f'result.scalar() {result.scalar()}')
#     print(f'result.first() {result.first()}')
#     print('-------------')
#
#     b = await session.get(Book, 5)
#     print(f'await session.get(Book, 5) {b}')
#     print('-------------')
#
#     result = await session.execute(select(Book).where(Book.id == 5))
#     print(f'result {result}')
#     print('-------------')
#
#     # await session.execute(select(Book).join(Book, Book.id == Book.id, isouter=True))
#     print('-------------')
#

if __name__ == '__main__':
    uvicorn.run(app, port=8000)
