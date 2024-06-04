package invertorrelation

import (
	"example/model"

	"gorm.io/gorm"
)

type InvestorRelationRepository interface {
	Create(values *model.InvestorRelation) (*model.InvestorRelation, error)
	Update(id int, values *model.InvestorRelation) (*model.InvestorRelation, error)
	Delete(id int) (resp *string, err error)
	FindOne(id int) (*model.InvestorRelation, error)
	FindAll() ([]*model.InvestorRelation, error)
}

type investorRelationRepository struct {
	db *gorm.DB
}

func NewInvestorRelationRepository(db *gorm.DB) investorRelationRepository {
	return investorRelationRepository{db}
}

func (repo *investorRelationRepository) Create(values *model.InvestorRelation) (*model.InvestorRelation, error) {
	if err := repo.db.Create(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *investorRelationRepository) Update(id int, values *model.InvestorRelation) (*model.InvestorRelation, error) {
	err := repo.db.Model(&model.InvestorRelation{ID: uint(id)}).Association("Banner").Replace(&values.Banner)
	if err != nil {
		return nil, err
	}

	if err := repo.db.Model(&model.InvestorRelation{}).Where("id =?", id).Updates(values).Error; err != nil {
		return nil, err
	}

	return values, nil
}
func (repo *investorRelationRepository) Delete(id int) (resp *string, err error) {
	if err := repo.db.Model(&model.InvestorRelation{}).Where("id =?", id).Delete(&model.InvestorRelation{}).Error; err != nil {
		return nil, err
	}

	status := "ok"

	return &status, nil
}
func (repo *investorRelationRepository) FindOne(id int) (*model.InvestorRelation, error) {
	var values model.InvestorRelation
	if err := repo.db.Preload("Banner").Model(&model.InvestorRelation{}).Where("id =?", id).First(&values).Error; err != nil {
		return nil, err
	}

	return &values, nil
}
func (repo *investorRelationRepository) FindAll() ([]*model.InvestorRelation, error) {
	var values []*model.InvestorRelation
	if err := repo.db.Preload("Banner").Model(&model.InvestorRelation{}).Find(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}
