package db

type SampleFruit struct {
	Id          string `json:"id" bson:"_id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type SampleFruitRequest Fruit
