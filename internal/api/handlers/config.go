package handlers

import (
	"fmt"

	"github.com/jsnfwlr/vexil/internal/db"
	"github.com/r3labs/sse/v2"
)

type Server struct {
	DBClient *db.Client
	EventSrv *sse.Server
}

func New(dbClient *db.Client, enableSSE bool) Server {
	var eventSrv *sse.Server
	if enableSSE {
		eventSrv = sse.NewWithCallback(AddSub, RemSub)
		eventSrv.AutoReplay = false
		// @TODO - use the dbClient to query the db for all the environments, and create streams for each one

		eventSrv.CreateStream("dev")
	}

	return Server{
		DBClient: dbClient,
		EventSrv: eventSrv,
	}
}

func AddSub(streamId string, sub *sse.Subscriber) {
	fmt.Printf("Subscriber added to stream %s\n", streamId)
}

func RemSub(streamId string, sub *sse.Subscriber) {
	fmt.Printf("Subscriber left stream %s\n", streamId)
}
