package main

import (
	"docker"
	"fmt"
)

func main() {
	fmt.Println(docker.RunDockerfile("./Dockerfiles", "test.Dockerfile"))
}
