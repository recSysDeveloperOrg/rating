package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"rating/constant"
	"rating/idl/gen/rating"
	"testing"
)

func TestQueryRateRecordsService_Query(t *testing.T) {
	ctx := NewQueryRateRecordsContext(context.Background(), &rating.QueryRateRecordsReq{
		UserId: testUser,
		Page:   0,
		Offset: 1000,
	})
	NewQueryRateRecordsService().Query(ctx)
	assert.Equal(t, ctx.Resp.BaseResp.ErrNo, constant.RetSuccess.Code)
	assert.NotZero(t, ctx.Resp.NRecords)
	assert.NotZero(t, len(ctx.Resp.Records))
}
