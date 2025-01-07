package dao

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	proto "github.com/listenGrey/lucianagRpcPKG/ask"
)

func Send(output string) error {
	conn, err := grpc.Dial(":8001", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect %v\n", err)
	}
	defer conn.Close()

	c := proto.NewRequestClient(conn)

	msg := &proto.Prompt{Cid: 0, Prompt: output}

	_, err = c.(proto.RequestClient).Request(context.Background(), msg)
	if err != nil {
		return fmt.Errorf("failed to receive %v\n", err)
	}

	return nil
}
