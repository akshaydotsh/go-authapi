package models

type FieldError struct {
	FieldName    string `json:"field"`
	ErrorMessage string `json:"message"`
}

type ValidationError struct {
	Errors []FieldError `json:"errors"`
}
