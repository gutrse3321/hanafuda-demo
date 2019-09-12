package tools

import (
	"sync"
	"time"
)

type Tools struct {
}

var (
	Tool = New()
	once sync.Once
)

func New() (t *Tools) {
	once.Do(func() {
		t = &Tools{}
	})
	return t
}

func (t *Tools) Ticker(s int, args ...func()) {
	ch := t.ticker(s, args...)
	time.Sleep(time.Duration(s) * time.Second)
	ch <- true
	close(ch)
}

func (t *Tools) ticker(s int, args ...func()) chan bool {
	ticker := time.NewTicker(time.Second)
	stopChan := make(chan bool)
	go func(ticker *time.Ticker) {
		num := s
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				num--
			case stop := <-stopChan:
				if stop {
					if len(args) != 0 {
						for _, function := range args {
							function()
						}
					}
					return
				}
			}
		}
	}(ticker)
	return stopChan
}
