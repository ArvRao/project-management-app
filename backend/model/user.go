package model

type ProjectUser struct {
	// gorm.Model
	Id       uint
	UserRgId string
	Name     string
	Age      uint8
	Email    string
	Password string
	Gender   string
}

type ErrorMsg struct {
	ErrMessage string
	AppError   error
	StatusCode uint8
}

type RequestApi struct {
	ProjUserStruct ProjectUser
	ApiMethod      string
	ApiAction      string
	ApiActionData  string
}

type ResponseApi struct {
	ProjUserStruct        ProjectUser   //update, get
	ErrorMsgStruct        ErrorMsg      //wrong inputs
	ProjectUserListStruct []ProjectUser // get all
}

type UserApiStruct struct {
	RequestApiStruct  RequestApi
	ResponseApiStruct ResponseApi
}
