package services

import (
	"errors"
	"net/http"

	"github.com/Alwin18/king-code/entity"
	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/repositories"
	"github.com/Alwin18/king-code/utils"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) RegisterUser(c *gin.Context, user *models.User) error {
	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	existData, err := s.GetUserByEmail(user.Email)
	if err != nil {
		if err.Error() == "record not found" {

		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return err
		}
	}

	if existData != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return errors.New("Email already exists")
	}

	// manipulate user data
	user.Default()
	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return err
	}
	user.Password = hashedPassword

	return s.Repo.CreateUser(user)
}

func (s *UserService) GetUserByID(c *gin.Context, id string) (*entity.GetUserByIdResponse, error) {
	data, err := s.Repo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}

	response := &entity.GetUserByIdResponse{
		ID:        id,
		Username:  data.Username,
		Email:     data.Email,
		XP:        data.XP,
		Level:     data.Level,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return response, nil
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.Repo.GetUserByEmail(email)
}
