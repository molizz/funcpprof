package profiler

import (
	"bufio"
	"bytes"
	"runtime/pprof"
	"time"

	"github.com/google/pprof/profile"
)

type CPUProfiler struct {
	pprofBuff   *bytes.Buffer
	pprofWriter *bufio.Writer
	startNano   int64
	stopNano    int64
}

func (p *CPUProfiler) StartProfile() error {
	p.pprofBuff = &bytes.Buffer{}
	p.pprofWriter = bufio.NewWriter(p.pprofBuff)
	p.startNano = time.Now().UnixNano()

	err := pprof.StartCPUProfile(p.pprofWriter)
	if err != nil {
		return err
	}
	return nil
}

func (p *CPUProfiler) StopProfile() (*profile.Profile, error) {
	pprof.StopCPUProfile()
	p.stopNano = time.Now().UnixNano()
	err := p.pprofWriter.Flush()
	if err != nil {
		return nil, err
	}

	if pf, err := profile.Parse(p.pprofBuff); err == nil {
		return pf, nil
	} else {
		return nil, err
	}
}
