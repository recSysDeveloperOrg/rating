package model

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRatingMetaDao_IncUserRating(t *testing.T) {
	assert.Nil(t, NewUserRatingMetaDao().IncUserRating(context.Background(), testUser, 1))
}
