package funcpprof

import (
	"sync"
	"time"
)

type Ticker struct {
	ticker     *time.Ticker  // 定时器
	tickerFunc func()        // 定时回调方法
	done       chan struct{} // 关闭
	doneOnce   sync.Once
}

func NewTicker(interval int, f func()) *Ticker {
	if interval < 0 {
		interval = 100 //s
	}
	if f == nil {
		panic("请输入回调方法")
	}
	return &Ticker{
		tickerFunc: f,
		done:       make(chan struct{}, 0),
		ticker:     time.NewTicker(time.Duration(interval) * time.Second),
	}
}

func (t *Ticker) Start() {
	go func() {
		done := false
		for !done {
			select {
			case <-t.done:
				done = true
				break
			case <-t.ticker.C:
				t.tickerFunc()
			}
		}
	}()
}

func (t *Ticker) Stop() {
	t.doneOnce.Do(func() {
		t.ticker.Stop()
		close(t.done)
	})
}
