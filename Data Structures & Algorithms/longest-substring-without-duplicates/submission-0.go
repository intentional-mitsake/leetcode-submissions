func lengthOfLongestSubstring(s string) int {
    subMap := make(map[byte]bool)
    longest := 0
    r, l := 0, 0
    for r = 0; r < len(s); r++ {
        if subMap[s[r]] {
            // if come across a duplicate, delete that 
            for subMap[s[r]] {
                // keep moving the left pointer and popping char
                // untill no duplicate left in the substring
                delete(subMap, s[l])
                l++
            }
            subMap[s[r]] = true // add the new on(not duplicat)
            continue
        }
        subMap[s[r]] = true
        // only update length if its longest
        if len(subMap) > longest {
            longest = len(subMap)
        }
    }
    return longest
}
