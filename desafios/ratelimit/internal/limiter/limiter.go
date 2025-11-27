package limiter

import (
	"errors"
	"fmt"
	"strings"
)

// Configuração para limites por IP e por Token
type RateLimiterConfig struct {
	IPLimit         int
	IPExpireSeconds int
	TokenLimits     map[string]int // token: limit
	TokenExpires    map[string]int // token: expire seconds
	BlockSeconds    int
}

type Limiter struct {
	storage Storage
	config  RateLimiterConfig
}

type Storage interface {
	Increment(key string, limit int, expireSeconds int) (int, int, error)
	Get(key string) (int, int, error)
	Block(key string, blockSeconds int) error
	IsBlocked(key string) (bool, int, error)
}

var ErrRateLimited = errors.New("you have reached the maximum number of requests or actions allowed within a certain time frame")

func NewLimiter(storage Storage, config RateLimiterConfig) *Limiter {
	return &Limiter{storage: storage, config: config}
}

// GetKeyAndLimit determina a chave e o limite a ser usado (token tem prioridade sobre IP)
func (l *Limiter) GetKeyAndLimit(ip, token string) (key string, limit, expire int) {
	if token != "" {
		if lim, ok := l.config.TokenLimits[token]; ok {
			exp := l.config.TokenExpires[token]
			return "token:" + token, lim, exp
		}
	}
	return "ip:" + ip, l.config.IPLimit, l.config.IPExpireSeconds
}

// Allow verifica se a requisição pode prosseguir, incrementa contador e bloqueia se necessário
func (l *Limiter) Allow(ip, token string) error {
	key, limit, expire := l.GetKeyAndLimit(ip, token)
	blocked, ttl, err := l.storage.IsBlocked(key)
	if err != nil {
		return err
	}
	if blocked {
		return fmt.Errorf("blocked for %ds: %w", ttl, ErrRateLimited)
	}
	count, _, err := l.storage.Increment(key, limit, expire)
	if err != nil {
		return err
	}
	if count > limit {
		_ = l.storage.Block(key, l.config.BlockSeconds)
		return ErrRateLimited
	}
	return nil
}

// ParseToken extrai o token do header API_KEY
func ParseToken(header string) string {
	return strings.TrimSpace(header)
}
