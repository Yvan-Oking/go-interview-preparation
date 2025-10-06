# Structs in Go

## What is a struct in Go?

A struct is a composite data type that groups together variables (fields) under a single name. Structs are useful for creating complex data types that represent real-world entities.

### Basic Syntax:

```go
type Person struct {
    Name string
    Age  int
}
```

### Creating Struct Instances:

```go
// Method 1: Field names
p1 := Person{
    Name: "Alice",
    Age:  30,
}

// Method 2: Positional (order matters)
p2 := Person{"Bob", 25}

// Method 3: Zero values
var p3 Person // Name: "", Age: 0
```

### Accessing Fields:

```go
p := Person{Name: "Charlie", Age: 35}
fmt.Println(p.Name) // "Charlie"
fmt.Println(p.Age)  // 35

p.Age = 36 // Modify field
```

### Methods on Structs:

```go
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

### Key Features:

- **Value Type**: Structs are value types (copied when assigned)
- **Zero Values**: Each field gets its zero value if not initialized
- **Methods**: Can define methods on struct types
- **Embedding**: Can embed other structs for composition
- **Tags**: Can add metadata using struct tags

### Example with Tags:

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```
