package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Category struct {
	db *sql.DB
	ID string
	Name string
	Description *string
}

func NewCategory(db *sql.DB) *Category  {
	return &Category{db: db}
}

func (c *Category) Create(name string, description *string) (Category, error) {
	id := uuid.New().String()
	if description != nil {
		_, err := c.db.Exec(
			"INSERT INTO category (id, name, description) VALUES ($1, $2, $3)", 
			id, name, *description,
		)
		if err != nil {
			return Category{}, err
		}
	} else {
		_, err := c.db.Exec(
			"INSERT INTO category (id, name) VALUES ($1, $2)", 
			id, name,
		)		
		if err != nil {
			return Category{}, err
		}
	}
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) GetByID(id string) (Category, error) {
	var category Category
	err := c.db.QueryRow(
		"SELECT * FROM category WHERE id = $1",
		id,
	).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (c *Category) GetAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT * FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		category := Category{}
		errRow := rows.Scan(&category.ID, &category.Name, &category.Description)
		if errRow != nil {
			return nil, errRow
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *Category) UpdateByID(id string, name string, description string) (error) {
	_, err := c.db.Exec(
		"UPDATE category " + 
		"SET name = $1, " + 
		"description = $2 " + 
		"WHERE id = $3",
		name,
		description,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}