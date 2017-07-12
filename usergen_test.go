package chipmunk

import "testing"

func TestUsernameGenerator(t *testing.T) {
	seed := 3
	ug := NewUsernameGenerator(seed)
	res := ug.convert(1000001)
	if res != "DW0001" {
		t.Errorf("expect %s, get %s", "DW0001", res)
	}

	m := map[string]bool{}
	for i := 0; i < M; i++ {
		m[ug.Generate()] = true
	}
	if len(m) != M {
		t.Errorf("sizeof m should be equal to %d, get %d", M, len(m))
	}
}
