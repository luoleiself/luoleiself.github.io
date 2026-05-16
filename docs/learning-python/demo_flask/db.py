import sqlite3
from pathlib import PurePath

from flask import current_app, g
from datetime import datetime


def get_db():
    db = getattr(g, '_database', None)
    if db is None:
        database_path = PurePath(__file__).parent / 'database.db'
        print('database_path', database_path)
        db = g._database = sqlite3.connect(database_path)
        db.row_factory = sqlite3.Row
    return db


def close_db(e=None):
    db = getattr(g, '_database', None)
    if db is not None:
        db.close()


def init_app(app):
    print('init_db...', PurePath(__file__))
    app.teardown_appcontext(close_db)
    with app.app_context():
        db = get_db()
        with app.open_resource('schema.sql', mode='r') as f:
            db.cursor().executescript(f.read())

        db.commit()


sqlite3.register_converter(
    'datetime',
    lambda v: datetime.fromisoformat(v.decode())
)
