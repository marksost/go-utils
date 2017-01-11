// Tests the net.go file
package goutils

import (
	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("net.go", func() {
	Describe("`GetEmptyPort` method", func() {
		Context("When one or more ports are free on a network device", func() {
			It("Returns the port", func() {
				// Call method
				_, err := GetEmptyPort()

				// Verify return values
				Expect(err).To(Not(HaveOccurred()))
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
})
