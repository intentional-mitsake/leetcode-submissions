func twoSum(numbers []int, target int) []int {
    // already sorted set, no two for loops(O(n^2)), two pointers(O(n))
    l, r := 0, len(numbers)-1 // 0, len-1
    for l < r {
        // only need cases where ind1<ind2
        if numbers[l]+numbers[r] == target {
            return []int{l+1, r+1}
        } else if numbers[l]+numbers[r] > target {
            // too large nums, reduce size, move r
            r--
        } else {
            // too small, incr size, move l
            l++
        }
    }
    return []int{}
}
