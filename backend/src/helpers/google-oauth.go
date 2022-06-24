package helpers

import (
	"context"
	"os"

	"google.golang.org/api/idtoken"
)

// https://cloud.google.com/docs/authentication/production#auth-cloud-implicit-go
// https://pkg.go.dev/google.golang.org/api/oauth2/v2
var dataClientId = os.Getenv("GOOGLE_DATA_CLIENT_ID")

func AuthenticateGoogle(token string) (map[string]interface{}, error) {

	payload, err := idtoken.Validate(context.Background(), token, dataClientId)
	if err != nil {
		return nil, err
	}
	return payload.Claims, nil
}
