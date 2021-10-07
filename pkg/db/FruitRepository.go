package db

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"strings"
)

const (
	ALL_ROWS = "ALL"
)

func (req *FruitRequest) GetDbSearchQuery() string {
	var query []string

	if req.Id != "" {
		query = append(query, fmt.Sprintf("id = '%s'", req.Id))
	}
	if req.Name != "" {
		query = append(query, fmt.Sprintf("name = '%s'", req.Name))
	}
	if req.Description != "" {
		query = append(query, fmt.Sprintf("description = '%s'", req.Description))
	}

	return strings.Join(query, " OR ")
}

func (receiver *FruitService) InsertFruit(fruit *FruitRequest) FruitResponse {
	log.Printf("Inserting Fruit (%+v)\n", fruit)

	db := GetPostgreSqlConnection()

	dbContext := GetContext()

	fruit.Id = uuid.NewString() // generate the ID ourselves

	_, err := db.ExecContext(
		dbContext,
		"INSERT INTO Fruit(id, name, description) VALUES($1, $2, $3)",
		fruit.Id, fruit.Name, fruit.Description)
	CheckErrorWithPanic(err, fmt.Sprintf("error while inserting Fruit. Data: %+v", fruit))

	fruitId := fruit.Id
	log.Printf("Inserted Fruit (id=%s, %+v)\n", fruitId, fruit)

	return FruitResponse{
		Id:      fruitId,
		Message: RESPOSNE_SUCCESS,
		Err:     nil,
	}
}

func (receiver *FruitService) DeleteFruits(req *FruitRequest) FruitResponse {
	log.Printf("Deleting Fruit(s) (%+v)\n", req)

	query := req.GetDbSearchQuery()

	db := GetPostgreSqlConnection()

	dbContext := GetContext()
	_, err := db.ExecContext(
		dbContext,
		"DELETE Fruit WHERE "+query)
	CheckErrorWithPanic(err, fmt.Sprintf("error while deleting Fruit (%+v)", req))

	log.Printf("Deleted Fruit(s) matching (%+v)\n", req)

	return FruitResponse{
		Message: RESPOSNE_SUCCESS + " in deleting fruits matching criteria",
		Err:     nil,
	}
}

func (receiver *FruitService) GetFruit(req *FruitRequest) Fruit {
	fruitResponse := receiver.GetFruits(req)

	var fruit Fruit
	if fruitResponse.Fruits != nil && len(fruitResponse.Fruits) > 0 {
		fruit = fruitResponse.Fruits[0]
	}

	return fruit
}

func (receiver *FruitService) GetFruits(req *FruitRequest) FruitResponse {
	log.Printf("Retrieving Fruits (+%v)\n", req)

	selectQuery := receiver.getSelectQuery(req)
	log.Printf("Query: [%s]\n", selectQuery)
	db := GetPostgreSqlConnection()

	rows, err := db.Query(selectQuery)
	CheckErrorWithPanic(err, fmt.Sprintf("error while retrieving Fruits (+%v)", req))
	defer rows.Close()

	var fruits []Fruit
	for rows.Next() {
		fruit := Fruit{}
		err := rows.Scan(&fruit.Id, &fruit.Name, &fruit.Description)
		CheckErrorWithPanic(err, "error while trying to decode Fruit")
		fruits = append(fruits, fruit)
	}

	var message = fmt.Sprintf("Found %d fruits", len(fruits))

	if len(fruits) == 0 {
		// Create default "no-result" response
		message = "No fruits found for given query"
		fruits = nil
	}

	response := FruitResponse{
		Message: message,
		Fruits:  fruits,
	}

	return response
}

func (receiver *FruitService) getSelectQuery(req *FruitRequest) string {
	var queryStr = "SELECT id, name, description FROM fruit"
	if req.Name != ALL_ROWS {
		query := req.GetDbSearchQuery()
		queryStr += " WHERE " + query
	}

	return queryStr
}
