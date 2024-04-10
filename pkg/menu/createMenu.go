package menu

import (
	"fmt"
	"menucli/pkg/cursor"
	"os/exec"
	"strings"
)

func CreateMenu(lines []string) error {
	menu := cursor.NewMenu("\n\nSelect Connection: [ \"q\" to quit] \n\n")
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