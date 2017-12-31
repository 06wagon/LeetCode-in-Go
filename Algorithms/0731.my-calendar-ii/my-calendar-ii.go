package Problem0731

import (
	"sort"
)

// MyCalendarTwo 第二个日历类
type MyCalendarTwo struct {
	events [][]int
}

// Constructor 返回 MyCalendarTwo
func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

// Book 可以预约事件，预约成功返回 true
func (m *MyCalendarTwo) Book(start, end int) bool {
	e := []int{start, end, 1}

	n := len(m.events)
	if n == 0 {
		m.events = append(m.events, e)
		return true
	}

	i := sort.Search(n, func(i int) bool {
		return e[0] < m.events[i][1]
	})

	j := sort.Search(n, func(j int) bool {
		return e[1] <= m.events[j][0]
	})

	if i == j {
		m.events = append(m.events, nil)
		copy(m.events[i+1:], m.events[i:])
		m.events[i] = e
		return true
	}

	for k := i; k < j; k++ {
		if m.events[k][2] == 2 {
			return false
		}
	}

	segements := overlap(m.events[i:j], e)
	temp := make([][]int, len(m.events)-j)
	copy(temp, m.events[j:])
	m.events = m.events[:i]
	m.events = append(m.events, segements...)
	m.events = append(m.events, temp...)

	return true
}

func overlap(events [][]int, e []int) [][]int {
	res := make([][]int, 0, len(events)*2+1)

	for i := 0; i < len(events); i++ {
		if e[0] < events[i][0] {
			e1 := []int{e[0], events[i][0], 1}
			res = append(res, e1)
			e[0] = events[i][0]
		}
		//
		if e[1] < events[i][1] {
			e[2] = 2
			e2 := []int{e[1], events[i][1], 1}
			res = append(res, e, e2)
		} else if e[1] == events[i][1] {
			e[2] = 2
			res = append(res, e)
		} else {
			e2 := []int{e[0], events[i][1], 2}
			res = append(res, e2)
			e[0] = events[i][1]
		}
	}

	return res
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
