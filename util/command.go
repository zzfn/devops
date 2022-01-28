package util

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

func Run(command string) {
	err := shell(command)
	if err != nil {
		fmt.Print(err)
	}
}
func shell(command string) error {
	args := strings.Fields(command)
	c := exec.Command(args[0], args[1:]...)
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	go func() {
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				break
			}
			log.Print(readString)
			fmt.Print(readString)
		}
	}()
	return c.Run()
}
