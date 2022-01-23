package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type Project struct {
	Project string
	Url     string
	Info    Info
}
type Info struct {
	Title       string
	description string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
	go runBuild("git pull")
	go runBuild("pnpm i")
	go runBuild("npm run build:prod")
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
	defer filePtr.Close()
	var info []Project
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		for p := range info {
			fmt.Println(info[p].Project)
			fmt.Println(info[p].Url)
			fmt.Println(info[p].Info.Title)
		}
	}

}
func readYaml() {
	var project Project //定义一个结构体变量

	//读取yaml文件到缓存中
	config, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Print(err)
	}
	err1 := yaml.Unmarshal(config, &project)
	if err1 != nil {
		fmt.Println("error")
	}
	//for p := range project {
	//	fmt.Println(project[p].Project)
	//	fmt.Println(project[p].Url)
	//}
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
			fmt.Print(readString)
		}
	}()
	return c.Run()
}
