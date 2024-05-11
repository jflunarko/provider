package directives

import (
	"context"
	"myapp/middleware"

	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func IsLogin(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	var (
		user = middleware.AuthContext(ctx)
	)

	if user == nil {
		return nil, &gqlerror.Error{
			Message: "User not logged in",
			Extensions: map[string]interface{}{
				"code": http.StatusUnauthorized,
			},
		}
	}

	return next(ctx)
}
