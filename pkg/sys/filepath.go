package sys

import (
	"path/filepath"
)

func Filenames(p []string) <-chan string {
	jobs := make(chan string)

	go func() {
		for _, pattern := range p {
			filenames, _ := filepath.Glob(pattern)
			for _, filename := range filenames {
				jobs <- filename
			}
		}
		close(jobs)
	}()

	return jobs
}
