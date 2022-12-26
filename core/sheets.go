package core

import (
	"context"
	"os"
	"regexp"
	"strconv"

	oauth2JWT "golang.org/x/oauth2/jwt"
	option "google.golang.org/api/option"
	sheets "google.golang.org/api/sheets/v4"
)

var (
	conf = &oauth2JWT.Config{
		Email: os.Getenv("GOOGLE_SERVICE_ACCOUNT_EMAIL"),
		PrivateKeyID: os.Getenv("GOOGLE_SERVICE_ACCOUNT_PRIVATE_KEY_ID"),
		PrivateKey: []byte(os.Getenv("GOOGLE_SERVICE_ACCOUNT_PRIVATE_KEY")),
		TokenURL: os.Getenv("GOOGLE_SERVICE_ACCOUNT_TOKEN_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets",
		},
	}
	spreadsheetId = os.Getenv("GOOGLE_SHEET_ID")
)

func AddUser(ctx context.Context, user *User) (*SheetsAddUserResponse, error) {
	client := conf.Client(ctx)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	usersRange := UsersRangefromRowId(1)

	res, err := srv.Spreadsheets.Values.Append(
		spreadsheetId,
		usersRange,
		&sheets.ValueRange{
			MajorDimension: "ROWS",
			Range: usersRange,
			Values: [][]interface{}{
				{
					user.Name,
					user.Email,
					user.Avatar,
				},
			},
		},
	).ValueInputOption("RAW").Do()
	if err != nil {
		return nil, err
	}

	rowrange := res.Updates.UpdatedRange
	re := regexp.MustCompile(`[a-zA-Z!]+(?P<x>\d+):.*`)
	rowid, err := strconv.Atoi(
		re.FindStringSubmatch(rowrange)[1],
	)
	if err != nil {
		return nil, err
	}

	return &SheetsAddUserResponse{
		RowId:    rowid,
		RowRange: rowrange,
	}, nil
}

func GetUser(ctx context.Context, rowId int) (*User, error) {
	client := conf.Client(ctx)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	rowrange := UsersRangefromRowId(rowId)

	res, err := srv.Spreadsheets.Values.Get(
		spreadsheetId,
		rowrange,
	).Do()
	if err != nil {
		return nil, err
	}

	return &User{
		Name:   res.Values[0][0].(string),
		Email:  res.Values[0][1].(string),
		Avatar: res.Values[0][2].(string),
	}, nil
}
