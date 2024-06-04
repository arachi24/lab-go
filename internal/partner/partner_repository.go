package partner

import (
	"example/model"

	"gorm.io/gorm"
)

type PartnerRepository interface {
	Create(values *model.Partner) (*model.Partner, error)
	Update(id int, values *model.Partner) (*model.Partner, error)
	Delete(id int) (resp *string, err error)
	FindOne(id int) (*model.Partner, error)
	FindAll() ([]*model.Partner, error)
}

type partnerRepository struct {
	db *gorm.DB
}

func NewPartnerRepository(db *gorm.DB) partnerRepository {
	return partnerRepository{db}
}

func (repo *partnerRepository) Create(values *model.Partner) (*model.Partner, error) {
	if err := repo.db.Create(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *partnerRepository) Update(id int, values *model.Partner) (*model.Partner, error) {
	if err := repo.db.Model(&model.Partner{}).Where("id =?", id).Updates(values).Error; err != nil {
		return nil, err
	}

	return values, nil
}

func (repo *partnerRepository) Delete(id int) (resp *string, err error) {
	if err := repo.db.Model(&model.Partner{}).Where("id =?", id).Delete(&model.Partner{}).Error; err != nil {
		return nil, err
	}

	status := "ok"

	return &status, nil
}
func (repo *partnerRepository) FindOne(id int) (*model.Partner, error) {
	var values model.Partner
	if err := repo.db.Model(&model.Partner{}).Where("id =?", id).First(&values).Error; err != nil {
		return nil, err
	}

	return &values, nil
}
func (repo *partnerRepository) FindAll() ([]*model.Partner, error) {
	var values []*model.Partner
	if err := repo.db.Model(&model.Partner{}).Find(&values).Error; err != nil {
		return nil, err
	}

	return values, nil
}
