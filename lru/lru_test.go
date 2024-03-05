// LRU Cache
// https://leetcode.com/problems/lru-cache/description/
// #hashmap
package lru

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want LRUCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Get(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		c    *LRUCache
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LRUCache{}
			if got := c.Get(tt.args.key); got != tt.want {
				t.Errorf("LRUCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Put(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name string
		c    *LRUCache
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LRUCache{}
			c.Put(tt.args.key, tt.args.value)
		})
	}
}

func TestFlow(t *testing.T) {
	c := New(2)
	//t.Errorf("testing %#v", c)
	c.Put(1, 1)                    // cache is {1=1}
	c.Put(2, 2)                    // cache is {1=1, 2=2}
	if got := c.Get(1); got != 1 { // return 1
		t.Errorf("wanted %d got %d", 1, got)
	}
	c.Put(3, 3)                     // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	if got := c.Get(2); got != -1 { // returns -1 (not found)
		t.Errorf("wanted %d got %d", -1, got)
	}
	c.Put(4, 4)                     // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	if got := c.Get(1); got != -1 { // return -1 (not found)
		t.Errorf("wanted %d got %d", -1, got)
	}
	if got := c.Get(3); got != 3 { // return 3
		t.Errorf("wanted %d got %d", 3, got)
	}
	if got := c.Get(4); got != 4 { // return 4
		t.Errorf("wanted %d got %d", 4, got)
	}
}
