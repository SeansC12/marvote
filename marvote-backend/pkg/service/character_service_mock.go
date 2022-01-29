package service

import (
	"github.com/SeansC12/marvote/pkg/model"
	"github.com/stretchr/testify/mock"
)

type MockedCharacterRepository struct {
	mock.Mock
}

func (m *MockedCharacterRepository) FindAll() ([]model.CharacterInfo, error) {
	args := m.Called()
	return args.Get(0).([]model.CharacterInfo), args.Error(1)
}

func (m *MockedCharacterRepository) FindById(id int) (model.CharacterInfo, error) {
	args := m.Called(id)
	return args.Get(0).(model.CharacterInfo), args.Error(1)
}
