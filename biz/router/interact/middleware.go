// Code generated by hertz generator.

package Interact

import (
	"douyin_backend/biz/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{jwt.Auth}
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _comment_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoriteMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{jwt.Auth}
}

func _action0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favorite_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoritelistMw() []app.HandlerFunc {
	// your code...
	return nil
}
