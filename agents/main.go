package main

import (
	"log"
	"postTap/communicator"
)

var queryComm *communicator.AmqpComm

func init() {
	queryComm = new(communicator.AmqpComm)
}

func main() {

	go WaitForCommand()
	initNode := &stap{scriptPath: "./stp_scripts/exec_init_node.stp", pid: 0, timeout: 0}
	initNode.Run()
}

func WaitForCommand() {
	commandQueue := new(communicator.AmqpComm)
	if err := commandQueue.Connect("amqp://guest:guest@localhost:5672"); err != nil {
		log.Fatalf("%s", err)
		return
	}
	defer commandQueue.Close()
	commandProcessor := new(Command)
	commandProcessor.RunningStp = map[int]*stap{}
	commandQueue.Receive("command", commandProcessor)
}