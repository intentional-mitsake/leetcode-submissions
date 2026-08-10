func longestConsecutive(nums []int) int {
    check := make(map[int]bool)
    for _, num := range nums{
        check[num] = true
    }
    longest := 0
    for key, _ := range check {
        if _, found := check[key - 1]; !found {
            length := 0
            for {
                if _, exists := check[key+length]; exists {
                    length++
                } else {
                    break
                }
            }
            if length > longest{
                longest = length
            }
        }
    }
    return longest
}
