package groupanagrams

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "a",
			args: args{
				strs: []string{
					"eat", "tea", "tan", "ate", "nat", "bat",
				},
			},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
