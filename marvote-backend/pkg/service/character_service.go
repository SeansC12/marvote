package service

import (
	"github.com/SeansC12/marvote/pkg/model"
	"github.com/SeansC12/marvote/pkg/repository"
)

type CharacterService struct {
	characterRepository repository.ICharacterRepository
}

func NewCharacterService(characterRepository repository.ICharacterRepository) *CharacterService {
	return &CharacterService{
		characterRepository: characterRepository,
	}
}

func (cs *CharacterService) GetAll() ([]model.CharacterInfo, error) {
	characters, err := cs.characterRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return characters, nil
}

func (cs *CharacterService) Get(characterId int) (model.CharacterInfo, error) {
	marvelCharacter, err := cs.characterRepository.FindById(characterId)
	if err != nil {
		return model.CharacterInfo{}, err
	}
	return marvelCharacter, nil
}
