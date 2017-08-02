package redis

// Options defines Redis queue options.
type Options struct {
	addr      string
	password  string
	db        int
	queueName string
}

// Option configures the Redis client.
type Option func(*Options)

// WithProject configures the Pubsub client with the named project.
func WithAddr(addr string) Option {
	return func(opts *Options) {
		opts.addr = addr
	}
}

func WithPassword(password string) Option {
	return func(opts *Options) {
		opts.password = password
	}
}

func WithDB(db int) Option {
	return func(opts *Options) {
		opts.db = db
	}
}

func WithQueueName(queueName string) Option {
	return func(opts *Options) {
		opts.queueName = queueName
	}
}
