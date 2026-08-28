package src

import(
	"net"
	"log"
	"fmt"
)

func StartServer(){
	for{//an infinite loop to listen to any upcoming connection at anytime
		listener, err := net.Listen("tcp", "127.0.1.1:8080") // localhost <=> my laptop acting as my own server
		if err != nil{
			log.Fatal(err)
		}
		defer listener.Close() //ensures the port gets freed when the function returns
		for {
			client, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
				continue
			}
			HandleConnection(client)
		}
	}
}

func HandleConnection(conn net.Conn){
	fmt.Printf("Redirecting into %s\n", conn.RemoteAddr().String())
	defer conn.Close()

}