package leetcodepractice

func IsPalindrome(x int) bool {
	tmp := x
	rev := 0
	for tmp > 0 {
		rev = (rev * 10) + (tmp % 10)
		tmp /= 10
	}
	if x == rev {
		return true
	}
	return false
}
