package dailybonus

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Daily Bonus User tests", func() {

	Describe("Creating a DbUser", func() {
		It("should succeed", func() {
			By("Creating DbUser")
			dbUser := &DbUser{}
			Expect(dbUser).NotTo(BeNil())
		})
	})

})
