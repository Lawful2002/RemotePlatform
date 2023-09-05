package docker

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"
	"systemcaller"
	"time"
)

func isDockerDaemonRunning() bool {
	// Input: nil
	// Output: bool -> whether the docker daemon is running
	log.SetPrefix("docker daemon: ")
	output, err := systemcaller.RunSystemCommand("ps")
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		return strings.Contains(output, "not running")
	}
}

func wasImageBuilt(dump string) bool {
	log.SetPrefix("docker build: ")
	return !strings.Contains(dump, "failed")
}

func isDockerContainerRunning(dump string) bool {
	// Input: dump string -> output recieved from the command runner
	// Output: bool -> whether the docker daemon is running
	log.SetPrefix("docker container: ")
	return !strings.Contains(dump, "Error response from daemon")
}

func buildDockerImage(filename string) (string, error) {
	// Input: filename string -> name of the Dockerfile to be built
	// Output: (string, error) -> name of the image and the error (if any)
	log.SetPrefix("docker build: ")
	name := fmt.Sprintf("%d-%d", time.Now().Nanosecond(), rand.Intn(math.MaxInt32))
	if !isDockerDaemonRunning() {
		return "", errors.New("daemon not running")
	}

	command := "build -t " + name + " ./Dockerfiles/Dockerfile-" + filename
	out, err := systemcaller.RunSystemCommand(command)
	if err != nil || !wasImageBuilt(out) {
		log.Fatal(err)
		return "", err
	}
	log.Print(out)
	return name, nil
}

func runDockerImage(name string) error {
	// Input: name string -> name of the image to be run
	// Output: error -> errors (if any)
	log.SetPrefix("docker run: ")
	if !isDockerDaemonRunning() {
		log.Fatal("daemon not running")
		return errors.New("docker daemon is not running")
	}
	command := "run -d --rm --name " + name
	out, err := systemcaller.RunSystemCommand(command)
	if err != nil || !isDockerContainerRunning(out) {
		log.Fatal(err)
		return errors.New("error encountered while running the dockerfile")
	}
	fmt.Println(out)
	log.Print(out)
	return nil
}

func RunDockerfile(filename string) error {
	// Input: name string -> name of the dockefile to be executed
	// Output: error -> errors (if any)
	image, err := buildDockerImage(filename)
	if err != nil {
		log.Fatal(err)
		return errors.New(err.Error())
	}
	runnerErr := runDockerImage(image)
	if runnerErr != nil {
		log.Fatal(runnerErr)
		return errors.New(runnerErr.Error())
	}

	return nil
}
