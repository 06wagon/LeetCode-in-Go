package Problem0130

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	board [][]byte
	ans   [][]byte
}{

	{
		[][]byte{
			[]byte("OXXOX"),
			[]byte("XOOXO"),
			[]byte("XOXOX"),
			[]byte("OXOOO"),
			[]byte("XXOXO"),
		},
		[][]byte{
			[]byte("OXXOX"),
			[]byte("XXXXO"),
			[]byte("XXXOX"),
			[]byte("OXOOO"),
			[]byte("XXOXO"),
		},
	},

	{
		[][]byte{
			[]byte("XXXX"),
			[]byte("XOOX"),
			[]byte("XXOX"),
			[]byte("XOXX"),
		},
		[][]byte{
			[]byte("XXXX"),
			[]byte("XXXX"),
			[]byte("XXXX"),
			[]byte("XOXX"),
		},
	},

	{
		[][]byte{
			[]byte("XXXX"),
			[]byte("XOOX"),
			[]byte("XOOX"),
			[]byte("XXXX"),
		},
		[][]byte{
			[]byte("XXXX"),
			[]byte("XXXX"),
			[]byte("XXXX"),
			[]byte("XXXX"),
		},
	},

	{
		[][]byte{
			[]byte("XXXX"),
		},
		[][]byte{
			[]byte("XXXX"),
		},
	},

	{
		[][]byte{
			[]byte("X"),
			[]byte("X"),
			[]byte("X"),
		},
		[][]byte{
			[]byte("X"),
			[]byte("X"),
			[]byte("X"),
		},
	},

	// 可以有多个 testcase
}

func Test_solve(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		solve(tc.board)
		ast.Equal(tc.ans, tc.board, "输入:%v", tc)
	}
}

func Benchmark_solve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			solve(tc.board)
		}
	}
}
