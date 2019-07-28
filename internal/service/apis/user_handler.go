package apis

import (
	"github.com/gin-gonic/gin"
	"smartapi/internal/log"
	"smartapi/internal/service/models"
)

// get user
func GetUser(c *gin.Context) {
	var h Header
	uuid, ok := c.Get("uuid")
	if !ok {
		log.Error("uuid get failed, please check middleware")
		ErrorMsgWithExtraInfo(c, UuidNotFound, h)
		return
	}
	h.Uuid = uuid.(string)

	// 1.解析参数
	id := c.Query("id") //查询请求URL后面的参数
	var user models.User
	err := models.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Errorf("err: %s,uuid: %s", err, uuid)
		ErrorMsgWithExtraInfo(c, UserNotFound, h)
		return
	}

	h.Code = Success
	h.Msg = getErrCodeMsg(Success)
	c.JSON(200, GetUserResp{
		Header: h,
		Data:   user,
	})
}

// add user
func AddUser(c *gin.Context) {
	var h Header
	uuid, ok := c.Get("uuid")
	if !ok {
		log.Error("uuid get failed, please check middleware")
		ErrorMsgWithExtraInfo(c, UuidNotFound, h)
		return
	}
	h.Uuid = uuid.(string)

	// 1.解析参数
	p := struct {
		UserName string `form:"user_name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	err := c.BindJSON(&p)
	if err != nil {
		log.Errorf("err: %s,uuid: %s", err, uuid)
		ErrorMsgWithExtraInfo(c, ParamInvalid, h)
		return
	}

	// 2.新增
	user := models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	err = models.DB.Create(&user).Error
	if err != nil {
		log.Errorf("err: %s,uuid: %s", err, uuid)
		ErrorMsgWithExtraInfo(c, UuidNotFound, h)
		return
	}

	h.Code = Success
	h.Msg = ErrCode[Success]
	c.JSON(200, h)
}

//update user
func UpdateUser(c *gin.Context) {

}

// delete user
func DeleteUser(c *gin.Context) {

}

// get user list
func GetUserList(c *gin.Context) {

}
