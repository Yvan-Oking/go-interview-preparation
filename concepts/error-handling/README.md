# Error Handling in Go

## How does error handling work in Go?

In Go, errors are treated as values. Functions that may produce an error typically return an error as the last return value. It's the caller's responsibility to check this value.

### Example:

```go
func doSomething() error {
    // Some logic
    return nil // or an error
}

if err := doSomething(); err != nil {
    fmt.Println("Error", err)
}
```

### Key Points:

- **Error as Value**: Errors are not exceptions, they are regular values
- **Explicit Handling**: The caller must explicitly check for errors
- **Last Return Value**: Error is typically the last return value
- **No Panic**: Errors don't cause panics by default (unlike exceptions in other languages)
