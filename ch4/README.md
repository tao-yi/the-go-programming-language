# Composite Types

- `arrays`
- `slices`
- `maps`
- `structs`

Arrays and structs are aggregate types; their values are concatenations of other values in memory.

- Both arrays and structs are fixed size.
- In contrast, slices and maps are dynamic data structures that grow as values are added.

### Arrays

arrays are rarely used directly in Go. Slices, which can grow and shrink, are much more versatile, but to understand slices we must understand arrays first.

```go
// 如果用...表示长度，那么数组长度会由初始值决定
var arr1 [3]int = [3]int{1, 2, 3}
arr2 := [...]int{1, 2, 3}
```

- **the size of an array is part of its type, so `[3]int` and `[4]int` are different types**

```go
r := [...]int{99: -1}
```

defins an array r with 100 elements, all zero except for the last, which has value `-1`

We may directly compare two arrays of that type using the `==` operator, which reports whether all corresponding elements are equal.

> When a function as called, a copy of each argument value is assigned to the corresponding parameter variable, so the function receives a copy, not the original.

Passing large array in this way can be inefficient, and any changes that the function makes to array elements affect only the copy, not the original.
This behavior is different from languages that implicitly pass arrays by reference.

Of course, we can explicitly pass a pointer to an array so that any modifications the function makes to array elements will be visible to the caller.

### Slices

Slices represent variable-length sequences whose elements all have the same type. `[]T`, it looks like an array type without a size.

A slice has three components:

- a pointer: points to the first element of the array that is reachable through the slice, which is not necessarily the array's first element
- a length
- a capacity

Multiple slices can share the same underlying array and may refer to overlapping parts of that array.

- `copy`: a built-in function which copies elements from one slice to another of the same type.

```go
// copy elements from one slice to another of the same type
copy(dest, src)
```

**usually we don't know whether a given call to append will cause a reallocation, so we can't assume that the original slice refers to the same array as the resulting slice, nor that it refers to a different one.**

```go
// the built-in append let us add more than one new element, or even a whole slice of them
var x []int
x = append(x, 1)
x = append(x, 2, 3)
x = append(x, 4, 5, 6)
x = append(x, x...) // append the slice x
```

a slice can be used to implement a stack.

```go
stack = append(stack, v) // push v

// the top of the stack is the last element
top := stack[len(stack)-1]

stack = stack[:len(stack)-1] // pop
```

## Maps

In Go, a `map` is a reference to a hash table, and a map type is written `map[K]V`

The key type `K` must be comparable using `==`, so that the map ca test whether a given key is equal to one already within it.

Though floating-point numbers are comparable, it's a bad idea to compare floats for equality.

```go
ages := make(map[string]int) // mapping from strings to ints

// or
ages := map[string]int{}

m1["alice"] = 32
delete(m1, "alice") // remove element m1["alice"]
```

a map element is not a variable, one reason that we can't take the address of a map element is that growing a map might cause rehashing of existing elements into new storage locations, thus potentiall invalidating the address.

```go
val := &m1["alice"] // compile error: cannot take address of map element
```

```go
m1["a"] = 1
m1["b"] = 2
for k, v := range m1 {
  fmt.Printf("key: %s, val: %d\n", k, v)
}
```

the order of map iteration is unspecified, and different implementations might use a different hash function, leading to a different ordering.

To enumerate the key/value pairs in order, we must sort the keys explicitly.

```go
names := make([]string, 0, len(ages))
for name := range m1 {
  names = append(names, name)
}

sort.Strings(names)
for _, name := range names {
  fmt.Printf("%s\t%d\n", name, m1[name])
}
```

accessing a map element by subscripting always yields a value, you get the zero value for the element type of not present in the map

```go
v, ok := m1["bob"]
if !ok {
  fmt.Println("bob is not a key in this map, age === 0")
}
```

Subscripting a map in this context yields two values, the second is a boolean that reports whether the element was present.

As with slices, maps cannot be compared to each other; the only legal comparison is with `nil`

```go
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
```

Go does not provide a `set` type, but since the keys of a map are distinct, a map can serve this purpose.

```go
set := make(map[string]bool) // a set of strings
```

### Structs

```go
type Employee struct {
	ID            string
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}
```

For efficiency, larger struct types are usually passed to or returned from functions indirectly using a pointer.

If all the files of a struct are comparable, the struct itself is comparable, so two expressions of that type may be compared using `==` or `!=`

```go
p := Point{x: 1, y: 2}
q := Point{x: 1, y: 2}
// If all the files of a struct are comparable, the struct itself is comparable
fmt.Println(p == q)

c1 := Complex{}
c2 := Complex{}
fmt.Println(c1 == c2) // compile error, (operator == not defined for Complex)

type Point struct {
  x, y int
}

type Complex struct {
  num []int // array is not comparable
  x   int
  y   int
}
```

struct embedding is just a syntactic sugar on the dot notation used to select struct fields.
