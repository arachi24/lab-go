package banner

import (
	"example/model"

	"gorm.io/gorm"
)

type BannerRepository interface {
	BulkCreate(values []model.Banner) error
	Update(id int, values *model.Banner) (*model.Banner, error)
	// Delete(id int) (resp *string, err error)
	// FindOne(id int) (*model.Homepage, error)
	// FindAll() ([]*model.Homepage, error)
}

type bannerRepository struct {
	db *gorm.DB
}

func NewBannerRepository(db *gorm.DB) bannerRepository {
	return bannerRepository{db}
}

func (repo *bannerRepository) BulkCreate(values []model.Banner) error {
	if len(values) > 0 {
		if err := repo.db.Create(&values).Error; err != nil {
			return err
		}
	}

	return nil
}

func (repo *bannerRepository) Update(id int, values *model.Banner) (*model.Banner, error) {
	if err := repo.db.Model(&model.Banner{}).Where("id =?", id).Updates(values).Error; err != nil {
		return nil, err
	}

	return values, nil
}
