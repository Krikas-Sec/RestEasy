package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type RequestConfig struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

func main() {
	// CLI flags
	method := flag.String("method", "GET", "HTTP method (GET, POST, PUT, DELETE)")
	url := flag.String("url", "", "API URL to request")
	headers := flag.String("headers", "", "HTTP headers as key:value,key:value")
	body := flag.String("body", "", "Request body")
	save := flag.String("save", "", "Save request configuration to file")
	load := flag.String("load", "", "Load request configuration from file")
	flag.Parse()

	if *load != "" {
		config, err := loadRequestConfig(*load)
		if err != nil {
			fmt.Println("Error loading request configuration:", err)
			os.Exit(1)
		}
		executeRequest(config)
		return
	}

	if *url == "" {
		fmt.Println("Error: URL is required. Use -url to specify the endpoint.")
		os.Exit(1)
	}

	headersMap := parseHeaders(*headers)
	config := RequestConfig{
		Method:  *method,
		URL:     *url,
		Headers: headersMap,
		Body:    *body,
	}

	if *save != "" {
		err := saveRequestConfig(config, *save)
		if err != nil {
			fmt.Println("Error saving request configuration:", err)
			os.Exit(1)
		}
		fmt.Printf("Request configuration saved to %s\n", *save)
		return
	}

	executeRequest(config)
}

func parseHeaders(headerStr string) map[string]string {
	headers := make(map[string]string)
	if headerStr == "" {
		return headers
	}
	pairs := strings.Split(headerStr, ",")
	for _, pair := range pairs {
		kv := strings.SplitN(pair, ":", 2)
		if len(kv) == 2 {
			headers[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}
	return headers
}

func executeRequest(config RequestConfig) {
	client := &http.Client{}
	var reqBody *bytes.Buffer
	if config.Body != "" {
		reqBody = bytes.NewBuffer([]byte(config.Body))
	} else {
		reqBody = &bytes.Buffer{}
	}

	req, err := http.NewRequest(config.Method, config.URL, reqBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	for key, value := range config.Headers {
		req.Header.Add(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	fmt.Printf("Status: %s\n", res.Status)
	fmt.Printf("Headers: %v\n", res.Header)
	fmt.Printf("Body: %s\n", string(body))
}

func saveRequestConfig(config RequestConfig, filename string) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func loadRequestConfig(filename string) (RequestConfig, error) {
	var config RequestConfig
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}
