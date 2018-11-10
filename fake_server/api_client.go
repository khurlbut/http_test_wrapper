package fake_server

import (
	"io/ioutil"
	"net/http"
)

type APIClient struct{}

func NewAPIClient(fakeServer FakeServer) APIClient {
	return APIClient{}
}

func (c APIClient) Get(target string, response chan APIResponse) {
	resp, err := http.Get(target)

	if err != nil {
		response <- NewAPIErrorResponse(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	status := resp.Status
	statusCode := resp.StatusCode

	response <- NewAPIResponse(status, statusCode, string(body))
}
