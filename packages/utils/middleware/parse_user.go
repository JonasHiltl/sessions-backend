package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/jonashiltl/sessions-backend/packages/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ParseUser(c context.Context) (types.JwtPayload, error) {
	payload := ""
	if md, ok := metadata.FromIncomingContext(c); ok {
		if jwt, ok := md["jwt_payload"]; ok {
			payload = strings.Join(jwt, ",")
		}
	}

	stringBytes, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return types.JwtPayload{}, status.Error(codes.Unauthenticated, "Not logged in")
	}

	result := types.JwtPayload{}
	json.Unmarshal(stringBytes, &result)

	return result, nil
}
