package ransomnote

// space O(n)
// time O(n*m)
func Solution(ransomNote, magazine string) bool {
	letters := make([]int, 26)
	for _, r := range magazine {
		letters[index(r)]++
	}
	for _, r := range ransomNote {
		if letters[index(r)] < 1 {
			return false
		}
		letters[index(r)]--
	}
	return true
}

func index(r rune) rune {
	return r - 'a'
}
