package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type UserRatingMetaDao struct {
}

var userRatingMetaDao *UserRatingMetaDao
var userRatingMetaDaoOnce sync.Once

func NewUserRatingMetaDao() *UserRatingMetaDao {
	userRatingMetaDaoOnce.Do(func() {
		userRatingMetaDao = &UserRatingMetaDao{}
	})

	return userRatingMetaDao
}

func (*UserRatingMetaDao) IncUserRating(ctx context.Context, userID string, inc int64) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	whereMap := bson.D{{"user_id", userObjectID}}
	updatesMap := bson.D{{"$inc", bson.D{{"total_rating", inc}}}}
	if _, err := GetClient().Collection(CollectionUserRatingMeta).
		UpdateOne(ctx, whereMap, updatesMap, options.Update().SetUpsert(true)); err != nil {
		return err
	}

	return nil
}
