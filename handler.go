package main

import (
	"context"
	"rating/idl/gen/rating"
)

type Handler struct {
	*rating.UnimplementedRatingServiceServer
}

func (*Handler) RateMovie(ctx context.Context, req *rating.RateReq) (*rating.RateResp, error) {
	return nil, nil
}

func (*Handler) QueryRateRecords(ctx context.Context, req *rating.QueryRateRecordsReq) (*rating.QueryRateRecordsResp, error) {
	return nil, nil
}

func (*Handler) BatchQueryMovieRating(ctx context.Context, req *rating.QueryMovieRatingReq) (*rating.QueryMovieRatingResp, error) {
	return nil, nil
}
