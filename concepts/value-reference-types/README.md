# Value Types vs Reference Types in Go

## Can you explain the difference between value types and reference types?

### Value Types

Value types include:
- Structs
- Arrays
- Basic types (int, float, string, bool, etc.)

**Behavior**: When you assign a value type to a new variable or pass it to a function, a **copy** is made.

### Reference Types

Reference types include:
- Slices
- Maps
- Channels
- Pointers
- Functions

**Behavior**: When you assign a reference type to a new variable or pass it to a function, both variables point to the **same underlying data**.

### Example:

```go
// Value type - struct
type Person struct {
    Name string
    Age  int
}

p1 := Person{Name: "Alice", Age: 30}
p2 := p1  // Copy is made
p2.Name = "Bob"  // p1.Name is still "Alice"

// Reference type - slice
slice1 := []int{1, 2, 3}
slice2 := slice1  // Both point to same underlying array
slice2[0] = 999   // slice1[0] is now 999
```

### Key Differences:

| Aspect | Value Types | Reference Types |
|--------|-------------|-----------------|
| Assignment | Creates copy | Shares reference |
| Memory | Each has its own memory | Share underlying memory |
| Modification | Changes don't affect original | Changes affect all references |
| Performance | Higher memory usage | Lower memory usage |
