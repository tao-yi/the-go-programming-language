# Methods

```go
type Point struct {
  x, y float64
}

func (p Point) Distance(q Point) float64 {}
```

- `p` is called the method's `receiver`
- `p.Distance` is called a `selector`

Method may be declared on any named type defined in the same package, so long as its underlying type is netier a pointer nor an interface

```go
type Name string

func (n Name) Change() {
}

type Unvalid *string

func (u Unvalid) A() { // compile error: invalid receiver Unvalid (pointer or interface type)compiler
}
```

### Methods with a Pointer Receiver

Because calling a function makes a copy of each argument value, if a function needs to update a variable, or if an argument is so large that we wish to avoid copying, we must pass the address of the variable using a pointer. The same goes for methods that need to update the receiver variable: we attach them to a pointer type, such as `*Pointer`.

If the receiver p is a variable of type Point but the method requires a \*Point receiver, we can use this shorthand:

```go
func (p *Point) ScaleBy(factor float64) { }

// the compiler will perform an implicit &p on the variable
p.ScaleBy(2)
```

- if the receiver argument is a `variable` of type `T` and the receiver parameter has type `*T`, the compiler implicitly takes the address of the variable
- if the receiver argument has the type `*T` and the receiver parameter has type `T`, the compiler implicitly dereferences the receiver

### Nil is a valid receiver value

### Method values and expressions

```go
p := Point{1,2}
q := Point{4,6}
distanceFromP := p.Distance // method value
fmt.Println(distanceFromP(q))
```

the selector `p.Distance` yields a `method value`, a function that binds a method (Point.Distance) to a specific receiver value p.

```go
type Rocket struct { /* ... */ }
func (r *Rocket) Launch() { /* ... */ }

r := new Rocket()
time.AfterFunc(10 * time.Second, r.Launch)
```
