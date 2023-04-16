package service

import (
	"database/sql"
	"fmt"
	"github.com/cluster05/linktree/api/model"
	"github.com/cluster05/linktree/api/repository"
	"github.com/google/uuid"
	"time"
)

type LinkService interface {
	CreateLink(model.JWTPayload, model.CreateLinkDTO) (model.Link, error)
	ReadLink(model.JWTPayload) ([]model.Link, error)
	UpdateLink(model.JWTPayload, model.UpdateLinkDTO) (model.Link, error)
	DeleteLink(model.JWTPayload, model.DeleteLinkDTO) error
}

type linkService struct {
	linkRepository repository.LinkRepository
}

type LinkServiceConfig struct {
	LinkRepository repository.LinkRepository
}

var (
	errLinkNotFound = fmt.Errorf("link not found")
)

func NewLinkService(config *LinkServiceConfig) LinkService {
	return &linkService{
		linkRepository: config.LinkRepository,
	}
}

func (ls *linkService) CreateLink(user model.JWTPayload, createLinkDTO model.CreateLinkDTO) (model.Link, error) {

	link := model.Link{
		LinkId:    uuid.NewString(),
		AuthId:    user.AuthId,
		Title:     createLinkDTO.Title,
		URL:       createLinkDTO.URL,
		ImageUrl:  createLinkDTO.URL,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		IsDeleted: false,
	}

	return ls.linkRepository.CreateLink(link)

}

func (ls *linkService) ReadLink(user model.JWTPayload) ([]model.Link, error) {
	return ls.linkRepository.ReadLink(user.AuthId)
}

func (ls *linkService) UpdateLink(user model.JWTPayload, updateLinkDTO model.UpdateLinkDTO) (model.Link, error) {

	_, err := ls.linkRepository.FindLink(user.AuthId, updateLinkDTO.LinkId)
	if err == sql.ErrNoRows {
		return model.Link{}, errLinkNotFound
	}

	link := model.Link{
		LinkId:   updateLinkDTO.LinkId,
		AuthId:   user.AuthId,
		Title:    updateLinkDTO.Title,
		URL:      updateLinkDTO.URL,
		ImageUrl: updateLinkDTO.URL,
	}
	return ls.linkRepository.UpdateLink(link)
}

func (ls *linkService) DeleteLink(user model.JWTPayload, deleteLinkDTO model.DeleteLinkDTO) error {

	_, err := ls.linkRepository.FindLink(user.AuthId, deleteLinkDTO.LinkId)
	if err == sql.ErrNoRows {
		return errLinkNotFound
	}

	return ls.linkRepository.DeleteLink(deleteLinkDTO.LinkId)
}
