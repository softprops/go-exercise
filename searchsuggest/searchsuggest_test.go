package searchsuggest

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		products   []string
		searchWord string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "a",
			args: args{
				products: []string{
					"mobile", "mouse", "moneypot", "monitor", "mousepad",
				},
				searchWord: "mouse",
			},
			want: [][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			},
		},
		{
			name: "b",
			args: args{
				products: []string{
					"havana",
				},
				searchWord: "havana",
			},
			want: [][]string{
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.products, tt.args.searchWord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
