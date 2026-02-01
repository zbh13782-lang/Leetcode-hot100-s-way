package hot

func isValid(s string) bool {
	stack := []byte{}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, byte(v))
		} else {
			if len(stack) == 0 {
				return false
			}
			if v == ')' {
				if stack[len(stack)-1] != '(' {
					return false
				}

			}

			if v == ']' {
				if stack[len(stack)-1] != '[' {
					return false
				}
			}

			if v == '}' {
				if stack[len(stack)-1] != '{' {
					return false
				}
			}

			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
