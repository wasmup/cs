import asyncio

async def print_numbers():
    for i in range(1, 11):
        print(i)
        await asyncio.sleep(0.1)

async def print_letters():
    for letter in ['a', 'b', 'c', 'd', 'e']:
        print(letter)
        await asyncio.sleep(0.2)

async def main():
    task1 = asyncio.create_task(print_numbers())
    task2 = asyncio.create_task(print_letters())

    await task1
    await task2

asyncio.run(main())
