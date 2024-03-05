package containsdup

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: true,
		},
		{
			name: "b",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "c",
			args: args{
				nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertEql(t, "Solution() = %v, want %v", Solution(tt.args.nums), tt.want)
		})
	}
}

func AssertEql[T comparable](t *testing.T, format string, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf(format, got, want)
	}
}
