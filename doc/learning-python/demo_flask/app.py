import time

from markupsafe import escape
from flask import Flask, url_for, request, render_template, make_response, redirect, session, flash, \
    render_template_string, g
import string

from demo_flask.login import login
from demo_flask import blog
from demo_flask import comment
from demo_flask import db

'''
Flask: WSGI 应用, 一个 Flask 应用就是一个 Flask 类的实例, 内置 http 服务器
    Flask 自动添加一个 static 视图(函数), 用于提供静态文件, 如 /static/style.css

应用上下文: 在整个应用生命周期中存在
请求上下文: 在每个请求生命周期中存在
current_app: 是一个线程安全的代理对象, 指向当前处理请求的 Flask 应用实例, 可以在没有直接访问应用对象的地方使用应用配置和资源
g: 一个线程安全的请求级别的全局对象, 用于在同一请求的多个函数之间共享数据, 在每个请求开始时创建, 请求结束时销毁

动态路由参数 <type:variable_name>, 可指定类型 int, float, path, uuid, 默认为 string
url_for() 函数用于构建指定函数的 url, 如 url_for('user_id', name='Tome', age=18) => /user/18?name=Tome
    第一个参数为函数名字符串表示, 如果未定义 endpoint 则默认使用视图函数名
    剩余参数为可变关键字参数对应到 url 中的变量, 未知变量将添加到 url 中作为查询参数

render_template() 使用模板文件渲染, 默认在 templates 目录下查找
render_template_string() 使用字符串模板渲染
abort() 用于中止请求处理
redirect() 用于重定向
make_response() 用于创建一个响应对象

request 请求对象, 包含了 headers, method, cookies, args, form, files, json 等属性

响应:
    1. 如果返回一个响应对象, 就直接返回它
    2. 如果返回一个字符串, 根据这个字符串和缺省参数生成一个用于返回的响应对象
    3. 如果返回一个迭代器或生成器, 那么返回字符串或字节, 作为流响应对待
    4. 如果返回一个字典或列表, 将自动使用 jsonify() 创建一个响应对象
    5. 如果返回一个元组, 那么根据元组中的项可以提供额外的信息, 元组中必须至少包含一个项目, 
        且项目应当由 response, status 或 header, 或三者组成,
        status 会重载状态码, headers 是一个由额外头部值组成的列表或字典
    6. 如果以上都不是, 那么 Flask 会假定返回值是一个有效的 WSGI 应用并把它转换为一个响应对象

会话: session, 可以在不同请求之间存储共享信息

消息闪现(flash): 在请求结束时记录一个消息, 提供且只提供给下一个请求使用, 通常使用布局模板来展现闪现的消息
    flash() 用于闪现消息
    get_flashed_messages() 用于获取闪现消息

日志: app.logger 对象基于 logging 模块

视图是一个应用对请求进行响应的函数.
蓝图: blueprint 是一种组织一组相关视图及其它代码的方式, 与把视图及其它代码直接注册到应用的方式不同,
    蓝图是把它们注册到蓝图, 然后在把蓝图注册到应用.
    蓝图实例的名称作为视图函数的端点前缀, 蓝图的名称不修改 url, 只修改端点
    url_for('login.static', filename='login.css')    # 解析自定义的资源目录
    
端点(endpoint): 内部用于反向生成 URL 的标识符, 每个视图函数默认使用函数名作为 endpoint, 也可以自定义
    @app.route('/user/<name>', endpoint='user_profile')
    def show_user(name):
        # 使用 url_for 生成 URL
        profile_url = url_for('user_profile', name='wanglaowu')
        return f'用户页面url: {profile_url}'
    
    # 蓝图自定义的 endpoint 生成 URL
    url_for('blog.create_blog', name='hash', topic='new', tags=('plus', 'plug'))
子域名(subdomain): 根据域名前缀路由到不同的(蓝图)视图函数, 实现多租户、多模块应用, 需配置 SERVER_NAME
    app.config['SERVER_NAME'] = 'api.example.com'
    @app.route('/users', subdomain="v1")
    def users():
        return {'data': 'API v1'}
        
    api_url = url_for('users')
    # 生成: https://v1.api.example.com/users
    
    
类视图(继承 View): 使用类定义视图函数, 重写 dispatch_request 方法处理请求
    methods 属性限制请求方法
    dispatch_request() 等同于视图函数, 动态参数将作为关键字参数传入此方法
    View.as_view()
        第一个参数是用于 url_for 的指向视图的名称
        创建视图函数, 任何传给 as_view 函数除第一个参数以外的其它参数都会在创建类时传递

方法视图(继承 MethodView): 基于类定义的请求方法自动设置 View.methods 


内置配置变量:
    SECRET_KEY: 用于会话管理的密钥
    SESSION_COOKIE_NAME: 会话 cookie 的名称
    SESSION_COOKIE_PATH: 会话 cookie 的路径
    SESSION_COOKIE_DOMAIN: 会话 cookie 的域名
    SESSION_COOKIE_SECURE: 会话 cookie 是否仅通过 https 访问
    SESSION_COOKIE_HTTPONLY: 会话 cookie 是否仅通过 http 访问
    SESSION_COOKIE_SAMESITE: 会话 cookie 的 SameSite 属性
    PERMANENT_SESSION_LIFETIME: 永久会话的生命周期
    SERVER_NAME: 服务器的主机名和端口
    APPLICATION_ROOT: 应用的根路径
    EXPLAIN_TEMPLATE_LOADING: 是否启用输出 render_template 渲染模板文件时的查找路径

特殊对象和函数:
    app.config  应用全局配置
    app.test_request_context() 用于创建一个测试请求上下文器
    app.app_context() 创建一个应用上下文管理器
    app.register_blueprint() 注册蓝图
    app.add_url_rule() 注册路由

装饰器:
    @app.before_request  # 在所有请求之前执行
    @app.after_request(response)  # 在所有请求之后执行
    @app.teardown_request   # 在退出请求情境时执行, 即使发生未处理的异常也会执行
    @app.teardown_appcontext    # 在退出应用情境时执行, 即使发生未处理的异常也会执行
    @app.endpoint('index')  # 关联端点名称 index 和 URL, 覆盖外层相同路由
    @app.errorhandler(404)  # 注册错误处理函数
    @app.template_filter    # 自定义模板过滤器
    @app.context_processor  # 自定义环境处理器
'''
'''
jinja2

{% ... %} 用于模板语句
{{ ... }} 用于输出表达式
{# ... #} 用于模板注释
{% autoescape false %}
    临时关闭自动转义
{% endautoescape %}

自定义过滤器: 使用注册的名字 upper
@app.template_filter('upper')
def my_filter(value):
    return value.upper()
环境处理器: 在模板渲染前运行, 自动将新的变量引入模板变量中
@app.context_processor
def context_processor():
    return dict(diy_name='variable', diy_flag='Flask')
'''

