package chipmunk

import "testing"

func TestLcg(t *testing.T) {
	seed := 3
	l := newlcg(seed)
	m := map[int]bool{}
	for i := 0; i < M; i++ {
		m[l.Next()] = true
	}
	if len(m) != M {
		t.Errorf("sizeof m should be equal to %d, get %d", M, len(m))
	}
}
