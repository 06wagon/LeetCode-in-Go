package Problem0306

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num string
	ans bool
}{

	{"112358", true},
	{"199100199", true},
	{"1991001991", false},

	// 可以有多个 testcase
}

func Test_isAdditiveNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isAdditiveNumber(tc.num), "输入:%v", tc)
	}
}

func Benchmark_isAdditiveNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isAdditiveNumber(tc.num)
		}
	}
}
