package userdata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Data Config tests", func() {

	Describe("Creating a UdConfig", func() {
		It("should succeed", func() {
			By("Creating UdConfig")
			udConfig := &UdConfig{}
			Expect(udConfig).NotTo(BeNil())
		})
	})

})
