package funcpprof

import (
	"time"
)

/*
每次监控到的数据应存起来(最大数量)
*/

var MaxProfiles = 400
var stackData *StackData

func init() {
	stackData = &StackData{
		list: NewQueueList(MaxProfiles),
	}
}

func GetProfiles(stamp int64) []*Profiling {
	profiles := make([]*Profiling, 0)
	stackData.list.Each(func(v interface{}) {
		pro, ok := v.(*Profiling)
		if !ok {
			return
		}
		if pro.stamp > stamp {
			profiles = append(profiles, pro)
		}
	})
	return profiles
}

type Profiling struct {
	stacks []*FunctionStack
	stamp  int64
}

type StackData struct {
	list *QueueList
}

func (p *StackData) AddNewStack(newStack []*FunctionStack) {
	profile := &Profiling{
		stacks: newStack,
		stamp:  time.Now().Unix(),
	}
	p.list.Push(profile)
}
