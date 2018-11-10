package fake_server

type APIClient struct{}

func NewAPIClient(fakeServer FakeServer) APIClient {
	return APIClient{}
}

func (c APIClient) Get(target string, response chan APIResponse) {
	response <- NewAPIResponse()
}
