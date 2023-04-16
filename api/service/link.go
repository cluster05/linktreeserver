package service

import (
	"errors"
	"time"

	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"

	"github.com/cluster05/linktree/api/model"
	"github.com/cluster05/linktree/api/repository"
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

func NewLinkService(config *LinkServiceConfig) LinkService {
	return &linkService{
		linkRepository: config.LinkRepository,
	}
}

func (ls *linkService) CreateLink(user model.JWTPayload, createLinkDTO model.CreateLinkDTO) (model.Link, error) {

	link := model.Link{
		LinkId:   shortuuid.New(),
		AuthId:   user.AuthId,
		Title:    createLinkDTO.Title,
		URL:      createLinkDTO.URL,
		ImageUrl: createLinkDTO.URL,
	}

	return ls.linkRepository.CreateLink(link)

}

func (ls *linkService) ReadLink(user model.JWTPayload) ([]model.Link, error) {
	return ls.linkRepository.ReadLink(user.AuthId)
}

func (ls *linkService) UpdateLink(user model.JWTPayload, updateLinkDTO model.UpdateLinkDTO) (model.Link, error) {

	findLink, err := ls.linkRepository.FindLink(user.AuthId, updateLinkDTO.LinkId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Link{}, gorm.ErrRecordNotFound
	}

	link := model.Link{
		LinkId:    updateLinkDTO.LinkId,
		AuthId:    user.AuthId,
		Title:     updateLinkDTO.Title,
		URL:       updateLinkDTO.URL,
		ImageUrl:  updateLinkDTO.URL,
		CreatedAt: findLink.CreatedAt,
		UpdatedAt: time.Now().Unix(),
	}
	return ls.linkRepository.UpdateLink(link)
}

func (ls *linkService) DeleteLink(user model.JWTPayload, deleteLinkDTO model.DeleteLinkDTO) error {

	_, err := ls.linkRepository.FindLink(user.AuthId, deleteLinkDTO.LinkId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	return ls.linkRepository.DeleteLink(deleteLinkDTO.LinkId)
}
