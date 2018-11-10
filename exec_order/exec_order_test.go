package exec_order_test

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"

	"fmt"
)

var _ = BeforeSuite(func() {
	fmt.Printf("--- BeforeSuite ---\n")
	TopLevelFunc()
})

func init() {
	fmt.Printf("--- init ---\n")
}

func TopLevelFunc() {
	fmt.Printf("--- TopLevelFunc ---\n")
}

var _ = Describe("ExecOrder", func() {
	fmt.Printf("--- ExecOrderDescribe ---\n")

	BeforeEach(func() {
		fmt.Printf("--- BeforeEach ---\n")
		TopLevelFunc()
	})

	Context("ExecOrder Context", func() {
		fmt.Printf("--- ExecOrderContext---\n")
		It("ExecOrderIt", func() {
			fmt.Printf("--- ExecOrderIt ---\n")
		})
	})
})
