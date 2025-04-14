package mock

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

type Platform struct {
	name     string
	client   pulsar.Client
	producer pulsar.Producer
	campign  *Campign
	ctx      context.Context
}

// init the ad platform
func (p *Platform) initialize() error {

	// init client
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://pulsar:6650",
	})

	if err != nil {
		return errors.New(fmt.Sprintf("Error in initializing pulsar client for %s. Error: %v\n", p.name, err))
	}
	p.client = client

	// init producer
	topicName := "persistent://public/default/" + p.name
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topicName,
	})

	if err != nil {
		return errors.New(fmt.Sprintf("Error in initializing producer for %s. Error: %v\n", p.name, err))
	}
	p.producer = producer

	return nil
}

// create and send view event to pulsar
func (p *Platform) pushViewEvent() {
	p.campign.NoOfViews++
	p.campign.DailyBudgetRemaining = p.campign.DailyBudgetRemaining - p.campign.Cpv
	cpv := strconv.Itoa(p.campign.Cpv)

	// generate event
	event := ViewEvent{
		EventType: "views",
		UserID:    p.campign.UserId,
		CampignID: p.campign.CampignId,
		Cpv:       cpv,
		SessionID: "Sess_123",
		AdID:      "10",
		PageURL:   "https://abc.com",
		Timestamp: time.Now().Format("02/01.2006"),
		Device:    "pc",
	}

	eventInbytes, _ := json.Marshal(event)
	p.producer.Send(p.ctx, &pulsar.ProducerMessage{Payload: eventInbytes})
}

// create and send click event to pulsar
func (p *Platform) pushClickEvent() {
	p.campign.NoOfClicks++

	// generate event
	event := ViewEvent{
		EventType: "clicks",
		UserID:    p.campign.UserId,
		CampignID: p.campign.CampignId,
		SessionID: "Sess_123",
		AdID:      "10",
		PageURL:   "https://abc.com",
		Timestamp: time.Now().Format("02/01.2006"),
		Device:    "pc",
	}

	eventInbytes, _ := json.Marshal(event)
	p.producer.Send(p.ctx, &pulsar.ProducerMessage{Payload: eventInbytes})
}

// create and send conversion event to pulsar
func (p *Platform) pushConversionEvent() {
	p.campign.NoOfConversions++

	// generate event
	event := ViewEvent{
		EventType: "conversions",
		UserID:    p.campign.UserId,
		CampignID: p.campign.CampignId,
		SessionID: "Sess_123",
		AdID:      "10",
		PageURL:   "https://abc.com",
		Timestamp: time.Now().Format("02/01.2006"),
		Device:    "pc",
	}

	eventInbytes, _ := json.Marshal(event)
	p.producer.Send(p.ctx, &pulsar.ProducerMessage{Payload: eventInbytes})

}
