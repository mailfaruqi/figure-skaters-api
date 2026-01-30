package services

import (
	"figure-skaters-api/models"
	"figure-skaters-api/repositories"
)

type ElementService struct {
repo *repositories.ElementRepository
}

func NewElementService(repo *repositories.ElementRepository) *ElementService {
return &ElementService{repo: repo}
}

func (s *ElementService) GetAll() ([]models.Element, error) {
return s.repo.GetAll()
}

func (s *ElementService) GetByID(id int) (*models.ElementDetail, error) {
return s.repo.GetByID(id)
}

func (s *ElementService) Create(element *models.Element) error {
return s.repo.Create(element)
}

func (s *ElementService) Update(element *models.Element) error {
return s.repo.Update(element)
}

func (s *ElementService) Delete(id int) error {
return s.repo.Delete(id)
}