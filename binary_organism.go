package gago

import "math/rand"

// BinaryOrganism - A simple organism represented by 1s and 0s
type BinaryOrganism struct {
	dna string
}

// NewBinaryOrganism - Creates a new binary organism of a certain length
func NewBinaryOrganism() *BinaryOrganism {
	o := new(BinaryOrganism)
	o.dna = ""
	for i := 0; i < 20; i++ {
		v := rand.Intn(2)
		if v == 0 {
			o.dna += "0"
		} else {
			o.dna += "1"
		}
	}
	return o
}

// NewBinaryOrganismFromString - Creates a new binary organism from a string
func NewBinaryOrganismFromString(dna string) *BinaryOrganism {
	o := new(BinaryOrganism)
	o.dna = dna
	return o
}

// Mutate - Change the organism
func (o *BinaryOrganism) Mutate() {
	index := rand.Intn(len(o.dna))
	c := string(o.dna[index])
	if c == "1" {
		o.dna = o.dna[:index] + "0" + o.dna[index+1:]
	} else {
		o.dna = o.dna[:index] + "1" + o.dna[index+1:]
	}
}

// Strength - Get the strength based on number of 1s
func (o *BinaryOrganism) Strength() int {
	strength := 0
	for i, l := 0, len(o.dna); i < l; i++ {
		c := string(o.dna[i])
		if c == "1" {
			strength++
		}
	}
	return strength
}

// Mate - Create a new BinaryOrganism from mating with another
func (o *BinaryOrganism) Mate(mate interface{}) Organism {
	binaryMate := mate.(*BinaryOrganism)
	splitPoint := rand.Intn(len(o.dna))
	dna := o.dna[:splitPoint] + binaryMate.dna[splitPoint:]
	child := NewBinaryOrganismFromString(dna)
	return child
}

func (o *BinaryOrganism) String() string {
	return o.dna
}
