package handler

import (
	"github.com/arvrao/project-management-app/model"
	"github.com/gofiber/fiber/v2"
)

type UserApiStruct model.UserApiStruct

// get all users
func (UserApiStructPtr *UserApiStruct) UserRouteMthd(c *fiber.Ctx) error {
	switch c.Method() {
	case "CREATE":
		// SignUpUserMthd()
		break
	case "PUT":
	case "GET":
		switch UserApiStructPtr.RequestApiStruct.ApiAction {
		// case "LoginUser":
		// 	LoginUserMthd()
		// 	break
		// case "GetUser":
		// 	GetUser()

		// case "ActivateUser":
		// 	ActivateUser()
		// case "LogoutUser":
		// 	LogoutUser()
		// case "ForgotPassword":
		// 	ForgotPassword()
		// case "ActivateAccount":
		// 	ActivateAccount()
		}
	case "DELETE":
	default:
	}
	return c.Status(200).JSON("signed in successfully")
}

func (UserApiStructPtr *UserApiStruct) ProjectUserCnstr(c *fiber.Ctx) {
	UserApiStructPtr.RequestApiStruct.ApiMethod = c.Method()
	UserApiStructPtr.RequestApiStruct.ApiAction = c.Params("id")
	UserApiStructPtr.RequestApiStruct.ApiActionData = c.Params("IdData")

}

func (UserApiStructPtr *UserApiStruct) SignUpUserMthd() {

}

func (UserApiStructPtr *UserApiStruct) GetUserMthd() {

}
func (UserApiStructPtr *UserApiStruct) LogoutUserMthd() {

}
func (UserApiStructPtr *UserApiStruct) LoginUserMthd() {

}

func (UserApiStructPtr *UserApiStruct) ForgotPasswordMthd() {

}

func (UserApiStructPtr *UserApiStruct) ActivateAccountMthf() {

}

func (UserApiStructPtr *UserApiStruct) ActivateUserMthd() {

}
