# Setup

Following along with course at [Complete Backend API in Golang (JWT, MySQL & Tests)](https://www.youtube.com/watch?v=7VLmLOiQ3ck&ab_channel=Tiago)

```shell
go mod init github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom
go get github.com/gorilla/mux
go get github.com/go-sql-driver/mysql
go get github.com/joho/godotenv
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
