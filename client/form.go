package client

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/anond0rf/vecchioclient/model"
)

func parseForm(body io.Reader) (map[string]string, error) {
	formFields := make(map[string]string)
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, fmt.Errorf("error while reading HTML document: %w", err)
	}

	doc.Find(`form[name="post"] input, form[name="post"] select, form[name="post"] textarea`).Each(func(_ int, field *goquery.Selection) {
		if name, exists := field.Attr("name"); exists {
			var value string
			switch goquery.NodeName(field) {
			case "input":
				value, _ = field.Attr("value")
			case "textarea":
				value = field.Text()
			case "select":
				field.Find("option[selected]").Each(func(_ int, option *goquery.Selection) {
					value, _ = option.Attr("value")
				})
			}
			if name != "" {
				formFields[name] = value
			}
		}
	})

	return formFields, nil
}

func (c *VecchioClient) constructPostData(formFields map[string]string, post model.Post) (*bytes.Buffer, string, error) {
	var postData bytes.Buffer
	writer := multipart.NewWriter(&postData)

	defer func() {
		if err := writer.Close(); err != nil {
			c.logger.Println("Error closing multipart writer: ", err)
		}
	}()

	err := writeStandardFormFields(writer, formFields)
	if err != nil {
		return nil, "", err
	}

	email := post.GetEmail()
	if post.GetSage() {
		email = "rabbia"
	}
	if post.GetSpoiler() {
		writer.WriteField("spoiler", "on")
	}

	writer.WriteField("board", strings.TrimSpace(post.GetBoard()))
	writer.WriteField("name", post.GetName())
	writer.WriteField("email", email)
	writer.WriteField("body", post.GetBody())
	writer.WriteField("embed", post.GetEmbed())
	writer.WriteField("password", post.GetPassword())
	writer.WriteField("json_response", "1")

	if post.GetThread() > 0 {
		writer.WriteField("thread", strconv.Itoa(post.GetThread()))
		writer.WriteField("post", "Nuova Risposta")
	} else {
		writer.WriteField("subject", post.GetSubject())
		writer.WriteField("post", "Nuovo Filo")
	}

	if c.verbose {
		c.logger.Println("Form data (excl. files): \n", postData.String())
	}

	err = writeFileFields(writer, post)
	if err != nil {
		return nil, "", fmt.Errorf("error writing files to form: %w", err)
	}

	return &postData, writer.FormDataContentType(), nil
}

func writeStandardFormFields(writer *multipart.Writer, formFields map[string]string) error {
	skipKeys := map[string]struct{}{
		"thread":        {},
		"board":         {},
		"name":          {},
		"email":         {},
		"spoiler":       {},
		"subject":       {},
		"body":          {},
		"embed":         {},
		"password":      {},
		"json_response": {},
		"post":          {},
		"file":          {},
	}

	for key, value := range formFields {
		if _, skip := skipKeys[key]; !skip {
			if err := writer.WriteField(key, value); err != nil {
				return fmt.Errorf("error writing form field %s: %w", key, err)
			}
		}
	}
	return nil
}

func writeFileFields(writer *multipart.Writer, post model.Post) error {
	for i, filePath := range post.GetFiles() {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("error opening file %s: %w", filePath, err)
		}
		defer file.Close()

		fileName := filepath.Base(filePath)
		var fieldName string
		if i == 0 {
			fieldName = "file"
		} else {
			fieldName = fmt.Sprintf("file%d", i+1)
		}
		part, err := writer.CreateFormFile(fieldName, fileName)
		if err != nil {
			return fmt.Errorf("error creating form file for %s: %w", filePath, err)
		}

		if _, err := io.Copy(part, file); err != nil {
			return fmt.Errorf("error copying file %s: %w", filePath, err)
		}
	}
	return nil
}
