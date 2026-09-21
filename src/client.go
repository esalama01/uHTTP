package src

import(
	"net"
	"fmt"
)

func C_conn() {
	conn, err := net.Dial("tcp", "localhost:8989")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("Hello, server"))
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println(string(buffer))
}