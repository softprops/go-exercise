// Bulls and Cows
//
// You are playing the Bulls and Cows game with your friend.
//
// You write down a secret number and ask your friend to guess what the number is. When your friend makes a guess, you provide a hint with the following info:
//
// The number of "bulls", which are digits in the guess that are in the correct position.
// The number of "cows", which are digits in the guess that are in your secret number but are located in the wrong position. Specifically, the non-bull digits in the guess that could be rearranged such that they become bulls.
// Given the secret number secret and your friend's guess guess, return the hint for your friend's guess.
//
// The hint should be formatted as "xAyB", where x is the number of bulls and y is the number of cows. Note that both secret and guess may contain duplicate digits.
//
// #medium
//
// https://leetcode.com/problems/bulls-and-cows/description/
package bullcows

import (
	"fmt"
)

func Solution(secret, guess string) string {
	bulls, cows := 0, 0
	nums := make([]int, 10) // one slot for each digit, 0-9
	for i, s := range secret {
		if rune(guess[i]) == s {
			bulls++
			continue
		}
		is, ig := int(s-'0'), int(guess[i]-'0')
		if nums[is] < 0 {
			cows++
		}
		if nums[ig] > 0 {
			cows++
		}
		nums[is]++
		nums[ig]--
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
