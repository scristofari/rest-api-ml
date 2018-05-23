# WIP - rest-api-ml

Create a REST api which will execute some machine learning algorithms.

## Ideas

- [ ] After uploading the image or file, launch a goroutine to run it and get the result.
- [ ] Save in a textfile the result of each uploaded object. 
- [ ] Mock the request / response for test purpose.

## Requirements

    - dep => https://github.com/golang/dep

## Dependencies:

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
