from flask import Blueprint, render_template, request, current_app, g, redirect, url_for, flash

from demo_flask.db import get_db

bp = Blueprint('blog', __name__)


@bp.route('/')
@bp.endpoint('index')  # 关联端点和路由
def index():
    print('g.test_name', getattr(g, 'test_name', None))
    db = get_db()
    cursor = db.cursor()
    result = cursor.execute('select * from users')
    users = result.fetchall()
    for user in users:
        print(f'id: {user[0]} username: {user[1]} password: {user[2]}')

    return render_template('blog.html', users=users)


# 自定义 endpoint
@bp.route('/create', methods=('GET', 'POST'), endpoint='create_blog')
def create():
    if request.method == 'GET':
        return render_template('create_blog.html')
    elif request.method == 'POST':
        username = request.form['username']
        password = request.form['password']
        if not username or not password:
            raise ValueError('username or password is required')

        db = get_db()
        cursor = db.cursor()
        res = cursor.execute('select * from users where username = ?', (username,))
        one = res.fetchone()
        if one is not None:
            flash(f'username: {username} already exists')
            return render_template('create_blog.html')

        res = cursor.execute('insert into users(username, password) values (?, ?)', (username, password))
        db.commit()
        print(f'res.rowcount {res.rowcount} res.lastrowid {res.lastrowid}')
        return redirect(url_for('blog.index'))
    else:
        return 'Not found', 404


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
