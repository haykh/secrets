# Secrets Package and CLI

## Using the CLI

### Build
First, build the application with:

```go build -o ./goencode cmd/cli.go```

### Commands:
To execute:
```./goencode```

This will show all commands available via the command line interface.
Supported commands are:
- GET
- SET
- LIST
- DELETE

#### Example usage:
```bash
./goencode set "key" "value" -k "encoding-key"

./goencode get "key" -k "encoding-key"

./goencode list -k "encoding-key"

./goencode delete "key" -k "encoding-key
```

## Using the Package
Developers can use the package by doing the following:

```go get github.com/haykh/goencode```

Usage of the package in Go code:
```go
v := secret.File("encoding-key", "path/to/file")

err := v.Set("key", "value")

value, err := v.Get("key")

keys, err := v.List()

err := v.Delete("key")
```

#

*Project extended from Jon Calhoun's Gophercises Secrets API and CLI.*