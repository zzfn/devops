package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("调用成功"))
	go runBuild("git pull")
	go runBuild("go build")
}
func main() {
	readJson()
	httpServer()
}
func readJson() {
	filePtr, err := os.Open("./config.json")
	if err != nil {
		fmt.Println("文件打开失败 [Err:%s]", err.Error())
		return
	}
	defer func(filePtr *os.File) {
		err := filePtr.Close()
		if err != nil {
			return
		}
	}(filePtr)
	var info Demo
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		for p := range info.Projects {
			fmt.Println(info.Projects[p].Name)
			fmt.Println(info.Projects[p].Url)
		}
	}

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
