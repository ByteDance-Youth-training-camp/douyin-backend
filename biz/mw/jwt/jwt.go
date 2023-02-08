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

func SignUser(username string, expired time.Duration) (*string, error) {
	// Create a JWT token with claims and signing method
	claim := jwt.StandardClaims{
		Audience:  username,
		ExpiresAt: time.Now().Unix() + int64(expired.Seconds()),
		IssuedAt:  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(Secret)
	if err != nil {
		hlog.DefaultLogger().Warn("sign jwt token failed", err)
		return nil, err
	}
	return &signedToken, nil
}

func ExtractClaims(signedToken string) (jwt.MapClaims, error) {
	decodedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := decodedToken.Claims.(jwt.MapClaims)
	if !ok || !decodedToken.Valid {
		return nil, errors.New("Decode jwt token faild or invalid token")
	}
	return claims, nil
}

type token struct {
	Token string `query:"token" json:"token" form:"token"`
}

func Auth(ctx context.Context, c *app.RequestContext) {
	abort := func(err error) {
		hlog.DefaultLogger().Debug(err)
		c.JSON(consts.StatusUnauthorized, map[string]interface{}{
			"status_code": consts.StatusUnauthorized,
			"status_msg":  "authorization error",
		})
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
	username := claims["aud"]
	c.Set("username", username)
	c.Next(ctx)
}
