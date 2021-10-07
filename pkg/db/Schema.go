package db

import (
	"context"
	"fmt"
)

type Fruit struct {
	Id          string `json:"id" bson:"_id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type FruitRequest Fruit

func (fruit *Fruit) String() string {
	return fmt.Sprintf("id=%s, name=%s, description=%s", fruit.Id, fruit.Name, fruit.Description)
}

type FruitResponse struct {
	Id      interface{} `json:"id,omitempty" bson:"_id"`
	Message string      `json:"message,omitempty"`
	Fruits  []Fruit     `json:"fruits,omitempty"` // Do not display if empty
	Err     *error      `json:"error,omitempty"`  // Do not display if empty
}

type Service interface {
	InsertFruit(req *FruitRequest) FruitResponse
	GetFruits(req *FruitRequest) FruitResponse
	GetFruit(req *FruitRequest) Fruit
	DeleteFruits(req *FruitRequest) FruitResponse
}

type FruitService struct {
	fruit *Fruit
}

func GetContext() context.Context {
	return context.Background()
}
