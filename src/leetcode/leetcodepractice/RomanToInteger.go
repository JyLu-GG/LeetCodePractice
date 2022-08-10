package leetcodepractice

func RomanToInt(s string) int {
	romandic := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	sum := 0
	for i := 0; i < len(s); {
		a := string(s[i])
		if i+1 < len(s) {
			b := string(s[i+1])
			if v, ok := romandic[a+b]; ok {
				sum += v
				i += 2
			} else if v, ok := romandic[a]; ok {
				sum += v
				i += 1
			}
		} else {
			if v, ok := romandic[a]; ok {
				sum += v
				i += 1
			}
		}
	}

	return sum
}

func RomanToInt2(s string) int {
	sum := 0
	romandic2 := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	romandic3 := map[byte]byte{
		'V': 'I',
		'X': 'I',
		'L': 'X',
		'C': 'X',
		'D': 'C',
		'M': 'C',
	}
	var per byte
	for i := len(s) - 1; i >= 0; i-- {
		v, _ := romandic3[per]
		if v == s[i] {
			sum -= romandic2[v]
		} else {
			sum += romandic2[s[i]]
		}

		per = s[i]
	}
	return sum
}
