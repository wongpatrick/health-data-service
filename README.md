# Health Data Service

Pocket Health Challenge
TODO: 
- Add details about the service
- Figure out naming convention
- Clean test files
- Talk about your decisions like database etc

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

## Upload dicom file
### Request
```POST /v1/dicom ```

    curl 'http://localhost:8080/v1/dicom'

### Response

    HTTP/1.1 201 OK
    Date: Mon, 03 March 2024 18:36:33 GMT
    Status: 201 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 45

    {
        "id": "6073f0cd-350c-4a1b-82c3-0ee708563fed"
    }


## Get attribute for file id
### Request
```GET /v1/dicom/{id}/attribute ```

    curl --location 'http://localhost:8080/v1/dicom/{id}/attribute?tag(0002,0001)'

### Response

    HTTP/1.1 200 OK
    Date: Mon, 03 March 2024 18:36:33 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 82

    {
        "tag": {
            "Group": 2,
            "Element": 1
        },
        "VR": 1,
        "rawVR": "OB",
        "valueLength": 2,
        "value": "AAE="
    }

## Get Image for file id
### Request
```GET /v1/dicom/{id}/image ```

    curl --location 'http://localhost:8080/v1/dicom/{id}/image'

### Response

    HTTP/1.1 200 OK
    Date: Mon, 03 March 2024 18:36:33 GMT
    Status: 200 OK
    Connection: close
    Content-Type: image/png
    Transfer-Encoding: chunked

<img src="https://github.com/wongpatrick/health-data-service/blob/main/testfiles/testimage.png?raw=true" alt="" width="300">
