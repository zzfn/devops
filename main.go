package main

import (
	"bufio"
	"devops/router"
	"devops/ssh"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("调用成功"))
	go runBuild("git pull")
	go runBuild("go build")
}
func cmd(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(ssh.Cmd(r.FormValue("cmd"))))
}
func main() {
	router.Router()
	http.ListenAndServe(":8080", nil)
}

func httpServer() {
	port := ":8080"
	log.Printf("Serving at: http://localhost%s\n", port)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("web"))))
	http.HandleFunc("/test", indexHandler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe fail:", err)
	}
}
func runBuild(command string) {
	err := Command(command)
	if err != nil {
		fmt.Print(err)
	}
}
func Command(command string) error {
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
