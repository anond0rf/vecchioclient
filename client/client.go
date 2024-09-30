package client

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const host = "vecchiochan.com"

type VecchioClient struct {
	client    *http.Client
	verbose   bool
	userAgent string
	logger    *log.Logger
}

type Config struct {
	Client    *http.Client
	Verbose   bool
	UserAgent string
	Logger    *log.Logger
}

var DefaultConfig = Config{
	Client:    &http.Client{Timeout: 30 * time.Second},
	Verbose:   false,
	UserAgent: "",
	Logger:    log.New(os.Stderr, "", log.LstdFlags),
}

func NewVecchioClient() *VecchioClient {
	return &VecchioClient{
		client:    DefaultConfig.Client,
		verbose:   DefaultConfig.Verbose,
		userAgent: DefaultConfig.UserAgent,
		logger:    DefaultConfig.Logger,
	}
}

func NewVecchioClientWithConfig(config Config) *VecchioClient {
	if config.Client == nil {
		config.Client = DefaultConfig.Client
	}
	if strings.TrimSpace(config.UserAgent) == "" {
		config.UserAgent = DefaultConfig.UserAgent
	}
	if config.Logger == nil {
		config.Logger = DefaultConfig.Logger
	}
	return &VecchioClient{
		client:    config.Client,
		verbose:   config.Verbose,
		userAgent: config.UserAgent,
		logger:    config.Logger,
	}
}

func (c *VecchioClient) NewThread(thread Thread) (int, error) {
	return c.sendPost(thread)
}

func (c *VecchioClient) PostReply(reply Reply) (int, error) {
	return c.sendPost(reply)
}
