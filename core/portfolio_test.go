package core

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Portfolio", func() {

	var (
		portfolio Portfolio
	)

	BeforeEach(func() {
		portfolio = Portfolio{}
	})

	Describe("Getting started with ginko", func() {
		Context("With 101 test", func() {
			It("should work", func() {
				Expect(1).To(Equal(1))
			})
		})
	})

})