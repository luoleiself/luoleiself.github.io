from concurrent.futures import ProcessPoolExecutor, ThreadPoolExecutor, as_completed

print('进程池任务必须是模块顶层可 pickle 的函数，lambda 无法被 pickle（Windows spawn 会报错）。')
print('----------------')


def process_cube(x: int) -> int:
    """进程池任务必须是模块顶层可 pickle 的函数，lambda 无法被 pickle（Windows spawn 会报错）。"""
    return x**3


if __name__ == '__main__':
    print('ThreadPoolExecutor: 线程池执行器, executor.submit() 提交任务返回 Future 对象')
    print('as_completed(): 返回迭代器对象，等待所有线程完成并返回结果')
    print('future.result(): 返回线程执行结果')
    with ThreadPoolExecutor(max_workers=5, thread_name_prefix='my_thread') as thread_executor:
        futures = []
        for i in range(4, 9):
            future = thread_executor.submit(lambda x: x ** 2, i)
            futures.append(future)
        for future in as_completed(futures):
            print(f'future.result(): {future.result()}')
        print('---------')

        print('executor.map(): 返回迭代器对象，提交任务并等待所有线程完成并返回结果')
        for data in thread_executor.map(lambda x: x ** 2, range(4, 9)):
            print(f'data: {data}')
    print('----------------------------')

    print('ProcessPoolExecutor: 进程池执行器, executor.submit() 提交任务返回 Future 对象')
    process_executor = ProcessPoolExecutor(max_workers=5)

    future1 = process_executor.submit(process_cube, 4)
    future2 = process_executor.submit(process_cube, 5)
    future3 = process_executor.submit(process_cube, 6)
    future4 = process_executor.submit(process_cube, 7)
    future5 = process_executor.submit(process_cube, 8)
    print('as_completed(): 返回迭代器对象，等待所有线程完成并返回结果')
    print('future.result(): 返回线程执行结果')
    for future in as_completed([future1, future2, future3, future4, future5]):
        print(f'future.result(): {future.result()}')
    print('---------')

    print('executor.map(): 返回迭代器对象，提交任务并等待所有线程完成并返回结果')
    for data in process_executor.map(process_cube, range(4, 9)):
        print(f'data: {data}')
    process_executor.shutdown(wait=True)
    print('----------------------------')
