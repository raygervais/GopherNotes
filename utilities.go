package main

func Index(vs []Note, t Note) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []Note, t Note) bool {
	return Index(vs, t) >= 0
}

func Any(vs []Note, f func(Note) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []Note, f func(Note) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []Note, f func(Note) bool) []Note {
	vsf := make([]Note, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []Note, f func(Note) String) []String {
	vsm := make([]Note, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
