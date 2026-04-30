import hashlib
import uuid
import time
import json
from contextlib import asynccontextmanager
from typing import Annotated, Any, Iterable, AsyncIterable

import fastapi
import uvicorn
from fastapi import FastAPI, Path, Query, Request, UploadFile, HTTPException
from fastapi.encoders import jsonable_encoder
from fastapi.params import Depends, Cookie, Header, Form, File
from fastapi.middleware.cors import CORSMiddleware
from fastapi.security import OAuth2PasswordBearer, OAuth2PasswordRequestForm
from fastapi.templating import Jinja2Templates

from pydantic import BaseModel, Field, EmailStr, UUID4, Secret, FileUrl, HttpUrl, AliasPath
# from pydantic_extra_types.coordinate import Coordinate
# from pydantic_extra_types.currency_code import Currency
# from pydantic_extra_types.mac_address import MacAddress
# from pydantic_extra_types.timezone_name import TimeZoneName

from starlette import status
from starlette.responses import FileResponse, HTMLResponse, JSONResponse, StreamingResponse, Response
from starlette.staticfiles import StaticFiles

from routers import news  # 导入分组路由

"""
# 配置应用位置
[tool.fastapi]
entrypoint = 'main:app' # from main import app  '模块名:应用名'

直接安装 fastapi 不包含 standard 标准包的中扩展
FastAPI: annotated-doc, pydantic, starlette, typing-extensions, typing-inspection
    [standard] 包含 email-validator, fastapi-cli, httpx, jinja2, pydantic-settings, python-multipart, uvicorn

    root_path=''    # 期望客户端通过代理服务器访问应用的路径前缀, 代理服务器仍然以正常路径访问应用.
                    # 可通过 fastapi cli 参数 --root-path 传入, uvicorn 直接传入应用.
                    # 通过 FastAPI 构造函数配置.

路径操作按顺序依次运行, 当同时存在同层级路径中包含静态路径和动态路径参数时, 静态路径优先在前.
fastapi 通过参数名称、类型、默认值声明检测(顺序无关紧要)参数, 自动识别同时声明的多个路径参数和查询参数, pydantic 模型请求体.
如果非路径参数声明默认值, 则不是必选, 如果只想把参数设为可选, 但又不想指定参数的值, 则把默认值设置为 None.
  函数参数识别规则:(未使用 Path, Query, Body 声明参数类型时)
    - 如果该参数也在路径中声明了, 它就是路径参数.
    - 如果该参数是（int、float、str、bool 等）单一类型, 它会被当作查询参数.
    - 如果该参数的类型声明为 Pydantic 模型, 它会被当作请求体.
  Path: 为路径参数声明类型校验, Query: 为查询参数声明类型校验, Body: 为请求体参数声明类型校验, 
  Form: 声明 Form 参数, 必须使用 Form 声明, 否则该参数将被解释为查询参数或请求体参数. 继承 Body 类
  File: 声明 File 参数, 必须使用 File 声明, 否则该参数将被解释为查询参数或请求体参数. 继承 Form 类
  UploadFile:
  Cookie: 声明 cookie 参数, 必须使用 Cookie 声明, 否则该参数将被解释为查询参数.
  Header: 声明 header 参数, 必须使用 Header 声明, 否则该参数将被解释为查询参数, convert_underscores=False 禁用 Header 的自动转换下划线为连字符.
    - 别名参数: Query(alias='') fastapi 使用别名在 URL 中查找参数值
    - 弃用参数: Query(deprecated=True) 在文档上标识参数已弃用
    - 排除参数: Query(include_in_schema=False)  标识参数不会出现在生成的 OpenAPI 模式中(也不会出现在自动文档系统中)
    - 自定义校验函数: 使用 pydantic 的 AfterValidator 在内部验证逻辑之后应用
  查询参数模型: 规则适用于 Cookie 参数模型和 Header 参数模型
    class FilterParams(BaseModel):
        model_config = {'extra': 'forbid'}  # 限制 pydantic 模型接收查询参数
        limit: int = Field(100, gt=0, le=1000)
        offset: int = Field(0, ge=0)
        order_by: Literal['created_at', 'updated_at'] = 'created_at'
        tags: list[str] = []
    @app.get('/items')
    async def read_items(filter_query: Annotated[FilterParams, Query()]):
        return filter_query
  嵌入单个请求体参数: Body(embed=True) 请求体参数嵌入到键中, embed=True
    class Item(BaseModel):
        name: str
        description: str | None = None
        price: float
        tax: float | None = None
    @app.get('/items/{item_id}')
    async def read_items(
        item_id: Annotated[int, Path(gt=0, description='must be greater than 0')],
        item: Annotated[Item, Body(embed=True)]
    ):
        results = {'item_id': item_id, 'item': item}
        return results
    # 期望请求体: { "item": {"name': "", "description": "", "price": 0.1, "tax": 0.5}}
  返回类型: fastapi 对返回数据进行校验, 在 OpenAPI 路径操作中为响应添加 JSON Schema, 使用 pydantic 将返回数据序列化为 JSON, 
    将输出数据限制并过滤为返回类型中定义的内容.
  response_model: 装饰器参数, 希望返回的数据和声明的类型不完全一致, 优先级高于声明的返回类型.
    如果路径操作函数的返回类型使用 pydantic 模型会被自动完成文档、校验等工作. 返回的数据不符合返回类型时在 IDE 中会报错.
    response_model=None # 禁用响应模型生成
    response_model_exclude_unset=True   # 设置响应中只包含实际设置的值而忽略默认值
    response_model_exclude_defaults=True
    response_model_exclude_none=True
    response_model_include={''}   # 设置响应中包含指定的属性, 传入 list 或 tuple 时, fastapi 会自动转换为 set.
    response_model_exclude={''}   # 设置响应中排除指定的属性, 传入 list 或 tuple 时, fastapi 会自动转换为 set.
  response_class: 响应类
    执行顺序(性能降序): response_model -> jsonable_encoder -> response_class
  responses: 为主响应附加其他响应信息  
  
  流式传输 JSON Lines: 在路径操作函数中使用 yield 逐个产生数据项, fastapi 自动识别并确保正确运行

Request, Response
    
异常处理装饰器
    app.add_exception_handler() # 添加异常处理函数
    @app.exception_handlers()   # 添加异常处理装饰器
        - RequestValidationError # 内部默认请求包含无效数据异常
        - HTTPException

    @app.exception_handlers(UnicornException)    # 注册 unicorn 异常处理器
    async def unicorn_exception_handler(request: Request, exc: UnicornException):
        return JSONResponse(status_code=418, content={'message': f'{exec.name} did something')

jsonable_encoder(): 将数据类型(如 pydantic 模型)转换为 JSON 兼容的数据类型
  
中间件: 是一个函数, 在每个特定的路径操作函数处理每个请求之前运行, 在返回每个响应之前运行.
    app.add_middleware():   # 添加中间件
    @app.middleware('http') # 装饰器定义 http 中间件, 按定义顺序自下而上执行
        request: 请求实例
        call_next: 将 request 传递给下一个中间件并返回 response
    
依赖注入(Depends): 声明运行所需的依赖项, 由框架执行时自动注入声明的依赖项
    执行过程: 嵌套依赖项会优先调用最内层的依赖项并获取结果传递给外层依赖项
        1. 接收到新的请求时, fastapi 会负责用正确的参数调用依赖项
        2. 获取依赖项返回的结果
        3. 将结果赋值给路径操作函数中的参数
    类作为依赖项: Annotated声明类作为依赖项类型时, Depends 参数可省略, fastapi 自动识别.
    子依赖项: fastapi 支持创建任意深度的子依赖项.
    路径操作装饰器依赖项: dependencies, 接收依赖注入组成的 list, 忽略依赖项的返回值(有无).
    全局依赖项: app = FastAPI(dependencies=[Depends(verify_token)]), 为整个应用(所有路径操作)添加依赖项, 和路径操作装饰器依赖项一致.
    使用 yield 依赖项: 在完成后执行一些额外步骤的依赖项, 在依赖项里只使用一次 yield.
        async def get_db():
            db = DBSession()
            try:
                yield db
                await db.commit()
            finally:
                db.close()

    Depends(get_value, use_cache=False) # 禁用同一个请求中多次调用同一个依赖项时使用缓存值.
    Depends(get_value, scope='function')   # 在路径操作函数返回之后, 但在响应发送之前关闭该依赖.
        scope='function',     # 在处理请求的路径操作函数之前启动依赖, 在路径操作函数结束后结束依赖, 但在响应发送给客户端之前.
        scope='request', 默认, # 在处理请求的路径操作函数之前启动依赖, 但在响应发送给客户端之后结束依赖.
    
    Security()  # 当需要声明 OAuth2 作用域时, 使用 Security() 代替 Depends().
        与 Depends() 的唯一区别是可以声明将与 OpenAPI 和自动UI文档集成的 OAuth2 作用域, 默认情况位于 /docs.

安全:
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl="token")  # 创建发送用户信息认证获取令牌
    OAuth2PasswordRequestForm: 使用 username, password, scope, grant_type, client_id, client_secret 组成的声明表单请求体的类.

app.mount() # 挂载到指定路径
app.include_router()    # 注册模块路由
    prefix=''   # 全局路由前缀   
APIRouter: 定义模块化路由, 选项和 FastAPI 相同
    prefix=''   # 模块化路由前缀

静态文件: StaticFiles: # 静态文件类
    app.mount('/static', StaticFiles(directory='static'), name='static')

ServerSentEvent: 配置 SSE 响应的字段: event, id, retry, comment.


后台任务: 返回响应之后后台运行的任务
    BackgroundTasks:
        .add_task() # 添加后台任务
            - 接收后台运行任务函数.
            - 应按顺序传递给任务函数的任意参数序列.
            - 应传递给任务函数的任意关键字参数.
        
        async def send_notification(email: str, background_tasks: BackgroundTasks):
            background_tasks.add_task(send_email, email, message='some notification')
            return {'message': 'Notification sent in the background'}


pydantic: 数据校验库, 定义包含属性的继承 BaseModel 的类, 实例化这个类会自动校验这些属性值, 
    内置支持数据类(可以不显式声明继承 BaseModel), 数据类不具有 pydantic 模型的所有能力
    并在需要时把它们转换为合适的类型, 返回一个包含所有数据的对象
    fastapi 支持自动识别与路径参数匹配的函数参数从路径中获取, 声明为 pydantic 模型的函数参数从请求体中获取.
        
    pydantic 模型配置
    model_config = {
        'extra': 'forbid',   # 限制 pydantic 模型接收额外的参数
        'json_schema_extra': {  # 添加到该模型输出的 JSON Schema 中的额外信息
            'examples':     # 添加示例数据
        }, 
    }
    
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
    
    create_model(): 动态创建数据模型
        __base__: 配置继承的父类
        __config__: 配置
        __doc__: 模型文档
        __validators__: 属性验证器字典

        def alphanum(cls, v):
            assert v.isalnum(), 'must be alphanumeric'
            return v

        DynamicModel = create_model(
            'DynamicModel',
            foo=(str, Field(alias='FOO')),
            bar=Annotated[str, Field(description='Bar field')],
            baz=(int, 123), # 添加默认值 123
            _private=(int, PrivateAttr(default=1)),
            __base__=FooModel,
            __validators__={
                'foo': field_validator('foo')(alphanum),
            }
        )

    BaseModel: 所有模型校验和序列化操作的基类
    RootModel: 用于校验没有字段名的顶层 json 值, 比如纯列表 [1, 2, 3], 纯字典{'a': 1}, 或任意单一层级的原始值,
        不是用来替代 BaseModel 的通用容器, 而是解决 整个输入就是一个值, 而非键值对集合 的场景
        Pets = RootModel[list[str]]
        print(Pets(['dog', 'cat']).model_dump_json())   # 格式化为 JSON 字符串
        class Pets(RootModel):
            root: list[str]
            def __iter__(self):
                return iter(self.root)

            def __getitem__(self, item):
                return self.root[item]

    AliasPath: 字段别名
    Field: 在 pydantic 模型内部声明校验和元数据
    验证器:
        装饰器验证器:
            # 模型校验装饰器
            @model_validator(mode='before' | 'after' | 'wrap')
            # 字段校验装饰器
            @field_validator(field_name, mode='before' | 'after' | 'wrap' | ‘plain’) 
        注解验证器:
            WrapValidator:  包装验证器, 是所有验证器中最灵活的, 可以在其他验证器处理输入之前或之后运行代码, 也可以通过提前返回值或引发错误来立即终止验证.
            PlainValidator:  普通验证器, 类似于 BeforeValidator, 在返回之后立即终止验证, 因此不会调用其他验证器.
            BeforeValidator: 在内部验证逻辑之前应用验证.
            AfterValidator: 在内部验证逻辑之后应用验证.
                # 自定义校验器, 检查 id 是否以 isbn- 或 imdb- 开始
                def check_valid_id(id: str):
                    if not id.startswith(('isbn-', 'imdb-')):
                        raise raise ValueError('Invalid ID format, it must start with "isbn-" or "imdb-"')
                    return id
                
                @app.get('/items')
                async def read_item(id: Annotated[str | None, AfterValidator(check_valid_id)] = None):
                    if id:
                        # ...
                    else:
                        # ...

    .model_dump(): 返回包含模型数据的字典
        include={}, # 包含的字段的集合
        exclude={}, # 排除的字段的集合
        exclude_unset=True,  # 排除未设置的字段
        exclude_defaults=True,  # 排除默认值的字段
    .model_dump_json()  # 返回数据的 JSON 字符串格式, 参数同 model_dump() 方法
    .model_copy(): 为已有模型创建副本
        update=update_data,  # 使用指定的值更新副本中的数据
    @classmethod
    .model_fields() # 返回模型的字段名称及字段名称的 dict
    @classmethod
    .model_computed_fields()    # 返回模型的计算字段名称及字段名称实例的 dict
    @classmethod
    .model_construct()  # 使用已验证的数据实例化
    @classmethod
    .model_json_schema()    # 为模型生成 JSON 模式
    @classmethod
    .model_rebuild()    # 尝试重建模型 pydantic 核心模式
    @classmethod
    .model_validate()   # 验证模型实例
    @classmethod
    .model_validate_json()  # 验证模型 JSON 字符串
    @classmethod
    .model_validate_strings()   # 使用字符串数据验证给定对象
        
模板: 
    templates = Jinja2Templates(director='templates', context_processors)
        director='',    # 模板目录
        context_processors=[],  # 添加环境处理器
    templates.env.filters['filter_name'] = filter    # 注册自定义过滤器
    
    templates.TemplateResponse()    # 渲染模板
"""


