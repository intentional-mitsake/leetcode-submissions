class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        rems = set()
        for i in range(len(nums)):
            rem = target - nums[i]
            if rem in rems:
                return [nums.index(rem), i]
            rems.add(nums[i])
        return []
