package services

import (
	"time"
	"user_service/models"
	"user_service/repository"
	"user_service/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type ServiceUser interface {
	RegisterUser(user models.RegisterUser) error
	LoginUser(login models.LoginUser) (string, error)
}

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo}
}

func (s *UserService) RegisterUser(user models.RegisterUser) error {

	data, _ := s.repo.Find(user.Email)

	if data.Email != "" {
		return fiber.NewError(fiber.StatusBadRequest, "user already created")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err

	}

	user.Password = string(hash)
	user.Role = "user"
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	if err := s.repo.Create(user); err != nil {
		return err

	}

	return nil
}

func (s *UserService) LoginUser(login models.LoginUser) (string, error) {
	user, err := s.repo.Find(login.Email)

	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return "", err
	}

	token, err := utils.TokenJwt(user)

	if err != nil {
		return "", err
	}

	return token, nil

}
