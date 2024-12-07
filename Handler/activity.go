package Handler

import (
	"WeiYangWork/Model"
	"WeiYangWork/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func JoinActivity(c *gin.Context) {
	claims, err := c.Get("UserClaims")
	if !err {
		c.JSON(401, gin.H{"error": "未识别到token!"})
		return
	}
	userClaims := claims.(*Model.UserClaims)
	teamID := c.Param("teamID")
	var team Model.Team
	if result := global.Db.Where("id=?", teamID).First(&team); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	//验证权限
	if team.Leader != userClaims.Username {
		c.JSON(403, gin.H{"error": "权限不足！"})
		return
	}
	//获取活动ID
	ActID := c.Param("ActID")
	var activity Model.Activity
	if result := global.Db.Where("id=?", ActID).First(&activity); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if activity.Status == 2 {
		c.JSON(403, gin.H{"error": "活动已经结束！"})
		return
	}
	activity.TeamID = append(activity.TeamID, team)
	result := global.Db.Save(&activity)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "加入活动成功！"})
}

func GetActivity(c *gin.Context) {
	_, ok := c.Get("UserClaims")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未识别到token!"})
		return
	}
	ActID := c.Query("ActID")
	var activity Model.Activity
	if result := global.Db.Where("id=?", ActID).First(&activity); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"activity": activity})
}

func GetAllActivity(c *gin.Context) {
	// 获取查询参数
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pagesize", "10")
	sort := c.DefaultQuery("sort", "")

	// 将查询参数转换为整数
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// 计算偏移量
	offset := (page - 1) * pageSize
	// 使用分页和排序查询数据库
	var activity []Model.Activity
	result := global.Db.Order(sort).Offset(offset).Limit(pageSize).Find(&activity)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	//返回结果
	c.JSON(http.StatusOK, gin.H{"data": activity})
}
