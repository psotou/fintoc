package fintoc

import "fmt"

type Error struct {
	ErrorObject `json:"error"`
}

type ErrorObject struct {
	Type    string `json:"type"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Param   string `json:"param"`
	DocUrl  string `json:"doc_url"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s\n%s\n", e.Type, e.Message)
}
