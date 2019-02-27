package funcpprof

import (
	"fmt"
	"sort"
	"strings"

	"github.com/google/pprof/profile"
)

// 传入name, 用户可以自行实现该方法, 并实现过滤掉不需要的数据
// name: like "github.com/xxx/xxx/vendor/github.com/iris-contrib/middleware/logger.(*loggerMiddleware).Serve"
//
var IgnoreFilter func(name string) bool

type FunctionStack struct {
	ID       uint64
	Position int
	Name     string
	Duration int64 // 持续时间
}

func Parse(profile *profile.Profile, ignoreRuntime bool) ([]*FunctionStack, error) {
	var cpuIndex, sampleIndex = -1, -1

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

	functionsMap := make(map[uint64]*FunctionStack)

	/*
		相同的方法在采样中持续出现,那么说明该方法一直存在, 则应该被统计
	*/
	for _, sample := range profile.Sample {
		// stackSample := sample.Value[sampleIndex]
		stackDuration := sample.Value[cpuIndex]

		for i := len(sample.Location) - 1; i >= 0; i-- {
			location := sample.Location[i]
			funcName := parseFuncInfo(location)

			if ignoreRuntime && isRuntimeFunc(funcName) {
				continue
			}
			if IgnoreFilter != nil && IgnoreFilter(funcName) {
				continue
			}

			if len(location.Line) > 0 {
				line := location.Line[0]
				if line.Function != nil {
					id := line.Function.ID
					if _, ok := functionsMap[id]; !ok {
						functionsMap[id] = &FunctionStack{
							Position: i,
							ID:       id,
							Name:     funcName,
							Duration: stackDuration,
						}
					} else {
						functionsMap[id].Duration += stackDuration
					}
				}
			}

		}

	}

	functions := make([]*FunctionStack, 0)
	for _, fn := range functionsMap {
		functions = append(functions, fn)
	}
	sort.Slice(functions, func(i, j int) bool {
		return functions[i].Position > functions[j].Position
	})
	return functions, nil
}

func parseFuncInfo(l *profile.Location) string {
	name, line := funcInfo(l)
	return fmt.Sprintf("%s#%d", name, line)
}

func funcInfo(l *profile.Location) (funcName string, funcLine int64) {
	for _, line := range l.Line {
		if line.Function == nil {
			continue
		}
		name := fmt.Sprintf("%s", line.Function.Name)
		return name, line.Line
	}
	return
}

func isRuntimeFunc(name string) bool {
	builtinNames := []string{
		"golang.org/x",
		"runtime.",
		"syscall.",
		"net.",
		"internal/",
		"database/sql",
	}
	for _, sub := range builtinNames {
		if strings.Index(name, sub) == 0 {
			return true
		}
	}
	return false
}
