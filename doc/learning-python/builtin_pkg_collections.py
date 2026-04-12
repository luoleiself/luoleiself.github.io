import random
import threading
from collections import Counter, deque, namedtuple
from queue import Queue
import time

print('''
具名元组: 可使用 .属性 或 下标访问
    参数 field_names: 格式支持包含字符的序列, 或者是以空格或者逗号分割的字符串形式
    namedtuple('Point', ['a', 'b'])
    namedtuple('Point', 'a b')
    namedtuple('Point', 'a, b')
''')
Point = namedtuple('Point', ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'])
p = Point(1, 2, 3, 4, 2, 3, 4, 5)
# p.z   # 访问不存在的属性报错
print(f'p {p} p 类型 {type(p)}')
print(f'p.a {p.a} p.b {p.b}')
print(f'p[0] {p[0]} p[1] {p[1]}')
print(f'p.count(1) {p.count(1)}')  # 统计元素出现的次数
print(f'p.count(2) {p.count(2)}')  # 统计元素出现的次数
dt = p._asdict()
print(f'p._asdict() {dt}')

p2 = Point(**dt)
print(f'Point(**dt) {p2}')
print('------------------------------------')

print('计数器')
s = 'abcabcdefgadefghihbefacdhgadedc'
c = Counter(s)
print(f's {s} len {len(s)}, c {c}')
print(f'c.total() {c.total()}')
print(f'c.most_common(5) {c.most_common(5)}')
print(f'c.keys() {c.keys()}')
print(f'c.values() {c.values()}')
print(f'c.items() {c.items()}')
c.update('abc')
print(f'c.update("abc") c.most_common(5) {c.most_common(5)}')
print('------------------------------------')

print('双端队列:')
print(f'未指定 maxlen 时, 队列的长度会无限增长')
print(f'如果设置了 maxlen, 当队列长度大于 maxlen 时, 会自动删除另一侧的元素')
d = deque([1, 2, 3, 4, 5], maxlen=6)
print(f'类型 {type(d)} d.长度 {len(d)} d.maxlen {d.maxlen}')
print(f'右侧追加元素, 删除左侧超出队列长度的元素')
for i in 'abc':
    d.append(i)
    print(f'd.append({i}) d {d}')
print('------------')
print(f'左侧追加元素, 删除右侧超出队列长度的元素')
for i in 'ABC':
    d.appendleft(i)
    print(f'd.appendleft({i}) d {d}')
print('------------------------------------')

# 使用多线程读写队列
print('Queue: 当队列缓冲满时则阻塞等待取出元素, 当队列为空时则阻塞等待加入元素')
print(f'Queue: 使用 while 遍历获取队列元素')
thread = threading.current_thread()
q = Queue(maxsize=4)
print(f'{thread.name} q 类型 {type(q)} q.qsize() {q.qsize()} q.maxsize {q.maxsize}')
print(f'{thread.name} q.full() {q.full()} q.empty() {q.empty()} ')


# 哨兵：生产线程结束后每个消费者各放一个，收到则退出（勿用 empty() 判断，多线程不可靠）
_QUEUE_SENTINEL = object()


def thread_put_to_queue(q: Queue, num_workers: int):
    thread = threading.current_thread()
    print(f'{thread.name} q.qsize() {q.qsize()}')
    for i in range(1, 30):
        q.put(i)
        print(f'\033[0;42m{thread.name}\033[0m q.put({i})')
    for _ in range(num_workers):
        q.put(_QUEUE_SENTINEL)
        print(f'\033[0;42m{thread.name}\033[0m q.put(SENTINEL 结束信号)')


def thread_get_from_queue(q: Queue):
    thread = threading.current_thread()
    frontend_color = random.randint(30, 37)
    print(f'\033[0;{frontend_color}m{thread.name}\033[0m q.qsize() {q.qsize()}')
    while True:
        time.sleep(0.3)  # 线程休眠 0.3 秒
        v = q.get()
        if v is _QUEUE_SENTINEL:
            q.task_done()
            print(
                f'\033[0;{frontend_color}m{thread.name}\033[0m 收到结束信号，退出')
            break
        print(f'\033[0;{frontend_color}m{thread.name}\033[0m q.get() {v}')
        q.task_done()


NUM_GET_THREADS = 3
thread_put = threading.Thread(
    target=thread_put_to_queue,
    name='thread_put',
    args=(q, NUM_GET_THREADS),
)
thread_put.start()
thread_get_list: list[threading.Thread] = []
for i in range(1, NUM_GET_THREADS + 1):
    thread_get = threading.Thread(
        target=thread_get_from_queue, args=(q,), name=f'thread_get_{i}'
    )
    thread_get_list.append(thread_get)
    thread_get.start()
q.join()  # 等待所有 put 对应的 task_done（含哨兵）
for t in thread_get_list:
    t.join()
print('------------------------------------')
