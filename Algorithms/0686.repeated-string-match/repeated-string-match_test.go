package Problem0686

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	a, b string
	ans  int
}{

	{
		"bb",
		"bbbbbbb",
		4,
	},

	{
		"aa",
		"a",
		1,
	},

	{
		"a",
		"aa",
		2,
	},

	{
		"abcd",
		"bcabcdbc",
		-1,
	},

	{
		"abcd",
		"cdabdab",
		-1,
	},

	{
		"abcd",
		"cdabcdab",
		3,
	},

	// 可以有多个 testcase
}

func Test_repeatedStringMatch(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, repeatedStringMatch(tc.a, tc.b), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			repeatedStringMatch(tc.a, tc.b)
		}
	}
}
