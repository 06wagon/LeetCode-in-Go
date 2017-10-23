package Problem0363

import "sort"

func maxSumSubmatrix(mat [][]int, target int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}

	m, n := len(mat), len(mat[0])
	// 本算法的复杂度是 O(M*M*N*logN)
	// 让 N = max(m,n) 可以加快程序
	M, N := min(m, n), max(m, n)
	// 从此开始称， M 为行，N 为列

	ans := -1 << 63

	var findTarget func([]int, int, int) int
	// 运用归并排序的思想，在 sums 中查找 <=target 的最大元素
	findTarget = func(sums []int, beg, end int) int {
		if beg+1 >= end {
			return -1 << 63
		}

		mid := (beg + end) >> 1

		res := max(findTarget(sums, beg, mid), findTarget(sums, mid, end))
		if res == target {
			return res
		}

		i := mid
		for _, l := range sums[beg:mid] {
			for i < len(sums) && sums[i]-l <= target {
				res = max(res, sums[i]-l)
				i++
			}
		}

		sort.Ints(sums[beg:end])
		return res
	}

	var iFirst, iLast, j, low, maxArea int
	var temp, sums []int

	for iFirst = 0; iFirst < M; iFirst++ {
		// temp[j] 表示 mat[iFirst:iLast+1][j] 中所有元素之和
		temp = make([]int, N)
		for iLast = iFirst; iLast < M; iLast++ {
			// sums[j] 是 mat[iFirst:iLast+1][:j+1] 中所有元素之和
			// sums[j2] - sums[j1], j1 < j2 表示
			// mat[iFisrt:iLast+1][j1:j2+1] 中所有元素之和
			// sums 中添加 0 为了 sums[j] - 0  , 表示
			// mat[iFisrt:iLast+1][:j+1] 中所有元素之和
			sums = []int{0}
			// maxArea 是 mat[iFirst:iLast+1][:] 中所有子矩阵中，所有元素之和的 最大值
			maxArea = -1 << 63
			low = 1<<63 - 1
			for j = 0; j < N; j++ {
				// 分情况更新 temp[j]
				if m < n {
					temp[j] += mat[iLast][j]
				} else {
					temp[j] += mat[j][iLast]
				}

				sums = append(sums, sums[len(sums)-1]+temp[j])

				maxArea = max(maxArea, sums[len(sums)-1]-low)

				low = min(low, sums[len(sums)-1])
			}

			if maxArea <= ans {
				// mat[iFirst:iLast+1][:] 中不可能有比 ans 更接近 target 的数
				// 可以开始下一轮了
				continue
			}

			if maxArea == target {
				// mat[iFirst:iLast+1][:] 中有一个子矩阵的所有元素之和，正好就是 target
				// 找到答案，可以结束程序了
				return target
			}

			if maxArea > target {
				// mat[iFirst:iLast+1][:] 中可能有个子矩阵的所有元素之和，
				// 比 ans 更接近 target 需要运行 findTarget 查找
				tempAns := findTarget(sums, 0, N+1)
				if tempAns == target {
					// 找到了 target
					// 直接结束程序
					return target
				}
				ans = max(ans, tempAns)
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
