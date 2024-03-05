package strgcd

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a",
			args: args{
				str1: "ABCABC",
				str2: "ABC",
			},
			want: "ABC",
		},
		{
			name: "b",
			args: args{
				str1: "ABABAB",
				str2: "ABAB",
			},
			want: "AB",
		},
		{
			name: "c",
			args: args{
				str1: "LEET",
				str2: "CODE",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
