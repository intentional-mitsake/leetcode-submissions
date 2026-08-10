func twoSum(nums []int, target int) []int {
    rems := make(map[int]int);
    for i, num := range nums {
        rem := target - num;
        if index, exists := rems[rem]; exists {
            return []int{index, i};
        }
        rems[num] = i;
    }
    return nil;
}
