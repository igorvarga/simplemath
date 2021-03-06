[![<igorvarga>](https://circleci.com/gh/igorvarga/simplemath.svg?style=svg)](https://circleci.com/gh/igorvarga/simplemath)

# Simple Math App
Simple web application in Go which accepts math problems via the URL and returns the response in JSON.

## How to run it
Before proceeding to the next steps, make sure you have Go [installed](https://golang.org/doc/install) and GOPATH 
environment variable set.

### Get the project
```shell script
go get https://github.com/igorvarga/simplemath
```

### Configure Server parameters (Optional step)
Server can be configured using environment variables listed below. In case that environment variables are missing the 
default values will be used.

Variable | Description
------------ | -------------
SM_ADDR | Server address string in addreess:port format, default value: ":8080"
SM_CACHE_EXPIRATION | Cache item expiry in seconds, default value: 60
SM_CACHE_SWEEPINTERVAL | Expired cache items sweep interval in seconds, default value 5

***Windows example***
```batch
SET SM_ADDR=":9090"
SET SM_CACHE_EXPIRATION=30
SET SM_CACHE_SWEEPINTERVAL=5
```

***Linux example***

Substitute SET with EXPORT keyword

### Run
```batch
cd %GOPATH%\src\github.com\igorvarga\simplemath
go run main.go
```

## Build
```batch
cd %GOPATH%\src\github.com\igorvarga\simplemath
go build .
```

## Try it
```shell-script
curl "http://localhost:8080/add?x=2&y=7"
```

## Running tests
```batch
cd %GOPATH%\src\github.com\igorvarga\simplemath
go test ./...
```

## Deploying to cloud
Connect CircleCI for building and deploying steps to any cloud provider supported by CircleCI.

## TODO
- [ ] Use .env for configuration
- [ ] Init project as go module
- [ ] Add docker (dev/docker branch) 
- [ ] Add client cache HTTP headers
- [ ] Introduce context cancellation to all layers
- [ ] Add leaky bucket feature to cache, map memory limit
- [ ] Explore RWMutex
- [ ] Refactor server and cache structs and interfaces to be more specific about their functionality
- [ ] Explore sync.Map
- [ ] Benchmark, sync race and memory profiling tests
- [ ] Write better godoc and tests
- [ ] REST API versioning