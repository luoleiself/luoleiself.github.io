from markupsafe import escape
from flask import Flask, url_for, request, render_template, make_response, redirect, session, flash, \
    get_flashed_messages, render_template_string
import string

import auth
import blog
import comment

'''
Flask: WSGI 应用, 一个 Flask 应用就是一个 Flask 类的实例, 内置 http 服务器
    Flask 自动添加一个 static 视图(函数), 用于提供静态文件, 如 /static/style.css

动态路由参数 <type:variable_name>, 可指定类型 int, float, path, uuid, 默认为 string
url_for() 函数用于构建指定函数的 url, 如 url_for('user_id', name='Tome', age=18) => /user/18?name=Tome
    第一个参数为函数名字符串表示
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

特殊对象和函数:
    app.config  应用全局配置
    app.test_request_context() 用于创建一个测试请求上下文
    app.teardown_request() 用于注册一个在请求结束时执行的函数
    app.register_blueprint() 用于注册蓝图
    app.before_request()

    g: 特殊全局对象, 独立于每一个请求, 在处理请求过程中可以用于存储可能多个函数都会用到的数据
    current_app: 当前应用的代理对象, 可以在应用上下文中使用

特殊装饰器:
    @app.endpoint('index')  # 关联端点名称 index 和 URL, 覆盖外层相同路由
    @app.errorhandler(404)  # 注册错误处理函数
    @app.before_request  # 在所有请求之前执行
    @app.after_request  # 在所有请求之后执行
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
    return dict(name='variable', flag='Flask')
'''

app = Flask(__name__)
app.config['SECRET_KEY'] = string.ascii_letters + \
    string.digits + string.punctuation
app.config['DATABASE'] = './db.sqlite'


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


@app.route('/user/<int:user_id>')
def user_id(user_id):
    """动态路由参数, 格式: <int:user_id>, 类型为 int"""
    return f'动态路由参数, 格式为 {escape('<int:user_id>')}: {user_id}</h1>'


@app.route('/template/<name>')
def template(name):
    res = make_response(render_template('template1.html', name=name))
    res.headers['Access-Control-Allow-Origin'] = '*'
    res.headers['Access-Control-Allow-Methods'] = '*'
    res.headers['Access-Control-Allow-Headers'] = '*'
    res.set_cookie('name', 'Flask', path='/',
                   max_age=7*24*60*60)
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
    return redirect(url_for('login'))


@app.route('/login')
def login():
    # 中止请求处理
    # abort(401)
    # 获取上一个请求闪现的消息
    msg = get_flashed_messages()
    print('msg', msg)
    return '401 Unauthorized {}'.format(msg), 401


@app.errorhandler(404)
def handle_404(error):
    print('error', error)
    return '404 Not Found', 404


# 自定义过滤器
@app.template_filter('upper')
def my_filter(value):
    return value.upper()


# 在所有请求之前执行
@app.before_request
def before_request():
    print('before_request', request.method, request.path)


@app.after_request
def after_request(response):
    print('after_request', request.method, request.path)
    return response


@app.teardown_request
def tear_down_request(response):
    print('teardown_request', request.method, request.path)
    return response


class User:
    name = 'Zhang San'
    age = 18


if __name__ == '__main__':
    # 注册蓝图
    app.register_blueprint(auth.bp)
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
        print('--------------')
        print(url_for('index'))
        print(url_for('hi', name='Flask', age=18,
                      addr='Shanghai', email='flask@example.com'))
        print(url_for('user', username='Flask', age=18, addr='Beijing'))
        print(url_for('user_id', name='Jerry', user_id=18))
        print('--------------')
        print('蓝图自定义 endpoint 生成 URL', url_for('blog.create_blog',
              name='hash', topic='new', tags=('hehe', 'haha')))
        print('--------------')

    # 测试指定请求
    with app.test_request_context('/hi'):
        print('test request hi...', request.method, request.path, request.args)
        assert request.method == 'GET'

    app.run(debug=True)

    import db
    db.init_app(app)
