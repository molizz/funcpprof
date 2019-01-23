package funcpprof

import (
	"errors"
	"fmt"
	"time"

	"github.com/molizz/funcpprof/profiler"
)

var ticker *TickerProfiler

var (
	TickerInterval = 100 //s, 每隔n秒进行一次采集
	TickerDelay    = 5   //s, 每次采集多久
)

type TickerProfiler struct {
	ticker        *Ticker
	ignoreRuntime bool
}

func (f *TickerProfiler) tickerFunc() {
	cpuProfiler := profiler.CPUProfiler{}
	err := cpuProfiler.StartProfiler()
	if err != nil {
		fmt.Println(fmt.Errorf("start profiler is err: %v", err))
		return
	}

	timer := time.NewTimer(time.Duration(TickerDelay) * time.Second)
	<-timer.C

	profile, err := cpuProfiler.StopProfiler()
	if err != nil {
		fmt.Println(fmt.Errorf("stop profiler is err: %v", err))
		return
	}
	fmt.Println("profile .. ", len(profile.Sample))
	profileMap, err := Parse(profile, f.ignoreRuntime)
	if err != nil {
		fmt.Println(fmt.Errorf("parse profile is err: %v", err))
		return
	}

	profileCollect.AddNewProfile(profileMap)
}

func (f *TickerProfiler) Stop() {
	f.ticker.Stop()
}

func (f *TickerProfiler) validate() error {
	if TickerInterval <= 0 {
		return errors.New("invalid ticker interval, must >= 1")
	}
	if TickerDelay <= 0 {
		return errors.New("invalid ticker delay, must >= 1")
	}
	return nil
}

func StartTickerProfiler(ignoreRuntime bool) *TickerProfiler {
	ticker = &TickerProfiler{}
	ticker.ticker = NewTicker(TickerInterval, ticker.tickerFunc)
	ticker.ignoreRuntime = ignoreRuntime
	if err := ticker.validate(); err != nil {
		panic(err)
	}
	ticker.ticker.Start()
	return ticker
}
