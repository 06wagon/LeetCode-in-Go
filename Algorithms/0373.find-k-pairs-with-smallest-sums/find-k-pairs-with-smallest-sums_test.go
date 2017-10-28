package Problem0373

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums1 []int
	nums2 []int
	k     int
	ans   [][]int
}{

	{
		[]int{1, 7, 11},
		[]int{2, 4, 6},
		3,
		[][]int{
			[]int{1, 2},
			[]int{1, 4},
			[]int{1, 6},
		},
	},

	{
		[]int{1, 1, 2},
		[]int{1, 2, 3},
		2,
		[][]int{
			[]int{1, 1},
			[]int{1, 1},
		},
	},

	{
		[]int{1, 2},
		[]int{3},
		3,
		[][]int{
			[]int{1, 3},
			[]int{2, 3},
		},
	},

	{
		[]int{1, 3, 5, 7},
		[]int{2, 4, 6, 8},
		16,
		[][]int{
			[]int{1, 2},
			[]int{3, 2},
			[]int{1, 4},
			[]int{5, 2},
			[]int{3, 4},
			[]int{1, 6},
			[]int{7, 2},
			[]int{5, 4},
			[]int{3, 6},
			[]int{1, 8},
			[]int{7, 4},
			[]int{5, 6},
			[]int{3, 8},
			[]int{7, 6},
			[]int{5, 8},
			[]int{7, 8},
		},
	},

	// 可以有多个 testcase
}

func Test_kSmallestPairs(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, kSmallestPairs(tc.nums1, tc.nums2, tc.k), "输入:%v", tc)
	}
}

func Benchmark_kSmallestPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			kSmallestPairs(tc.nums1, tc.nums2, tc.k)
		}
	}
}
