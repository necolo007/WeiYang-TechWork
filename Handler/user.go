package Handler

import (
	M "WeiYangWork/Model"
	"WeiYangWork/global"
	"WeiYangWork/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c *gin.Context) {
	var user M.User
_:
	c.BindJSON(&user)
	var ExistingUser M.User
	res := global.Db.Where("name=?", user.Name).First(&ExistingUser)
	if user.Name == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码不为空！"})
	} else if res.RowsAffected == 0 {
		c.JSON(404, gin.H{"msg": "登录失败，用户名不存在！"})
	} else {
		//对比加密后的密码，识别进入
		err := bcrypt.CompareHashAndPassword([]byte(ExistingUser.Password), []byte(user.Password))
		if err != nil {
			c.JSON(400, gin.H{"msg": "登录失败，用户名或者密码错误!"})
			return
		} else {
			var TokenString string
			TokenString, err = utils.GenerateToken(ExistingUser)
			c.JSON(200, gin.H{"msg": "登录成功！", "TokenString": TokenString})
		}
	}
}

// Register 注册
func Register(c *gin.Context) {
	var user M.User
_:
	c.BindJSON(&user)
	res := global.Db.Where("name=?", user.Name).First(&user)
	if user.Name == "" || user.Password == "" {
		c.JSON(400, gin.H{"error": "用户名和密码不能为空！"})
	} else if res.RowsAffected != 0 {
		c.JSON(400, gin.H{"msg": "注册失败，用户名字已经存在!"})
	} else {
		//对密码进行加密储存
		Hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"msg": "密码加密错误"})
		}
		user.Password = string(Hashedpassword)
		user.Role = "user"
		user.AdminLevel = 0
		global.Db.Save(&user)
		c.JSON(http.StatusOK, gin.H{"msg": "注册成功！", "UserInfo": user})
	}
}
