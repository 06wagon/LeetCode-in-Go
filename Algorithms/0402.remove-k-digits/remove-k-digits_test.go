package Problem0402

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num string
	k   int
	ans string
}{

	{
		"102030",
		3,
		"0",
	},

	{
		"1432219",
		3,
		"1219",
	},

	{
		"10200",
		1,
		"200",
	},

	{
		"10",
		2,
		"0",
	},

	// 可以有多个 testcase
}

func Test_removeKdigits(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, removeKdigits(tc.num, tc.k), "输入:%v", tc)
	}
}

func Benchmark_removeKdigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeKdigits(tc.num, tc.k)
		}
	}
}
