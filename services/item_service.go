package services

import (
	"github.com/shingoasou-0804/oshin-go-gin/models"
	"github.com/shingoasou-0804/oshin-go-gin/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindById(itemId uint) (*models.Item, error) {
	return s.repository.FindById(itemId)
}
