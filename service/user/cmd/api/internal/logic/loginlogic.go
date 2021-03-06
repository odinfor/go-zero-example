package logic

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	errorx "go-zero-example/common"
	"go-zero-example/service/user/model"
	"strings"
	"time"

	"go-zero-example/service/user/cmd/api/internal/svc"
	"go-zero-example/service/user/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(req types.LoginRequest) (*types.LoginReply, error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByName(req.Username)
	switch err {
	case nil:
		if userInfo.Password != req.Password {
			return nil, errorx.NewDefaultError("用户密码不正确")
		}

		// 鉴权
		now := time.Now().Unix()
		accessExpire := l.svcCtx.Config.Auth.AccessExpire
		jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
		if err != nil {
			logx.Errorf("jwt token error: %v", err)
			return &types.LoginReply{}, errorx.NewDefaultError(err.Error())
		}

		fmt.Printf("get jwt token: %s", jwtToken)

		return &types.LoginReply{
			Id:           userInfo.Id,
			Name:         userInfo.Name,
			Gender:       userInfo.Gender,
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		}, nil
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("用户名不存在")
	default:
		return nil, err
	}
}
