package book_test

import (
	"testing"

	"fmt"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
)

func TestBook(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Book Suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("Executing BeforeSuite - this code is executed before anything else in the entire test suite.")
	fmt.Println("If you need a unique param (like a port number) for each parallel execution path")
	fmt.Println("you can use the ParallelNode parameter (import github.com/onsi/ginkgo/config) like this:")
	fmt.Printf("Parallel Node: %d", config.GinkgoConfig.ParallelNode)
})

var _ = AfterSuite(func() {
	fmt.Println("Executing AfterSuite - this code is executed after all of the tests. This code will execute even if you <cntl-C> out of the tests!")
})
