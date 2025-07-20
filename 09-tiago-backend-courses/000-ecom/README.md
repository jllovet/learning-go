# Setup

Following along with course at [Complete Backend API in Golang (JWT, MySQL & Tests)](https://www.youtube.com/watch?v=7VLmLOiQ3ck&ab_channel=Tiago)

```shell
go mod init github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom
go get github.com/gorilla/mux
```

```shell
mkdir cmd
mkdir cmd/api
mkdir cmd/migrate
touch cmd/api/api.go
touch cmd/main.go
mkdir services
mkdir services/user
touch services/user/routes.go
touch Makefile
```

# Running

```shell
make run
```
