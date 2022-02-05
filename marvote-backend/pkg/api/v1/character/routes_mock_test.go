package character

import (
	"context"

	"github.com/SeansC12/marvote/pkg/model"
	"github.com/stretchr/testify/mock"
)

type MockedCharacterService struct {
	mock.Mock
}

func (m *MockedCharacterService) GetAll(ctx context.Context) ([]model.CharacterInfo, error) {
	args := m.Called()
	return args.Get(0).([]model.CharacterInfo), args.Error(1)
}

func (m *MockedCharacterService) Get(ctx context.Context, id string) (model.CharacterInfo, error) {
	args := m.Called(id)
	return args.Get(0).(model.CharacterInfo), args.Error(1)
}

func (m *MockedCharacterService) Save(ctx context.Context, charInfo model.CharacterInfo) (model.CharacterInfo, error) {
	args := m.Called(charInfo)
	return args.Get(0).(model.CharacterInfo), args.Error(1)
}

func (m *MockedCharacterService) Delete(ctx context.Context, id string) (int64, error) {
	args := m.Called(id)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockedCharacterService) CastVote(ctx context.Context, id string) error {
	args := m.Called(id)
	return args.Error(0)
}
