package main

import (
	"fmt"
	"github.com/alexsuslov/messages"
	"log"
)

func main() {
	M:=messages.New("msg")
	err:=M.ReLoad("templates", ".tpl")
	if err!= nil{
		panic(err)
	}
	err=fmt.Errorf("error message")
	msg := M.Execute("error.tpl", struct{Err error; PID string}{err, "PID1"})
	log.Print(msg)

	M.SetDebugger(true)
	M.Debug("error_debug.tpl", struct{Err error; PID string}{err, "PID1"})
	M.SetDebugger(true)
	M.Debug("debug", struct{Err error; PID string}{err, "PID1"})
}