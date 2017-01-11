// A collection of useful Golang utility methods and libraries
package goutils

import (
	// Standard lib
	"fmt"
	"net"
	"regexp"
)

// GetEmptyPort returns a number to be used as a new server's port
// NOTE: Uses tcp to allow the kernel to give an open port
func GetEmptyPort(host string) (int, error) {
	// Set default host if needed
	if host == "" {
		host = "localhost"
	}

	// Create regex for extracting port
	r, _ := regexp.Compile("\\d+$")

	// NOTE: Uses "port" 0 to allow the kernal to chose a port for itself
	if l, err := net.Listen("tcp", fmt.Sprintf("%s:0", host)); err == nil {
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
