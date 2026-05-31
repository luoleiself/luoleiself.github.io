from concurrent.futures import ProcessPoolExecutor, ThreadPoolExecutor, as_completed, Future

"""
ThreadPoolExecutor: 线程池执行器
    InterpreterPoolExecutor: 解释器执行器, 每个执行器相互隔离不能共享任何可变对象和其他数据,
                                每个解释器都有自己的 GIL, 这意味着真正的多核并行执行. 3.14 支持
ProcessPoolExecutor: 进程池执行器
Future: 封装异步执行的可调用对象, 通过 executor.submit() 创建
    - .result() # 获取结果
    - .done()   # 检查任务状态
    - .running()    # 返回任务是否被执行没有被取消
    - .cancel() # 尝试取消任务
    - .cancelled()  # 返回任务是否被成功取消
    - .add_done_callback(fn)  # 当 Future 完成(正常结束、被取消或抛出异常)时调用回调函数

as_completed(futures, timeout=None)  # 返回迭代器对象, 等待所有任务完成并返回结果
"""

print('进程池任务必须是模块顶层可 pickle 的函数，lambda 无法被 pickle（Windows spawn 会报错）。')
print('-' * 30)


def process_cube(x: int) -> int:
    """进程池任务必须是模块顶层可 pickle 的函数，lambda 无法被 pickle（Windows spawn 会报错）。"""
    return x ** 3


if __name__ == '__main__':
    print('ThreadPoolExecutor: 线程池执行器, executor.submit() 提交任务返回 Future 对象')
    print('as_completed(futures): 返回迭代器对象，等待所有线程完成并返回结果')
    print('future.result(): 返回线程执行结果')
    with ThreadPoolExecutor(max_workers=5, thread_name_prefix='my_thread') as thread_executor:
        futures = [thread_executor.submit(
            lambda x: x ** 2, i) for i in range(4, 9)]
        for future in as_completed(futures):
            print(
                f'future.result(): {future.result()} future.done(): {future.done()} future.cancelled(): {future.cancelled()}')

        print('-' * 10)

        print('executor.map(): 返回迭代器对象，提交任务并等待所有线程完成并返回结果')
        for data in thread_executor.map(lambda x: x ** 2, range(4, 9)):
            print(f'data: {data}')

        print('executor.shutdown(wait=True, cancel_future=False) 关闭执行器, 使用 with 语句时会自动调用此方法')
        # thread_executor.shutdown(wait=True, cancel_future=False)

    print('-' * 30)

    print('ProcessPoolExecutor: 进程池执行器, executor.submit() 提交任务返回 Future 对象')
    print('as_completed(futures): 返回迭代器对象，等待所有线程完成并返回结果')
    print('future.result(): 返回线程执行结果')
    with ProcessPoolExecutor(max_workers=5) as process_executor:
        futures = [process_executor.submit(process_cube, i) for i in range(4, 9)]
        for future in as_completed(futures):
            print(f'future.result(): {future.result()}')

        print('-' * 10)

        print('executor.map(): 返回迭代器对象，提交任务并等待所有线程完成并返回结果')
        for data in process_executor.map(process_cube, range(4, 9)):
            print(f'data: {data}')

        print('process_executor.shutdown(wait=True, cancel_futures=False) 关闭执行器, 使用 with 语句时会自动调用此方法')
        # process_executor.shutdown(wait=True, cancel_futures=False)
        print('-' * 10)

    print('-' * 30)
