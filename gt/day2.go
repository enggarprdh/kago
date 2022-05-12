package gt

//IsVocal is function to check word is vocal or consonant
func IsVocal(a rune) bool {
	switch a {
	case 'A', 'I', 'U', 'E', 'O', 'a', 'i', 'u', 'e', 'o':
		return true
	default:
		return false
	}
}

//GetVocalInAWord function to count all vocal in a word
func GetVocalInAWord(txt string) int {
	vocal := 0
	for _, c := range txt {
		if IsVocal(c) {
			vocal++
		}
	}
	return vocal
}

//CharCounts is alias
type CharCounts map[string]int

//GetCharCount is function to get char and count
func GetCharCount(txt string) (vocals CharCounts, consonans CharCounts) {
	v := make(map[string]int)
	c := make(map[string]int)

	for index, w := range txt {
		if IsVocal(w) {
			v[string(w)] = index
		} else {
			c[string(w)] = index
		}
	}
	vocals = v
	consonans = c

	return vocals, consonans
}
