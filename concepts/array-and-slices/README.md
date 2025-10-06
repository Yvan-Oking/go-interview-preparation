# Arrays and Slices in Go

## What are Arrays and Slices?

Arrays and slices are fundamental data structures in Go for storing collections of elements. Understanding their internal workings is crucial for Go interviews.

### Arrays

Arrays in Go are **fixed-size** collections of elements of the same type.

```go
var arr [5]int                    // Array of 5 integers, zero-initialized
arr := [5]int{1, 2, 3, 4, 5}     // Array with initial values
arr := [...]int{1, 2, 3, 4, 5}   // Compiler determines size
```

**Key Points:**
- Size is part of the type: `[5]int` and `[10]int` are different types
- Arrays are **value types** - copied when assigned
- Zero-initialized by default
- Size must be known at compile time

### Slices

Slices are **dynamic arrays** built on top of arrays. They provide a more flexible way to work with collections.

```go
var slice []int                   // Zero value is nil
slice := make([]int, 5)           // Length 5, capacity 5
slice := make([]int, 5, 10)       // Length 5, capacity 10
slice := []int{1, 2, 3, 4, 5}     // Slice literal
```

## Internal Structure of Slices

**CRITICAL FOR INTERVIEWS:** Understanding slice internals is essential.

A slice is a **struct** with three fields:

```go
type slice struct {
    ptr *T    // Pointer to underlying array
    len int   // Length (number of visible elements)
    cap int   // Capacity (total space in underlying array)
}
```

### Visual Representation:

```
slice := make([]int, 3, 5)

slice = {
    ptr: 0x1040a0c0,  // Points to underlying array
    len: 3,           // 3 visible elements
    cap: 5            // 5 total capacity
}

Underlying array: [1, 2, 3, 0, 0]
                   ↑     ↑
                   len=3 cap=5
```

## Key Concepts for Interviews

### 1. Length vs Capacity

```go
slice := make([]int, 3, 5)
fmt.Println(len(slice))  // 3 - visible elements
fmt.Println(cap(slice))  // 5 - total capacity
```

**Interview Gotcha:** `len()` returns visible elements, NOT the size of the underlying array.

### 2. Slice Operations

```go
// Creating slices
slice := []int{1, 2, 3, 4, 5}

// Slicing (creates new slice header, shares underlying array)
subSlice := slice[1:3]  // [2, 3] - len=2, cap=4

// Appending
slice = append(slice, 6)  // May reallocate if capacity exceeded
```

### 3. Memory Sharing and Gotchas

**DANGER ZONE - Common Interview Traps:**

```go
// Example 1: Shared underlying array
original := []int{1, 2, 3, 4, 5}
slice1 := original[1:3]  // [2, 3]
slice2 := original[2:4]  // [3, 4]

slice1[0] = 999
fmt.Println(original)  // [1, 999, 3, 4, 5] - MODIFIED!
fmt.Println(slice2)    // [999, 4] - MODIFIED!
```

**Why?** Both slices share the same underlying array.

### 4. Append Behavior

```go
slice := make([]int, 3, 5)
slice = append(slice, 4, 5)  // Still fits in capacity
// len=5, cap=5

slice = append(slice, 6)  // Exceeds capacity - REALLOCATION!
// len=6, cap=10 (usually doubles)
```

**Interview Question:** What happens when you append beyond capacity?

**Answer:** Go creates a new underlying array, copies existing elements, and updates the slice header.

### 5. Zero Value and Nil Slices

```go
var slice []int
fmt.Println(slice == nil)  // true
fmt.Println(len(slice))    // 0
fmt.Println(cap(slice))    // 0

// But this is NOT nil:
slice = make([]int, 0)
fmt.Println(slice == nil)  // false
fmt.Println(len(slice))    // 0
fmt.Println(cap(slice))    // 0
```

## Common Interview Questions and Answers

### Q1: What's the difference between arrays and slices?

**Answer:**
- Arrays: Fixed size, value type, size part of type
- Slices: Dynamic size, reference type, built on arrays

