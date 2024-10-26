package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/anond0rf/vecchioclient/internal/model"
)

func (c *VecchioClient) sendPost(post model.Post) (int, error) {
	if err := validatePost(post); err != nil {
		return 0, fmt.Errorf("error validating post: %w", err)
	}

	userAgent := getUserAgent()
	if c.userAgent != "" {
		userAgent = c.userAgent
	}

	getResp, err := c.sendGetRequest(userAgent, post)
	if err != nil {
		return 0, fmt.Errorf("error performing GET request: %w", err)
	}
	defer getResp.Body.Close()

	if c.verbose {
		c.logger.Println("GET Response Status Code: ", getResp.StatusCode)
	}

	if getResp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("GET received non-200 response: %d", getResp.StatusCode)
	}

	formFields, err := parseForm(getResp.Body)
	if err != nil {
		return 0, fmt.Errorf("error parsing HTML form: %w", err)
	}

	postData, contentType, err := c.constructPostData(formFields, post)
	if err != nil {
		return 0, fmt.Errorf("error constructing POST data: %w", err)
	}

	postResp, err := c.sendPostRequest(postData, contentType, userAgent, post)
	if err != nil {
		return 0, fmt.Errorf("error performing POST request: %w", err)
	}
	defer postResp.Body.Close()

	if c.verbose {
		c.logger.Println("POST Response Status Code: ", postResp.StatusCode)
	}

	return c.handlePostResponse(postResp)
}

func (c *VecchioClient) handlePostResponse(resp *http.Response) (int, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("unable to read POST response body: %w", err)
	}

	if c.verbose {
		c.logger.Println("POST Response body: ", string(body))
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("POST received non-200 response: %d", resp.StatusCode)
	}

	var sResp model.SuccessResponse
	if err := json.Unmarshal(body, &sResp); err != nil {
		var eResp model.ErrorResponse
		if err := json.Unmarshal(body, &eResp); err != nil {
			return 0, fmt.Errorf("failed to decode POST response: %w", err)
		}
		return 0, fmt.Errorf("POST responded with error: %s", eResp.Error)
	}

	id, err := strconv.Atoi(sResp.ID)
	if err != nil {
		return 0, fmt.Errorf("unable to read id from response body: %w", err)
	}

	return id, nil
}
