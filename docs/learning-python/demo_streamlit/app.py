import streamlit as st
from openai import OpenAI
from datetime import datetime
from pathlib import Path, PurePath
import json

# streamlit: 是一个基于 tornado 框架的快速搭建 Web应用的 Python 库, 支持大量数据表、图表等对象的渲染, 支持网格化、响应式布局.
#     可以帮助用户以高效、灵活的方式可视化数据, 分析结果.
#     同时, streamlit 还支持自定义 CSS 样式, 可以实现更丰富的页面效果.
#
#     session_state: 缓存数据, 页面刷新仍然存在

sessions_dir = Path(PurePath(__file__).parent, 'sessions')


# 生成会话文件 id
def generate_session_id():
    return datetime.now().strftime("%Y-%m-%d_%H-%M-%S")


# 保存会话
def save_session():
    if st.session_state.current_session:
        session_data = {
            'nick_name': st.session_state.nick_name,
            'current_session': st.session_state.current_session,
            'messages': st.session_state.messages
        }

    if not sessions_dir.exists():
        sessions_dir.mkdir()

    with open(sessions_dir / f'{session_data["current_session"]}.json', 'w', encoding='utf-8') as f:
        json.dump(session_data, f, ensure_ascii=False, indent=2)


# 加载会话记录列表
def load_sessions() -> list[str]:
    if not sessions_dir.exists():
        return []

    # 文件倒叙排序
    session_list = sorted([datetime.strptime(f.name.replace('.json', ''), "%Y-%m-%d_%H-%M-%S").timestamp() for f in
                           sessions_dir.glob('*.json')], reverse=True)
    return [datetime.fromtimestamp(ts).strftime("%Y-%m-%d_%H-%M-%S") for ts in session_list]


# 加载会话
def load_session(session_id: str):
    session_file = sessions_dir / f'{session_id}.json'
    try:
        if session_file.exists():
            with open(session_file, 'r', encoding='utf-8') as f:
                session_data = json.load(f)
                st.session_state.nick_name = session_data['nick_name']
                st.session_state.current_session = session_data['current_session']
                st.session_state.messages = session_data['messages']
    except Exception as e:
        st.error(f"加载会话失败: {e}")


# 删除会话
@st.dialog('删除会话')
def delete_session(session_id: str):
    session_file = sessions_dir / f'{session_id}.json'
    if session_file.exists():
        st.warning(f"确定要删除会话 {session_id} 吗?")
        if st.button('确定'):
            if session_id == st.session_state.current_session:
                st.session_state.current_session = generate_session_id()
                st.session_state.messages = []
                st.session_state.nick_name = '小甜甜'
            session_file.unlink()
            st.success(f"会话 {session_id} 删除成功")
            st.rerun()
    else:
        st.error(f"会话 {session_id} 不存在")


# 页面设置
st.set_page_config(
    page_title="AI智能伴侣",
    page_icon=":tada:",
    layout="wide",
    initial_sidebar_state="expanded",
)

# 标题
st.title("AI智能伴侣")

# logo
st.logo('🦉')

if 'nick_name' not in st.session_state:
    st.session_state.nick_name = '小甜甜'

if 'current_session' not in st.session_state:
    st.session_state.current_session = generate_session_id()

# 消息缓存, 刷新页面仍然存在
if 'messages' not in st.session_state:
    st.session_state.messages = []

# 显示历史消息
for message in st.session_state.messages:
    st.chat_message(message['role']).write(message['content'])
    # if message['role'] == 'user':
    #     st.chat_message('user').write(message['content'])
    # else:
    #     st.chat_message('assistant').write(message['content'])

# 侧边栏
with st.sidebar:
    st.sidebar.subheader("AI会话管理")
    # 创建新的会话, 清空当前会话记录
    if st.button('新会话', icon='➕️', width="stretch"):
        save_session()
        if st.session_state.messages:
            st.session_state.messages = []
            st.session_state.current_session = generate_session_id()
            save_session()
            st.rerun()  # 重新运行当前页面

    st.sidebar.caption('历史会话')
    # 加载历史记录
    session_list = load_sessions()
    for session in session_list:
        col1, col2 = st.columns([4, 1])  # 创建布局
        with col1:
            if st.button(session, icon='📄', width="stretch", key=f"load_{session}",
                         type='primary' if session == st.session_state.current_session else 'secondary'):
                if st.session_state.messages:
                    save_session()
                load_session(session)
                st.rerun()
        with col2:
            if st.button('', icon='❌', width="stretch", key=f"delete_{session}"):
                delete_session(session)

    st.divider()
    st.sidebar.subheader("伴侣信息")
    st.session_state.nick_name = st.sidebar.text_input("昵称:", placeholder="请输入昵称",
                                                       value=st.session_state.nick_name)

api_key = "sk-ea772930e62f4006a46ef1ba9f961a08"
client = OpenAI(
    api_key=api_key,
    base_url="https://dashscope.aliyuncs.com/compatible-mode/v1",
    timeout=60,
    max_retries=3,
)

system_prompt = """
你是一个优秀的智能助手, 你的名字是%s, 请使用可爱和调皮的语气回答问题
"""

# 消息输入框, 获取用户输入的问题
prompt = st.chat_input('请输入您要问的问题?')
if prompt:
    st.chat_message('user').write(prompt)
    # 缓存输入的问题
    st.session_state.messages.append({'role': 'user', 'content': prompt})
    # 调用大模型
    response = client.chat.completions.create(
        model='MiniMax-M2.5',
        messages=[
            {"role": "system", "content": system_prompt % st.session_state.nick_name},
            *st.session_state.messages,  # 将历史对话加入到上下文, 用于大模型参考
        ],
        stream=True,  # 是否流式返回
        extra_body={"thinking": {"type": "enabled"}}
    )
    # 一次获取所有返回内容
    # st.chat_message('assistant').write(response.choices[0].message.content)
    # 缓存结果
    # st.session_state.messages.append({'role': 'assistant', 'content': response.choices[0].message.content})\

    # 创建一个空的组件, 用于显示大模型返回的内容
    response_content = st.empty()
    # 流式返回
    thinking_content = ''
    content = ""
    for chunk in response:
        if chunk.choices:
            delta = chunk.choices[0].delta
            # thinking
            if hasattr(delta, "reasoning_content") and delta.reasoning_content:
                thinking_content += delta.reasoning_content
            # content
            if hasattr(delta, "content") and delta.content:
                content += delta.content
            response_content.chat_message('assistant').write(content)

    print(thinking_content)
    print('=' * 20)
    # 缓存结果
    st.session_state.messages.append({'role': 'assistant', 'content': content})
