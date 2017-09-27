package Problem0179

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  string
}{

	{
		[]int{1, 20},
		"201",
	},

	{
		[]int{3, 30, 34, 5, 9},
		"9534330",
	},

	// 可以有多个 testcase
}

func Test_largestNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, largestNumber(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_largestNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			largestNumber(tc.nums)
		}
	}
}
