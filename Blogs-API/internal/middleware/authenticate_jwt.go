package middleware

import (
	m "blog/pkg/models"
	h "blog/pkg/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientToken := strings.Split(r.Header.Get("authorization"), " ")

		if len(clientToken) != 2 {
			errorMessage := h.ErrorMessageJson(h.UNAUTHORIZED_ACCESS, h.UNAUTHORIZED_ACCESS_CODE, h.UNAUTHORIZED_ACCESS_DETAIL)
			h.Response(w, r, errorMessage, http.StatusUnauthorized)
			return
		}

		token := clientToken[1]

		claims := &m.SignedUser{}

		jwtToken, err2 := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Method mismatch.")
			}

			return os.Getenv("JWT_SECRET_KEY"), nil
		})
		if err2 != nil {
			errorMessage := h.ErrorMessageJson(h.UNAUTHORIZED_ACCESS, h.UNAUTHORIZED_ACCESS_CODE, h.UNAUTHORIZED_ACCESS_DETAIL)
			h.Response(w, r, errorMessage, http.StatusUnauthorized)
			return
		}

		if !jwtToken.Valid {
			errorMessage := h.ErrorMessageJson(h.UNAUTHORIZED_ACCESS, h.UNAUTHORIZED_ACCESS_CODE, h.UNAUTHORIZED_ACCESS_DETAIL)
			h.Response(w, r, errorMessage, http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", claims.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
