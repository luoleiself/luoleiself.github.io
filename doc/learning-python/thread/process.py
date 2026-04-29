from multiprocessing import Pipe, Queue

"""
multiprocessing.Pipe() 创建一个双向管道返回一组连管道连接对象, 参数 duplex 默认为 True 表示管道是双向的,
    为 False 表示管道是单向的, 第一个对象只能接收消息, 第二个对象只能发送消息
    
    (conn1, conn2) = Pipe(duplex=True)   # duplex 默认为 True, 创建双向管道
    
multiprocessing.Queue() 创建一个使用管道和一些锁/信号量实现的进程共享队列, 
    当一个进程首次将一个项目放入队列时, 会启动一个馈送线程, 将对象从缓冲区传输到管道中.
"""
print('Pipe() 管道')
print('双向管道')
(conn1, conn2) = Pipe()

conn1.send('hello world')
res = conn2.recv()
print(f'Pipe conn2 received: {res}')

conn2.send_bytes(b'hello world gg')
res = conn1.recv_bytes()
print(f'Pipe conn1 received bytes: {res}')
print('-' * 20)

print('单向管道')
(conn1, conn2) = Pipe(duplex=False)
# conn1.send('hello world') # 报错, 第一个对象只能接收消息
# res = conn2.recv()
# print(f'Pipe conn2 received: {res}')
conn2.send('hello unidirectional')
res = conn1.recv()
print(f'Pipe conn1 received: {res}')
conn2.send_bytes(b'hello unidirectional')
res = conn1.recv_bytes()
print(f'Pipe conn1 received bytes: {res}')
print('-' * 20)

if __name__ == '__main__':
    q = Queue(10)
    q.put(1)
