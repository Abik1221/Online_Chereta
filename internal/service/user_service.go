package services

import (
    "ONLINE_CHERETA/internal/models"
    "ONLINE_CHERETA/internal/repositories"
    "ONLINE_CHERETA/internal/utils"
    "errors"
)

type UserService struct {
    userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

// RegisterUser handles user registration
func (s *UserService) RegisterUser(username, email, password string) (*models.User, error) {
    // Check if the user already exists
    existingUser, err := s.userRepo.FindByEmail(email)
    if err == nil && existingUser != nil {
        return nil, errors.New("user already exists")
    }

    // Hash the password
    hashedPassword, err := utils.HashPassword(password)
    if err != nil {
        return nil, err
    }

    // Create the user
    user := &models.User{
        Username:     username,
        Email:        email,
        PasswordHash: hashedPassword,
    }

    if err := s.userRepo.Create(user); err != nil {
        return nil, err
    }

    return user, nil
}

// LoginUser handles user login
func (s *UserService) LoginUser(email, password string) (*models.User, error) {
    // Find the user by email
    user, err := s.userRepo.FindByEmail(email)
    if err != nil {
        return nil, errors.New("invalid credentials")
    }

    // Check the password
    if err := utils.CheckPassword(password, user.PasswordHash); err != nil {
        return nil, errors.New("invalid credentials")
    }

    return user, nil
}

// GetUserByID fetches a user by ID
func (s *UserService) GetUserByID(userID uint) (*models.User, error) {
    var user models.User
    if err := s.userRepo.db.First(&user, userID).Error; err != nil {
        return nil, err
    }
    return &user, nil
}