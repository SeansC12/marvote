package service

import (
	"context"

	"github.com/SeansC12/marvote/pkg/model"
	"github.com/stretchr/testify/mock"
)

type MockedCharacterRepository struct {
	mock.Mock
}

func (m *MockedCharacterRepository) FindAll(ctx context.Context) ([]model.CharacterInfo, error) {
	args := m.Called()
	return args.Get(0).([]model.CharacterInfo), args.Error(1)
}

func (m *MockedCharacterRepository) FindById(ctx context.Context, id string) (model.CharacterInfo, error) {
	args := m.Called(id)
	return args.Get(0).(model.CharacterInfo), args.Error(1)
}

func (m *MockedCharacterRepository) Save(ctx context.Context, character model.CharacterInfo) (model.CharacterInfo, error) {
	args := m.Called(character)
	return args.Get(0).(model.CharacterInfo), args.Error(1)
}

func (m *MockedCharacterRepository) Delete(ctx context.Context, id string) (int64, error) {
	args := m.Called(id)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockedCharacterRepository) CastVote(ctx context.Context, id string) error {
	args := m.Called(id)
	return args.Error(0)
}
