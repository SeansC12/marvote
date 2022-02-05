package repository

import (
	"context"
	"testing"

	"github.com/SeansC12/marvote/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

type CharacterRepositoryTestSuite struct {
	suite.Suite
}

func (ts *CharacterRepositoryTestSuite) SetupTest() {
}

func (ts *CharacterRepositoryTestSuite) TestMustSuccessfullySaveCharacter() {
	ctx := context.TODO()
	var characterCollection *mongo.Collection
	mt := mtest.New(ts.T(), mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(m *mtest.T) {
		characterCollection = m.Coll
		m.AddMockResponses(mtest.CreateSuccessResponse())

		repo := NewCharacterRepository(ctx, characterCollection)
		data := model.CharacterInfo{
			Aka:  "Natasha Romanov",
			Name: "Black Widow",
		}
		response, err := repo.Save(ctx, data)
		assert.Nil(ts.T(), err)
		assert.Equal(ts.T(), "Black Widow", response.Name, "Must have the same name")
	})
}
func (ts *CharacterRepositoryTestSuite) TestFindAllCharactersSuccess() {
	ctx := context.TODO()
	var characterCollection *mongo.Collection
	mt := mtest.New(ts.T(), mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(m *mtest.T) {
		characterCollection = m.Coll

		id1 := primitive.NewObjectID()
		id2 := primitive.NewObjectID()

		first := mtest.CreateCursorResponse(1, "marvel.marvel_characters", mtest.FirstBatch, bson.D{
			primitive.E{Key: "_id", Value: id1},
			primitive.E{Key: "name", Value: "Spiderman"},
			primitive.E{Key: "aka", Value: "Peter Parker"},
		})
		second := mtest.CreateCursorResponse(1, "marvel.marvel_characters", mtest.NextBatch, bson.D{
			primitive.E{Key: "_id", Value: id2},
			primitive.E{Key: "name", Value: "Wolverine"},
			primitive.E{Key: "aka", Value: "Logan"},
		})
		endCursor := mtest.CreateCursorResponse(0, "marvel.marvel_characters", mtest.NextBatch)
		m.AddMockResponses(first, second, endCursor)

		repo := NewCharacterRepository(ctx, characterCollection)
		response, err := repo.FindAll(ctx)
		assert.Nil(ts.T(), err)

		assert.Equal(ts.T(), 2, len(response), "Must be of size 2")
		assert.Equal(ts.T(), "Spiderman", response[0].Name, "Must have the same name")
		assert.Equal(ts.T(), "Wolverine", response[1].Name, "Must have the same name")
	})
}

func (ts *CharacterRepositoryTestSuite) TestGetOneCharacterSuccess() {
	var characterCollection *mongo.Collection
	ctx := context.TODO()
	mt := mtest.New(ts.T(), mtest.NewOptions().ClientType(mtest.Mock))
	id1 := primitive.NewObjectID()
	defer mt.Close()
	mt.Run("success", func(m *mtest.T) {
		characterCollection = m.Coll
		first := mtest.CreateCursorResponse(1, "marvel.marvel_characters", mtest.FirstBatch, bson.D{
			primitive.E{Key: "_id", Value: id1},
			primitive.E{Key: "name", Value: "Spiderman"},
			primitive.E{Key: "aka", Value: "Peter Parker"},
		})
		m.AddMockResponses(first)

		repo := NewCharacterRepository(ctx, characterCollection)
		response, err := repo.FindById(ctx, id1.Hex())
		assert.Nil(ts.T(), err)
		assert.Equal(ts.T(), "Spiderman", response.Name, "Must have the same name")
	})
}

func (ts *CharacterRepositoryTestSuite) TestInvalidIdWhileGetOneCharacter() {
	var characterCollection *mongo.Collection
	ctx := context.TODO()
	mt := mtest.New(ts.T(), mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(m *mtest.T) {
		repo := NewCharacterRepository(ctx, characterCollection)
		_, err := repo.FindById(ctx, "invalidId")
		assert.NotNil(ts.T(), err)
	})
}

func (ts *CharacterRepositoryTestSuite) TestSuccessfulDeleteOneCharacter() {
	var characterCollection *mongo.Collection
	ctx := context.TODO()
	mt := mtest.New(ts.T(), mtest.NewOptions().ClientType(mtest.Mock))
	id1 := primitive.NewObjectID()
	defer mt.Close()
	mt.Run("success", func(m *mtest.T) {
		characterCollection = m.Coll
		m.AddMockResponses(bson.D{
			primitive.E{Key: "ok", Value: 1},
			primitive.E{Key: "acknowledged", Value: true},
			primitive.E{Key: "n", Value: 1},
		})

		repo := NewCharacterRepository(ctx, characterCollection)
		response, err := repo.Delete(ctx, id1.Hex())
		assert.Nil(ts.T(), err)
		assert.Equal(ts.T(), int64(1), response, "Must have the same name")
	})
}

func (ts *CharacterRepositoryTestSuite) TestInvalidIdWhileDeletingCharacter() {
	var characterCollection *mongo.Collection
	ctx := context.TODO()
	mt := mtest.New(ts.T(), mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(m *mtest.T) {
		repo := NewCharacterRepository(ctx, characterCollection)
		_, err := repo.Delete(ctx, "invalidId")
		assert.NotNil(ts.T(), err)
	})
}

func (ts *CharacterRepositoryTestSuite) TestSuccessfulCastVote() {
	var characterCollection *mongo.Collection
	ctx := context.TODO()
	mt := mtest.New(ts.T(), mtest.NewOptions().ClientType(mtest.Mock))
	id1 := primitive.NewObjectID()
	defer mt.Close()
	mt.Run("success", func(m *mtest.T) {
		characterCollection = m.Coll
		m.AddMockResponses(bson.D{
			primitive.E{Key: "ok", Value: 1},
			primitive.E{Key: "acknowledged", Value: true},
			primitive.E{Key: "n", Value: 1},
		})

		repo := NewCharacterRepository(ctx, characterCollection)
		err := repo.CastVote(ctx, id1.Hex())
		assert.Nil(ts.T(), err)
	})
}

func (ts *CharacterRepositoryTestSuite) TestInvalidIdWhileCastingVote() {
	var characterCollection *mongo.Collection
	ctx := context.TODO()
	mt := mtest.New(ts.T(), mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(m *mtest.T) {
		repo := NewCharacterRepository(ctx, characterCollection)
		err := repo.CastVote(ctx, "invalidId")
		assert.NotNil(ts.T(), err)
	})
}
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(CharacterRepositoryTestSuite))
}
