# Health Data Service

Health Data Service is a API that would allows you to upload and view a DICOM file.

The creation of this service was for the pocket health challenge. The requirement was to build a HTTP microservice using Go to do the following:
- Upload and store a DICOM file locally
- Extract and return any DICOM header attributes based on a DICOM Tag as a query parameter
- Convert the file into a PNG for browser-based viewing

For parsing and processing the DICOM file, the following library was used: `https://github.com/suyashkumar/dicom`

Feel free to use the `IM000001` file in `testfiles` directory.

## Install

    go mod vendor
    go build

## Run the service

    go run .

## Run with Docker

    docker build ./
    docker run -d -p 8080:8080 {id}

## REST API
The REST API for a health data service is described below. For easy access, feel free to use the [Postman](https://www.postman.com/) [collection](https://github.com/wongpatrick/health-data-service/blob/main/Health-Data-Service.postman_collection.json). 

## Upload and store a dicom file and returns the id for viewing
### Request
```POST /v1/dicom ```

    curl --location 'http://localhost:8080/v1/dicom' \
    --form 'file=@"/C:/Users/wongp/Documents/Projects/health-data-service/testfiles/IM000001"'

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


## Extract and return attribute for file id based on the tag query parameter
### Request
```GET /v1/dicom/{id}/attribute ```

    curl --location 'http://localhost:8080/v1/dicom/{id}/attribute?tag=(0002,0001)'

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

## Convert file to a png image based on the file id
### Request
```GET /v1/dicom/{id}/image ```

    curl --location -o 'http://localhost:8080/v1/dicom/{id}/image' -o image.png

### Response

    HTTP/1.1 200 OK
    Date: Mon, 03 March 2024 18:36:33 GMT
    Status: 200 OK
    Connection: close
    Content-Type: image/png
    Transfer-Encoding: chunked

<img src="https://github.com/wongpatrick/health-data-service/blob/main/assets/image.png?raw=true" alt="" width="300">

## Next Steps
- Features/Function
    - Bulk upload/query
    - Build out user and access control limiting DICOM file access to their own
    - Build out relational DB for storing user info
    - Implement object storage such as Azure Blob Storage
- Testing/Devops
    - Increase code coverage
    - Add and clean up mocks
    - Add Integration Tests
    - Add E2E Test
    - Add Health Check route
    - Add better logging and monitoring