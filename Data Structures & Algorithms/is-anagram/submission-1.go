func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    freqM1 := make(map[byte]int) // [char]freq
    freqM2 := make(map[byte]int)
    for i:=0; i<len(s); i++ {
        freqM1[s[i]]++
        freqM2[t[i]]++
    }
    // no comparison btwn two maps in go unlike python
    for k, v := range freqM1 {
        if freqM2[k] != v {
            return false
        }
    }
    return true
}
