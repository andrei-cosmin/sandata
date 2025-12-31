# sandata

A collection of generic data structures for Go.

## Install

```bash
go get github.com/andrei-cosmin/sandata
```

## Packages

| Package    | Description                                      |
|------------|--------------------------------------------------|
| `array`    | Auto-growing array with bitmask clearing         |
| `bit`      | Read-only bitmask wrapper with set operations    |
| `chain`    | Double linked list nodes                         |
| `flag`     | Simple boolean flag                              |
| `pool`     | Fixed-capacity stack pool                        |
| `set`      | Generic set with union, difference, intersection |
| `trie`     | Prefix trie with iterator support                |
| `mathutil` | Math utilities (next power of two)               |

## Usage

```go
import "github.com/andrei-cosmin/sandata/set"

s := set.New[string](10)
s.Insert("foo")
s.Insert("bar")
s.Has("foo") // true
```

## License

MIT
