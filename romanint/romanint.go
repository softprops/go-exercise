// Roman to Integer
//
// #hashtable
//
// https://leetcode.com/problems/roman-to-integer/description/
package romanint

var Lookup = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func Solution(s string) int {
	sum := 0
	for i, r := range s {
		v := Lookup[r]
		if i+1 < len(s) && v < Lookup[rune(s[i+1])] {
			sum -= v
		} else {
			sum += v
		}
	}
	return sum
}
