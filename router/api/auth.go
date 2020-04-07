package api

import (
	"net/http"
	"um_sys/pkg/app"
	"um_sys/pkg/e"
	"um_sys/pkg/util"
	"um_sys/service/auth_service"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Email    string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Register(c *gin.Context) {
	appG := app.Gin{Ctx: c}

	email := c.PostForm("email")
	password := c.PostForm("password")

	//参数检测
	a := auth{Email: email, Password: password}
	valid := validation.Validation{}
	ok, _ := valid.Valid(&a)
	if !ok {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAM, nil)
		return
	}

	//检测该email是否存在
	auth := auth_service.Auth{Email: email}
	exist, err := auth.ExistEmail()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if exist {
		appG.Response(http.StatusOK, e.ERROR_EXIST_USER, nil)
		return
	}

	//插入accounts account_members
	auth = auth_service.Auth{Email: email, Password: password}
	id, err := auth.Register()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]uint64{
		"id": id,
	})

	return
}

func Login(c *gin.Context) {
	//获取email/password
	//进行校验
	//校验通过，修改登录时间和ip
	//校验通过，生成token下发
	appG := app.Gin{Ctx: c}

	email := c.PostForm("email")
	password := c.PostForm("password")
	a := auth{Email: email, Password: password}
	valid := validation.Validation{}

	ok, _ := valid.Valid(a)
	if !ok {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAM, nil)
		return
	}

	authservice := auth_service.Auth{Email: email, Password: password}
	exist, err := authservice.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if !exist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(email, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})

	return
}

//todo: 第三方登录
func PlatformLogin(c *gin.Context) {
}
