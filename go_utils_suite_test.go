// Test suite setup for the go-utils package
package goutils

import (
	// Standard lib
	"io/ioutil"
	"testing"

	// Third-party
	log "github.com/Sirupsen/logrus"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type (
	// Struct representing SliceContains input data
	SliceContainsTestData struct {
		Needle   string
		Haystack []string
	}
)

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
