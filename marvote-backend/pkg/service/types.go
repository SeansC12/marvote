package service

import (
	"context"

	"github.com/SeansC12/marvote/pkg/model"
)

type ICharacterService interface {
	GetAll(ctx context.Context) ([]model.CharacterInfo, error)
	Get(ctx context.Context, characterId string) (model.CharacterInfo, error)
	Save(context.Context, model.CharacterInfo) (model.CharacterInfo, error)
	Delete(ctx context.Context, characterId string) (int64, error)
}
