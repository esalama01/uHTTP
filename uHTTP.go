package main

import(
	"exec"
	"os"
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
	serverCmd := exec.Command("go", "run", "server.go")
	serverCmd.Stdout = os.Stdout
	serverCmd.Stderr = os.Stderr
	
	src.S_conn()
	time.Sleep(1 * time.Second)
	
}
