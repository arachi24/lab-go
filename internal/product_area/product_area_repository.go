package productarea

import (
	"example/model"

	"gorm.io/gorm"
)

type ProductareaRepository interface {
	Create(values *model.ProductArea) (*model.ProductArea, error)
	Update(id int, values *model.ProductArea) (*model.ProductArea, error)
	Delete(id int) (resp *string, err error)
	FindOne(id int) (*model.ProductArea, error)
	FindAll() ([]*model.ProductArea, error)
}

type productareaRepository struct {
	db *gorm.DB
}

func NewProductreaRepository(db *gorm.DB) productareaRepository {
	return productareaRepository{db}
}

func (repo *productareaRepository) Create(values *model.ProductArea) (*model.ProductArea, error) {
	if err := repo.db.Create(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *productareaRepository) Update(id int, values *model.ProductArea) (*model.ProductArea, error) {
	uid := uint(id)
	err := repo.db.Model(&model.Banner{ProductAreaID: &uid}).Association("Banner").Replace(&values.Banner)
	if err != nil {
		return nil, err
	}

	if err := repo.db.Model(&model.ProductArea{}).Where("id =?", id).Updates(values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *productareaRepository) Delete(id int) (resp *string, err error) {
	if err := repo.db.Model(&model.ProductArea{}).Where("id =?", id).Delete(&model.ProductArea{}).Error; err != nil {
		return nil, err
	}

	status := "ok"

	return &status, nil
}
func (repo *productareaRepository) FindOne(id int) (*model.ProductArea, error) {
	var values model.ProductArea
	if err := repo.db.Model(&model.ProductArea{}).Preload("Banner").Where("id =?", id).First(&values).Error; err != nil {
		return nil, err
	}

	return &values, nil
}
func (repo *productareaRepository) FindAll() ([]*model.ProductArea, error) {
	var values []*model.ProductArea
	if err := repo.db.Model(&model.ProductArea{}).Preload("Banner").Find(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}
