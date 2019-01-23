package funcpprof

var profileCollect *ProfileCollect

func init() {
	profileCollect = &ProfileCollect{}
}

func GetProfile() map[string]int64 {
	return profileCollect.profilesMap
}

type ProfileCollect struct {
	profilesMap map[string]int64
	// mu          sync.Mutex
}

func (p *ProfileCollect) AddNewProfile(newProfile map[string]int64) {
	// p.mu.Lock()
	// defer p.mu.Unlock()

	p.profilesMap = newProfile
}
