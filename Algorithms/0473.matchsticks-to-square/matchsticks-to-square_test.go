package Problem0473

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  bool
}{

	{
		[]int{1, 1, 2, 2, 2},
		true,
	},

	{
		[]int{3, 3, 3, 3, 4},
		false,
	},

	// 可以有多个 testcase
}

func Test_makesquare(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, makesquare(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_makesquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			makesquare(tc.nums)
		}
	}
}
