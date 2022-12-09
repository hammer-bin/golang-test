package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c *gin.Context) {
	var request authModel.EmailLoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		// メールアドレスでDBからユーザ取得（詳細は割愛）
		user, err := authRepository.GetUserByEmail(request.email)
		// ハッシュ値でのパスワード比較
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.password))
		if err != nil {
			c.Status(http.StatusBadRequest)
		} else {
			session := sessions.Default(c)
			// セッションに格納する為にユーザ情報をJson化
			loginUser, err := json.Marshal(auth)
			if err == nil {
				session.Set("loginUser", string(loginUser))
				session.Save()
				c.Status(http.StatusOK)
			} else {
				c.Status(http.StatusInternalServerError)
			}
		}
	}
}
