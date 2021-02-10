package main

import (
	"github.com/PierreKieffer/mongoStream/dataModel"
	"github.com/PierreKieffer/mongoStream/services"
	"log"
)

var exit = make(chan bool)

func main() {

	mongoUri := "mongodb://localhost:27017"

	// Init oplog buffer channel
	var buffer = services.InitBuffer(10)

	// Start consumer
	go Consumer(buffer)

	// Start producers
	go services.ListenerWorker(mongoUri, "database", "collection1", buffer)
	go services.ListenerWorker(mongoUri, "database", "collection2", buffer)

	<-exit
}

func Consumer(logBuffer chan dataModel.Oplog) {
	for {
		log.Println("reveived data : ", <-logBuffer)
	}

}
