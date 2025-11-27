package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/neberson/pos-go-expert-fullcycle/desafios/ratelimit/internal/limiter"
	"github.com/neberson/pos-go-expert-fullcycle/desafios/ratelimit/internal/middleware"
	"github.com/neberson/pos-go-expert-fullcycle/desafios/ratelimit/internal/storage"
)

func main() {
	_ = godotenv.Load()

	ipLimit, _ := strconv.Atoi(getEnv("IP_LIMIT", "10"))
	ipExpire, _ := strconv.Atoi(getEnv("IP_EXPIRE_SECONDS", "1"))
	blockSeconds, _ := strconv.Atoi(getEnv("BLOCK_SECONDS", "300"))

	tokenLimits := make(map[string]int)
	tokenExpires := make(map[string]int)
	for _, pair := range strings.Split(getEnv("TOKEN_LIMITS", ""), ",") {
		if pair == "" {
			continue
		}
		kv := strings.Split(pair, ":")
		if len(kv) == 2 {
			lim, _ := strconv.Atoi(kv[1])
			tokenLimits[kv[0]] = lim
		}
	}
	for _, pair := range strings.Split(getEnv("TOKEN_EXPIRES", ""), ",") {
		if pair == "" {
			continue
		}
		kv := strings.Split(pair, ":")
		if len(kv) == 2 {
			exp, _ := strconv.Atoi(kv[1])
			tokenExpires[kv[0]] = exp
		}
	}

	redisAddr := getEnv("REDIS_ADDR", "localhost:6379")
	redisPass := getEnv("REDIS_PASSWORD", "")
	redisDB, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))
	store := storage.NewRedisStorage(redisAddr, redisPass, redisDB)

	limiterCfg := limiter.RateLimiterConfig{
		IPLimit:         ipLimit,
		IPExpireSeconds: ipExpire,
		TokenLimits:     tokenLimits,
		TokenExpires:    tokenExpires,
		BlockSeconds:    blockSeconds,
	}
	lim := limiter.NewLimiter(store, limiterCfg)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", middleware.RateLimitMiddleware(lim)(mux))
}

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
