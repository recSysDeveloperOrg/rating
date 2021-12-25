package model

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRatingDao_InsertRating(t *testing.T) {
	dao := NewRatingDao()
	movieID := generateStr()
	assert.Equal(t, nil, dao.InsertRating(context.Background(), testUser, movieID, 5.0))
	assert.Equal(t, ErrDuplicateRating, dao.InsertRating(context.Background(), testUser, movieID, 5.0))
}

func TestRatingDao_QueryRating(t *testing.T) {
	ratings, err := NewRatingDao().QueryRating(context.Background(), testUser, []string{"917791850604129841576568"})
	assert.Nil(t, err)
	for _, r := range ratings {
		t.Logf("%+v", r)
	}
}

func TestRatingDao_QueryRecentRatings(t *testing.T) {
	ratings, nRecords, err := NewRatingDao().QueryRecentRatings(context.Background(), testUser, 0, 999999999)
	assert.Nil(t, err)
	assert.Equal(t, int64(len(ratings)), nRecords)
}

func TestRatingDao_UpdateRating(t *testing.T) {
	score := generateScore()
	assert.Nil(t, NewRatingDao().UpdateRating(context.Background(), testUser, "917791850604129841576568", score))
	rating, err := NewRatingDao().QueryRating(context.Background(), testUser, []string{"917791850604129841576568"})
	assert.Nil(t, err)
	assert.Equal(t, rating[0].Rating, score)
}
