package grpc

import (
	"context"

	errors "github.com/Bit-Bridge-Source/BitBridge-CommonService-Go/public/error"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const userIDKey contextKey = "user_id"

func AuthUnaryInterceptor(secret []byte, publicMethods map[string]struct{}) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Check if the method is public; if it is, bypass the authentication
		if _, ok := publicMethods[info.FullMethod]; ok {
			return handler(ctx, req)
		}

		md, _ := metadata.FromIncomingContext(ctx)
		authHeader := md.Get("authorization")
		if len(authHeader) == 0 {
			return nil, status.Errorf(errors.Unauthorized, "Authorization token is required")
		}

		tokenStr := authHeader[0]
		if len(tokenStr) <= 7 {
			return nil, status.Errorf(errors.Unauthorized, "Invalid authorization token")
		}

		tokenStr = tokenStr[7:]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil {
			return nil, status.Errorf(errors.Unauthorized, "Invalid authorization token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return nil, status.Errorf(errors.Unauthorized, "Invalid authorization token")
		}

		ctx = context.WithValue(ctx, userIDKey, claims["user_id"])
		return handler(ctx, req)
	}
}
