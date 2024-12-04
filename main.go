package main

import (
	H "WeiYangWork/Handler"
	M "WeiYangWork/Middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/user")
	{
		v1.POST("/login", H.Login)
		v1.POST("/register", H.Register)
	}
	v2 := r.Group("/team")
	{
		v2.GET("/getTeamInfo", M.AuthMiddleware(), H.GetTeamInfo)
		v2.POST("/createTeam", M.AuthMiddleware(), H.CreateTeam)
		v2_1 := v2.Group("/:teamID")
		{
			v2_1.PUT("/updateTeam", M.AuthMiddleware(), H.UpdateTeam)
			v2_1.DELETE("/deleteTeam", M.AuthMiddleware(), H.DeleteTeam)
		}
	}
	v3 := r.Group("/Admin")
	{

	}
	v4 := r.Group("/message/:userID")
	{
		v4.GET("/getMessage", M.AuthMiddleware(), H.GetMessage)
		v4.POST("/sendMessage/:towards", M.AuthMiddleware(), H.SendMessage)
	}
	err := r.Run("localhost:8080")
	if err != nil {
		fmt.Println("服务器启动失败！")
	}
}
