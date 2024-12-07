package Handler

import (
	M "WeiYangWork/Model"
	"WeiYangWork/global"
	"github.com/gin-gonic/gin"
)

func CreateActivity(c *gin.Context) {
	claims, err := c.Get("UserClaims")
	if !err {
		c.JSON(401, gin.H{"error": "未识别到token!"})
		return
	}
	userClaims := claims.(*M.UserClaims)
	if userClaims.Role != "admin" {
		c.JSON(403, gin.H{"error": "权限不足"})
		return
	}
	var activity M.Activity
	_ = c.BindJSON(&activity)
	activity.ActivityLeader = userClaims.Username
	result := global.Db.Save(&activity)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"activity": activity})
}
