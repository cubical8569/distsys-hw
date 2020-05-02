package storage

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	db *gorm.DB
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

func (db *DB) AddProduct(product *models.Product) error {
	return db.db.Create(product).Error
}

func (db *DB) Products(params *GetParams) ([]models.Product, error) {
	var products []models.Product

	var err error
	if params.Limit != nil && params.Offset != nil {
		err = db.db.Offset(*params.Offset).Limit(*params.Limit).Find(&products).Error
	} else if params.Offset != nil {
		err = db.db.Offset(*params.Offset).Find(&products).Error
	} else if params.Limit != nil {
		err = db.db.Limit(*params.Limit).Find(&products).Error
	} else {
		err = db.db.Find(&products).Error
	}

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (db *DB) GetProduct(id int) (*models.Product, error) {
	var product models.Product
	if err := db.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (db *DB) UpdateProduct(product *models.Product) error {
	return db.db.Save(product).Error
}

func (db *DB) DeleteProduct(id uint) error {
	return db.db.Delete(&models.Product{
		Model: gorm.Model{ID: id},
	}).Error
}
