package engine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	redis "github.com/redis/go-redis/v9"
)

// one pulsar client per campign not per platform
type Platform struct {
	name      string
	ctx       context.Context
	pclient   pulsar.Client
	pconsumer pulsar.Consumer
	rclient   *redis.Client
}

func (p *Platform) initialize() error {
	// create pulsar client
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: PulsarHostName,
	})

	if err != nil {
		return errors.New(fmt.Sprintf("Error in initializing pulsar client for %s. Error: %v\n", p.name, err))
	}
	p.pclient = client

	// subscribe to the topic
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            PulsarTopicPrefix + p.name,
		SubscriptionName: "my-app",
		Type:             pulsar.Exclusive,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("Error in initializing consumer for %s. Error: %v\n", p.name, err))
	}
	p.pconsumer = consumer

	// create a redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:       RedisHostName,
		ClientName: "kavin-Redis",
		DB:         0,
	})
	p.rclient = rdb

	payload := RedisEvent{
		Platform:      p.name,
		UserID:        UserId,
		CampignID:     CampignID,
		Views:         0,
		Clicks:        0,
		Conversions:   0,
		Lasteventtime: time.Now().Format("02/01.2006"),
	}

	var fields map[string]interface{}

	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &fields)
	if err != nil {
		log.Fatal(err)
	}

	key := fmt.Sprintf("campaign:%s:%s:%s", payload.UserID, payload.CampignID, payload.Platform)
	fmt.Println("Redis Key: ", key)
	if err := rdb.HSet(p.ctx, key, fields).Err(); err != nil {
		return errors.New(fmt.Sprintf("Error in initializing redis client for %s. Error: %v\n", p.name, err))
	}
	return nil
}
