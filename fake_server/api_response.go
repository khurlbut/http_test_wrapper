package fake_server

type APIResponse struct {
	Status       string
	StatusCode   int
	Body         string
	IsError      bool
	ErrorMessage string
}

func NewAPIResponse(status string, statusCode int, body string) APIResponse {
	return APIResponse{
		Status:       status,
		StatusCode:   statusCode,
		Body:         body,
		IsError:      false,
		ErrorMessage: "",
	}
}

func NewAPIErrorResponse(errorMessage string) APIResponse {
	return APIResponse{
		Status:       "",
		StatusCode:   -1,
		Body:         "",
		IsError:      true,
		ErrorMessage: errorMessage,
	}
}
