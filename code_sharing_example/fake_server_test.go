package fake_server_test

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "kdh.com/code_sharing_example"
)

var _ = Describe("FakeServer API Client Test Suite 1", func() {
	var client APIClient
	var fakeServer FakeServer
	var response chan APIResponse

	BeforeEach(func() {
		response = make(chan APIResponse, 1)
		fakeServer = NewFakeServer()
		client = NewAPIClient(fakeServer)
		client.Get("some/endpoint", response)
	})

	Describe("failure modes", func() {
		AssertFailedBehavior := func() {
			It("should not include JSON in the response", func() {
				Ω((<-response).JSON).Should(BeZero())
			})

			It("should not report success", func() {
				Ω((<-response).Success).Should(BeFalse())
			})
		}

		Context("when the server does not return a 200", func() {
			BeforeEach(func() {
				fakeServer.Respond(404)
			})

			AssertFailedBehavior()
		})

		Context("when the server returns unparseable JSON", func() {
			BeforeEach(func() {
				fakeServer.Succeed("{I'm not JSON!")
			})

			AssertFailedBehavior()
		})

		Context("when the request errors", func() {
			BeforeEach(func() {
				fakeServer.Error(errors.New("oops!"))
			})

			AssertFailedBehavior()
		})

	})
})

var _ = Describe("FakeServer API Client Test Suite 2", func() {
	var client APIClient
	var fakeServer FakeServer
	var response chan APIResponse

	BeforeEach(func() {
		response = make(chan APIResponse, 1)
		fakeServer = NewFakeServer()
		client = NewAPIClient(fakeServer)
		client.Get("/some/endpoint", response)
	})

	Describe("failure modes", func() {
		AssertNoJSONInResponse := func() func() {
			return func() {
				Ω((<-response).JSON).Should(BeZero())
			}
		}

		AssertDoesNotReportSuccess := func() func() {
			return func() {
				Ω((<-response).Success).Should(BeFalse())
			}
		}
		Context("when the server does not return a 200", func() {
			BeforeEach(func() {
				fakeServer.Respond(404)
			})

			It("should not include JSON in the response", AssertNoJSONInResponse())
			It("should not report success", AssertDoesNotReportSuccess())
		})

		Context("when the server ruturns uparseable JSON", func() {
			BeforeEach(func() {
				fakeServer.Succeed("{I'm not JSON!")
			})

			It("should not include JSON in the response", AssertNoJSONInResponse())
			It("should not report success", AssertDoesNotReportSuccess())
		})

		Context("when the request errors", func() {
			BeforeEach(func() {
				fakeServer.Error(errors.New("oops!"))
			})

			It("should not include JSON in the response", AssertNoJSONInResponse())
			It("should not report success", AssertDoesNotReportSuccess())
		})
	})
})
