package acts_channel

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"maps"

	gonanoid "github.com/matoous/go-nanoid/v2"
	acts_grpc "github.com/yaojianpin/acts-channel-go/acts.grpc"
	"github.com/yaojianpin/acts-channel-go/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ActsChannel struct {
	conn *grpc.ClientConn
	stub acts_grpc.ActsServiceClient
}

func Connect(url string) (ActsChannel, error) {
	var stub acts_grpc.ActsServiceClient
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err == nil {
		log.Printf("%s connected\n", url)
		stub = acts_grpc.NewActsServiceClient(conn)
	}

	return ActsChannel{conn, stub}, err
}

func (c *ActsChannel) Send(name string, options any) (any, error) {
	log.Printf("Send %s %v\n", name, options)
	ctx := context.Background()
	var ret any
	var err error
	if data, err := json.Marshal(options); err == nil {
		if seq, err := gonanoid.New(); err == nil {
			var message = acts_grpc.Message{Seq: seq, Name: name, Data: data}
			if resp, err := c.stub.Send(ctx, &message); err == nil {
				err = json.Unmarshal(resp.Data, &ret)
				if err == nil {

				}
			}
		}
	}
	return ret, err
}

func (c *ActsChannel) Ack(id string) error {
	seq, err := gonanoid.New()
	if err == nil {
		ctx := context.Background()
		var message = acts_grpc.Message{Seq: seq, Name: "msg:ack", Ack: &id, Data: nil}
		c.stub.Send(ctx, &message)
	}
	return err
}

func (c *ActsChannel) Publish(pack any) (any, error) {
	return c.Send("pack:publish", pack)
}

func (c *ActsChannel) Deploy(model string, mid *string) (any, error) {
	log.Println("deploy")
	data := make(map[string]any)
	data["model"] = model
	data["mid"] = mid
	return c.Send("model:deploy", data)
}

func (c *ActsChannel) Start(id string, options *any) (any, error) {
	log.Println("Start")
	data := make(map[string]any)
	data["id"] = id
	if options != nil {
		maps.Copy(data, (*options).(map[string]any))
	}
	return c.Send("proc:start", data)
}

func (c *ActsChannel) Subscribe(clientid string, callback options.Callback, opts ...options.Options) error {
	log.Println("Subscribe")
	ctx := context.Background()
	var err error
	options := options.DefaultOptions()
	ack := true
	for _, opt := range opts {
		opt(&options)
	}

	messageOptions := acts_grpc.MessageOptions{Type: options.Type, State: options.State, Tag: options.Tag, Key: options.Key}
	if stream, err := c.stub.OnMessage(ctx, &messageOptions); err == nil {
		go func() {
			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Printf("recv error:%v", err)
					continue
				}
				if ack {
					c.Ack(resp.Seq)
				}
				var data any
				if err := json.Unmarshal(resp.Data, &data); err == nil {
					callback(data)
				}
			}
		}()
	}

	return err
}

func (c *ActsChannel) Act(name string, pid string, tid string, options *any) (any, error) {
	log.Println("Act:", name)
	data := make(map[string]any)
	data["pid"] = pid
	data["tid"] = tid
	if options != nil {
		maps.Copy(data, (*options).(map[string]any))
	}
	return c.Send(fmt.Sprintf("act:%s", name), data)
}

func (c *ActsChannel) Close() error {
	log.Println("Close")
	return c.conn.Close()
}
