# Setup

Following along with course at [Complete Backend API in Golang (JWT, MySQL & Tests)](https://www.youtube.com/watch?v=7VLmLOiQ3ck&ab_channel=Tiago)

```shell
go mod init github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom
go get github.com/gorilla/mux
go get github.com/go-sql-driver/mysql
go get github.com/joho/godotenv
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
mkdir db
touch db/db.go
mkdir config
touch config/env.go
touch .env
touch docker-compose.yaml
```

# Running

To run the API, we can use the Makefile we have built.

```shell
make run
```

To set up the database that the API is going to talk to in the backend, we will need to spin up a few Docker containers. The services in those containers are defined in `docker-compose.yaml`. For this example scenario, they will be reading from the same location as the API, namely from `.env`. To start those services, run:

```shell
docker compose up --build
```
