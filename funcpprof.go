package funcpprof

import (
	"time"
)

/*
每次监控到的数据应存起来(最大数量)
*/

var MaxProfiles = 400
var profilesData *ProfilesData

func init() {
	profilesData = &ProfilesData{
		list: &QueueList{
			maxQueue: MaxProfiles,
		},
	}
}

func GetProfiles(stamp int64) []*Profile {
	profiles := make([]*Profile, 0)
	profilesData.list.Each(func(v interface{}) {
		pro, ok := v.(*Profile)
		if !ok {
			return
		}
		if pro.stamp > stamp {
			profiles = append(profiles, pro)
		}
	})
	return profiles
}

type Profile struct {
	profilesSet map[string]int64
	stamp       int64
}

type ProfilesData struct {
	list *QueueList
}

func (p *ProfilesData) AddNewProfile(newProfile map[string]int64) {
	profile := &Profile{
		profilesSet: newProfile,
		stamp:       time.Now().Unix(),
	}
	p.list.Push(profile)
}
