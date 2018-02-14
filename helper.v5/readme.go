package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func rebuildReadme() {
	// 由于网络原因，有时候 signin 比较慢
	signin()

	lc := lastestLeetCode()
	makeREADME(lc)
}

func makeREADME(lc *leetcode) {
	file := "README.md"
	os.Remove(file)

	// 更新 README.md 的内容
	template := `%s

## 进度

> **注意**：从 743 题开始，API 中的题号与页面中的出现错位。
> 本表格以 API 中的题号为准。
> 统计规则：1.免费题，2.算法题，3.能用 Go 解答

%s
## 题解

%s
以下免费的算法题，暂时不能使用 Go 解答

%s
#%s
`

	head := getHead(lc)

	progressTable := lc.Algorithms.progressTable()

	availableTable := lc.Problems.available().table()

	unavailList := lc.Problems.unavailable().list()

	tail := read("README_TAIL.md")

	content := fmt.Sprintf(template, head, progressTable, availableTable, unavailList, tail)

	// 保存 README.md 文件
	err := ioutil.WriteFile(file, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func getHead(lc *leetcode) string {
	headFormat := string(read("README_HEAD.md"))
	return fmt.Sprintf(headFormat,
		lc.Username, lc.Username,
		lc.Ranking, lc.Username,
		lc.Progress, lc.Username,
	)
}
