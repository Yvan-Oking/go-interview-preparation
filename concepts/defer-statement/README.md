# Defer Statement in Go

## What is the purpose of the defer statement in Go?

The `defer` statement is used to postpone the execution of a function until the surrounding function returns. It's commonly used for cleanup tasks, such as closing files or unlocking mutexes.

### Key Characteristics:

- **Deferred Execution**: Function runs when the surrounding function returns
- **LIFO Order**: Defer statements are executed in Last In, First Out order
- **Cleanup Tasks**: Perfect for resource cleanup
- **Panic Safety**: Deferred functions run even if the function panics

### Example:

```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // Will be called when function returns
    
    // Process file...
    
    return nil
}
```

### Multiple Defer Statements:

```go
func example() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    defer fmt.Println("Third defer")
    
    fmt.Println("Function body")
}

// Output:
// Function body
// Third defer
// Second defer
// First defer
```

### Common Use Cases:

- **File Operations**: Closing files
- **Mutex Operations**: Unlocking mutexes
- **Database Connections**: Closing connections
- **Logging**: Recording function entry/exit
- **Recovery**: Panic recovery
