package Problem0432

import "container/heap"

// AllOne 是解题所需的结构
type AllOne struct {
	m   map[string]*entry
	max maxPQ
	min minPQ
}

// Constructor initialize your data structure here.
func Constructor() AllOne {
	return AllOne{m: make(map[string]*entry),
		max: maxPQ{},
		min: minPQ{},
	}
}

// Inc inserts a new key <Key> with value 1. Or increments an existing key by 1.
func (a *AllOne) Inc(key string) {
	if ep, ok := a.m[key]; ok {
		ep.value++
		heap.Fix(&a.max, ep.maxIndex)
		heap.Fix(&a.min, ep.maxIndex)
	} else {
		ep = &entry{key: key, value: 1}
		a.m[key] = ep
		heap.Push(&a.max, ep)
		heap.Push(&a.min, ep)
	}
}

// Dec decrements an existing key by 1. If Key's value is 1, remove it from the data structure.
func (a *AllOne) Dec(key string) {
	if ep, ok := a.m[key]; !ok {
		return
	} else if ep.value == 1 {
		heap.Remove(&a.max, ep.maxIndex)
		heap.Remove(&a.min, ep.minIndex)
	} else {
		ep.value--
		heap.Fix(&a.max, ep.maxIndex)
		heap.Fix(&a.min, ep.minIndex)
	}
}

// GetMaxKey returns one of the keys with maximal value.
func (a *AllOne) GetMaxKey() string {
	if len(a.max) == 0 {
		return ""
	}
	return a.max[0].key
}

// GetMinKey returns one of the keys with Minimal value.
func (a *AllOne) GetMinKey() string {
	if len(a.min) == 0 {
		return ""
	}
	return a.min[0].key
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */

// entry 是 priorityQueue 中的元素
type entry struct {
	key   string
	value int
	// index 是 item 在 heap 中的索引号
	// item 加入 Priority Queue 后， Priority 会变化时，很有用
	// 如果 item.priority 一直不变的话，可以删除 index
	maxIndex int
	minIndex int
}

// priorityQueue implements heap.Interface and holds items.
type minPQ []*entry

func (pq minPQ) Len() int { return len(pq) }

func (pq minPQ) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq minPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].minIndex = i
	pq[j].minIndex = j
}

// Push 往 pq 中放 item
func (pq *minPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*entry)
	item.minIndex = n
	*pq = append(*pq, item)
}

// Pop 从 pq 中取出最优先的 item
func (pq *minPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.minIndex = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// priorityQueue implements heap.Interface and holds items.
type maxPQ []*entry

func (pq maxPQ) Len() int { return len(pq) }

func (pq maxPQ) Less(i, j int) bool {
	return pq[i].value > pq[j].value
}

func (pq maxPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].maxIndex = i
	pq[j].maxIndex = j
}

// Push 往 pq 中放 item
func (pq *maxPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*entry)
	item.maxIndex = n
	*pq = append(*pq, item)
}

// Pop 从 pq 中取出最优先的 item
func (pq *maxPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.maxIndex = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
