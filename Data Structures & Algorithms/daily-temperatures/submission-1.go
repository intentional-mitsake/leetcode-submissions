func dailyTemperatures(temperatures []int) []int {
	// store diff
	res := make([]int, len(temperatures))
	var stck []int 
	for indx, t := range temperatures {
		for len(stck) > 0 && t > temperatures[stck[len(stck)-1]] {
			tos := stck[len(stck)-1]
			stck = stck[:len(stck)-1]
			res[tos] = indx - tos
		}
		stck = append(stck, indx)
	}
	return res
}
