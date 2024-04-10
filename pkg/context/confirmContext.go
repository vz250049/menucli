package context

import (
	"fmt"
	"os/exec"
	"strings"
)

func ConfirmContext() error {
	c, b := exec.Command("kubectl", "config", "current-context"), new(strings.Builder)
	c.Stdout = b
	err := c.Run()
	if err != nil {
		return err
	}
	if c.ProcessState.ExitCode() != 0 {
		return fmt.Errorf("error: %s", b.String())
	}
	
	fmt.Println("\033[1mConnected to:\033[0m", b.String())
	return nil
}
