package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/yadunut/CVWO/backend/gateway/internal/graph"
	"github.com/yadunut/CVWO/backend/proto"
)

type key string

const userIdKey = "user_id"

func AuthMiddleware(resolver *graph.Resolver) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			splitToken := strings.Split(token, "Bearer ")
			// if token is not available
			if len(splitToken) < 2 {
				next.ServeHTTP(w, r)
				return
			}
			res, err := resolver.AuthClient.Verify(r.Context(), &proto.VerifyRequest{Token: splitToken[1]})
			// if error, do nothing
			if err != nil || res.Status == proto.ResponseStatus_FAILURE {
				next.ServeHTTP(w, r)
				return
			}
			id, err := uuid.Parse(res.Id)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			// if not error, add user id to context. This can be used to inject
			ctx := context.WithValue(r.Context(), userIdKey, id)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func FromContext(ctx context.Context) uuid.UUID {
	raw, _ := ctx.Value(userIdKey).(uuid.UUID)
	return raw
}
