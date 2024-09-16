package server

import (
	"context"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type validatorInterceptor struct {
	validate *validator.Validate
}

func NewValidatorInterceptor(v *validator.Validate) grpc.UnaryServerInterceptor {
	interceptor := &validatorInterceptor{
		validate: v,
	}
	return interceptor.interceptor
}

func (vi *validatorInterceptor) interceptor(
	ctx context.Context,
	req any,
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	if err := vi.validate.Struct(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation failed: %v", err)
	}
	return handler(ctx, req)
}
