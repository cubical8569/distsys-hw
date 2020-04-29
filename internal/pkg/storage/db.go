package storage

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	db *gorm.DB
}

func (db *DB) AddProduct(product *models.Product) error {
	return db.db.Create(product).Error
}

func (db *DB) Products() ([]models.Product, error) {
	var products []models.Product
	if err := db.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func NewDB() (*DB, error) {
	db, err := gorm.Open("postgres", "host=172.17.0.2 port=5432 user=postgres dbname=postgres password=mysecretpassword sslmode=disable")
	if err != nil {
		return nil, err
	}
	// TODO: close db

	db.AutoMigrate(&models.Product{})

	return &DB{db: db}, nil
}
