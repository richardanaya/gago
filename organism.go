package gago

//Organism - an interface for an organism
type Organism interface {
	Mutate()
	Strength() int
	Mate(mate interface{}) Organism
}
