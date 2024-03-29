package types

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserParams struct {
	FirstName string `json:"firstName" validate:"required,min=2"`
	LastName  string `json:"lastName" validate:"required,min=2"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

type UpdateUserParams struct {
	FirstName string `json:"firstName,omitempty" validate:"omitempty,min=2"`
	LastName  string `json:"lastName,omitempty" validate:"omitempty,min=2"`
}

type User struct {
	ID                string `json:"id,omitempty"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"-"`
}

var errorMessages = map[string]string{
	"firstName": "First name must be at least 2 characters",
	"lastName":  "Last name must be at least 2 characters",
	"email":     "Invalid email address",
	"password":  "Password must be at least 8 characters",
}

func (params *UpdateUserParams) Validate() map[string]string {

	params.FirstName = strings.TrimSpace(params.FirstName)
	params.LastName = strings.TrimSpace(params.LastName)

	validate := validator.New()
	errs := validate.Struct(params)

	if errs != nil {
		errorMap := make(map[string]string)
		for _, err := range errs.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(params).Elem().FieldByName(err.Field())
			tag := string(field.Tag.Get("json"))
			errorMap[tag] = errorMessages[tag]
		}
		return errorMap
	}
	return nil
}

func (params *CreateUserParams) Validate() map[string]string {
	params.FirstName = strings.TrimSpace(params.FirstName)
	params.LastName = strings.TrimSpace(params.LastName)
	params.Email = strings.TrimSpace(params.Email)
	params.Password = strings.TrimSpace(params.Password)

	validate := validator.New()
	errs := validate.Struct(params)
	if errs != nil {
		errorMap := make(map[string]string)
		for _, err := range errs.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(params).Elem().FieldByName(err.Field())
			tag := string(field.Tag.Get("json"))
			errorMap[tag] = errorMessages[tag]
		}
		return errorMap
	}
	return nil
}

func NewUserfromParams(params *CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}