# # pydantic types
# class TestPydanticType(BaseModel):
#     id: UUID4
#     name: str = Field(title="Name of the user", max_length=20)
#     age: int = Field(title="Age of the user", gt=0, lt=200)
#     email: EmailStr = Field(title="Email of the user", pattern=r"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$")
#     pwd: Secret = Field(title="Password of the user", min_length=8, max_length=20)
#     description: str = Field(title="Description of the user")
#     avatar: FileUrl = Field(title="Avatar of the user")
#     blog: HttpUrl = Field(title="Blog of the user")
#     salary: float = Field(title="Salary of the user", gt=0, lt=1000000)
#     currency: Currency = Field(title="Currency of the user")
#     position: Coordinate = Field(title="Position of the user")
#     mac_addr: MacAddress = Field(title="MAC address of the user")
#     zone: TimeZoneName = Field(title="Time zone name of the user")

# # pydantic AliasPath
# class User(BaseModel):
#     # 下标 0 表示取别名字段的第一个值
#     first_name: str = Field(validation_alias=AliasPath('names', 0))
#     # 下标 1 表示取别名字段的第二个值
#     last_name: str = Field(validation_alias=AliasPath('names', 1))
#     address: str = Field(validation_alias=AliasPath('contract', 'address'))
#
# user = User.model_validate({
#     'name': ['John', 'Doe'],
#     'contract': {'address': '123 Main St'}
# })
# print(user)
# # first_name='John' last_name='Doe' address='123 Main St'

