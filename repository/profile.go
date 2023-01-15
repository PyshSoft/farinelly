package repository

import "github.com/PyshSoft/farinelly/entity"

type ProfileRepository interface {
	Create(entity.Profile) error
	GetById(string) (entity.Profile, bool)
	GetByGenger() []entity.Profile
}