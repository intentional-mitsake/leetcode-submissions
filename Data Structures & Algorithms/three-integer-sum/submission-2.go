func threeSum(nums []int) [][]int {
	// sort the array
	sort.Ints(nums)
	// result
	var res [][]int
	var i int
	for i = 0; i<len(nums); i++ {
		a := nums[i]
		l, r := i+1, len(nums)-1
		if a > 0 {
			// means rest of the list is also >0 so no point
			break
		} 
		
		if i>0 && a == nums[i-1] {
			// duplicate
			continue
		}
		for l < r {
			
				// all need to be distinct
				target := a+nums[l]+nums[r]
				if target == 0 {
					res = append(res, []int{a, nums[l], nums[r]})
					// move on to the next nums
					// cant have duplicates, and if move only one and works its prob duplicate
					r--
					l++
					// to avoid using the same num multi times
					for l < r && nums[l] == nums[l-1] {
						l++
                    }
				} else if target > 0 {
					// too high, reduce size, move r
					r--
				} else {
					// too small
					l++
				}
		}
	}
	return res
}
