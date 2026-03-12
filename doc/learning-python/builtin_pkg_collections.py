from collections import Counter, deque, namedtuple
from queue import Queue

print('具名元组: 可使用 .属性 或 下标访问')
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

print('双端队列')
d = deque([1, 2, 3, 4, 5])
print(f'd {d}')
for i in d:
    print(i)
d.pop()
print(f'右侧弹出 d.pop() {d}')
d.popleft()
print(f'左侧弹出 d.popleft() {d}')
d.extendleft([0, -1, -2])
print(f'左侧扩展 d.extendleft([0, -1, -2]) {d}')
d.append(8)
print(f'右侧添加 d.append(8) {d}')
d.appendleft(7)
print(f'左侧添加 d.appendleft(7) {d}')
print('------------------------------------')

print('队列: 当队列满时则阻塞等待取出元素, 当队列为空时则阻塞等待加入元素')
q = Queue(5)
print(f'q 类型 {type(q)} q.empty() {q.empty()}')
q.put(1)
print(f'q.put(1) q.qsize() {q.qsize()}')
q.put(2)
print(f'q.put(2) q.qsize() {q.qsize()}')
q.put(3)
q.put(4)
print(f'q.full() {q.full()} q.empty() {q.empty()}')
q.put(5)
print(f'q.qsize() {q.qsize()} q.full() {q.full()}')
# q.put(3)    # 待队列有空位
# 从队列中取出一个元素
print(f'q.get() q.qsize() {q.get()} {q.qsize()}')
print('------------------------------------')
