package service

import (
	"context"
	"errors"
	"fmt"
	. "rating/constant"
	"rating/idl/gen/rating"
	"rating/model"
	"rating/rpc"
	"strings"
	"sync"
)

type RatingService struct {
}

type RatingContext struct {
	Ctx     context.Context
	Req     *rating.RateReq
	Resp    *rating.RateResp
	ErrCode *ErrorCode
}

var ratingService *RatingService
var ratingServiceOnce sync.Once

func NewRatingService() *RatingService {
	ratingServiceOnce.Do(func() {
		ratingService = &RatingService{}
	})

	return ratingService
}

func NewRatingContext(ctx context.Context, req *rating.RateReq) *RatingContext {
	return &RatingContext{
		Ctx: ctx,
		Req: req,
		Resp: &rating.RateResp{
			BaseResp: &rating.BaseResp{},
		},
	}
}

func (s *RatingService) Rate(ctx *RatingContext) {
	defer s.buildResponse(ctx)
	if s.checkParams(ctx); ctx.ErrCode != nil {
		return
	}
	s.rate(ctx)
}

func (*RatingService) checkParams(ctx *RatingContext) {
	if ctx.Req.Rating <= 0 {
		ctx.ErrCode = BuildErrCode("rating must be above 0", RetParamsErr)
		return
	}
	if strings.TrimSpace(ctx.Req.UserId) == "" {
		ctx.ErrCode = BuildErrCode("empty user id", RetParamsErr)
		return
	}
	if strings.TrimSpace(ctx.Req.MovieId) == "" {
		ctx.ErrCode = BuildErrCode("empty movie id", RetParamsErr)
	}
}

func (*RatingService) rate(ctx *RatingContext) {
	isNewUser, nModify := true, ctx.Req.Rating
	if err := model.NewRatingDao().InsertRating(ctx.Ctx, ctx.Req.UserId, ctx.Req.MovieId, ctx.Req.Rating); err != nil {
		if !errors.Is(err, model.ErrDuplicateRating) {
			ctx.ErrCode = BuildErrCode(err, RetWriteRepoErr)
			return
		}

		isNewUser = false
		oldRating, err := model.NewRatingDao().QueryRating(ctx.Ctx, ctx.Req.UserId, []string{ctx.Req.MovieId})
		if len(oldRating) == 0 || err != nil {
			ctx.ErrCode = BuildErrCode(fmt.Sprintf("no rating, %s:%s:%v", ctx.Req.UserId, ctx.Req.MovieId, err),
				RetReadRepoErr)
			return
		}
		nModify = nModify - oldRating[0].Rating
		if err := model.NewRatingDao().UpdateRating(ctx.Ctx, ctx.Req.UserId, ctx.Req.MovieId, ctx.Req.Rating); err != nil {
			ctx.ErrCode = BuildErrCode(err, RetWriteRepoErr)
			return
		}
	}

	if _, err := rpc.ModifyMovieRating(ctx.Ctx, ctx.Req.MovieId, nModify, isNewUser); err != nil {
		ctx.ErrCode = BuildErrCode(err, RetRpcCallFailed)
	}
}

func (*RatingService) buildResponse(ctx *RatingContext) {
	errCode := RetSuccess
	if ctx.ErrCode != nil {
		errCode = ctx.ErrCode
	}

	ctx.Resp.BaseResp.ErrNo, ctx.Resp.BaseResp.ErrMsg = errCode.Code, errCode.Msg
}
