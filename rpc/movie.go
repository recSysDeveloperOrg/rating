package rpc

import (
	"context"
	"log"
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

	return resp, nil
}
