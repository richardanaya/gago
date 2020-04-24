# gago
A genetics algorithm library with speciation written in Go

# Example

In order to use gago, all you need to do is implement an structure that implements the Organism interface

```go
type Organism interface {
	Mutate()
	Strength() int
	Mate(mate interface{}) Organism
}
```

Once you do that you can generate a species using your organism and breed it to an optimal solution. An example organism BinaryOrganism has been provided for testing:

```go
//Define a function that can create Organisms
createOrganism := func() Organism { 
	return NewBinaryOrganism()
}
//Create a species of 20 members
s := NewSpecies(createOrganism, 20) 
//The average strength of the species should be fairly average at this point
fmt.Println("Start:",s.AverageStrength())
//Get the strongest 2
strongest := s.GetStrongest(2) 
//Breed the species from the 2 strongest viable members
s.Breed(strongest, 20) 
//Do this 40 times
for i := 0; i < 40; i++ { 
	strongest := s.GetStrongest(2)
	s.Breed(strongest, 20)
}
//The average strength of the species should be much higher
fmt.Println("End:",s.AverageStrength())
```

# TODO
* Implement speciated population growth of a population of multiple species using goroutines
* Better control over configuration values
