package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	MYWebSocket "vtoken_digiccy_go/MyWebSocket/proto/MyWebSocket"
)

type MyWebSocket struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *MyWebSocket) MyWebSocketCall(ctx context.Context, req *MYWebSocket.Request, rsp *MYWebSocket.Response) error {
	log.Log("Received MyWebSocket.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *MyWebSocket) Stream(ctx context.Context, req *MYWebSocket.StreamingRequest, stream MYWebSocket.MyWebSocket_StreamStream) error {
	log.Logf("Received MyWebSocket.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&MYWebSocket.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *MyWebSocket) PingPong(ctx context.Context, stream MYWebSocket.MyWebSocket_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&MYWebSocket.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
