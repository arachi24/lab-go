package location

import (
	"example/model"

	"gorm.io/gorm"
)

type LocationRepository interface {
	Create(values *model.Location) (*model.Location, error)
	Update(id int, values *model.Location) (*model.Location, error)
	Delete(id int) (resp *string, err error)
	FindOne(id int) (*model.Location, error)
	FindAll() ([]*model.Location, error)
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) locationRepository {
	return locationRepository{db}
}

func (repo *locationRepository) Create(values *model.Location) (*model.Location, error) {
	if err := repo.db.Create(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *locationRepository) Update(id int, values *model.Location) (*model.Location, error) {
	if err := repo.db.Model(&model.Location{}).Where("id =?", id).Updates(values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *locationRepository) Delete(id int) (resp *string, err error) {
	if err := repo.db.Model(&model.Location{}).Where("id =?", id).Delete(&model.Location{}).Error; err != nil {
		return nil, err
	}

	status := "ok"

	return &status, nil
}
func (repo *locationRepository) FindOne(id int) (*model.Location, error) {
	var values model.Location
	if err := repo.db.Model(&model.Location{}).Where("id =?", id).First(&values).Error; err != nil {
		return nil, err
	}

	return &values, nil
}
func (repo *locationRepository) FindAll() ([]*model.Location, error) {
	var values []*model.Location
	if err := repo.db.Model(&model.Location{}).Find(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}
