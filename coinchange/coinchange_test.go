package coinchange

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
		{
			name: "b",
			args: args{
				coins:  []int{2},
				amount: 11,
			},
			want: -1,
		},
		{
			name: "c",
			args: args{
				coins:  []int{1},
				amount: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
