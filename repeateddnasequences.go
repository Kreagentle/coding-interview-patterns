package main

func findRepeatedSequences(dna string, k int) Set {
  var left, right int
  right += k
  test := *NewSet()
	output := *NewSet()
	for i := 0; i <= len(dna) - k; i++ {
	  if test.Exists(dna[left:right]) {
	    output.Add(dna[left:right])
	  }
	  test.Add(dna[left:right])
	  left += 1
	  right += 1
	}
	return output
}

type Set struct {
	hashMap map[interface{}]bool
}

func NewSet() *Set {
	s := new(Set)
	s.hashMap = make(map[interface{}]bool)
	return s
}

func (s *Set) Add(value interface{}) {
	s.hashMap[value] = true
}

func (s *Set) Exists(value interface{}) bool {
	_, ok := s.hashMap[value]
	return ok
}
