# WIP - rest-api-ml

Create a REST api which will execute some machine learning algorithms.

## Ideas

- [x] Create the webserver in go with the endpoints
- [x] Add the dependencies thanks to dep.
- [x] Use Makefile to create some tasks.
- [ ] Mock the request / response for test purpose.
- [ ] Add a workflow ( status : Upload / Build / Run )
- [ ] After uploading the artifact, launch some goroutines (background) to build / run it and get the result.
- [ ] Store the result of each uploaded object / Mapping : uuid upload - type (build / run) - container id - result / Create an interface to help adding X differents storage types.
- [ ] Stream the build / run tasks output - Event source - Log of the container

## Requirements

    - dep => https://github.com/golang/dep

## Dependencies

    - Gorilla Mux   => https://github.com/gorilla/mux
    - godotenv      => https://github.com/joho/godotenv
    - docker (moby) => https://github.com/joho/godotenv

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

The tests will take a long time with 1GB by image.

 ```Makefile
    make test
    make bench
```
