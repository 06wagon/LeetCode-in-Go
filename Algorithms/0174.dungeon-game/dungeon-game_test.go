package Problem0174

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	dungeon [][]int
	ans     int
}{

	{
		[][]int{
			[]int{-2, -3, 3},
			[]int{-5, -10, 1},
			[]int{10, 30, -5},
		},
		7,
	},

	// 可以有多个 testcase
}

func Test_calculateMinimumHP(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, calculateMinimumHP(tc.dungeon), "输入:%v", tc)
	}
}

func Benchmark_calculateMinimumHP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			calculateMinimumHP(tc.dungeon)
		}
	}
}
