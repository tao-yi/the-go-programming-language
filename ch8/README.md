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

As with maps, a channel is a reference to the data structure created by make.

```go
ch := make(chan int) // ch has type 'chan int'
```

`close` operation sets a flag indicating that no more values will ever be sent on this channel; subsequent attempts to send will panic.

## Unbuffered Channels

A send operation on an unbuffered channel blocks the sending goroutine until another goroutine executes a corresponding receive on the same channel, at which point the value is transmitted and both goroutines may continue.

Conversely, if the receive operation was attempted first, the receiving goroutine is blocked until another goroutine performs a send on the same channel.

Communication over an unbuffered channel causes the sending and receiving goroutines to synchronize. Because of this, unbuffered channels are sometimes called synchronous channels.

You needn't close every channel when you've finished with it. It's only necessary to close a channel when it is important to tell the receiving goroutines that all data have been sent.

## Buffered Channels

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
