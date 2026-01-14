package postgres

import "time"

// Option применяет конфигурацию к Config
type Option func(*Config)

// WithMaxPoolSize устанавливает максимальный размер пула подключений
func WithMaxPoolSize(size int) Option {
	return func(c *Config) {
		if size > 0 {
			c.MaxPoolSize = size
		}
	}
}

// WithConnAttempts устанавливает количество попыток подключения
func WithConnAttempts(attempts int) Option {
	return func(c *Config) {
		if attempts > 0 {
			c.ConnAttempts = attempts
		}
	}
}

// WithConnTimeout устанавливает timeout подключения
func WithConnTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		if timeout > 0 {
			c.ConnTimeout = timeout
		}
	}
}

// WithMaxIdleConns устанавливает максимальное количество неиспользуемых подключений
func WithMaxIdleConns(conns int) Option {
	return func(c *Config) {
		if conns > 0 {
			c.MaxIdleConns = conns
		}
	}
}
