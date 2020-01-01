package dailybonus

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Daily Bonus Config tests", func() {

	Describe("Creating a DbConfig", func() {
		It("should succeed", func() {
			By("Creating DbConfig")
			dbConfig := &DbConfig{}
			Expect(dbConfig).NotTo(BeNil())
		})
	})

})
