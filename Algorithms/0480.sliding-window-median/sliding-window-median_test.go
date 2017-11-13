package Problem0480

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	k    int
	ans  []float64
}{

	{
		[]int{1, 3, -1, -3, 5, 3, 6, 7},
		3,
		[]float64{1, -1, -1, 3, 5, 6},
	},

	// 可以有多个 testcase
}

func Test_medianSlidingWindow(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, medianSlidingWindow(tc.nums, tc.k), "输入:%v", tc)
	}
}

func Benchmark_medianSlidingWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			medianSlidingWindow(tc.nums, tc.k)
		}
	}
}
