package fanout

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"
)

//taggingDispatcher implement our dispatcher interface
type taggingDispatcher struct {
	Address string
	//	stream  proto.StreamClient
	conn *grpc.ClientConn
}
type messageContent struct {
	content  string
	priority int
}

func TestComplexStreamingFanOut(t *testing.T) {

	builder := func() IDispatcher {
		return &taggingDispatcher{Address: "127.0.0.2"}
	}
	tagging := &Tagging{
		topic:    "new topic",
		pipeline: NewPipeline(builder, 2, true),
	}

	tagging.pipeline.Start(context.Background())

	tagging.pipeline.Dispatch(messageContent{"all,please stay home", 1000})

}

type Tagging struct {
	topic    string
	pipeline *Pipeline
}

func (d *taggingDispatcher) Before(ctx context.Context) error {

	fmt.Println("i'm doing somthing before processing")

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
	fmt.Println("i'm doing somthing After processing")
	return nil
}

func (d *taggingDispatcher) Process(msg interface{}) error {

	content := msg.(messageContent)
	fmt.Println("i'm doing processing,with conten", content)
	return nil
}
