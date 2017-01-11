// A collection of useful Golang utility methods and libraries
package goutils

import (
	// Standard lib
	"fmt"
	"net"
	"regexp"
)

var (
	// The host to run the Redis server on
	// NOTE: Public variable to allow package authors the ability
	// to change this before starting the Redis server
	ServerHost string = "localhost"
)

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
