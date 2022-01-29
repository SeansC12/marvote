package repository

import "github.com/SeansC12/marvote/pkg/model"

type ICharacterRepository interface {
	FindAll() ([]model.CharacterInfo, error)
	FindById(characterId int) (model.CharacterInfo, error)
}
