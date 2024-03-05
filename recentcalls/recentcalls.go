// Number of Recent Calls
//
// https://leetcode.com/problems/number-of-recent-calls/description
//
// #queue
package recentcalls

type RecentCounter struct {
	requests []int
}

func Constructor() *RecentCounter {
	return &RecentCounter{}
}

func (r *RecentCounter) Ping(t int) int {
	r.requests = append(r.requests, t)
	for r.requests[0] < t-3000 {
		r.requests = r.requests[1:]
	}
	return len(r.requests)
}
