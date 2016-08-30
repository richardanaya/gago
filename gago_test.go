package gago

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestCreateBinaryOrganism(t *testing.T) {
	rand.Seed(1)
	o := NewBinaryOrganism()
	expected := "11111010000101001110"
	result := o.String()
	if result != expected {
		t.Error("Expected " + expected + " , got " + result)
	}
}

func TestMutateBinaryOrganism(t *testing.T) {
	rand.Seed(1)
	o := NewBinaryOrganism()
	o.Mutate()
	expected := "11111010000101011110"
	result := o.String()
	if result != expected {
		t.Error("Expected " + expected + " , got " + result)
	}
}

func TestStrengthBinaryOrganism(t *testing.T) {
	rand.Seed(1)
	o := NewBinaryOrganism()
	o.Mutate()
	expected := 12
	result := o.Strength()
	if result != expected {
		t.Error("Expected: ", expected, " Result: ", result)
	}
}

func TestSpecies(t *testing.T) {
	rand.Seed(1)
	create := func() Organism {
		return NewBinaryOrganism()
	}
	s := NewSpecies(create, 20)

	expected := 20
	result := len(s.members)
	if result != expected {
		t.Error("Expected: ", expected, " Result: ", result)
	}
}

func TestSpeciesGetStrongest(t *testing.T) {
	rand.Seed(1)
	create := func() Organism {
		return NewBinaryOrganism()
	}
	s := NewSpecies(create, 20)
	strongest := s.GetStrongest(2)

	expectedLength := 2
	resultLength := len(strongest)
	if resultLength != expectedLength {
		t.Error("Expected: ", expectedLength, " Result: ", resultLength)
	}

	expectedFirstStrength := 14
	resultFirstStrength := strongest[0].Strength()
	if resultFirstStrength != expectedFirstStrength {
		t.Error("Expected: ", expectedFirstStrength, " Result: ", resultFirstStrength)
	}

	expectedSecondStrength := 13
	resultSecondStrength := strongest[1].Strength()
	if resultSecondStrength != expectedSecondStrength {
		t.Error("Expected: ", expectedSecondStrength, " Result: ", resultSecondStrength)
	}
}

func TestSpeciesBreed(t *testing.T) {
	rand.Seed(1)
	create := func() Organism {
		return NewBinaryOrganism()
	}
	s := NewSpecies(create, 20)
	fmt.Println(s.AverageStrength())
	strongest := s.GetStrongest(2)
	s.Breed(strongest, 20)
	fmt.Println(s.AverageStrength())
	for i := 0; i < 40; i++ {
		strongest = s.GetStrongest(2)
		s.Breed(strongest, 20)
		fmt.Println(s.AverageStrength())
	}
}
