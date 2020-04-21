package sync

import (
	"sync"
)

type WaitGroup interface {
	Add(delta int)
	Done()
	Wait()
}

type SemaphoredWaitGroup struct {
	Size     int
	ch       chan bool
	wg       sync.WaitGroup
	initOnce sync.Once
	addOnce  sync.Once
	doneOnce sync.Once
}

func (s *SemaphoredWaitGroup) Add(delta int) {
	s.initOnce.Do(func() {
		if s.Size == 0 {
			s.Size = 1
		}
		s.ch = make(chan bool, s.Size)
	})

	s.ch <- true
	s.wg.Add(delta)

	s.addOnce.Do(func() {
		s.wg.Add(1)
	})
	s.doneOnce.Do(func() {
		s.wg.Done()
	})
}
func (s *SemaphoredWaitGroup) Done() {
	<-s.ch
	s.wg.Done()

	s.doneOnce.Do(func() {
		s.wg.Done()
	})
}
func (s *SemaphoredWaitGroup) Wait() {
	s.addOnce.Do(func() {
		s.wg.Add(1)
	})

	s.wg.Wait()
	close(s.ch) // disables ruse
}
