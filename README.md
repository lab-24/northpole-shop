The **Northpole Shop** is a simple RESTful API built using Go.
The implementation is inspired by the tutorial:  
[Building a Dockerized RESTful API Application in Go](https://learning-cloud-native-go.github.io/docs/building-a-dockerized-restful-api-application-in-go/).

## Run the Backend
Before running the Docker image for the first time, create an `.app.env` file.
You can copy the contents of the `.app.env.template` file and configure the
database settings accordingly.

Use `make run` to build and start the Docker container, which includes both
the PostgreSQL database and the backend API.

## Test Endpoints
### List devices

* Get all devices:
  `curl -H "Accept: application/json" localhost:8080/api/devices | jq`
* Get all devices from location_id `e7f1f3c0-0b6b-11ec-82a8-0242ac130008`
  (Location garage):  
  `curl -H "Accept: application/json" localhost:8080/api/devices?location_id=e7f1f3c0-0b6b-11ec-82a8-0242ac130008 | jq`
* Get all devices before GMT `Wed Dec 18 2024 00:16:51 GMT+0000`
  (Unix Timestamp: `1734481011`)  
  `curl -H "Accept: application/json" localhost:8080/api/devices?end_time=1734481011 | jq`

## Endpoints
| Name         | HTTP Method | Route          | Query Parameter                 |
|--------------|-------------|----------------|---------------------------------|
| List Devices | GET         | /api/devices   | start_date end_date location_id |

💡 [OpenApi Documentation](docs/openapi/index.html)
