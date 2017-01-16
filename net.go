// Package goutils contains a collection of useful Golang utility methods and libraries
package goutils

import (
	// Standard lib
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"
	"time"
)

type (
	// RequestConfig contains a set of configuration settings
	// to be used with the methods that make HTTP requests
	RequestConfig struct {
		Body    io.Reader
		Client  *http.Client
		Method  string
		Timeout int
		URL     string
	}
)

var (
	// ServerHost is the host to run the Redis server on
	// NOTE: Public variable to allow package authors the ability
	// to change this before starting the Redis server
	ServerHost = "localhost"
)

// NewRequestConfig returns a RequestConfig struct with
// default settings set for each of it's properties
func NewRequestConfig() *RequestConfig {
	return &RequestConfig{
		Body:    nil,
		Client:  nil,
		Method:  "GET",
		Timeout: 1,
		URL:     "",
	}
}

// GetEmptyPort returns a number to be used as a new server's port
// NOTE: Uses tcp to allow the kernel to give an open port
func GetEmptyPort() (int, error) {
	// Create regex for extracting port
	r, _ := regexp.Compile("\\d+$")

	// NOTE: Uses "port" 0 to allow the kernal to chose a port for itself
	if l, err := net.Listen("tcp", fmt.Sprintf("%s:0", ServerHost)); err == nil {
		// Close listener
		defer l.Close()

		// Use regex to extract port
		port := r.FindString(l.Addr().String())

		if len(port) != 0 {
			return String2Int(port), nil
		}
	}

	return 0, fmt.Errorf("No random ports were found")
}

// GetStatusCodeForRequest attempts to make an HTTP request against a URL
// with a given HTTP method (ex: GET) and returns it's status code if successful,
// an error otherwise
func GetStatusCodeForRequest(c *RequestConfig) (int, error) {
	// Make new request
	req, err := http.NewRequest(c.Method, c.URL, c.Body)
	if err != nil {
		return 0, err
	}

	// Create client
	// NOTE: Allow a client to be passed in (like during testing)
	if c.Client == nil {
		c.Client = &http.Client{
			Timeout: time.Duration(c.Timeout) * time.Second,
		}
	}

	// Send request
	res, err := c.Client.Do(req)
	if err != nil {
		return 0, err
	}

	return res.StatusCode, nil
}
