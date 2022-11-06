package sit

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("/createCustomer -- 200", func() {
	It("create customer with no previous ", func() {
		fmt.Println("hello, ginkgo!")
		Expect(true).To(BeTrue())
	})
	It("create customer with no previous history", Label("Smoke", "Regression"), func() {
		fmt.Println("hello, ginkgo!")
		Expect(true).To(BeTrue())
	})
	It("create customer with no previous history", Label("SIT1", "Regression"), func() {
		fmt.Println("hello, ginkgo!")
		Expect(true).To(BeTrue())
	})

})
