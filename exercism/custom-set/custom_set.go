package stringset

import "fmt"

type Set map[string]bool

func (s Set) Add(a string) {
	s[a] = true
}

func (s Set) Delete(r string) {
	delete(s, r)
}

func (s Set) Has(r string) bool {
	return s[r]
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Len() int {
	return len(s)
}

func (s Set) Slice() []string {
	st := make([]string, len(s))
	i := 0
	for v := range s {
		st[i] = v
		i++
	}
	return st
}

func (s Set) String() string {
	out := "{"
	for k := range s {
		out += fmt.Sprintf("%+q, ", k)
	}
	if !s.IsEmpty() {
		out = out[:len(out)-2]
	}
	return out + "}"
}

func Difference(s1, s2 Set) Set {
	s := Set{}
	for v := range s1 {
		if !s2[v] {
			s[v] = true
		}
	}
	return s
}

func Disjoint(s1, s2 Set) bool {
	for v := range s1 {
		if s2.Has(v) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if s1.Len() != s2.Len() {
		return false
	}
	for v := range s1 {
		if !s2.Has(v) {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	s := Set{}
	for v := range s1 {
		if s2[v] {
			s[v] = true
		}
	}
	return s
}

func New() Set {
	return Set{}
}

func NewFromSlice(sl []string) Set {
	s := make(Set, len(sl))
	for _, v := range sl {
		s[v] = true
	}
	return s
}

func Subset(s1, s2 Set) bool {
	for v := range s1 {
		if !s2[v] {
			return false
		}
	}
	return true
}

func Union(s1, s2 Set) Set {
	s := Set{}
	for k := range s1 {
		s[k] = true
	}
	for k := range s2 {
		s[k] = true
	}
	return s
}
