package ransomnote

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a",
			args: args{
				ransomNote: "aa",
				magazine:   "b",
			},
			want: false,
		},
		{
			name: "b",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			want: false,
		},
		{
			name: "c",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
