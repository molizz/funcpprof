package funcpprof

import (
	"fmt"
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

func GetProfiles(stamp int64) Profiles {
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
	return Profiles(profiles)
}

type Profiling struct {
	stacks []*FunctionStack
	stamp  int64
}

type Profiles []*Profiling

func (p Profiles) Each(fn func(string)) {
	for _, pf := range p {
		for _, s := range pf.stacks {
			f := fmt.Sprintf("%d\t%s", s.Duration, s.Name)
			fn(f)
		}
	}
}

func (p Profiles) Flush(path string) error {
	return NewWrite(p).Flush(path)
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
