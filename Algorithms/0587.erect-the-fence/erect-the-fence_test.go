package Problem0587

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	points [][]int
	ans    [][]int
}{

	{
		[][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}},
		[][]int{{1, 1}, {2, 0}, {4, 2}, {3, 3}, {2, 4}},
	},

	{
		[][]int{{1, 2}, {2, 2}, {4, 2}},
		[][]int{{1, 2}, {2, 2}, {4, 2}},
	},

	// 可以有多个 testcase
}

func Test_outerTrees(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		points := kit.Intss2Points(tc.points)
		ans := kit.Intss2Points(tc.ans)
		ast.Equal(ans, outerTrees(points), "输入:%v", tc)
	}
}

func Benchmark_outerTrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			points := kit.Intss2Points(tc.points)
			outerTrees(points)
		}
	}
}
