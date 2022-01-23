package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
	go runBuild("git pull")
	go runBuild("pnpm i")
	go runBuild("npm run build:prod")
}
func main() {
	port := ":8080"
	log.Printf("Serving at: http://localhost%s\n", port)
	http.HandleFunc("/", http.StripPrefix("/", http.FileServer(http.Dir("web"))).ServeHTTP)
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
			fmt.Print(readString)
		}
	}()
	return c.Run()
}
