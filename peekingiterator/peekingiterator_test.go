// Peeking Iterator
//
// https://leetcode.com/problems/peeking-iterator/description/
package peekingiterator

import (
	"reflect"
	"testing"
)

func TestNewIterator(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want *Iterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIterator(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIterator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIterator_next(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iterator{
				data: tt.fields.data,
			}
			if got := i.next(); got != tt.want {
				t.Errorf("Iterator.next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIterator_hasNext(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iterator{
				data: tt.fields.data,
			}
			if got := i.hasNext(); got != tt.want {
				t.Errorf("Iterator.hasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPeekingIterator(t *testing.T) {
	type args struct {
		it *Iterator
	}
	tests := []struct {
		name string
		args args
		want *PeekingIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPeekingIterator(tt.args.it); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPeekingIterator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeekingIterator_HasNext(t *testing.T) {
	type fields struct {
		it   *Iterator
		next *int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &PeekingIterator{
				it:   tt.fields.it,
				next: tt.fields.next,
			}
			if got := i.HasNext(); got != tt.want {
				t.Errorf("PeekingIterator.HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeekingIterator_Next(t *testing.T) {
	type fields struct {
		it   *Iterator
		next *int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &PeekingIterator{
				it:   tt.fields.it,
				next: tt.fields.next,
			}
			if got := i.Next(); got != tt.want {
				t.Errorf("PeekingIterator.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeekingIterator_Peek(t *testing.T) {
	type fields struct {
		it   *Iterator
		next *int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &PeekingIterator{
				it:   tt.fields.it,
				next: tt.fields.next,
			}
			if got := i.Peek(); got != tt.want {
				t.Errorf("PeekingIterator.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlow(t *testing.T) {
	p := NewPeekingIterator(
		NewIterator(
			[]int{1, 2, 3},
		),
	)
	if got := p.Next(); got != 1 {
		t.Errorf("got %v wanted %v", got, 1)
	}

	if got := p.Peek(); got != 2 {
		t.Errorf("got %v wanted %v", got, 2)
	}

	if got := p.Next(); got != 2 {
		t.Errorf("got %v wanted %v", got, 2)
	}
}
