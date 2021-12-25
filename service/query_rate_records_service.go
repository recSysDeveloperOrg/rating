package service

import (
	"context"
	. "rating/constant"
	"rating/idl/gen/rating"
	"rating/model"
	"strings"
	"sync"
)

type QueryRateRecordsService struct {
}

type QueryRateRecordsContext struct {
	Ctx     context.Context
	Req     *rating.QueryRateRecordsReq
	Resp    *rating.QueryRateRecordsResp
	ErrCode *ErrorCode

	Records  []*model.Rating
	NRecords int64
}

var queryRateRecordsService *QueryRateRecordsService
var queryRateRecordsServiceOnce sync.Once

func NewQueryRateRecordsService() *QueryRateRecordsService {
	queryRateRecordsServiceOnce.Do(func() {
		queryRateRecordsService = &QueryRateRecordsService{}
	})

	return queryRateRecordsService
}

func NewQueryRateRecordsContext(ctx context.Context, req *rating.QueryRateRecordsReq) *QueryRateRecordsContext {
	return &QueryRateRecordsContext{
		Ctx: ctx,
		Req: req,
		Resp: &rating.QueryRateRecordsResp{
			BaseResp: &rating.BaseResp{},
		},
	}
}

func (s *QueryRateRecordsService) Query(ctx *QueryRateRecordsContext) {
	defer s.buildResponse(ctx)
	if s.checkParams(ctx); ctx.ErrCode != nil {
		return
	}
	s.query(ctx)
}

func (*QueryRateRecordsService) checkParams(ctx *QueryRateRecordsContext) {
	if strings.TrimSpace(ctx.Req.UserId) == "" {
		ctx.ErrCode = BuildErrCode("empty user id", RetParamsErr)
		return
	}
	if ctx.Req.Page < 0 {
		ctx.ErrCode = BuildErrCode("page must > 0", RetParamsErr)
		return
	}
	if ctx.Req.Offset < 0 {
		ctx.ErrCode = BuildErrCode("offset must > 0", RetParamsErr)
	}
}

func (*QueryRateRecordsService) query(ctx *QueryRateRecordsContext) {
	records, nRecords, err := model.NewRatingDao().QueryRecentRatings(ctx.Ctx, ctx.Req.UserId, ctx.Req.Page, ctx.Req.Offset)
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetReadRepoErr)
		return
	}
	ctx.Records = records
	ctx.NRecords = nRecords
}

func (*QueryRateRecordsService) buildResponse(ctx *QueryRateRecordsContext) {
	errCode := RetSuccess
	if ctx.ErrCode != nil {
		errCode = ctx.ErrCode
	}

	ctx.Resp.BaseResp.ErrNo, ctx.Resp.BaseResp.ErrMsg = errCode.Code, errCode.Msg
	for _, record := range ctx.Records {
		ctx.Resp.Records = append(ctx.Resp.Records, &rating.RateRecord{
			Movie: &rating.Movie{
				Id: record.MovieID,
			},
			Rating: record.Rating,
		})
	}
	ctx.Resp.NRecords = ctx.NRecords
}
