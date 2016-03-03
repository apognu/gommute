package gommute

import (
	"fmt"
	"strings"
	"time"
)

type Coordinates struct {
	Longitude string `json:"lon"`
	Latitude  string `json:"lat"`
}

type OutJourney struct {
	Duration int          `json:"duration"`
	Section  []OutSection `json:"sections"`
}

type OutSection struct {
	Mode            string       `json:"mode"`
	Name            string       `json:"name,omitempty"`
	Color           string       `json:"color,omitempty"`
	Direction       string       `json:"direction,omitempty"`
	From            string       `json:"from,omitempty"`
	FromCoordinates *Coordinates `json:"from_coordinates,omitempty"`
	To              string       `json:"to,omitempty"`
	ToCoordinates   *Coordinates `json:"to_coordinates,omitempty"`
	DepartureTime   time.Time    `json:"departure_time"`
	ArrivalTime     time.Time    `json:"arrival_time"`
	Duration        int          `json:"duration"`
}

type JourneysResponse struct {
	Journeys []Journey `json:"journeys"`
}

func (j JourneysResponse) Normalize() *OutJourney {
	if len(j.Journeys) == 0 {
		return nil
	}

	journey := j.Journeys[0]
	out := &OutJourney{Duration: journey.Duration}
	for _, s := range journey.Sections {
		if s.Type == "waiting" {
			continue
		}

		dep, _ := time.Parse("20060102T150405", s.DepartureTime)
		arr, _ := time.Parse("20060102T150405", s.ArrivalTime)

		sec := OutSection{
			FromCoordinates: s.From.Info.Coordinates,
			ToCoordinates:   s.To.Info.Coordinates,
			DepartureTime:   dep,
			ArrivalTime:     arr,
			Duration:        s.Duration,
		}

		if s.From.Address != nil {
			sec.FromCoordinates = s.From.Address.Coordinates
		}
		if s.To.Address != nil {
			sec.ToCoordinates = s.To.Address.Coordinates
		}

		if s.Type == "street_network" {
			sec.Mode = "walking"
			sec.From = s.From.Name
			sec.To = s.To.Name
		} else if s.Type == "public_transport" {
			sec.Mode = strings.ToLower(s.DisplayInformation.Mode)
			sec.Name = s.DisplayInformation.Label
			sec.Color = fmt.Sprintf("#%s", s.DisplayInformation.Color)
			sec.Direction = s.DisplayInformation.Direction
			sec.From = s.From.Info.Name
			sec.To = s.To.Info.Name
		} else if s.Type == "transfer" {
			sec.Mode = "transfer"
			sec.FromCoordinates = nil
			sec.ToCoordinates = nil
		}

		out.Section = append(out.Section, sec)
	}

	return out
}

type Journey struct {
	Sections []Section `json:"sections"`
	Duration int       `json:"duration"`
}

type Section struct {
	Type               string             `json:"type"`
	DisplayInformation DisplayInformation `json:"display_informations,omitempty"`
	From               StopPoint          `json:"from"`
	To                 StopPoint          `json:"to"`
	DepartureTime      string             `json:"departure_date_time"`
	ArrivalTime        string             `json:"arrival_date_time"`
	Duration           int                `json:"duration"`
}

type DisplayInformation struct {
	Mode      string `json:"commercial_mode"`
	Label     string `json:"label"`
	Direction string `json:"direction"`
	Color     string `json:"color"`
}

type StopPoint struct {
	Type    string        `json:"embedded_type"`
	Name    string        `json:"name"`
	Info    StopPointInfo `json:"stop_point,omitempty"`
	Address *Address      `json:"address,omitempty"`
}

type StopPointInfo struct {
	Name        string       `json:"name"`
	Coordinates *Coordinates `json:"coord,omitempty"`
}

type Address struct {
	Coordinates *Coordinates `json:"coord"`
}
