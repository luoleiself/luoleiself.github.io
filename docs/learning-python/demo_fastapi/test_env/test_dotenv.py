from pathlib import PurePath
from dotenv import load_dotenv, dotenv_values, get_key
import os

from pydantic import Field
from pydantic_settings import BaseSettings

# 加载解析 .env 环境变量
load_dotenv()


def test_load_dotenv():
    """
    使用 python-dotenv 管理环境变量
    load_dotenv()   # 加载解析 .env 环境变量
    dotenv_values() # 加载解析 .env 环境变量作为 dict 返回
    get_key()   # 从 .env 环境变量中获取指定键
    """
    port = os.getenv('PORT')
    db_url = os.getenv('DATABASE_URL')
    secret_key = os.getenv('SECRET_KEY')
    debug = os.getenv('DEBUG')

    print(f'\nport {port}')
    print(f'db_url {db_url}')
    print(f'secret_key {secret_key}')
    print(f'debug {debug}')

    assert port is not None, 'PORT is None'
    assert db_url is not None, 'DB_URL is None'
    assert secret_key == 'secret-key', 'SECRET_KEY is not secret-key'
    assert True if debug else False, 'DEBUG is not True'
    print('-' * 10)

    env_config = dotenv_values()
    print(f'env_config {env_config}')
    print('-' * 10)

    env_file_path = PurePath(__file__).parent / '.env'
    port = get_key(env_file_path, 'PORT')
    print(f'port {port}')
    db_url = get_key(env_file_path, 'DATABASE_URL')
    print(f'db_url {db_url}')


class Settings(BaseSettings):
    port: int = Field(..., ge=1024, lt=65535)
    database_url: str
    secret_key: str
    debug: bool = Field(default=True)

    class Config:
        env_file = PurePath(__file__).parent / '.env'  # 自动加载 .env
        # env_prefix = 'APP_'


settings = Settings()


def test_pydantic_env():
    """
    使用 pydantic-settings 管理环境变量
    """
    print(f'settings.port {settings.port}')
    print(f'settings.database_url {settings.database_url}')
    print(f'settings.secret_key {settings.secret_key}')
    print(f'settings.debug {settings.debug}')
