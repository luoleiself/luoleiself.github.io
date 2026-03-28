from flask import Blueprint, render_template
from flask.views import View, MethodView
'''
类视图: 不能使用 @app.route() 装饰器装饰类
    使用类继承 View 类定义视图函数, 重写 dispatch_request 方法处理请求

# 添加 URL 规则, 绑定视图函数
    [View].as_view()
        函数创建视图函数, 任何传给 as_view 函数第一个参数以外的其它参数都会在创建类时传递
        
app.add_url_rule('/comment/<int:id>', view_func=Comment.as_view('comment'))  
'''


bp = Blueprint('comment', __name__)


class Comment(View):
    # 使用 methods 属性限制请求方法
    # 等价于在 add_url_rule() 中使用 methods 参数传入
    methods = ['GET', 'POST']

    def __init__(self, model, template):
        self.model = model
        self.template = template

    # 等同于视图函数
    def dispatch_request(self, id):
        # if request.method == 'POST':
        #     return self.create_comment()
        # elif request.method == 'GET':
        #     return self.get_comments()

        # 查询结果
        # result = self.model.query.all()
        result = {}

        return render_template(self.template, result=result)


# MethodView 扩展 View 类, 根据请求方法调用类的不同方法
# 动态参数将作为关键字参数传入此方法
class CommentView(MethodView):
    def get(self, id):
        return 'GET request'

    def post(self, id):
        return 'POST request'

    def put(self, id):
        return 'PUT request'

    def patch(self, id):
        return 'PATCH request'

    def delete(self, id):
        return 'DELETE request'
