// Package client provides a client for interacting with vecchiochan.
//
// It includes functionality for creating new threads and  posting replies to existing threads, and allows
// for custom http.Client, logger, logging behavior, and User-Agent header.
package client

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const host = "vecchiochan.com"

// VecchioClient is a client for interacting with vecchiochan.
//
// It allows for posting new threads and replies.
type VecchioClient struct {
	client    *http.Client
	logger    *log.Logger
	userAgent string
	verbose   bool
}

// Config is used to customize the creation of a VecchioClient.
//
// It holds configurations for the underlying HTTP client, logger, user-agent, and logging verbosity.
//
// When Verbose is set to true, HTTP response codes as well as the submitted form and response are logged.
type Config struct {
	Client    *http.Client
	Logger    *log.Logger
	UserAgent string
	Verbose   bool
}

// DefaultConfig provides default values for a VecchioClient.
var DefaultConfig = Config{
	Client:    &http.Client{Timeout: 30 * time.Second},
	Logger:    log.New(os.Stderr, "", log.LstdFlags),
	UserAgent: "",
	Verbose:   false,
}

// NewVecchioClient creates a new VecchioClient with default configuration.
//
// It uses a default HTTP client, logger, and user-agent and only base logging will be shown unless otherwise configured.
//
// See NewVecchioClientWithConfig to pass a custom configuration.
func NewVecchioClient() *VecchioClient {
	return &VecchioClient{
		client:    DefaultConfig.Client,
		logger:    DefaultConfig.Logger,
		userAgent: DefaultConfig.UserAgent,
		verbose:   DefaultConfig.Verbose,
	}
}

// NewVecchioClientWithConfig creates a new VecchioClient using the provided configuration.
//
// If any field in the config is nil or empty, it will fall back to the corresponding default value.
//
// If no custom configuration is needed calling NewVecchioClient instead is enough.
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

// NewThread submits a new thread to vecchiochan.
//
// The thread argument should contain all the necessary information for a new thread.
//
// It returns the id of the newly-created thread or an error if the operation fails.
func (c *VecchioClient) NewThread(thread Thread) (int, error) {
	c.logger.Printf("Submitting new thread... (thread: %+v)\n", thread)
	id, err := c.sendPost(thread)
	if err != nil {
		c.logger.Printf("Unable to post thread %+v. Error: %v\n", thread, err)
	} else {
		c.logger.Printf("Thread posted successfully (id: %d) - %+v\n", id, thread)
	}
	return id, err
}

// PostReply submits a reply to an existing thread on vecchiochan.
//
// The reply argument should contain all the necessary information for a reply.
//
// It returns the id of the submitted reply or an error if the operation fails.
func (c *VecchioClient) PostReply(reply Reply) (int, error) {
	c.logger.Printf("Submitting new reply... (reply: %+v)\n", reply)
	id, err := c.sendPost(reply)
	if err != nil {
		c.logger.Printf("Unable to post reply %+v. Error: %v\n", reply, err)
	} else {
		c.logger.Printf("Reply posted successfully (id: %d) - %+v\n", id, reply)
	}
	return id, err
}
