package Problem0282

import (
	"fmt"
)

func addOperators(s string, target int) []string {
	res := []string{}

	ss := make([]string, 0, len(s)*2)

	var dfs func(int, []string)

	split := func(begin, end int, strs []string) {
		for i := begin + 1; i <= end; i++ {
			dfs(i, append(strs, s[begin:i], "+"))
			dfs(i, append(strs, s[begin:i], "-"))
			dfs(i, append(strs, s[begin:i], "*"))
		}
	}

	dfs = func(strs []string) {
		if len(s) == 0 && compute(strs) == target {
			res = append(res, connect(strs))
		}

	}

	dfs(s, ss)

	return res
}

func compute(ss []string) int {
	var n1, n2, n3 int
	var opt1, opt2 string
	// n1 opt1 n2 opt2 n3
	// ↓   ↓   ↓   ↓   ↓
	// 1   +   2   *   3
	n1, opt1, n2 = 0, "+", integer(ss[0])

	for i := 1; i < len(ss); i += 2 {
		fmt.Println(ss)
		opt2, n3 = ss[i], integer(ss[i+1])
		if opt2 == "*" {
			n2 = operate(n2, n3, opt2)
		} else {
			n1 = operate(n1, n2, opt1)
			opt1 = opt2
			n2 = n3
		}
	}

	return operate(n1, n2, opt1)
}

func connect(ss []string) string {
	res := ss[0]
	for i := 1; i < len(ss); i++ {
		res += ss[i]
	}
	return res
}

// 把 string 转换成 int
func integer(s string) int {
	res := int(s[0] - '0')
	for i := 1; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
	}
	return res
}

func operate(a, b int, opt string) int {
	switch opt[0] {
	case '+':
		return a + b
	case '-':
		return a - b
	default:
		return a * b
	}
}
