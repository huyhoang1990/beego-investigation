package controllers

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/server/web"
	"github.com/huyhoang1990/beego-investigation/conf"
	"github.com/huyhoang1990/beego-investigation/entity"
	"github.com/huyhoang1990/beego-investigation/repo"
)

const (
	errorStatusCode   = 1
	successStatusCode = 0
)

type BaseController struct {
	web.Controller
	UserRepo repo.UserRepo
	User     *entity.User
}

type CookieRemember struct {
	UserId string
	Time   time.Time
}

func (c *BaseController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "404.tpl"
}

func (c *BaseController) Prepare() {
	if user, ok := c.GetSession(conf.LoginSessionName).(entity.User); ok && user.ID != "" {
		c.User = &user
	}

	if c.Ctx.Request.Method == "GET" && c.Ctx.Request.RequestURI != conf.LOGIN_PATH {
		c.Layout = "index.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Header"] = "header.html"
		c.LayoutSections["Footer"] = "footer.html"
	}
}

func (c *BaseController) ResponseJsonSuccess() {
	errorMessage := &entity.ErrorMessage{
		ErrorCode: successStatusCode,
		Message:   "success",
	}
	c.Data["json"] = errorMessage
	c.ServeJSON()
}

func (c *BaseController) ResponseJsonFailed(message string) {
	errorMessage := &entity.ErrorMessage{
		ErrorCode: errorStatusCode,
		Message:   message,
	}
	c.Data["json"] = errorMessage
	c.ServeJSON()
}

func (c *BaseController) SetUser(user entity.User) {
	if user.ID != "" {
		fmt.Println("set sesssionnnn")
		c.SetSession(conf.LoginSessionName, user)
		c.SetSession("uid", user.ID)
	} else {
		c.DelSession(conf.LoginSessionName)
		c.DelSession("uid")
		c.DestroySession()
	}
}
