## Chapter 1 Hello, World!

- [ ] `go run` compiles the source code from one or more source files whose names end in `.go`, links it with libraries, then runs the resulting executable file.
- [ ] package `main` is special. It defines a standalone executable program, not a library.
- [ ] the function `main` is also special -- it's where execution of the program begins. Whatever main does is what the program does.
- [ ] `os.Args` are used to take user input from command line
- [ ] `os.Args[0]` is the name of the command itself
- [ ] `os.Args[1:]` since os.Args is a slice of strings, `[m:n]` means starting from `m` including mth element, to `n` but not including nth element
- [ ] `var s` variables are implicitly initialized to the `zero value`, which is `0` for numeric types and `""` for strings
- [ ] string concatenation will create new string each time, use `strings.Builder` instead

```
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

```

- [ ] a `map` holds a set of key/value pairs and provides constant-time operations to store, retrieve, or test for an item in hte set.

  - the value may be of any type whose values can be compared with `==`
  - the order of map iteration

- `%d` decimal integer
- `%x`, `%o`, `%b`, integer in hex, octal, binary
- `%f`, `%g`, `%e` floating-point number
- `%t` boolean
- `%c` rune (Unicode code point)
- `%s` string
- `%q` quoted string "abc" or rune 'c'
- `%v` any vlaue in a natural format
- `%T` type of any value
- `%%` literal percent sign

- [ ] a `map` is a `reference` to the data structure created by `make`. When a map is passed to a function, the function receives a copy of the reference, so any changes the called function makes to the underlying data structure will be visible through the caller's map reference too.

- [ ] `bufio.NewScanner(f)` reads data in a streaming way, however, `ioutil.ReadFile()` reads entire data into memory and return a byte slice.
- [ ] `io.Copy(dst, src)` reads from src and writes to dst. Use it instead of `ioutil.ReadAll` to pipe data from src to dst without requiring a buffer large enough to hold the entire stream
- [ ] the function `main` runs in a goroutine and the go statement creates additional goroutines.
- [ ] behind the scenes, the server runs the handler for each incoming request in a separate goroutine so that it can serve multiple requests simultaneously
- [ ] `mu.Lock()` is used to ensure that at most one goroutine accesses the variable at a time

```go
switch coinflip() {
case "heads":
	heads++;
case "tails":
	tails++;
default:
	return
}
```

- [ ] cases do not fall through from one to the next as in C-like languages (although there is a rarely used `fallthrough` statement)

```go
func Signum(x int) int {
	switch {
		case x > 0:
			return +1
		default:
			return 0
		case x < 0:
			return -1
	}
}
```

- [ ] this form is called a `tagless switch`, it's equivalent to `switch true {}`
- [ ] go provides pointers, that is, values that contain the address of a variable
- [ ] `go doc` tool makes documents easily accessible from the command line
