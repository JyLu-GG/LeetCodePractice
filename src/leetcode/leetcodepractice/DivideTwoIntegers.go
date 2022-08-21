package leetcodepractice

import (
	"math"
)

func Divide(dividend int, divisor int) int {
	ret := 0
	d := dividend
	r := divisor
	if dividend < 0 {
		d = -d
	}
	if divisor < 0 {
		r = -r
	}
	for d >= r {
		d -= r
		ret++
	}

	if dividend < 0 || divisor < 0 {
		if !(dividend < 0 && divisor < 0) {
			ret = -ret
		}
	}

	if ret < -2147483648 {
		return -2147483648
	} else if ret > 2147483647 {
		return 2147483647
	}

	return ret
}

func Divide2(dividend int, divisor int) int {
	res := 0
	isNegative := !(dividend < 0 && divisor < 0) && (dividend < 0 || divisor < 0)
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	for dividend >= divisor {
		x := 0
		for dividend >= divisor<<x {
			x++
		}
		dividend -= divisor << (x - 1)
		res += 1 << (x - 1)
	}

	if isNegative {
		res = -res
	}
	if res >= math.MaxInt32 {
		return math.MaxInt32
	}
	if res <= math.MinInt32 {
		return math.MinInt32
	}
	return res
}