# 已废弃
# @app.on_event('startup')
# async def startup_event():
#     async with async_engine.begin() as conn:
#         await conn.run_sync(Base.metadata.create_all)
#


# 异步上下文管理器
@asynccontextmanager
async def start_app(app: FastAPI):
    print('start app...')
    yield
    print('shutdown app...')


app = FastAPI(lifespan=start_app, dependencies=[])


# cookie pydantic 模型
class CookieModel(BaseModel):
    process_time: str


# header pydantic 模型
class HeaderModel(BaseModel):
    user_agent: str
    cache_control: str
    accept: str
    accept_encoding: str
    accept_language: str
    host: str
    connection: str


# 响应模型, 继承 CookieModel 和 HeaderModel, 添加 version 字段
class IndexResponseModel(CookieModel, HeaderModel):
    version: str


# 路径操作装饰器
@app.get('/', response_model=IndexResponseModel, status_code=status.HTTP_200_OK)
async def index(
        cookie_model: Annotated[CookieModel, Cookie()],  # cookie pydantic model
        header_model: Annotated[HeaderModel, Header()],  # header pydantic model
        request: Request,
        response: Response,
) -> Any:
    """
    request: 临时请求对象, 可以访问请求对象的其他信息
    response: 临时响应对象, 可以设置 cookie, header 等信息
        fastapi 将使用这个临时响应来提取cookie和头部信息, 并将它们放入路径操作函数返回的值的最终响应中,
        该响应由任何 response_model 过滤.
    """
    print(f'@app.get("/") cookie_model: {cookie_model}')
    print(f'@app.get("/") header_model: {header_model}')
    cookie_dict = cookie_model.model_dump()
    header_dict = header_model.model_dump()

    print(f'@app.get("/") request.url: {request.url}')
    # 临时响应对象设置 cookie
    response.set_cookie(key='cookie_key', value='cookie_value', domain='/', expires=10 * 60 * 60)

    return {
        **cookie_dict,
        **header_dict,
        'version': fastapi.__version__,
        1: 2  # 返回值会被过滤掉
    }


