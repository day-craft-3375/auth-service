package postgres

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = 5 * time.Second
	_defaultMaxIdleConns = 1
)

// Config конфигурация для подключения к БД
type Config struct {
	MaxPoolSize  int
	ConnAttempts int
	ConnTimeout  time.Duration
	MaxIdleConns int
}

// Postgres клиент для работы с PostgreSQL БД
type Postgres struct {
	db     *sqlx.DB
	config Config
	closed bool
}

// New создает новый Postgres клиент с указанной URL и опциями
func New(url string, opts ...Option) (*Postgres, error) {
	config := Config{
		MaxPoolSize:  _defaultMaxPoolSize,
		ConnAttempts: _defaultConnAttempts,
		ConnTimeout:  _defaultConnTimeout,
		MaxIdleConns: _defaultMaxIdleConns,
	}

	// применяем опции
	for _, opt := range opts {
		opt(&config)
	}

	pg := &Postgres{
		config: config,
	}

	// инициализируем подключение
	if err := pg.connect(url); err != nil {
		return nil, err
	}

	return pg, nil
}

// connect устанавливает подключение к БД с повторными попытками
func (p *Postgres) connect(url string) error {
	var err error
	attempts := p.config.ConnAttempts

	for attempts > 0 {
		p.db, err = sqlx.Open("pgx", url)
		if err == nil {
			// проверяем подключение
			ctx, cancel := context.WithTimeout(context.Background(), p.config.ConnTimeout)
			err = p.db.PingContext(ctx)
			cancel()

			if err == nil {
				// конфигурируем пул
				p.db.SetMaxOpenConns(p.config.MaxPoolSize)
				p.db.SetMaxIdleConns(p.config.MaxIdleConns)
				return nil
			}
		}

		attempts--
		if attempts > 0 {
			time.Sleep(p.config.ConnTimeout)
		}
	}

	return fmt.Errorf("postgres - connect failed after %d attempts: %w", p.config.ConnAttempts, err)
}

// DB возвращает sqlx.DB для выполнения запросов
func (p *Postgres) DB() *sqlx.DB {
	if p.closed {
		return nil
	}
	return p.db
}

// Health проверяет здоровье подключения
func (p *Postgres) Health(ctx context.Context) error {
	if p.closed {
		return fmt.Errorf("postgres client is closed")
	}

	if err := p.db.PingContext(ctx); err != nil {
		return fmt.Errorf("postgres - health check failed: %w", err)
	}

	return nil
}

// Close закрывает подключение к базе данных
func (p *Postgres) Close() error {
	if p.closed {
		return nil
	}

	p.closed = true

	if p.db != nil {
		if err := p.db.Close(); err != nil {
			return fmt.Errorf("postgres - close failed: %w", err)
		}
	}

	return nil
}
