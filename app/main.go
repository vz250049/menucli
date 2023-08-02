package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const path = ".goctx"

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func createMenu(lines []string) error {
	menu := gocliselect.NewMenu("Select Connection")
	for _, line := range lines {
		stringSlice := strings.Split(line, ", ")
		menu.AddItem(stringSlice[0], stringSlice[1])
	}
	choice := menu.Display()

	arg := strings.Split(choice, " ")

	c, b := exec.Command("gcloud", arg[1], arg[2], arg[3], arg[4], arg[5], arg[6], arg[7], arg[8]), new(strings.Builder)
	c.Stdout = b
	err := c.Run()
	if err != nil {
		return err
	}
	fmt.Println(b.String())
	return nil
}

func confirmContext() error {

	c, b := exec.Command("kubectl", "config", "current-context"), new(strings.Builder)
	c.Stdout = b
	err := c.Run()
	if err != nil {
		return err
	}
	fmt.Println("Connected to: ", b.String())
	return nil
}

func main() {

	lines, err := readLines(path)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	err = createMenu(lines)
	if err != nil {
		log.Fatalf("setContext: %s", err)
	}
	err = confirmContext()
	if err != nil {
		log.Fatalf("confirmContext: %s", err)
	}
}
