func productExceptSelf(nums []int) []int {
    prefix := 1
    output := make([]int, len(nums))
    for indx, num := range nums {
        output[indx] = prefix
        prefix *= num 
    }
    postfix := 1
    for j:=len(nums)-1; j>=0; j-- {
        output[j] *= postfix
        postfix *= nums[j]
    }
    return output
}
