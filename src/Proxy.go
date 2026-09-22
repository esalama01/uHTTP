package src

import(
	"net"
	"fmt"
	"bufio"
	"io"
)

func S_conn(){
	listener, err := net.Listen("tcp", ":8989")
	if err != nil{
		fmt.Println("Error Listening: ",err)
		return
	}
	defer listener.Close()
	fmt.Println("Server Running On Port 8989..")
	for {
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("Error Accepting: ",err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()
	addr := conn.RemoteAddr()
	addr_type := addr.Network()
	addr_name := addr.String()
	C_conn(addr_type, addr_name)
	/*
	reader := bufio.NewReader(conn)
	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil{
			if err != io.EOF {
				fmt.Println("Failed to read the data.")
			}
			return
		}
		fmt.Printf("request: %s", bytes)
		line := fmt.Sprintf("Echo: %s", bytes)
		fmt.Printf("response: %s", line)
		_, err = conn.Write([]byte(line))
		if err != nil{
			fmt.Println("Failed to write the data.")
			return
		}
	}
	*/
}

func C_conn(a_type string, a_name string) {
	conn, err := net.Dial(a_type, a_name)
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