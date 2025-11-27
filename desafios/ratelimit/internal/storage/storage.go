package storage

// LimiterStorage define a interface para persistência do rate limiter.
type LimiterStorage interface {
	// Incrementa o contador de requisições para uma chave (IP ou Token) e retorna o novo valor e o tempo restante de expiração.
	Increment(key string, limit int, expireSeconds int) (count int, ttlSeconds int, err error)
	// Retorna o contador atual e o tempo restante de expiração para uma chave.
	Get(key string) (count int, ttlSeconds int, err error)
	// Bloqueia explicitamente uma chave por um tempo determinado.
	Block(key string, blockSeconds int) error
	// Verifica se a chave está bloqueada.
	IsBlocked(key string) (bool, int, error)
}
