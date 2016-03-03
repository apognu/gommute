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

type outJourney struct {
	Duration int          `json:"duration"`
	Section  []outSection `json:"sections"`
}

type outSection struct {
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

type journeysResponse struct {
	Journeys []journey `json:"journeys"`
}

func (j journeysResponse) normalize() *outJourney {
	if len(j.Journeys) == 0 {
		return nil
	}

	journey := j.Journeys[0]
	out := &outJourney{Duration: journey.Duration}
	for _, s := range journey.Sections {
		if s.Type == "waiting" {
			continue
		}

		dep, _ := time.Parse("20060102T150405", s.DepartureTime)
		arr, _ := time.Parse("20060102T150405", s.ArrivalTime)

		sec := outSection{
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
			sec.From = normalizeName(s.From.Name)
			sec.To = normalizeName(s.To.Name)
		} else if s.Type == "public_transport" {
			sec.Mode = strings.ToLower(s.DisplayInformation.Mode)
			sec.Name = s.DisplayInformation.Label
			sec.Color = fmt.Sprintf("#%s", s.DisplayInformation.Color)
			sec.Direction = normalizeName(s.DisplayInformation.Direction)
			sec.From = normalizeName(s.From.Info.Name)
			sec.To = normalizeName(s.To.Info.Name)
		} else if s.Type == "transfer" {
			sec.Mode = "transfer"
			sec.FromCoordinates = nil
			sec.ToCoordinates = nil
		}

		out.Section = append(out.Section, sec)
	}

	return out
}

func normalizeName(s string) string {
	return strings.Title(strings.ToLower(strings.TrimSpace(strings.Split(s, "(")[0])))
}

type journey struct {
	Sections []section `json:"sections"`
	Duration int       `json:"duration"`
}

type section struct {
	Type               string             `json:"type"`
	DisplayInformation displayInformation `json:"display_informations,omitempty"`
	From               stopPoint          `json:"from"`
	To                 stopPoint          `json:"to"`
	DepartureTime      string             `json:"departure_date_time"`
	ArrivalTime        string             `json:"arrival_date_time"`
	Duration           int                `json:"duration"`
}

type displayInformation struct {
	Mode      string `json:"commercial_mode"`
	Label     string `json:"label"`
	Direction string `json:"direction"`
	Color     string `json:"color"`
}

type stopPoint struct {
	Type    string        `json:"embedded_type"`
	Name    string        `json:"name"`
	Info    stopPointInfo `json:"stop_point,omitempty"`
	Address *address      `json:"address,omitempty"`
}

type stopPointInfo struct {
	Name        string       `json:"name"`
	Coordinates *Coordinates `json:"coord,omitempty"`
}

type address struct {
	Coordinates *Coordinates `json:"coord"`
}
