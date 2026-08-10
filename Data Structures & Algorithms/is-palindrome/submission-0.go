func isPalindrome(s string) bool {
    lwr := strings.ToLower(s)
    noSpace := strings.ReplaceAll(lwr, " ", "")
    slicedStr := []rune(noSpace)
    var rev string
    var seq string
    for i:= 0; i<len(slicedStr); i++ {
        if unicode.IsLetter(slicedStr[i]) || unicode.IsDigit(slicedStr[i]) {
                seq += string(slicedStr[i])
            } 
    }
    for j:=len(slicedStr)-1; j>=0; j-- {
            if unicode.IsLetter(slicedStr[j]) || unicode.IsDigit(slicedStr[j]) {
                rev += string(slicedStr[j])
            } 
        }
    return rev == seq
}
