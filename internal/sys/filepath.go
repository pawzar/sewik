package sys

import (
	"log"
	"path/filepath"
)

func Filenames(p []string, size int) <-chan string {
	jobs := make(chan string, size)

	go func() {
		defer close(jobs)
		log.Printf("[START] Filenames: %q", p)
		defer log.Printf("[DONE] Filenames.")

		for _, pattern := range p {
			filenames, _ := filepath.Glob(pattern)
			for _, filename := range filenames {
				jobs <- filename
			}
		}
	}()

	return jobs
}
