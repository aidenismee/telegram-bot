package repository

import (
	"errors"
	"github.com/nekizz/telegram-bot/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	ModelReadByConditionRepository[model.User]
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) ReadByCondition(cond interface{}, preloads ...string) (*model.User, error) {
	var user *model.User

	query := r.db.Model(&model.User{}).Select("*")

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Where(cond).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}
