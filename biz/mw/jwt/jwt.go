package jwt

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v4"
)

var Secret = []byte("a secret key")

/* About Standard JWT Claims
aud: 接收jwt的一方
exp: jwt的过期时间，这个过期时间必须要大于签发时间
jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
iat: jwt的签发时间
iss: jwt签发者
nbf: 定义在什么时间之前，该jwt都是不可用的.就是这条token信息生效时间.这个值可以不设置,但是设定后,一定要大于当前Unix UTC,否则token将会延迟生效.
sub: jwt所面向的用户
*/

type uClaims struct {
	Uid int64 `json:"uid"`
	jwt.RegisteredClaims
}

func SignUser(uid int64, expired time.Duration) (*string, error) {
	// Create a JWT token with claims and signing method
	claim := uClaims{
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expired)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(Secret)
	if err != nil {
		hlog.Warn("sign jwt token failed", err)
		return nil, err
	}

	return &signedToken, nil
}

func ExtractClaims(signedToken string) (*uClaims, error) {
	
	decodedToken, err := jwt.ParseWithClaims(signedToken, &uClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !decodedToken.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := decodedToken.Claims.(*uClaims)
	if !ok {
		return nil, errors.New("jwt claims type error")
	}
	return claims, nil
}

type token struct {
	Token string `query:"token" json:"token" form:"token"`
}

func Auth(ctx context.Context, c *app.RequestContext) {
	abort := func(err error) {
		hlog.Debug(err)
		c.JSON(consts.StatusOK, map[string]interface{}{
			"status_code": consts.StatusUnauthorized,
			"status_msg":  "authorization error",
		})
		c.Abort()
	}

	tk := token{}
	err := c.Bind(&tk)
	if err != nil {
		abort(err)
		return
	}
	
	claims, err := ExtractClaims(tk.Token)
	if err != nil {
		abort(err)
		return
	}
	c.Set("uid", claims.Uid)
	c.Next(ctx)
}
