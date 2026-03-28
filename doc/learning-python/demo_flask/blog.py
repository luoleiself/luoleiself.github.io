from flask import Blueprint, render_template, request

bp = Blueprint('blog', __name__, url_defaults={'username': 'wanglaowu'})


@bp.route('/')
@bp.endpoint('index')   # 关联端点和路由
def index():
    return render_template('blog.html')


# 自定义 endpoint
@bp.route('/create', endpoint='create_blog')
def create(username):
    print('create blueprint', username)
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
