package util

import (
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionData struct {
	UID   uint   `json:"id"`
	UName string `json:"name"`
	Role  string `json:"role"`
}

func (sd *SessionData) Save(c *gin.Context) error {
	session := sessions.Default(c)
	sessionDataBytes, err := json.Marshal(sd)
	if err != nil {
		return err
	}

	session.Set("data", string(sessionDataBytes))
	return session.Save()
}

func GetSession(c *gin.Context) *SessionData {
	ret := &SessionData{}

	session := sessions.Default(c)
	sessionDataStr := session.Get("data")
	if sessionDataStr == nil {
		return ret
	}

	err := json.Unmarshal([]byte(sessionDataStr.(string)), ret)
	if err != nil {
		return ret
	}

	return ret
}
