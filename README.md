The **Northpole Shop** is a simple RESTful API built using Go. The implementation is inspired by the tutorial [Building a Dockerized RESTful API Application in Go](https://learning-cloud-native-go.github.io/docs/building-a-dockerized-restful-api-application-in-go/).

## Run the Backend
Before running the Docker image for the first time, create an `.app.env` file. You can copy the contents of the `.app.env.template` file and configure the database settings accordingly.

Use `make run` to build and start the Docker container, which includes both the PostgreSQL database and the backend API.

## Test the Backend
`curl localhost:8080`
