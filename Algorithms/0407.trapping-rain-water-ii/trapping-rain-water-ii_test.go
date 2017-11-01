package Problem0407

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	heightMap [][]int
	ans       int
}{

	{
		[][]int{
			[]int{1, 4, 3, 1, 3, 2},
			[]int{3, 2, 1, 3, 2, 4},
			[]int{2, 3, 3, 2, 3, 1},
		},
		4,
	},

	{
		[][]int{
			[]int{5, 5, 5, 1},
			[]int{5, 1, 1, 5},
			[]int{5, 1, 5, 5},
			[]int{5, 2, 5, 8},
		},
		3,
	},

	// 可以有多个 testcase
}

func Test_trapRainWater(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, trapRainWater(tc.heightMap), "输入:%v", tc)
	}
}

func Benchmark_trapRainWater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			trapRainWater(tc.heightMap)
		}
	}
}
