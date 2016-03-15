package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Help() revel.Result {
	// 获取前端get请求操作
	var UserName, Password string
	c.Params.Bind(&UserName, "userName")
	c.Params.Bind(&Password, "password")
	//在控制台打印
	revel.INFO.Printf("userName:%s,password:%s\n", UserName, Password)
	//在视图(模板)渲染
	return c.Render(UserName, Password)
}

func (c App) Form() revel.Result {
	// 获取前端post请求操作
	c.Request.ParseForm()
	UserName := c.Request.Form.Get("userName")
	Password := c.Request.Form.Get("password")
	revel.INFO.Printf("userName:%s,password:%s\n", UserName, Password)
	//返回空
	return nil
}

func (c App) showVersion() revel.Result {
	return nil
}
