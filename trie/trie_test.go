package trie

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Trie
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_Insert(t *testing.T) {
	type fields struct {
		root     *Trie
		data     rune
		terminal bool
		children [26]*Trie
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Trie{
				root:     tt.fields.root,
				data:     tt.fields.data,
				terminal: tt.fields.terminal,
				children: tt.fields.children,
			}
			this.Insert(tt.args.word)
		})
	}
}

func Test_insert(t *testing.T) {
	type args struct {
		root *Trie
		word string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insert(tt.args.root, tt.args.word)
		})
	}
}

func TestTrie_Search(t *testing.T) {
	type fields struct {
		root     *Trie
		data     rune
		terminal bool
		children [26]*Trie
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Trie{
				root:     tt.fields.root,
				data:     tt.fields.data,
				terminal: tt.fields.terminal,
				children: tt.fields.children,
			}
			if got := tr.Search(tt.args.word); got != tt.want {
				t.Errorf("Trie.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search(t *testing.T) {
	type args struct {
		root *Trie
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.root, tt.args.word); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_StartsWith(t *testing.T) {
	type fields struct {
		root     *Trie
		data     rune
		terminal bool
		children [26]*Trie
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Trie{
				root:     tt.fields.root,
				data:     tt.fields.data,
				terminal: tt.fields.terminal,
				children: tt.fields.children,
			}
			if got := tr.StartsWith(tt.args.prefix); got != tt.want {
				t.Errorf("Trie.StartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_startsWith(t *testing.T) {
	type args struct {
		root *Trie
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := startsWith(tt.args.root, tt.args.word); got != tt.want {
				t.Errorf("startsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	root := New()
	root.Insert("apple")
	if !root.Search("apple") {
		t.Errorf("did not contain apple")
	}
	if root.Search("app") {
		t.Errorf("contained app")
	}
	if !root.StartsWith("app") {
		t.Errorf("did not start with app")
	}
	root.Insert("app")
	if !root.Search("app") {
		t.Errorf("did not contain app")
	}
}
