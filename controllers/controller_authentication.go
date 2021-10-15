package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/huyhoang1990/beego-investigation/conf"
	"github.com/huyhoang1990/beego-investigation/entity"
	"github.com/huyhoang1990/beego-investigation/service"
)

func NewAuthenticationController(
	baseController *BaseController,
	authService service.AuthService,
) *AuthenticationController {
	return &AuthenticationController{
		BaseController: baseController,
		AuthService:    authService,
	}
}

type AuthenticationController struct {
	*BaseController
	AuthService service.AuthService
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	RepeatedPassword string `json:"repeated_password"`
}

func (c *AuthenticationController) Login() {
	if c.Ctx.Request.Method == "GET" {
		fmt.Println("vaoooooo")
		fmt.Println(c.GetSession(conf.LoginSessionName))
		if user, ok := c.GetSession(conf.LoginSessionName).(entity.User); ok && user.ID != "" {
			fmt.Println("vaoooooo")
			fmt.Println(user)
			c.Ctx.Redirect(302, conf.HOME_PATH)
		} else {
			c.TplName = "login.html"
			_ = c.Render()
		}

	} else {
		req := LoginReq{}
		if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&req); err != nil {
			c.ResponseJsonFailed(err.Error())
		}
		user, err := c.AuthService.ValidateUser(c.Ctx.Request.Context(), req.Username, req.Password)

		if err != nil {
			c.ResponseJsonFailed(err.Error())
		}

		c.SetUser(*user)

		c.ResponseJsonSuccess()
	}
}
