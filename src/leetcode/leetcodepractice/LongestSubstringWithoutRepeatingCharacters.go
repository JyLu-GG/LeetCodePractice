package leetcodepractice

func LengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	mp := make(map[byte]int)
	ret := 0

	for i, j := 0, 0; i < len(s); i++ {
		if _, ok := mp[s[i]]; ok {
			j = max(j, mp[s[i]]+1)
		}
		mp[s[i]] = i

		ret = max(ret, i-j+1)
		//fmt.Printf("Runing s : %s, i : %d, j : %d, max : %d\n", string(s[i]), i, j, ret)
	}
	//fmt.Println(mp)
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
