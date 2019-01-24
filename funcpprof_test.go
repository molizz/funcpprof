package funcpprof

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestTickerProfiler(t *testing.T) {
	TickerInterval = 1
	TickerDelay = 2
	ticker := StartTickerProfiler(false)
	defer ticker.Stop()

	time.Sleep(time.Duration(10+rand.Intn(10)) * time.Millisecond)
	cpuWork2()

	profiles := GetProfiles(0)
	// for n, t := range profiles {
	// 	fmt.Println(t, n)
	// }

	if len(profiles) == 0 {
		t.Fatalf("profile count is 0")
	}
}

func cpuWork2() {
	for i := 0; i < 45000000; i++ {
		str := "str" + strconv.Itoa(i)
		str = str + "a"
	}
}
