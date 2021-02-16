# Interfaces

```go
package io

type Writer interface {
  write(p []byte) (n int, err error)
}
```

- `*os.File` or `*bytes.Buffer` has a method called `Write`

```go
package fmt

type Stringer interface {
  String() string
}
```

```go
package io

type Reader interface {
  Read(p []byte) (n int, err error)
}

type Closer interface {
  Close() error
}

// embedding an interface
type ReadWriter interface {
  Reader
  Writer
}

type ReadWriteCloser interface {
  Reader
  Writer
  Closer
}
```

```go
func main() {
	var w io.Writer
	var rwc io.ReadWriteCloser

	rwc = os.Stdout
	rwc.Close()
	rwc.Read()
	rwc.Write()

	w = rwc
	w.Write()
}
```

### Interface Values

Conceptually, a value of an interface type, or interface value, has two components, a concrete type and a value of that type. These are called the interface’s `dynamic type` and `dynamic value`.

For a statically typed language like Go, types are a compile-time concept, so a type is not a value.

In Go, variables are always initialized to a well-defined value, and interfaces are no exception.

```go
// the zero value for an interface has both its type and value components set to nil
var w io.Writer
// this assignment involves an implicit conversion from a concrete type to an interface type
// and is equivalent to the explicit conversion `io.Writer(os.Stdout)`
w = os.Stdout
```

In general, we cannot know at compile time what the dynamic type of an interface value will be, so a call through an interface must use `dynamic dispatch`.
Instead of a direct call, the compiler must generate code to obtain the address of the method named `Write` from the type descriptor, then make an indirect call to that address.

Interface values may be compared using `==` and `!=`. Two interface values are equal if both are nil, or if their dynamic tpyes are identical and their dynamic values are equal according to the usual behavior of `==` for that type

### An Interface Containing a Nil Pointer Is Non-nil

A nil interface value, which contains no value at all, is not the same as an interface value containing a pointer that happens to be nil.

```go
const debug = false

func main() {
  // default value: nil <*bytes.Buffer>
	var buf *bytes.Buffer

	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}

	f(buf) // note: subtly wrong
	if debug {
		fmt.Println(buf.String())
	}
}

func f(out io.Writer) {
  // out: { data:nil <*bytes.Buffer> }
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
```

When main calls `f`, it assigns a `nil` pointer of type `*bytes.Buffer` to the out parameter, so the dynamic value of out is nil. However, its dynamic type is `*bytes.Buffer`, meaning that out is a non-nil interface containing a nil pointer value.

The solution is to change the type of `buf` in main to `io.Writer`, thereby avoiding the assignment of the `dysfunctional` value to the interface in the first place:

```go
var buf io.Writer
if debug {
  // ...
}
f(buf) // ok
```

### The error interface

```go
type error interface{
  Error() string
}
```

the simplest way to create an error:

```go
errors.New("")
```

the `syscall` package provides Go's low-level system call API.

```go
func main() {
	var err error
	err = errors.New("hello world!")
	fmt.Println("%v", err)
	err = syscall.Errno(1)
	fmt.Println("%v", err)
}
```

### Type Assertions

A type assertion checks that the dynamic type of its operand matches the asserted type.

There are two possibilities:

1. if the asserted type `T` is a concrete type, then the type assertion checks whether x's dynamic type is `identical` to T

```go
var w io.Writer
w = os.Stdout
f := w.(*os.File) // success: f == os.Stdout
c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
```

2. if instead the asserted type `T` is an interface type, then the type assertion checks whether x's dynamic type satisfies T,

### type switches

```go
package os

type PathError struct {
	Op string
	Path string
	Err error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.path + ": " + e.Err.Error()
}
```

```go
func sqlQuoteV2(x interface{}) string {
	switch x.(type) {
	case nil:
	case int, uint:
	case bool:
	case string:
	default:
	}
}
```
