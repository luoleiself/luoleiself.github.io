import threading
from multiprocessing import Process
from threading import Lock, Thread
import time
import os

count = 0


def sing():
    for _ in range(1_000_000):
        global count
        count += 1
        # time.sleep(1.8)


def dance():
    for _ in range(1_000_000):
        global count
        count -= 1
        # time.sleep(1)


print(f'main process pid: {os.getpid()}')
t1 = Thread(target=sing, name='SingThread')
t2 = Thread(target=dance, name='DanceThread')
# 设置线程为守护线程，主线程运行结束后会立即终止守护线程, 忽略守护线程是否执行完成
# t1.daemon = True
# t2.daemon = True
t1.start()
t2.start()
print(t1.is_alive())
t1.join()   # 阻塞线程直到线程终止退出或超时
t2.join()
print(f'{count=:>10}')
print('--------------------------------------')


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


if __name__ == '__main__':
    detailHtml = DetailHtml('wadaxing', Lock())
    start_time = time.time()
    detailHtml.start()
    detailHtml.join()
    print(f'cost time: {time.time() - start_time}')
print('--------------------------------------')
