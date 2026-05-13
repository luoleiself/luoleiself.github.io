import jwt
from datetime import datetime, timedelta, timezone
from typing import Optional

# openssl rand -hex 32
SECRET = 'c9dee47fc2c4e1d5f37015f4b8cdd5ea6adcd12d3b9e50630ddc96eccf44ae86'
ALGORITHM = 'HS256'
ACCESS_TOKEN_EXPIRE_MINUTES = 24 * 60


def create_access_token(data: dict, expires_delta: Optional[timedelta] = None) -> str:
    """
    创建 JWT token

    Args:
        data: 要编码的数据字典
        expires_delta: 可选的过期时间增量，如果不提供则使用默认值

    Returns:
        编码后的 JWT token 字符串
    """
    to_encode = data.copy()
    now = datetime.now(timezone.utc)
    if expires_delta:
        expire = now + expires_delta
    else:
        expire = now + timedelta(minutes=ACCESS_TOKEN_EXPIRE_MINUTES)

    to_encode.update({'iat': now, "exp": expire})
    encoded_jwt = jwt.encode(to_encode, SECRET, algorithm=ALGORITHM, headers={"typ": "JWT", "alg": ALGORITHM})

    return encoded_jwt


def decode_token(token: str) -> dict:
    """
    解码并验证 JWT token

    Args:
        token: JWT token 字符串

    Returns:
        解码后的数据字典

    Raises:
        jwt.ExpiredSignatureError: token 已过期
        jwt.InvalidTokenError: token 无效
    """
    try:
        payload = jwt.decode(token, SECRET, algorithms=[ALGORITHM])
        return payload
    except jwt.ExpiredSignatureError:
        raise Exception("Token 已过期")
    except jwt.InvalidTokenError:
        raise Exception("无效的 Token")


def verify_token(token: str) -> bool:
    """
    验证 JWT token 是否有效

    Args:
        token: JWT token 字符串

    Returns:
        True 如果 token 有效，否则 False
    """
    try:
        decode_token(token)
        return True
    except Exception:
        return False


# 使用示例
if __name__ == "__main__":
    # 1. 创建 token
    user_data = {
        "user_id": 123,
        "username": "test_user",
        "email": "test@example.com"
    }

    token = create_access_token(data=user_data)
    print(f"生成的 Token: {token}")

    # 2. 解码 token
    try:
        decoded_data = decode_token(token)
        print(f"解码后的数据: {decoded_data}")
    except Exception as e:
        print(f"解码失败: {e}")

    # 3. 验证 token
    is_valid = verify_token(token)
    print(f"Token 是否有效: {is_valid}")

    # 4. 测试过期 token
    expired_token = create_access_token(
        data=user_data,
        expires_delta=timedelta(seconds=-1)  # 已过期的 token
    )
    print(f"\n测试过期 Token:")
    print(f"Token 是否有效: {verify_token(expired_token)}")

    try:
        decode_token(expired_token)
    except Exception as e:
        print(f"解码过期 Token 错误: {e}")

    # 5. 测试无效 token
    invalid_token = "invalid.token.here"
    print(f"\n测试无效 Token:")
    print(f"Token 是否有效: {verify_token(invalid_token)}")
