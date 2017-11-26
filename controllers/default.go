package controllers

import "github.com/astaxie/beego"
import "fmt"

type MainController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	// check username and password
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "admin" {
		fmt.Println("login")
	} else {
		fmt.Println("wrong username")
	}
	fmt.Println(username)
	fmt.Println(password)
	c.Data["username"] = username
	c.TplName = "login.tpl"
}
