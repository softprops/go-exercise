// Can Place Flowers
//
// You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.
//
// Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.
//
// https://leetcode.com/problems/can-place-flowers/descriptio
package canplaceflowers

func Solution(flowerbed []int, n int) bool {
	for i, l := 0, false; i < len(flowerbed) && n > 0; i++ {
		switch flowerbed[i] {
		case 0:
			if !l {
				if i == len(flowerbed)-1 || (i < len(flowerbed) && flowerbed[i+1] == 0) {
					n--
					l = true
					continue
				}
			}
			l = false
		case 1:
			l = true
		}
	}
	return n < 1
}
