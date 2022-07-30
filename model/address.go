package model

import (
	DB "TreinoAgiota/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	DB.Model
	Rua         string `json:"rua" binding:"-"`
	Cep         string `json:"cep" binding:"required"`
	Complemento string `json:"complemento" binding:"required"`
}

func (a *Address) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
