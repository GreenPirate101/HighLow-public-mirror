package core

import (
	"context"
	"errors"
	"os"

	idtoken "google.golang.org/api/idtoken"
)


func GoogleVerify(ctx context.Context, code string) (*User, error)  {
	clientId, present := os.LookupEnv("GOOGLE_CLIENT_ID")
	if !present {
		return nil, errors.New("invalid conf: `GOOGLE_CLIENTID` not set")
	}

	payload, err := idtoken.Validate(ctx, code, clientId)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name: payload.Claims["name"].(string),
		Email: payload.Claims["email"].(string),
		Avatar: payload.Claims["picture"].(string),
	}

	return user, nil
}