# Form 参数, 需要显式声明, 否则被解释为查询参数或请求体参数.
class LoginFormModel(BaseModel):
    model_config = {
        'extra': 'forbid',  # 禁止额外的参数字段
    }
    username: str
    password: str


@app.post('/login', status_code=status.HTTP_200_OK)
async def login(form_data: Annotated[LoginFormModel, Form()]):
    print(f'login username: {form_data.username}, password: {form_data.password}')
    return {'username': form_data.username}


# File 参数, 需要显式声明, 否则被解释为查询参数或请求体参数.
# Bytes, 以二进制形式保存在内存中, 适用于小型文件
# UploadFile:
#   无需在参数类型声明中使用 File(),
#   使用 spooled 文件, 文件优先存储在内存中, 当达到最大上限时, 会写入磁盘
#   可以获取上传文件的元数据
#   提供 file-like 的 async 接口
#   暴露一个实际的 SpooledTemporaryFile 对象, 可以直接传递给期望的 file-like 对象的其他库
@app.post('/files', tags=['file'])
async def create_file(file: Annotated[bytes, File()]):
    return {'file_size': len(file)}


@app.post('/uploadfile', tags=['file'])
async def upload_file(
        org: Annotated[str, Form(description="organize code")],
        file: Annotated[UploadFile, File(description="file as UploadFile")]
):
    print(f'org: {org}, files: {file}')
    files = [file]

    file_content = await file.read()  # 读取文件内容
    print(f'read file content is: {file_content}')
    return {
        'org': org,
        'files': [{'file_name': file.filename, 'file_content_type': file.content_type, 'file_size': file.size} for file
                  in files],
    }


