package mergealt

import "bytes"

func Solution(word1, word2 string) string {
	buf := bytes.NewBufferString("")
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			buf.WriteRune(rune(word1[i]))
		}
		if i < len(word2) {
			buf.WriteRune(rune(word2[i]))
		}
	}
	return buf.String()
}
