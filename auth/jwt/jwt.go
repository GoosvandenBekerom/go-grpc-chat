package jwt

import (
	"context"
	"github.com/GoosvandenBekerom/go-grpc-chat/config"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/credentials"
)

// Currently unused, keeping this file here because at some point I want to make the gRPC connection tls secured
// and send a jwt token for user specific information using a credentials.PerRPCCredentials for each rpc call

type token struct {
	token string
}

func GenerateAndSign(user string) (credentials.PerRPCCredentials, error) {
	encoded, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user": user,
	}).SignedString([]byte(config.JwtSecret))
	return token{encoded}, err
}

func (t token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": t.token,
	}, nil
}

func (t token) RequireTransportSecurity() bool {
	return true
}
