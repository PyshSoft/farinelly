package service

import (
	"github.com/PyshSoft/farinelly/entity"
	"github.com/PyshSoft/farinelly/repository"
)

type ProfileService struct {
	repo repository.ProfileRepository
}

func NewProfileService(repo repository.ProfileRepository) *ProfileService {
	return &ProfileService{repo}
}

func (ps *ProfileService) Create(profile entity.Profile) error {
	return ps.repo.Create(profile)
}