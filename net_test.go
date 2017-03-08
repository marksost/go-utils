// Tests the net.go file
package goutils

import (
	// Standard lib
	"bytes"
	"net/http"
	"net/url"
	"time"

	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("net.go", func() {
	Describe("`NewRequestConfig` method", func() {
		It("Returns a valid request config struct", func() {
			// Call method
			c := NewRequestConfig()

			// Verify request config was properly created and returned
			Expect(c.Body).To(BeNil())
			Expect(c.Client).To(BeNil())
			Expect(c.Method).To(Equal("GET"))
			Expect(c.Timeout).To(Equal(5))
			Expect(c.URL).To(Equal(""))
		})
	})

	Describe("`GetEmptyPort` method", func() {
		// NOTE: This test is not exact. Some CI runner won't free up ports,
		// which can cause this test to fail there
		Context("When one or more ports are free on a network device", func() {
			It("Returns the port", func() {
				// Call method
				_, err := GetEmptyPort()

				// Verify return values
				Expect(port).To(Not(BeNil()))
			})
		})

		Context("When no ports are free on a network device", func() {
			BeforeEach(func() {
				// Set invalid server host
				ServerHost = "invalid-address"
			})

			It("Returns an error", func() {
				// Call method with invalid server host
				port, err := GetEmptyPort()

				// Verify return values
				Expect(port).To(Equal(0))
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("`GetStatusCodeForRequest` method", func() {
		var (
			// Input for `GetStatusCodeForRequest` input
			input *RequestConfig
		)

		Context("An error occurred when forming the HTTP request", func() {
			BeforeEach(func() {
				// Set input
				input = &RequestConfig{
					Method: "GET",
					URL:    ":",
				}
			})

			It("Returns an error", func() {
				// Call method
				_, err := GetStatusCodeForRequest(input)

				// Verify return value
				Expect(err).To(HaveOccurred())
			})
		})

		Context("An error occurred when forming the HTTP request via a non-mocked client", func() {
			BeforeEach(func() {
				// Set input
				input = &RequestConfig{
					Method:  "GET",
					Timeout: 1,
					URL:     getMockServer("timeout").URL,
				}
			})

			It("Returns an error", func() {
				// Call method
				_, err := GetStatusCodeForRequest(input)

				// Verify return value
				Expect(err).To(HaveOccurred())
			})
		})

		Context("An error occurred when forming the HTTP request via a mocked client", func() {
			BeforeEach(func() {
				// Create test server to mock responses
				server := getMockServer("timeout")

				// Set input
				input = &RequestConfig{
					Client: &http.Client{
						Timeout: time.Duration(1) * time.Second,
						Transport: &http.Transport{
							Proxy: func(req *http.Request) (*url.URL, error) {
								return url.Parse(server.URL)
							},
						},
					},
					Method: "GET",
					URL:    server.URL,
				}
			})

			It("Returns an error", func() {
				// Call method
				_, err := GetStatusCodeForRequest(input)

				// Verify return value
				Expect(err).To(HaveOccurred())
			})
		})

		Context("The server returned a valid response", func() {
			BeforeEach(func() {
				// Create test server to mock responses
				server := getMockServer("default")

				// Set input
				input = &RequestConfig{
					Client: &http.Client{
						Transport: &http.Transport{
							Proxy: func(req *http.Request) (*url.URL, error) {
								return url.Parse(server.URL)
							},
						},
					},
					Method: "GET",
					URL:    server.URL,
				}
			})

			It("Returns no error", func() {
				// Call method
				code, err := GetStatusCodeForRequest(input)

				// Verify return value
				Expect(code).To(Equal(200))
				Expect(err).To(Not(HaveOccurred()))
			})
		})
	})

	Describe("`MakeRequest` method", func() {
		var (
			// Input for `MakeRequest` input
			input *RequestConfig
		)

		Context("An error occurred when forming the HTTP request", func() {
			BeforeEach(func() {
				// Set input
				input = &RequestConfig{
					Method: "GET",
					URL:    ":",
				}
			})

			It("Returns an error", func() {
				// Call method
				_, err := MakeRequest(input)

				// Verify return value
				Expect(err).To(HaveOccurred())
			})
		})

		Context("An error occurred when forming the HTTP request via a non-mocked client", func() {
			BeforeEach(func() {
				// Set input
				input = &RequestConfig{
					Body:        bytes.NewBuffer([]byte("foo")), // Ensures header adding coverage
					ContentType: "application/json",
					Method:      "POST",
					Timeout:     1,
					URL:         getMockServer("timeout").URL, // Test server to ensure client error
				}
			})

			It("Returns an error", func() {
				// Call method
				_, err := MakeRequest(input)

				// Verify return value
				Expect(err).To(HaveOccurred())
			})
		})

		Context("An error occurred when forming the HTTP request via a mocked client", func() {
			BeforeEach(func() {
				// Create test server to mock responses
				server := getMockServer("timeout")

				// Set input
				input = &RequestConfig{
					Body: bytes.NewBuffer([]byte("foo")), // Ensures header adding coverage
					Client: &http.Client{
						Timeout: time.Duration(1) * time.Second,
						Transport: &http.Transport{
							Proxy: func(req *http.Request) (*url.URL, error) {
								return url.Parse(server.URL)
							},
						},
					},
					Method: "PATCH",
					URL:    server.URL,
				}
			})

			It("Returns an error", func() {
				// Call method
				_, err := MakeRequest(input)

				// Verify return value
				Expect(err).To(HaveOccurred())
			})
		})

		Context("The server returned a valid response", func() {
			BeforeEach(func() {
				// Create test server to mock responses
				server := getMockServer("default")

				// Set input
				input = &RequestConfig{
					Client: &http.Client{
						Transport: &http.Transport{
							Proxy: func(req *http.Request) (*url.URL, error) {
								return url.Parse(server.URL)
							},
						},
					},
					Method: "GET",
					URL:    server.URL,
				}
			})

			It("Returns no error", func() {
				// Call method
				res, err := MakeRequest(input)

				// Verify return value
				Expect(err).To(Not(HaveOccurred()))
				Expect(res.StatusCode).To(Equal(200))
			})
		})
	})
})
