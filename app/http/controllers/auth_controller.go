package controllers

import (
	"goblog/app/models/user"
	"goblog/app/requests"
	"goblog/pkg/view"
	"net/http"
)

// AuthController 处理用户认证
type AuthController struct {
}

// Register 注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 处理注册逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {

	// 1. 初始化数据
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	// 2. 表单规则
	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		// 5.1 表单不通过 —— 重新显示表单
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
	} else {
		// _user.Create()
		// if _user.ID > 0 {
		// 	fmt.Fprint(w, "插入成功，ID为"+_user.GetStringID())
		// } else {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	fmt.Fprint(w, "创建用户失败，请联系管理员")
		// }
	}

}
