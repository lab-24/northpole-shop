The **Northpole Shop** is simple RESTful API built using Go.
The implementation is inspired by the tutorial:  
[Building a Dockerized RESTful API Application in Go](https://learning-cloud-native-go.github.io/docs/building-a-dockerized-restful-api-application-in-go/).

## Run the Backend

Before running the Docker image the first time, create an `.app.env` file.
You can copy the contents of the `.app.env.template` file and configure the
database settings accordingly.

Use `make run` to build and start the Docker container, which includes both
the PostgreSQL database and the backend API.

## Endpoints

Only API Endpoints for devices are implemented in this demo:
| Name          | HTTP Method | Route             | Query Parameter                 |
|---------------|-------------|-------------------|---------------------------------|
| List Devices  | GET         | /api/devices      | start_date end_date location_id |
| Create Device | POST        | /api/devices      |                                 |
| Read Device   | GET         | /api/devices/{id} |                                 |
| Update Device | PUT         | /api/devices/{id} |                                 |
| Delete Device | DELETE      | /api/devices/{id} |                                 |

💡 [OpenApi Documentation](/docs/openapi/)

## Test Endpoints

### JWT Authorization

The API is secured using JWT Authorization. Currently, there are no endpoints
or user tables implemented. For testing purposes, set `JWT_DEBUG=true` in the
`.app.env` file. This will generate a valid token during the application startup.
Export the token and include it in your requests.

```
export TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjN9.V6upMbO3L7r2AZqj36dOronI0hONPCAmvO1QB2JXUX0
```

![Get your JWT test token](./docs/startup_jwt_token.png "Startup Application")

### List devices

* Get all devices:  

```
  curl -H "Accept: application/json" -H "Authorization: Bearer ${TOKEN}" localhost:8080/api/devices
```

* Get all devices from location_id `e7f1f3c0-0b6b-11ec-82a8-0242ac130008`
  (Location garage):  

```
  curl -H "Accept: application/json" -H "Authorization: Bearer ${TOKEN}" localhost:8080/api/devices?location_id=e7f1f3c0-0b6b-11ec-82a8-0242ac130008
```

* Get all devices before GMT `Wed Dec 18 2024 00:16:51 GMT+0000`
  (Unix Timestamp: `1734481011`)  

```
  curl -H "Accept: application/json" -H "Authorization: Bearer ${TOKEN}" localhost:8080/api/devices?end_time=1734481011
```

### CRUD Device

* Create a device:  

```
  curl -H "Accept: application/json" -H "Authorization: Bearer ${TOKEN}" -X POST \
  -d '{"name":"Rudy the red nose","serial_number":"RUDY","location_id":"e7f1f3c0-0b6b-11ec-82a8-0242ac130008","device_type_id":"e7f1f3c0-0b6b-11ec-82a8-0242ac130006"}' \
  localhost:8080/api/devices
```

* Update a device:  

```
  curl -H "Accept: application/json" -H "Authorization: Bearer ${TOKEN}" -X PUT \
  -d '{"name":"Rudy the red nose","serial_number":"RUDY","location_id":"e7f1f3c0-0b6b-11ec-82a8-0242ac130008","device_type_id":"e7f1f3c0-0b6b-11ec-82a8-0242ac130006"}' \
  localhost:8080/api/devices/e7f1f3c0-0b6b-11ec-82a8-0242ac130012 
```

* Delete a device:  

```
  curl -H "Accept: application/json" -H "Authorization: Bearer ${TOKEN}" \
  -X DELETE localhost:8080/api/devices/e7f1f3c0-0b6b-11ec-82a8-0242ac130012
```

* Read a device:  

```
  curl -H "Accept: application/json" -H "Authorization: Bearer ${TOKEN}" \
  -X GET localhost:8080/api/devices/e7f1f3c0-0b6b-11ec-82a8-0242ac130012
```

## Database Schema

![Database schema](./docs/db_schema.png "Database schema")
