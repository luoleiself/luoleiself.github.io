import uvicorn
from fastapi import FastAPI, Path, Query, Request
from fastapi.encoders import jsonable_encoder
from fastapi.params import Depends
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel, Field
from starlette.responses import FileResponse, HTMLResponse, JSONResponse
import time
from routers import news  # 导入分组路由

"""
直接安装 fastapi 不包含 standard 标准包的中扩展
FastAPI: annotated-doc, pydantic, starlette, typing-extensions, typing-inspection
    [standard] 包含 email-validator, fastapi-cli, httpx, jinja2, pydantic-settings, python-multipart, uvicorn

APIRouter: 定义模块化路由
    
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


# 添加 CORS 中间件
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# @app.on_event('startup')
# async def startup_event():
#     async with async_engine.begin() as conn:
#         await conn.run_sync(Base.metadata.create_all)
#
#

# 注册路由
app.include_router(news.router)

if __name__ == '__main__':
    uvicorn.run(app, port=8000)
