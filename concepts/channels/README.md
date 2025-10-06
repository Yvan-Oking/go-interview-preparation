# Channels in Go

## What are channels in Go, and how are they used?

Channels are a way for goroutines to communicate with each other and synchronize their execution. They allow you to send and receive values between goroutines safely, preventing race conditions.

### Creating Channels:

```go
ch := make(chan int)           // Unbuffered channel
ch := make(chan int, 10)       // Buffered channel with capacity 10
```

### Sending and Receiving:

```go
ch := make(chan int)

// Send data
go func() {
    ch <- 42 // Send data
}()

// Receive data
value := <-ch // Receive data
```

### Channel Types:

1. **Unbuffered Channels**:
   - Synchronous communication
   - Sender blocks until receiver is ready
   - Receiver blocks until sender is ready

2. **Buffered Channels**:
   - Asynchronous communication
   - Sender blocks only when buffer is full
   - Receiver blocks only when buffer is empty

### Example:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        ch <- "Hello from goroutine!"
    }()
    
    msg := <-ch
    fmt.Println(msg)
}
```

### Channel Operations:

- **Send**: `ch <- value`
- **Receive**: `value := <-ch`
- **Close**: `close(ch)`
- **Check if closed**: `value, ok := <-ch`

### Benefits:

- **Thread Safety**: Prevents race conditions
- **Synchronization**: Coordinates goroutine execution
- **Communication**: Enables data sharing between goroutines
