package rpc

import (
	"context"
	"errors"
	"log"
	"rating/constant"
	"rating/idl/gen/movie"
)

func ModifyMovieRating(ctx context.Context, movieID string, nModify float64, isNewUser bool) (
	*movie.ModifyMovieRatingResp, error) {
	req := &movie.ModifyMovieRatingReq{
		MovieId:   movieID,
		NUpdate:   nModify,
		IsNewUser: isNewUser,
	}
	log.Printf("req from rating to movie, req:%+v", req)
	resp, err := MovieClient.ModifyMovieRating(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp == nil || resp.BaseResp.ErrNo != constant.RetSuccess.Code {
		return nil, errors.New("resp ErrNo not success")
	}

	return resp, nil
}
