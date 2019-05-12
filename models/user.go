package models

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"gopkg.in/go-playground/validator.v9"
	_ "gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

type User struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty" validate:"-"`
	Name      string        `json:"name" bson:"name" validate:"required,min=3,max=15"`
	Email     string        `json:"email" bson:"email" validate:"required,email"`
	Password  string        `json:"password" bson:"password" validate:"required,min=3,max=15"`
	CreatedAt int           `json:"created_at" bson:"created_at" validate:"numeric,omitempty"`
	Role      string        `json:"role" bson:"role" validate:"required,oneof=user admin superadmin"`
}

type UserLoginCreds struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=user admin superadmin"`
}

func constructValidationError(err error) ValidationError {
	var validationError ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		fmt.Println(err.Tag())
		var fieldError FieldError
		fieldError.FieldName = err.Tag()
		fieldError.ErrorMessage = "Validation Error in " + err.Tag() + ": " + err.Value().(string)
		validationError.Errors = append(validationError.Errors, fieldError)
	}
	return validationError
}

func (user *User) Validate() (ValidationError, bool) {
	err := validate.Struct(user)
	if err != nil {
		return constructValidationError(err), false
	}
	return ValidationError{}, true
}

func (user *UserLoginCreds) Validate() (ValidationError, bool) {
	err := validate.Struct(user)
	if err != nil {
		return constructValidationError(err), false
	}
	return ValidationError{}, true
}
