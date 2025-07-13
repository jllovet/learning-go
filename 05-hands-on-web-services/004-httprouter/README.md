# Setup

```shell
go mod init github.com/jllovet/learning-go/05-hands-on-web-services/004-httprouter
go get github.com/julienschmidt/httprouter
```

# Testing

```shell
curl http://localhost:8000/api/v1/go-version
curl http://localhost:8000/api/v1/show-file/main.go
```