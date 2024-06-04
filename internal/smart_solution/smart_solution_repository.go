package smart_solution

import (
	"example/model"
	"fmt"

	"gorm.io/gorm"
)

type SmartSolutionRepository interface {
	Create(values *model.SmartSolution) (*model.SmartSolution, error)
	Update(id int, values *model.SmartSolution) (*model.SmartSolution, error)
	Delete(id int) (resp *string, err error)
	FindOne(id int) (*model.SmartSolution, error)
	FindAll(offset int, limit int, sort string) ([]model.SmartSolution, int64, error)
}

type smartSolutionRepository struct {
	db *gorm.DB
}

func NewSmartSolutionRepository(db *gorm.DB) smartSolutionRepository {
	return smartSolutionRepository{db}
}

func (repo *smartSolutionRepository) Create(values *model.SmartSolution) (*model.SmartSolution, error) {
	if err := repo.db.Create(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *smartSolutionRepository) Update(id int, values *model.SmartSolution) (*model.SmartSolution, error) {
	// uid := uint(id)
	// err := repo.db.Model(&model.SolutionGroup{SmartSolutionID: uint(id)}).Association("Banner").Replace(&values.Banner)
	// if err != nil {
	// 	return nil, err
	// }

	// if err := repo.db.Model(&model.ProductArea{}).Where("id =?", id).Updates(values).Error; err != nil {
	// 	return nil, err
	// }
	repo.db.Where(model.SmartSolution{ID: uint(id)}).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&values)

	return values, nil
}

func (repo *smartSolutionRepository) Delete(id int) (resp *string, err error) {
	if err := repo.db.Model(&model.SmartSolution{}).Where("id =?", id).Delete(&model.SmartSolution{}).Error; err != nil {
		return nil, err
	}

	status := "ok"

	return &status, nil
}
func (repo *smartSolutionRepository) FindOne(id int) (*model.SmartSolution, error) {
	var values model.SmartSolution
	if err := repo.db.Preload("SolutionGroups").Where(&model.SmartSolution{ID: uint(id)}).First(&values).Error; err != nil {
		return nil, err
	}

	return &values, nil
}
func (repo *smartSolutionRepository) FindAll(offset int, limit int, sort string) ([]model.SmartSolution, int64, error) {
	var list []model.SmartSolution
	query := repo.db.Model(&model.SmartSolution{})

	err := query.Preload("SolutionGroups").Limit(limit).Offset(offset).Order(fmt.Sprintf("created_at %s", sort)).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	var count int64
	repo.db.Model(&model.SmartSolution{}).Where(query).Count(&count)

	return list, count, nil
}
