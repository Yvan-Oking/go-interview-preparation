# Concurrency in Go

## Concurrency in Go

Concurrency in Go is handled using goroutines. A goroutine is a lightweight, independent execution thread managed by Go's runtime. Goroutines allow you to run multiple functions concurrently without managing threads explicitly.

### Key Concepts:

- **Concurrency**: Managing many tasks independently, even if they don't all run simultaneously
- **Parallelism**: Achieving simultaneous execution by distributing tasks across multiple CPU cores
- **Goroutines**: Lightweight threads that enable easy concurrency

### Concurrency vs Parallelism:

#### Concurrency:
- Managing many tasks independently
- Tasks may or may not run at the same time
- Go achieves concurrency easily with goroutines
- Focus on program structure and design

#### Parallelism:
- Achieving simultaneous execution
- Distributing tasks across multiple CPU cores
- Go can run tasks in parallel using available cores
- Controlled by `runtime.GOMAXPROCS()`

### Example:

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    // Set number of OS threads
    runtime.GOMAXPROCS(4)
    
    // Concurrent execution
    go task("A")
    go task("B")
    go task("C")
    
    time.Sleep(2 * time.Second)
}

func task(name string) {
    for i := 0; i < 5; i++ {
        fmt.Printf("Task %s: %d\n", name, i)
        time.Sleep(100 * time.Millisecond)
    }
}
```

### Summary:

- **Concurrency**: Managing many tasks independently, even if they don't all run simultaneously. Go achieves concurrency easily with goroutines.
- **Parallelism**: Achieving simultaneous execution by distributing tasks across multiple CPU cores. Go can run tasks in parallel by using the system's available cores.
- Go's concurrency model makes it easy to build programs that handle many tasks, and with the right hardware and settings (`runtime.GOMAXPROCS()`), it can achieve true parallel execution.
