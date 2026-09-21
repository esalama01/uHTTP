package main

import(
	//"golang.org/x/sys/unix"
	//"log"
	"uHTTP/src"
	"time"
)

func main(){
	/*
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM,0)
	if err != nil {
		log.Fatalf("Failed to create socket: %v", err)
	}
	defer unix.Close(fd)
	*/
	for {
		src.S_conn()
		time.Sleep(1 * time.Second)
	}
}
