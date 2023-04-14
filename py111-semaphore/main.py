import threading
import time

# Create a semaphore with an initial value of 5
semaphore = threading.Semaphore(5)

print_lock = threading.Lock()

def worker():
    # Acquire the semaphore
    semaphore.acquire()

    with print_lock:
        print("Worker acquired semaphore")
    
    time.sleep(0.1)
    
    # Release the semaphore
    semaphore.release()
    
    with print_lock:
        print("Worker released semaphore")

# Create 10 worker threads
for i in range(10):
    t = threading.Thread(target=worker)
    t.start()
