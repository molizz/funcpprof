package funcpprof

/*
容量有限的队列
*/
import (
	"container/list"
	"sync"
)

func NewQueueList(max int) *QueueList {
	return &QueueList{
		maxQueue: max,
		list:     list.New(),
	}
}

type QueueList struct {
	maxQueue int
	list     *list.List
	muList   sync.Mutex
}

func (q *QueueList) Push(e interface{}) {
	q.muList.Lock()
	defer q.muList.Unlock()

	q.list.PushBack(e)

	for {
		if q.list.Len() > q.maxQueue {
			q.pop()
		} else {
			break
		}
	}
}

func (q *QueueList) Each(fn func(v interface{})) {
	for e := q.list.Front(); e != nil; e = e.Next() {
		fn(e.Value)
	}
}

func (q *QueueList) pop() interface{} {
	front := q.list.Front()
	if front != nil {
		result := q.list.Remove(front)
		return result
	} else {
		return nil
	}
}

func (q *QueueList) Pop() interface{} {
	q.muList.Lock()
	defer q.muList.Unlock()

	return q.pop()
}
