package gommute

import (
	"fmt"
	"net/url"
)

type gomute struct {
	APIKey string
	From   Coordinates
	To     Coordinates
}

func New(key string, f, t Coordinates) *gomute {
	return &gomute{
		APIKey: key,
		From:   f,
		To:     t,
	}
}

func (g *gomute) Journey() (*OutJourney, error) {
	p := &url.Values{}
	p.Add("from", fmt.Sprintf("%s;%s", g.From.Longitude, g.From.Latitude))
	p.Add("to", fmt.Sprintf("%s;%s", g.To.Longitude, g.To.Latitude))

	var j JourneysResponse
	err := g.fetch("/journeys", p, &j)
	if err != nil {
		return nil, err
	}

	return j.Normalize(), nil
}
