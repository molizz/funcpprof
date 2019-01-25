package funcpprof

import (
	"fmt"

	"github.com/molizz/funcpprof/profiler"
)

func StartProfiler() error {
	cpuProfiler := profiler.CPUProfiler{}
	err := cpuProfiler.StartProfiler()
	if err != nil {
		return err
	}
	return nil
}

func StopProfiler(ignoreRuntime bool) error {
	cpuProfiler := profiler.CPUProfiler{}
	profile, err := cpuProfiler.StopProfiler()
	if err != nil {
		return err
	}

	profileMap, err := Parse(profile, ignoreRuntime)
	if err != nil {
		return fmt.Errorf("parse profile is err: %v", err)
	}

	profilesData.AddNewProfile(profileMap)
	return nil
}
