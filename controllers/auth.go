package controllers

import (
	"errors"
	"go_jwt/controllers/dao"
	"go_jwt/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReqRegister struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ReqLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginCheck(username, password string) (string, error) {
	var err error

	u := dao.User{}
	uData := u.CheckHaveUserName(username)

	if (uData.Id == 0) || (password != uData.Password) {
		err = errors.New("no this user or password error!")
		return "", err
	}

	token, err := tools.GenerateToken(uData.Id)
	if err != nil {
		return "", err
	}
	return token, err
}

func GetInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": "ok",
	})
}

func Login(c *gin.Context) {
	var req ReqLogin

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := LoginCheck(req.UserName,req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username or password error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Register(c *gin.Context) {

	var req ReqRegister

	u := new(dao.User)
	

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":err.Error(),
		})
		return
	}
	u.UserName=req.UserName
	u.Password=req.Password
	userId, err :=u.Add()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userId,
	})
}