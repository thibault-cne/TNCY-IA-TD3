package exercise1

import (
	"testing"
)

func TestMvMultiply(t *testing.T) {
	cases := []struct {
		v Individual
		e []int
	}{
		{
			v: Individual{V: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			e: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			v: Individual{V: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			e: []int{1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, c := range cases {
		v := c.v.MvMultiply(MATRIX)

		for i := 0; i < 12; i++ {
			if c.e[i] != v[i] {
				t.Errorf("%s expected %d actual %d\n", "MvMultiply", c.e, v)
			}
		}
	}
}

func TestFitness(t *testing.T) {
	cases := []struct {
		v Individual
		e int
	}{
		{
			v: Individual{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			e: 0,
		},
		{
			v: Individual{[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			e: 9,
		},
	}

	for _, c := range cases {
		s := c.v.fitness()
		if s != c.e {
			t.Errorf("%s expected %d actual %d\n", "fitness", c.e, s)
		}
	}
}

func TestCross(t *testing.T) {
	cases := []struct {
		i1 Individual
		i2 Individual
		r1 Individual
		r2 Individual
	}{
		{
			i1: Individual{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			i2: Individual{[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			r1: Individual{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			r2: Individual{[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
		{
			i1: Individual{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			i2: Individual{[]int{1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0}},
			r1: Individual{[]int{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}},
			r2: Individual{[]int{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}},
		},
	}

	for _, c := range cases {
		r1, r2 := c.i1.cross(c.i2)

		if !r1.compare(c.r1) {
			t.Errorf("%s expected %d actual %d\n", "cross", c.r1, r1)
		}

		if !r2.compare(c.r2) {
			t.Errorf("%s expected %d actual %d\n", "cross", c.r2, r2)
		}
	}
}

func TestSort(t *testing.T) {
	cases := []struct {
		s Population
		e Population
	}{
		{
			s: Population{[]Individual{Individual{[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, Individual{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}}},
			e: Population{[]Individual{Individual{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, Individual{[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}}},
		},
	}

	for _, c := range cases {
		c.s.sort(func(i Individual) int { return i.fitness() })

		if !c.s.compare(c.e) {
			t.Errorf("%s expected %d actual %d\n", "sort", c.e, c.s)
		}
	}
}
