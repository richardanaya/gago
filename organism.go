package gago

//Organism - an interface for an organism
type Organism interface {
	Mutate()
	Strength() int
	Similarity(other interface{}) float64
	Mate(mate interface{}) Organism
	String() string
}
