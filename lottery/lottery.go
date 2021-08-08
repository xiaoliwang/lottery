package lottery

import (
	"math/rand"
	"time"
)

type Sampling interface {
	Lot() string
}

type Possibility map[string]int

func Init() {
	rand.Seed(time.Now().UnixNano())
}

type Sample struct {
	possibility     Possibility
	all_possibility int
	middle          []string
}

func NewSample(p Possibility) *Sample {
	s := new(Sample)
	s.possibility = p
	middle := make([]string, 0, 1000)
	for key, value := range p {
		middle = sameElementArray(middle, key, value)
	}
	s.middle = middle
	s.all_possibility = len(middle)
	return s
}

func (s *Sample) Lot() string {
	i := rand.Intn(s.all_possibility)
	return s.middle[i]
}

func sameElementArray(s []string, elem string, length int) []string {
	for i := 0; i < length; i++ {
		s = append(s, elem)
	}
	return s
}

type Sample2 struct {
	possibility     Possibility
	all_possibility int
	middle          []poss
}

type poss struct {
	threshold int
	val       string
}

func NewSample2(p Possibility) *Sample2 {
	s := new(Sample2)
	s.possibility = p
	threshold := 0
	for key, value := range p {
		threshold += value
		poss := poss{
			threshold,
			key,
		}
		s.middle = append(s.middle, poss)
	}
	s.all_possibility = threshold
	return s
}

func (s *Sample2) Lot() string {
	i := rand.Intn(s.all_possibility)
	for _, val := range s.middle {
		if i < val.threshold {
			return val.val
		}
	}
	return "wrong"
}

// alias sampling
type Sample3 struct {
	possibility     Possibility
	extend          int
	all_possibility int
	middle          []alias
}

type alias struct {
	value     string
	value2    string
	threshold int
}

func NewSample3(p Possibility) *Sample3 {
	s := new(Sample3)

	extend := len(p)
	s.extend = extend
	s.possibility = p

	all_possibility := 0
	for _, value := range p {
		all_possibility += value
	}
	s.all_possibility = all_possibility

	larges := make([]alias, 0, extend)
	smalls := make([]alias, 0, extend)

	for key, value := range p {
		new_value := value * extend
		if new_value > all_possibility {
			larges = append(larges, alias{key, "", new_value})
		} else {
			smalls = append(smalls, alias{key, "", new_value})
		}
	}

	for len(larges) > 0 && len(smalls) > 0 {
		var large, small alias
		large, larges = larges[0], larges[1:]
		small, smalls = smalls[0], smalls[1:]

		large.threshold = large.threshold + small.threshold - s.all_possibility
		small.value2 = large.value

		if large.threshold > all_possibility {
			larges = append(larges, large)
		} else if large.threshold == all_possibility {
			large.value2 = large.value
			s.middle = append(s.middle, large)
		} else {
			smalls = append(smalls, large)
		}
		s.middle = append(s.middle, small)
	}
	return s
}

func (s *Sample3) Lot() string {
	i := rand.Intn(s.extend)
	poss := rand.Intn(s.all_possibility)
	mid := s.middle[i]
	if mid.threshold <= poss {
		return mid.value2
	}
	return mid.value
}
