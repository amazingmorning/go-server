package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:9898")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		_, err = conn.Write([]byte("hello go-server!"))
		if err != nil {
			fmt.Sprintf("Write error:%s", err)
		}

		buf := make([]byte, 1024)

		count, err := conn.Read(buf)

		if err != nil {
			fmt.Sprintf("Read error :%s", err)
		}

		fmt.Printf(" server call back : %s, cnt = %d\n", buf, count)

		time.Sleep(1 * time.Second)

	}

}
