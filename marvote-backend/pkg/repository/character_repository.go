package repository

import "github.com/SeansC12/marvote/pkg/model"

type CharacterRepository struct {
	characters []model.CharacterInfo
}

func NewCharacterRepository() *CharacterRepository {
	var ci model.CharacterInfo
	allCharacters := make([]model.CharacterInfo, 0, 1)
	ci = model.CharacterInfo{
		Id:   0,
		Name: "Spiderman",
		Aka:  "Peter Parker",
	}
	allCharacters = append(allCharacters, ci)
	ci = model.CharacterInfo{
		Id:   1,
		Name: "Doctor Strange",
		Aka:  "Stephen Strange",
	}
	allCharacters = append(allCharacters, ci)
	return &CharacterRepository{characters: allCharacters}
}

func (cs *CharacterRepository) FindAll() ([]model.CharacterInfo, error) {
	return cs.characters, nil
}

func (cs *CharacterRepository) FindById(characterId int) (model.CharacterInfo, error) {
	for _, value := range cs.characters {
		if value.Id == characterId {
			return value, nil
		}
	}
	return model.CharacterInfo{}, nil
}
