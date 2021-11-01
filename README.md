# SliceX

SliceX provides functional operations on Go slices using Go 1.18 type parameters.

## Get started

### Prerequisites

- Make sure your local `go.mod` declares `go 1.18`
- Go 1.18 not being relaesed yet, I suggest using the wrapper [gotip](https://pkg.go.dev/golang.org/dl/gotip):
  > The gotip command compiles and runs the go command from the development tree.

### Installation

Assuming the above prerequisites are fulfilled:

```sh
gotip get github.com/drykit-go/slicex
```

## Available functions

- `Map`
- `Filter`
- `Reduce`
- `Apply`
- `ApplyUntil`
- `KeysOf`
- `ValuesOf`

## Examples

See [unit tests](./slicex_test.go) for some usage examples.
