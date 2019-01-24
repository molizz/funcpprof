package funcpprof

import (
	"testing"
)

func TestNewQueueList(t *testing.T) {
	l := NewQueueList(10)
	for i := 0; i < 20; i++ {
		l.Push(i)
	}

	if l.list.Len() != 10 {
		t.Fatalf("list len is %d", l.list.Len())
	}
}
