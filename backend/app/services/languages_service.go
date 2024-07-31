package services

import (
	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/linguai/app/models"
)

type LanguagesService struct {
	raptor.Service
}

func (ts *LanguagesService) All() []models.Language {
	return []models.Language{
		{
			Code: "hr",
			Name: "Croatian",
		},
		{
			Code: "cz",
			Name: "Czech",
		},
		{
			Code: "en",
			Name: "English",
		},
		{
			Code: "de",
			Name: "German",
		},
	}
}