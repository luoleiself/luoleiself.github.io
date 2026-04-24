from threading import Barrier, Event, Lock, RLock, Condition, Semaphore, Thread
import random
import time

print('Lock: 互斥锁, 不支持同一线程多次获取锁（会导致死锁）')
print(f'只有 锁定 和 非锁定 状态')
# 互斥, 不可重入
lock = Lock()
count = 0


def add(l: Lock):
    """
    使用 with 语句管理锁

    等价于使用 lock.acquire() 和 lock.release() 方法
    """
    global count
    for _ in range(1000000):
        # with l:
        #     count += 1
        l.acquire()  # 不能重复获取锁, 会导致死锁
        count += 1
        l.release()


def sub(l: Lock):
    global count
    for _ in range(1000000):
        with l:  # 使用 with 语句管理锁
            count -= 1


t1 = Thread(target=add, args=(lock,))
t2 = Thread(target=sub, args=(lock,))
t1.start()
t2.start()
t1.join()
t2.join()
print(f'最终结果: {count=}')
print('-' * 30)

print('RLock: 可重入锁, 基于 Lock 实现, 支持同一线程多次获取锁')
print(f'通过记录锁的获取次数, 但必须对应相同次数的 release()')
# 可重入, 同一线程可多次获取
rlock = RLock()
counter = 0


def nested_incre():
    """嵌套函数也需要锁"""
    global counter  # 声明全局变量
    with rlock:  # 使用 with 管理锁
        counter += 1


def incre_twice():
    """外层函数先获取锁"""
    global counter  # 声明全局变量
    with rlock:  # 使用 width 管理锁
        nested_incre()
        nested_incre()


threads = []

for _ in range(5):
    t = Thread(target=incre_twice)
    threads.append(t)
    t.start()

for t in threads:
    t.join()

print(f'最终结果: {counter=}')
print('-' * 30)

print(f'Condition: 条件变量, 内部基于 RLock 实现')
print(f'允许线程在特定条件满足时才继续执行')
print(f'wait() 释放锁并等待通知')
print(f'notify()/notify_all() 唤醒等待的线程')
# 条件等待 + 通知
cond = Condition()
MAX_QUEUE = 5
queue = []


def producer():
    """
    生产者线程
    """
    for _ in range(10):
        with cond:  # 使用 with 管理锁
            while len(queue) >= MAX_QUEUE:
                print(f'队列已满, 生产者等待...')
                cond.wait()
            item = f'商品-{i}'
            queue.append(item)
            print(f'生产者生产: {item}')

            # 通知消费者线程
            cond.notify()
        time.sleep(random.uniform(0.5, 1.1))


def consumer():
    """
    消费者线程
    """
    for _ in range(10):
        with cond:  # 使用 with 管理锁
            while not queue:
                print(f'队列为空, 消费者等待...')
                cond.wait()
            item = queue.pop(0)
            print(f'消费者消费: {item}')

            # 通知生产者线程
            cond.notify()
        time.sleep(random.uniform(1, 2.5))


t1 = Thread(target=producer)
t2 = Thread(target=consumer)
t1.start()
t2.start()
t1.join()
t2.join()
print('-' * 30)

print(f'Semaphore: 信号, 基于 Condition 实现, 初始化时指定允许的最大线程数')
print(f'允许一定数量的线程同时访问共享资源')
# 计数器, 允许多线程同时访问
sem = Semaphore(3)  # 最大同时运行线程数


def access_database(thread_id):
    """
    模拟访问数据库

    使用 with 管理信号量, 等价于使用 sem.acquire() 和 sem.release() 方法
    """
    with sem:
        print(f'{thread_id} 开始访问数据库')
        # 模拟访问数据库耗时
        time.sleep(random.uniform(1, 3))
        print(f'{thread_id} 结束访问数据库')


threads = []
for i in range(1, 11):
    t = Thread(target=access_database, args=(i,))
    threads.append(t)
    t.start()

for t in threads:
    t.join()
print('-' * 30)

print(f'Event: 事件, 内部基于 Condition 实现, 允许一个或多个线程等待某个事件发生')
print(f'set() 设置标志为 True, 唤醒所有等待线程')
print(f'clear() 清除标志为 False')
print(f'wait() 等待标志变为 True')
print(f'is_set() 检查标志是否为 True')
print(
    f'单一 Event: 初始 False → worker wait() 阻塞；set() 表示“开始工作”；'
    f'工作期间保持 True；clear() 表示“停止”，while is_set() 退出。'
)
# 标志信号
run_evt = Event()


def worker(name):
    """
    工作线程：先 wait() 直到主线程 set() 再进入工作循环；主线程 clear() 后循环结束。
    """
    print(f'线程{name} 已就绪, 等待开始信号...')
    run_evt.wait()
    print(f'线程{name} 开始执行...')

    while run_evt.is_set():
        print(f'线程{name} 工作中...')
        time.sleep(random.uniform(0.5, 1))
    print(f'线程{name} 收到停止信号(已 clear), 停止工作...')


threads = []
for i in range(1, 5):
    t = Thread(target=worker, args=(i,))
    threads.append(t)
    t.start()

print(f'主线程准备 5 秒后启动工作线程...')
time.sleep(5)
print(f'发送开始信号! run_evt.set()')
run_evt.set()

# 主线程等待 10 秒
time.sleep(10)
print(f'发送停止信号! run_evt.clear()')
run_evt.clear()

for t in threads:
    t.join()
print('-' * 30)

print(f'Barrier: 障碍, 内部基于 Condition 实现, 初始化时指定允许互相等待的最大线程数')
print(f'wait() 阻塞线程, 直到所有线程都到达障碍点')
# 线程互相等待
barrier = Barrier(5)


def stage(name, stage_num):
    print(f'线程{name} 正在阶段{stage_num}')
    time.sleep(random.uniform(1, 3))
    print(f'线程{name} 已完成阶段{stage_num}，等待其他线程...')
    barrier.wait()
    print(f'线程{name} 通过阶段{stage_num}, 进入下一个阶段')


threads = []

for i in range(1, 6):
    t = Thread(target=stage, args=(i, 1))
    threads.append(t)
    t.start()

for t in threads:
    t.join()
print('-' * 30)
