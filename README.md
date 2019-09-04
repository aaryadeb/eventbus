# EventBus
Minimal Event Bus written in Go

## Example
```go
package main

import (
	"fmt"
	"github.com/byliuyang/eventbus/eventbus"
	"time"
)

func main() {
	bus := eventbus.NewEventBus()

	notificationChannel := make(eventbus.DataChannel)
	notification := "notification"
	bus.Subscribe(notification, notificationChannel)

	go func() {
		for {
			select {
			case data := <-notificationChannel:
				fmt.Println(data)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	bus.Publish(notification, "Hello!")
	bus.UnSubscribe(notification, notificationChannel)
	time.Sleep(1 * time.Second)
}
```

## Author
Harry Liu - [byliuyang](https://github.com/byliuyang)

## License
This project is maintained under MIT license
