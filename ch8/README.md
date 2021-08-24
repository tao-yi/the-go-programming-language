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

下面的代码有一个 bug，当遇到第一个 error 的时候，整个函数会 return，导致 errors channel 没有被消费干净。而其他 goroutine 还在执行，会往 errors channel 中发送数据，`errors <- err`，会导致这些 goroutine 永远阻塞。

```go
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			// each remaining goroutine will block forever
			// when it tries to send a value on that channel
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

解决这个问题最简单的方法是用 buffered channel。
The simplest solution is to use a buffered channel with sufficient capacity that no worker goroutine will block when it sends a message. (An alternative solution is to create another goroutine to drain the channel while the main goroutine returns the first error without delay)

The next version of makeThumbnails uses a buffered channel to return the names of the generated image files along with any errors.

```go
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	// 这里用buffered channel，确保goroutine往里发送数据时不会阻塞
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			// ch是buffered，所以这里不会阻塞
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

```

## Multiplexing with `select

The `time.Tick` function returns a channel on which it sends events periodically, acting like a metronome. The value of each event is a timestamp.

```go
func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}

	fmt.Println("launching...")
}

```

```go
func main() {
	abort := make(chan struct{})
	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}

	fmt.Println("launching...")
}

```

The `time.Tick` function behaves as if it creates a goroutine that calls `time.Sleep` in a loop, sending an event each time it wakes up. When the countdown function above returns, it stops receiving events from the tick, but the ticker goroutine is still there, trying in vain to send on a channel from which no goroutine is receiving -- a `goroutine leak`.

The `Tick` function is convenient, but it's appropriate only when the ticks will be needed through the lifetime of the application. Otherwise, we should use this pattern.

```go
ticker := time.NewTicker(1 * time.Second)
<- ticker.C // receive from the ticker's channel
ticker.Stop() // cause the ticker's goroutine to terminate
```

We have to use a counting semaphore to prevent from opening too many files at once.

```go
var sema = make(chan struct{}, 4)

// dirents returns the entries of directory dir
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{} // acquire token
	defer func() {
		<-sema // release token
	}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3: %v\n", err)
		return nil
	}
	return entries
}

```

## Cancellation

Sometimes we need to instruct a goroutine to stop what it is doing. For example, in a web server performing a computation on behalf of a client that hs disconnected.

There is no way for one goroutine to terminate another directly.

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
