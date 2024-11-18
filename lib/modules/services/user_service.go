package services

import logger "GoPorject/log"

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetWelcomeMessage() string {
	logger.Info("UserService.GetWelcomeMessage")
	return "Welcome to your first Go service!"
}
