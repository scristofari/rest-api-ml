package main

import (
	"fmt"
	"log"

	"github.com/scristofari/rest-api-ml/ml"
)

func main() {
	err := ml.BuildImageFromDockerfile("./ml/fixture/archive.tar", []string{"api-ml-test"})
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Done")
}
