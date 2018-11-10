package fake_server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "kdh.com/fake_server"
)

var targetURL = "http://example.com"

var _ = Describe("Error Scenarios", func() {
	var (
		client     APIClient
		fakeServer FakeServer
		response   chan APIResponse
	)

	BeforeEach(func() {
		response = make(chan APIResponse, 1)
		fakeServer = NewFakeServer()
		client = NewAPIClient(fakeServer)
	})

	Context("Missing Protocol", func() {
		JustBeforeEach(func() {
			client.Get("missing/protocol", response)
		})

		It("should report connection error", func() {
			Ω((<-response).IsError).Should(BeTrue())
		})

		It("should give proper desription of error", func() {
			Ω((<-response).ErrorMessage).Should(Equal("Get missing/protocol: unsupported protocol scheme \"\""))
		})
	})

	Context("Server Unreachable", func() {
		JustBeforeEach(func() {
			client.Get("http://unreachable", response)
		})

		It("should report connection error", func() {
			Ω((<-response).IsError).Should(BeTrue())
		})

		It("should give proper desription of error", func() {
			Ω((<-response).ErrorMessage).Should(Equal("Get http://unreachable: dial tcp: lookup unreachable: no such host"))
		})
	})

	Context("Page Not Found (404)", func() {
		JustBeforeEach(func() {
			client.Get("http://example.com/page/does/not/exist", response)
		})

		It("should return status code 404", func() {
			Ω((<-response).StatusCode).Should(Equal(404))
		})

		It("should return status 404 message", func() {
			Ω((<-response).Status).Should(Equal("404 Not Found"))
		})

		It("should not report connection error", func() {
			Ω((<-response).IsError).Should(BeFalse())
		})

		It("should have an empty connection error message ", func() {
			Ω((<-response).ErrorMessage).Should(BeEmpty())
		})
	})
})

var _ = Describe("Correct Scenarios", func() {
	var client APIClient
	var fakeServer FakeServer
	var response chan APIResponse

	BeforeEach(func() {
		response = make(chan APIResponse, 1)
		fakeServer = NewFakeServer()
		client = NewAPIClient(fakeServer)
		client.Get("http://example.com", response)
	})

	It("should return status code 200", func() {
		Ω((<-response).StatusCode).Should(Equal(200))
	})

	It("should return status message OK", func() {
		Ω((<-response).Status).Should(Equal("200 OK"))
	})

	It("should have data in the body", func() {
		Ω((<-response).Body).ShouldNot(BeEmpty())
	})

	It("should not report connection error", func() {
		Ω((<-response).IsError).Should(BeFalse())
	})

	It("should have an empty connection error message ", func() {
		Ω((<-response).ErrorMessage).Should(BeEmpty())
	})
})
