# Mutex and Channel basics

### What is an atomic operation?
> *An atomic operation is an operation that is guaranteed to complete before any other operation is executed. This is often used in an operating system when doing critial work.*

### What is a semaphore?
> *A semaphore is a for of integer which is often used for syncronization or as a mutex. When a thread tries to 'get' a semaphore and it is 0, then the process is blocked until the someone else unlocks the semaphore. 'Geting' a semaphore decreases it by 1, 'signaling' a semaphore increments it by 1*

### What is a mutex?
> *A mutex is a binary semaphore with owenership. It can either be locked or unlocked, and only the thread which locked it can unlock it. It is often used to protect a shared resource which can only be used by one thread at a time.*

### What is the difference between a mutex and a binary semaphore?
> *A mutex has the concept of ownership. Only the thread who locked it can unlock it. For a binary semaphore, anyone can 'unlock' it. Semaphores are often used for syncronization, while mutexes are used for protecting critical sections.*

### What is a critical section?
> *A critical section is a section of code where multiple threads do work on the same shared resource. This is a section that needs protection, otherwise it can/will result in an error.*

### What is the difference between race conditions and data races?
 > *A data race occurs when 2 instructions from different thread access the same memory location, at least one of the instructions is a write, and there is no way of knowing which instruction will execute first (no protection of critical section). If the write happens first, then the read value will be different from if the read was the first instruction. A race condition is when the outcome of some operation is determined by the order of execution. This can happen in multithreaded or distributes software programs if no syncronization method is used to determine the order of execution.*

### List some advantages of using message passing over lock-based synchronization primitives.
> *Some advantages of using message passing is that it can be more modular, because the modules interface is defined by the messages it can recieve and send. It can make adding/changing modules easier.*

### List some advantages of using lock-based synchronization primitives over message passing.
> *....*
