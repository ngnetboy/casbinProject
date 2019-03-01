package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
	"service"
	"util"
)

func LoginAction(c *gin.Context) {
	var account model.Account
	var user *model.Account
	var inputHansh string
	var session *util.SessionData

	if err := c.BindJSON(&account); err != nil {
		goto ERROR
	}

	user = service.Account.GetAccountByName(account.Name)
	if user == nil {
		goto ERROR
	}

	inputHansh, _ = util.GetSha512([]byte(account.Password), []byte(user.Password))
	if inputHansh != user.Password {
		goto ERROR
	}

	session = &util.SessionData{
		UID:   user.ID,
		UName: user.Name,
		Role:  user.Role,
	}
	session.Save(c)

	c.String(http.StatusOK, "login success")
	return

ERROR:
	c.String(http.StatusOK, "login failed")
}

func LogoutAction(c *gin.Context) {
	session := sessions.Default(c)
	session.Options(sessions.Options{
		MaxAge: -1,
	})
	session.Clear()
	session.Save()
	c.String(http.StatusOK, "logout success.")
}
