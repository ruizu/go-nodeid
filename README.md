# go-nodeid

`go-nodeid` is a lightweight Go module for generating unique identifiers based on the first non-loopback IPv4 address of the machine. This can be useful for identifying nodes in a distributed system or other scenarios where a machine-specific identifier is needed.

## Features

- Retrieves a unique identifier based on the machine's IPv4 address.
- Provides both error-handling (`Get`) and panic-based (`MustGet`) retrieval methods.
- Includes comprehensive test coverage.

## Installation

To use `go-nodeid` in your project, add it as a dependency:

```bash
go get github.com/ruizu/go-nodeid
```

## Usage

### Import the package

```go
import "github.com/ruizu/go-nodeid"
```

### Retrieve the Node ID

#### Using `Get`

`Get` retrieves the unique identifier and returns an error if no valid IPv4 address is found:

```go
id, err := nodeid.Get()
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Node ID:", id)
}
```

#### Using `MustGet`

`MustGet` retrieves the unique identifier and panics if it fails:

```go
id := nodeid.MustGet()
fmt.Println("Node ID:", id)
```

## License

This project is licensed under the MIT License. See the [LICENSE.md](LICENSE.md) file for details.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the module.

## Testing

To run the tests, use the following command:

```bash
go test ./...
```