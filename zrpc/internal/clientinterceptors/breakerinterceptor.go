package clientinterceptors

import (
	"context"
	"path"

	"github.com/Holyson/test-go-zero-cors/core/breaker"
	"github.com/Holyson/test-go-zero-cors/zrpc/internal/codes"
	"google.golang.org/grpc"
)

// BreakerInterceptor is an interceptor that acts as a circuit breaker.
func BreakerInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	breakerName := path.Join(cc.Target(), method)
	return breaker.DoWithAcceptable(breakerName, func() error {
		return invoker(ctx, method, req, reply, cc, opts...)
	}, codes.Acceptable)
}
