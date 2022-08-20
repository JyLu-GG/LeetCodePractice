package leetcodepractice

func AddBinary(a string, b string) string {
	indexa := len(a) - 1
	indexb := len(b) - 1
	ret := ""
	var tmpa, tmpb, tmp byte = 48, 48, 48
	for {
		tmpa = 48
		tmpb = 48
		if indexa >= 0 {
			tmpa = a[indexa]
		}
		if indexb >= 0 {
			tmpb = b[indexb]
		}

		if tmpa == 49 && tmpb == 49 && tmp == 49 {
			tmp = 49
			ret = "1" + ret
		} else if tmpa == 49 && tmpb == 48 && tmp == 49 {
			tmp = 49
			ret = "0" + ret
		} else if tmpa == 48 && tmpb == 49 && tmp == 49 {
			tmp = 49
			ret = "0" + ret
		} else if tmpa == 49 && tmpb == 49 && tmp == 48 {
			tmp = 49
			ret = "0" + ret
		} else if tmpa == 49 && tmpb == 48 && tmp == 48 {
			tmp = 48
			ret = "1" + ret
		} else if tmpa == 48 && tmpb == 49 && tmp == 48 {
			tmp = 48
			ret = "1" + ret
		} else if tmpa == 48 && tmpb == 48 && tmp == 49 {
			tmp = 48
			ret = "1" + ret
		} else {
			tmp = 48
			ret = "0" + ret
		}
		indexa -= 1
		indexb -= 1
		if indexa < 0 && indexb < 0 {
			if tmp == 49 {
				ret = "1" + ret
			}
			break
		}
	}

	return ret
}
