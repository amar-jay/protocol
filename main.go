package main

import (
	"github.com/livekit/protocol/livekit"
	"google.golang.org/protobuf/proto"
)

func main() {
	_, err := proto.Marshal(&livekit.WebhookEvent{})
	if err != nil {
		println(err)
	}
	//--- NOTHING HERE ---
}
