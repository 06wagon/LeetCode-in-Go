package Problem0759

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	intervals [][]int
	ans       int
}{

	{
		[][]int{{1, 3}, {8, 9}},
		4,
	},

	{
		[][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}},
		3,
	},

	{
		[][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
		5,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, intersectionSizeTwo(tc.intervals), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			intersectionSizeTwo(tc.intervals)
		}
	}
}