app = Flask(__name__)
app.config['SECRET_KEY'] = string.ascii_letters + \
                           string.digits + string.punctuation


@app.route('/', methods=('GET',))
def index():
    app.logger.info('index')
    return 'Hello, World!'


@app.route('/hi')
def hi():
    app.logger.warning('hi hi hi')
    return render_template_string('render_template_string request.method: {{method}}', method=request.method)


@app.route('/user/<username>', methods=('GET', 'POST'))
def user(username):
    """动态路由参数, 支持参数指定类型"""
    app.logger.info('user is %s', username)
    session['username'] = username

    # result = g.db.execute('SELECT * FROM users')
    # print(result)

    if request.method == 'POST':
        print(f'request.method: {request.method}')
        return f'request.method: {request.method}'
    else:
        return f'''<h1>动态路由参数: {username}</h1>
        <h3>参数可指定类型, 将匹配指定的路由</h3>
        <ul> 
            <li>/user/<username>: 缺省类型为 string, 匹配所有类型参数</li>
            <li>/user/<int:user_id>: 匹配整数</li>
            <li>/user/<float:user_id>: 匹配浮点数</li>
            <li>/user/<path:user_id>: 匹配路径, 可以包含斜线 </li>
            <li>/user/<uuid:user_id>: 匹配 UUID</li>
        </ul>
        '''


