package controller

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"service"
	"util"
)

func IndexAction(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func Authorizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := util.GetSession(c)
		if session.UID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		log.Debugln("Authorizer role:", session.Role,
			" path:", c.Request.URL.Path,
			" method:", c.Request.Method)

		ok := service.CheckPolicy(session.Role, c.Request.URL.Path, c.Request.Method)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg":   "没有权限",
				"error": 0,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func MapRouters() *gin.Engine {
	router := gin.New()
	store := cookie.NewStore([]byte("ngnetboy"))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("casbin", store))

	router.GET("/", IndexAction)
	router.POST("/api/v1/aaa/login", LoginAction)
	router.POST("/api/v1/aaa/logout", LogoutAction)

	admin := router.Group("/api/v1/admin")
	admin.Use(Authorizer())
	{
		//管理 policy
		admin.POST("/policy", AddPolicyAction)
		admin.GET("/policy", GetPolicyAction)
		admin.DELETE("/policy", DeletePolicyByRoleAction)
	}
	return router
}
