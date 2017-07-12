package chipmunk

import (
	"fmt"
	"sync/atomic"
)

// https://en.wikipedia.org/wiki/Linear_congruential_generator
const (
	M = 6760000
	A = 261
	B = 1000001
)

type lcg struct {
	n int64
}

func newlcg(seed int) *lcg {
	if seed >= 6760000 {
		panic(fmt.Sprintf("seed exceed %d", M))
	}
	return &lcg{int64(seed)}
}

// thread-safe
func (l *lcg) Next() int {
	var old, new int64
	for {
		old = l.n
		new = (A*old + B) % M
		if atomic.CompareAndSwapInt64(&l.n, old, new) {
			break
		}
	}
	return int(new)
}
