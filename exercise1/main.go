package exercise1

import (
	"fmt"
	"math/rand"
)

type Mv struct {
	M [][]int
}

type Individual struct {
	V []int
}

type Population struct {
	I []Individual
}

var (
	// Shop cost
	SHOP_COST  = 6
	BOROUGH_NB = 12

	POPULATION = 100
	// Nb of generation to be generated
	GEN_NB = 100

	// Make sure the mutation probability is a float
	MUTATION_P = 0.0
)

// Shop prices
var SHOP_PI = []int{1, 8, 6, 3, 2, 4, 2, 2, 1, 1, 1, 3}

// Adjacency matrix
var MATRIX = Mv{M: [][]int{
	{1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
	{1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0},
	{1, 0, 0, 0, 1, 1, 0, 1, 0, 1, 1, 0},
	{1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
	{0, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0},
	{0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1},
	{0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0},
	{0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1},
}}

// Multiplication between a row vector and a matrix
func (v Individual) MvMultiply(m Mv) []int {
	r := make([]int, len(v.V))

	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r); j++ {
			r[i] += v.V[j] * m.M[j][i]
		}
	}

	return r
}

// Implement fitness function
func (m Individual) fitness() int {
	var r int
	b := m.MvMultiply(MATRIX)

	for i := 0; i < BOROUGH_NB; i++ {
		// If borough is covered by a shop we add it's income
		if b[i] != 0 {
			r += SHOP_PI[i]
		}

		// If there is a shop in the borough, we remove
		// the cost of the shop
		if m.V[i] != 0 {
			r -= SHOP_COST
		}
	}

	return r
}

// Max fitness in a given population
func (p Population) maxFitness() int {
	var m int

	for _, i := range p.I {
		t := i.fitness()
		if t > m {
			m = t
		}
	}

	return m
}

// Generate a random individual
func generateIndividual() Individual {
	I := Individual{V: make([]int, BOROUGH_NB)}

	for i := 0; i < BOROUGH_NB; i++ {
		I.V[i] = rand.Intn(2)
	}

	return I
}

// Mutate an individual
func (i Individual) mutation() {
	p := rand.Float64()

	if p < MUTATION_P {
		r := rand.Intn(BOROUGH_NB)
		i.V[r] = rand.Intn(2)
	}
}

// Mutate a population
func (p Population) mutation() {
	for _, i := range p.I {
		i.mutation()
	}
}

// Compare two individuals
func (i1 Individual) compare(i2 Individual) bool {
	for i, v := range i1.V {
		if v != i2.V[i] {
			return false
		}
	}

	return true
}

func (p1 Population) compare(p2 Population) bool {
	for i, ind := range p1.I {
		if !ind.compare(p2.I[i]) {
			return false
		}
	}

	return true
}

// Cross two individuals to create two new ones.
// For two individuals [A B] and [C D] we create [A D] and [C B]
func (i1 Individual) cross(i2 Individual) (Individual, Individual) {
	var n1, n2 []int

	t := BOROUGH_NB / 2

	n1 = append(n1, i1.V[:t]...)
	n1 = append(n1, i2.V[t:]...)

	n2 = append(n2, i2.V[:t]...)
	n2 = append(n2, i1.V[t:]...)

	return Individual{V: n1}, Individual{V: n2}
}

// Generate first generation
func firstGen() Population {
	var p Population

	for i := 0; i < POPULATION; i++ {
		p.I = append(p.I, generateIndividual())
	}

	return p
}

// Sort the popultation with a given function.
// We use fitness here
func (p Population) sort(f func(Individual) int) {
	var n []int

	for i := 0; i < len(p.I); i++ {
		n = append(n, f(p.I[i]))
	}

	var isDone = false

	for !isDone {
		isDone = true
		var i = 0
		for i < len(n)-1 {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
				p.I[i], p.I[i+1] = p.I[i+1], p.I[i]
				isDone = false
			}
			i++
		}
	}
}

// Select only the top half portion of the population
func (p Population) selection() {
	p.sort(func(i Individual) int { return i.fitness() })
	t := p.I

	p.I = t[POPULATION/2:]
}

// Generate the next generation
func (p Population) nextGen() {
	p.selection()

	for i := 0; i < POPULATION/4; i += 2 {
		n1, n2 := p.I[i].cross(p.I[i+1])
		p.I = append(p.I, n1, n2)
	}
}

// Run the whole process
func Run() {
	p := firstGen()

	for i := 0; i < GEN_NB; i++ {
		fmt.Printf("Generation %d max fitness : %d\n", i, p.maxFitness())
		p.nextGen()
		p.mutation()
	}

	p.sort(func(i Individual) int { return i.fitness() })
	fmt.Printf("Final population : %+v\n", p.I)
}
