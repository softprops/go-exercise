package pascalstriangle

func Solution(numRows int) [][]int {
	rows := [][]int{}
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[len(row)-1] = 1, 1
		for j := 1; j < i; j++ {
			prev := rows[i-1]
			row[j] = prev[j-1] + prev[j]
		}
		rows = append(rows, row)
	}
	return rows
}
