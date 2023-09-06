package systemcaller

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func RunSystemCommand(command string) (string, error) {
	// Input: command string -> the command to be executed
	// Output: (string, error) -> the output of the run command and error(if any)
	log.SetPrefix("command: ")

	if len(command) == 0 {
		return "", errors.New("empty command")
	}

	os := runtime.GOOS
	switch os {
	case "windows":
		runner, err := exec.Command("cmd", "/c", command).Output()
		if err == nil {
			return string(runner), nil
		} else {
			fmt.Println(err.Error())
			log.Fatal(err)
			return "", err
		}
	case "linux":
		fmt.Println("Linux")
	default:
		return "", errors.New("os not supported")
	}

	return "", nil
}
