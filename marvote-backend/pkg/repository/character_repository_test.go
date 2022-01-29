package repository

import (
	"testing"

	"github.com/SeansC12/marvote/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CharacterRepositoryTestSuite struct {
	suite.Suite
	fakeData model.CharacterInfo
}

func (ts *CharacterRepositoryTestSuite) SetupTest() {
	ts.fakeData = model.CharacterInfo{
		Id:   0,
		Name: "Spiderman",
		Aka:  "Peter Parker",
	}
}

func (ts *CharacterRepositoryTestSuite) TestFindAllCharactersSuccess() {

	repo := NewCharacterRepository()
	response, err := repo.FindAll()
	assert.Nil(ts.T(), err)

	assert.Equal(ts.T(), 2, len(response), "Must be of size 1")
	assert.Equal(ts.T(), "Spiderman", response[0].Name, "Must have the same name")
}

func (ts *CharacterRepositoryTestSuite) TestGetOneCharactersSuccess() {
	repo := NewCharacterRepository()
	response, err := repo.FindById(0)
	assert.Nil(ts.T(), err)
	assert.Equal(ts.T(), "Spiderman", response.Name, "Must be of the same name")
	assert.Equal(ts.T(), 0, response.Id, "Must have the same id")
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(CharacterRepositoryTestSuite))
}
