package service

import (
	"time"

	"github.com/elgiavilla/mc_user/models"
	"github.com/elgiavilla/mc_user/users"
)

type MongoService struct {
	repo           users.Repository
	contextTimeout time.Duration
}

func NewService(r users.Repository, t time.Duration) users.Service {
	return &MongoService{
		repo:           r,
		contextTimeout: t,
	}
}

func (s *MongoService) Find(id models.ID) (*models.User, error) {
	return s.repo.Find(id)
}

func (s *MongoService) FindAll() ([]*models.User, error) {
	return s.repo.FindAll()
}

func (s *MongoService) Store(b *models.User) (models.ID, error) {
	b.ID = models.NewID()
	return s.repo.Store(b)
}

func (s *MongoService) Delete(id models.ID) error {
	return s.repo.Delete(id)
}
