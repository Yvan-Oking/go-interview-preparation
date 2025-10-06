# Dependency Management in Go

## How do you manage dependencies in a Go project?

Go uses modules to manage dependencies. This provides a modern, reliable way to handle external packages and their versions.

### Key Commands:

1. **Initialize Module**:
   ```bash
   go mod init <module-name>
   ```

2. **Add Dependencies**:
   ```bash
   go get <package-name>
   ```

3. **Update Dependencies**:
   ```bash
   go get -u <package-name>
   ```

4. **Remove Unused Dependencies**:
   ```bash
   go mod tidy
   ```

### Files Created:

- **`go.mod`**: Records module path, Go version, and direct dependencies
- **`go.sum`**: Contains cryptographic checksums for dependency verification

### Example `go.mod`:

```go
module example.com/myproject

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/lib/pq v1.10.9
)
```

### Benefits:

- **Version Control**: Ensures consistent builds across environments
- **Security**: Cryptographic verification of dependencies
- **Reproducibility**: Same dependencies across different machines
- **Minimal**: Only includes necessary dependencies
