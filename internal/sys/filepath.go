package sys

import (
	"log"
	"path/filepath"
)

func Filenames(patterns []string, size int) <-chan string {
	ch := make(chan string, size)

	go func() {
		defer close(ch)
		log.Printf("[START] Filenames: %q", patterns)
		defer log.Printf("[DONE] Filenames.")

		for _, pattern := range patterns {
			filenames, _ := filepath.Glob(pattern)
			for _, filename := range filenames {
				ch <- filename
			}
		}
	}()

	return ch
}
