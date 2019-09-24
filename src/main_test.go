package main

import (
	"log"
	"os/exec"
	"testing"
)

func Test_main(t *testing.T) {

	expected :=
		`1	apple
2	banana
3	orange
4	watermelon
`

	cmd := exec.Command("/bin/sh", "-c", "cat ./fruits.txt | go run main.go")
	rawOutput, err := cmd.Output()
	output := string(rawOutput)

	if err != nil {
		log.Fatal(err)
	}

	if output != expected {
		t.Errorf("Expecting: %s, receive %v", expected, output)
	}

}

func Test_mainWithoutArguments(t *testing.T) {

	expected := "No pipe data was received\n"

	cmd := exec.Command("/bin/sh", "-c", "go run main.go")
	rawOutput, err := cmd.Output()
	output := string(rawOutput)

	if err != nil {
		log.Fatal(err)
	}

	if output != expected {
		t.Errorf("Expecting: %s, receive %v", expected, output)
	}

}
