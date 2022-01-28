package util

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
)

func Send() {
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("root"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	sshClient, err := ssh.Dial("tcp", "127.0.0.1:22", config)
	if err != nil {
		log.Fatal("创建ssh client 失败", err)
	}
	defer sshClient.Close()
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer sftpClient.Close()
	cwd, err := sftpClient.Getwd()
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("当前目录：", cwd)
	fi, err := sftpClient.Lstat(cwd)
	log.Println(fi)

	remoteFileName := "index.html"
	remoteFile, err := sftpClient.Create(sftp.Join(cwd, remoteFileName))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer remoteFile.Close()

	localFileName := "web/index.html"
	//打开本地文件file.dat
	localFile, err := os.Open(localFileName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer localFile.Close()

	//本地文件流拷贝到上传文件流
	n, err := io.Copy(remoteFile, localFile)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//获取本地文件大小
	localFileInfo, err := os.Stat(localFileName)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("文件上传成功[%s->%s]本地文件大小：%s，上传文件大小：%s", localFileName, remoteFileName, formatFileSize(localFileInfo.Size()), formatFileSize(n))
}
func formatFileSize(s int64) (size string) {
	if s < 1024 {
		return fmt.Sprintf("%.2fB", float64(s)/float64(1))
	} else if s < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(s)/float64(1024))
	} else if s < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(s)/float64(1024*1024))
	} else if s < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(s)/float64(1024*1024*1024))
	} else if s < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(s)/float64(1024*1024*1024*1024))
	} else { //if s < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(s)/float64(1024*1024*1024*1024*1024))
	}
}
