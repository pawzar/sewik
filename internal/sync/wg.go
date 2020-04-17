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
	Size            int
	ch              chan bool
	wg              sync.WaitGroup
	onlyOneTime     sync.Once
	onlyOneTimeAdd  sync.Once
	onlyOneTimeDone sync.Once
}

func (x *SemaphoredWaitGroup) Add(delta int) {
	x.onlyOneTime.Do(func() {
		if x.Size == 0 {
			x.Size = 1
		}
		x.ch = make(chan bool, x.Size)
	})

	x.ch <- true
	x.wg.Add(delta)

	x.onlyOneTimeAdd.Do(func() {
		x.wg.Add(1)
	})
	x.onlyOneTimeDone.Do(func() {
		x.wg.Done()
	})
}
func (x *SemaphoredWaitGroup) Done() {
	<-x.ch
	x.wg.Done()

	x.onlyOneTimeDone.Do(func() {
		x.wg.Done()
	})
}
func (x *SemaphoredWaitGroup) Wait() {
	x.onlyOneTimeAdd.Do(func() {
		x.wg.Add(1)
	})

	x.wg.Wait()
	close(x.ch) // disables ruse
}
