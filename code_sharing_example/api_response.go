package fake_server

type APIResponse struct {
	JSON    string
	Success bool
}

func NewAPIResponse() APIResponse {
	return APIResponse{
		JSON:    "",
		Success: false,
	}
}
