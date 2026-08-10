type Solution struct{
    encoded string;
}

func (s *Solution) Encode(strs []string) string {
    for _, str := range strs{
        s.encoded += strconv.Itoa(len(str)) + "!" + str
    }
    return s.encoded;
}

func (s *Solution) Decode(encoded string) []string {
    var length string;
    var soln []string;
    i := 0;
    for(i<len(encoded)) {
        j := i;
        for(encoded[j] != '!'){
            //length += string(encoded[j]);
            j++;
        }
        length = encoded[i:j];
        l, _ := strconv.Atoi(length);
        soln = append(soln, encoded[j+1:j+1+l]);
        i = j+1+l;//next word
    }
    return soln
}
