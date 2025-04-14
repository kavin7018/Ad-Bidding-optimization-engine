package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"personal-documents/zocket/app/healthmetrics"
	"strconv"
	"sync"
)

func Start(ctx context.Context, exitChan chan<- bool) {
	fmt.Println("Starting application engine...")

	AdPlatforms := []string{"google", "facebook", "instagram"}

	var wg sync.WaitGroup
	for _, platform := range AdPlatforms {
		wg.Add(1)
		go process(&wg, platform, ctx, exitChan)
	}

	defer wg.Wait()
}

func process(wg *sync.WaitGroup, platform string, ctx context.Context, exitChan chan<- bool) {
	defer wg.Done()

	// initialize the client for platform
	p := Platform{name: platform, ctx: ctx}
	if err := p.initialize(); err != nil {
		fmt.Printf("Error in initializing clients for mock server. Err: %v", err)
		exitChan <- true
	}
	defer p.pconsumer.Close()

	// creating a worker pool to process the messages
	// var msgChan chan string
	// for i := 1; i <= 3; i++ {
	// 	wg.Add(1)
	// 	go processMsg(msgChan, wg)
	// }

	// consumer messages
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Shutdown signal received. Exiting consumer loop.")
			return
		default:
			msg, err := p.pconsumer.Receive(ctx)
			if err != nil {
				fmt.Println("Error receiving message:", err)
				continue
			}
			fmt.Println("Msg: ", string(msg.Payload()))
			p.processMsg(msg.Payload())
			p.pconsumer.Ack(msg)
		}
	}

}

// func processMsg(msgChan chan string, wg *sync.WaitGroup) {
func (p *Platform) processMsg(msg []byte) {
	baseEvent := make(map[string]string)
	if err := json.Unmarshal(msg, &baseEvent); err != nil {
		log.Println("Error unmarshalling message:", err)
		return
	}

	platform := p.name
	eventType := baseEvent["event_type"]
	healthmetrics.PlatformActivity.WithLabelValues(platform, eventType).Inc()

	val, _ := strconv.ParseFloat(baseEvent["cpv"], 64)
	if eventType == "views" {
		healthmetrics.Bid.WithLabelValues(platform).Set(val)
	}

	fmt.Printf("Received a %s event for platform %s\n", baseEvent, p.name)

	key := "campaign" + ":" + UserId + ":" + CampignID + ":" + p.name
	fmt.Printf("Incrementing %s event for key %s\n", eventType, key)
	_, err := p.rclient.HIncrBy(p.ctx, key, eventType, 1).Result()
	if err != nil {
		fmt.Printf("Error in updating the counter in redis. Err: %v", err)
		return
	}
}
