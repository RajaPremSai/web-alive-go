package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string,port string)string{
	addr := destination+":"+port
	timeOut :=time.Duration(5*time.Second)
	conn,err:=net.DialTimeout("tcp",addr,timeOut)
	var status string

	if err!=nil{
		status=fmt.Sprintf("[DOWN] %v website is not reachable, Error : %v" ,destination,err)
	}else{
		status=fmt.Sprintf("[UP] %v is reachable,'n From : %v to %v",destination,conn.LocalAddr(),conn.RemoteAddr())
	}
	return status
}