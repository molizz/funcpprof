package funcpprof

import (
	"fmt"

	"github.com/molizz/funcpprof/profiler"
)

func StartProfile() error {
	cpuProfiler := profiler.CPUProfiler{}
	err := cpuProfiler.StartProfile()
	if err != nil {
		return err
	}
	return nil
}

func StopProfile(ignoreRuntime bool) error {
	cpuProfiler := profiler.CPUProfiler{}
	profile, err := cpuProfiler.StopProfile()
	if err != nil {
		return err
	}

	stackNodes, err := Parse(profile, ignoreRuntime)
	if err != nil {
		return fmt.Errorf("parse profile is err: %v", err)
	}

	stackData.AddNewStack(stackNodes)
	return nil
}
