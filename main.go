package main

import (
	"bufio"
	"github.com/spf13/cobra"
	"log"
	"menucli/pkg/context"
	"menucli/pkg/menu"
	"os"
)

const config = "/usr/local/bin/.ctxgo.config"


func readLines() ([]string, error) {

	file, err := os.Open(config)
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


func main() {
    var rootCmd = &cobra.Command{
        Use:   "ctxgo",
		Version: "v2.0",
        Short: "GCP context switcher",
        Long:  "ctxgo delivers a menu selection to switch between GCP contexts.",
        Run: func(cmd *cobra.Command, args []string) {
            lines, err := readLines()
            if err != nil {
                log.Fatalf("readLines: %s", err)
            }
            err = menu.CreateMenu(lines)
            if err != nil {
                log.Fatalf("Error: Access issue %s", err)
            }
            err = context.ConfirmContext()
            if err != nil {
                log.Fatalf("confirmContext: %s", err)
            }
        },
    }
	
    rootCmd.Execute()
}