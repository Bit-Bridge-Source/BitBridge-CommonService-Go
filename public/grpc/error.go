package grpc

import (
	"context"
	"log"

	common_error "github.com/Bit-Bridge-Source/BitBridge-CommonService-Go/public/error"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GRPCErrorHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("Error occurred during method %v execution: %v", info.FullMethod, err)

		// Check if the error is of type *ServiceError
		if serr, ok := err.(*common_error.ServiceError); ok {
			switch serr.Code {
			case common_error.InternalServerError:
				return nil, status.Error(codes.Internal, serr.Message)
			case common_error.NotFound:
				return nil, status.Error(codes.NotFound, serr.Message)
			case common_error.BadRequest:
				return nil, status.Error(codes.InvalidArgument, serr.Message)
			case common_error.Unauthorized:
				return nil, status.Error(codes.Unauthenticated, serr.Message)
			case common_error.Forbidden:
				return nil, status.Error(codes.PermissionDenied, serr.Message)
			case common_error.Conflict:
				return nil, status.Error(codes.AlreadyExists, serr.Message)
			case common_error.TooManyRequests:
				return nil, status.Error(codes.ResourceExhausted, serr.Message)
			case common_error.ServiceUnavailable:
				return nil, status.Error(codes.Unavailable, serr.Message)
			case common_error.NotImplemented:
				return nil, status.Error(codes.Unimplemented, serr.Message)
			case common_error.Timeout:
				return nil, status.Error(codes.DeadlineExceeded, serr.Message)
			default:
				return nil, status.Error(codes.Unknown, serr.Message)
			}
		}

		// If it's not a ServiceError or another recognized error, return a generic error
		return nil, status.Error(codes.Unknown, "An unknown error occurred")
	}

	return resp, nil
}
