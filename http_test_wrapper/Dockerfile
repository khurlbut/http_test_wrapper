FROM golang:1.11.1-alpine
WORKDIR /go
COPY ./*.go src/kdh.com/http_test_wrapper/
COPY ./fake_server/entry.go src/kdh.com/http_test_wrapper/fake_server/entry.go
CMD ["go", "install", "kdh.com/http_test_wrapper/fake_server"]
