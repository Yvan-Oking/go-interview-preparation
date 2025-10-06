# Interfaces in Go

## What are interfaces in Go?

An interface is a type that specifies a contract by defining methods without implementing them. Types that implement those methods satisfy the interface. This allows for polymorphism in Go.

### Key Concepts:

- **Contract Definition**: Interfaces define what methods a type must have
- **Implicit Implementation**: Types implement interfaces implicitly (no explicit declaration needed)
- **Polymorphism**: Different types can be used interchangeably if they implement the same interface
- **Method Signatures**: Only method signatures are defined, not implementations

### Example:

```go
type Writer interface {
    Write([]byte) (int, error)
}

type File struct {
    name string
}

func (f File) Write(data []byte) (int, error) {
    // Implementation
    return len(data), nil
}

// File now implements the Writer interface
```

### Benefits:

- **Flexibility**: Allows different types to be used in the same context
- **Testability**: Easy to create mock implementations for testing
- **Composition**: Interfaces can be composed together
