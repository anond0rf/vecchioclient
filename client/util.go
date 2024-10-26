package client

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/anond0rf/vecchioclient/internal/model"
)

func getUserAgent() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	uaList := []string{
		"Mozilla/5.0 (Windows NT %d.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT %d.0; Win64; x64; rv:130.0) Gecko/20100101 Firefox/%d.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X %d_%d_0) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/%d.1 Safari/605.1.15"}

	idx := r.Intn(3)

	if idx < 2 {
		osVer := r.Intn(2) + 10
		brVer := r.Intn(31) + 100
		return fmt.Sprintf(uaList[idx], osVer, brVer)
	}

	osMajVer := r.Intn(2) + 12
	osMinVer := r.Intn(10)
	brVer := r.Intn(4) + 15

	return fmt.Sprintf(uaList[idx], osMajVer, osMinVer, brVer)
}

func validatePost(post model.Post) error {
	if strings.TrimSpace(post.GetBoard()) == "" {
		return fmt.Errorf("board cannot be empty")
	}
	if p, ok := post.(Reply); ok && p.Thread < 1 {
		return fmt.Errorf("thread must be >= 1")
	}
	return nil
}
