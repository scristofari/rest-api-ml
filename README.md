[![Go Report Card](https://goreportcard.com/badge/github.com/scristofari/rest-api-ml)](https://goreportcard.com/report/github.com/scristofari/rest-api-ml)

[![GoDoc](https://godoc.org/github.com/scristofari/rest-api-ml/box?status.svg)](https://godoc.org/github.com/scristofari/rest-api-ml/box)



# WIP - rest-api-ml

Create a REST api which will execute some machine learning algorithms.

## WIKI:
   [WIKI](https://github.com/scristofari/rest-api-ml/wiki)

## Ideas

- [x] Create the webserver in go with the endpoints
- [x] Add the dependencies thanks to dep.
- [x] Use Makefile to create some tasks.
- [ ] Mock the request / response for test purpose.
- [ ] Add a workflow ( status : Upload / Build / Run )
- [ ] After uploading the artifact, launch some goroutines (background) to build / run it and get the result. Mount a volume for the run task to get the score.json.
- [ ] Store the result of each uploaded object / Mapping : uuid upload - type (build / run) - container id - result / Create an interface to help adding X differents storage types.
- [ ] Stream the build / run tasks output - Event source - Log of the container
- [ ] Pull an image from the hub, just need to run.
- [ ] Add a absolute path for the artifact.  
- [ ] The image's tag is the artifact's hash
- [ ] Grpc / Protobuffer / Grpc Http
- [ ] Endpoint with a version, ex: /v1/upload, /v1/aritifact/uuid
- [ ] Algo: Do not create a json file, ouptut the result in the standart output - No need to mount or copy the file.


## Requirements

    - dep => https://github.com/golang/dep

## Dependencies

    - Gorilla Mux   => https://github.com/gorilla/mux
    - godotenv      => https://github.com/joho/godotenv
    - docker (moby) => https://github.com/moby/moby/tree/master/client

## Installation

We use [dep](https://github.com/golang/dep) to install the dependencies.

 ```Makefile
    make install
 ```

Build and save the docker image in the 'ml' directory.
This image is the docker + algo.

 ```Makefile
    make build
    make save
 ```


## Run the project

 ```Makefile
    make run
 ```

## Test and Benchmark

 ```Makefile
    make test
    make bench
```
