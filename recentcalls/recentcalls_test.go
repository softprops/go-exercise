package recentcalls

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want *RecentCounter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecentCounter_Ping(t *testing.T) {
	type args struct {
		t int
	}
	tests := []struct {
		name string
		r    *RecentCounter
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RecentCounter{}
			if got := r.Ping(tt.args.t); got != tt.want {
				t.Errorf("RecentCounter.Ping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlow(t *testing.T) {
	rc := Constructor()
	if rc.Ping(1) != 1 {
		t.Errorf("expected 1")
	}
	if rc.Ping(100) != 2 {
		t.Errorf("expected 2")
	}
	if rc.Ping(3001) != 3 {
		t.Errorf("expected 3")
	}
	if rc.Ping(3002) != 3 {
		t.Errorf("expected 3")
	}
}
