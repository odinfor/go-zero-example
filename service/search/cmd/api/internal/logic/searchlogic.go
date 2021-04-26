package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"go-zero-example/service/user/cmd/rpc/userclient"

	"go-zero-example/service/search/cmd/api/internal/svc"
	"go-zero-example/service/search/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchLogic {
	return SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req types.SearchRequest) (*types.SearchReply, error) {
	// todo: add your logic here and delete this line
	logx.Infof("userId: %v", l.ctx.Value("userId"))

	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	userId, err := userIdNumber.Int64()
	if err != nil {
		return nil, err
	}

	// 调用user rpc
	_, err = l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.SearchReply{
		Name: req.Name,
		Count: 100,
	}, nil
}
