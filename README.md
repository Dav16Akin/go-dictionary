# go-dictionary

A small, idiomatic Go package that provides simple dictionary (map-like) utilities and helpers for working with words and definitions. Designed for learning, small projects, and as a starting point for building more advanced dictionary-based tools.

## Features

- Simple, type-safe Go package
- Basic operations: add, get, delete, update
- Lookup helpers for definitions and synonyms (if provided)
- Easy to extend and test

## Installation

Install the package with go get:

    go get github.com/Dav16Akin/go-dictionary

Or use Go modules (recommended):

    go get github.com/Dav16Akin/go-dictionary@latest

## Quick start

Import the package and use it:

```go
package main

import (
    "fmt"
    "github.com/Dav16Akin/go-dictionary/dictionary"
)

func main() {
    d := dictionary.New()

    d.Add("gopher", "The Go mascot and an affectionate name for Go programmers.")
    def, ok := d.Get("gopher")
    if ok {
        fmt.Println("gopher:", def)
    }

    d.Update("gopher", "A friendly mascot and common term for Go developers.")
    d.Delete("obsolete")
}
```

Adjust the import path above to match the package structure in this repo if different.

## API

This project exposes a small, focused API (examples):

- New() *Dictionary — create a new dictionary instance
- (d *Dictionary) Add(key, value string) error — add a new entry
- (d *Dictionary) Get(key string) (string, bool) — retrieve an entry
- (d *Dictionary) Update(key, value string) error — update an entry
- (d *Dictionary) Delete(key string) error — delete an entry

Refer to the source code for exact function signatures and additional helpers.

## Tests

Run unit tests with:

    go test ./...

Add tests for new features and edge cases.

## Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a branch for your feature or bugfix: `git checkout -b feature/name`
3. Make your changes and add tests
4. Run `go test ./...` and ensure the test suite passes
5. Open a pull request describing your changes

Please keep changes small and focused and follow Go idioms.

## License

This project is provided under the MIT License. See LICENSE for details.

## Contact

Created by Dav16Akin. Feel free to open issues or pull requests on GitHub: https://github.com/Dav16Akin/go-dictionary
