package Problem0169

func majorityElement(nums []int) int {
	n := len(nums)
	times := n / 2
	m := make(map[int]int, times)

	for _, n := range nums {
		if m[n]+1 > times {
			return n
		}
		m[n]++
	}

	max, res := 0, 0
	for i, v := range m {
		if max < v {
			max = v
			res = i
		}
	}

	return res
}
