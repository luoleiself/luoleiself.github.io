import json
from decimal import Decimal
from datetime import datetime, date

"""
json 模块默认只能序列化基本类型（dict、list、str、int、float、bool、None），遇到自定义对象或特殊类型时会报错
"""

dt = {'a': 'A', 'c': 1, 'd': True, 'f': '张三'}
print(f'ascii 编码 json.dumps(dt): {json.dumps(dt)}')
print(
    f'ascii 不编码 json.dumps(dt, ensure_ascii=False): {json.dumps(dt, ensure_ascii=False)}')
print('-' * 10)
print('''
# 自定义编码类继承 json.JSONEncoder, 关键字参数 cls 传入 json.dumps
class CustomJSONEncoder(json.JSONEncoder):
    def default(self, obj):
        if isinstance(obj, datetime):
            return obj.isoformat()

        if isinstance(obj, date):
            return obj.isoformat()

        if isinstance(obj, Decimal):
            return float(obj)

        if hasattr(obj, '__dict__'):
            # 递归处理 __dict__ 中的属性
            return {
                key: value
                for key, value in obj.__dict__.items()
                if not key.startswith('_')  # 过滤私有属性
            }

        return super().default(obj)


# 自定义编码函数(简单场景), 关键字参数 default 传入 json.dumps
def customize_json(obj):
    if isinstance(obj, datetime):
        return obj.isoformat()

    if isinstance(obj, date):
        return obj.isoformat()

    if isinstance(obj, Decimal):
        return float(obj)

    if hasattr(obj, '__dict__'):
        return obj.__dict__

    raise TypeError(f'Object of type {type(obj)} is not JSON serializable')
''')


class CustomJSONEncoder(json.JSONEncoder):
    def default(self, obj):
        if isinstance(obj, datetime):
            return obj.isoformat()

        if isinstance(obj, date):
            return obj.isoformat()

        if isinstance(obj, Decimal):
            return float(obj)

        if hasattr(obj, '__dict__'):
            # 递归处理 __dict__ 中的属性
            return {
                key: value
                for key, value in obj.__dict__.items()
                if not key.startswith('_')  # 过滤私有属性
            }

        return super().default(obj)


def customize_json(obj):
    if isinstance(obj, datetime):
        return obj.isoformat()

    if isinstance(obj, date):
        return obj.isoformat()

    if isinstance(obj, Decimal):
        return float(obj)

    if hasattr(obj, '__dict__'):
        return obj.__dict__

    raise TypeError(f'Object of type {type(obj)} is not JSON serializable')


class User:
    def __init__(self, name, age, price, birthday):
        self.name = name
        self.age = age
        self.price = price
        self.birthday = birthday


org_data = {
    'name': '张三',
    'age': 18,
    'price': Decimal(18.88),
    'birthday': datetime.now(),
    'user': User('张三', 18, Decimal(18.88), datetime.now())
}
json_str = json.dumps(org_data, cls=CustomJSONEncoder)
print(f'自定义编码类, 传入关键字参数 cls:\n {json_str}')
json_str = json.dumps(org_data, default=customize_json)
print(f'自定义编码函数, 传入关键字参数 default:\n {json_str}')
print('-' * 10)
