import "math/rand"

type RandomizedSet struct {
	mapping map[int]int
	values  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		mapping: make(map[int]int),
		values:  make([]int, 0),
	}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.mapping[val]; ok {
		return false
	}

	s.values = append(s.values, val)
	s.mapping[val] = len(s.values) - 1
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	if _, ok := s.mapping[val]; !ok {
		return false
	}

	if len(s.values) == 1 {
		s.values = []int{}
		s.mapping = map[int]int{}
		return true
	}

	index := s.mapping[val]
	delete(s.mapping, val)

	if index != len(s.values)-1 {
		s.values[index] = s.values[len(s.values)-1]
		s.mapping[s.values[len(s.values)-1]] = index
	}

	s.values = s.values[:len(s.values)-1]
	return true
}

func (s *RandomizedSet) GetRandom() int {
	return s.values[rand.Intn(len(s.values))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
