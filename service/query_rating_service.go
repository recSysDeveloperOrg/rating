package service

import (
	"context"
	. "rating/constant"
	"rating/idl/gen/rating"
	"rating/model"
	"strings"
	"sync"
)

type QueryRatingService struct {
}

type QueryRatingContext struct {
	Ctx     context.Context
	Req     *rating.QueryMovieRatingReq
	Resp    *rating.QueryMovieRatingResp
	ErrCode *ErrorCode

	Ratings []*model.Rating
}

var queryRatingService *QueryRatingService
var queryRatingServiceOnce sync.Once

func NewQueryRatingService() *QueryRatingService {
	queryRatingServiceOnce.Do(func() {
		queryRatingService = &QueryRatingService{}
	})

	return queryRatingService
}

func NewQueryRatingContext(ctx context.Context, req *rating.QueryMovieRatingReq) *QueryRatingContext {
	return &QueryRatingContext{
		Ctx: ctx,
		Req: req,
		Resp: &rating.QueryMovieRatingResp{
			BaseResp: &rating.BaseResp{},
		},
	}
}

func (s *QueryRatingService) Query(ctx *QueryRatingContext) {
	defer s.buildResponse(ctx)
	if s.checkParams(ctx); ctx.ErrCode != nil {
		return
	}
	s.query(ctx)
}

func (*QueryRatingService) checkParams(ctx *QueryRatingContext) {
	if strings.TrimSpace(ctx.Req.UserId) == "" {
		ctx.ErrCode = BuildErrCode("empty user id", RetParamsErr)
		return
	}
	for _, movieID := range ctx.Req.MovieIdList {
		if strings.TrimSpace(movieID) == "" {
			ctx.ErrCode = BuildErrCode("empty movie id", RetParamsErr)
			return
		}
	}
}

func (*QueryRatingService) query(ctx *QueryRatingContext) {
	rates, err := model.NewRatingDao().QueryRating(ctx.Ctx, ctx.Req.UserId, ctx.Req.MovieIdList)
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetReadRepoErr)
		return
	}
	ctx.Ratings = rates
}

func (*QueryRatingService) buildResponse(ctx *QueryRatingContext) {
	errCode := RetSuccess
	if ctx.ErrCode != nil {
		errCode = ctx.ErrCode
	}

	ctx.Resp.BaseResp.ErrNo, ctx.Resp.BaseResp.ErrMsg = errCode.Code, errCode.Msg
	movieID2Rating := make(map[string]float64)
	for _, movieRating := range ctx.Ratings {
		movieID2Rating[movieRating.MovieID] = movieRating.Rating
	}
	ctx.Resp.MovieId2PersonalRating = movieID2Rating
}
