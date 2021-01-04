package day_2

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"

)

func Test(t *testing.T) {
	fmt.Printf("test")
}

var _ = Describe("Stuff2", func() {

	Context("initially", func() {

		s := S2{}
		s.Name = "Zoe"

		It("Has to have name", func() {
			Expect(s.Name).Should(Equal("Zoe"))
		})

	})

})
