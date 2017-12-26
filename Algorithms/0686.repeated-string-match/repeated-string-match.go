package Problem0686

import (
	"strings"
)

func repeatedStringMatch(a, b string) int {
	if strings.Contains(a, b) {
		return 1
	}

	i := strings.Index(b, a)
	if i == -1 {
		return -1
	}
	r := b[:i]

	res := 0
	lenA := len(a)
	b = b[i:]
	for len(b) >= lenA && b[:lenA] == a {
		b = b[lenA:]
		res++
	}

	b += r
	if b == "" {
		return res
	} else if b == a {
		return res + 2
	}
	return -1
}
