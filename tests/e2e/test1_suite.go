package e2e

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("tests for learning ginkgo", func() {
	It("should run the tests", func() {
		fmt.Println("hello, ginkgo!")
		Expect(true).To(BeTrue())
	})
})
