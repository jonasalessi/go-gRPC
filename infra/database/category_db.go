package database

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type CategoryEntity struct {
	Id          string
	Name        string
	Description string
}

type CategoryDB struct {
	Data []*CategoryEntity
}

var CategoryNotFoundError = errors.New("Category was not found")

func (c *CategoryDB) Save(name, description string) (category *CategoryEntity) {
	id := uuid.New().String()
	category = &CategoryEntity{
		Id:          id,
		Name:        name,
		Description: description,
	}
	c.Data = append(c.Data, category)
	fmt.Println("Now it has ", len(c.Data), " records!")
	return
}

func (c *CategoryDB) GetById(id string) (*CategoryEntity, error) {
	for _, cat := range c.Data {
		if cat.Id == id {
			return cat, nil
		}
	}
	return nil, CategoryNotFoundError
}

func (c *CategoryDB) GetAll() []*CategoryEntity {
	return c.Data
}
