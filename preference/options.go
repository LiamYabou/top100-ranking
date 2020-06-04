package preference

// The place that the user can define the globle options among the packages.

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/streadway/amqp"
)

// Option represents the optional function.
type Option func(opts *Options)

type Options struct {
	DB *pgxpool.Pool
	AMQP *amqp.Connection
}

func LoadOptions(options ...Option) *Options {
	opts := new(Options)
	for _, option := range options {
		option(opts)
	}
	return opts
}

func WithDB(db *pgxpool.Pool) Option {
	return func(opts *Options) {
		opts.DB = db
	}
}

func WithOptions(options Options) Option {
	return func(opts *Options) {
		*opts = options
	}
}
