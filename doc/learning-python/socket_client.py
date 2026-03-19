import signal
import socket
import sys

# 结束通信的协议消息（需与服务器一致）
END_MSG = '__END__'
shutdown_flag = False

client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client.connect(('127.0.0.1', 65432))

def signal_handler(signum, frame):
    """处理 Ctrl+C (SIGINT) 和 kill (SIGTERM) 信号，先发送结束消息再退出"""
    global shutdown_flag
    shutdown_flag = True
    print(f'\n收到信号 {signum}，正在发送结束消息并关闭连接...')
    try:
        client.send(END_MSG.encode('utf-8'))
    except (OSError, BrokenPipeError):
        pass
    finally:
        try:
            client.close()
        except OSError:
            pass
    sys.exit(0)


# 注册信号处理
signal.signal(signal.SIGINT, signal_handler)
signal.signal(signal.SIGTERM, signal_handler)

print('已连接到服务器，输入 "exit" 结束通信')

while not shutdown_flag:
    try:
        input_data = input('输入消息: ')
        if shutdown_flag:
            break
        if input_data == 'exit' or input_data == END_MSG:
            client.send(END_MSG.encode('utf-8'))
            print('已发送结束消息，关闭连接')
            client.close()
            break
        client.send(input_data.encode('utf-8'))
        data = client.recv(1024)
        if not data:
            print('服务器已关闭连接')
            break
        msg = data.decode('utf-8')
        if msg == END_MSG:
            print('服务器请求结束连接')
            client.close()
            break
        print(f'服务器: {msg}')
    except (ConnectionResetError, BrokenPipeError, OSError) as e:
        print(f'连接异常: {e}')
        break
