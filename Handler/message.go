package Handler

import (
	"WeiYangWork/Model"
	"WeiYangWork/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMessage(c *gin.Context) {
	claims, exist := c.Get("UserClaims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未识别到token!"})
		return
	}
	//获取用户信息
	userClaims := claims.(*Model.UserClaims)
	//获取请求参数
	var message []Model.Message
	Towards := c.Param("userID")
	//查询消息
	result := global.Db.Where("ReceiverID = ? AND SenderID = ?", Towards, userClaims.UserId).Find(&message)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func SendMessage(c *gin.Context) {
	claims, exist := c.Get("UserClaims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未识别到token!"})
		return
	}
	userClaims := claims.(Model.UserClaims)
	//
}
