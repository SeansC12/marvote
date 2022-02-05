package repository

import (
	"context"
	"log"

	"github.com/SeansC12/marvote/pkg/logging"
	"github.com/SeansC12/marvote/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CharacterRepository struct {
	characterCollections *mongo.Collection
}

func NewCharacterRepository(ctx context.Context, collection *mongo.Collection) *CharacterRepository {

	return &CharacterRepository{
		characterCollections: collection,
	}
}

func (cs *CharacterRepository) FindAll(ctx context.Context) ([]model.CharacterInfo, error) {
	var characters []model.CharacterInfo
	cursor, err := cs.characterCollections.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &characters); err != nil {
		log.Fatal(err)
	}
	return characters, nil
}

func (cs *CharacterRepository) FindById(ctx context.Context, characterId string) (model.CharacterInfo, error) {
	objectId, err := primitive.ObjectIDFromHex(characterId)
	if err != nil {
		return model.CharacterInfo{}, err
	}

	filter := bson.M{"_id": objectId}
	var characterInfo model.CharacterInfo

	if err := cs.characterCollections.FindOne(ctx, filter).Decode(&characterInfo); err != nil {
		return model.CharacterInfo{}, err
	}
	return characterInfo, nil
}

func (cs *CharacterRepository) Save(ctx context.Context, data model.CharacterInfo) (model.CharacterInfo, error) {
	result, err := cs.characterCollections.InsertOne(ctx, data)
	if err != nil {
		return model.CharacterInfo{}, err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return model.CharacterInfo{
			Name: data.Name,
			Aka:  data.Aka,
			Id:   oid.Hex(),
		}, nil

	}
	return model.CharacterInfo{}, nil
}

func (cs *CharacterRepository) Delete(ctx context.Context, characterId string) (int64, error) {

	objectId, err := primitive.ObjectIDFromHex(characterId)
	if err != nil {
		return 0, err
	}

	filter := bson.M{"_id": objectId}
	result, err := cs.characterCollections.DeleteOne(ctx, filter)
	if err != nil {
		logging.Errorf("%v", err)
		return 0, err
	}
	return result.DeletedCount, nil
}

func (cs *CharacterRepository) CastVote(ctx context.Context, characterId string) error {
	objectId, err := primitive.ObjectIDFromHex(characterId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectId}
	_, err = cs.characterCollections.UpdateOne(ctx, filter, bson.D{
		{"$inc", bson.D{{"votes", 1}}},
	}, options.Update().SetUpsert(true))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
