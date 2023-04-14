from multiprocessing import Process

def square_numbers():
    for i in range(100):
        result = i * i
        print(f"Number {i} squared is {result}")

if __name__ == '__main__':
    processes = []
    num_processes = 4

    for i in range(num_processes):
        p = Process(target=square_numbers)
        processes.append(p)

    for p in processes:
        p.start()

    for p in processes:
        p.join()

    print('Done')
