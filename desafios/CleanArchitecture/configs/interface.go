package config

type ConfInterface interface {
	LoadConfig(path string) (*conf, error)
}
