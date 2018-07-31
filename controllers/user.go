package controllers


// Operations about Users
type UserController struct {
	BaseController
}


func (u *UserController) Login() {


	data := map[string]string{
		"username" : "admin",
		"password" : "123456",
	}

	u.Response(true, "login ok", data)


}


