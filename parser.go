package funcpprof

import (
	"fmt"
	"strings"

	"github.com/google/pprof/profile"
)

// 传入name, 用户可以自行实现该方法, 并实现过滤掉不需要的数据
// name: like "github.com/xxx/xxx/vendor/github.com/iris-contrib/middleware/logger.(*loggerMiddleware).Serve"
//
var IgnoreFilter func(name string) bool

func Parse(profile *profile.Profile, ignoreRuntime bool) (map[string]int64, error) {
	var cpuIndex, sampleIndex = -1, -1
	var result = make(map[string]int64)

	for i, st := range profile.SampleType {
		if st.Type == "samples" {
			sampleIndex = i
		} else if st.Type == "cpu" {
			cpuIndex = i
		}
	}

	if cpuIndex == -1 || sampleIndex == -1 {
		return nil, fmt.Errorf("invalid profile")
	}

	for _, sample := range profile.Sample {
		runtime := sample.Value[cpuIndex]
		for _, location := range sample.Location {
			name := location.Line[0].Function.Name
			if ignoreRuntime && IsRuntimeFunc(name) {
				continue
			}
			if IgnoreFilter != nil && IgnoreFilter(name) {
				continue
			}

			if _, ok := result[name]; ok {
				result[name] += runtime
			} else {
				result[name] = runtime
			}
		}
	}
	return result, nil
}

func IsRuntimeFunc(name string) bool {
	builtinNames := []string{
		"golang.org/x",
		"runtime.",
		"syscall.",
	}
	for _, sub := range builtinNames {
		if strings.Contains(name, sub) {
			return true
		}
	}
	return false
}
