package services

import (
	"context"

	"github.com/Alieksieiev0/user-service/internal/models"
	"gorm.io/gorm"
)

type UserService interface {
	GetById(ctx context.Context, id string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	Save(ctx context.Context, user *models.User) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		db: db,
	}
}

func (us *userService) GetById(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	return user, us.db.First(user, id).Error
}

func (us *userService) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}
	return user, us.db.First(user, "username = ?", username).Error
}

func (us *userService) Save(ctx context.Context, user *models.User) error {
	return us.db.Save(user).Error
}
