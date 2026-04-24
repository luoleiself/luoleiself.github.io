import subprocess

"""
subprocess.run()    # 运行命令, 返回一个 CompletedProcess 实例
    args    # 获取运行参数
    stdout  # 获取输出
    stderr  # 获取错误输出

subprocess.Popen()  # 实例化子进程
    poll()  # 检查命令是否完成
    wait()  # 等待命令完成
    communicate()   # 与子进程通信
    stdout/stderr   # 获取输出结果
"""

# 输入参数
arg = input('input parameter:')
p1 = subprocess.run(["dir", arg], shell=True, check=True, capture_output=True, text=True)
print(f'process: {p1}')
print(f'process.returncode: {p1.returncode} {p1.check_returncode()}')
print(f'process.args: {p1.args}')
print(f'process.stdout: {p1.stdout}')
print(f'process.stderr: {p1.stderr}')
try:
    p1.check_returncode()
except subprocess.CalledProcessError as e:
    print(f'CalledProcessError: {e}')
    print(f'process.stderr: {p1.stderr}')

print('-' * 30)

print('Popen: 高级用法')
with subprocess.Popen(['ping', '-n', '4', 'www.baidu.com'], text=True, universal_newlines=True, stdout=subprocess.PIPE,
                      stderr=subprocess.PIPE) as p:
    for line in p.stdout:
        print(line, end='')

    result = p.communicate(input='y\n')
    print(f'result: {result}')
