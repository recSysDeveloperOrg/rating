package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"rating/idl/gen/movie"
	"rating/idl/gen/rating"
	"rating/mock"
	"rating/rpc"
	"testing"
)

func TestRatingService_Rate(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	movieMockClient := mock.NewMockMovieServiceClient(mockCtl)
	req := &rating.RateReq{
		UserId:  testUser,
		MovieId: "917791850604129841576568",
		Rating:  generateScore(),
	}
	movieMockClient.EXPECT().ModifyMovieRating(gomock.Any(), gomock.Any()).Return(&movie.ModifyMovieRatingResp{
		BaseResp: &movie.BaseResp{
			ErrNo:  0,
			ErrMsg: "成功",
		},
	}, nil)
	rpc.MovieClient = movieMockClient

	ctx := NewRatingContext(context.Background(), req)
	NewRatingService().Rate(ctx)
}
