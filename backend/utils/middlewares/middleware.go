package middlewares

import (
	"backend/config"
	"backend/utils/auth"
	"context"
	"github.com/golang-jwt/jwt/v5/request"
	"github.com/sirupsen/logrus"
	"net/http"
)

type contextKey string

const userIDKey contextKey = "userId"

func AuthenticateUserMiddleware(cfg *config.LocalConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// get userId from header
			headerUserId := r.Header.Get("UserId")
			if headerUserId == "" {
				http.Error(w, "Unauthorized header", http.StatusUnauthorized)
				return
			}

			// get auth token
			bearerExtractor := request.BearerExtractor{}
			tokenString, err := bearerExtractor.ExtractToken(r)
			if err != nil {
				logrus.Error(err)
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// validate token
			token, err := auth.ValidateToken(tokenString, []byte(cfg.JWTSecretKey))
			if err != nil {
				logrus.Error(err)
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			// validate if userId in the header is the same one in the token
			claims, ok := token.Claims.(*auth.CustomClaims)
			if !ok || !token.Valid {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			if claims.UserData.Id != headerUserId {
				http.Error(w, "User Id mismatch", http.StatusUnauthorized)
				return
			}

			// add userId to context
			ctx := context.WithValue(r.Context(), userIDKey, headerUserId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
