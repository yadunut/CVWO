// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type BaseError interface {
	IsBaseError()
	GetMessage() string
}

type LoginResponse interface {
	IsLoginResponse()
}

type RegisterResponse interface {
	IsRegisterResponse()
}

type Comment struct {
	ID      string  `json:"id"`
	Author  *User   `json:"author"`
	Message string  `json:"message"`
	Parent  *Thread `json:"parent"`
}

type InvalidLoginError struct {
	Message string `json:"message"`
}

func (InvalidLoginError) IsBaseError()            {}
func (this InvalidLoginError) GetMessage() string { return this.Message }

func (InvalidLoginError) IsLoginResponse() {}

type InvalidRegistrationError struct {
	Message string `json:"message"`
}

func (InvalidRegistrationError) IsBaseError()            {}
func (this InvalidRegistrationError) GetMessage() string { return this.Message }

func (InvalidRegistrationError) IsRegisterResponse() {}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SuccessfulLogin struct {
	Token string `json:"token"`
}

func (SuccessfulLogin) IsLoginResponse() {}

type SuccessfulRegistration struct {
	Token string `json:"token"`
}

func (SuccessfulRegistration) IsRegisterResponse() {}

type Thread struct {
	ID       string     `json:"id"`
	Owner    *User      `json:"owner"`
	Title    string     `json:"title"`
	Comments []*Comment `json:"comments"`
	Likes    []*User    `json:"likes"`
}

type User struct {
	ID       string     `json:"id"`
	Username string     `json:"username"`
	Bio      string     `json:"bio"`
	Threads  []*Thread  `json:"threads"`
	Comments []*Comment `json:"comments"`
}
