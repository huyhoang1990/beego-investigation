package controllers

type HomeController struct {
	*BaseController
}

func NewHomeController(baseController *BaseController) *HomeController {
	return &HomeController{
		BaseController: baseController,
	}
}

func (c *HomeController) Home() {
	c.TplName = "home.html"
	_ = c.Render()
}
