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
	logger    *log.Logger
	userAgent string
	verbose   bool
}

type Config struct {
	Client    *http.Client
	Logger    *log.Logger
	UserAgent string
	Verbose   bool
}

var DefaultConfig = Config{
	Client:    &http.Client{Timeout: 30 * time.Second},
	Logger:    log.New(os.Stderr, "", log.LstdFlags),
	UserAgent: "",
	Verbose:   false,
}

func NewVecchioClient() *VecchioClient {
	return &VecchioClient{
		client:    DefaultConfig.Client,
		logger:    DefaultConfig.Logger,
		userAgent: DefaultConfig.UserAgent,
		verbose:   DefaultConfig.Verbose,
	}
}

func NewVecchioClientWithConfig(config Config) *VecchioClient {
	if config.Client == nil {
		config.Client = DefaultConfig.Client
	}
	if config.Logger == nil {
		config.Logger = DefaultConfig.Logger
	}
	if strings.TrimSpace(config.UserAgent) == "" {
		config.UserAgent = DefaultConfig.UserAgent
	}
	return &VecchioClient{
		client:    config.Client,
		logger:    config.Logger,
		userAgent: config.UserAgent,
		verbose:   config.Verbose,
	}
}

func (c *VecchioClient) NewThread(thread Thread) (int, error) {
	c.logger.Printf("Submitting new thread... (thread: %+v)\n", thread)
	return c.sendPost(thread)
}

func (c *VecchioClient) PostReply(reply Reply) (int, error) {
	c.logger.Printf("Submitting new reply... (reply: %+v)\n", reply)
	return c.sendPost(reply)
}
