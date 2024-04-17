package maxpopulationyear

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		logs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				logs: [][]int{
					{1993, 1999}, {2000, 2010},
				},
			},
			want: 1993,
		},
		{
			name: "b",
			args: args{
				logs: [][]int{
					{1950, 1961}, {1960, 1971}, {1970, 1981},
				},
			},
			want: 1960,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.logs); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
