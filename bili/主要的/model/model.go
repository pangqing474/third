package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB



func init(){
	db, _ =sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
}


//通过用户名和密码完成user表中注册操作
func UserSignup(username string,password string/*,phonenumber string,identify string*/)bool{
	stmt,err:=db.Prepare(
		"insert into alice(username,password/*,phonenumber,identify*/)values(?,?,?,?) ")
	if err!=nil{
		fmt.Println("fail to insert")
		return false
	}
	defer stmt.Close()

	_,err=stmt.Exec(username,password/*,phonenumber,identify*/)
	if err!=nil{
		fmt.Println("fail to insert")
		return false
	}

	return false
}

//登录
func UserSignin( username string,password string)bool{
	stmt,err:=db.Query("select password from alice where username=?",username)
	if err!=nil{
		log.Fatal(err)
		return false
	}

	defer stmt.Close()
	for stmt.Next() {
		var row string
		err = stmt.Scan(&row)
		if row==password{
			return true
		}
	}
	return false
}
//点赞关注数

// type Video struct {
//	username string `json:"username"`
	//image
//}
/*
//搜索
func Search(){
	rows, err := db.Query("SELECT username, image FROM video")
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		var user User
		rows.Scan(&user.Username, &user.Password)
		users = append(users, user)
	}
}
*/
type comment struct {
	Username string `json:"username"`
	Message string `json:"message"`
}

//发送评论
func SendMessage(username string,message string)bool{
	stmt,err:=db.Prepare(
		"insert into comment(username,message)values(?,?)")
	if err!=nil{
		fmt.Println("fail to insert the message")
		return false
	}
	defer stmt.Close()

	_,err=stmt.Exec(username,message)
	if err!=nil{
		fmt.Println("fail to insert")
		return false
	}
	return true
}

//展示评论
func ShowMesage(maxtime int)(error, []comment){
	stmt,err:=db.Query("select username,message from comment")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()
	var message []comment
	for stmt.Next() {
		var msg comment
		stmt.Scan(&msg.Username, &msg.Message)
		message=append(message,msg)
	}
	return err,message
}

//点赞
func Good(username string,video string,good string)bool{
	stmt,err:=db.Prepare(
		"insert into zan(username,video,good)values(?,?,?)")
	if err!=nil{
		fmt.Println("fail to insert")
		return false
	}
	defer stmt.Close()

	_,err=stmt.Exec(username,video,good)
	if err!=nil{
		fmt.Println("fail to 点赞")
		return false
	}
	return true
}


//收藏
func Shoucan(username string,video string,collect string)bool{
	stmt,err:=db.Prepare(
		"insert into collect(username,video,collect)values(?,?,?)")
	if err!=nil{
		fmt.Println("fail to insert")
		return false
	}
	defer stmt.Close()

	_,err=stmt.Exec(username,video,collect)
	if err!=nil{
		fmt.Println("fail to 收藏")
		return false
	}
	return true
}

//投币
func Toubi(username string,video string,coin string)bool{
	stmt,err:=db.Prepare(
		"insert into coin(username,video,coin)values(?,?,?)")
	if err!=nil{
		fmt.Println("fail to insert")
		return false
	}
	defer stmt.Close()

	_,err=stmt.Exec(username,video,coin)
	if err!=nil{
		fmt.Println("fail to 投币")
		return false
	}
	return true
}

//一键三连
func SanLian(username string,video string,good string,collect string,coin string)bool{
	stmt,err:=db.Prepare(
		"insert into person(username,video,good,collect,coin)values(?,?,?,?,?)")
	if err!=nil{
		fmt.Println("fail to insert")
		return false
	}
	defer stmt.Close()

	_,err=stmt.Exec(username,video,good,collect,coin)
	if err!=nil{
		fmt.Println("fail to 一键三连")
		return false
	}
	return true
}
/*
//判断该用户是不是up主
func (username string)bool{
	stmt, err := db.Query("select username from alice where name=?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()
	for stmt.Next() {
		var up string b


	}

	return err,message

}
 */


