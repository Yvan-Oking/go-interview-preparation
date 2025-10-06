# Goroutines in Go

## Explain the concept of goroutines

Goroutines are lightweight threads managed by the Go runtime. They allow functions to run concurrently, making it easy to write programs that can perform multiple tasks at once. Goroutines are inexpensive to create and manage, allowing you to scale your applications efficiently.

### Key Characteristics:

- **Lightweight**: Very small memory footprint (2KB initial stack)
- **Managed by Runtime**: Go runtime handles scheduling
- **Concurrent**: Can run thousands of goroutines simultaneously
- **Simple Syntax**: Just add `go` keyword before function call

### Basic Usage:

```go
func main() {
    go sayHello("World")
    go sayHello("Go")
    
    time.Sleep(1 * time.Second)
}

func sayHello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

### Creating Goroutines:

```go
// Anonymous function
go func() {
    fmt.Println("Running in goroutine")
}()

// Named function
go myFunction()

// Method call
go obj.Method()
```

### Goroutine Lifecycle:

1. **Creation**: `go` keyword creates new goroutine
2. **Execution**: Goroutine runs concurrently
3. **Completion**: Goroutine ends when function returns
4. **Cleanup**: Runtime automatically cleans up resources

### Example with Multiple Goroutines:

```go
package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collect results
    for a := 1; a <= 5; a++ {
        <-results
    }
}
```

### Benefits:

- **Concurrency**: Easy to write concurrent programs
- **Efficiency**: Low overhead compared to OS threads
- **Scalability**: Can handle thousands of goroutines
- **Simplicity**: Simple syntax and management
