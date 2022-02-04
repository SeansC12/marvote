package service

import (
	"context"
	"testing"

	"github.com/SeansC12/marvote/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CharacterServiceTestSuite struct {
	suite.Suite
	fakeData model.CharacterInfo
}

func (ts *CharacterServiceTestSuite) SetupTest() {
	ts.fakeData = model.CharacterInfo{
		Id:   "0",
		Name: "Spiderman",
		Aka:  "Peter Parker",
	}
}

func (ts *CharacterServiceTestSuite) TestGetAllCharactersSuccess() {
	ctx := context.TODO()
	mockCharacterRepo := new(MockedCharacterRepository)
	service := NewCharacterService(mockCharacterRepo)
	allCharacters := make([]model.CharacterInfo, 0, 1)
	allCharacters = append(allCharacters, ts.fakeData)
	mockCharacterRepo.On("FindAll").Return(allCharacters, nil)
	response, err := service.GetAll(ctx)
	assert.Nil(ts.T(), err)

	assert.Equal(ts.T(), 1, len(response), "Must be of size 1")
	assert.Equal(ts.T(), "Spiderman", response[0].Name, "Must have the same name")
}

func (ts *CharacterServiceTestSuite) TestGetOneCharactersSuccess() {
	ctx := context.TODO()
	mockCharacterRepo := new(MockedCharacterRepository)
	service := NewCharacterService(mockCharacterRepo)
	mockCharacterRepo.On("FindById", "0").Return(ts.fakeData, nil)
	response, err := service.Get(ctx, "0")
	assert.Nil(ts.T(), err)

	assert.Equal(ts.T(), "Spiderman", response.Name, "Must be of the same name")
	assert.Equal(ts.T(), "0", response.Id, "Must have the same id")
}

func (ts *CharacterServiceTestSuite) TestMustSaveCharacters() {
	ctx := context.TODO()
	mockCharacterRepo := new(MockedCharacterRepository)
	service := NewCharacterService(mockCharacterRepo)

	fakeData := model.CharacterInfo{
		Id:   "0",
		Name: "Daredevil",
		Aka:  "Matt Murdock",
	}
	charInfo := model.CharacterInfo{
		Name: "Daredevil",
		Aka:  "Matt Murdock",
	}
	mockCharacterRepo.On("Save", charInfo).Return(fakeData, nil)
	response, err := service.Save(ctx, charInfo)
	assert.Nil(ts.T(), err)

	assert.Equal(ts.T(), "Daredevil", response.Name, "Must be of the same name")
	assert.Equal(ts.T(), "0", response.Id, "Must not have the same id")
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(CharacterServiceTestSuite))
}
