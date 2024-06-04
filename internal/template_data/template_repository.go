package template

import (
	"example/model"

	"gorm.io/gorm"
)

type TemplateHandlerRepository interface {
	Create(values *model.Template) (*model.Template, error)
	FindOne(filter *ThemeQuery) (*model.Template, error)
	Update(id int, values *model.Template) (*model.Template, error)
}

type templateHandlerRepository struct {
	db *gorm.DB
}

func NewTemplateRepository(db *gorm.DB) templateHandlerRepository {
	return templateHandlerRepository{db}
}

func (repo *templateHandlerRepository) Create(values *model.Template) (*model.Template, error) {
	if err := repo.db.Create(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *templateHandlerRepository) Update(id int, values *model.Template) (*model.Template, error) {

	if err := repo.db.Model(&model.Template{}).Where("id =?", id).Updates(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *templateHandlerRepository) Delete(id int) (resp *string, err error) {
	if err := repo.db.Model(&model.Template{}).Where("id =?", id).Delete(&model.Template{}).Error; err != nil {
		return nil, err
	}

	status := "ok"

	return &status, nil
}
func (repo *templateHandlerRepository) FindOne(filter *ThemeQuery) (*model.Template, error) {
	var values model.Template
	themeModel := repo.db.Model(&model.Template{})
	if filter != nil {
		if filter.Slug != "" {
			themeModel = themeModel.Where("slug = ?", filter.Slug)
		}
	}
	if err := themeModel.First(&values).Error; err != nil {
		return nil, err
	}

	return &values, nil
}
