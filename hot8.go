
package hot


// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

// 示例 1:

// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。
// 示例 2:

// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
func lengthOfLongestSubstring(s string) int {
	left := 0
	m := make(map[byte]int)
	res := 0
	for right := 0; right < len(s); right++ {

		for left < right && m[s[right]] == 1 {
			delete(m, s[left])
			left++
		}
		m[s[right]] = 1
		res = max(res, right-left+1)
	}
	return res
}
