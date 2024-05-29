package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

var code int

//// if user exist
//func UserExist(c *gin.Context) {
//
//}

// add user
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

//search single user

// search user list
func GetUsers(c *gin.Context) {

}

// edit user
func EditUser(c *gin.Context) {

}

// delete user
func DeleteUser(c *gin.Context) {

}
