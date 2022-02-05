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

func (ts *CharacterServiceTestSuite) TestMustGetAllCharacters() {
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

func (ts *CharacterServiceTestSuite) TestMustNotGetAllCharacters() {
	ctx := context.TODO()
	reqErr := &ErrorFailedToLoadData{}
	mockCharacterRepo := new(MockedCharacterRepository)
	service := NewCharacterService(mockCharacterRepo)
	allCharacters := make([]model.CharacterInfo, 0, 1)

	mockCharacterRepo.On("FindAll").Return(allCharacters, reqErr)
	response, err := service.GetAll(ctx)
	assert.NotNil(ts.T(), err)

	assert.Equal(ts.T(), 0, len(response), "Must be of size 1")
}

func (ts *CharacterServiceTestSuite) TestMustGetOneCharacter() {
	ctx := context.TODO()
	mockCharacterRepo := new(MockedCharacterRepository)
	service := NewCharacterService(mockCharacterRepo)
	mockCharacterRepo.On("FindById", "0").Return(ts.fakeData, nil)
	response, err := service.Get(ctx, "0")
	assert.Nil(ts.T(), err)

	assert.Equal(ts.T(), "Spiderman", response.Name, "Must be of the same name")
	assert.Equal(ts.T(), "0", response.Id, "Must have the same id")
}

func (ts *CharacterServiceTestSuite) TestMustNotGetOneCharacter() {
	ctx := context.TODO()
	reqErr := &ErrorFailedToLoadData{}
	mockCharacterRepo := new(MockedCharacterRepository)
	service := NewCharacterService(mockCharacterRepo)
	mockCharacterRepo.On("FindById", "0").Return(model.CharacterInfo{}, reqErr)
	response, err := service.Get(ctx, "0")
	assert.NotNil(ts.T(), err)

	assert.Equal(ts.T(), "", response.Name, "Name must be empty")
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

func (ts *CharacterServiceTestSuite) TestMustNotSaveCharacters() {
	ctx := context.TODO()
	reqErr := &ErrorFailedToLoadData{}
	mockCharacterRepo := new(MockedCharacterRepository)
	service := NewCharacterService(mockCharacterRepo)
	charInfo := model.CharacterInfo{
		Name: "Daredevil",
		Aka:  "Matt Murdock",
	}
	mockCharacterRepo.On("Save", charInfo).Return(model.CharacterInfo{}, reqErr)
	response, err := service.Save(ctx, charInfo)
	assert.Equal(ts.T(), "", response.Name, "Name must be empty")
	assert.NotNil(ts.T(), err)
}
func (ts *CharacterServiceTestSuite) TestMustDeleteOneCharacter() {
	ctx := context.TODO()
	mockCharacterRepo := new(MockedCharacterRepository)
	service := NewCharacterService(mockCharacterRepo)
	mockCharacterRepo.On("Delete", "0").Return(int64(1), nil)
	response, err := service.Delete(ctx, "0")
	assert.Nil(ts.T(), err)

	assert.Equal(ts.T(), int64(1), response, "Must delete one record")
}

func (ts *CharacterServiceTestSuite) TestMustNotDeleteOneCharacter() {
	ctx := context.TODO()
	reqErr := &ErrorFailedToLoadData{}
	mockCharacterRepo := new(MockedCharacterRepository)
	service := NewCharacterService(mockCharacterRepo)
	mockCharacterRepo.On("Delete", "0").Return(int64(0), reqErr)
	response, err := service.Delete(ctx, "0")
	assert.NotNil(ts.T(), err)

	assert.Equal(ts.T(), int64(0), response, "Must delete no record")
}

func (ts *CharacterServiceTestSuite) TestMustAllowToCastVote() {
	ctx := context.TODO()
	mockCharacterRepo := new(MockedCharacterRepository)
	service := NewCharacterService(mockCharacterRepo)
	mockCharacterRepo.On("CastVote", "0").Return(nil)
	err := service.CastVote(ctx, "0")
	assert.Nil(ts.T(), err)
}

func (ts *CharacterServiceTestSuite) TestMustNotAllowToCastVote() {
	ctx := context.TODO()
	reqErr := &ErrorFailedToLoadData{}
	mockCharacterRepo := new(MockedCharacterRepository)
	service := NewCharacterService(mockCharacterRepo)
	mockCharacterRepo.On("CastVote", "0").Return(reqErr)
	err := service.CastVote(ctx, "0")
	assert.NotNil(ts.T(), err)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(CharacterServiceTestSuite))
}
