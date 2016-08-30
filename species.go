package gago

import (
	"math/rand"
	"sort"
)

//Organisms - A array of organisms that can be sorted descending
type Organisms []Organism

func (o Organisms) Len() int {
	return len(o)
}
func (o Organisms) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}
func (o Organisms) Less(i, j int) bool {
	return o[i].Strength() > o[j].Strength()
}

// Species - A simple organism represented by 1s and 0s
type Species struct {
	members      Organisms
	mutationRate float32
}

// NewSpecies - Create a species with a creator function and a specific size
func NewSpecies(create func() Organism, size int) *Species {
	s := new(Species)
	for i := 0; i < size; i++ {
		s.members = append(s.members, create())
	}
	s.mutationRate = 0.03
	return s
}

// GetStrongest - Get n strongest members of a species
func (s *Species) GetStrongest(n int) []Organism {
	sort.Sort(s.members)
	return s.members[:n]
}

// BreedReplace - Breed the species of same size from a number of viable organisms
func (s *Species) BreedReplace(viable Organisms) {
	s.Breed(viable, len(s.members))
}

// AverageStrength - Get the average strength of the population
func (s *Species) AverageStrength() int {
	avg := 0
	for _, o := range s.members {
		avg += o.Strength()
	}
	return avg / len(s.members)
}

// Breed - Breed the species from a number of viable organisms
func (s *Species) Breed(viable Organisms, size int) {
	newMembers := Organisms{}
	for i := 0; i < size; i++ {
		parentA := viable[rand.Intn(len(viable))]
		parentB := viable[rand.Intn(len(viable))]
		child := parentA.Mate(parentB)
		if rand.Float32() <= s.mutationRate {
			child.Mutate()
		}
		newMembers = append(newMembers, child)
	}
	s.members = newMembers
}
