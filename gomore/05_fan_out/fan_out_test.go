package fanout

import "concurrency"

type taggingDispatcher struct {
	Address string
	stream  proto.Havilah_StreamMetricClient
	conn    *grpc.ClientConn
}

func (d *taggingDispatcher) Before(ctx context.Context) error {
	conn, err := grpc.Dial(d.Address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	d.conn = conn
	client := proto.NewHavilahClient(conn)

	stream, err := client.StreamMetric(ctx)
	if err != nil {
		return err
	}
	d.stream = stream
	return nil
}

func (d *taggingDispatcher) After() error {
	_, err := d.stream.CloseAndRecv()

	e := d.conn.Close()
	if e != nil {
		log.Error("close havilah connection error", field.Error(e))
	}
	return err
}

func (d *taggingDispatcher) Process(msg interface{}) error {
	return d.stream.Send(msg.(*proto.Tagging))
}


tagging := &Tagging{
    topic: topic,
    pipeline: concurrency.NewPipeline(func() concurrency.Dispatcher {
        return &taggingDispatcher{Address: address}
    }, ch, idle, debug),
}
tagging.pipeline.Start()

func main(){
	tagging.pipeline.Dispatch(youStruct{})
}
