package sync

import (
	"sync"

	"sewik/pkg/sys"
)

type WaitGroup interface {
	Add(delta int)
	Done()
	Wait()
}
type LimitingWaitGroup struct {
	Limit           int
	ch              chan bool
	wg              sync.WaitGroup
	onlyOneTime     sync.Once
	onlyOneTimeAdd  sync.Once
	onlyOneTimeDone sync.Once
}

func (x *LimitingWaitGroup) Add(delta int) {
	x.onlyOneTime.Do(func() {
		if x.Limit == 0 {
			x.Limit = 1
		}
		x.ch = make(chan bool, x.Limit)
		sys.ChUtBo("wgr", x.ch)
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
func (x *LimitingWaitGroup) Done() {
	<-x.ch
	x.wg.Done()

	x.onlyOneTimeDone.Do(func() {
		x.wg.Done()
	})
}
func (x *LimitingWaitGroup) Wait() {
	x.onlyOneTimeAdd.Do(func() {
		x.wg.Add(1)
	})

	x.wg.Wait()
	close(x.ch) //disables ruse
}
