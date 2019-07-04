package users

import (
	"github.com/elgiavilla/mc_user/models"
)

type Repository interface {
	Find(id models.ID) (*models.User, error)
	FindAll() ([]*models.User, error)
	Store(b *models.User) (models.ID, error)
	Delete(id models.ID) error
}
