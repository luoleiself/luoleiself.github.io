import streamlit as st


# 装饰器, 缓存函数的返回数据, 通常用于数据副本
@st.cache_data(ttl='1d')
def get_data():
    return 'Hello world!'


# 装饰器, 缓存函数的返回结果对象, 通用用于数据库链接, ML 模型等
@st.cache_resource
def get_database_session(url):
    # Create a database session object that points to the URL.
    session = url
    return session


# st.context
st.write(st.context.url)
st.write(st.context.cookies)
st.write(st.context.locale)
st.write(st.context.theme)
st.write(st.context.headers)
st.write(st.context.ip_address)
st.write(st.context.is_embedded)
st.write(st.context.timezone)

# st.write(st.query_params['p'])
st.write('st.query_params.get_all("p")', st.query_params.get_all('p'))
st.write('st.query_params.to_dict()', st.query_params.to_dict())

# markdown 介绍
st.markdown("""
    ### 介绍
    本程序是一个基于 [OpenAI](https://openai.com/) 的智能对话程序,
    使用 [DashScope](https://dashscope.com/) 提供的 [MiniMax-M2.5](https://dashscope.com/model/detail?model=MiniMax-M2.5) 模型进行对话.

    ### 使用方法
    1. 输入问题
    2. 点击发送按钮
    3. 等待模型返回结果
    4. 聊天结束, 点击新会话按钮, 创建新的会话
""")

# 代码块
st.code('''
print('Hello world!')
lst = [i for i in range(10)]
for i in lst:
    print(i)
''', language="python")
st.code('''
const name = 'hello world'
console.log(name)
''', language='javascript')

# 图片
# st.image('owl.png')

# 链接
st.link_button("GitHub", "https://github.com/dashscope/examples")

# 折线图
st.line_chart([[10, 20, 30], [40, 50, 60]])
# 散点图
st.scatter_chart([[15, 22, 31], [44, 51, 62]])
# 柱状图
st.bar_chart([[10, 20, 30], [40, 50, 60]])
# 面积图
st.area_chart([[10, 20, 30], [40, 50, 60]])

# 表单
with st.form(key='my_form'):
    question = st.text_input(label="问题:", placeholder="请输入问题", value="")
    radio_button = st.radio("角色:", ['用户', '助手'], index=0)
    select_box = st.selectbox("模型:", ['MiniMax-M2.5', 'MiniMax-M2.5-Chat', 'qwen-3.6-plus', 'deepseek-v4-flash'],
                              index=0)
    checkbox = st.checkbox("保存会话", value=True)
    submit_button = st.form_submit_button(label="发送", type='primary')
    if submit_button:
        if question:
            st.chat_message(radio_button).write(question)

# 滑动条
st.slider("滑动条", 0, 100, 50)


# 片段
@st.fragment()
def my_fragment(num: int = 0):
    st.write(f"这是一个片段. {num}")
    st.write(get_data())
    st.write(get_database_session('hello streamlit'))


@st.fragment()
def my_fragment2():
    st.write("这是一个片段2")
    st.write(get_data())
    st.write(get_database_session('hello streamlit'))


# 容器
with st.container():
    my_fragment()
    my_fragment2()
