package main

import (
	"awesomeProject16/以前的/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r:= gin.Default()
	r.POST("/user/:login",Login)

	v1 := r.Group("/v1")
	{
		v1.POST("/register",Register )
		v1.POST("/login", Login)
		v1.GET("/ta",TA)
		v1.LoadHTMLFiles("D:/Goprojects/src/awesomeProject16/loading.html")
		v1.LoadHTMLFiles("D:/Goprojects/src/awesomeProject16/register.html")
		//v1.POST("/read", readEndpoint)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/search",Search)
		v2.POST("/sendmsg",SendMsg)
		v1.LoadHTMLFiles("D:/Goprojects/src/awesomeProject16/vedio.html")
		v2.GET("/showmsg",ShowMsg)
		v2.POST("/zan",Zan)
		v2.POST("/collect",Collect)
		v2.POST("/coin",Coin)
		v2.POST("/sanlian",Sanlian)

	}
	/*v3 := r.Group("/v3")
	{
		v3.POST("/upload",Upload)
	}

	}
*/

	r.Run(":8080")

}
type user struct {
	username string `json:"username"`
	password string `json:"password"`
}



//注册
func Register(c *gin.Context){
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	//phonenumber:=c.PostForm("phonenumber")
	//identify:=c.PostForm("identify")
	fmt.Println("user:"+username+password/*+phonenumber+identify*/)
	if model.UserSignup(username,password/*,phonenumber,identify*/){
		c.JSON(500,gin.H{"status":http.StatusInternalServerError,"message":"数据库Insert报错"})
	}else {
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "成功" })
	}
}
//登录
func Login(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	if model.UserSignin(username, password) {
		//c.SetCookie("username", username, 10, "localhost:8080", "localhost", false, true)
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "登录成功"})
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "登录失败，用户名或密码错误"})
	}
}

//TA的点赞关注数等
func TA(c *gin.Context) {


}

//搜索
func Search(c *gin.Context) {




}

//发送评论
func SendMsg(c *gin.Context){
	username:=c.PostForm("username")
	fmt.Println("username"+username)
	/*if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"cookie读取失败"})
		return
	}*/
	message:=c.PostForm("message")
	if model.SendMessage(username,message){
		c.JSON(200, gin.H{"内容":message,"用户名":username})
	}else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "发送失败"})
	}
}

//展示评论
func ShowMsg(c *gin.Context){
	err,message:=model.ShowMesage(10)
	if err !=nil{
		c.JSON(500,gin.H{"status": http.StatusInternalServerError,"message":"数据库读取失败"})
	}
	c.JSON(200,gin.H{"评论":message})
}

//点赞
func Zan(c*gin.Context){
	username:=c.PostForm("username")
	video:=c.PostForm("video")
	good:=c.PostForm("good")
	if model.Good(username, video,good) {
		//c.SetCookie("username", username, 10, "localhost:8080", "localhost", false, true)
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "点赞成功"})
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "点赞失败"})
	}
}

//关注
func Collect(c*gin.Context){
	username:=c.PostForm("username")
	video:=c.PostForm("video")
	collect:=c.PostForm("collect")
	if model.Shoucan(username, video,collect) {
		//c.SetCookie("username", username, 10, "localhost:8080", "localhost", false, true)
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "关注成功"})
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "关注失败"})
	}
}

//投币
func Coin(c*gin.Context){
	username:=c.PostForm("username")
	video:=c.PostForm("video")
	coin:=c.PostForm("coin")
	if model.Toubi(username, video,coin) {
		//c.SetCookie("username", username, 10, "localhost:8080", "localhost", false, true)
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "投币成功"})
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "投币失败"})
	}
}

//一键三连
func Sanlian(c*gin.Context){
	username:=c.PostForm("username")
	video:=c.PostForm("video")
	good:=c.PostForm("good")
	collect:=c.PostForm("collect")
	coin:=c.PostForm("coin")
	if model.SanLian(username, video,good,collect,coin) {
		//c.SetCookie("username", username, 10, "localhost:8080", "localhost", false, true)
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "三连成功"})
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "三连失败"})
	}
}

//