from fastapi import APIRouter
from starlette.responses import HTMLResponse

router = APIRouter(prefix="/api/news", tags=["news"])


@router.get('/', response_class=HTMLResponse, summary="获取新闻列表", description="获取新闻列表",
            response_description="返回新闻列表")
async def list():
    return 'hello world'
