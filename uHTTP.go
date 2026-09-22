package main

import(
	//"exec"
	"uHTTP/src"
	//"time"
)

func main(){
	/*
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM,0)
	if err != nil {
		log.Fatalf("Failed to create socket: %v", err)
	}
	defer unix.Close(fd)
	*/
	src.S_conn()
}