@app.route('/user/<int:uid>')
def user_id(uid):
    """动态路由参数, 格式: <int:user_id>, 类型为 int"""
    return f'动态路由参数, 格式为 {escape('<int:user_id>')}: {uid}</h1>'


@app.route('/template/<name>')
def template(name):
    res = make_response(render_template('template1.html', name=name))
    res.headers['Access-Control-Allow-Origin'] = '*'
    res.headers['Access-Control-Allow-Methods'] = '*'
    res.headers['Access-Control-Allow-Headers'] = '*'
    res.set_cookie('name', 'Flask', path='/',
                   max_age=7 * 24 * 60 * 60)
    print(f'request.method: {request.method}')
    print(f'request.cookies: {request.cookies}')
    print(f'request.args: {request.args}')
    return res


@app.route('/me')
def me():
    """
    返回字典或列表时, 自动使用 jsonify() 转换为 JSON 响应
    """
    return {
        "userName": session.get('username', 'hangman'),
        "userEmail": 'zhangsan@example.com'
    }


@app.route('/mine')
def mine():
    # 闪现消息
    flash('From /mine')
    flash('From /mine again')
    # 重定向
    return redirect(url_for('login.register'))


# 注册错误处理函数
@app.errorhandler(404)
def handle_404(error):
    print('error', error)
    return '404 Not Found', 404


# 自定义过滤器
@app.template_filter('upper')
def my_filter(value):
    return value.upper()


# 注册环境处理器
@app.context_processor
def context_processor():
    return dict(diy_name='hello flask', diy_flag='wadaxi')


# 在所有请求之前执行, 如果有返回值则不再调用视图函数
@app.before_request
def before_request():
    g.start_time = time.perf_counter()
    g.test_name = 'hello flask test'
    print('app before request...')


# 在所有请求之后执行, 返回 response
@app.after_request
def after_request(response):
    print('app after request...')
    end_time = time.perf_counter()
    start_time = g.start_time
    print(f'request time: {end_time - start_time}')
    return response


# 在退出请求情境时执行, 忽略未处理的异常
@app.teardown_request
def teardown_request(exception):
    print('teardown_request exception', exception)


# 在退出应用情境时执行, 忽略未处理的异常
@app.teardown_appcontext
def teardown_appcontext(exception):
    print('teardown_appcontext exception', exception)


class User:
    name = 'Zhang San'
    age = 18


print('__name__', __name__)
# 注册蓝图
app.register_blueprint(login.bp)
app.register_blueprint(blog.bp)
# 使用 @bp.endpoint('index') 装饰器, 关联端点名称 index 和 URL
# url_for('index') 或 url_for('blog.index') 都会有效
# app.add_url_rule('/', endpoint='index')

app.register_blueprint(comment.bp)
# as_view() 第一个参数是用于 url_for 的指向视图的名称
# 剩余参数在创建类时传递
app.add_url_rule(
    '/comment/<int:id>',
    view_func=comment.Comment.as_view('comment', User, 'comment.html')
)

# 测试请求上下文
with app.test_request_context():
    """测试请求上下文"""
    print('-' * 10)
    print(url_for('index'))
    print(url_for('hi', name='Flask', age=18,
                  addr='Shanghai', email='flask@example.com'))
    print(url_for('user', username='Flask', age=18, addr='Beijing'))
    print(url_for('user_id', name='Jerry', uid=18))
    print('-' * 10)
    print('蓝图 url_prefix: ', url_for('login.register'))
    print('蓝图自定义 endpoint 生成 URL', url_for('blog.create_blog', _method='POST',
                                                  username='hash', topic='new', tags=('hehe', 'haha')))
    print('-' * 10)

# 测试指定请求
with app.test_request_context('/hi'):
    print('test request hi...', request.method, request.path, request.args)
    assert request.method == 'GET'

# 手动创建应用上下文管理器
with app.app_context():
    db.init_app(app)

if __name__ == '__main__':
    app.run(debug=True)
