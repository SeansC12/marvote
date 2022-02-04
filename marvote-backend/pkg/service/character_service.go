package service

import (
	"context"

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

func (cs *CharacterService) GetAll(ctx context.Context) ([]model.CharacterInfo, error) {
	characters, err := cs.characterRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return characters, nil
}

func (cs *CharacterService) Get(ctx context.Context, characterId string) (model.CharacterInfo, error) {
	marvelCharacter, err := cs.characterRepository.FindById(ctx, characterId)
	if err != nil {
		return model.CharacterInfo{}, err
	}
	return marvelCharacter, nil
}

func (cs *CharacterService) Save(ctx context.Context, charInfo model.CharacterInfo) (model.CharacterInfo, error) {
	result, err := cs.characterRepository.Save(ctx, charInfo)
	if err != nil {
		return model.CharacterInfo{}, err
	}
	return result, nil
}
