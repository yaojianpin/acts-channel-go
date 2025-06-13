# acts-channel

Acts-channel is a workflow client to connect to [`acts-server`](https://github.com/yaojianpin/acts-server)

# Usage

Before connecting, please download [`acts-server`](https://github.com/yaojianpin/acts-server) and start it

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
	acts_channel "github.com/yaojianpin/acts-channel-go"
	"github.com/yaojianpin/acts-channel-go/options"
)

func main() {
	var model = `
    id: test
    steps:
        - name: step 1
          id: step1
          acts:
              - uses: acts.core.irq
                key: abc
    `

	if client, err := acts_channel.Connect("127.0.0.1:10080"); err == nil {
		defer func() {
			client.Close()
		}()

        // deploy a workflow model
		if ret, err := client.Deploy(model, nil); err == nil {
			fmt.Printf("deploy=%v\n", ret)
			if ret, err := client.Send("model:get", map[string]any{"id": "test", "fmt": "tree"}); err == nil {
				fmt.Printf("models=%v\n", ret)
			}

            // subscribe the messages from server
			client.Subscribe("client-1", func(message any) {
				fmt.Printf("message: %v\n", message)
				var data options.Message
				mapstructure.Decode(message, &data)
				fmt.Printf("data: %v\n", data)
				if data.Key == "abc" && data.State == "created" {

                    // execute act from client
					client.Act("complete", data.Pid, data.Tid, nil)
				}

			})

            // start a workflow by model id
			if pid, err := client.Start("test", nil); err == nil {
				fmt.Printf("%s\n", pid)
			}
		}
	}
	time.Sleep(3 * time.Second)
}

```