### Q2: What happens when you slice a slice?

**Answer:**
- Creates new slice header
- Shares underlying array
- New slice's capacity = original capacity - offset
- Modifications affect all slices sharing the array

### Q3: How does append work internally?

**Answer:**
1. Check if capacity is sufficient
2. If yes: add elements, update length
3. If no: allocate new array (usually 2x capacity), copy elements, add new elements

### Q4: What's the zero value of a slice?

**Answer:** `nil` - length 0, capacity 0, no underlying array

### Q5: Can you modify a slice passed to a function?

**Answer:**
- Yes, you can modify elements
- No, you cannot change the slice header (length, capacity, pointer)
- Use pointer to slice (`*[]T`) to modify the slice itself

## Performance Considerations

### Memory Allocation Patterns

```go
// Inefficient - multiple allocations
var result []int
for i := 0; i < 1000; i++ {
    result = append(result, i)  // May reallocate multiple times
}

// Efficient - pre-allocate
result := make([]int, 0, 1000)  // Pre-allocate capacity
for i := 0; i < 1000; i++ {
    result = append(result, i)  // No reallocations
}
```

### Copy vs Slice

```go
// Slice - shares memory
slice1 := original[1:3]

// Copy - independent memory
slice2 := make([]int, 2)
copy(slice2, original[1:3])
```

## Advanced Topics

### 1. Slice Tricks

```go
// Remove element at index i
slice = append(slice[:i], slice[i+1:]...)

// Insert element at index i
slice = append(slice[:i], append([]T{newElement}, slice[i:]...)...)

// Reverse slice
for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
    slice[i], slice[j] = slice[j], slice[i]
}
```

### 2. Slice as Stack

```go
// Push
slice = append(slice, element)

// Pop
element := slice[len(slice)-1]
slice = slice[:len(slice)-1]
```

### 3. Slice as Queue

```go
// Enqueue
slice = append(slice, element)

// Dequeue
element := slice[0]
slice = slice[1:]
```

## Common Pitfalls and Gotchas

### 1. Modifying Shared Slices

```go
func modifySlice(s []int) {
    s[0] = 999  // Modifies original
    s = append(s, 6)  // Doesn't affect original slice header
}
```

### 2. Range Loop Gotcha

```go
slice := []int{1, 2, 3}
for i, v := range slice {
    v = v * 2  // Doesn't modify slice!
    slice[i] = v * 2  // This modifies slice
}
```

### 3. Capacity After Slicing

```go
original := make([]int, 5, 10)  // len=5, cap=10
sliced := original[1:3]         // len=2, cap=9 (10-1)
```

## Best Practices

1. **Pre-allocate capacity** when you know the size
2. **Use copy()** when you need independent slices
3. **Be careful with slicing** - remember memory sharing
4. **Check for nil slices** before operations
5. **Use make()** for zero-length slices that will grow

## Example: Complete Slice Implementation

```go
package main

import "fmt"

func main() {
    // Create slice
    slice := make([]int, 0, 5)
    fmt.Printf("Initial: len=%d, cap=%d\n", len(slice), cap(slice))
    
    // Append elements
    for i := 1; i <= 7; i++ {
        slice = append(slice, i)
        fmt.Printf("After append %d: len=%d, cap=%d\n", i, len(slice), cap(slice))
    }
    
    // Slice operations
    subSlice := slice[2:5]
    fmt.Printf("Sub-slice: %v, len=%d, cap=%d\n", subSlice, len(subSlice), cap(subSlice))
    
    // Modify sub-slice (affects original)
    subSlice[0] = 999
    fmt.Printf("Original after modification: %v\n", slice)
    
    // Independent copy
    independent := make([]int, len(slice))
    copy(independent, slice)
    independent[0] = 888
    fmt.Printf("Original: %v\n", slice)
    fmt.Printf("Independent: %v\n", independent)
}
```

This comprehensive guide covers all the essential concepts, gotchas, and interview questions about arrays and slices in Go. The internal structure understanding is crucial for technical interviews and will help you explain how slices work under the hood.
