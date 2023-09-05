package main

import (
	"fmt"
	"systemcaller"
)

func main() {
	fmt.Println(systemcaller.RunSystemCommand("docker run --help"))
}
