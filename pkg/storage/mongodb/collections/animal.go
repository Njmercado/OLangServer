package collections

import "log"

//AnimalPopularNames : list of names for locals from given country.
type AnimalPopularNames struct {
	PopularNames []string
	Country      string
}

//Animal : estructura de animal dentro de la base de datos
//Type references from what kind of animal is --> Vertebrate or Invertebrate
type Animal struct {
	ScientificName string
	Description    string
	Amount         int32
	AnimalPopularNames
	animalType string
}

//SetAnimalType : this function allow to define animal type
func (a *Animal) setAnimalType(animalType string) {

	if animalType != "Vertebrate" && animalType != "Invertebrate" {
		log.Panic("Type of inserted animal is invalid")
		log.Panic("There are only two possible values: 'Vertebrate' and 'Invertebrate'")
		log.Panic("Please insert correct one")
	} else {
		(*a).animalType = animalType
	}
}
