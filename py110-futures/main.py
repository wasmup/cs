import concurrent.futures

def square(x):
    return x ** 2

with concurrent.futures.ThreadPoolExecutor() as executor:
    future = executor.submit(square, 5)
    result = future.result()
    print(result)
