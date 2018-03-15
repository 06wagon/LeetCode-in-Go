package problem0777

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	start string
	end   string
	ans   bool
}{

	{
		"RXXLRXRXL",
		"XRLXXRRLX",
		true,
	},

	// 可以有多个 testcase
}

func Test_canTransform(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, canTransform(tc.start, tc.end), "输入:%v", tc)
	}
}

func Benchmark_canTransform(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canTransform(tc.start, tc.end)
		}
	}
}
