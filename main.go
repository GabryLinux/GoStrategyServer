package main

import (
	//"GoServer/ActionProcedures/GameInteractions"
	//"fmt"
	"GoServer/ConnectionsListener"
	"log"
)


func main() {
	
	server := ConnectionListener.NewServer(":3000")

	//fmt.Println(gameinteractions.ProvaQUalcosa())
	log.Fatal(server.Start())
}
