from fastapi.testclient import TestClient
from .app import app

# 创建 TestClient 实例
client = TestClient(app)


def test_read_main():
    response = client.get("/html")
    assert response.status_code == 200
    assert response.content == b"<h1>HTMLResponse<h1>", '响应内容不符合要求'
