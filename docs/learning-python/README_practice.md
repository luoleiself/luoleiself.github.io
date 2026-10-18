# Python 项目开发实战心得与技术经验总结

> 按"踩坑—决策—习惯"三个维度整理的实战经验，适用于 Python 3.12+，
> 版本相关注意事项以 Python 3.14 为基准标注。

---

## 一、项目结构与工程化：先定骨架再写代码

### 1.1 src 布局优于平铺布局

```
myproj/
├── pyproject.toml
├── src/
│   └── myproj/          # 真正的包
│       ├── __init__.py  # 只放公开 API 的显式导出
│       ├── core.py
│       └── utils/
├── tests/
│   └── test_core.py
└── docs/
```

**为什么**：平铺布局（包目录直接在根下）容易让脚本"意外 import 到当前目录的同名模块"，测试时经常 import 到的是本地源码而非安装件，问题到部署才暴露。src 布局强制你只能 import 已安装的包，本地未 `pip install -e .` 立刻报错——**把环境问题提前到开发第一分钟**。

### 1.2 `__init__.py` 是 API 契约，不是杂物间

```python
# __init__.py —— 用 __all__ 显式声明，别人 from myproj import * 时才是受控的
from .core import sql, SQL
from .safe import Safe, Identifier

__all__ = ["sql", "SQL", "Safe", "Identifier"]
```

心得：内部模块随便重构，只要 `__all__` 不变就不算 breaking change。这也是工具库后来能安心删掉内部死代码的底气——它从来不在公开契约里。

### 1.3 pyproject.toml 是唯一事实来源

- `requires-python` 写死下限，别再写 `setup.py`
- 依赖只写"需要的最低版本"（`httpx>=0.27`），**锁版本交给 lock 文件**（uv.lock / poetry.lock）；库项目发布 sdist 时不带 upper bound，应用项目部署时永远带 lock
- 工具配置全部收进 pyproject（ruff/mypy/pytest 都支持 `[tool.xxx]` 段），杜绝 `.flake8`、`setup.cfg`、`pytest.ini` 三个文件打架

---

## 二、语言层经验：少即是多

### 2.1 数据结构选对了，代码量减半

最常犯的不是语法错，而是**用错容器**：

```python
# 用 dict 计数 → 用 Counter
from collections import Counter, defaultdict, deque

# 嵌套 dict 初始化地狱 → defaultdict
tree = defaultdict(dict)

# 头部弹出用 list.pop(0)（O(n)）→ deque
q = deque(maxlen=1000)     # maxlen 还是天然的滑动窗口

# "存在性判断"用 list → set（O(n) vs O(1)）

# 手写元组解包比较 → dataclass(frozen=True, slots=True)
```

一条铁律：**凡是 `if x in some_list:` 出现在循环里，先停下来把 some_list 换成 set。**

### 2.2 推导式有边界，超过就写循环

```python
# 好的：一层过滤 + 一层变换
lookup = {u.id: u for u in users if u.active}

# 坏的：嵌套两层以上 + walrus 缝合 —— 立刻改成显式循环或函数
result = [y for x in data if (m := f(x)) is not None for y in g(m)]
```

推导式的价值是"意图一目了然"，一旦需要注释才能看懂，它就失去了存在的意义。

### 2.3 异常处理的三条实战规则

```python
# 1. 捕获要精准，转抛要链接
try:
    cfg = tomllib.loads(text)
except tomllib.TOMLDecodeError as e:
    raise ConfigError(f"bad config in {path}") from e   # from e 保留因果链

# 2. 只在"能处理"的层捕获，否则让它冒泡——
#    库代码里 try/except: return None 是排障灾难的根源

# 3. 资源管理交给 with，永远不要手写 try/finally: close()
```

3.14 的新细节：多异常捕获无 `as` 时可省括号 `except TimeoutError, ConnectionRefusedError:`；但带 `as` 必须写括号——混用会让读代码的人停顿，团队里要么统一用旧风格，要么在 lint 里定死。

### 2.4 类型提示的实用主义

```python
from collections.abc import Sequence, Iterable   # 参数用抽象类型，返回值用具体类型

def top_k(items: Iterable[Item], k: int) -> list[Item]: ...
#      入参接受一切可迭代 ↑           返回承诺是 list ↑
```

- **边界处标注，内部信任推断**。公开 API、跨模块函数必须有签名；三行的小私有函数可以不写
- `X | None` 优先于 `Optional[X]`；3.14 里 `Union[int, str]` 和 `int | str` 已是同一个运行时类型，新代码别再引 `typing.Union`
- 用 mypy/pyright 做检查器，而不是人肉 review 类型。**没有检查器的类型提示会在半年内烂掉**
- 3.14 特别提醒：写库要读注解时，用 `annotationlib.get_annotations(..., format=Format.FORWARDREF)`，别再直接碰 `__annotations__`——这是 fastapi/pydantic 生态正在迁移的方式，也是前向引用不再需要加引号后代价最小的兼容姿势

---

## 三、性能与并发：先测量，再动手

### 3.1 并发模型选择表（实战中最常被问的决策）

| 场景 | 选择 | 一句话理由 |
|---|---|---|
| 大量网络 I/O（爬虫、网关） | `asyncio` + httpx | 单线程扛几千连接 |
| 阻塞库的 I/O（老 SDK、读盘） | `ThreadPoolExecutor` | I/O 期间 GIL 会释放，简单可靠 |
| CPU 密集（解析、计算） | `ProcessPoolExecutor` / 3.14 `InterpreterPoolExecutor` | 绕开 GIL；子解释器比进程轻 |
| CPU 密集 + 自由线程构建（3.14t） | `ThreadPoolExecutor` | 无 GIL 时线程直接跨核，但要求所有 C 扩展都是 ft 构建 |

