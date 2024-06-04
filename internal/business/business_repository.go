package business

import (
	"example/model"

	"gorm.io/gorm"
)

type BusinessRepository interface {
	Create(values *model.Business) (*model.Business, error)
	Update(id int, values *model.Business) (*model.Business, error)
	Delete(id int) (resp *string, err error)
	FindOne(id int) (*model.Business, error)
	FindAll() ([]*model.Business, error)
}

type businessRepository struct {
	db *gorm.DB
}

func NewBusinessRepository(db *gorm.DB) businessRepository {
	return businessRepository{db}
}

func (repo *businessRepository) Create(values *model.Business) (*model.Business, error) {
	if err := repo.db.Create(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *businessRepository) Update(id int, values *model.Business) (*model.Business, error) {
	if err := repo.db.Model(&model.Business{}).Where("id =?", id).Updates(values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *businessRepository) Delete(id int) (resp *string, err error) {
	if err := repo.db.Model(&model.Business{}).Where("id =?", id).Delete(&model.Business{}).Error; err != nil {
		return nil, err
	}

	status := "ok"

	return &status, nil
}
func (repo *businessRepository) FindOne(id int) (*model.Business, error) {
	var values model.Business
	if err := repo.db.Model(&model.Business{}).Where("id =?", id).First(&values).Error; err != nil {
		return nil, err
	}

	return &values, nil
}
func (repo *businessRepository) FindAll() ([]*model.Business, error) {
	var values []*model.Business
	if err := repo.db.Model(&model.Business{}).Find(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}
