package Problem0137

// Golang 中的 int 为 64 bit
func singleNumber(nums []int) int {
	n := len(nums)
	res := 0
	base := 1

	for i := 0; i < 64; i++ {
		temp := 0
		for j := 0; j < n; j++ {
			temp += nums[j] % 2 //首先把输入数字的第i位加起来。
			nums[j] /= 2
		}
		temp %= 3          //然后求它们除以3的余数。
		res += temp * base //把二进制表示的结果转化为十进制表示的结果
		base *= 2
	}

	return res
}
