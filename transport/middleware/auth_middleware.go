package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

const (
	externalPrefix = "external"
	internalPrefix = "internal"
)

func (rh *RestHandler) AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		clientPublicIP := r.Header.Get("public-ip")
		ctx := context.WithValue(r.Context(), constant.ClientPublicIP, clientPublicIP)
		ctx = context.WithValue(ctx, constant.IsAgrabaCall, true)

		r = r.WithContext(ctx)

		if isExcludedFromAuthAgrabaInterceptor(r.URL.Path) {
			c.Next(ctx)
			return
		}

		switch getFirstPrefix(string(c.URI().Path())) {
		case externalPrefix:
			authHeader := c.GetHeader("Authorization")

			ctx, err := s.agrabaAuth.Authorize(r.Context(), authHeader)
			if err != nil {
				response := &model.AgrabaHTTPResponse{Message: err.Error()}
				if errors.Is(err, internaluser.ErrInternal) {
					httpResponseWrite(w, response, http.StatusInternalServerError)
					return
				}

				httpResponseWrite(w, response, http.StatusUnauthorized)
				return
			}

			r = r.WithContext(ctx)
		case internalPrefix:
			authHeader := r.Header.Get("Authorization")

			ctx, err := s.agrabaAuth.Authorize(r.Context(), authHeader)
			if err != nil {
				response := &model.AgrabaHTTPResponse{Message: err.Error()}
				if errors.Is(err, internaluser.ErrInternal) {
					httpResponseWrite(w, response, http.StatusInternalServerError)
					return
				}

				httpResponseWrite(w, response, http.StatusUnauthorized)
				return
			}

			r = r.WithContext(ctx)
		default:
			c.JSON(http.StatusInternalServerError, "unauthorized access")
			return
		}

		c.Next(ctx)
	}
}

func getFirstPrefix(path string) string {
	paths := strings.Split(path, "/")
	if len(paths) > 1 {
		return paths[1]
	}
	return ""
}
