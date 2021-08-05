package lottery

import (
	"testing"
)

var possibility = Possibility{
	"a": 300,
	"b": 100,
	"c": 100,
	"d": 500,
}

func TestSample(t *testing.T) {
	sample := NewSample(possibility)
	test(sample, t)
}

func TestSample2(t *testing.T) {
	sample := NewSample2(possibility)
	test(sample, t)
}

func TestSample3(t *testing.T) {
	sample := NewSample3(possibility)
	test(sample, t)
}

func test(s Sampling, t *testing.T) {
	try := 100000
	gots := Possibility{"a": 0, "b": 0, "c": 0}
	for i := 0; i < try; i++ {
		got := s.Lot()
		gots[got]++
	}
	for key, value := range possibility {
		error := abs(value*(try/1000) - gots[key])
		if error > (try / 1000 * 3) {
			t.Errorf("偏差大于 %s %d", key, gots[key])
		}
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
