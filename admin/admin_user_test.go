package admin

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Admin User tests", func() {

	Describe("Creating an AdminUser", func() {
		It("should succeed", func() {
			By("Creating AdminUser")
			adminUser := &AdminUser{}
			Expect(adminUser).NotTo(BeNil())
		})
	})

})
