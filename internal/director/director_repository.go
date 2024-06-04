package director

import (
	"example/model"

	"gorm.io/gorm"
)

type DirectorRepository interface {
	Create(values *model.Director) (*model.Director, error)
	Update(id int, values *model.Director) (*model.Director, error)
	Delete(id int) (resp *string, err error)
	FindOne(id int) (*model.Director, error)
	FindAll() ([]*model.Director, error)
}

type directorRepository struct {
	db *gorm.DB
}

func NewDirectorRepository(db *gorm.DB) directorRepository {
	return directorRepository{db}
}

func (repo *directorRepository) Create(values *model.Director) (*model.Director, error) {
	if err := repo.db.Create(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *directorRepository) Update(id int, values *model.Director) (*model.Director, error) {
	if err := repo.db.Model(&model.Director{}).Where("id =?", id).Updates(values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *directorRepository) Delete(id int) (resp *string, err error) {
	if err := repo.db.Model(&model.Director{}).Where("id =?", id).Delete(&model.Director{}).Error; err != nil {
		return nil, err
	}

	status := "ok"

	return &status, nil
}
func (repo *directorRepository) FindOne(id int) (*model.Director, error) {
	var values model.Director
	if err := repo.db.Model(&model.Director{}).Where("id =?", id).First(&values).Error; err != nil {
		return nil, err
	}

	return &values, nil
}
func (repo *directorRepository) FindAll() ([]*model.Director, error) {
	var values []*model.Director
	if err := repo.db.Model(&model.Director{}).Find(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}
