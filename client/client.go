package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Event TODO import from events service rather than duplicate code
type Event struct {
	ID          int64
	UserID      int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
}

// Events service
type Events struct {
	timeout time.Duration
	url     string
}

func NewEvents() Events {
	return Events{
		timeout: time.Duration(1) * time.Second,
		url:     "http://localhost:8080/events/%d",
	}
}

func (service Events) GetEvent(eventId int) (*Event, error) {
	log.Printf("Fetching event with id: [%v]", eventId)
	log.Printf("Fetching event with id2: [%v]", eventId)
	c := http.Client{Timeout: service.timeout}
	url := fmt.Sprintf(service.url, eventId)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Printf("error creating request %s", err)
		return nil, err
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)

	if err != nil {
		log.Printf("error on event service request %s", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("error on closing request body %s", err)
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		var event Event
		err = json.Unmarshal(body, &event)
		if err != nil {
			log.Printf("error parsing json response %s", err)
			return nil, err
		}
		log.Printf("event found:%v", event)
		return &event, nil

	}
	return nil, fmt.Errorf("event could not be verified, reponse: %v", resp.StatusCode)
}
