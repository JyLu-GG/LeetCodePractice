package leetcodepractice

func IsValid(s string) bool {
	tmp := make(map[int]byte)
	var tmpindex = 0
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			tmp[tmpindex] = byte(v)
			tmpindex += 1
		} else if v == ')' {
			if tmp[tmpindex-1] == byte('(') {
				delete(tmp, tmpindex)
				tmpindex -= 1
			} else {
				return false
			}
		} else if v == ']' {
			if tmp[tmpindex-1] == byte('[') {
				delete(tmp, tmpindex)
				tmpindex -= 1
			} else {
				return false
			}
		} else if v == '}' {
			if tmp[tmpindex-1] == byte('{') {
				delete(tmp, tmpindex)
				tmpindex -= 1
			} else {
				return false
			}
		}
	}
	if tmpindex == 0 {
		return true
	} else {
		return false
	}
}

func IsValid2(s string) bool {
	var stack []rune
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, v := range s {
		switch v {
		case '{', '[', '(':
			stack = append(stack, v)
		case '}', ']', ')':
			ln := len(stack)
			if ln > 0 && stack[ln-1] == pairs[v] {
				stack = stack[:ln-1]
			} else {
				return false
			}
		default:
			return false
		}
	}

	return len(stack) == 0
}
