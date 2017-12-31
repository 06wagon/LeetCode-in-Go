package Problem0720

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	words []string
	ans   string
}{

	{
		[]string{"yo", "ew", "fc", "zrc", "yodn", "fcm", "qm", "qmo", "fcmz", "z", "ewq", "yod", "ewqz", "y"},
		"yodn",
	},

	{
		[]string{"w", "wo", "wor", "worl", "world"},
		"world",
	},

	{
		[]string{"a", "banana", "app", "appl", "ap", "apply", "apple"},
		"apple",
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, longestWord(tc.words), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestWord(tc.words)
		}
	}
}
