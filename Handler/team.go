package Handler

import (
	"WeiYangWork/Model"
	"WeiYangWork/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetTeamInfo 获取队伍信息
func GetTeamInfo(c *gin.Context) {
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
	var teams []Model.Team
	result := global.Db.Order(sort).Offset(offset).Limit(pageSize).Find(&teams)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{"data": teams})
}

func CreateTeam(c *gin.Context) {
	claims, exist := c.Get("UserClaims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未识别到token!"})
		return
	}
	//获取用户信息
	userClaims := claims.(*Model.UserClaims)
	//获取请求参数
	var team Model.Team
	_ = c.BindJSON(&team)
	//设置队伍创建者
	team.Leader = userClaims.Username
	team.ApprovalStatus = "pending"
	//保存队伍信息
	result := global.Db.Save(&team)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "队伍创建成功！", "team": team})
}

func UpdateTeam(c *gin.Context) {
	claims, ok := c.Get("UserClaims")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未识别到token!"})
		return
	}
	//获取用户信息
	userClaims := claims.(*Model.UserClaims)
	var team, ExistTeam Model.Team
	_ = c.BindJSON(&team)
	//查询队伍信息
	if result := global.Db.Where("id=?", team.ID).First(&ExistTeam); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if userClaims.Username != ExistTeam.Name {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "权限不足！"})
	}
	//更新队伍信息
	result := global.Db.Model(&team).Updates(team)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "队伍信息更新成功！", "team": team})
}

func DeleteTeam(c *gin.Context) {
	claims, ok := c.Get("UserClaims")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未识别到token!"})
		return
	}
	userClaims := claims.(*Model.UserClaims)
	teamID := c.Param("teamID")
	var team Model.Team
	if result := global.Db.Where("id=?", teamID).First(&team); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if userClaims.Username != team.Leader {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "权限不足！"})
		return
	}
	if result := global.Db.Delete(&team); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "队伍删除成功！"})
}
