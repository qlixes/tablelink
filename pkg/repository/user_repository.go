package repo

import (
	"tablelink/internal/models"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

type UserRepoInterface interface {
	Create(models.User) (models.User, error)
	Update(models.User) error
	Delete(id int64) error
	GetByEmail(email string) (models.User, error)
	Get() []models.User
}

func New(db *gorm.DB) UserRepoInterface {
	return &Repo{db}
}

func (repo *Repo) Create(user models.User) (models.User, error) {
	err := repo.db.Create(&user).Error

	return user, err
}

func (repo *Repo) Update(user models.User) error {
	var users models.User
	if err := repo.db.Where("id = ?", user.ID).First(&users).Error; err != nil {
		return err
	}
	users.Name = user.Name
	err := repo.db.Save(users).Error

	return err
}

func (repo *Repo) Delete(id int64) error {
	err := repo.db.Where("id = ?", id).Delete(&models.User{}).Error

	return err
}

func (repo *Repo) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := repo.db.Where("email = ?", email).First(&user).Error

	return user, err
}

func (repo *Repo) Get() []models.User {
	users := []models.User{}

	repo.db.Find(&users)

	return users
}