# 依赖注入
# async def depend_function(
#         skip: Annotated[int, Query(ge=0, description='must be greater than 0')] = 0,
#         limit: Annotated[int, Query(ge=0, description='must be greater than 0')] = 0,
# ):
#     return {'skip': skip, 'limit': limit}

oauth2_scheme = OAuth2PasswordBearer(tokenUrl="token")


# token 认证
@app.post('/token')
async def get_token(form_data: Annotated[OAuth2PasswordRequestForm, Depends()]):
    access_token = hashlib.sha256(b'zhangsan').hexdigest()
    user_uuid = uuid.uuid4()
    print(f'get_token username: {form_data.username}, password: {form_data.password}')
    print(f'get_token access_token: {access_token}, user_uuid: {user_uuid}')
    return {'access_token': access_token, 'token_type': 'bearer'}


# 类作为依赖项, Depends 参数可省略
# 子依赖项
class CommonsDep:
    def __init__(self, skip: Annotated[int, Query(ge=0, description='must be greater than 0')],
                 limit: Annotated[int, Query(ge=0, description='must be greater than 0')],
                 token: Annotated[str, Depends(oauth2_scheme)]):
        self.skip = skip
        self.limit = limit
        self.token = token


@app.get('/news/list')
async def get_news(commons: Annotated[CommonsDep, Depends()]):
    return commons


@app.get('/blog/list')
async def get_blog(commons: Annotated[CommonsDep, Depends()]):
    return commons


async def verify_token(x_token: Annotated[str, Header()]):
    if x_token != "fake-super-secret-token":
        raise HTTPException(status_code=status.HTTP_400_BAD_REQUEST, detail="Invalid X-Token header")


async def verify_key(x_key: Annotated[str, Header()]):
    if x_key != "fake-super-secret-key":
        raise HTTPException(status_code=status.HTTP_400_BAD_REQUEST, detail="Invalid X-Key header")
    return x_key


# 路径操作装饰器依赖项
@app.get('/items', dependencies=[Depends(verify_token), Depends(verify_key)])
async def read_items():
    return [{"item_id": "foo"}]


