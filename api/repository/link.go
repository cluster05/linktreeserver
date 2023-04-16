package repository

import (
	"github.com/cluster05/linktree/api/model"
	"github.com/jmoiron/sqlx"
	"time"
)

type LinkRepository interface {
	FindLink(authId string, linkId string) (model.Link, error)
	CreateLink(model.Link) (model.Link, error)
	ReadLink(string) ([]model.Link, error)
	UpdateLink(model.Link) (model.Link, error)
	DeleteLink(string) error
}

type linkRepository struct {
	MySqlDB *sqlx.DB
}

type LinkRepositoryConfig struct {
	MySqlDB *sqlx.DB
}

func NewLinkRepository(config *LinkRepositoryConfig) LinkRepository {
	return &linkRepository{
		MySqlDB: config.MySqlDB,
	}
}

func (repo *linkRepository) FindLink(authId string, linkId string) (model.Link, error) {
	findLinkQuery := `SELECT linkId,title,url,imageUrl
		FROM link
		WHERE  authId=? AND linkId=? AND isDeleted=false;
	`
	link := model.Link{}
	err := repo.MySqlDB.Get(&link, findLinkQuery, authId, linkId)
	if err != nil {
		return model.Link{}, err
	}

	return link, err

}

func (repo *linkRepository) CreateLink(link model.Link) (model.Link, error) {
	createLinkQuery := `INSERT INTO 
		link(authId,linkId,title,url,imageUrl,createdAt,updatedAt,isDeleted) VALUES
		(:authId,:linkId,:title,:url,:imageUrl,:createdAt,:updatedAt,:isDeleted);`

	_, err := repo.MySqlDB.NamedExec(createLinkQuery, link)
	if err != nil {
		return model.Link{}, err
	}
	return link, nil
}

func (repo *linkRepository) ReadLink(userId string) ([]model.Link, error) {
	readLinkQuery := `SELECT linkId,title,url,imageUrl
		FROM link
		WHERE authId=? AND isDeleted=false;
		`
	links := []model.Link{}
	err := repo.MySqlDB.Select(&links, readLinkQuery, userId)
	if err != nil {
		return []model.Link{}, err
	}
	return links, nil
}

func (repo *linkRepository) UpdateLink(link model.Link) (model.Link, error) {
	updateLinkQuery := `UPDATE link 
		SET title=? ,
			url=?,
			image=?,
			updatedAt=?,
		WHERE linkId=? isDeleted=false;
		`
	_, err := repo.MySqlDB.Exec(updateLinkQuery, link.Title, link.URL, link.ImageUrl, link.UpdatedAt, link.LinkId)
	if err != nil {
		return model.Link{}, err
	}
	return link, nil
}

func (repo *linkRepository) DeleteLink(linkId string) error {
	deleteLinkQuery := `UPDATE link
	SET isDeleted=true,
		updatedAt=?
	WHERE linkId=? isDeleted=false;
	`
	_, err := repo.MySqlDB.Exec(deleteLinkQuery, time.Now().Unix(), linkId)
	if err != nil {
		return err
	}
	return nil
}
