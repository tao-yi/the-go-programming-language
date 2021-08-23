# Goroutines and Channels

Go enables two sytels of concurrent programming:

- communicating sequential processes or CSP
- shared memory multithreading

## Goroutines

In Go, each concurrently executing activity is called a `goroutine`

When a program starts, its only goroutine is the one that calls the main function, so we call it the `main goroutine`

A `go` statement causes the function to be called in a newly created goroutine.

## Channels

If goroutines are the activities of a concurrent Go program, `channels` are the connections between them. A channel is a communication mechanism that lets one goroutine send values to another goroutine.

As with maps, a channel is a **reference** to the data structure created by make.

- when we copy a channel or pass one as an argument to a function, we are copying a reference, so caller and callee refer to the same data structure. As with other reference types, the zero value of a channel is `nil`

- Two channels of the same type may be compared using `==`. The comparison is true if both are references to the same channel data structure.

- A channel has two principle operations, `send` and `receive`, collectively known as `communications`. A send statement transmits a value from one goroutine, through the channel, to another goroutine exceuting a corresponding receive expression.

```go
ch := make(chan int) // ch has type 'chan int'
```

- Channels support a third operation, `close`, which sets a flag indicating that no more values will ever be sent on this channel; **subsequent attempts to send will panic**. Receive operations on a closed channel yield the values that have been sent until no more values are left; **any receive operations thereafter complete immediately** and yield the zero value of the channel's element type.

## Unbuffered Channels

A send operation on an unbuffered channel **blocks** the sending goroutine until another goroutine executes a corresponding receive on the same channel, at which point the value is transmitted and both goroutines may continue.

Conversely, if the receive operation was attempted first, the receiving goroutine is blocked until another goroutine performs a send on the same channel.

Communication over an unbuffered channel causes the sending and receiving goroutines to synchronize. Because of this, **unbuffered channels are sometimes called synchronous channels.**

**You needn't close every channel when you've finished with it.** It's only necessary to close a channel when it is important to tell the receiving goroutines that all data have been sent.

## Buffered Channels

A buffered channel has a queue of elements. The queue’s maximum size is determined when it is created, by the capacity argument to make.

```go
ch = make(chan string, 3)
```

## Looping In Parallel

```go
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errors <- err
		}(f)
	}

	// When it encounters the first non-nil error,
	// it returns the error to the caller,
	// leaving no goroutine draining the errors channel.
	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE: incorrect: goroutine leak!
		}
	}

	return nil
}

```

## Mutual Exclusion

```go
func Balance() int {
  mu.Lock()
  defer mu.Unlock()
  return balance
}


func Deposit(amount int) {
  mu.Lock()
  defer mu.Unlock()
  balance = balance + amount
}

// NOTE: not atomic!
func Withdraw(amount int) bool {
  Deposit(-amount)
  if Balance() < 0 {
    Deposit(amount)
    return false // insufficient funds
  }
  return true
}
```

```go
// NOTE: incorrect
func Withdraw(amount int) bool {
  // tries to acquire the mutex lock a second time => deadlock
  mu.Lock()
  defer mu.Unlock()
  Deposit(-amount)
  if Balance() < 0 {
    Deposit(amount)
    return false // insufficient funds
  }
  return true
}
```

mutext locks are not re-entrant: it's not possible to lock a mutext that's already locked - this leads to a deadlock where nothing can proceed, and Withdraw blocks forever

A common solution is to divide a function such as Deposit into two: an unexported function, `deposit`, that assumes the lock is already held and does the real work, and an exported function that acquires the lock before calling deposit.

```go
var mu sync.Mutex
var balance int = 0

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

// this function requires that the lock be held
func deposit(amount int) {
	balance += amount
}
```

### Read/Write Mutexes

We need a special kind of lock that allows read-only operations to proceed in parallel with each other, but write operations to have fully exclusive access.

This lock is called a `multiple readers, single writer` lock, and in Go it's provided by `sync.RWMutex`

```go
func Balance() int {
	mu.RLock() // locks for reading
	defer mu.RUnlock()
	return balance
}
```

### Memory Synchronization

Synchronization also affects memory.

In modern computer there may be dozes of processors, each with its own local cache of the main memory.

For efficiency, writes to memory are buffered within each processor and flushed out to main memory only when necessary. They may even be committed to main memory in a different order than they were written by the writing goroutine.

Synchronization primitives like channel communications and mutex operations cause the processor to flush out and commit all its accumulated writes so that the effects of goroutine execution up to that point are guaranteed to be visible to goroutines running on other processor.
