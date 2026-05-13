import multiprocessing
import threading
from multiprocessing import Process
from threading import Lock, Thread
import time
import os


def sing(l: Lock):
    t = threading.current_thread()  # 获取当前线程对象
    print(f'{t.name} is running...')
    for i in range(10):
        with l:
            time.sleep(0.1)

    print(f'{t.name} is end...')


print(f'main process pid: {os.getpid()}')

# 设置线程为守护线程，主线程运行结束后会立即终止守护线程, 忽略守护线程是否执行完成
# t1.daemon = True
# t2.daemon = True

lock = Lock()
thread_list = []
for i in range(10):
    t = Thread(target=sing, args=(lock,))
    t.name = f'thread-{i}'
    t.start()

for t in thread_list:
    print(f't.name: {t.name}, t.is_alive: {t.is_alive()}')
    t.join()  # 阻塞线程直到线程终止退出或超时
print('-' * 30)


# 继承 Thread 类实现多线程
class DetailHtml(Thread):
    def __init__(self, name, lock):
        super().__init__(name=name)
        self.lock = lock

    # 重写 run 方法
    def run(self):
        self.lock.acquire()
        time.sleep(2.5)
        print(f'{self.name} is running')
        self.lock.release()


def f(l: Lock):
    with l:
        current = multiprocessing.current_process()
        print(f'{current.name} is running...')


if __name__ == '__main__':
    detailHtml = DetailHtml('wadaxing', Lock())
    start_time = time.time()
    detailHtml.start()
    detailHtml.join()
    print(f'cost time: {time.time() - start_time}')
    print('-' * 20)

    lock = multiprocessing.Lock()
    p = Process(target=f, name='process-sing', args=(lock,))
    p.start()
    p.join()
