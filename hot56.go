package hot

func letterCombinations(digits string) (res []string) {
	m := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var dfs func(i int, cur []byte)
	dfs = func(i int, cur []byte) {
		if i == len(digits) {
			cp := make([]byte, len(cur))
			copy(cp, cur)
			res = append(res, string(cp))
			return
		}
		digit := int(digits[i]) - int('0')
		for j := 0; j < len(m[digit]); j++ {
			cur = append(cur, m[digit][j])
			dfs(i+1, cur)
			cur = cur[:len(cur)-1]
		}

	}
	dfs(0, []byte{})
	return
}
