package limiter

import (
	"errors"
	"testing"
)

type fakeStorage struct {
	count   map[string]int
	blocked map[string]bool
}

func newFakeStorage() *fakeStorage {
	return &fakeStorage{
		count:   make(map[string]int),
		blocked: make(map[string]bool),
	}
}

func (f *fakeStorage) Increment(key string, limit int, expireSeconds int) (int, int, error) {
	f.count[key]++
	return f.count[key], expireSeconds, nil
}
func (f *fakeStorage) Get(key string) (int, int, error) {
	return f.count[key], 1, nil
}
func (f *fakeStorage) Block(key string, blockSeconds int) error {
	f.blocked[key] = true
	return nil
}
func (f *fakeStorage) IsBlocked(key string) (bool, int, error) {
	return f.blocked[key], 1, nil
}

func TestLimiter_IPLimit(t *testing.T) {
	cfg := RateLimiterConfig{IPLimit: 2, IPExpireSeconds: 1, BlockSeconds: 10}
	lim := NewLimiter(newFakeStorage(), cfg)
	ip := "1.2.3.4"
	token := ""
	if err := lim.Allow(ip, token); err != nil {
		t.Fatalf("primeira req deve passar: %v", err)
	}
	if err := lim.Allow(ip, token); err != nil {
		t.Fatalf("segunda req deve passar: %v", err)
	}
	err := lim.Allow(ip, token)
	if !errors.Is(err, ErrRateLimited) {
		t.Fatalf("terceira req deve ser limitada, err: %v", err)
	}
}

func TestLimiter_TokenLimit(t *testing.T) {
	cfg := RateLimiterConfig{
		IPLimit: 1, IPExpireSeconds: 1, BlockSeconds: 10,
		TokenLimits:  map[string]int{"abc": 2},
		TokenExpires: map[string]int{"abc": 1},
	}
	lim := NewLimiter(newFakeStorage(), cfg)
	ip := "1.2.3.4"
	token := "abc"
	if err := lim.Allow(ip, token); err != nil {
		t.Fatalf("primeira req token deve passar: %v", err)
	}
	if err := lim.Allow(ip, token); err != nil {
		t.Fatalf("segunda req token deve passar: %v", err)
	}
	err := lim.Allow(ip, token)
	if !errors.Is(err, ErrRateLimited) {
		t.Fatalf("terceira req token deve ser limitada, err: %v", err)
	}
}

func TestLimiter_TokenSobrepoeIP(t *testing.T) {
	cfg := RateLimiterConfig{
		IPLimit: 1, IPExpireSeconds: 1, BlockSeconds: 10,
		TokenLimits:  map[string]int{"abc": 3},
		TokenExpires: map[string]int{"abc": 1},
	}
	lim := NewLimiter(newFakeStorage(), cfg)
	ip := "1.2.3.4"
	token := "abc"
	for i := 0; i < 3; i++ {
		if err := lim.Allow(ip, token); err != nil {
			t.Fatalf("req %d token deve passar: %v", i+1, err)
		}
	}
	err := lim.Allow(ip, token)
	if !errors.Is(err, ErrRateLimited) {
		t.Fatalf("quarta req token deve ser limitada, err: %v", err)
	}
}

func TestLimiter_Bloqueio(t *testing.T) {
	fs := newFakeStorage()
	cfg := RateLimiterConfig{IPLimit: 1, IPExpireSeconds: 1, BlockSeconds: 10}
	lim := NewLimiter(fs, cfg)
	ip := "1.2.3.4"
	_ = lim.Allow(ip, "")
	_ = lim.Allow(ip, "") // deve bloquear
	if ok, _, _ := fs.IsBlocked("ip:" + ip); !ok {
		t.Fatalf("ip deveria estar bloqueado")
	}
}
