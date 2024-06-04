package homepage

import (
	"example/model"

	"gorm.io/gorm"
)

type HomepageRepository interface {
	Create(values *model.Homepage) (*model.Homepage, error)
	Update(id int, values *model.Homepage) (*model.Homepage, error)
	Delete(id int) (resp *string, err error)
	FindOne(id int) (*model.Homepage, error)
	FindAll() ([]*model.Homepage, error)
}

type homepageRepository struct {
	db *gorm.DB
}

func NewHomepageRepository(db *gorm.DB) homepageRepository {
	return homepageRepository{db}
}

func (repo *homepageRepository) Create(values *model.Homepage) (*model.Homepage, error) {
	if err := repo.db.Create(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *homepageRepository) Update(id int, values *model.Homepage) (*model.Homepage, error) {
	if err := repo.db.Model(&model.Homepage{}).Where("id =?", id).Updates(values).Error; err != nil {
		return nil, err
	}

	return values, nil
}
func (repo *homepageRepository) Delete(id int) (resp *string, err error) {
	if err := repo.db.Model(&model.Homepage{}).Where("id =?", id).Delete(&model.Homepage{}).Error; err != nil {
		return nil, err
	}

	status := "ok"

	return &status, nil
}
func (repo *homepageRepository) FindOne(id int) (*model.Homepage, error) {
	var values model.Homepage
	if err := repo.db.Preload("MainBanner", "type =?", "MAIN").Preload("SubBanner", "type =?", "SUB").Model(&model.Homepage{}).Where("id =?", id).First(&values).Error; err != nil {
		return nil, err
	}

	return &values, nil
}
func (repo *homepageRepository) FindAll() ([]*model.Homepage, error) {
	var values []*model.Homepage
	if err := repo.db.Preload("MainBanner", "type =?", "MAIN").Preload("SubBanner", "type =?", "SUB").Model(&model.Homepage{}).Find(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}
