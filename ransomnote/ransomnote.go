package ransomnote

func Solution(ransomNote, magazine string) bool {
	letters := make(map[rune]int)
	for _, r := range magazine {
		letters[r] += 1
	}
	for _, r := range ransomNote {
		if count, ok := letters[r]; ok && count > 0 {
			letters[r] -= 1
		} else {
			return false
		}
	}
	return true
}
