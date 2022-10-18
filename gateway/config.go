package gateway

var (
	DefaultGatewayURL = "wss://www.guilded.gg/websocket/v1"
)

type Config struct {
	// The gateway URL to connect to. If not specified, the default gateway URL will be used.
	url               string
	version           string
	ratelimiter       *Ratelimiter
	lastID            string
	maxReconnectTries int
}

func NewConfig() *Config {
	return &Config{
		url:               DefaultGatewayURL,
		ratelimiter:       NewRatelimiter(NewRatelimiterConfig()),
		maxReconnectTries: 5,
	}
}

func (c *Config) WithURL(url string) *Config {
	c.url = url
	return c
}

func (c *Config) WithMaxReconnectTries(maxReconnectTries int) *Config {
	c.maxReconnectTries = maxReconnectTries
	return c
}

type RatelimiterConfig struct {
	// The maximum number of requests that can be made per second. If not specified, the default is 120.
	maxRequestsPerMinute int
}

func NewRatelimiterConfig() *RatelimiterConfig {
	return &RatelimiterConfig{
		maxRequestsPerMinute: 120,
	}
}

func (c *RatelimiterConfig) WithMaxRequestsPerMinute(maxRequestsPerMinute int) *RatelimiterConfig {
	c.maxRequestsPerMinute = maxRequestsPerMinute
	return c
}
