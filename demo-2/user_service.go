package main

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *User) error {
	return s.repo.Create(user)
}

func (s *UserService) ListUsers() ([]User, error) {
	return s.repo.List()
}
