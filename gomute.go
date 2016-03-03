package gommute

import (
	"fmt"
	"net/url"
	"time"
)

const (
	Departure = "departure"
	Arrival   = "arrival"
)

type TimeReference struct {
	Time      time.Time
	Reference string
}

func NewTimeReference(t time.Time, ref string) *TimeReference {
	return &TimeReference{Time: t, Reference: ref}
}

type gomute struct {
	APIKey string
	From   Coordinates
	To     Coordinates
	Time   *TimeReference
}

func New(key string, f, t Coordinates) *gomute {
	return &gomute{
		APIKey: key,
		From:   f,
		To:     t,
	}
}

func (g *gomute) Journey() (*outJourney, error) {
	p := &url.Values{}
	p.Add("from", fmt.Sprintf("%s;%s", g.From.Longitude, g.From.Latitude))
	p.Add("to", fmt.Sprintf("%s;%s", g.To.Longitude, g.To.Latitude))

	if g.Time != nil {
		p.Add("datetime", g.Time.Time.Format("20060102T150405"))
		p.Add("datetime_represents", g.Time.Reference)
	}

	var j journeysResponse
	err := g.fetch("/journeys", p, &j)
	if err != nil {
		return nil, err
	}

	return j.normalize(), nil
}
