package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"rating/constant"
	"rating/idl/gen/rating"
	"testing"
)

func TestQueryRatingService_Query(t *testing.T) {
	ctx := NewQueryRatingContext(context.Background(), &rating.QueryMovieRatingReq{
		UserId:      testUser,
		MovieIdList: []string{"917791850604129841576568"},
	})
	NewQueryRatingService().Query(ctx)
	assert.Equal(t, ctx.Resp.BaseResp.ErrNo, constant.RetSuccess.Code)
}
