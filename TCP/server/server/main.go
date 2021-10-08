package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		//转成字符串
		recvStr := string(buf[:n])

		fmt.Println("收到client端发来的数据：", recvStr)
		//conn.Write([]byte(recvStr)) // 发送数据
	}
}

func senfInfo(conn net.Conn) {
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err := conn.Write([]byte(inputInfo)) // 发送数据
        if err != nil {
            return
        }
		fmt.Println("服务端发送数据",inputInfo)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")

	if err != nil {
		fmt.Println("出粗啦")
		return
	}
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	//	go senfInfo(conn)
	}
}
