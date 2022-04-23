package interation

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func TestRepeat2(t *testing.T) {
	rCnt := 100
	rChr := "a"
	var expected string
	repeated := Repeat2(rChr, rCnt)
	for i := 0; i < rCnt; i++ {
		expected += rChr
	}
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
