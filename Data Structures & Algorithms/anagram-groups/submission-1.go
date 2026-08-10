func groupAnagrams(strs []string) [][]string {
	grouped := make(map[[26]int][]string);
	for _, str := range strs {
		var key [26]int;//store 26 alphabet indx
		for _, char := range str{
			key[char - 'a']++;
		}
		grouped[key] = append(grouped[key], str);
	}
	result := make([][]string, 0, len(grouped));
	for _, val := range grouped{
		result = append(result, val);
	}
	return result;
}
