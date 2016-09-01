package gago

import "strconv"

// Population - A population of serveral species
type Population struct {
	species             []*Species
	organismFactory     func() Organism
	similarityThreshold float64
}

//NewPopulation - create a population from an organism factory function of a given size and speciate it
func NewPopulation(create func() Organism, size int) *Population {
	p := new(Population)
	p.organismFactory = create
	p.similarityThreshold = .5
	for i := 0; i < size; i++ {
		p.CreateOrganism()
	}
	return p
}

//CreateOrganism - Create an organism without parents
func (p *Population) CreateOrganism() {
	p.AddOrganism(p.organismFactory())
}

//AddOrganism -- Add an organism to the right species
func (p *Population) AddOrganism(o Organism) {
	if len(p.species) == 0 {
		firstSpecies := new(Species)
		p.species = append(p.species, firstSpecies)
		firstSpecies.members = append(firstSpecies.members, o)
	} else {
		var similarSpecies *Species
		for j, speciesLen := 0, len(p.species); j < speciesLen; j++ {
			strongest := p.species[j].Strongest()
			if strongest.Similarity(o) > p.similarityThreshold {
				similarSpecies = p.species[j]
				break
			}
		}
		if similarSpecies != nil {
			similarSpecies.members = append(similarSpecies.members, o)
		} else {
			newSpecies := new(Species)
			p.species = append(p.species, newSpecies)
			newSpecies.members = append(newSpecies.members, o)
		}
	}
}

func (p *Population) String() string {
	s := "Population:\n"
	for i, l := 0, len(p.species); i < l; i++ {
		s += "    Species " + strconv.Itoa(i) + ":\n"
		for j, k := 0, len(p.species[i].members); j < k; j++ {
			s += "          :" + p.species[i].members[j].String() + "\n"
		}
	}
	return s
}
