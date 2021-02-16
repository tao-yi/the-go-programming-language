# Functions

Arguments are passed `by value`, so the function receives a copy of each arguments.

However, if the argument contains some kind of reference, like a `pointer`, `slice`, `map`, `function`, or `channel`, then the caller may be affected by any modifications the function makes to variables indirectly referred to by the argument.

You may occasionally encounter a function declaration without a body, indicating that the function is implemented in a language other than Go. Such a declaration defines the function signature.

```go
package main

func Sin(x float64) float64
```

Many Programming language implementations use a fixed-size function call stack; sizes from 64KB to 2MB are typical. Fixed-size stacks impose a limit on the depth of recursion, so one must be careful to avoid a `stack overflow` when traversing large data structures recursively.

In contrast, typical Go implementations use variable-size stacks that start small and grow as needed up to a limit on the order of a gigabyte. This lets us use recursion safely and without worrying about overflow.

In a functionn with named results, the operands of a return statement may be omitted. This is called a `bare return`

### Error-handling strategies

We use `fmt.Errorf` to build descriptive errors by successfively prefixing additional context information to the original error message.

When the error is ultimately handled by the program's main function, it should provide a clear causal chain from the root problem to the overall failure.

Because error messages are frequently chained together, message strings should not be capitalized and newlines should be avoided.

For errors that represent transient or unpredictable problems, it may make sense to retry the failed operation, possibly with a delay between tries, and perhaps with a limit on the number of attempts or the time spent trying before giving up entirely.

### Function Values

Functions are first-class values in Go: like other values, function values have types, and they may be assigned to variables or passed to or returned from functions.

The zero vlaue of a function type is `nil`. Calling a nil function value causes a panic

```go
func squares() func() int {
  var int
  return func() int {
    x++
    return x*x
  }
}
```

The above example demonstrates that function values are not just code but can have state.
The anonymous inner function can access and update the local variables of the enclosing function squares. These hidden variable references are why we classify functions as reference types and why function values are not comparable.

Function values like these are implemented using a technique called `closures`, and Go programmers often use this term for function values.

### variadic functions

```go
func sum(vals ...int) int {
  total := 0
  for _, val := range vals {
    total += val
  }
  return total
}
```

```go
func f(...int) {}
func g([]int) {}

// f和g的参数是不同的类型
fmt.Printf("%T\n", f) // func(...int)
fmt.Printf("%T\n", g) // func([]int)
```

### Deferred Function Calls

a `defer` statement is an ordinary function or method call prefixed by the keyword `defer`. The function and argument expressions are evaluated when the statement is executed, but the actual call is `deferred` until the function that contains the defer statement has finished, whether normally, by executing a return statement or falling off the end or abnormally by panicking.

Any number of calls may be deferred; they are executed in the reverse of the order in which they were deferred.

```go

func main() {
	defer say("one")
	defer say("two")
	defer say("three")
	fmt.Println("main end")
}

func say(word string) {
	fmt.Printf("%s\n", word)
}

// execution order:
// main end
// three
// two
// one
```

可以用来释放锁, unlock a mutex

```go
var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
  mu.Lock()
  defer mu.Unlock()
  return m[key]
}
```

### Panic

Go's type system catches many mistakes at compile time, but others, like an out-of-bounds array access or nil pointer dereference, require checks at run time.
When the Go runtime detects these mistakes, it `panics`

> A panic is ofhten the best thing to do when some "impossible" situation happens, e.g. execution reaches a case that logically can't happen

```go
switch s := suit(drawCard()); s {
  case "Spades": // ...
  case "Hearts": // ...
  case "Diamonds": // ...
  case "Clubs": // ...
  default:
    panic(fmt.Sprintf("invalid suit %q", s))
}
```

It's good practice to assert that the preconditions of a function hold, but this can easily be done to excess. There is not point asserting a condition that the runtime will check for you.

```go
func Reset(x *Buffer) {
  if x == nil {
    panic("x is nil") // unnecessary!
  }
  x.elements = nil
}
```

Although Go's panic mechanism resembles exceptions in other languages, the situations in which panic is used are quite different.

### Recover

A web server that encounters an unexpected problem could close the connection rather than leave the client hanging, and during development, it might report the error to the client too.

If the built-in `recover` function is called within a deferred function and the function containing the defer statement is panicking, recover ends the current state of panic and returns the panic value.

```go
type bailout struct{}

func soleTitle(doc *html.Node) (title string, err error) {
	defer func() {
		switch p := recover(); p {
		case nil: // no panic
		case bailout{}: // "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic; carry on panicking
		}
	}()

	// Bail out of recursion if we find more than one non-empty title.
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple title elements
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

```
