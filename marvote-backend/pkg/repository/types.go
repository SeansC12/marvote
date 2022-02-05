package repository

import (
	"context"

	"github.com/SeansC12/marvote/pkg/model"
)

type ICharacterRepository interface {
	FindAll(ctx context.Context) ([]model.CharacterInfo, error)
	FindById(ctx context.Context, characterId string) (model.CharacterInfo, error)
	Save(context.Context, model.CharacterInfo) (model.CharacterInfo, error)
	Delete(ctx context.Context, characterId string) (int64, error)
}
