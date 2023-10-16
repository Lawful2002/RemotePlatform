package main

import (
	"docker"
	"fmt"
)

func main() {
	fmt.Println(docker.StopDockerContainer("396299200-1298498081"))
}
