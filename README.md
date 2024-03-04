# Health Data Service

Pocket Health Challenge
TODO: Add details about the service

## Install

    go mod vendor
    go build

## Run the service

    go run .

## Run with Docker

    docker build ./
    docker run -d -p 8080:8080 {id}

## REST API
The REST API for a health data service is described below.

## Post upload dicom file
### Request
```POST /v1/dicom/upload ```

    curl 'http://localhost:8080/v1/dicom/upload'

### Response

    HTTP/1.1 200 OK
    Date: Mon, 03 March 2024 18:36:33 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 287

    [
        {
            "id": 2,
            "title": "Ace the first round",
            "description": "go through the first rounds of interview",
            "status": "A",
            "userId": 1,
            "createdAt": "2023-09-18T18:58:29.063666Z",
            "modifiedAt": "2023-09-18T18:58:29.063666Z"
        }
    ]
