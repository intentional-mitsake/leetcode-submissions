func checkInclusion(s1 string, s2 string) bool {
	if len(s1)>len(s2) { return false }
	charMap1 := make([]int, 26)
	charMap2 := make([]int, 26)
	// map to the window first
	for i := 0; i<len(s1); i++ {
		charMap1[s1[i]-'a']++ // a - a = 0
		// first window map for the s2
		charMap2[s2[i]-'a']++ // l - a = 12
	}
	matches := 0
	// check for each alph
	for i:=0; i<26; i++ {
		if charMap1[i] == charMap2[i] { matches++ }
	}
	l := 0
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}

		// adding right char to window
		index := s2[r] - 'a'
		charMap2[index]++
		if charMap1[index] == charMap2[index] {
			matches++
		} else if charMap1[index]+1 == charMap2[index] {
			matches--
		}
		// removing left char from window
		index = s2[l] - 'a'
		charMap2[index]--
		if charMap1[index] == charMap2[index] {
			matches++
		} else if charMap1[index]-1 == charMap2[index] {
			matches--
		}
		l++
	}
	return matches == 26
}
