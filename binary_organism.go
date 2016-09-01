package gago

import (
	"math"
	"math/rand"
)

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

// Similarity - The similarity of one organism to another from 0 to 1.0
func (o *BinaryOrganism) Similarity(other interface{}) float64 {
	binaryOther := other.(*BinaryOrganism)
	var myLen = float64(len(o.dna))
	var otherLen = float64(len(binaryOther.dna))

	maxLen := math.Max(myLen, otherLen)
	minLen := math.Min(myLen, otherLen)
	similarGenes := 0.0
	for i := 0.0; i < minLen; i++ {
		if string(o.dna[int(i)]) == string(binaryOther.dna[int(i)]) {
			similarGenes++
		}
	}

	return similarGenes / maxLen
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
