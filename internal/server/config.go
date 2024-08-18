package server

type Config interface {
	Port() int
}

type DefaultConfig struct{}

func NewConfig() DefaultConfig {
	return DefaultConfig{}
}

// Port is hardcoded for now due to developer laziness
func (c DefaultConfig) Port() int {
	return 4001
}
