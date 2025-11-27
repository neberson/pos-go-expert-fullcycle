package middleware

import (
	"net/http"
	"strings"

	"github.com/neberson/pos-go-expert-fullcycle/desafios/ratelimit/internal/limiter"
)

func RateLimitMiddleware(l *limiter.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			if ipHeader := r.Header.Get("X-Real-IP"); ipHeader != "" {
				ip = ipHeader
			} else if ipHeader = r.Header.Get("X-Forwarded-For"); ipHeader != "" {
				ip = strings.Split(ipHeader, ",")[0]
			}
			token := limiter.ParseToken(r.Header.Get("API_KEY"))
			err := l.Allow(ip, token)
			if err != nil {
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("you have reached the maximum number of requests or actions allowed within a certain time frame"))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
