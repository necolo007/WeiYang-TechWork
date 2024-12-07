package main

import (
	H "WeiYangWork/Handler"
	M "WeiYangWork/Middleware"
	C "WeiYangWork/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	C.InitConfig()
	r := gin.Default()
	v1 := r.Group("/user")
	{
		v1.POST("/login", H.Login)
		v1.POST("/register", H.Register)
		v1.GET("/activity/An", M.AuthMiddleware(), H.GetActivity)
		v1.GET("/activity/All", M.AuthMiddleware(), H.GetAllActivity)
	}
	v2 := r.Group("/team")
	{
		v2.GET("/getTeamInfo", M.AuthMiddleware(), H.GetTeamInfo)
		v2.POST("/createTeam", M.AuthMiddleware(), H.CreateTeam)
		v2_1 := v2.Group("/:teamID")
		{
			//Leader操作
			v2_1.PUT("/updateTeam", M.AuthMiddleware(), H.UpdateTeam)
			v2_1.DELETE("/deleteTeam", M.AuthMiddleware(), H.DeleteTeam)
			v2_1.POST("joinActivity/:ActID", M.AuthMiddleware(), H.JoinActivity)
			//Member操作
			v2_1.GET("/chat", M.AuthMiddleware(), H.ConnectWebSocket)
			v2_1.POST("/joinTeam", M.AuthMiddleware(), H.JoinTeam)
		}
	}
	v3 := r.Group("/admin")
	{
		v3.POST("/createActivity", M.AuthMiddleware(), H.CreateActivity)
	}
	err := r.Run("localhost:8080")
	if err != nil {
		fmt.Println("服务器启动失败！")
	}
}
