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

### JWT Authorization

The API has been protected with a JWT Authorization. There are no endpoints and
user tables implemented yet. For testing purposes, please set `JWT_DEBUG=true`
in `.app.env` file and you will get a valid Token during startup of the
application. Export this token an use it in your requests.  
`export TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjN9.V6upMbO3L7r2AZqj36dOronI0hONPCAmvO1QB2JXUX0`

![Get your JWT test token](./docs/startup_jwt_token.png "Startup Application")

### List devices

* Get all devices:  
  `curl -H "Accept: application/json" -H "Authorization: Bearer ${TOKEN}" localhost:8080/api/devices`
* Get all devices from location_id `e7f1f3c0-0b6b-11ec-82a8-0242ac130008`
  (Location garage):  
  `curl -H "Accept: application/json" -H "Authorization: Bearer ${TOKEN}" localhost:8080/api/devices?location_id=e7f1f3c0-0b6b-11ec-82a8-0242ac130008`
* Get all devices before GMT `Wed Dec 18 2024 00:16:51 GMT+0000`
  (Unix Timestamp: `1734481011`)  
  `curl -H "Accept: application/json" -H "Authorization: Bearer ${TOKEN}" localhost:8080/api/devices?end_time=1734481011`

## Endpoints
| Name         | HTTP Method | Route          | Query Parameter                 |
|--------------|-------------|----------------|---------------------------------|
| List Devices | GET         | /api/devices   | start_date end_date location_id |

💡 [OpenApi Documentation](/docs/openapi/)
