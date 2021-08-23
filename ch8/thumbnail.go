package main

import (
	"log"
	"math/rand"
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
	makeThumbnails4(filenames)
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
