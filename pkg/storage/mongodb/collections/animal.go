package collections

//AnimalPopularNames : list of names for locals from given country.
type AnimalPopularNames struct {
	PopularNames []string `json:"PopularNames" bson:"PopularNames"`
	Country      string   `json:"Country" bson:"Country"`
}

//Animal : estructura de animal dentro de la base de datos
//Type references from what kind of animal is --> Vertebrate or Invertebrate
type Animal struct {
	ScientificName     string               `json:"ScientificName" bson:"ScientificName"`
	Description        string               `json:"Description" bson:"Description"`
	Amount             int32                `json:"Amount" bson:"Amount"`
	AnimalPopularNames []AnimalPopularNames `json:"AnimalPopularNames" bson:"AnimalPopularNames"`
	AnimalType         string               `json:"AnimalType" bson:"AnimalType"`
}
