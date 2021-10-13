package db

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sgahlot/go-postgres-service-quickstart/pkg/common"
	"log"
	"strings"
)

const (
	ALL_ROWS = "ALL"
)

type FruitService struct {
	fruit *common.Fruit
}

func (receiver *FruitService) GetDbSearchQuery(req *common.Fruit) string {
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

func (receiver *FruitService) InsertFruit(fruit *common.Fruit) common.FruitResponse {
	log.Printf("Inserting Fruit (%+v)\n", fruit)

	db := GetPostgreSqlConnection()

	dbContext := common.GetContext()

	fruit.Id = uuid.NewString() // generate the ID ourselves

	_, err := db.ExecContext(
		dbContext,
		"INSERT INTO Fruit(id, name, description) VALUES($1, $2, $3)",
		fruit.Id, fruit.Name, fruit.Description)
	common.CheckErrorWithPanic(err, fmt.Sprintf("error while inserting Fruit. Data: %+v", fruit))

	fruitId := fruit.Id
	log.Printf("Inserted Fruit (id=%s, %+v)\n", fruitId, fruit)

	return common.FruitResponse{
		Id:      fruitId,
		Message: common.RESPOSNE_SUCCESS,
		Err:     nil,
	}
}

func (receiver *FruitService) DeleteFruits(req *common.Fruit) common.FruitResponse {
	log.Printf("Deleting Fruit(s) (%+v)\n", req)

	query := receiver.GetDbSearchQuery(req)

	db := GetPostgreSqlConnection()

	dbContext := common.GetContext()
	_, err := db.ExecContext(
		dbContext,
		"DELETE Fruit WHERE "+query)
	common.CheckErrorWithPanic(err, fmt.Sprintf("error while deleting Fruit (%+v)", req))

	log.Printf("Deleted Fruit(s) matching (%+v)\n", req)

	return common.FruitResponse{
		Message: common.RESPOSNE_SUCCESS + " in deleting fruits matching criteria",
		Err:     nil,
	}
}

func (receiver *FruitService) GetFruit(req *common.Fruit) common.Fruit {
	fruitResponse := receiver.GetFruits(req)

	var fruit common.Fruit
	if fruitResponse.Fruits != nil && len(fruitResponse.Fruits) > 0 {
		fruit = fruitResponse.Fruits[0]
	}

	return fruit
}

func (receiver *FruitService) GetFruits(req *common.Fruit) common.FruitResponse {
	log.Printf("Retrieving Fruits (+%v)\n", req)

	selectQuery := receiver.getSelectQuery(req)
	log.Printf("Query: [%s]\n", selectQuery)
	db := GetPostgreSqlConnection()

	rows, err := db.Query(selectQuery)
	common.CheckErrorWithPanic(err, fmt.Sprintf("error while retrieving Fruits (+%v)", req))
	defer rows.Close()

	var fruits []common.Fruit
	for rows.Next() {
		fruit := common.Fruit{}
		err := rows.Scan(&fruit.Id, &fruit.Name, &fruit.Description)
		common.CheckErrorWithPanic(err, "error while trying to decode Fruit")
		fruits = append(fruits, fruit)
	}

	var message = fmt.Sprintf("Found %d fruits", len(fruits))

	if len(fruits) == 0 {
		// Create default "no-result" response
		message = "No fruits found for given query"
		fruits = nil
	}

	response := common.FruitResponse{
		Message: message,
		Fruits:  fruits,
	}

	return response
}

func (receiver *FruitService) getSelectQuery(req *common.Fruit) string {
	var queryStr = "SELECT id, name, description FROM fruit"
	if req.Name != ALL_ROWS {
		query := receiver.GetDbSearchQuery(req)
		queryStr += " WHERE " + query
	}

	return queryStr
}
