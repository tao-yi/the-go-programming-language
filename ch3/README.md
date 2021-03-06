# Basic Data Types

Go's types fall into four categories:

- `basic types`: numbers, strings, booleans
- `aggregate types`: arrays, structs
- `reference types`: pointers, slices, maps, functions, channels
- `interface types`: interface

## Integers

- The type `rune` is a synonym for `int32` and conventionally indicates that a value is a Unicode code point.
- the type `byte` is an synonym for `unit8` and emphasizes that the value is a piece of raw data rather than a small numeric quantity

- `uintptr`: unsigned interger type whose width is not specified but is sufficient to hold all the bits of a pointer value.
  - is used only for low level programming such as at the boundary of a Go program with a C library or an operating system.

## Boolean

there is no implicit conversion from a boolean value to a numeric value like 0 or 1, or vice versa.

It might be worth writing a conversion function if this operation were needed often:

```go
// btoi returns 1 if b is true and 0 if false
func btoi(b bool) int {
  if b {
    return 1
  }
  return 0
}
```

## Strings

A string is an immutable sequence of bytes. Text strings are conventionally interpreted as UTF-8-encoded sequences of Unicode code points (runes).

The built-in `len` function returns the number of bytes (not runes) in a string, and the `index` operation `s[i]` retrieves the i-th byte of string `s`, where `0 <= i <= len(s)`

The `i-th` byte of a string is not necessarily the `i-th` character of a string, because the UTF-8 encoding of a non-ASCII code point requires two or more bytes.

A string `s` and a substring likbe `s[7:]` may safely safely share the same data, so the substring operation is also cheap. No new memory is allocated in either case.

In go, [a string is in effect a read-only slice of bytes](https://blog.golang.org/strings).

A string value can be written as a `string literal`, a sequence of bytes enclosed in double quotes:

```go
"Hello, 世界 "
```

A `raw string literal` is written using backquotes instead of doubgle quotes.

```go
const GoUsage  = `Go is a tool for managing Go source code.

Usage:
  go command [arguments]
`
```

### UTF-8

UTF-8 is a variable length encoding of Unicode code points as bytes. It uses between 1 and 4 bytes to represent each rune, but only 1 byte for ASCII characters, and only 2 or 3 bytes for most runes in common use.

## Rune & Unicode

Unicode collects all of the characters in all of the world's writing systems, and assigns each one a standard number called `Unicode code point` or in Go terminology `rune`

The natural data type to hold a single rune is `int32`, and that's what Go uses, it has the synonym `rune` for precisely this purpose.

### Strings and Byte Slices

four standard packages are particularly important for manipulating strings:

- `bytes`: because strings are immutable, building up strings incrementally can involve a lot of allocation and copying. In such cases, it's more efficient to use the `bytes.Buffer` type.
- `strings`
- `strconv`: convert bool, integer, float to and from their string representations.
- `unicode`: `IsDigit`, `IsLetter`, `IsUpper`, `IsLower` for classifying runes.

### Constants

constants are expressions whose value is known to the compiler and whos evaluation is guaranteed to occur at compile time, not at run time.

The underlying type of every constant is a basic type: `boolean`, `string` or `number`

Many computations on constants can be completely evaludated at compile time, reducing the work necessary at run time and enabling other compiler optimizations.

The results of all arithmetic, logical, and comparison operations applied to constant operands are themselves constants, as are the results of conversions and calls to certain built-in functions such as `len`, `cap`, `real`, `imag`, `unsafe.Sizeof`
