package techplay

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const apiBaseEndpoint = "https://api.techplay.jp/v1"

// New func
func New(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("missing token")
	}
	return &Client{token: token}, nil
}

// Client struct
type Client struct {
	token string
}

// EventRanking struct
type EventRanking struct {
	UpdatedAt string       `json:"updated_at"`
	Ranking   []rawRanking `json:"ranking"`
}

// rawRanking struct
type rawRanking struct {
	Rank        int    `json:"rank"`
	EventID     int    `json:"event_id"`
	EventURL    string `json:"event_url"`
	Title       string `json:"title"`
	StartedAt   string `json:"started_at"`
	Description string `json:"description"`
	Capacity    int    `json:"capacity"`
	Entered     int    `json:"entered"`
	Station     string `json:"station"`
}

// GetEventRanking func
func (c *Client) GetEventRanking() (*EventRanking, error) {
	res, err := http.Get(fmt.Sprintf("%v/ranking/event?token=%v", apiBaseEndpoint, c.token))
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var eve *EventRanking
	err = json.Unmarshal(body, &eve)
	if err != nil {
		return nil, err
	}

	return eve, nil
}
