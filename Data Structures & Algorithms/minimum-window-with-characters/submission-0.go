func minWindow(s string, t string) string {
	if t == "" {
		return ""
	}
	freqMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		freqMap[t[i]]++
	}
	wMap := make(map[byte]int)
	l := 0
	contains := 0 
	needed := len(freqMap)
	indx := []int{-1,-1}
	shortest := math.MaxInt
	for r := 0; r < len(s); r++ {
		wMap[s[r]]++
		// if the char was in t and the freq of it in the wndow is same as in t, contains incr
		if freqMap[s[r]] > 0 && freqMap[s[r]] == wMap[s[r]] {
			contains++
		}
		// contains only incr when enough freq of a char is in the wMap
		// if a window has all chars needed(dupli include), then contains = needed
		for contains == needed { // for to pop till cant pop(popping dels char in freqMap)
			// len(s) = 10, 9-0+1 = 10
			if r-l+1 < shortest {
				shortest = r-l+1
				indx = []int{l, r}
			}
			wMap[s[l]]--// pop from left 
			// if the char needs to be in sub AND has been popped
			if freqMap[s[l]] > 0 && freqMap[s[l]] > wMap[s[l]] {
				// incr of l after this isnt registered
				// so the prev l to r window is used
				contains--
			}
			l++
		}
	}
	if indx[0] == -1 {
		return ""
	}
	return s[indx[0]:indx[1]+1]
}
