package client

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/anond0rf/vecchioclient/model"
)

func (c *VecchioClient) sendGetRequest(userAgent string, post model.Post) (*http.Response, error) {
	reqURL := generateGetURL(post)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating GET request to %s: %w", reqURL, err)
	}

	u, err := url.Parse(reqURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse URL from string: %s %w", reqURL, err)
	}

	setGetHeaders(req, userAgent, u)
	return c.client.Do(req)
}

func generateGetURL(post model.Post) string {
	switch p := post.(type) {
	case Reply:
		return fmt.Sprintf("https://%s/%s/res/%d.html", host, p.Board, p.Thread)
	case Thread:
		return fmt.Sprintf("https://%s/%s/", host, p.Board)
	default:
		return ""
	}
}

func setGetHeaders(req *http.Request, userAgent string, u *url.URL) {
	pathSegments := strings.Split(strings.Trim(u.Path, "/"), "/")

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Referer", fmt.Sprintf("https://%s/%s/", u.Host, pathSegments[0]))
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "it-IT,it;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Set("Accept-Encoding", "identity")
}

func (c *VecchioClient) sendPostRequest(postData *bytes.Buffer, contentType, userAgent string, post model.Post) (*http.Response, error) {
	reqURL := fmt.Sprintf("https://%s/post.php", host)
	req, err := http.NewRequest("POST", reqURL, postData)
	if err != nil {
		return nil, fmt.Errorf("creating POST request to %s: %w", reqURL, err)
	}

	u, err := url.Parse(reqURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse URL from string: %s %w", reqURL, err)
	}

	setPostHeaders(req, post, userAgent, u, contentType)
	return c.client.Do(req)
}

func setPostHeaders(req *http.Request, post model.Post, userAgent string, u *url.URL, contentType string) {
	switch p := post.(type) {
	case Reply:
		req.Header.Set("Referer", fmt.Sprintf("https://%s/%s/res/%d.html", u.Host, p.GetBoard(), p.GetThread()))
	case Thread:
		req.Header.Set("Referer", fmt.Sprintf("https://%s/%s/", u.Host, p.GetBoard()))
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Origin", fmt.Sprintf("https://%s", u.Host))
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "it-IT,it;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Set("Accept-Encoding", "identity")
}
