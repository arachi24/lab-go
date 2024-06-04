package news

import (
	"example/model"

	"gorm.io/gorm"
)

type NewsRepository interface {
	Create(values *model.News) (*model.News, error)
	Update(id int, values *model.News) (*model.News, error)
	Delete(id int) (resp *string, err error)
	FindOne(id int) (*model.News, error)
	FindAll() ([]*model.News, error)
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) newsRepository {
	return newsRepository{db}
}

func (repo *newsRepository) Create(values *model.News) (*model.News, error) {
	if err := repo.db.Create(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *newsRepository) Update(id int, values *model.News) (*model.News, error) {
	if err := repo.db.Model(&model.News{}).Where("id =?", id).Updates(values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *newsRepository) Delete(id int) (resp *string, err error) {
	if err := repo.db.Model(&model.News{}).Where("id =?", id).Delete(&model.News{}).Error; err != nil {
		return nil, err
	}

	status := "ok"

	return &status, nil
}
func (repo *newsRepository) FindOne(id int) (*model.News, error) {
	var values model.News
	if err := repo.db.Model(&model.News{}).Where("id =?", id).First(&values).Error; err != nil {
		return nil, err
	}

	return &values, nil
}
func (repo *newsRepository) FindAll() ([]*model.News, error) {
	var values []*model.News
	if err := repo.db.Model(&model.News{}).Find(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}
