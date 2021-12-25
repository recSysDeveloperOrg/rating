package main

import (
	"context"
	"log"
	"rating/idl/gen/rating"
	"rating/service"
)

type Handler struct {
	*rating.UnimplementedRatingServiceServer
}

func (*Handler) RateMovie(ctx context.Context, req *rating.RateReq) (*rating.RateResp, error) {
	log.Printf("req:%+v", req)
	rCtx := service.NewRatingContext(ctx, req)
	service.NewRatingService().Rate(rCtx)
	log.Printf("resp:%+v", rCtx.Resp)

	return rCtx.Resp, nil
}

func (*Handler) QueryRateRecords(ctx context.Context, req *rating.QueryRateRecordsReq) (*rating.QueryRateRecordsResp, error) {
	log.Printf("req:%+v", req)
	qCtx := service.NewQueryRateRecordsContext(ctx, req)
	service.NewQueryRateRecordsService().Query(qCtx)
	log.Printf("resp:%+v", qCtx.Resp)

	return qCtx.Resp, nil
}

func (*Handler) BatchQueryMovieRating(ctx context.Context, req *rating.QueryMovieRatingReq) (*rating.QueryMovieRatingResp, error) {
	log.Printf("req:%+v", req)
	qCtx := service.NewQueryRatingContext(ctx, req)
	service.NewQueryRatingService().Query(qCtx)
	log.Printf("resp:%+v", qCtx.Resp)

	return qCtx.Resp, nil
}
