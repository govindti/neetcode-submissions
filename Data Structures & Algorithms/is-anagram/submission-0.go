func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	count_s := map[rune]int{}
	count_t := map[rune]int{}

	// Count characters in string s
    // We use '_' because we don't need the index
    for _, char := range s {
        count_s[char]++ // Go safely defaults missing keys to 0
    }

    // Count characters in string t
    for _, char := range t {
        count_t[char]++
    }

	for char, count := range count_s {
		if count_t[char] != count {
			return false
		}
	}
	return true
	
}
