package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

var items = []item{
	{ID: "1", Name: "Blue shirt", Type: "shirt", Description: "A blue shirt", Quantity: 3},
	{ID: "2", Name: "Green shirt", Type: "shirt", Description: "A green shirt", Quantity: 5},
	{ID: "3", Name: "Black shorts", Type: "shorts", Description: "A black short", Quantity: 2},
	{ID: "4", Name: "Red jacket", Type: "jacket", Description: "A red jacket", Quantity: 1},
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func lookUpItemById(c *gin.Context) {
	id := c.Param("id")
	item, err := getItemById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, item)
}

func getItemById(id string) (*item, error) {
	for i, item := range items {
		if item.ID == id {
			return &items[i], nil
		}
	}
	return nil, errors.New("item_not_found")
}

func createItem(c *gin.Context) {
	var newItem item
	err := c.BindJSON(&newItem)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Incorrect payload"})
		return
	}
	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

func searchItemByType(c *gin.Context) {
	itemType, hasTypeQuery := c.GetQuery("type")

	if hasTypeQuery == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing type query"})
		return
	}

	itemList, err := getItemByType(itemType)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, itemList)
}

func getItemByType(itemType string) ([]item, error) {
	newItems := []item{}
	for i, item := range items {
		if item.Type == itemType {
			newItems = append(newItems, items[i])
		}
	}
	if len(newItems) != 0 {
		return newItems, nil
	}
	return nil, errors.New("item_not_found")
}
func main() {
	router := gin.Default()
	router.GET("/items", getItems)
	router.GET("/items/:id", lookUpItemById)
	router.POST("/items", createItem)
	router.GET("/items/search", searchItemByType)
	router.Run("localhost:8080")
}
