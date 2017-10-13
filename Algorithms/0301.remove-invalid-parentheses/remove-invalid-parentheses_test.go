package Problem0301

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans []string
}{
	{")(", []string{""}},
	{"()())()", []string{"()()()", "(())()"}},
	{"(a)())()", []string{"(a)()()", "(a())()"}},
	{"((())))))((()", []string{"((()))()"}},
	{")d))", []string{"d"}},

	// 可以有多个 testcase
}

func Test_removeInvalidParentheses(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ans := removeInvalidParentheses(tc.s)
		sort.Strings(ans)
		sort.Strings(tc.ans)
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_removeInvalidParentheses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeInvalidParentheses(tc.s)
		}
	}
}
