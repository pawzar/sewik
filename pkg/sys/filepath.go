package sys

import (
	"log"
	"path/filepath"
)

func Filenames(p []string, size int) <-chan string {
	jobs := make(chan string, size)

	go func() {
		log.Printf("[START] Filenames: %q", p)
		for _, pattern := range p {
			filenames, _ := filepath.Glob(pattern)
			for _, filename := range filenames {
				jobs <- filename
			}
		}
		close(jobs)
		log.Printf("[DONE] Filenames.")
	}()

	return jobs
}