心得：**90% 的"要不要上异步"争论，答案取决于依赖库里有没有阻塞调用**。一个 `requests.get` 混进 async 代码，整个事件循环停摆。混用时的逃生舱是 `asyncio.to_thread(fn, *args)`。

### 3.2 微优化的优先顺序

1. **算法与容器**（set 查重、排序一次复用、避免 O(n²) 拼接——`''.join(parts)` 而不是 `s += x`）
2. **少跨边界**：减少函数调用、属性查找，热循环里 `d = data.get` 局部化
3. `functools.cache` / `lru_cache(maxsize=...)`——最便宜的性能改进，但注意**实例方法上加 cache 会导致内存泄漏**（self 被强引用），实例级缓存用 `cached_property`
4. 最后才考虑 C 扩展 / Cython / numpy 向量化

调试侧：`python -X importtime` 看导入开销，`cProfile` + `snakeviz` 看 CPU，`py-spy top --pid` 看生产（3.14 还有 `sys.remote_exec` / `python -m pdb -p <PID>` 附加运行中进程，值得学）。

### 3.3 一个反直觉经验：字符串拼接的成本

```python
# 逐段 append 再 join —— O(n)
parts = []
for tok in stream:
    parts.append(tok.text)
text = ''.join(parts)

# 但如果是"每帧都要看全文"（流式渲染场景）：
# 用 list 存块、只在需要时 join；别每帧 "".join(全量) —— 这正是 O(n²) 的来源
```

这条与前端流式渲染的"增量累积、只重渲尾块"策略是同构问题：**凡是"每次到达都处理全量历史"的代码结构，都要警惕二次方复杂度。**

---

## 四、测试与质量：把信心建立在测试上，不是 review 上

### 4.1 pytest 的高性价比用法

```python
import pytest

@pytest.mark.parametrize("raw,expect", [
    ("```js\nx=1\n```", "code"),
    ("# h\n\ntext", "md"),
])
def test_split(raw, expect):
    assert first_block_type(raw) == expect

@pytest.fixture
def session(mocker): ...     # 依赖注入优于 monkeypatch 全局
```

- **表驱动测试是性价比之王**："输入→期望输出"型的渲染器/解析器，parametrize 二十行覆盖全部边界
- 测**行为**不测实现；重构时测试应保持绿色
- 时间/随机/网络三件套用 `freezegun` / `random.seed` / 注入假连接——**构造函数注入依赖（如 socket 打开函数作为参数传入）是测试友好的核心手段**

### 4.2 lint/format 组合拳

```bash
ruff check --fix . && ruff format . && mypy .
```

- Ruff 一个工具顶 flake8+isort+bandit，配 `target-version = "py314"` 后它还会提示你升级旧写法
- pre-commit 钩子让风格争议不再消耗 review 时间
- 经验：**每修一个 bug，先写一个能复现它的失败测试**，再动手改代码。这一条守住了，回归 bug 会肉眼可见地减少

---

## 五、避坑清单（血泪向）

```python
# 1. 可变默认参数
def f(items=[]):        # ❌ 默认对象在 def 执行时创建一次，所有调用共享
def f(items=None): items = items or []   # ✅ 或更推荐 sentinel 模式

# 2. 闭包晚绑定
funcs = [lambda: i for i in range(3)]        # 全返回 2
funcs = [lambda i=i: i for i in range(3)]    # 默认参数捕获

# 3. == 比较浮点
abs(a - b) < 1e-9   # 或 math.isclose(a, b)

# 4. is 只用于 None/True/False；小整数缓存会骗你

# 5. dict 在 3.7+ 保插入序是语言规范，但"依赖 dict 顺序"仍是坏设计

# 6. json.dumps 默认允许 NaN/Infinity —— 对接严格 JSON 解析器时 allow_nan=False

# 7. datetime.now() 无时区；服务端一律 datetime.now(timezone.utc)

# 8. finally 里 return/break/continue（3.14 起 SyntaxWarning）——
#    它会静默吞异常，是历史级 bug 温床
```

版本迁移专项（3.12/3.13 → 3.14）：

- `from __future__ import annotations` 在 3.14 已被**弃用**，纯 3.14 项目删掉它
- 旧式 Union 不再缓存，`Union[int, str] is Union[int, str]` 可能为 False，**用 `==` 比较**
- itertools 迭代器不再可 pickle/copy——缓存层里靠前求值的写法要检查
- `functools.partial` 变成方法描述符：类属性上的 partial 现在会绑定 self，需要旧行为就包 `staticmethod()`
- GC 行为在微版本间有变（3.14.0–3.14.4 增量式，3.14.5+ 回退分代式）——**调过 gc.freeze/threshold 的项目必须在目标微版本实测**

---

## 六、习惯层面的心得（比技巧更值钱）

1. **虚拟环境 + 锁文件是底线**，uv 现在是默认选择（创建/安装快一个数量级）：`uv init && uv add httpx && uv run pytest`
2. **小步提交、可运行提交**：每个 commit 的测试套件是绿的。回滚成本决定你敢不敢重构
3. **日志用 `%s` 惰性格式化**（`logger.debug("x %s", x)`），热点路径里 f-string 日志是白烧 CPU
4. **写库时"防御在边界，信任在内部"**：参数校验放在公开函数入口，内部函数之间不再重复检查
5. **文档即测试**：README 里的每个示例都应能跑（Doctest 或 `pytest --doctest-modules` 可选执行）
6. 读官方 *What's New* 只花 20 分钟，但能避免 80% 的"用三年前的写法写新代码"

---

## 相关文档

- 字符串专题（f-string / t-string 详解）：[README_str.md](./README_str.md)
- 环境工具专题：[README.uv.md](./README.uv.md)
