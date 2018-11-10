package fake_server

type FakeServer struct{}

func NewFakeServer() FakeServer {
	return FakeServer{}
}

func (fakeServer FakeServer) Respond(status int16) {}

func (fakeServer FakeServer) Succeed(s string) {}

func (fakeServer FakeServer) Error(e error) {}
