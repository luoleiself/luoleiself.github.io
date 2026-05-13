from flask import Blueprint, request, session, redirect, url_for, flash, render_template, get_flashed_messages

bp = Blueprint(
'login',
    __name__,
    template_folder="templates",    # 蓝图的模板目录
    static_folder="static", # 蓝图的资源目录
    url_prefix='/login'  # url 前缀
)


@bp.route('/register', methods=('GET', 'POST'))
def register():
    msg = get_flashed_messages()
    print('flashed messages', msg)
    if request.method == 'GET':
        return render_template('login.html')
    elif request.method == 'POST':
        return {"code": 200, 'msg': 'success', 'data': {}}
    else:
        return "register good", 200


@bp.errorhandler(404)
def page_not_found(e):
    return "Not Found", 404