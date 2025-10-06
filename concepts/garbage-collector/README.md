# Garbage Collector

## Three Color Marking Algorithm

The three color marking algorithm starts with gray root objects and all white objects on the heap. Each object has references to an array of objects.

### Process:

1. **Gray Root Objects**: The garbage collector starts with gray root objects
2. **White Objects**: All other objects on the heap start as white
3. **Traversal**: The garbage collector picks up a gray object and traverses through all its references
4. **Color Change**: Once those objects are processed, they change their color to gray
5. **Black Objects**: Processed objects become black
6. **Continuation**: The process continues until there are no more gray objects
7. **Collection**: White objects that remain are unreachable and are collected by the garbage collector

### Object States:
- **Gray**: Objects being processed in the future
- **Black**: Objects that are reachable and have been processed
- **White**: Unreachable objects that will be collected
