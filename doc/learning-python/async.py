import asyncio
import time

'''
- async 定义协程函数, 直接调用时返回协程对象, 必须由 asyncio.run() 运行
- await 可等待的对象, 协程对象, Future 对象, Task 对象
  - 直接 await 协程对象不会调度协程只是让出执行权, 需要包装成任务对象才会被调度
- asyncio.create_task() 创建 Task 对象加入任务调度列表
- asyncio.gather() 聚合多个协程对象或任务对象加入任务调度列表
- asyncio.wait() 等待任务列表全部完成
- asyncio.run() 运行协程对象
async for 用于迭代异步可迭代对象 (实现了 __aiter__ 的对象), 只能用在 async 协程函数中, 否则报错
async with 用于获取异步上下文管理器, 只能用在 async 协程函数中, 否则报错
'''


async def foo():
    print('foo step 1')
    await asyncio.sleep(2)
    print('foo step 2   ')
    return 'foo'


async def bar():
    print('bar step 3')
    await asyncio.sleep(1)
    print('bar step 4')
    return 'bar'

# 直接调用协程函数返回协程对象
print(foo(), bar())
print('---------------------------------')

print(r'''
方式1: 已过时
使用 asyncio.ensure_future() 创建 Future 对象
使用 asyncio.get_event_loop() 获取事件循环
使用 asyncio.wait() 等待任务列表全部完成
使用 loop.run_until_complete() 事件循环执行任务列表

works = [
    asyncio.ensure_future(foo()),
    asyncio.ensure_future(bar())
]
loop = asyncio.get_event_loop()
loop.run_until_complete(asyncio.wait(works))
print(works[0].result(), works[1].result())
''')
works = [
    asyncio.ensure_future(foo()),
    asyncio.ensure_future(bar())
]
loop = asyncio.get_event_loop()
loop.run_until_complete(asyncio.wait(works))
print(works[0].result(), works[1].result())
print('---------------------------------')

print(r'''
方式2:
使用 asyncio.create_task() 创建 Task 对象加入任务调度列表
使用 asyncio.run() 运行协程对象

async def main():
    task1 = asyncio.create_task(foo())
    task2 = asyncio.create_task(bar())
    ret1 = await task1  # 等待任务对象完成
    ret2 = await task2
    print(ret1, ret2)

asyncio.run(main())
''')


async def main():
    task1 = asyncio.create_task(foo())
    task2 = asyncio.create_task(bar())
    ret1 = await task1  # 等待任务对象完成
    ret2 = await task2
    print(ret1, ret2)

asyncio.run(main())
print('-------------')

print(r'''
直接 await 协程对象不会调度协程只是让出执行权, 需要包装成任务对象才会被调度

async def main():
    ret1 = await foo()
    ret2 = await bar()
    print(ret1, ret2)

asyncio.run(main())
''')


async def main():
    ret1 = await foo()
    ret2 = await bar()
    print(ret1, ret2)

asyncio.run(main())
print('---------------------------------')

print(r'''
方式3:
使用 asyncio.create_task() 创建 Task 对象加入任务调度列表
使用 asyncio.wait() 等待任务列表全部完成
使用 asyncio.run() 运行协程对象

async def main():
    task_lists = [
        asyncio.create_task(foo()),
        asyncio.create_task(bar())
    ]
    done, pending = await asyncio.wait(task_lists)
    print(done, pending)
    for task in done:
        print(task.result())

asyncio.run(main())
''')


async def main():
    task_lists = [
        asyncio.create_task(foo()),
        asyncio.create_task(bar())
    ]
    done, pending = await asyncio.wait(task_lists)
    print(done, pending)
    for task in done:
        print(task.result())

asyncio.run(main())
print('---------------------------------')

print(r'''
方式4:
使用 asyncio.gather() 聚合多个协程对象或任务对象加入任务调度列表
使用 asyncio.run() 运行协程对象

async def main():
    # 聚合多个协程对象
    ret = await asyncio.gather(foo(), bar())
    print(ret)

asyncio.run(main())
''')


async def main():
    # 聚合多个协程对象
    ret = await asyncio.gather(foo(), bar())
    print(ret)

asyncio.run(main())
print('-------------')

print(r'''
方式4-1:
使用 asyncio.create_task() 创建任务对象
使用 asyncio.gather() 聚合多个协程对象或任务对象加入任务调度列表
使用 asyncio.run() 运行协程对象

async def main():
    task1 = asyncio.create_task(foo())
    task2 = asyncio.create_task(bar())
    # 聚合多个任务对象
    ret = await asyncio.gather(task1, task2)
    print(ret)

asyncio.run(main())
''')


async def main():
    task1 = asyncio.create_task(foo())
    task2 = asyncio.create_task(bar())
    # 聚合多个任务对象
    ret = await asyncio.gather(task1, task2)
    print(ret)

asyncio.run(main())
print('---------------------------------')

print(r'''
async for 用于迭代异步可迭代对象 (实现了 __aiter__ 的对象), 不是直接迭代协程或 gather 的返回值
gather() 返回的是「单个协程」, 应 await; 若要用 async for, 需借助 async 生成器等产出异步迭代
只能在 async 协程函数中使用, 否则报错: Use of "async" not allowed outside of async function

async def gather_async_iter():
    """把 await gather 的结果逐个 yield, 得到异步迭代器供 async for 使用。"""
    for value in await asyncio.gather(foo(), bar()):
        yield value


async def main():
    async for item in gather_async_iter():
        print(item)

asyncio.run(main())
''')


async def gather_async_iter():
    """把 await gather 的结果逐个 yield, 得到异步迭代器供 async for 使用。"""
    for value in await asyncio.gather(foo(), bar()):
        yield value


async def main():
    async for item in gather_async_iter():
        print(item)

asyncio.run(main())
print('---------------------------------')

print(r'''
async with 用于获取异步上下文管理器
只能用在 async 协程函数中, 否则报错: Use of "async" not allowed outside of async function

async def main():
    async with asyncio.Lock() as lock:
        print('async with lock 1...')
        await asyncio.sleep(2)
        print('async with lock 2...')

asyncio.run(main())
''')


async def main():
    async with asyncio.Lock() as lock:
        start = time.time()
        print('async with lock 1...')
        await asyncio.sleep(2)
        print('async with lock 2...')
        end = time.time()
        print(f'cost time: {end - start}')

asyncio.run(main())
print('---------------------------------')

# async def main():
#     print('main step m1')
#     loop = asyncio.get_event_loop()
#     future = loop.run_in_executor(None, foo)
#     ret = await asyncio.wait([future])
#     print(ret)
#     print('main step m2')

# asyncio.run(main())
# print('---------------------------------')
