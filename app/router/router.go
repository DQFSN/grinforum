package router

import (
	"net/http"

	"github.com/chalvern/apollo/app/controllers"
	"github.com/chalvern/apollo/configs/initializer"
	"github.com/gin-gonic/gin"
)

// pong for ping
func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// 定义 router
func routerInit() {
	get("ping_pong", "/ping", pong)
	get("home_page", "/", controllers.HomeIndex)
	get("about", "/about", pong)

	// captcha
	get("captcha_get", initializer.Captcha.URLPrefix+":id", controllers.GetCaptcha)

	// account
	get("signup", "/signup", controllers.SignupGet)
	post("signup_post", "/signup", controllers.SigninPost)
	get("signin", "/signin", controllers.SigninGet)
	get("signout", "/signout", pong)

	// user
	get("user_info", "/user/info", pong)
}
