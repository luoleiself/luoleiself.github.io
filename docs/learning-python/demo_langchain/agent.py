from langchain.agents import create_agent
from langchain.chat_models import init_chat_model
from langchain_core.messages import SystemMessage, AIMessage
from langchain_core.prompts import ChatPromptTemplate, MessagesPlaceholder
from langchain_core.tools import tool
from langchain_openai import ChatOpenAI

"""
提示词:
    1. 给大模型设定角色和能力
    2. 明确核心请求和任务
    3. 按步骤拆解复杂任务
    4. 指定风格与语气
    5. 明确要求输出格式
    6. 提供输入输出的示例

LangChain:
    - Prompts   # 提示词工程
    - Models    # 调用各类模型
    - History   # 管理会话历史记录
    - Indexes   # 管理和分析各类文档
    - Chains    # 构建功能的执行链条
    - Agent # 构建智能体

核心组件:
    - create_agent(): 创建智能体
    - init_chat_model(): 创建模型

调用方式:
    - llm.invoke() / llm.ainvoke()          # 普通调用
    - llm.stream() / llm.astream()          # 流式调用
    - llm.batch() / llm.abatch()            # 批量调用,顺序返回
    - llm.batch_as_completed()              # 批量调用,乱序返回

工具绑定:
    - llm.bind_tools([tool])                # 绑定工具
    - llm.with_structured_output(schema)    # 结构化输出
    - llm.with_retry()                      # 重试机制

工具定义:
    - @tool 装饰器定义工具函数
"""


@tool
def get_weather(city: str) -> str:
    """
    查询指定城市的天气信息。

    Args:
        city: 城市名称,如 "北京", "上海"

    Returns:
        该城市的天气描述
    """
    weather_data = {
        "北京": "晴天,温度 25°C",
        "上海": "多云,温度 28°C",
        "广州": "小雨,温度 30°C",
        "深圳": "晴天,温度 29°C",
    }
    return weather_data.get(city, f"未找到 {city} 的天气信息")


@tool
def calculate(expression: str) -> str:
    """
    执行数学计算。

    Args:
        expression: 数学表达式,如 "2 + 3 * 4"

    Returns:
        计算结果
    """
    try:
        result = eval(expression, {"__builtins__": {}}, {})
        return f"计算结果: {result}"
    except Exception as e:
        return f"计算错误: {str(e)}"


@tool
def get_time_info(query: str) -> str:
    """
    获取当前时间信息。

    Args:
        query: 时间查询类型,如 "current_time", "date"

    Returns:
        时间信息字符串
    """
    from datetime import datetime
    now = datetime.now()

    if "date" in query.lower():
        return f"当前日期: {now.strftime('%Y年%m月%d日')}"
    elif "time" in query.lower():
        return f"当前时间: {now.strftime('%H:%M:%S')}"
    else:
        return f"当前日期时间: {now.strftime('%Y年%m月%d日 %H:%M:%S')}"


# 初始化模型并绑定工具
llm = ChatOpenAI(model="gpt-3.5-turbo", temperature=0, api_key=None)
model_with_tools = llm.bind_tools([get_weather, calculate, get_time_info])

# 测试工具调用
print("=" * 50)
print("测试工具调用:")
print("=" * 50)

response = model_with_tools.invoke("北京的天气怎么样?")
print(f"问题: 北京的天气怎么样?")
print(f"回答: {response.content}")
if response.tool_calls:
    print(f"工具调用: {response.tool_calls}")

print()

response = model_with_tools.invoke("计算 2 + 3 * 4")
print(f"问题: 计算 2 + 3 * 4")
print(f"回答: {response.content}")
if response.tool_calls:
    print(f"工具调用: {response.tool_calls}")

print()

response = model_with_tools.invoke("现在几点了?")
print(f"问题: 现在几点了?")
print(f"回答: {response.content}")
if response.tool_calls:
    print(f"工具调用: {response.tool_calls}")

# 创建提示词模板
print("\n" + "=" * 50)
print("提示词模板:")
print("=" * 50)

prompt = ChatPromptTemplate.from_messages([
    ("system", "你是一个智能助手,可以查询天气、执行计算、提供时间信息。请根据用户的问题选择合适的工具。"),
    MessagesPlaceholder("chat_history", optional=True),
    ("human", "{input}"),
])

print(f"提示词模板结构:\n{prompt}")

# 使用提示词模板
print("\n" + "=" * 50)
print("使用提示词模板:")
print("=" * 50)

formatted_prompt = prompt.format_messages(
    input="帮我计算一下 100 除以 3 的结果"
)
print(f"格式化后的提示词: {formatted_prompt}")

# 可选: 创建完整的智能体 (需要取消注释)
"""
agent = create_agent(
    model=init_chat_model("gpt-3.5-turbo", temperature=0),
    tools=[get_weather, calculate, get_time_info],
    system_prompt="你是一个智能助手,可以查询天气、执行计算、提供时间信息。",
)

result = agent.invoke({"messages": [{"role": "user", "content": "北京天气如何?"}]})
print(result)
"""
