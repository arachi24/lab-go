package product

import (
	"example/model"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(values *model.Product) (*model.Product, error)
	Update(id int, values *model.Product) (*model.Product, error)
	Delete(id int) (resp *string, err error)
	FindOne(id int) (*model.Product, error)
	FindAll(offset int, limit int, sort string) ([]model.Product, int64, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) productRepository {
	return productRepository{db}
}

func (repo *productRepository) Create(values *model.Product) (*model.Product, error) {
	if err := repo.db.Create(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *productRepository) Delete(id int) (resp *string, err error) {
	if err := repo.db.Model(&model.Product{}).Where("id =?", id).Delete(&model.Product{}).Error; err != nil {
		return nil, err
	}

	status := "ok"

	return &status, nil
}

func (repo *productRepository) Update(id int, values *model.Product) (*model.Product, error) {
	// err := repo.db.Model(&model.ProductDetail{ProductID: uint(id)}).Updates(values.ProductDetail).Error
	// if err != nil {
	// 	return nil, err
	// }

	// if err := repo.db.Model(&model.Product{}).Where("id =?", id).Updates(values).Error; err != nil {
	// 	return nil, err
	// }
	repo.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&values)

	return values, nil
}

func (repo *productRepository) FindOne(id int) (*model.Product, error) {
	var values model.Product
	if err := repo.db.Preload("ProductDetail").Where(&model.Product{ID: uint(id)}).First(&values).Error; err != nil {
		return nil, err
	}

	return &values, nil
}
func (repo *productRepository) FindAll(offset int, limit int, sort string) ([]model.Product, int64, error) {
	var list []model.Product
	query := repo.db.Model(&model.Product{})

	err := query.Preload("ProductDetail").Limit(limit).Offset(offset).Order(fmt.Sprintf("created_at %s", sort)).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	var count int64
	repo.db.Model(&model.Product{}).Where(query).Count(&count)

	return list, count, nil
}
