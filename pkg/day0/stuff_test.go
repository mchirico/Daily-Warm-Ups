package day0_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mchirico/Daily-Warm-Ups/pkg/day0"
)

var _ = Describe("Stuff", func() {
	Context("initially", func() {

		s := day0.S{}
		s.Name = "Zoe"

		It("Has to have name", func() {
			Expect(s.Name).Should(Equal("Zoe"))
		})

	})
})
