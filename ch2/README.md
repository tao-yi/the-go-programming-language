## Chapter 2 Program Strucure

- zero value:
  - `0` for numbers
  - `""` for strings
  - `nil` for interfaces and reference types (slice, pointer, map, channel, function)

### Pointers

- A variable is a piece of storage containing a value
- a pointer value is the address of a variable, a pointer is thus the location a value is stored

```go
func main() {
	x := 1
	addr := &x

	var nullptr *int
	fmt.Printf("%v\n", addr)    // the address of variable x: 0xc0000140a0
	fmt.Printf("%v\n", *addr)   // the value addr points to: 1
	fmt.Printf("%v\n", nullptr) // the zero value for a pointer of any type is <nil>
}
```

### The `new` function

Another way to create a variable is to use the built-in function new.

`new(T)` creates an unnamed variable of type `T`, initializes it to the zero value of `T`, and returns its address, which is a value of type `*T`

```go
func newInt() *int{
  return new(int)
}

// 等价于
func newInt() *int {
  var dummy int
  return &dummy
}
```

### Lifetime of Variables

The lifetime of a package-level variable is the entire execution of the program.

How does the garbage collector know that a variable's storage can be reclaimed?
The basic idea is that every package-level variable and every local variable of each currently active function can potentially be the start or root of a path to the variable in question, following pointers and other kinds of references that ultimately lead to the variable. If no such path exists, the variable has become unreachable, so it can no longer affect the rest of the computation.

**Because the lifetime of a variable is determined only by whether or not it is reachable**, a local variable may outlive a single iteration of the enclosing loop. It may continue to exist even after its enclosing function has returned.

**A compiler may choose to allocate local variables on the heap or on the stack** but, perhaps surprisingly, this choice is not determined by whether `var` or `new` was used to declare the variable.

```go

var global *int

func f() {
  // x must be heap-allocated
	var x int
	x = 1
	global = &x
}
```

`x` must be heap-allocated because it is still reachable from the variable `global` after `f()` has returned, despite being declared as a local variable; we say **x secapes from f**.

```
func g() {
  y := new(int)
  *y = 1
}
```

Conversely, when `g` returns, the variable `y` becomes unreachable and can be recycled. Since `*y` does not escape from `g`, it's safe for the compiler to allocate `*y` on the stack.

Each variable that escapes requires an extra memory allocation.

### Tuple Assignment

`tuple assignments` allows several variables to be assigned at once.

All of the right-hand side expressions are evaluated before any of the variables are updated, making this form most useful when some of the variables appears on both sides of assignment

```go
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

```

For every type `T` there is a corresponding conversion operation `T(x)` that converts the value `x` to type `T`.
A conversion from one type to another is allowed if both have the same underlying type, or if both are unnamed pointer types that point to variables of the same underlying type;

- In any case, a conversion never fails at run time.

**Named types also make it possible to defined new behaviors for values of the type.**

```go
func (c Celsius) String() string {
	return fmt.Sprintf("%g`C", c)
}
```

Many types declare a `String` method of this form because it controls how values of the type appear when printed as a string by the `fmt` package.

### Package initialization

Package initialization begins by initializing package-level variables in the order in which they are declared, except that dependencies are resolved first:

```go
var a = b + c // a initialized third, to 3
var b = f() // b initialized second, to 2, by calling f
var c = 1 // c initialized first, to 1
```

If the package has multiple `.go` files, they are initialized in the order in which the files are given to the compiler; the go tool sorts `.go` files by name before invoking the compiler.

导入的执行顺序

1. 所有的 import 文件会按照 package 名字排序
2. 初始化 package 中的 global variable
3. 执行 package 的 init 函数

### Scope

scope 和 lifetime 是不一样的

- scope 是 compile-time property
- lifetime 是 run-time property

There is a lexical block for the entire source code, called the `universe block`

The declarations of built-in types, functions, and constants like `int`, `len`, and `true` are in the universe block and can be referred to throughout the entire program.

Declarations outside any function, that is, at `package level`, can be referred to from any file in the same package.

when the compiler encounters a reference to a name, it looks for a declaration, starting with the innermost enclosing lexical block and working up to the universe block. If the compiler finds no declaration, it reports an "undeclared name" error.
