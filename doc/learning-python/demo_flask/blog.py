from flask import Blueprint, render_template, request, current_app, g

from demo_flask.db import get_db

bp = Blueprint('blog', __name__)


@bp.route('/')
@bp.endpoint('index')   # 关联端点和路由
def index():
    print('g.test_name', getattr(g, 'test_name', None))
    db = get_db()
    cursor = db.cursor()
    result = cursor.execute('select * from users')
    user = result.fetchall()
    print('user', user)
    return render_template('blog.html', user=user)


# 自定义 endpoint
@bp.route('/create/<username>', endpoint='create_blog')
def create(username):
    print('blog bp create', username)
    # cursor = get_db().cursor()
    # result = cursor.execute('insert into users (username, password) values (?, ?)', (username, '123456'))
    # print('result', result)
    # print('result.fetchall()', result.fetchall())
    cursor = get_db().execute('insert into users ("username", "password") values (?, ?)', (username, '123456'))
    rv = cursor.fetchall()
    cursor.close()
    print('rv', rv)
    return 'Create a new blog post'


# 多路由视图函数
@bp.route('/info/<username>', endpoint='blog_info')
@bp.route('/profile/<username>', endpoint='blog_profile')
def show_user(username):
    if request.endpoint == 'blog_info':
        return f'request.endpoint: blog_info: {username}'
    elif request.endpoint == 'blog_profile':
        return f'request.endpoint: blog_profile: {username}'
    else:
        return 'Not found', 404


# blog 蓝图请求之前执行
@bp.before_request
def bp_before_request():
    print('blog blueprint before request...')
    pass


# blog 蓝图请求之后执行
@bp.after_request
def bp_after_request(response):
    print('blog blueprint after request...')
    return response


# blog 蓝图退出请求情境时执行, 忽略未处理的异常
@bp.teardown_request
def bp_teardown_request(exception):
    print('blog blueprint teardown request...')

