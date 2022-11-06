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
	It("create customer with no previous history", func() {
		fmt.Println("hello, ginkgo!")
		Expect(true).To(BeTrue())
	})
	It("create customer with no previous history", func() {
		fmt.Println("hello, ginkgo!")
		Expect(true).To(BeTrue())
	})

})

var _ = Describe("/createCustomer -- 400", func() {
	It("create customer with missing history", func() {
		fmt.Println("hello, ginkgo!")
		Expect(true).To(BeTrue())
	})
	It("create customer with no previous history", func() {
		fmt.Println("hello, ginkgo!")
		Expect(true).To(BeTrue())
	})
	It("create customer with no previous history", func() {
		fmt.Println("hello, ginkgo!")
		Expect(true).To(BeTrue())
	})

})
