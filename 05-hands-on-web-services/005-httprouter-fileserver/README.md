# Setup

```shell
go mod init github.com/jllovet/learning-go/05-hands-on-web-services/005-httprouter-fileserver
go get github.com/julienschmidt/httprouter
```

# Testing

```shell
curl http://localhost:8000/static/greek.txt
curl http://localhost:8000/static/latin.txt
```