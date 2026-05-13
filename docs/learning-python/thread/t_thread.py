import multiprocessing
import os
import threading
import time


def complete_heavy_task():
    result = 0
    for i in range(10_000_000):
        result += i
    print('Task completed')


if __name__ == '__main__':
    cpu_count = os.cpu_count()
    cpu_count = cpu_count if cpu_count else 5
    print(f'cpu_count: {cpu_count}')

    threads = []
    for i in range(cpu_count):
        t = threading.Thread(target=complete_heavy_task)
        threads.append(t)

    start_time = time.perf_counter()
    for t in threads:
        t.start()

    for t in threads:
        t.join()

    end_time = time.perf_counter()
    print(f'Total time: {end_time - start_time}')

    processes = []
    for i in range(cpu_count):
        p = multiprocessing.Process(target=complete_heavy_task)
        processes.append(p)

    start_time = time.perf_counter()
    for p in processes:
        p.start()

    for p in processes:
        p.join()

    end_time = time.perf_counter()
    print(f'Total time: {end_time - start_time}')
