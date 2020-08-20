package eventbus_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/short-d/eventbus"
)

func TestAbs(t *testing.T) {
	bus := eventbus.NewEventBus()
	notificationChannel := make(eventbus.DataChannel)
	notificationChannel2 := make(eventbus.DataChannel)
	notification := "notification"
	bus.Subscribe(notification, notificationChannel)
	bus.Subscribe(notification, notificationChannel2)

	// messageEvent := "messageEvent"
	// // messageEventChannel := make(eventbus.DataChannel)
	// bus.Subscribe(messageEvent, notificationChannel)

	// fmt.Println(bus)

	var wg sync.WaitGroup
	wg.Add(2)

	// runSubscribers()
	go func() {
		// fmt.Println("starting subscriber")
		for {
			// fmt.Println("iinside for")
			select {
			case data := <-notificationChannel:
				// fmt.Println("data: chan1", data, len(notificationChannel))
				if data != nil {
					bus.UnSubscribe(notification, notificationChannel)
					// fmt.Println(bus)
					wg.Done()
					return
				} else {
					// fmt.Println("getting nil data from eventBus")
				}
			}
		}
	}()

	go func() {
		// fmt.Println("starting subscriber")
		for {
			select {
			case data := <-notificationChannel2:
				// fmt.Println("data: chan2 ", data, len(notificationChannel2))
				if data != nil {
					bus.UnSubscribe(notification, notificationChannel2)
					wg.Done()
					return
				} else {
					// fmt.Println("getting nil data from eventBus")
				}
			}
		}
	}()

	go func() {
		fmt.Println("Publishing")
		// time.Sleep(2 * time.Second)
		bus.Publish(notification, "Hello!")
		fmt.Println("done publishing")
	}()
	wg.Wait()

	// fmt.Println(bus)
}
