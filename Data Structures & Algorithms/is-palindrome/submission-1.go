func isPalindrome(s string) bool {
    lwr := strings.ToLower(s)
    noSpace := strings.ReplaceAll(lwr, " ", "")
    slicedStr := []rune(noSpace)
    // trying a two pointer soln
    l, r := 0, len(slicedStr)-1
    for l < r {
        if !unicode.IsLetter(slicedStr[l]) && !unicode.IsDigit(slicedStr[l]) {
            // get to the first alphanumeric char
            l++
            continue
            
        } 
        if !unicode.IsLetter(slicedStr[r]) && !unicode.IsDigit(slicedStr[r]) {
            r--
            continue        
        } 
        // getting here means alphanumeric reached
        if slicedStr[l] != slicedStr[r] {
            return false
        } 
        // simul moving both
        l++
        r--
    }
    return true
}
