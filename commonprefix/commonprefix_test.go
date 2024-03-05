package commonprefix

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a",
			args: args{
				strs: []string{
					"flower", "flow", "flight",
				},
			},
			want: "fl",
		},
		{
			name: "b",
			args: args{
				strs: []string{
					"dog", "racecar", "car",
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.strs); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
