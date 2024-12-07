package Handler

import (
	"WeiYangWork/Model"
	"WeiYangWork/global"
	"WeiYangWork/service"
	"WeiYangWork/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// ConnectWebSocket 连接 WebSocket
func ConnectWebSocket(c *gin.Context) {
	claims, ok := c.Get("UserClaims")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未识别到token!"})
		return
	}
	teamID, err := strconv.Atoi(c.Param("teamID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}
	// 验证用户是否在队伍中
	if !utils.IsUserInTeam(claims.(*Model.UserClaims), uint(teamID)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not in this team"})
		return
	}
	conn, err := global.UP.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	service.WsHandler(conn, uint(teamID))
}
