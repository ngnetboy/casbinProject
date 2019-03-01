package controller

import (
	"github.com/Sirupsen/logrus"
	"model"
	"service"

	"github.com/gin-gonic/gin"
	"net/http"
)

func AddPolicyAction(c *gin.Context) {
	var pr model.PolicyRequest
	var ok bool

	if err := c.BindJSON(&pr); err != nil {
		goto ERROR
	}
	if len(pr.Role) == 0 || len(pr.Path) == 0 || len(pr.Method) == 0 {
		goto ERROR
	}

	ok = service.Policy.AddPolicy(&pr)
	if ok {
		c.String(http.StatusOK, "Add policy success.")
		return
	}
ERROR:
	c.String(http.StatusOK, "Add policy failed.")
}

func GetPolicyAction(c *gin.Context) {
	policies := service.Policy.GetPolicy()
	actions := service.Policy.GetAction()
	subjects := service.Policy.GetSubject()
	objects := service.Policy.GetObject()

	c.JSON(http.StatusOK, gin.H{
		"policies": policies,
		"actions":  actions,
		"subjects": subjects,
		"objects":  objects,
	})
	logrus.Infoln("getpolicy")
}

func DeletePolicyByRoleAction(c *gin.Context) {
	type content struct {
		Role string `json:"role"`
	}
	var ct content
	var ok bool
	if err := c.BindJSON(&ct); err != nil {
		goto ERROR
	}

	ok = service.Policy.DeletePolicy(0, ct.Role)
	if ok {
		c.String(http.StatusOK, "Delete policy success.")
		return
	}

ERROR:
	c.String(http.StatusOK, "Delete policy failed.")
}
