package fanout

import (
	"context"
	"testing"

	"google.golang.org/grpc"
)

//taggingDispatcher implement our dispatcher interface
type taggingDispatcher struct {
	Address string
	//	stream  proto.StreamClient
	conn *grpc.ClientConn
}
type messageContent struct{}

func TestComplexStreamingFanOut(t *testing.T) {

	builder := func() IDispatcher {
		return &taggingDispatcher{Address: "SH"}
	}
	tagging := &Tagging{
		topic:    "new topic",
		pipeline: NewPipeline(builder, 2, true),
	}
	tagging.pipeline.Dispatch(messageContent{})

	tagging.pipeline.Start(context.Background())
}

type Tagging struct {
	topic    string
	pipeline *Pipeline
}

func (d *taggingDispatcher) Before(ctx context.Context) error {
	conn, err := grpc.Dial(d.Address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	d.conn = conn
	// // //	client := proto.NewClient(conn)
	// // 	stream, err := client.StreamMetric(ctx)
	// // 	if err != nil {
	// // 		return err
	// // 	}
	// // 	d.stream = stream
	return nil
}

func (d *taggingDispatcher) After() error {
	// _, err := d.stream.CloseAndRecv()
	// e := d.conn.Close()
	// if e != nil {
	// 	log.Error("close connection error", field.Error(e))
	// }
	//return err
	return nil
}

func (d *taggingDispatcher) Process(msg interface{}) error {
	//return d.stream.Send(msg.(*proto.Tagging))
	return nil
}
