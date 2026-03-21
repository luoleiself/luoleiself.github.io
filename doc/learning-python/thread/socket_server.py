import signal
import socket
import sys
import threading
import time

host = '127.0.0.1'
port = 65432

# 结束通信的协议消息
END_MSG = '__END__'
shutdown_flag = False

# 存储所有客户端连接，用于信号处理时广播结束消息
client_sockets = []
client_lock = threading.Lock()


def signal_handler(signum, frame):
    """处理 Ctrl+C (SIGINT) 和 kill (SIGTERM) 信号，先发送结束消息再退出"""
    global shutdown_flag
    shutdown_flag = True
    print(f'\n收到信号 {signum}，正在发送结束消息并关闭连接...')
    with client_lock:
        for sock in client_sockets[:]:
            try:
                sock.send(END_MSG.encode('utf-8'))
            except (OSError, BrokenPipeError):
                pass
            finally:
                try:
                    sock.close()
                except OSError:
                    pass
            client_sockets.remove(sock)
    try:
        server.close()
    except OSError:
        pass
    sys.exit(0)


def handle_client(sock, addr):
    """处理单个客户端的通信（每个客户端一个线程）"""
    with client_lock:
        client_sockets.append(sock)
    try:
        while not shutdown_flag:
            try:
                data = sock.recv(1024)
                if not data:
                    break
                msg = data.decode('utf-8')
                if msg == END_MSG or msg == 'exit':
                    print(f'客户端 {addr} 请求结束连接')
                    break
                print(f'客户端 {addr}: {msg}')
                if msg == 'time':
                    gmt_time = time.gmtime(time.time())
                    reply = time.strftime('%Y-%m-%dT%H:%M:%S.000Z', gmt_time)
                    sock.send(reply.encode('utf-8'))
                else:
                    input_data = input(f'回复 {addr}: ')
                    if shutdown_flag:
                        break
                    sock.send(input_data.encode('utf-8'))
                    if input_data == END_MSG:
                        break
            except (ConnectionResetError, BrokenPipeError, OSError):
                break
    finally:
        try:
            sock.close()
        except OSError:
            pass
        with client_lock:
            if sock in client_sockets:
                client_sockets.remove(sock)
        print(f'客户端 {addr} 已断开连接')


# 创建 TCP 服务器
server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
server.bind((host, port))
server.listen(512)

# 注册信号处理
signal.signal(signal.SIGINT, signal_handler)
signal.signal(signal.SIGTERM, signal_handler)

print(f'服务器已启动: {host}:{port}')
print('等待客户端连接...')

while not shutdown_flag:
    try:
        sock, addr = server.accept()
        print(f'新客户端连接: {addr}')
        t = threading.Thread(target=handle_client,
                             args=(sock, addr), daemon=True)
        t.start()
    except OSError:
        if shutdown_flag:
            break
        raise
