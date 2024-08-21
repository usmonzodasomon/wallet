package middlewares

import (
	"context"
	"encoding/json"
	"github/usmonzodasomon/wallet/internal/config"
	"github/usmonzodasomon/wallet/pkg/helpers"
	"io"
	"net/http"
	"strings"
)

func CheckHashMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-UserId")
		if userID == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "X-UserId header is required"})
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid body"})
			return
		}

		r.Body = io.NopCloser(strings.NewReader(string(body)))

		hash := r.Header.Get("X-Digest")
		if hash != helpers.ToSha1(string(body), config.Cfg.SecretKey) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid hash"})
			return
		}

		ctx := context.WithValue(r.Context(), "X-UserId", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
