// Pascal's Triangle
//
// Given an integer numRows, return the first numRows of Pascal's triangle.
//
// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
//
// #array #dp
// https://leetcode.com/problems/pascals-triangle/description/
package pascalstriangle

func Solution(numRows int) [][]int {
	rows := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[len(row)-1] = 1, 1
		for j := 1; j < i; j++ {
			prev := rows[i-1]
			row[j] = prev[j-1] + prev[j]
		}
		rows[i] = row
	}
	return rows
}
