# User Creation Golang Project!

Create a POST API in Go using any Golang framework gorilla/mux. The API should
accept a JSON payload with the following fields:
● name: string
● pan: string (PAN number)
● mobile: number (mobile number)
● email: string (email ID)

## Build the Docker image

```docker build -t myapp .```

## Run the Docker container

``` docker run -p 8080:8080 myapp```

## Run Without using Dockerfile
You can use below command to run this program locally

```go run /cmd/main.go```

## Test the API  
You can use tools like `curl` or Postman to test the API:

``` curl -X POST http://localhost:8080/user -H "Content-Type: application/json" -d '{"name":"John Doe", "pan":"ABCDE1234F", "mobile":"9876543210", "email":"john@example.com"}'```

## Run the Tests
You can run your tests using the Go test command:

``` go test ./internal/handlers -v```

## Check Coverage for test cases

``` go test ./... -cover ```