package service

import "github.com/SeansC12/marvote/pkg/model"

type ICharacterService interface {
	GetAll() ([]model.CharacterInfo, error)
	Get(characterId int) (model.CharacterInfo, error)
	Save(model.CharacterInfo) (model.CharacterInfo, error)
}