"""............................."""


# 异常处理
@app.get('/exception')
async def raise_exception():
    """
    raise 抛出异常, 并添加异常状态码, 异常详情, 响应头
    """
    raise HTTPException(status_code=status.HTTP_400_BAD_REQUEST, detail={'method': 'GET'},
                        headers={'X-Error': 'There goes my error', 'X-Error-Code': '400',
                                 'x-Error-Content': 'Error Content'})


"""............................."""


# deprecated 标识已废弃的接口
@app.get('/fruit/{id}', response_class=JSONResponse, deprecated=True)
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
# 以 JSON 格式读取请求体
# 在必要时把请求体转换为对应的模型
# 校验数据, 数据错误时返回清晰的错误信息, 并指出错误数据的确切位置和内容
class User(BaseModel):
    username: str = Field(min_length=5, max_length=20,
                          description='username minlength 5 and maxlength 20')
    password: str
    email: EmailStr  # 邮箱
    address: Address | None = None


# 装饰器定义响应模型, 返回结果必须符合响应模型定义的值
@app.post('/user', response_model=User)
async def register(user: User):
    print(f'received: {user}')
    u = jsonable_encoder(user)
    print(f'jsonable_encoder(user): {u}')
    u = user.model_dump()
    print(f'.model_dump: {u}')
    return {'id': '001', 'username': 'zhangsan', 'password': '123456', 'email': 'zhangsan@fa.com'}


"""............................."""

stream_data = [{"id": 1, "name": f"name-{1}"},
               {"id": 2, "name": f"name-{2}"},
               {"id": 3, "name": f"name-{3}"},
               {"id": 4, "name": f"name-{4}"}]


# 流式响应 JSON Lines
@app.get('/stream/no-async', tags=['stream'])
def stream_no_async():
    async def generate():
        for item in stream_data:
            yield json.dumps(item) + '\n'

    return StreamingResponse(
        generate(),
        media_type='application/x-ndjson',
        headers={'Content-Disposition': 'inline'})


@app.get('/stream/streaming', response_class=StreamingResponse, tags=['stream'])
async def stream_streaming() -> AsyncIterable[str]:
    message = """
    Rick: (stumbles in drunkenly, and turns on the lights) Morty! You gotta come on. You got--... you gotta come with me.
    Morty: (rubs his eyes) What, Rick? What's going on?
    Rick: I got a surprise for you, Morty.
    Morty: It's the middle of the night. What are you talking about?
    Rick: (spills alcohol on Morty's bed) Come on, I got a surprise for you. (drags Morty by the ankle) Come on, hurry up. (pulls Morty out of his bed and into the hall)
    Morty: Ow! Ow! You're tugging me too hard!
    Rick: We gotta go, gotta get outta here, come on. Got a surprise for you Morty.
    """
    for line in message.splitlines():
        yield line


"""............................."""


# 装饰器定义响应类
@app.get('/html', response_class=HTMLResponse, tags=['template'])
async def get_html():
    return '<h1>HTMLResponse<h1>'


# 返回响应类实例
@app.get('/file/name', tags=['template'])
async def get_file():
    return FileResponse('./file.html')


# 添加环境处理器
def app_context(request: Request):
    return {'context_processors_name': 'c_p_n', 'url': request.url, 'method': request.method, 'cookie': request.cookies,
            'user_agent': request.headers.get('user-agent')}


def marked_filter(text):
    return '<p>--gg--' + text + '--gg--</p>'


# 实例化模板渲染对象
templates = Jinja2Templates(directory='templates', context_processors=[app_context])
# 自定义过滤器
templates.env.filters['marked'] = marked_filter


@app.get('/template/{id}', response_class=HTMLResponse, tags=['template'])
async def get_template(id: int, request: Request):
    return templates.TemplateResponse(request=request, name='index.html', context={'id': id})


"""............................."""


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
    start_time = time.perf_counter()
    response = await call_next(request)
    process_time = time.perf_counter() - start_time
    response.headers['X-Process-Time'] = str(process_time)
    response.set_cookie('process_time', str(process_time), max_age=24 * 60 * 60, path="/")
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

# 挂载静态文件
app.mount('/static', StaticFiles(directory='static'), name='static')

# 注册路由
app.include_router(news.router, prefix='/api')  # 全局路由前缀

if __name__ == '__main__':
    uvicorn.run(app, port=8000)
