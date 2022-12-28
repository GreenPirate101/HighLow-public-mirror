package core

import (
	"errors"
	"strings"
)

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

type User struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

type UsersLoginArgsT struct {
	Code string
}

func UsersLoginArgs(args map[string]interface{}) (*UsersLoginArgsT, error) {
	code, exist := args["code"].(string)
	if !exist {
		return nil, errors.New("invalid request: `code` is required")
	}

	ula := &UsersLoginArgsT{
		code,
	}

	return ula, nil
}

type UsersProfileArgsT struct {
	Token string
}

func UsersProfileArgs(args map[string]map[string]string) (*UsersProfileArgsT, error) {
	headers := args["__ow_headers"]
	authorization, exist := headers["authorization"]
	if !exist {
		return nil, errors.New("invalid request: Bearer `token` is required")
	}
	token := strings.TrimSpace(
		strings.Replace(
			authorization, "Bearer", "", 1,
		),
	)
	upa := &UsersProfileArgsT{
		token,
	}

	return upa, nil
}

type SheetsAddUserResponse struct {
	RowId    int    `json:"rowid"`
	RowRange string `json:"rowrange"`
}


