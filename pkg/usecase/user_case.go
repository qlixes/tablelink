package usecase

type UserCase struct{
	repo repository.UserRepoInterface
}

type UserCaseInterface struct {
	Login(email string, password string) models.User
	Store(role_id int64, name string, email string, password string) error
	Update(id int64, name string) error
	Delete(id int64) error
	Show() []models.User
}

func (uc *UserCase) Login(email string, password string) models.User {

}

func (uc *UserCase) Store(role_id int64, name string, email string, password string) error {

}

func (uc *UserCase) Update(id int64, name string) error {

}

func (uc *UserCase) Delete(id int64) error {

}

func (uc *UserCase) Show() []models.User {

}