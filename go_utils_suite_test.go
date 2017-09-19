// Test suite setup for the go-utils package
package goutils

import (
	// Standard lib
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

type (
	// Struct representing IntSlice2StringSlice input data
	IntSlice2StringSliceTestData struct {
		Input  []int
		Output []string
	}
	// Struct representing SliceContains input data
	SliceContainsTestData struct {
		Needle   string
		Haystack []string
	}
)

// getMockServer returns a httptest server with the desired handler function
// based on the key passed in
func getMockServer(key string) *httptest.Server {
	var handler http.Handler

	switch key {
	case "bad-request":
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Write headers and body
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, `{"code":400}`)
		})
	case "timeout":
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			select {} // NOTE: Allows timeout error to occur within client.Do calls
		})
	default:
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Write headers and body
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, `{}`)
		})
	}

	return httptest.NewServer(handler)
}

// Tests the go-utils package
func TestConfig(t *testing.T) {
	// Register gomega fail handler
	RegisterFailHandler(Fail)

	// Have go's testing package run package specs
	RunSpecs(t, "go-utils suite")
}

func init() {
	// Set logger output so as not to log during tests
	log.SetOutput(ioutil.Discard)
}
