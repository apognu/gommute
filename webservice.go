package gommute

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const endpoint = "http://api.navitia.io/v1"

func (g *gomute) fetch(uri string, p *url.Values, obj interface{}) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s?%s", endpoint, uri, p.Encode()), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", g.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(obj)
	if err != nil {
		return err
	}

	return nil
}
