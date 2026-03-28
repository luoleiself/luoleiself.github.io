from flask import Blueprint, request, session, redirect, url_for, flash


bp = Blueprint('auth', __name__, url_prefix='/auth')


@bp.route('/register', methods=('GET', 'POST'))
def register():
    return {"code": 200, 'msg': 'success', 'data': {}}
