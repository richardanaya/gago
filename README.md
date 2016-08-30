# gago
A genetics algorithm library with speciation written in Go

#Example

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
create := func() Organism {
	return NewBinaryOrganism()
}
s := NewSpecies(create, 20)
strongest := s.GetStrongest(2)
s.Breed(strongest, 20)
for i := 0; i < 40; i++ {
	s.Breed(strongest, 20)
}
fmt.Println(s.AverageStrength())
```
