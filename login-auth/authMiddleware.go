package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		// Json文字列がinterdace型で格納されている。dproxyのライブラリを使用して値を取り出す
		loginUserJson, err := dproxy.New(session.Get("loginUser")).String()

		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		} else {
			var loginInfo model.AuthUser
			// Json文字列のアンマーシャル
			err := json.Unmarshal([]byte(loginUserJson), &loginInfo)
			if err != nil {
				c.Status(http.StatusUnauthorized)
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}
