package http

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func gRPCtoHTTPErr(err error) (int, *ErrorResponse) {
	status, ok := status.FromError(err)
	if !ok {
		return 500, &ErrorResponse{
			Error: ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: "Internal server error",
			},
		}
	}

	switch status.Code() {
	case codes.InvalidArgument:
		return 400, &ErrorResponse{
			Error: ErrorDetail{
				Code:    "INVALID_ARGUMENT",
				Message: status.Message(),
			},
		}
	case codes.NotFound:
		return 404, &ErrorResponse{
			Error: ErrorDetail{
				Code:    "NOT_FOUND",
				Message: status.Message(),
			},
		}
	case codes.Unauthenticated:
		return 401, &ErrorResponse{
			Error: ErrorDetail{
				Code:    "UNAUTHENTICATED",
				Message: status.Message(),
			},
		}
	case codes.AlreadyExists:
		return 409, &ErrorResponse{
			Error: ErrorDetail{
				Code:    "ALREADY_EXISTS",
				Message: status.Message(),
			},
		}
	default:
		return 500, &ErrorResponse{
			Error: ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: "Internal server error",
			},
		}
	}
}
