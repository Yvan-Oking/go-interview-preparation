# Synchronization in Go

## Channels and Synchronization

Channels are used to communicate between goroutines in Go. They allow you to send and receive values between goroutines safely, preventing race conditions.

### Basic Channel Usage:

```go
ch := make(chan int) // Create channel

// Send data
go func() {
    ch <- 42 // Send data
}()

// Receive data
value := <-ch // Receive data
```

## Synchronization Primitives

### 1. sync.Mutex

#### What is `sync.Mutex`?

`sync.Mutex` is a mutual exclusion lock. It is used to protect shared data from being accessed by multiple goroutines simultaneously, which can lead to race conditions and inconsistent data. The Mutex ensures that only one goroutine can access the critical section of code at a time.

#### How to Use `sync.Mutex`:

- **Locking**: Before accessing the shared resource, a goroutine must call the `Lock` method
- **Unlocking**: After finishing the operation, it must call the `Unlock` method to release the lock

#### Example:

```go
package main

import (
    "fmt"
    "sync"
)

type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

func main() {
    counter := &Counter{}
    
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    
    wg.Wait()
    fmt.Println("Final value:", counter.Value())
}
```

### 2. sync.WaitGroup

#### What is `sync.WaitGroup`?

`sync.WaitGroup` is used to wait for a collection of goroutines to finish executing. It maintains a counter that tracks how many goroutines are currently running. You can increment the counter with `Add`, decrement it with `Done`, and block the main goroutine until the counter is zero using `Wait`.

#### How to Use `sync.WaitGroup`:

- **Add**: Use `wg.Add(n)` to add `n` to the wait group counter
- **Done**: Call `wg.Done()` when a goroutine finishes its work, which decrements the counter
- **Wait**: Call `wg.Wait()` to block until the counter is zero

#### Example:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    
    wg.Wait()
    fmt.Println("All workers completed")
}
```

## Summary

- **`sync.Mutex`** is used to provide exclusive access to shared resources, preventing race conditions
- **`sync.WaitGroup`** is used to wait for a set of goroutines to finish, ensuring that the main program doesn't exit prematurely
- **Channels** provide a safe way for goroutines to communicate and synchronize
- Choose the right synchronization primitive based on your specific needs:
  - Use `Mutex` for protecting shared data
  - Use `WaitGroup` for waiting for goroutines to complete
  - Use channels for communication between goroutines
