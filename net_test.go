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
				port, err := GetEmptyPort("")

				// Verify return values
				Expect(port).To(Not(Equal(0)))
				Expect(err).To(Not(HaveOccurred()))
			})
		})

		Context("When no ports are free on a network device", func() {
			It("Returns an error", func() {
				// Call method with invalid server host
				port, err := GetEmptyPort("invalid-address")

				// Verify return values
				Expect(port).To(Equal(0))
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
