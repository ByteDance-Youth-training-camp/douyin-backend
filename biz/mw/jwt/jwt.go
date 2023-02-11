package jwt

import (
	"context"
	"douyin_backend/biz/config"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v4"
)

/* About Standard JWT Claims
aud: 接收jwt的一方
exp: jwt的过期时间，这个过期时间必须要大于签发时间
jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
iat: jwt的签发时间
iss: jwt签发者
nbf: 定义在什么时间之前，该jwt都是不可用的.就是这条token信息生效时间.这个值可以不设置,但是设定后,一定要大于当前Unix UTC,否则token将会延迟生效.
sub: jwt所面向的用户
*/

func SignUser(userId int64, username string, expired time.Duration) (*string, error) {
	// Create a JWT token with claims and signing method
	claim := jwt.RegisteredClaims{
		ID:        strconv.FormatInt(userId, 10),
		Issuer:    "minitok",
		Subject:   "token",
		Audience:  jwt.ClaimStrings{username},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expired)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(config.Cfg.Jwt.Secret))
	if err != nil {
		hlog.DefaultLogger().Warn("sign jwt token failed", err)
		return nil, err
	}
	return &signedToken, nil
}

type token struct {
	Token string `query:"token" json:"token" form:"token"`
}

func Auth(ctx context.Context, c *app.RequestContext) {
	abort := func(err error) {
		hlog.DefaultLogger().Debug(err)
		c.JSON(consts.StatusOK, map[string]interface{}{
			"status_code": consts.StatusOK,
			"status_msg":  "authorization error",
		})
	}

	tk := token{}
	err := c.Bind(&tk)
	if err != nil {
		abort(err)
		return
	}

	claims, err := ParseToken(tk.Token)
	hlog.DefaultLogger().Debug(tk.Token)
	if err != nil {
		abort(err)
		return
	}

	c.Set("username", claims.Audience)
	c.Set("userId", claims.ID)
	c.Next(ctx)
}

func ParseToken(token string) (*jwt.RegisteredClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.Jwt.Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.RegisteredClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
