package hot

import (
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	if s == "" {
		return s
	}

	//是字母，往后继续
	if unicode.IsLetter(rune(s[0])) {
		return s[:1] + decodeString(s[1:])
	}

	//必定是数字
	i := strings.IndexByte(s, '[')
	balance := 1
	for j := i + 1; ; j++ {
		if s[j] == '[' {
			balance++
		} else if s[j] == ']' {
			balance--
			if balance == 0 {
				number, _ := strconv.Atoi(s[:i])
				return strings.Repeat(decodeString(s[i+1:j]), number) + decodeString(s[j+1:])
			}
		}
	}
}
