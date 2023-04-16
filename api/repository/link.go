package repository

import (
	"gorm.io/gorm"

	"github.com/cluster05/linktree/api/model"
)

type LinkRepository interface {
	FindLink(authId string, linkId string) (model.Link, error)
	CreateLink(model.Link) (model.Link, error)
	ReadLink(string) ([]model.Link, error)
	UpdateLink(model.Link) (model.Link, error)
	DeleteLink(string) error
}

type linkRepository struct {
	MySqlDB *gorm.DB
}

type LinkRepositoryConfig struct {
	MySqlDB *gorm.DB
}

func NewLinkRepository(config *LinkRepositoryConfig) LinkRepository {
	return &linkRepository{
		MySqlDB: config.MySqlDB,
	}
}

func (repo *linkRepository) FindLink(authId string, linkId string) (model.Link, error) {
	link := model.Link{}
	result := repo.MySqlDB.Where("authId=? AND linkId=? AND isDeleted=false", authId, linkId).Find(&link)
	if result.Error != nil {
		return model.Link{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Link{}, gorm.ErrRecordNotFound
	}

	return link, nil
}

func (repo *linkRepository) CreateLink(link model.Link) (model.Link, error) {
	result := repo.MySqlDB.Create(&link)
	if result.Error != nil {
		return model.Link{}, result.Error
	}
	return link, nil
}

func (repo *linkRepository) ReadLink(userId string) ([]model.Link, error) {
	links := []model.Link{}
	result := repo.MySqlDB.Where("authId=? AND isDeleted=false", userId).Find(&links)

	if result.Error != nil {
		return []model.Link{}, result.Error
	}
	return links, nil
}

func (repo *linkRepository) UpdateLink(link model.Link) (model.Link, error) {
	result := repo.MySqlDB.Model(&model.Link{}).
		Where("linkId=? AND isDeleted=false", link.LinkId).
		Updates(map[string]interface{}{"title": link.Title, "url": link.URL, "imageUrl": link.ImageUrl})
	if result.Error != nil {
		return model.Link{}, result.Error
	}
	return link, nil
}

func (repo *linkRepository) DeleteLink(linkId string) error {
	result := repo.MySqlDB.
		Where("linkId=? AND isDeleted=false", linkId).
		Delete(&model.Link{})

	if result.Error != nil {
		return result.Error
	}
	return nil
}
