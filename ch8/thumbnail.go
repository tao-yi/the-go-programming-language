package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {
	filenames := []string{
		"a.jpeg",
		"b.jpeg",
		"c.jpeg",
		"d.jpeg",
	}

	// makeThumbnails3(filenames)
	// makeThumbnails4(filenames)
	fs, _ := makeThumbnails5(filenames)
	for _, f := range fs {
		fmt.Println(f)
	}
}

// ImageFile reads an image from infile and writes
// a thumbnail-size version of it in the same directoty
// It returns the generated file name, e.g., "foo.thumb.jpg"
func ImageFile(infile string) (string, error) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	return infile, nil
}

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// NOTE: incorrect
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go ImageFile(f)
	}
	// it starts all the goroutine but doesn't wait for them to finish
}

func makeThumbnails3(filenames []string) {
	ch := make(chan string)
	for _, f := range filenames {
		go func(f string) {
			r, _ := ImageFile(f)
			ch <- r
		}(f)
	}

	// wait for goroutines to complete
	for range filenames {
		log.Println(<-ch)
	}
}

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

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
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

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}
