package model

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type RatingDao struct {
}

type Rating struct {
	MovieID string  `bson:"movie_id"`
	Rating  float64 `bson:"rating"`
	UserID  string  `bson:"user_id"`
}

var ratingDao *RatingDao
var ratingDaoOnce sync.Once

var ErrDuplicateRating = errors.New("you've already rated this movie")

const errDuplicateCode = 11000

func NewRatingDao() *RatingDao {
	ratingDaoOnce.Do(func() {
		ratingDao = &RatingDao{}
	})

	return ratingDao
}

func (*RatingDao) InsertRating(ctx context.Context, userID, movieID string, nRating float64) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	movieObjectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return err
	}
	if _, err := GetClient().Collection(CollectionRating).InsertOne(ctx, struct {
		UserID  primitive.ObjectID `bson:"user_id"`
		MovieID primitive.ObjectID `bson:"movie_id"`
		Rating  float64            `bson:"rating"`
	}{userObjectID, movieObjectID, nRating}); err != nil {
		var writeErr mongo.WriteException
		if errors.As(err, &writeErr) {
			if len(writeErr.WriteErrors) > 0 {
				if writeErr.WriteErrors[0].Code == errDuplicateCode {
					return ErrDuplicateRating
				}
			}
		}

		return err
	}

	return nil
}

func (*RatingDao) UpdateRating(ctx context.Context, userID, movieID string, nRating float64) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	movieObjectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return err
	}
	whereMap := bson.D{{"user_id", userObjectID}, {"movie_id", movieObjectID}}
	updatesMap := bson.D{{"$set", bson.D{{"rating", nRating}}}}
	if _, err := GetClient().Collection(CollectionRating).UpdateOne(ctx, whereMap, updatesMap); err != nil {
		return err
	}

	return nil
}

func (*RatingDao) QueryRating(ctx context.Context, userID string, movieIDs []string) ([]*Rating, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	movieObjectIDs := make([]primitive.ObjectID, len(movieIDs))
	for i, movieID := range movieIDs {
		movieObjectID, err := primitive.ObjectIDFromHex(movieID)
		if err != nil {
			return nil, err
		}
		movieObjectIDs[i] = movieObjectID
	}

	whereMap := bson.D{{"user_id", userObjectID},
		{"movie_id", bson.D{{"$in", movieObjectIDs}}}}
	c, err := GetClient().Collection(CollectionRating).Find(ctx, whereMap)
	if err != nil {
		return nil, err
	}
	var ratings []*Rating
	if err := c.All(ctx, &ratings); err != nil {
		return nil, err
	}

	return ratings, nil
}

func (*RatingDao) QueryRecentRatings(ctx context.Context, userID string, page, offset int64) ([]*Rating, int64, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, 0, err
	}
	c, err := GetClient().Collection(CollectionRating).Find(ctx, bson.D{{"user_id", userObjectID}},
		options.Find().SetSkip(page*offset).SetLimit(offset).SetSort(bson.D{{"_id", -1}}))
	if err != nil {
		return nil, 0, err
	}
	var ratings []*Rating
	if err := c.All(ctx, &ratings); err != nil {
		return nil, 0, err
	}

	nRecords, err := GetClient().Collection(CollectionRating).CountDocuments(ctx, nil)
	if err != nil {
		return nil, 0, err
	}

	return ratings, nRecords, nil
}
