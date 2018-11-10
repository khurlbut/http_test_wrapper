package exec_order_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestExecOrder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ExecOrder Suite")
}
