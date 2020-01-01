package userdata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Data User tests", func() {

	Describe("Creating a UdUser", func() {
		It("should succeed", func() {
			By("Creating UdUser")
			udUser := &UdUser{}
			Expect(udUser).NotTo(BeNil())
		})
	})

})
