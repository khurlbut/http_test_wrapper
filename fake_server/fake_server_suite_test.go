package fake_server_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFakeServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FakeServer Suite")
}
